// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/api/computers/usecases/computer.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	entities "backend/internal/app/api/computers/domain/entities"
	requests "backend/internal/app/api/computers/interface/requests"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockComputerUsecase is a mock of ComputerUsecase interface.
type MockComputerUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockComputerUsecaseMockRecorder
}

// MockComputerUsecaseMockRecorder is the mock recorder for MockComputerUsecase.
type MockComputerUsecaseMockRecorder struct {
	mock *MockComputerUsecase
}

// NewMockComputerUsecase creates a new mock instance.
func NewMockComputerUsecase(ctrl *gomock.Controller) *MockComputerUsecase {
	mock := &MockComputerUsecase{ctrl: ctrl}
	mock.recorder = &MockComputerUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputerUsecase) EXPECT() *MockComputerUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockComputerUsecase) Create(arg0 requests.CreateComputerInput) (*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockComputerUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockComputerUsecase)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockComputerUsecase) Delete(arg0 uint) (*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockComputerUsecaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockComputerUsecase)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockComputerUsecase) Get(arg0 uint) (*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockComputerUsecaseMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockComputerUsecase)(nil).Get), arg0)
}

// Search mocks base method.
func (m *MockComputerUsecase) Search(arg0 *requests.SearchComputersQuery) (*[]entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].(*[]entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockComputerUsecaseMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockComputerUsecase)(nil).Search), arg0)
}

// Update mocks base method.
func (m *MockComputerUsecase) Update(arg0 uint, arg1 requests.UpdateComputerInput) (*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockComputerUsecaseMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockComputerUsecase)(nil).Update), arg0, arg1)
}

// Wake mocks base method.
func (m *MockComputerUsecase) Wake(arg0 uint) (*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wake", arg0)
	ret0, _ := ret[0].(*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wake indicates an expected call of Wake.
func (mr *MockComputerUsecaseMockRecorder) Wake(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wake", reflect.TypeOf((*MockComputerUsecase)(nil).Wake), arg0)
}
