// Code generated by mockery. DO NOT EDIT.

package core_mock

import mock "github.com/stretchr/testify/mock"

// Cmd is an autogenerated mock type for the ICmd type
type Cmd struct {
	mock.Mock
}

type Cmd_Expecter struct {
	mock *mock.Mock
}

func (_m *Cmd) EXPECT() *Cmd_Expecter {
	return &Cmd_Expecter{mock: &_m.Mock}
}

// Output provides a mock function with given fields:
func (_m *Cmd) Output() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Output")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cmd_Output_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Output'
type Cmd_Output_Call struct {
	*mock.Call
}

// Output is a helper method to define mock.On call
func (_e *Cmd_Expecter) Output() *Cmd_Output_Call {
	return &Cmd_Output_Call{Call: _e.mock.On("Output")}
}

func (_c *Cmd_Output_Call) Run(run func()) *Cmd_Output_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Cmd_Output_Call) Return(output []byte, err error) *Cmd_Output_Call {
	_c.Call.Return(output, err)
	return _c
}

func (_c *Cmd_Output_Call) RunAndReturn(run func() ([]byte, error)) *Cmd_Output_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *Cmd) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Cmd_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type Cmd_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *Cmd_Expecter) String() *Cmd_String_Call {
	return &Cmd_String_Call{Call: _e.mock.On("String")}
}

func (_c *Cmd_String_Call) Run(run func()) *Cmd_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Cmd_String_Call) Return(cmd string) *Cmd_String_Call {
	_c.Call.Return(cmd)
	return _c
}

func (_c *Cmd_String_Call) RunAndReturn(run func() string) *Cmd_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewCmd creates a new instance of Cmd. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCmd(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cmd {
	mock := &Cmd{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
