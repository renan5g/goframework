// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	http "github.com/renan5g/goframework/contracts/http"
	mock "github.com/stretchr/testify/mock"
)

// ResourceController is an autogenerated mock type for the ResourceController type
type ResourceController struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: _a0
func (_m *ResourceController) Destroy(_a0 http.Context) http.Response {
	ret := _m.Called(_a0)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(http.Context) http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Index provides a mock function with given fields: _a0
func (_m *ResourceController) Index(_a0 http.Context) http.Response {
	ret := _m.Called(_a0)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(http.Context) http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Show provides a mock function with given fields: _a0
func (_m *ResourceController) Show(_a0 http.Context) http.Response {
	ret := _m.Called(_a0)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(http.Context) http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Store provides a mock function with given fields: _a0
func (_m *ResourceController) Store(_a0 http.Context) http.Response {
	ret := _m.Called(_a0)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(http.Context) http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *ResourceController) Update(_a0 http.Context) http.Response {
	ret := _m.Called(_a0)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(http.Context) http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// NewResourceController creates a new instance of ResourceController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResourceController(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResourceController {
	mock := &ResourceController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
