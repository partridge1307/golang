// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/hasher.go
//
// Generated by this command:
//
//	mockgen -source pkg/hasher.go -destination services/auth/test/mock/hasher.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHasher is a mock of Hasher interface.
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher.
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance.
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// CompareHash mocks base method.
func (m *MockHasher) CompareHash(hashedPassword, password string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHash", hashedPassword, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareHash indicates an expected call of CompareHash.
func (mr *MockHasherMockRecorder) CompareHash(hashedPassword, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHash", reflect.TypeOf((*MockHasher)(nil).CompareHash), hashedPassword, password)
}

// HashPassword mocks base method.
func (m *MockHasher) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockHasherMockRecorder) HashPassword(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockHasher)(nil).HashPassword), password)
}
