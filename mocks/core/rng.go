// Code generated by mockery. DO NOT EDIT.

package core_mock

import mock "github.com/stretchr/testify/mock"

// RNG is an autogenerated mock type for the IRNG type
type RNG struct {
	mock.Mock
}

type RNG_Expecter struct {
	mock *mock.Mock
}

func (_m *RNG) EXPECT() *RNG_Expecter {
	return &RNG_Expecter{mock: &_m.Mock}
}

// Read provides a mock function with given fields: block
func (_m *RNG) Read(block []byte) error {
	ret := _m.Called(block)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RNG_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type RNG_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - block []byte
func (_e *RNG_Expecter) Read(block interface{}) *RNG_Read_Call {
	return &RNG_Read_Call{Call: _e.mock.On("Read", block)}
}

func (_c *RNG_Read_Call) Run(run func(block []byte)) *RNG_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *RNG_Read_Call) Return(err error) *RNG_Read_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *RNG_Read_Call) RunAndReturn(run func([]byte) error) *RNG_Read_Call {
	_c.Call.Return(run)
	return _c
}

// NewRNG creates a new instance of RNG. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRNG(t interface {
	mock.TestingT
	Cleanup(func())
}) *RNG {
	mock := &RNG{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
