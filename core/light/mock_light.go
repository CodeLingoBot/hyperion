// Code generated by mockery v1.0.0. DO NOT EDIT.

// NOTE: run 'make update-mocks' to regenerate

package light

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockLight is an autogenerated mock type for the Light type
type MockLight struct {
	mock.Mock
}

// GetID provides a mock function with given fields:
func (_m *MockLight) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *MockLight) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetType provides a mock function with given fields:
func (_m *MockLight) GetType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetState provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockLight) SetState(_a0 context.Context, _a1 Manager, _a2 TargetState) {
	_m.Called(_a0, _a1, _a2)
}
