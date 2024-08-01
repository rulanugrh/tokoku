// Code generated by mockery v2.40.2. DO NOT EDIT.

package repomocks

import (
	domain "github.com/rulanugrh/tokoku/product/internal/entity/domain"
	mock "github.com/stretchr/testify/mock"
)

// CategoryInterface is an autogenerated mock type for the CategoryInterface type
type CategoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: req
func (_m *CategoryInterface) Create(req domain.Category) (*domain.Category, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Category) (*domain.Category, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(domain.Category) *domain.Category); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.Category) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *CategoryInterface) FindAll() (*[]domain.Category, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 *[]domain.Category
	var r1 error
	if rf, ok := ret.Get(0).(func() (*[]domain.Category, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *[]domain.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Category)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindID provides a mock function with given fields: id
func (_m *CategoryInterface) FindID(id uint) (*domain.Category, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindID")
	}

	var r0 *domain.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.Category, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.Category); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCategoryInterface creates a new instance of CategoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryInterface {
	mock := &CategoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}