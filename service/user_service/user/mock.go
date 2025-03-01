// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=mock.go -package=user
//

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	models "github.com/syedomair/backend-microservices/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetAllUserDB mocks base method.
func (m *MockRepository) GetAllUserDB(limit, offset int, orderby, sort string) ([]*models.User, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserDB", limit, offset, orderby, sort)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllUserDB indicates an expected call of GetAllUserDB.
func (mr *MockRepositoryMockRecorder) GetAllUserDB(limit, offset, orderby, sort any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserDB", reflect.TypeOf((*MockRepository)(nil).GetAllUserDB), limit, offset, orderby, sort)
}

// GetUserAvgAge mocks base method.
func (m *MockRepository) GetUserAvgAge() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAvgAge")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAvgAge indicates an expected call of GetUserAvgAge.
func (mr *MockRepositoryMockRecorder) GetUserAvgAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAvgAge", reflect.TypeOf((*MockRepository)(nil).GetUserAvgAge))
}

// GetUserAvgSalary mocks base method.
func (m *MockRepository) GetUserAvgSalary() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAvgSalary")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAvgSalary indicates an expected call of GetUserAvgSalary.
func (mr *MockRepositoryMockRecorder) GetUserAvgSalary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAvgSalary", reflect.TypeOf((*MockRepository)(nil).GetUserAvgSalary))
}

// GetUserHighAge mocks base method.
func (m *MockRepository) GetUserHighAge() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserHighAge")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserHighAge indicates an expected call of GetUserHighAge.
func (mr *MockRepositoryMockRecorder) GetUserHighAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserHighAge", reflect.TypeOf((*MockRepository)(nil).GetUserHighAge))
}

// GetUserHighSalary mocks base method.
func (m *MockRepository) GetUserHighSalary() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserHighSalary")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserHighSalary indicates an expected call of GetUserHighSalary.
func (mr *MockRepositoryMockRecorder) GetUserHighSalary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserHighSalary", reflect.TypeOf((*MockRepository)(nil).GetUserHighSalary))
}

// GetUserLowAge mocks base method.
func (m *MockRepository) GetUserLowAge() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLowAge")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLowAge indicates an expected call of GetUserLowAge.
func (mr *MockRepositoryMockRecorder) GetUserLowAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLowAge", reflect.TypeOf((*MockRepository)(nil).GetUserLowAge))
}

// GetUserLowSalary mocks base method.
func (m *MockRepository) GetUserLowSalary() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLowSalary")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLowSalary indicates an expected call of GetUserLowSalary.
func (mr *MockRepositoryMockRecorder) GetUserLowSalary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLowSalary", reflect.TypeOf((*MockRepository)(nil).GetUserLowSalary))
}
