import os
import pathlib
import platform
import sys
import tarfile

from invoke import Context, task
from invoke.exceptions import Exit

from tasks.libs.common.color import Color, color_message
from tasks.libs.common.git import get_commit_sha, get_main_parent_commit
from tasks.libs.common.utils import collapsed_section, get_distro

PROFILE_COV = "coverage.out"
TMP_PROFILE_COV_PREFIX = "coverage.out.rerun"
GO_COV_TEST_PATH = "test_with_coverage"
COV_ARCHIVE_NAME = f"coverage_{get_distro()}.tgz"
AWS_CMD = "aws.cmd" if sys.platform == 'win32' else "aws"
BUCKET_CI_VAR = "S3_PERMANENT_ARTIFACTS_URI"


class CodecovWorkaround:
    """
    The CodecovWorkaround class wraps the gotestsum cmd execution to fix codecov reports inaccuracy,
    according to https://github.com/gotestyourself/gotestsum/issues/274 workaround.
    Basically unit tests' reruns rewrite the whole coverage file, making it inaccurate.
    We use the --raw-command flag to tell each `go test` iteration to write coverage in a different file.
    """

    def __init__(self, ctx: Context, module_path: str, coverage: bool, packages: str, args: dict[str, str]):
        self.ctx = ctx
        self.module_path = module_path
        self.coverage = coverage
        self.packages = packages
        self.args = args
        self.cov_test_path_sh = os.path.join(self.module_path, GO_COV_TEST_PATH) + ".sh"
        self.cov_test_path_ps1 = os.path.join(self.module_path, GO_COV_TEST_PATH) + ".ps1"
        self.call_ps1_from_bat = os.path.join(self.module_path, GO_COV_TEST_PATH) + ".bat"
        self.cov_test_path = self.cov_test_path_sh if platform.system() != 'Windows' else self.cov_test_path_ps1

    def __enter__(self):
        coverage_script = ""
        if self.coverage:
            if platform.system() == 'Windows':
                coverage_script = f"""$tempFile = (".\\{TMP_PROFILE_COV_PREFIX}." + ([guid]::NewGuid().ToString().Replace("-", "").Substring(0, 10)))
go test $($args | select -skip 1) -json -coverprofile="$tempFile" {self.packages}
exit $LASTEXITCODE
"""
            else:
                coverage_script = f"""#!/usr/bin/env bash
set -eu
go test "${{@:2}}" -json -coverprofile=\"$(mktemp {TMP_PROFILE_COV_PREFIX}.XXXXXXXXXX)\" {self.packages}
"""
            with open(self.cov_test_path, 'w', encoding='utf-8') as f:
                f.write(coverage_script)

            with open(self.call_ps1_from_bat, 'w', encoding='utf-8') as f:
                f.write(
                    f"""@echo off
powershell.exe -executionpolicy Bypass -file {GO_COV_TEST_PATH}.ps1 %*"""
                )

            os.chmod(self.cov_test_path, 0o755)
            os.chmod(self.call_ps1_from_bat, 0o755)

        return self.cov_test_path_sh if platform.system() != 'Windows' else self.call_ps1_from_bat

    def __exit__(self, *_):
        if self.coverage:
            # Removing the coverage script.
            try:
                os.remove(self.cov_test_path)
                os.remove(self.call_ps1_from_bat)
            except FileNotFoundError:
                print(
                    f"Error: Could not find the coverage script {self.cov_test_path} or {self.call_ps1_from_bat} while trying to delete it.",
                    file=sys.stderr,
                )
            # Merging the unit tests reruns coverage files, keeping only the merged file.
            files_to_delete = [
                os.path.join(self.module_path, f)
                for f in os.listdir(self.module_path)
                if f.startswith(f"{TMP_PROFILE_COV_PREFIX}.")
            ]
            if not files_to_delete:
                print(
                    f"Error: Could not find coverage files starting with '{TMP_PROFILE_COV_PREFIX}.' in {self.module_path}",
                    file=sys.stderr,
                )
            else:
                self.ctx.run(
                    f"gocovmerge {' '.join(files_to_delete)} > \"{os.path.join(self.module_path, PROFILE_COV)}\""
                )
                for f in files_to_delete:
                    os.remove(f)


@task
def codecov(
    ctx: Context,
    pull_coverage_cache: bool = False,
    push_coverage_cache: bool = False,
    debug: bool = False,
):
    """
    Uploads coverage data of all modules to Codecov.
    This expects that the coverage files have already been generated by
    inv test --coverage.

    Flags:   --pull-coverage-cache: [For dev branches] Pull the coverage cache from main parent commit.
             --push-coverage-cache: [For main]         Push the coverage cache to the S3 bucket.
             --debug:                                  Print debug information.
    """
    if pull_coverage_cache and push_coverage_cache:
        raise Exit(
            color_message("Error: Can't use both --pull-missing-coverage and --push-coverage-cache flags.", "red"),
            code=1,
        )
    distro_tag = get_distro()
    codecov_binary = "codecov" if platform.system() != "Windows" else "codecov.exe"
    with collapsed_section("Upload coverage reports to Codecov"):
        if pull_coverage_cache:
            apply_missing_coverage(ctx, from_commit_sha=get_main_parent_commit(ctx), debug=debug)
        if push_coverage_cache:
            upload_coverage_to_s3(ctx)
        ctx.run(f"{codecov_binary} -f {PROFILE_COV} -F {distro_tag}", warn=True)


def produce_coverage_tar(files, archive_name):
    """
    Produce a tgz file containing all coverage files.
    """
    with tarfile.open(archive_name, "w:gz") as tgz:
        for f in files:
            tgz.add(f)
    print(color_message(f"Successfully created {archive_name}", Color.GREEN))


def _get_coverage_cache_uri():
    if BUCKET_CI_VAR not in os.environ:
        raise Exit(color_message(f"Error: the {BUCKET_CI_VAR} environment variable is not set.", "red"), code=1)
    return f"{os.environ[BUCKET_CI_VAR]}/coverage-cache"


@task
def upload_coverage_to_s3(ctx: Context):
    """
    Create an archive with all the coverage.out files from the inv test --coverage command.
    Then upload the archive to the dd-ci-persistent-artefacts-build-stable S3 bucket.
    """

    # Find all coverage files in the project and put them in a tgz archive
    cov_files = sorted(pathlib.Path(".").rglob(PROFILE_COV))
    produce_coverage_tar(cov_files, COV_ARCHIVE_NAME)

    # Upload the archive to S3
    cache_uri = _get_coverage_cache_uri()
    commit_sha = os.getenv("CI_COMMIT_SHA") or get_commit_sha(ctx)
    if ctx.run(f"{AWS_CMD} s3 cp {COV_ARCHIVE_NAME} {cache_uri}/{commit_sha}/", echo=True, warn=True):
        print(
            color_message(
                f'Successfully uploaded coverage cache to {cache_uri}/{commit_sha}/{COV_ARCHIVE_NAME}', Color.GREEN
            )
        )
    else:
        raise Exit(
            color_message(f"Failed to upload coverage cache to {cache_uri}/{commit_sha}/{COV_ARCHIVE_NAME}", "red"),
            code=1,
        )

    # Remove the local archive
    os.remove(COV_ARCHIVE_NAME)
    print(color_message(f'Successfully removed the local {COV_ARCHIVE_NAME}', Color.GREEN))


@task
def apply_missing_coverage(ctx: Context, from_commit_sha: str, debug: bool = False):
    """
    Download the coverage cache archive from S3 for the given commit SHA
    and extract it to the right folders.

    :param from_commit_sha: The commit SHA from which to restore the coverage cache. It needs at least the 8 first characters.
    """
    if not from_commit_sha or len(from_commit_sha) < 8:
        raise Exit(color_message("Error: the commit SHA is missing or invalid.", "red"), code=1)

    # Download the coverage archive from S3
    cache_uri = _get_coverage_cache_uri()
    cache_key = f"{cache_uri}/{from_commit_sha}/{COV_ARCHIVE_NAME}"
    downloaded_archive = f"coverage_{from_commit_sha[:8]}.tgz"
    if ctx.run(f"{AWS_CMD} s3 cp {cache_key} ./{downloaded_archive}", echo=True, warn=True):
        print(color_message(f'Successfully retrieved coverage cache from commit {from_commit_sha}', Color.GREEN))
    else:
        raise Exit(color_message(f'Failed to restore coverage cache from {cache_key}', "red"), code=1)

    # Extract only the missing coverage files from the archive
    with tarfile.open(f"{downloaded_archive}", "r:gz") as tgz:
        curr_cov_files = [str(p) for p in pathlib.Path(".").rglob(PROFILE_COV)]
        main_cov_files = tgz.getnames()
        missing_cov_files = list(set(main_cov_files) - set(curr_cov_files))
        if debug:
            print(f"Current coverage files: {sorted(curr_cov_files)}")
            print(f"Main coverage files: {main_cov_files}")
            print(f"Missing coverage files: {sorted(missing_cov_files)}")
        missing_members = [f for f in tgz.getmembers() if f.name in missing_cov_files]
        tgz.extractall(path='.', members=missing_members)

    # Remove the local archive
    print(color_message(f'Successfully extracted coverage cache from {downloaded_archive}', Color.GREEN))
    os.remove(downloaded_archive)
