// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	models "github.com/addamb/todo-api/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// Queries is an autogenerated mock type for the Queries type
type Queries struct {
	mock.Mock
}

type Queries_Expecter struct {
	mock *mock.Mock
}

func (_m *Queries) EXPECT() *Queries_Expecter {
	return &Queries_Expecter{mock: &_m.Mock}
}

// CreateTodoItem provides a mock function with given fields: name
func (_m *Queries) CreateTodoItem(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queries_CreateTodoItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTodoItem'
type Queries_CreateTodoItem_Call struct {
	*mock.Call
}

// CreateTodoItem is a helper method to define mock.On call
//   - name string
func (_e *Queries_Expecter) CreateTodoItem(name interface{}) *Queries_CreateTodoItem_Call {
	return &Queries_CreateTodoItem_Call{Call: _e.mock.On("CreateTodoItem", name)}
}

func (_c *Queries_CreateTodoItem_Call) Run(run func(name string)) *Queries_CreateTodoItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Queries_CreateTodoItem_Call) Return(_a0 error) *Queries_CreateTodoItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queries_CreateTodoItem_Call) RunAndReturn(run func(string) error) *Queries_CreateTodoItem_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTodoItems provides a mock function with given fields: ids
func (_m *Queries) DeleteTodoItems(ids []int) error {
	ret := _m.Called(ids)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int) error); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queries_DeleteTodoItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTodoItems'
type Queries_DeleteTodoItems_Call struct {
	*mock.Call
}

// DeleteTodoItems is a helper method to define mock.On call
//   - ids []int
func (_e *Queries_Expecter) DeleteTodoItems(ids interface{}) *Queries_DeleteTodoItems_Call {
	return &Queries_DeleteTodoItems_Call{Call: _e.mock.On("DeleteTodoItems", ids)}
}

func (_c *Queries_DeleteTodoItems_Call) Run(run func(ids []int)) *Queries_DeleteTodoItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int))
	})
	return _c
}

func (_c *Queries_DeleteTodoItems_Call) Return(_a0 error) *Queries_DeleteTodoItems_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queries_DeleteTodoItems_Call) RunAndReturn(run func([]int) error) *Queries_DeleteTodoItems_Call {
	_c.Call.Return(run)
	return _c
}

// GetTodos provides a mock function with given fields:
func (_m *Queries) GetTodos() ([]models.Todo, error) {
	ret := _m.Called()

	var r0 []models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_GetTodos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTodos'
type Queries_GetTodos_Call struct {
	*mock.Call
}

// GetTodos is a helper method to define mock.On call
func (_e *Queries_Expecter) GetTodos() *Queries_GetTodos_Call {
	return &Queries_GetTodos_Call{Call: _e.mock.On("GetTodos")}
}

func (_c *Queries_GetTodos_Call) Run(run func()) *Queries_GetTodos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Queries_GetTodos_Call) Return(_a0 []models.Todo, _a1 error) *Queries_GetTodos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_GetTodos_Call) RunAndReturn(run func() ([]models.Todo, error)) *Queries_GetTodos_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTodoItem provides a mock function with given fields: id, name, status
func (_m *Queries) UpdateTodoItem(id int, name string, status bool) error {
	ret := _m.Called(id, name, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, bool) error); ok {
		r0 = rf(id, name, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Queries_UpdateTodoItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTodoItem'
type Queries_UpdateTodoItem_Call struct {
	*mock.Call
}

// UpdateTodoItem is a helper method to define mock.On call
//   - id int
//   - name string
//   - status bool
func (_e *Queries_Expecter) UpdateTodoItem(id interface{}, name interface{}, status interface{}) *Queries_UpdateTodoItem_Call {
	return &Queries_UpdateTodoItem_Call{Call: _e.mock.On("UpdateTodoItem", id, name, status)}
}

func (_c *Queries_UpdateTodoItem_Call) Run(run func(id int, name string, status bool)) *Queries_UpdateTodoItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *Queries_UpdateTodoItem_Call) Return(_a0 error) *Queries_UpdateTodoItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queries_UpdateTodoItem_Call) RunAndReturn(run func(int, string, bool) error) *Queries_UpdateTodoItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueries creates a new instance of Queries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueries(t interface {
	mock.TestingT
	Cleanup(func())
}) *Queries {
	mock := &Queries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
