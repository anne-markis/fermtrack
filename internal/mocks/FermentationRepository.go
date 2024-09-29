// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	repository "github.com/anne-markis/fermtrack/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// FermentationRepository is an autogenerated mock type for the FermentationRepository type
type FermentationRepository struct {
	mock.Mock
}

type FermentationRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *FermentationRepository) EXPECT() *FermentationRepository_Expecter {
	return &FermentationRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: fermentation
func (_m *FermentationRepository) Create(fermentation *repository.Fermentation) error {
	ret := _m.Called(fermentation)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Fermentation) error); ok {
		r0 = rf(fermentation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FermentationRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type FermentationRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - fermentation *repository.Fermentation
func (_e *FermentationRepository_Expecter) Create(fermentation interface{}) *FermentationRepository_Create_Call {
	return &FermentationRepository_Create_Call{Call: _e.mock.On("Create", fermentation)}
}

func (_c *FermentationRepository_Create_Call) Run(run func(fermentation *repository.Fermentation)) *FermentationRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*repository.Fermentation))
	})
	return _c
}

func (_c *FermentationRepository_Create_Call) Return(_a0 error) *FermentationRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FermentationRepository_Create_Call) RunAndReturn(run func(*repository.Fermentation) error) *FermentationRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: uuid
func (_m *FermentationRepository) Delete(uuid string) error {
	ret := _m.Called(uuid)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FermentationRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FermentationRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - uuid string
func (_e *FermentationRepository_Expecter) Delete(uuid interface{}) *FermentationRepository_Delete_Call {
	return &FermentationRepository_Delete_Call{Call: _e.mock.On("Delete", uuid)}
}

func (_c *FermentationRepository_Delete_Call) Run(run func(uuid string)) *FermentationRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FermentationRepository_Delete_Call) Return(_a0 error) *FermentationRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FermentationRepository_Delete_Call) RunAndReturn(run func(string) error) *FermentationRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *FermentationRepository) FindAll() ([]repository.Fermentation, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []repository.Fermentation
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]repository.Fermentation, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []repository.Fermentation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Fermentation)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FermentationRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type FermentationRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *FermentationRepository_Expecter) FindAll() *FermentationRepository_FindAll_Call {
	return &FermentationRepository_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *FermentationRepository_FindAll_Call) Run(run func()) *FermentationRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FermentationRepository_FindAll_Call) Return(_a0 []repository.Fermentation, _a1 error) *FermentationRepository_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FermentationRepository_FindAll_Call) RunAndReturn(run func() ([]repository.Fermentation, error)) *FermentationRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindByID provides a mock function with given fields: uuid
func (_m *FermentationRepository) FindByID(uuid string) (*repository.Fermentation, error) {
	ret := _m.Called(uuid)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *repository.Fermentation
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.Fermentation, error)); ok {
		return rf(uuid)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.Fermentation); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Fermentation)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FermentationRepository_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type FermentationRepository_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - uuid string
func (_e *FermentationRepository_Expecter) FindByID(uuid interface{}) *FermentationRepository_FindByID_Call {
	return &FermentationRepository_FindByID_Call{Call: _e.mock.On("FindByID", uuid)}
}

func (_c *FermentationRepository_FindByID_Call) Run(run func(uuid string)) *FermentationRepository_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FermentationRepository_FindByID_Call) Return(_a0 *repository.Fermentation, _a1 error) *FermentationRepository_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FermentationRepository_FindByID_Call) RunAndReturn(run func(string) (*repository.Fermentation, error)) *FermentationRepository_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: fermentation
func (_m *FermentationRepository) Update(fermentation *repository.Fermentation) error {
	ret := _m.Called(fermentation)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Fermentation) error); ok {
		r0 = rf(fermentation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FermentationRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type FermentationRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - fermentation *repository.Fermentation
func (_e *FermentationRepository_Expecter) Update(fermentation interface{}) *FermentationRepository_Update_Call {
	return &FermentationRepository_Update_Call{Call: _e.mock.On("Update", fermentation)}
}

func (_c *FermentationRepository_Update_Call) Run(run func(fermentation *repository.Fermentation)) *FermentationRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*repository.Fermentation))
	})
	return _c
}

func (_c *FermentationRepository_Update_Call) Return(_a0 error) *FermentationRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FermentationRepository_Update_Call) RunAndReturn(run func(*repository.Fermentation) error) *FermentationRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewFermentationRepository creates a new instance of FermentationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFermentationRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FermentationRepository {
	mock := &FermentationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
