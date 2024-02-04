// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/renan5g/goframework/contracts/log"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Entry is an autogenerated mock type for the Entry type
type Entry struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *Entry) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Level provides a mock function with given fields:
func (_m *Entry) Level() log.Level {
	ret := _m.Called()

	var r0 log.Level
	if rf, ok := ret.Get(0).(func() log.Level); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(log.Level)
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *Entry) Message() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Time provides a mock function with given fields:
func (_m *Entry) Time() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NewEntry creates a new instance of Entry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEntry(t interface {
	mock.TestingT
	Cleanup(func())
}) *Entry {
	mock := &Entry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
