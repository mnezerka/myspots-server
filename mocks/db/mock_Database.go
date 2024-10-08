// Code generated by mockery v2.44.1. DO NOT EDIT.

package db

import (
	db "mnezerka/myspots-server/db"

	mock "github.com/stretchr/testify/mock"
)

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

type MockDatabase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatabase) EXPECT() *MockDatabase_Expecter {
	return &MockDatabase_Expecter{mock: &_m.Mock}
}

// Client provides a mock function with given fields:
func (_m *MockDatabase) Client() db.Client {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 db.Client
	if rf, ok := ret.Get(0).(func() db.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Client)
		}
	}

	return r0
}

// MockDatabase_Client_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Client'
type MockDatabase_Client_Call struct {
	*mock.Call
}

// Client is a helper method to define mock.On call
func (_e *MockDatabase_Expecter) Client() *MockDatabase_Client_Call {
	return &MockDatabase_Client_Call{Call: _e.mock.On("Client")}
}

func (_c *MockDatabase_Client_Call) Run(run func()) *MockDatabase_Client_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDatabase_Client_Call) Return(_a0 db.Client) *MockDatabase_Client_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Client_Call) RunAndReturn(run func() db.Client) *MockDatabase_Client_Call {
	_c.Call.Return(run)
	return _c
}

// Collection provides a mock function with given fields: _a0
func (_m *MockDatabase) Collection(_a0 string) db.Collection {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Collection")
	}

	var r0 db.Collection
	if rf, ok := ret.Get(0).(func(string) db.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Collection)
		}
	}

	return r0
}

// MockDatabase_Collection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Collection'
type MockDatabase_Collection_Call struct {
	*mock.Call
}

// Collection is a helper method to define mock.On call
//   - _a0 string
func (_e *MockDatabase_Expecter) Collection(_a0 interface{}) *MockDatabase_Collection_Call {
	return &MockDatabase_Collection_Call{Call: _e.mock.On("Collection", _a0)}
}

func (_c *MockDatabase_Collection_Call) Run(run func(_a0 string)) *MockDatabase_Collection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDatabase_Collection_Call) Return(_a0 db.Collection) *MockDatabase_Collection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Collection_Call) RunAndReturn(run func(string) db.Collection) *MockDatabase_Collection_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDatabase creates a new instance of MockDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatabase {
	mock := &MockDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
