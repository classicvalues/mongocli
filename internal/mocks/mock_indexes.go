// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/indexes.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockIndexCreator is a mock of IndexCreator interface
type MockIndexCreator struct {
	ctrl     *gomock.Controller
	recorder *MockIndexCreatorMockRecorder
}

// MockIndexCreatorMockRecorder is the mock recorder for MockIndexCreator
type MockIndexCreatorMockRecorder struct {
	mock *MockIndexCreator
}

// NewMockIndexCreator creates a new mock instance
func NewMockIndexCreator(ctrl *gomock.Controller) *MockIndexCreator {
	mock := &MockIndexCreator{ctrl: ctrl}
	mock.recorder = &MockIndexCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexCreator) EXPECT() *MockIndexCreatorMockRecorder {
	return m.recorder
}

// CreateIndex mocks base method
func (m *MockIndexCreator) CreateIndex(arg0, arg1 string, arg2 *mongodbatlas.IndexConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndex indicates an expected call of CreateIndex
func (mr *MockIndexCreatorMockRecorder) CreateIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockIndexCreator)(nil).CreateIndex), arg0, arg1, arg2)
}