// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	http "github.com/renan5g/goframework/contracts/http"
	mock "github.com/stretchr/testify/mock"
)

// Limit is an autogenerated mock type for the Limit type
type Limit struct {
	mock.Mock
}

// By provides a mock function with given fields: key
func (_m *Limit) By(key string) http.Limit {
	ret := _m.Called(key)

	var r0 http.Limit
	if rf, ok := ret.Get(0).(func(string) http.Limit); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Limit)
		}
	}

	return r0
}

// Response provides a mock function with given fields: _a0
func (_m *Limit) Response(_a0 func(http.Context)) http.Limit {
	ret := _m.Called(_a0)

	var r0 http.Limit
	if rf, ok := ret.Get(0).(func(func(http.Context)) http.Limit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Limit)
		}
	}

	return r0
}

// NewLimit creates a new instance of Limit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLimit(t interface {
	mock.TestingT
	Cleanup(func())
}) *Limit {
	mock := &Limit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
