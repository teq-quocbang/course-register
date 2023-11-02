// Code generated by mockery v2.36.0. DO NOT EDIT.

package class

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

// BatchDeCreMember provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) BatchDeCreMember(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_BatchDeCreMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchDeCreMember'
type MockRepository_BatchDeCreMember_Call struct {
	*mock.Call
}

// BatchDeCreMember is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *MockRepository_Expecter) BatchDeCreMember(_a0 interface{}, _a1 interface{}) *MockRepository_BatchDeCreMember_Call {
	return &MockRepository_BatchDeCreMember_Call{Call: _e.mock.On("BatchDeCreMember", _a0, _a1)}
}

func (_c *MockRepository_BatchDeCreMember_Call) Run(run func(_a0 context.Context, _a1 string)) *MockRepository_BatchDeCreMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_BatchDeCreMember_Call) Return(_a0 error) *MockRepository_BatchDeCreMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_BatchDeCreMember_Call) RunAndReturn(run func(context.Context, string) error) *MockRepository_BatchDeCreMember_Call {
	_c.Call.Return(run)
	return _c
}

// BatchInCreMember provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) BatchInCreMember(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_BatchInCreMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchInCreMember'
type MockRepository_BatchInCreMember_Call struct {
	*mock.Call
}

// BatchInCreMember is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *MockRepository_Expecter) BatchInCreMember(_a0 interface{}, _a1 interface{}) *MockRepository_BatchInCreMember_Call {
	return &MockRepository_BatchInCreMember_Call{Call: _e.mock.On("BatchInCreMember", _a0, _a1)}
}

func (_c *MockRepository_BatchInCreMember_Call) Run(run func(_a0 context.Context, _a1 string)) *MockRepository_BatchInCreMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_BatchInCreMember_Call) Return(_a0 error) *MockRepository_BatchInCreMember_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_BatchInCreMember_Call) RunAndReturn(run func(context.Context, string) error) *MockRepository_BatchInCreMember_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Create(_a0 context.Context, _a1 *model.Class) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Class) error); ok {
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
//   - _a1 *model.Class
func (_e *MockRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *model.Class)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Class))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, *model.Class) error) *MockRepository_Create_Call {
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

// GetByID provides a mock function with given fields: ctx, classID
func (_m *MockRepository) GetByID(ctx context.Context, classID string) (model.Class, error) {
	ret := _m.Called(ctx, classID)

	var r0 model.Class
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Class, error)); ok {
		return rf(ctx, classID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Class); ok {
		r0 = rf(ctx, classID)
	} else {
		r0 = ret.Get(0).(model.Class)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, classID)
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
//   - classID string
func (_e *MockRepository_Expecter) GetByID(ctx interface{}, classID interface{}) *MockRepository_GetByID_Call {
	return &MockRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, classID)}
}

func (_c *MockRepository_GetByID_Call) Run(run func(ctx context.Context, classID string)) *MockRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetByID_Call) Return(_a0 model.Class, _a1 error) *MockRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetByID_Call) RunAndReturn(run func(context.Context, string) (model.Class, error)) *MockRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetListBySemester provides a mock function with given fields: ctx, semesterID
func (_m *MockRepository) GetListBySemester(ctx context.Context, semesterID string) ([]model.Class, error) {
	ret := _m.Called(ctx, semesterID)

	var r0 []model.Class
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.Class, error)); ok {
		return rf(ctx, semesterID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Class); ok {
		r0 = rf(ctx, semesterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Class)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, semesterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetListBySemester_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListBySemester'
type MockRepository_GetListBySemester_Call struct {
	*mock.Call
}

// GetListBySemester is a helper method to define mock.On call
//   - ctx context.Context
//   - semesterID string
func (_e *MockRepository_Expecter) GetListBySemester(ctx interface{}, semesterID interface{}) *MockRepository_GetListBySemester_Call {
	return &MockRepository_GetListBySemester_Call{Call: _e.mock.On("GetListBySemester", ctx, semesterID)}
}

func (_c *MockRepository_GetListBySemester_Call) Run(run func(ctx context.Context, semesterID string)) *MockRepository_GetListBySemester_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetListBySemester_Call) Return(_a0 []model.Class, _a1 error) *MockRepository_GetListBySemester_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetListBySemester_Call) RunAndReturn(run func(context.Context, string) ([]model.Class, error)) *MockRepository_GetListBySemester_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Update(_a0 context.Context, _a1 *model.Class) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Class) error); ok {
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
//   - _a1 *model.Class
func (_e *MockRepository_Expecter) Update(_a0 interface{}, _a1 interface{}) *MockRepository_Update_Call {
	return &MockRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *MockRepository_Update_Call) Run(run func(_a0 context.Context, _a1 *model.Class)) *MockRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Class))
	})
	return _c
}

func (_c *MockRepository_Update_Call) Return(_a0 error) *MockRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Update_Call) RunAndReturn(run func(context.Context, *model.Class) error) *MockRepository_Update_Call {
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