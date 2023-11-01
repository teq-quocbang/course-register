// Code generated by mockery v2.36.0. DO NOT EDIT.

package semester

import (
	context "context"

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

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Create(_a0 context.Context, _a1 *model.Semester) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Semester) error); ok {
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
//   - _a1 *model.Semester
func (_e *MockRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *model.Semester)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Semester))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, *model.Semester) error) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *MockRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *MockRepository_Delete_Call {
	return &MockRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *MockRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 string)) *MockRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_Delete_Call) Return(_a0 error) *MockRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, semesterID
func (_m *MockRepository) GetByID(ctx context.Context, semesterID string) (model.Semester, error) {
	ret := _m.Called(ctx, semesterID)

	var r0 model.Semester
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Semester, error)); ok {
		return rf(ctx, semesterID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Semester); ok {
		r0 = rf(ctx, semesterID)
	} else {
		r0 = ret.Get(0).(model.Semester)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, semesterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - semesterID string
func (_e *MockRepository_Expecter) GetByID(ctx interface{}, semesterID interface{}) *MockRepository_GetByID_Call {
	return &MockRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, semesterID)}
}

func (_c *MockRepository_GetByID_Call) Run(run func(ctx context.Context, semesterID string)) *MockRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetByID_Call) Return(_a0 model.Semester, _a1 error) *MockRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.Semester, error)) *MockRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetListByYear provides a mock function with given fields: ctx, year
func (_m *MockRepository) GetListByYear(ctx context.Context, year string) ([]model.Semester, error) {
	ret := _m.Called(ctx, year)

	var r0 []model.Semester
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.Semester, error)); ok {
		return rf(ctx, year)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Semester); ok {
		r0 = rf(ctx, year)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Semester)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetListByYear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListByYear'
type MockRepository_GetListByYear_Call struct {
	*mock.Call
}

// GetListByYear is a helper method to define mock.On call
//   - ctx context.Context
//   - year string
func (_e *MockRepository_Expecter) GetListByYear(ctx interface{}, year interface{}) *MockRepository_GetListByYear_Call {
	return &MockRepository_GetListByYear_Call{Call: _e.mock.On("GetListByYear", ctx, year)}
}

func (_c *MockRepository_GetListByYear_Call) Run(run func(ctx context.Context, year string)) *MockRepository_GetListByYear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetListByYear_Call) Return(_a0 []model.Semester, _a1 error) *MockRepository_GetListByYear_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetListByYear_Call) RunAndReturn(run func(context.Context, string) ([]model.Semester, error)) *MockRepository_GetListByYear_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Update(_a0 context.Context, _a1 *model.Semester) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Semester) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *model.Semester
func (_e *MockRepository_Expecter) Update(_a0 interface{}, _a1 interface{}) *MockRepository_Update_Call {
	return &MockRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *MockRepository_Update_Call) Run(run func(_a0 context.Context, _a1 *model.Semester)) *MockRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Semester))
	})
	return _c
}

func (_c *MockRepository_Update_Call) Return(_a0 error) *MockRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Update_Call) RunAndReturn(run func(context.Context, *model.Semester) error) *MockRepository_Update_Call {
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
