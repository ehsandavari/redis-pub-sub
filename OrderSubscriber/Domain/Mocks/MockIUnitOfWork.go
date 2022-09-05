// Code generated by MockGen. DO NOT EDIT.
// Source: OrderSubscriber/Domain/Interfaces (interfaces: IUnitOfWork)

// Package Mock is a generated GoMock package.
package Mock

import (
	Interfaces "OrderSubscriber/Domain/Interfaces"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUnitOfWork is a mock of IUnitOfWork interface.
type MockIUnitOfWork struct {
	ctrl     *gomock.Controller
	recorder *MockIUnitOfWorkMockRecorder
}

// MockIUnitOfWorkMockRecorder is the mock recorder for MockIUnitOfWork.
type MockIUnitOfWorkMockRecorder struct {
	mock *MockIUnitOfWork
}

// NewMockIUnitOfWork creates a new mock instance.
func NewMockIUnitOfWork(ctrl *gomock.Controller) *MockIUnitOfWork {
	mock := &MockIUnitOfWork{ctrl: ctrl}
	mock.recorder = &MockIUnitOfWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUnitOfWork) EXPECT() *MockIUnitOfWorkMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockIUnitOfWork) Do(arg0 Interfaces.UnitOfWorkBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockIUnitOfWorkMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockIUnitOfWork)(nil).Do), arg0)
}

// OrderRepository mocks base method.
func (m *MockIUnitOfWork) OrderRepository() Interfaces.IOrderRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderRepository")
	ret0, _ := ret[0].(Interfaces.IOrderRepository)
	return ret0
}

// OrderRepository indicates an expected call of OrderRepository.
func (mr *MockIUnitOfWorkMockRecorder) OrderRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderRepository", reflect.TypeOf((*MockIUnitOfWork)(nil).OrderRepository))
}