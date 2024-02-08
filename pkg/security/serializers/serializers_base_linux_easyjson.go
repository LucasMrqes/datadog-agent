//go:build linux
// +build linux

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package serializers

import (
	json "encoding/json"
	utils "github.com/DataDog/datadog-agent/pkg/security/utils"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers(in *jlexer.Lexer, out *ProcessContextSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.ProcessSerializer = new(ProcessSerializer)
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "parent":
			if in.IsNull() {
				in.Skip()
				out.Parent = nil
			} else {
				if out.Parent == nil {
					out.Parent = new(ProcessSerializer)
				}
				(*out.Parent).UnmarshalEasyJSON(in)
			}
		case "ancestors":
			if in.IsNull() {
				in.Skip()
				out.Ancestors = nil
			} else {
				in.Delim('[')
				if out.Ancestors == nil {
					if !in.IsDelim(']') {
						out.Ancestors = make([]*ProcessSerializer, 0, 8)
					} else {
						out.Ancestors = []*ProcessSerializer{}
					}
				} else {
					out.Ancestors = (out.Ancestors)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ProcessSerializer
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ProcessSerializer)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Ancestors = append(out.Ancestors, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pid":
			out.Pid = uint32(in.Uint32())
		case "ppid":
			if in.IsNull() {
				in.Skip()
				out.PPid = nil
			} else {
				if out.PPid == nil {
					out.PPid = new(uint32)
				}
				*out.PPid = uint32(in.Uint32())
			}
		case "tid":
			out.Tid = uint32(in.Uint32())
		case "uid":
			out.UID = int(in.Int())
		case "gid":
			out.GID = int(in.Int())
		case "user":
			out.User = string(in.String())
		case "group":
			out.Group = string(in.String())
		case "path_resolution_error":
			out.PathResolutionError = string(in.String())
		case "comm":
			out.Comm = string(in.String())
		case "tty":
			out.TTY = string(in.String())
		case "fork_time":
			if in.IsNull() {
				in.Skip()
				out.ForkTime = nil
			} else {
				if out.ForkTime == nil {
					out.ForkTime = new(utils.EasyjsonTime)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ForkTime).UnmarshalJSON(data))
				}
			}
		case "exec_time":
			if in.IsNull() {
				in.Skip()
				out.ExecTime = nil
			} else {
				if out.ExecTime == nil {
					out.ExecTime = new(utils.EasyjsonTime)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ExecTime).UnmarshalJSON(data))
				}
			}
		case "exit_time":
			if in.IsNull() {
				in.Skip()
				out.ExitTime = nil
			} else {
				if out.ExitTime == nil {
					out.ExitTime = new(utils.EasyjsonTime)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ExitTime).UnmarshalJSON(data))
				}
			}
		case "credentials":
			if in.IsNull() {
				in.Skip()
				out.Credentials = nil
			} else {
				if out.Credentials == nil {
					out.Credentials = new(ProcessCredentialsSerializer)
				}
				(*out.Credentials).UnmarshalEasyJSON(in)
			}
		case "user_session":
			if in.IsNull() {
				in.Skip()
				out.UserSession = nil
			} else {
				if out.UserSession == nil {
					out.UserSession = new(UserSessionContextSerializer)
				}
				(*out.UserSession).UnmarshalEasyJSON(in)
			}
		case "executable":
			if in.IsNull() {
				in.Skip()
				out.Executable = nil
			} else {
				if out.Executable == nil {
					out.Executable = new(FileSerializer)
				}
				(*out.Executable).UnmarshalEasyJSON(in)
			}
		case "interpreter":
			if in.IsNull() {
				in.Skip()
				out.Interpreter = nil
			} else {
				if out.Interpreter == nil {
					out.Interpreter = new(FileSerializer)
				}
				(*out.Interpreter).UnmarshalEasyJSON(in)
			}
		case "container":
			if in.IsNull() {
				in.Skip()
				out.Container = nil
			} else {
				if out.Container == nil {
					out.Container = new(ContainerContextSerializer)
				}
				(*out.Container).UnmarshalEasyJSON(in)
			}
		case "argv0":
			out.Argv0 = string(in.String())
		case "args":
			if in.IsNull() {
				in.Skip()
				out.Args = nil
			} else {
				in.Delim('[')
				if out.Args == nil {
					if !in.IsDelim(']') {
						out.Args = make([]string, 0, 4)
					} else {
						out.Args = []string{}
					}
				} else {
					out.Args = (out.Args)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Args = append(out.Args, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "args_truncated":
			out.ArgsTruncated = bool(in.Bool())
		case "envs":
			if in.IsNull() {
				in.Skip()
				out.Envs = nil
			} else {
				in.Delim('[')
				if out.Envs == nil {
					if !in.IsDelim(']') {
						out.Envs = make([]string, 0, 4)
					} else {
						out.Envs = []string{}
					}
				} else {
					out.Envs = (out.Envs)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Envs = append(out.Envs, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "envs_truncated":
			out.EnvsTruncated = bool(in.Bool())
		case "is_thread":
			out.IsThread = bool(in.Bool())
		case "is_kworker":
			out.IsKworker = bool(in.Bool())
		case "is_exec_child":
			out.IsExecChild = bool(in.Bool())
		case "source":
			out.Source = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers(out *jwriter.Writer, in ProcessContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Parent != nil {
		const prefix string = ",\"parent\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Parent).MarshalEasyJSON(out)
	}
	if len(in.Ancestors) != 0 {
		const prefix string = ",\"ancestors\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v4, v5 := range in.Ancestors {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					(*v5).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Pid != 0 {
		const prefix string = ",\"pid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Pid))
	}
	if in.PPid != nil {
		const prefix string = ",\"ppid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(*in.PPid))
	}
	if in.Tid != 0 {
		const prefix string = ",\"tid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Tid))
	}
	{
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.UID))
	}
	{
		const prefix string = ",\"gid\":"
		out.RawString(prefix)
		out.Int(int(in.GID))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		out.String(string(in.Group))
	}
	if in.PathResolutionError != "" {
		const prefix string = ",\"path_resolution_error\":"
		out.RawString(prefix)
		out.String(string(in.PathResolutionError))
	}
	if in.Comm != "" {
		const prefix string = ",\"comm\":"
		out.RawString(prefix)
		out.String(string(in.Comm))
	}
	if in.TTY != "" {
		const prefix string = ",\"tty\":"
		out.RawString(prefix)
		out.String(string(in.TTY))
	}
	if in.ForkTime != nil {
		const prefix string = ",\"fork_time\":"
		out.RawString(prefix)
		(*in.ForkTime).MarshalEasyJSON(out)
	}
	if in.ExecTime != nil {
		const prefix string = ",\"exec_time\":"
		out.RawString(prefix)
		(*in.ExecTime).MarshalEasyJSON(out)
	}
	if in.ExitTime != nil {
		const prefix string = ",\"exit_time\":"
		out.RawString(prefix)
		(*in.ExitTime).MarshalEasyJSON(out)
	}
	if in.Credentials != nil {
		const prefix string = ",\"credentials\":"
		out.RawString(prefix)
		(*in.Credentials).MarshalEasyJSON(out)
	}
	if in.UserSession != nil {
		const prefix string = ",\"user_session\":"
		out.RawString(prefix)
		(*in.UserSession).MarshalEasyJSON(out)
	}
	if in.Executable != nil {
		const prefix string = ",\"executable\":"
		out.RawString(prefix)
		(*in.Executable).MarshalEasyJSON(out)
	}
	if in.Interpreter != nil {
		const prefix string = ",\"interpreter\":"
		out.RawString(prefix)
		(*in.Interpreter).MarshalEasyJSON(out)
	}
	if in.Container != nil {
		const prefix string = ",\"container\":"
		out.RawString(prefix)
		(*in.Container).MarshalEasyJSON(out)
	}
	if in.Argv0 != "" {
		const prefix string = ",\"argv0\":"
		out.RawString(prefix)
		out.String(string(in.Argv0))
	}
	if len(in.Args) != 0 {
		const prefix string = ",\"args\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v6, v7 := range in.Args {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if in.ArgsTruncated {
		const prefix string = ",\"args_truncated\":"
		out.RawString(prefix)
		out.Bool(bool(in.ArgsTruncated))
	}
	if len(in.Envs) != 0 {
		const prefix string = ",\"envs\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.Envs {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.EnvsTruncated {
		const prefix string = ",\"envs_truncated\":"
		out.RawString(prefix)
		out.Bool(bool(in.EnvsTruncated))
	}
	if in.IsThread {
		const prefix string = ",\"is_thread\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsThread))
	}
	if in.IsKworker {
		const prefix string = ",\"is_kworker\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsKworker))
	}
	if in.IsExecChild {
		const prefix string = ",\"is_exec_child\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsExecChild))
	}
	if in.Source != "" {
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		out.String(string(in.Source))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProcessContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProcessContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers1(in *jlexer.Lexer, out *NetworkContextSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "device":
			if in.IsNull() {
				in.Skip()
				out.Device = nil
			} else {
				if out.Device == nil {
					out.Device = new(NetworkDeviceSerializer)
				}
				(*out.Device).UnmarshalEasyJSON(in)
			}
		case "l3_protocol":
			out.L3Protocol = string(in.String())
		case "l4_protocol":
			out.L4Protocol = string(in.String())
		case "source":
			(out.Source).UnmarshalEasyJSON(in)
		case "destination":
			(out.Destination).UnmarshalEasyJSON(in)
		case "size":
			out.Size = uint32(in.Uint32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers1(out *jwriter.Writer, in NetworkContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Device != nil {
		const prefix string = ",\"device\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Device).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"l3_protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.L3Protocol))
	}
	{
		const prefix string = ",\"l4_protocol\":"
		out.RawString(prefix)
		out.String(string(in.L4Protocol))
	}
	{
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		(in.Source).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"destination\":"
		out.RawString(prefix)
		(in.Destination).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Size))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NetworkContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NetworkContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers1(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers2(in *jlexer.Lexer, out *MatchedRuleSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Tags = append(out.Tags, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "policy_name":
			out.PolicyName = string(in.String())
		case "policy_version":
			out.PolicyVersion = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers2(out *jwriter.Writer, in MatchedRuleSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	if len(in.Tags) != 0 {
		const prefix string = ",\"tags\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Tags {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if in.PolicyName != "" {
		const prefix string = ",\"policy_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PolicyName))
	}
	if in.PolicyVersion != "" {
		const prefix string = ",\"policy_version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PolicyVersion))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MatchedRuleSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MatchedRuleSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers2(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers3(in *jlexer.Lexer, out *IPPortSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ip":
			out.IP = string(in.String())
		case "port":
			out.Port = uint16(in.Uint16())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers3(out *jwriter.Writer, in IPPortSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ip\":"
		out.RawString(prefix[1:])
		out.String(string(in.IP))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Port))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IPPortSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IPPortSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers3(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers4(in *jlexer.Lexer, out *IPPortFamilySerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "family":
			out.Family = string(in.String())
		case "ip":
			out.IP = string(in.String())
		case "port":
			out.Port = uint16(in.Uint16())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers4(out *jwriter.Writer, in IPPortFamilySerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"family\":"
		out.RawString(prefix[1:])
		out.String(string(in.Family))
	}
	{
		const prefix string = ",\"ip\":"
		out.RawString(prefix)
		out.String(string(in.IP))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Port))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IPPortFamilySerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IPPortFamilySerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers4(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers5(in *jlexer.Lexer, out *ExitEventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cause":
			out.Cause = string(in.String())
		case "code":
			out.Code = uint32(in.Uint32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers5(out *jwriter.Writer, in ExitEventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cause\":"
		out.RawString(prefix[1:])
		out.String(string(in.Cause))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Code))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ExitEventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ExitEventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers5(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers6(in *jlexer.Lexer, out *EventContextSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "outcome":
			out.Outcome = string(in.String())
		case "async":
			out.Async = bool(in.Bool())
		case "matched_rules":
			if in.IsNull() {
				in.Skip()
				out.MatchedRules = nil
			} else {
				in.Delim('[')
				if out.MatchedRules == nil {
					if !in.IsDelim(']') {
						out.MatchedRules = make([]MatchedRuleSerializer, 0, 0)
					} else {
						out.MatchedRules = []MatchedRuleSerializer{}
					}
				} else {
					out.MatchedRules = (out.MatchedRules)[:0]
				}
				for !in.IsDelim(']') {
					var v13 MatchedRuleSerializer
					(v13).UnmarshalEasyJSON(in)
					out.MatchedRules = append(out.MatchedRules, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "origin":
			out.Origin = string(in.String())
		case "suppressed":
			out.Suppressed = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers6(out *jwriter.Writer, in EventContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Category != "" {
		const prefix string = ",\"category\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Category))
	}
	if in.Outcome != "" {
		const prefix string = ",\"outcome\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Outcome))
	}
	if in.Async {
		const prefix string = ",\"async\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Async))
	}
	if len(in.MatchedRules) != 0 {
		const prefix string = ",\"matched_rules\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.MatchedRules {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.Origin != "" {
		const prefix string = ",\"origin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Origin))
	}
	if in.Suppressed {
		const prefix string = ",\"suppressed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Suppressed))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers6(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers6(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers7(in *jlexer.Lexer, out *DNSQuestionSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "class":
			out.Class = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "size":
			out.Size = uint16(in.Uint16())
		case "count":
			out.Count = uint16(in.Uint16())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers7(out *jwriter.Writer, in DNSQuestionSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"class\":"
		out.RawString(prefix[1:])
		out.String(string(in.Class))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Size))
	}
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Count))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DNSQuestionSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers7(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DNSQuestionSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers7(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers8(in *jlexer.Lexer, out *DNSEventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint16(in.Uint16())
		case "question":
			(out.Question).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers8(out *jwriter.Writer, in DNSEventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint16(uint16(in.ID))
	}
	{
		const prefix string = ",\"question\":"
		out.RawString(prefix)
		(in.Question).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DNSEventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers8(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DNSEventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers8(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers9(in *jlexer.Lexer, out *DDContextSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "span_id":
			out.SpanID = uint64(in.Uint64())
		case "trace_id":
			out.TraceID = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers9(out *jwriter.Writer, in DDContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SpanID != 0 {
		const prefix string = ",\"span_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.SpanID))
	}
	if in.TraceID != 0 {
		const prefix string = ",\"trace_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.TraceID))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DDContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers9(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DDContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers9(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers10(in *jlexer.Lexer, out *ContainerContextSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "created_at":
			if in.IsNull() {
				in.Skip()
				out.CreatedAt = nil
			} else {
				if out.CreatedAt == nil {
					out.CreatedAt = new(utils.EasyjsonTime)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreatedAt).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers10(out *jwriter.Writer, in ContainerContextSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.CreatedAt != nil {
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.CreatedAt).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerContextSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers10(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerContextSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers10(l, v)
}
func easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers11(in *jlexer.Lexer, out *BaseEventSerializer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.FileEventSerializer = new(FileEventSerializer)
	out.ExitEventSerializer = new(ExitEventSerializer)
	out.ProcessContextSerializer = new(ProcessContextSerializer)
	out.ContainerContextSerializer = new(ContainerContextSerializer)
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "evt":
			(out.EventContextSerializer).UnmarshalEasyJSON(in)
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "file":
			if in.IsNull() {
				in.Skip()
				out.FileEventSerializer = nil
			} else {
				if out.FileEventSerializer == nil {
					out.FileEventSerializer = new(FileEventSerializer)
				}
				(*out.FileEventSerializer).UnmarshalEasyJSON(in)
			}
		case "exit":
			if in.IsNull() {
				in.Skip()
				out.ExitEventSerializer = nil
			} else {
				if out.ExitEventSerializer == nil {
					out.ExitEventSerializer = new(ExitEventSerializer)
				}
				(*out.ExitEventSerializer).UnmarshalEasyJSON(in)
			}
		case "process":
			if in.IsNull() {
				in.Skip()
				out.ProcessContextSerializer = nil
			} else {
				if out.ProcessContextSerializer == nil {
					out.ProcessContextSerializer = new(ProcessContextSerializer)
				}
				(*out.ProcessContextSerializer).UnmarshalEasyJSON(in)
			}
		case "container":
			if in.IsNull() {
				in.Skip()
				out.ContainerContextSerializer = nil
			} else {
				if out.ContainerContextSerializer == nil {
					out.ContainerContextSerializer = new(ContainerContextSerializer)
				}
				(*out.ContainerContextSerializer).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers11(out *jwriter.Writer, in BaseEventSerializer) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"evt\":"
		first = false
		out.RawString(prefix[1:])
		(in.EventContextSerializer).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Date).MarshalEasyJSON(out)
	}
	if in.FileEventSerializer != nil {
		const prefix string = ",\"file\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.FileEventSerializer).MarshalEasyJSON(out)
	}
	if in.ExitEventSerializer != nil {
		const prefix string = ",\"exit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ExitEventSerializer).MarshalEasyJSON(out)
	}
	if in.ProcessContextSerializer != nil {
		const prefix string = ",\"process\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ProcessContextSerializer).MarshalEasyJSON(out)
	}
	if in.ContainerContextSerializer != nil {
		const prefix string = ",\"container\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ContainerContextSerializer).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseEventSerializer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1e47abeEncodeGithubComDataDogDatadogAgentPkgSecuritySerializers11(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseEventSerializer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1e47abeDecodeGithubComDataDogDatadogAgentPkgSecuritySerializers11(l, v)
}
