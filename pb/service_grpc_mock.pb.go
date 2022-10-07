// Code generated by MockGen. DO NOT EDIT.
// Source: service_grpc.pb.go

// Package pb is a generated GoMock package.
package pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockServiceClient is a mock of ServiceClient interface
type MockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceClientMockRecorder
}

// MockServiceClientMockRecorder is the mock recorder for MockServiceClient
type MockServiceClientMockRecorder struct {
	mock *MockServiceClient
}

// NewMockServiceClient creates a new mock instance
func NewMockServiceClient(ctrl *gomock.Controller) *MockServiceClient {
	mock := &MockServiceClient{ctrl: ctrl}
	mock.recorder = &MockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceClient) EXPECT() *MockServiceClientMockRecorder {
	return m.recorder
}

// Status mocks base method
func (m *MockServiceClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockServiceClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockServiceClient)(nil).Status), varargs...)
}

// MockServiceServer is a mock of ServiceServer interface
type MockServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockServiceServerMockRecorder
}

// MockServiceServerMockRecorder is the mock recorder for MockServiceServer
type MockServiceServerMockRecorder struct {
	mock *MockServiceServer
}

// NewMockServiceServer creates a new mock instance
func NewMockServiceServer(ctrl *gomock.Controller) *MockServiceServer {
	mock := &MockServiceServer{ctrl: ctrl}
	mock.recorder = &MockServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceServer) EXPECT() *MockServiceServerMockRecorder {
	return m.recorder
}

// Status mocks base method
func (m *MockServiceServer) Status(arg0 context.Context, arg1 *Empty) (*StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockServiceServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockServiceServer)(nil).Status), arg0, arg1)
}

// mustEmbedUnimplementedServiceServer mocks base method
func (m *MockServiceServer) mustEmbedUnimplementedServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedServiceServer")
}

// mustEmbedUnimplementedServiceServer indicates an expected call of mustEmbedUnimplementedServiceServer
func (mr *MockServiceServerMockRecorder) mustEmbedUnimplementedServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedServiceServer", reflect.TypeOf((*MockServiceServer)(nil).mustEmbedUnimplementedServiceServer))
}

// MockUnsafeServiceServer is a mock of UnsafeServiceServer interface
type MockUnsafeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeServiceServerMockRecorder
}

// MockUnsafeServiceServerMockRecorder is the mock recorder for MockUnsafeServiceServer
type MockUnsafeServiceServerMockRecorder struct {
	mock *MockUnsafeServiceServer
}

// NewMockUnsafeServiceServer creates a new mock instance
func NewMockUnsafeServiceServer(ctrl *gomock.Controller) *MockUnsafeServiceServer {
	mock := &MockUnsafeServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnsafeServiceServer) EXPECT() *MockUnsafeServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedServiceServer mocks base method
func (m *MockUnsafeServiceServer) mustEmbedUnimplementedServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedServiceServer")
}

// mustEmbedUnimplementedServiceServer indicates an expected call of mustEmbedUnimplementedServiceServer
func (mr *MockUnsafeServiceServerMockRecorder) mustEmbedUnimplementedServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedServiceServer", reflect.TypeOf((*MockUnsafeServiceServer)(nil).mustEmbedUnimplementedServiceServer))
}
