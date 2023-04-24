// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../pkg/helper/finder.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	helper "github.com/bsn-si/IPEHR-gateway/src/pkg/helper"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockFinder is a mock of Finder interface.
type MockFinder struct {
	ctrl     *gomock.Controller
	recorder *MockFinderMockRecorder
}

// MockFinderMockRecorder is the mock recorder for MockFinder.
type MockFinderMockRecorder struct {
	mock *MockFinder
}

// NewMockFinder creates a new mock instance.
func NewMockFinder(ctrl *gomock.Controller) *MockFinder {
	mock := &MockFinder{ctrl: ctrl}
	mock.recorder = &MockFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinder) EXPECT() *MockFinderMockRecorder {
	return m.recorder
}

// GetEhrUUIDByUserID mocks base method.
func (m *MockFinder) GetEhrUUIDByUserID(ctx context.Context, userID, systemID string) (*uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEhrUUIDByUserID", ctx, userID, systemID)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEhrUUIDByUserID indicates an expected call of GetEhrUUIDByUserID.
func (mr *MockFinderMockRecorder) GetEhrUUIDByUserID(ctx, userID, systemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEhrUUIDByUserID", reflect.TypeOf((*MockFinder)(nil).GetEhrUUIDByUserID), ctx, userID, systemID)
}

// IsExist mocks base method.
func (m *MockFinder) IsExist(ctx context.Context, args ...string) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsExist", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExist indicates an expected call of IsExist.
func (mr *MockFinderMockRecorder) IsExist(ctx interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockFinder)(nil).IsExist), varargs...)
}

// MockSearcher is a mock of Searcher interface.
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher.
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance.
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// GetEhrUUIDByUserID mocks base method.
func (m *MockSearcher) GetEhrUUIDByUserID() (*uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEhrUUIDByUserID")
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEhrUUIDByUserID indicates an expected call of GetEhrUUIDByUserID.
func (mr *MockSearcherMockRecorder) GetEhrUUIDByUserID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEhrUUIDByUserID", reflect.TypeOf((*MockSearcher)(nil).GetEhrUUIDByUserID))
}

// IsEhrBelongsToUser mocks base method.
func (m *MockSearcher) IsEhrBelongsToUser() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEhrBelongsToUser")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEhrBelongsToUser indicates an expected call of IsEhrBelongsToUser.
func (mr *MockSearcherMockRecorder) IsEhrBelongsToUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEhrBelongsToUser", reflect.TypeOf((*MockSearcher)(nil).IsEhrBelongsToUser))
}

// IsExist mocks base method.
func (m *MockSearcher) IsExist(ID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", ID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExist indicates an expected call of IsExist.
func (mr *MockSearcherMockRecorder) IsExist(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockSearcher)(nil).IsExist), ID)
}

// UseService mocks base method.
func (m *MockSearcher) UseService(s helper.Finder) *helper.Search {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseService", s)
	ret0, _ := ret[0].(*helper.Search)
	return ret0
}

// UseService indicates an expected call of UseService.
func (mr *MockSearcherMockRecorder) UseService(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseService", reflect.TypeOf((*MockSearcher)(nil).UseService), s)
}
