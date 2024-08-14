// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	mock "github.com/stretchr/testify/mock"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, task
func (_m *TaskUsecase) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	ret := _m.Called(ctx, task)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) (*domain.Task, error)); ok {
		return rf(ctx, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) *domain.Task); ok {
		r0 = rf(ctx, task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TaskUsecase) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *TaskUsecase) GetAll(ctx context.Context) ([]domain.Task, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Task, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TaskUsecase) GetByID(ctx context.Context, id int) (*domain.Task, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*domain.Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, task
func (_m *TaskUsecase) Update(ctx context.Context, id int, task *domain.Task) (*domain.Task, error) {
	ret := _m.Called(ctx, id, task)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *domain.Task) (*domain.Task, error)); ok {
		return rf(ctx, id, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *domain.Task) *domain.Task); ok {
		r0 = rf(ctx, id, task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *domain.Task) error); ok {
		r1 = rf(ctx, id, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
