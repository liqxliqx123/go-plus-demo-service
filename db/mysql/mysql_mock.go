// Code generated by MockGen. DO NOT EDIT.
// Source: db/mysql/mysql.go

// Package mysql is a generated GoMock package.
package mysql

import (
	context "context"
	models "my-demo-service/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockxxxMySQLInterface is a mock of xxxMySQLInterface interface.
type MockxxxMySQLInterface struct {
	ctrl     *gomock.Controller
	recorder *MockxxxMySQLInterfaceMockRecorder
}

// MockxxxMySQLInterfaceMockRecorder is the mock recorder for MockxxxMySQLInterface.
type MockxxxMySQLInterfaceMockRecorder struct {
	mock *MockxxxMySQLInterface
}

// NewMockxxxMySQLInterface creates a new mock instance.
func NewMockxxxMySQLInterface(ctrl *gomock.Controller) *MockxxxMySQLInterface {
	mock := &MockxxxMySQLInterface{ctrl: ctrl}
	mock.recorder = &MockxxxMySQLInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockxxxMySQLInterface) EXPECT() *MockxxxMySQLInterfaceMockRecorder {
	return m.recorder
}

// CheckDailyxxxEndTime mocks base method.
func (m *MockxxxMySQLInterface) CheckDailyxxxEndTime(ctx context.Context) []*models.xxxSetup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDailyxxxEndTime", ctx)
	ret0, _ := ret[0].([]*models.xxxSetup)
	return ret0
}

// CheckDailyxxxEndTime indicates an expected call of CheckDailyxxxEndTime.
func (mr *MockxxxMySQLInterfaceMockRecorder) CheckDailyxxxEndTime(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDailyxxxEndTime", reflect.TypeOf((*MockxxxMySQLInterface)(nil).CheckDailyxxxEndTime), ctx)
}

// Create mocks base method.
func (m *MockxxxMySQLInterface) Create(d models.xxxTemplate) (models.xxxTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", d)
	ret0, _ := ret[0].(models.xxxTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockxxxMySQLInterfaceMockRecorder) Create(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockxxxMySQLInterface)(nil).Create), d)
}

// CreatexxxResult mocks base method.
func (m *MockxxxMySQLInterface) CreatexxxResult(result *models.xxxResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatexxxResult", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatexxxResult indicates an expected call of CreatexxxResult.
func (mr *MockxxxMySQLInterfaceMockRecorder) CreatexxxResult(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatexxxResult", reflect.TypeOf((*MockxxxMySQLInterface)(nil).CreatexxxResult), result)
}

// CreatexxxSetup mocks base method.
func (m *MockxxxMySQLInterface) CreatexxxSetup(d models.xxxSetup) (models.xxxSetup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatexxxSetup", d)
	ret0, _ := ret[0].(models.xxxSetup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatexxxSetup indicates an expected call of CreatexxxSetup.
func (mr *MockxxxMySQLInterfaceMockRecorder) CreatexxxSetup(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatexxxSetup", reflect.TypeOf((*MockxxxMySQLInterface)(nil).CreatexxxSetup), d)
}

// CreatexxxSetupLayout mocks base method.
func (m *MockxxxMySQLInterface) CreatexxxSetupLayout(d models.xxxSetupLayout) (models.xxxSetupLayout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatexxxSetupLayout", d)
	ret0, _ := ret[0].(models.xxxSetupLayout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatexxxSetupLayout indicates an expected call of CreatexxxSetupLayout.
func (mr *MockxxxMySQLInterfaceMockRecorder) CreatexxxSetupLayout(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatexxxSetupLayout", reflect.TypeOf((*MockxxxMySQLInterface)(nil).CreatexxxSetupLayout), d)
}

// GetByID mocks base method.
func (m *MockxxxMySQLInterface) GetByID(id int64) (models.xxxTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.xxxTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockxxxMySQLInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockxxxMySQLInterface)(nil).GetByID), id)
}

// GetList mocks base method.
func (m *MockxxxMySQLInterface) GetList(conditionMap map[string]interface{}) []models.xxxTemplate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", conditionMap)
	ret0, _ := ret[0].([]models.xxxTemplate)
	return ret0
}

// GetList indicates an expected call of GetList.
func (mr *MockxxxMySQLInterfaceMockRecorder) GetList(conditionMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockxxxMySQLInterface)(nil).GetList), conditionMap)
}

// GetxxxSetupByID mocks base method.
func (m *MockxxxMySQLInterface) GetxxxSetupByID(id int64) (models.xxxSetup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetxxxSetupByID", id)
	ret0, _ := ret[0].(models.xxxSetup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetxxxSetupByID indicates an expected call of GetxxxSetupByID.
func (mr *MockxxxMySQLInterfaceMockRecorder) GetxxxSetupByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetxxxSetupByID", reflect.TypeOf((*MockxxxMySQLInterface)(nil).GetxxxSetupByID), id)
}

// GetxxxSetupLayoutByID mocks base method.
func (m *MockxxxMySQLInterface) GetxxxSetupLayoutByID(id int64) (models.xxxSetupLayout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetxxxSetupLayoutByID", id)
	ret0, _ := ret[0].(models.xxxSetupLayout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetxxxSetupLayoutByID indicates an expected call of GetxxxSetupLayoutByID.
func (mr *MockxxxMySQLInterfaceMockRecorder) GetxxxSetupLayoutByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetxxxSetupLayoutByID", reflect.TypeOf((*MockxxxMySQLInterface)(nil).GetxxxSetupLayoutByID), id)
}

// GetxxxSetupLayoutList mocks base method.
func (m *MockxxxMySQLInterface) GetxxxSetupLayoutList(conditionMap map[string]interface{}) []models.xxxSetupLayout {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetxxxSetupLayoutList", conditionMap)
	ret0, _ := ret[0].([]models.xxxSetupLayout)
	return ret0
}

// GetxxxSetupLayoutList indicates an expected call of GetxxxSetupLayoutList.
func (mr *MockxxxMySQLInterfaceMockRecorder) GetxxxSetupLayoutList(conditionMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetxxxSetupLayoutList", reflect.TypeOf((*MockxxxMySQLInterface)(nil).GetxxxSetupLayoutList), conditionMap)
}

// GetxxxSetupList mocks base method.
func (m *MockxxxMySQLInterface) GetxxxSetupList(conditionMap map[string]interface{}) []models.xxxSetup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetxxxSetupList", conditionMap)
	ret0, _ := ret[0].([]models.xxxSetup)
	return ret0
}

// GetxxxSetupList indicates an expected call of GetxxxSetupList.
func (mr *MockxxxMySQLInterfaceMockRecorder) GetxxxSetupList(conditionMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetxxxSetupList", reflect.TypeOf((*MockxxxMySQLInterface)(nil).GetxxxSetupList), conditionMap)
}

// Update mocks base method.
func (m *MockxxxMySQLInterface) Update(d models.xxxTemplate) (models.xxxTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", d)
	ret0, _ := ret[0].(models.xxxTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockxxxMySQLInterfaceMockRecorder) Update(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockxxxMySQLInterface)(nil).Update), d)
}

// UpdatexxxSetup mocks base method.
func (m *MockxxxMySQLInterface) UpdatexxxSetup(d models.xxxSetup) (models.xxxSetup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatexxxSetup", d)
	ret0, _ := ret[0].(models.xxxSetup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatexxxSetup indicates an expected call of UpdatexxxSetup.
func (mr *MockxxxMySQLInterfaceMockRecorder) UpdatexxxSetup(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatexxxSetup", reflect.TypeOf((*MockxxxMySQLInterface)(nil).UpdatexxxSetup), d)
}

// UpdatexxxSetupLayout mocks base method.
func (m *MockxxxMySQLInterface) UpdatexxxSetupLayout(d models.xxxSetupLayout) (models.xxxSetupLayout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatexxxSetupLayout", d)
	ret0, _ := ret[0].(models.xxxSetupLayout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatexxxSetupLayout indicates an expected call of UpdatexxxSetupLayout.
func (mr *MockxxxMySQLInterfaceMockRecorder) UpdatexxxSetupLayout(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatexxxSetupLayout", reflect.TypeOf((*MockxxxMySQLInterface)(nil).UpdatexxxSetupLayout), d)
}

// UpdateResultStatus mocks base method.
func (m *MockxxxMySQLInterface) UpdateResultStatus(arg0 *models.xxxResult) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateResultStatus", arg0)
}

// UpdateResultStatus indicates an expected call of UpdateResultStatus.
func (mr *MockxxxMySQLInterfaceMockRecorder) UpdateResultStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResultStatus", reflect.TypeOf((*MockxxxMySQLInterface)(nil).UpdateResultStatus), arg0)
}
