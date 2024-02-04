// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	translation "github.com/renan5g/goframework/contracts/translation"
	mock "github.com/stretchr/testify/mock"
)

// Translator is an autogenerated mock type for the Translator type
type Translator struct {
	mock.Mock
}

// Choice provides a mock function with given fields: key, number, options
func (_m *Translator) Choice(key string, number int, options ...translation.Option) string {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, number)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int, ...translation.Option) string); ok {
		r0 = rf(key, number, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Get provides a mock function with given fields: key, options
func (_m *Translator) Get(key string, options ...translation.Option) string {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...translation.Option) string); ok {
		r0 = rf(key, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetFallback provides a mock function with given fields:
func (_m *Translator) GetFallback() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLocale provides a mock function with given fields:
func (_m *Translator) GetLocale() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Has provides a mock function with given fields: key, options
func (_m *Translator) Has(key string, options ...translation.Option) bool {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...translation.Option) bool); ok {
		r0 = rf(key, options...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetFallback provides a mock function with given fields: locale
func (_m *Translator) SetFallback(locale string) context.Context {
	ret := _m.Called(locale)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(string) context.Context); ok {
		r0 = rf(locale)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// SetLocale provides a mock function with given fields: locale
func (_m *Translator) SetLocale(locale string) context.Context {
	ret := _m.Called(locale)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(string) context.Context); ok {
		r0 = rf(locale)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// NewTranslator creates a new instance of Translator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTranslator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Translator {
	mock := &Translator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
