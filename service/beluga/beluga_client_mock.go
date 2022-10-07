// Package beluga is a generated GoMock package.
package beluga

import (
	context "context"
	reflect "reflect"

	"gitlab.xxx.com/xxx-xxx/go-kit/services/beluga"

	gomock "github.com/golang/mock/gomock"
	beluga0 "gitlab.xxx.com/xxx-xxx/go-kit/pb/beluga"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateStatefulDagTask mocks base method.
func (m *MockClient) CreateStatefulDagTask(arg0 context.Context, arg1 string, arg2 *beluga0.StatefulDagTaskSpec, arg3 *beluga0.Metadata, arg4 *beluga0.CrontabStrategy, arg5 *beluga0.ErrorStrategy, arg6 int64, arg7 *beluga0.ResourceStrategy, arg8 ...beluga.TaskOption) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7}
	for _, a := range arg8 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStatefulDagTask", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStatefulDagTask indicates an expected call of CreateStatefulDagTask.
func (mr *MockClientMockRecorder) CreateStatefulDagTask(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}, arg8 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7}, arg8...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStatefulDagTask", reflect.TypeOf((*MockClient)(nil).CreateStatefulDagTask), varargs...)
}

// DeleteTask mocks base method.
func (m *MockClient) DeleteTask(arg0 context.Context, arg1 int64, arg2 ...beluga.TaskOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockClientMockRecorder) DeleteTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockClient)(nil).DeleteTask), varargs...)
}

// GetTask mocks base method.
func (m *MockClient) GetTask(arg0 context.Context, arg1 int64, arg2 ...beluga.TaskOption) (*beluga0.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTask", varargs...)
	ret0, _ := ret[0].(*beluga0.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockClientMockRecorder) GetTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockClient)(nil).GetTask), varargs...)
}

// GetTaskList mocks base method.
func (m *MockClient) GetTaskList(arg0 context.Context, arg1 string, arg2 beluga0.TaskType, arg3 beluga0.TaskStatus, arg4 *beluga0.MetadataFilter, arg5 beluga.Pagination, arg6 ...beluga.TaskOption) ([]*beluga0.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4, arg5}
	for _, a := range arg6 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTaskList", varargs...)
	ret0, _ := ret[0].([]*beluga0.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskList indicates an expected call of GetTaskList.
func (mr *MockClientMockRecorder) GetTaskList(arg0, arg1, arg2, arg3, arg4, arg5 interface{}, arg6 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5}, arg6...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskList", reflect.TypeOf((*MockClient)(nil).GetTaskList), varargs...)
}

// ModifyStatefulDagTask mocks base method.
func (m *MockClient) ModifyStatefulDagTask(arg0 context.Context, arg1 int64, arg2 string, arg3 *beluga0.StatefulDagTaskSpec, arg4 *beluga0.Metadata, arg5 *beluga0.CrontabStrategy, arg6 *beluga0.ErrorStrategy, arg7 beluga0.TaskStatus, arg8 int64, arg9 *beluga0.ResourceStrategy, arg10 ...beluga.TaskOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9}
	for _, a := range arg10 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModifyStatefulDagTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyStatefulDagTask indicates an expected call of ModifyStatefulDagTask.
func (mr *MockClientMockRecorder) ModifyStatefulDagTask(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}, arg10 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9}, arg10...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyStatefulDagTask", reflect.TypeOf((*MockClient)(nil).ModifyStatefulDagTask), varargs...)
}

// TriggerTask mocks base method.
func (m *MockClient) TriggerTask(arg0 context.Context, arg1 int64, arg2 ...beluga.TaskOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// TriggerTask indicates an expected call of TriggerTask.
func (mr *MockClientMockRecorder) TriggerTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerTask", reflect.TypeOf((*MockClient)(nil).TriggerTask), varargs...)
}
