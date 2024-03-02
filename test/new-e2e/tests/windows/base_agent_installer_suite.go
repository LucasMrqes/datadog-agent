// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package windows contains the code to run the e2e tests on Windows
package windows

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/components"
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/e2e"
	"github.com/DataDog/datadog-agent/test/new-e2e/pkg/runner"
	platformCommon "github.com/DataDog/datadog-agent/test/new-e2e/tests/agent-platform/common"
	windowsAgent "github.com/DataDog/datadog-agent/test/new-e2e/tests/windows/common/agent"

	"testing"
)

// AgentInstallerSuite is the interface for the Windows Agent installer suites
type AgentInstallerSuite[Env any] interface {
	e2e.Suite[Env]

	// GetStackNamePart returns a string that will be included in the stack name.
	// Use this to differentiate the stacks when running multiple suites in parallel.
	GetStackNamePart() (string, error)

	// GetAgentPackage returns the Agent package to install.
	GetAgentPackage() (*windowsAgent.Package, error)
}

// BaseAgentInstallerSuite is a base class for the Windows Agent installer suites
type BaseAgentInstallerSuite[Env any] struct {
	e2e.BaseSuite[Env]

	AgentPackage *windowsAgent.Package
	OutputDir    string
}

// InstallAgent installs the Agent on a given Windows host. It will pass all the parameters to the MSI installer.
func (b *BaseAgentInstallerSuite[Env]) InstallAgent(host *components.RemoteHost, options ...windowsAgent.InstallAgentOption) (string, error) {
	b.T().Helper()
	opts := []windowsAgent.InstallAgentOption{
		windowsAgent.WithInstallLogFile(filepath.Join(b.OutputDir, "install.log")),
	}
	opts = append(opts, options...)
	return windowsAgent.InstallAgent(host, opts...)
}

// NewTestClientForHost creates a new TestClient for a given host.
func (b *BaseAgentInstallerSuite[Env]) NewTestClientForHost(host *components.RemoteHost) *platformCommon.TestClient {
	// We could bring the code from NewWindowsTestClient here
	return platformCommon.NewWindowsTestClient(b.T(), host)
}

// BeforeTest overrides the base BeforeTest to perform some additional per-test setup like configuring the output directory.
func (b *BaseAgentInstallerSuite[Env]) BeforeTest(suiteName, testName string) {
	b.BaseSuite.BeforeTest(suiteName, testName)

	var err error
	b.OutputDir, err = runner.GetTestOutputDir(runner.GetProfile(), b.T())
	if err != nil {
		b.T().Fatalf("should get output dir")
	}
	b.T().Logf("Output dir: %s", b.OutputDir)
}

// SetupSuite overrides the base SetupSuite to perform some additional setups like setting the package to install.
func (b *BaseAgentInstallerSuite[Env]) SetupSuite() {
	b.BaseSuite.SetupSuite()

	var err error
	b.AgentPackage, err = b.GetAgentPackage()
	if err != nil {
		b.T().Fatal(err)
	}
	b.T().Logf("Using Agent: %#v", b.AgentPackage)
}

// GetAgentPackage returns the Agent package to install.
// This method is called automatically by SetupSuite, and only needs to be called explicitly
// if you need to get the package before SetupSuite is called.
func (b *BaseAgentInstallerSuite[Env]) GetAgentPackage() (*windowsAgent.Package, error) {
	if b.AgentPackage == nil {
		var err error
		b.AgentPackage, err = windowsAgent.GetPackageFromEnv()
		if err != nil {
			return nil, fmt.Errorf("failed to get MSI URL from env: %w", err)
		}
	}

	return b.AgentPackage, nil
}

// GetStackNamePart implements AgentInstallerSuite, returning an empty string by default.
func (b *BaseAgentInstallerSuite[Env]) GetStackNamePart() (string, error) {
	return "", nil
}

// Run runs an AgentInstallerSuite test suite.
// It extends e2e.Run by setting a default stack name and including the Agent major version in the test name.
// These help deconflict stacks in parallel tests and differentiate tests in the junit reports.
func Run[Env any, T AgentInstallerSuite[Env]](t *testing.T, s T, options ...e2e.SuiteOption) {
	s.SetT(t)

	opts := []e2e.SuiteOption{}

	// default stack name. This will be overridden if the WithStackName option is provided.
	stackName, err := getDefaultStackName(s)
	if err != nil {
		t.Fatalf("failed to get stack name: %v", err)
	}
	opts = append(opts, e2e.WithStackName(stackName))

	// give precedence to provided options
	opts = append(opts, options...)

	// Include the agent major version in the test name so junit reports will differentiate
	// Agent 6 and Agent 7 tests run by the CI.
	majorVersion, err := getAgentMajorVersion(s)
	if err != nil {
		t.Fatal(err)
	}
	testName := fmt.Sprintf("Windows Agent v%s", majorVersion)

	t.Run(testName, func(t *testing.T) {
		e2e.Run(t, s, opts...)
	})
}

// getdefaultStackName returns the stack name for the given AgentInstallerSuite,
// including information to differentiate the stacks betweenn jobs.
func getDefaultStackName[Env any, T AgentInstallerSuite[Env]](s T) (string, error) {
	agentPackage, err := s.GetAgentPackage()
	if err != nil {
		return "", err
	}
	majorVersion, err := getAgentMajorVersion(s)
	if err != nil {
		return "", err
	}

	suitePart, err := s.GetStackNamePart()
	if err != nil {
		return "", err
	}
	if suitePart != "" {
		suitePart = fmt.Sprintf("-%s", suitePart)
	}

	// E2E auto includes the pipeline ID in the stack name, so we don't need to do that here.
	stackName := fmt.Sprintf("windows-msi-test%s-v%s-%s", suitePart, majorVersion, agentPackage.Arch)

	// If running in CI, append the CI job ID to the stack name to ensure uniqueness between jobs
	ciJobID := os.Getenv("CI_JOB_ID")
	if ciJobID != "" {
		stackName = fmt.Sprintf("%s-%s", stackName, ciJobID)
	}

	return stackName, nil
}

// GetAgentMajorVersion returns the major version of the Agent package.
func getAgentMajorVersion[Env any, T AgentInstallerSuite[Env]](s T) (string, error) {
	agentPackage, err := s.GetAgentPackage()
	if err != nil {
		return "", err
	}
	return strings.Split(agentPackage.Version, ".")[0], nil
}
