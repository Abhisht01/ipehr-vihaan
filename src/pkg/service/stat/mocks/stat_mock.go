// Code generated by MockGen. DO NOT EDIT.
// Source: ./stat.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPatientsRepository is a mock of PatientsRepository interface.
type MockPatientsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPatientsRepositoryMockRecorder
}

// MockPatientsRepositoryMockRecorder is the mock recorder for MockPatientsRepository.
type MockPatientsRepositoryMockRecorder struct {
	mock *MockPatientsRepository
}

// NewMockPatientsRepository creates a new mock instance.
func NewMockPatientsRepository(ctrl *gomock.Controller) *MockPatientsRepository {
	mock := &MockPatientsRepository{ctrl: ctrl}
	mock.recorder = &MockPatientsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPatientsRepository) EXPECT() *MockPatientsRepositoryMockRecorder {
	return m.recorder
}

// StatDocumentsCountGet mocks base method.
func (m *MockPatientsRepository) StatDocumentsCountGet(ctx context.Context, start, end int64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatDocumentsCountGet", ctx, start, end)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatDocumentsCountGet indicates an expected call of StatDocumentsCountGet.
func (mr *MockPatientsRepositoryMockRecorder) StatDocumentsCountGet(ctx, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatDocumentsCountGet", reflect.TypeOf((*MockPatientsRepository)(nil).StatDocumentsCountGet), ctx, start, end)
}

// StatPatientsCountGet mocks base method.
func (m *MockPatientsRepository) StatPatientsCountGet(ctx context.Context, start, end int64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatPatientsCountGet", ctx, start, end)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatPatientsCountGet indicates an expected call of StatPatientsCountGet.
func (mr *MockPatientsRepositoryMockRecorder) StatPatientsCountGet(ctx, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatPatientsCountGet", reflect.TypeOf((*MockPatientsRepository)(nil).StatPatientsCountGet), ctx, start, end)
}
