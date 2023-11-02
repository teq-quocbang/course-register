// Code generated by mockery v2.36.0. DO NOT EDIT.

package register

import (
	context "context"

	codetype "github.com/teq-quocbang/course-register/codetype"

	mock "github.com/stretchr/testify/mock"

	model "github.com/teq-quocbang/course-register/model"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// BatchUpdateSwapIsCanCeledStatus provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) BatchUpdateSwapIsCanCeledStatus(_a0 context.Context, _a1 *model.Register) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Register) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_BatchUpdateSwapIsCanCeledStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchUpdateSwapIsCanCeledStatus'
type MockRepository_BatchUpdateSwapIsCanCeledStatus_Call struct {
	*mock.Call
}

// BatchUpdateSwapIsCanCeledStatus is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Register
func (_e *MockRepository_Expecter) BatchUpdateSwapIsCanCeledStatus(_a0 interface{}, _a1 interface{}) *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call {
	return &MockRepository_BatchUpdateSwapIsCanCeledStatus_Call{Call: _e.mock.On("BatchUpdateSwapIsCanCeledStatus", _a0, _a1)}
}

func (_c *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call) Run(run func(_a0 context.Context, _a1 *model.Register)) *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Register))
	})
	return _c
}

func (_c *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call) Return(_a0 error) *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call) RunAndReturn(run func(context.Context, *model.Register) error) *MockRepository_BatchUpdateSwapIsCanCeledStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Create(_a0 context.Context, _a1 *model.Register) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Register) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Register
func (_e *MockRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *model.Register)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Register))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, *model.Register) error) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Get(_a0 context.Context, _a1 *model.Register) (*model.Register, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Register
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Register) (*model.Register, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Register) *model.Register); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Register)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Register) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Register
func (_e *MockRepository_Expecter) Get(_a0 interface{}, _a1 interface{}) *MockRepository_Get_Call {
	return &MockRepository_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *MockRepository_Get_Call) Run(run func(_a0 context.Context, _a1 *model.Register)) *MockRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Register))
	})
	return _c
}

func (_c *MockRepository_Get_Call) Return(_a0 *model.Register, _a1 error) *MockRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Get_Call) RunAndReturn(run func(context.Context, *model.Register) (*model.Register, error)) *MockRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetListByFirstCourseChar provides a mock function with given fields: ctx, firstChar, accountID, semesterID
func (_m *MockRepository) GetListByFirstCourseChar(ctx context.Context, firstChar string, accountID uint, semesterID string) ([]model.Register, error) {
	ret := _m.Called(ctx, firstChar, accountID, semesterID)

	var r0 []model.Register
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint, string) ([]model.Register, error)); ok {
		return rf(ctx, firstChar, accountID, semesterID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint, string) []model.Register); ok {
		r0 = rf(ctx, firstChar, accountID, semesterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Register)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint, string) error); ok {
		r1 = rf(ctx, firstChar, accountID, semesterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetListByFirstCourseChar_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListByFirstCourseChar'
type MockRepository_GetListByFirstCourseChar_Call struct {
	*mock.Call
}

// GetListByFirstCourseChar is a helper method to define mock.On call
//   - ctx context.Context
//   - firstChar string
//   - accountID uint
//   - semesterID string
func (_e *MockRepository_Expecter) GetListByFirstCourseChar(ctx interface{}, firstChar interface{}, accountID interface{}, semesterID interface{}) *MockRepository_GetListByFirstCourseChar_Call {
	return &MockRepository_GetListByFirstCourseChar_Call{Call: _e.mock.On("GetListByFirstCourseChar", ctx, firstChar, accountID, semesterID)}
}

func (_c *MockRepository_GetListByFirstCourseChar_Call) Run(run func(ctx context.Context, firstChar string, accountID uint, semesterID string)) *MockRepository_GetListByFirstCourseChar_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint), args[3].(string))
	})
	return _c
}

func (_c *MockRepository_GetListByFirstCourseChar_Call) Return(_a0 []model.Register, _a1 error) *MockRepository_GetListByFirstCourseChar_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetListByFirstCourseChar_Call) RunAndReturn(run func(context.Context, string, uint, string) ([]model.Register, error)) *MockRepository_GetListByFirstCourseChar_Call {
	_c.Call.Return(run)
	return _c
}

// GetListBySemesterID provides a mock function with given fields: ctx, accountID, semesterID, order, paginator
func (_m *MockRepository) GetListBySemesterID(ctx context.Context, accountID uint, semesterID string, order []string, paginator codetype.Paginator) ([]model.Register, int64, error) {
	ret := _m.Called(ctx, accountID, semesterID, order, paginator)

	var r0 []model.Register
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, []string, codetype.Paginator) ([]model.Register, int64, error)); ok {
		return rf(ctx, accountID, semesterID, order, paginator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, []string, codetype.Paginator) []model.Register); ok {
		r0 = rf(ctx, accountID, semesterID, order, paginator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Register)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, string, []string, codetype.Paginator) int64); ok {
		r1 = rf(ctx, accountID, semesterID, order, paginator)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint, string, []string, codetype.Paginator) error); ok {
		r2 = rf(ctx, accountID, semesterID, order, paginator)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepository_GetListBySemesterID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListBySemesterID'
type MockRepository_GetListBySemesterID_Call struct {
	*mock.Call
}

// GetListBySemesterID is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID uint
//   - semesterID string
//   - order []string
//   - paginator codetype.Paginator
func (_e *MockRepository_Expecter) GetListBySemesterID(ctx interface{}, accountID interface{}, semesterID interface{}, order interface{}, paginator interface{}) *MockRepository_GetListBySemesterID_Call {
	return &MockRepository_GetListBySemesterID_Call{Call: _e.mock.On("GetListBySemesterID", ctx, accountID, semesterID, order, paginator)}
}

func (_c *MockRepository_GetListBySemesterID_Call) Run(run func(ctx context.Context, accountID uint, semesterID string, order []string, paginator codetype.Paginator)) *MockRepository_GetListBySemesterID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(string), args[3].([]string), args[4].(codetype.Paginator))
	})
	return _c
}

func (_c *MockRepository_GetListBySemesterID_Call) Return(_a0 []model.Register, _a1 int64, _a2 error) *MockRepository_GetListBySemesterID_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepository_GetListBySemesterID_Call) RunAndReturn(run func(context.Context, uint, string, []string, codetype.Paginator) ([]model.Register, int64, error)) *MockRepository_GetListBySemesterID_Call {
	_c.Call.Return(run)
	return _c
}

// GetListRegistered provides a mock function with given fields: ctx, accountID, semesterID, order, paginator
func (_m *MockRepository) GetListRegistered(ctx context.Context, accountID uint, semesterID string, order []string, paginator codetype.Paginator) ([]model.Register, int64, error) {
	ret := _m.Called(ctx, accountID, semesterID, order, paginator)

	var r0 []model.Register
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, []string, codetype.Paginator) ([]model.Register, int64, error)); ok {
		return rf(ctx, accountID, semesterID, order, paginator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, []string, codetype.Paginator) []model.Register); ok {
		r0 = rf(ctx, accountID, semesterID, order, paginator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Register)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, string, []string, codetype.Paginator) int64); ok {
		r1 = rf(ctx, accountID, semesterID, order, paginator)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint, string, []string, codetype.Paginator) error); ok {
		r2 = rf(ctx, accountID, semesterID, order, paginator)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepository_GetListRegistered_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListRegistered'
type MockRepository_GetListRegistered_Call struct {
	*mock.Call
}

// GetListRegistered is a helper method to define mock.On call
//   - ctx context.Context
//   - accountID uint
//   - semesterID string
//   - order []string
//   - paginator codetype.Paginator
func (_e *MockRepository_Expecter) GetListRegistered(ctx interface{}, accountID interface{}, semesterID interface{}, order interface{}, paginator interface{}) *MockRepository_GetListRegistered_Call {
	return &MockRepository_GetListRegistered_Call{Call: _e.mock.On("GetListRegistered", ctx, accountID, semesterID, order, paginator)}
}

func (_c *MockRepository_GetListRegistered_Call) Run(run func(ctx context.Context, accountID uint, semesterID string, order []string, paginator codetype.Paginator)) *MockRepository_GetListRegistered_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(string), args[3].([]string), args[4].(codetype.Paginator))
	})
	return _c
}

func (_c *MockRepository_GetListRegistered_Call) Return(_a0 []model.Register, _a1 int64, _a2 error) *MockRepository_GetListRegistered_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepository_GetListRegistered_Call) RunAndReturn(run func(context.Context, uint, string, []string, codetype.Paginator) ([]model.Register, int64, error)) *MockRepository_GetListRegistered_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}