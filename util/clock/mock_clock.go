// Code generated by mockery v1.0.0. DO NOT EDIT.

// NOTE: run 'make update-mocks' to regenerate

package clock

import mock "github.com/stretchr/testify/mock"
import time "time"

// MockClock is an autogenerated mock type for the Clock type
type MockClock struct {
	mock.Mock
}

// After provides a mock function with given fields: d
func (_m *MockClock) After(d time.Duration) <-chan time.Time {
	ret := _m.Called(d)

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func(time.Duration) <-chan time.Time); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// Now provides a mock function with given fields:
func (_m *MockClock) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Sleep provides a mock function with given fields: d
func (_m *MockClock) Sleep(d time.Duration) {
	_m.Called(d)
}