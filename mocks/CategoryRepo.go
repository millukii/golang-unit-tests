// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "go-tests/models"

	mock "github.com/stretchr/testify/mock"
)

// CategoryRepo is an autogenerated mock type for the CategoryRepo type
type CategoryRepo struct {
	mock.Mock
}

// GetById provides a mock function with given fields: id
func (_m *CategoryRepo) GetById(id int) (*models.Category, error) {
	ret := _m.Called(id)

	var r0 *models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Category, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Category); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCategoryRepo creates a new instance of CategoryRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryRepo {
	mock := &CategoryRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}