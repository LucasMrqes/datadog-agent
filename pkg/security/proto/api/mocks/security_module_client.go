// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/DataDog/datadog-agent/pkg/security/proto/api"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// SecurityModuleClient is an autogenerated mock type for the SecurityModuleClient type
type SecurityModuleClient struct {
	mock.Mock
}

// DumpActivity provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) DumpActivity(ctx context.Context, in *api.ActivityDumpParams, opts ...grpc.CallOption) (*api.ActivityDumpMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DumpActivity")
	}

	var r0 *api.ActivityDumpMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpParams, ...grpc.CallOption) (*api.ActivityDumpMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpParams, ...grpc.CallOption) *api.ActivityDumpMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ActivityDumpMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.ActivityDumpParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DumpDiscarders provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) DumpDiscarders(ctx context.Context, in *api.DumpDiscardersParams, opts ...grpc.CallOption) (*api.DumpDiscardersMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DumpDiscarders")
	}

	var r0 *api.DumpDiscardersMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.DumpDiscardersParams, ...grpc.CallOption) (*api.DumpDiscardersMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.DumpDiscardersParams, ...grpc.CallOption) *api.DumpDiscardersMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.DumpDiscardersMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.DumpDiscardersParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DumpNetworkNamespace provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) DumpNetworkNamespace(ctx context.Context, in *api.DumpNetworkNamespaceParams, opts ...grpc.CallOption) (*api.DumpNetworkNamespaceMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DumpNetworkNamespace")
	}

	var r0 *api.DumpNetworkNamespaceMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.DumpNetworkNamespaceParams, ...grpc.CallOption) (*api.DumpNetworkNamespaceMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.DumpNetworkNamespaceParams, ...grpc.CallOption) *api.DumpNetworkNamespaceMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.DumpNetworkNamespaceMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.DumpNetworkNamespaceParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DumpProcessCache provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) DumpProcessCache(ctx context.Context, in *api.DumpProcessCacheParams, opts ...grpc.CallOption) (*api.SecurityDumpProcessCacheMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DumpProcessCache")
	}

	var r0 *api.SecurityDumpProcessCacheMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.DumpProcessCacheParams, ...grpc.CallOption) (*api.SecurityDumpProcessCacheMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.DumpProcessCacheParams, ...grpc.CallOption) *api.SecurityDumpProcessCacheMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.SecurityDumpProcessCacheMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.DumpProcessCacheParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActivityDumpStream provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) GetActivityDumpStream(ctx context.Context, in *api.ActivityDumpStreamParams, opts ...grpc.CallOption) (api.SecurityModule_GetActivityDumpStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetActivityDumpStream")
	}

	var r0 api.SecurityModule_GetActivityDumpStreamClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpStreamParams, ...grpc.CallOption) (api.SecurityModule_GetActivityDumpStreamClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpStreamParams, ...grpc.CallOption) api.SecurityModule_GetActivityDumpStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.SecurityModule_GetActivityDumpStreamClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.ActivityDumpStreamParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfig provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) GetConfig(ctx context.Context, in *api.GetConfigParams, opts ...grpc.CallOption) (*api.SecurityConfigMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetConfig")
	}

	var r0 *api.SecurityConfigMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetConfigParams, ...grpc.CallOption) (*api.SecurityConfigMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetConfigParams, ...grpc.CallOption) *api.SecurityConfigMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.SecurityConfigMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.GetConfigParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) GetEvents(ctx context.Context, in *api.GetEventParams, opts ...grpc.CallOption) (api.SecurityModule_GetEventsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetEvents")
	}

	var r0 api.SecurityModule_GetEventsClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetEventParams, ...grpc.CallOption) (api.SecurityModule_GetEventsClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetEventParams, ...grpc.CallOption) api.SecurityModule_GetEventsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.SecurityModule_GetEventsClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.GetEventParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRuleSetReport provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) GetRuleSetReport(ctx context.Context, in *api.GetRuleSetReportParams, opts ...grpc.CallOption) (*api.GetRuleSetReportResultMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRuleSetReport")
	}

	var r0 *api.GetRuleSetReportResultMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetRuleSetReportParams, ...grpc.CallOption) (*api.GetRuleSetReportResultMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetRuleSetReportParams, ...grpc.CallOption) *api.GetRuleSetReportResultMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.GetRuleSetReportResultMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.GetRuleSetReportParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatus provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) GetStatus(ctx context.Context, in *api.GetStatusParams, opts ...grpc.CallOption) (*api.Status, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 *api.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetStatusParams, ...grpc.CallOption) (*api.Status, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetStatusParams, ...grpc.CallOption) *api.Status); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.GetStatusParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListActivityDumps provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) ListActivityDumps(ctx context.Context, in *api.ActivityDumpListParams, opts ...grpc.CallOption) (*api.ActivityDumpListMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListActivityDumps")
	}

	var r0 *api.ActivityDumpListMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpListParams, ...grpc.CallOption) (*api.ActivityDumpListMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpListParams, ...grpc.CallOption) *api.ActivityDumpListMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ActivityDumpListMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.ActivityDumpListParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecurityProfiles provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) ListSecurityProfiles(ctx context.Context, in *api.SecurityProfileListParams, opts ...grpc.CallOption) (*api.SecurityProfileListMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSecurityProfiles")
	}

	var r0 *api.SecurityProfileListMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.SecurityProfileListParams, ...grpc.CallOption) (*api.SecurityProfileListMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.SecurityProfileListParams, ...grpc.CallOption) *api.SecurityProfileListMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.SecurityProfileListMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.SecurityProfileListParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReloadPolicies provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) ReloadPolicies(ctx context.Context, in *api.ReloadPoliciesParams, opts ...grpc.CallOption) (*api.ReloadPoliciesResultMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReloadPolicies")
	}

	var r0 *api.ReloadPoliciesResultMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.ReloadPoliciesParams, ...grpc.CallOption) (*api.ReloadPoliciesResultMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.ReloadPoliciesParams, ...grpc.CallOption) *api.ReloadPoliciesResultMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ReloadPoliciesResultMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.ReloadPoliciesParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunSelfTest provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) RunSelfTest(ctx context.Context, in *api.RunSelfTestParams, opts ...grpc.CallOption) (*api.SecuritySelfTestResultMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RunSelfTest")
	}

	var r0 *api.SecuritySelfTestResultMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.RunSelfTestParams, ...grpc.CallOption) (*api.SecuritySelfTestResultMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.RunSelfTestParams, ...grpc.CallOption) *api.SecuritySelfTestResultMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.SecuritySelfTestResultMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.RunSelfTestParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSecurityProfile provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) SaveSecurityProfile(ctx context.Context, in *api.SecurityProfileSaveParams, opts ...grpc.CallOption) (*api.SecurityProfileSaveMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SaveSecurityProfile")
	}

	var r0 *api.SecurityProfileSaveMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.SecurityProfileSaveParams, ...grpc.CallOption) (*api.SecurityProfileSaveMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.SecurityProfileSaveParams, ...grpc.CallOption) *api.SecurityProfileSaveMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.SecurityProfileSaveMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.SecurityProfileSaveParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopActivityDump provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) StopActivityDump(ctx context.Context, in *api.ActivityDumpStopParams, opts ...grpc.CallOption) (*api.ActivityDumpStopMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopActivityDump")
	}

	var r0 *api.ActivityDumpStopMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpStopParams, ...grpc.CallOption) (*api.ActivityDumpStopMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.ActivityDumpStopParams, ...grpc.CallOption) *api.ActivityDumpStopMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ActivityDumpStopMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.ActivityDumpStopParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TranscodingRequest provides a mock function with given fields: ctx, in, opts
func (_m *SecurityModuleClient) TranscodingRequest(ctx context.Context, in *api.TranscodingRequestParams, opts ...grpc.CallOption) (*api.TranscodingRequestMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TranscodingRequest")
	}

	var r0 *api.TranscodingRequestMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.TranscodingRequestParams, ...grpc.CallOption) (*api.TranscodingRequestMessage, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.TranscodingRequestParams, ...grpc.CallOption) *api.TranscodingRequestMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.TranscodingRequestMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.TranscodingRequestParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSecurityModuleClient creates a new instance of SecurityModuleClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSecurityModuleClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SecurityModuleClient {
	mock := &SecurityModuleClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
