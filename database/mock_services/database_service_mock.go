// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mungujn/weather/database/services (interfaces: DatabaseServiceClient)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	services "github.com/mungujn/weather/database/services"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockDatabaseServiceClient is a mock of DatabaseServiceClient interface
type MockDatabaseServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseServiceClientMockRecorder
}

// MockDatabaseServiceClientMockRecorder is the mock recorder for MockDatabaseServiceClient
type MockDatabaseServiceClientMockRecorder struct {
	mock *MockDatabaseServiceClient
}

// NewMockDatabaseServiceClient creates a new mock instance
func NewMockDatabaseServiceClient(ctrl *gomock.Controller) *MockDatabaseServiceClient {
	mock := &MockDatabaseServiceClient{ctrl: ctrl}
	mock.recorder = &MockDatabaseServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseServiceClient) EXPECT() *MockDatabaseServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDatabaseServiceClient) Create(arg0 context.Context, arg1 *services.Data, arg2 ...grpc.CallOption) (*services.Result, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*services.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDatabaseServiceClientMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDatabaseServiceClient)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *MockDatabaseServiceClient) Delete(arg0 context.Context, arg1 *services.Data, arg2 ...grpc.CallOption) (*services.Result, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*services.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockDatabaseServiceClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabaseServiceClient)(nil).Delete), varargs...)
}

// Read mocks base method
func (m *MockDatabaseServiceClient) Read(arg0 context.Context, arg1 *services.Data, arg2 ...grpc.CallOption) (*services.Data, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*services.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockDatabaseServiceClientMockRecorder) Read(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDatabaseServiceClient)(nil).Read), varargs...)
}

// Update mocks base method
func (m *MockDatabaseServiceClient) Update(arg0 context.Context, arg1 *services.Data, arg2 ...grpc.CallOption) (*services.Result, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*services.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDatabaseServiceClientMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDatabaseServiceClient)(nil).Update), varargs...)
}