// Code generated by MockGen. DO NOT EDIT.
// Source: repository/postgres/store.go
//
// Generated by this command:
//
//	mockgen -source repository/postgres/store.go -destination test/mock/postgres.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entity "order_service/services/auth/entity"
	entity0 "order_service/services/user/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// AddAuth mocks base method.
func (m *MockAuthRepository) AddAuth(ctx context.Context, data entity.Auth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAuth", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAuth indicates an expected call of AddAuth.
func (mr *MockAuthRepositoryMockRecorder) AddAuth(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAuth", reflect.TypeOf((*MockAuthRepository)(nil).AddAuth), ctx, data)
}

// GetAuth mocks base method.
func (m *MockAuthRepository) GetAuth(ctx context.Context, username string) (*entity0.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuth", ctx, username)
	ret0, _ := ret[0].(*entity0.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuth indicates an expected call of GetAuth.
func (mr *MockAuthRepositoryMockRecorder) GetAuth(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*MockAuthRepository)(nil).GetAuth), ctx, username)
}
