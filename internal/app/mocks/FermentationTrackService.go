// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/anne-markis/fermtrack/internal/app/repository"
	mock "github.com/stretchr/testify/mock"
)

// FermentationTrackService is an autogenerated mock type for the FermentationTrackService type
type FermentationTrackService struct {
	mock.Mock
}

type FermentationTrackService_Expecter struct {
	mock *mock.Mock
}

func (_m *FermentationTrackService) EXPECT() *FermentationTrackService_Expecter {
	return &FermentationTrackService_Expecter{mock: &_m.Mock}
}

// GetFermentationAdvice provides a mock function with given fields: ctx, question
func (_m *FermentationTrackService) GetFermentationAdvice(ctx context.Context, question string) (string, error) {
	ret := _m.Called(ctx, question)

	if len(ret) == 0 {
		panic("no return value specified for GetFermentationAdvice")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, question)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, question)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, question)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FermentationTrackService_GetFermentationAdvice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFermentationAdvice'
type FermentationTrackService_GetFermentationAdvice_Call struct {
	*mock.Call
}

// GetFermentationAdvice is a helper method to define mock.On call
//   - ctx context.Context
//   - question string
func (_e *FermentationTrackService_Expecter) GetFermentationAdvice(ctx interface{}, question interface{}) *FermentationTrackService_GetFermentationAdvice_Call {
	return &FermentationTrackService_GetFermentationAdvice_Call{Call: _e.mock.On("GetFermentationAdvice", ctx, question)}
}

func (_c *FermentationTrackService_GetFermentationAdvice_Call) Run(run func(ctx context.Context, question string)) *FermentationTrackService_GetFermentationAdvice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FermentationTrackService_GetFermentationAdvice_Call) Return(_a0 string, _a1 error) *FermentationTrackService_GetFermentationAdvice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FermentationTrackService_GetFermentationAdvice_Call) RunAndReturn(run func(context.Context, string) (string, error)) *FermentationTrackService_GetFermentationAdvice_Call {
	_c.Call.Return(run)
	return _c
}

// GetFermentationByUUID provides a mock function with given fields: ctx, uuid
func (_m *FermentationTrackService) GetFermentationByUUID(ctx context.Context, uuid string) (*repository.Fermentation, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for GetFermentationByUUID")
	}

	var r0 *repository.Fermentation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*repository.Fermentation, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *repository.Fermentation); ok {
		r0 = rf(ctx, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Fermentation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FermentationTrackService_GetFermentationByUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFermentationByUUID'
type FermentationTrackService_GetFermentationByUUID_Call struct {
	*mock.Call
}

// GetFermentationByUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - uuid string
func (_e *FermentationTrackService_Expecter) GetFermentationByUUID(ctx interface{}, uuid interface{}) *FermentationTrackService_GetFermentationByUUID_Call {
	return &FermentationTrackService_GetFermentationByUUID_Call{Call: _e.mock.On("GetFermentationByUUID", ctx, uuid)}
}

func (_c *FermentationTrackService_GetFermentationByUUID_Call) Run(run func(ctx context.Context, uuid string)) *FermentationTrackService_GetFermentationByUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FermentationTrackService_GetFermentationByUUID_Call) Return(_a0 *repository.Fermentation, _a1 error) *FermentationTrackService_GetFermentationByUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FermentationTrackService_GetFermentationByUUID_Call) RunAndReturn(run func(context.Context, string) (*repository.Fermentation, error)) *FermentationTrackService_GetFermentationByUUID_Call {
	_c.Call.Return(run)
	return _c
}

// GetFermentations provides a mock function with given fields: ctx
func (_m *FermentationTrackService) GetFermentations(ctx context.Context) ([]repository.Fermentation, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFermentations")
	}

	var r0 []repository.Fermentation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]repository.Fermentation, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []repository.Fermentation); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Fermentation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FermentationTrackService_GetFermentations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFermentations'
type FermentationTrackService_GetFermentations_Call struct {
	*mock.Call
}

// GetFermentations is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FermentationTrackService_Expecter) GetFermentations(ctx interface{}) *FermentationTrackService_GetFermentations_Call {
	return &FermentationTrackService_GetFermentations_Call{Call: _e.mock.On("GetFermentations", ctx)}
}

func (_c *FermentationTrackService_GetFermentations_Call) Run(run func(ctx context.Context)) *FermentationTrackService_GetFermentations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FermentationTrackService_GetFermentations_Call) Return(_a0 []repository.Fermentation, _a1 error) *FermentationTrackService_GetFermentations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FermentationTrackService_GetFermentations_Call) RunAndReturn(run func(context.Context) ([]repository.Fermentation, error)) *FermentationTrackService_GetFermentations_Call {
	_c.Call.Return(run)
	return _c
}

// NewFermentationTrackService creates a new instance of FermentationTrackService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFermentationTrackService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FermentationTrackService {
	mock := &FermentationTrackService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
