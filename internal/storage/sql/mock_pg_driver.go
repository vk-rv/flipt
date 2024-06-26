// Code generated by mockery v2.42.1. DO NOT EDIT.

package sql

import (
	"context"
	"database/sql/driver"
	"errors"

	"github.com/stretchr/testify/mock"
)

// MockConnector is an autogenerated mock type for the Connector type
type mockConnector struct {
	mock.Mock
}

// Connect provides a mock function with given fields: _a0
func (_m *mockConnector) Connect(_a0 context.Context) (driver.Conn, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 driver.Conn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (driver.Conn, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) driver.Conn); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(driver.Conn)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver provides a mock function with given fields:
func (_m *mockConnector) Driver() driver.Driver {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Driver")
	}

	var r0 driver.Driver
	if rf, ok := ret.Get(0).(func() driver.Driver); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(driver.Driver)
	}

	return r0
}

// newMockConnector creates a new instance of MockConnector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockConnector(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockConnector {
	mock := &mockConnector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

type mockDriverConn struct{}

func (*mockDriverConn) Prepare(query string) (driver.Stmt, error) {
	return nil, errors.New("unexpected call")
}
func (*mockDriverConn) Close() error              { return nil }
func (*mockDriverConn) Begin() (driver.Tx, error) { return nil, errors.New("unexpected call") }

// MockDriverContext is an autogenerated mock type for the DriverContext type
type MockDriverContext struct {
	mock.Mock
}

// OpenConnector provides a mock function with given fields: name
func (_m *MockDriverContext) OpenConnector(name string) (driver.Connector, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for OpenConnector")
	}

	var r0 driver.Connector
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (driver.Connector, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) driver.Connector); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Connector)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockDriverContext creates a new instance of MockDriverContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDriverContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDriverContext {
	mock := &MockDriverContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
