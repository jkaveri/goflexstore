// Code generated by mockery v2.40.1. DO NOT EDIT.

package mockquery

import mock "github.com/stretchr/testify/mock"

// Param is an autogenerated mock type for the Param type
type Param struct {
	mock.Mock
}

type Param_Expecter struct {
	mock *mock.Mock
}

func (_m *Param) EXPECT() *Param_Expecter {
	return &Param_Expecter{mock: &_m.Mock}
}

// ParamType provides a mock function with given fields:
func (_m *Param) ParamType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ParamType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Param_ParamType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParamType'
type Param_ParamType_Call struct {
	*mock.Call
}

// ParamType is a helper method to define mock.On call
func (_e *Param_Expecter) ParamType() *Param_ParamType_Call {
	return &Param_ParamType_Call{Call: _e.mock.On("ParamType")}
}

func (_c *Param_ParamType_Call) Run(run func()) *Param_ParamType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Param_ParamType_Call) Return(_a0 string) *Param_ParamType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Param_ParamType_Call) RunAndReturn(run func() string) *Param_ParamType_Call {
	_c.Call.Return(run)
	return _c
}

// NewParam creates a new instance of Param. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParam(t interface {
	mock.TestingT
	Cleanup(func())
}) *Param {
	mock := &Param{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
