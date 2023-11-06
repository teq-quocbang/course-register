// Code generated by mockery v2.36.0. DO NOT EDIT.

package cache

import mock "github.com/stretchr/testify/mock"

// MockICache is an autogenerated mock type for the ICache type
type MockICache struct {
	mock.Mock
}

type MockICache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockICache) EXPECT() *MockICache_Expecter {
	return &MockICache_Expecter{mock: &_m.Mock}
}

// Register provides a mock function with given fields:
func (_m *MockICache) Register() RegisterService {
	ret := _m.Called()

	var r0 RegisterService
	if rf, ok := ret.Get(0).(func() RegisterService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RegisterService)
		}
	}

	return r0
}

// MockICache_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockICache_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
func (_e *MockICache_Expecter) Register() *MockICache_Register_Call {
	return &MockICache_Register_Call{Call: _e.mock.On("Register")}
}

func (_c *MockICache_Register_Call) Run(run func()) *MockICache_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockICache_Register_Call) Return(_a0 RegisterService) *MockICache_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockICache_Register_Call) RunAndReturn(run func() RegisterService) *MockICache_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockICache creates a new instance of MockICache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockICache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockICache {
	mock := &MockICache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}