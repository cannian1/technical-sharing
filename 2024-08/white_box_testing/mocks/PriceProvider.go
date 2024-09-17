// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	model "white_box_testing/model"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PriceProvider is an autogenerated mock type for the PriceProvider type
type PriceProvider struct {
	mock.Mock
}

// Latest provides a mock function with given fields:
func (_m *PriceProvider) Latest() (*model.PriceData, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Latest")
	}

	var r0 *model.PriceData
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.PriceData, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.PriceData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PriceData)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: date
func (_m *PriceProvider) List(date time.Time) ([]*model.PriceData, error) {
	ret := _m.Called(date)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*model.PriceData
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time) ([]*model.PriceData, error)); ok {
		return rf(date)
	}
	if rf, ok := ret.Get(0).(func(time.Time) []*model.PriceData); ok {
		r0 = rf(date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PriceData)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPriceProvider creates a new instance of PriceProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceProvider {
	mock := &PriceProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}