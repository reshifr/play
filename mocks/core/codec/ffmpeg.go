// Code generated by mockery. DO NOT EDIT.

package codec_mock

import (
	core "github.com/reshifr/play/core"
	codec "github.com/reshifr/play/core/codec"

	mock "github.com/stretchr/testify/mock"
)

// FFmpeg is an autogenerated mock type for the IFFmpeg type
type FFmpeg struct {
	mock.Mock
}

type FFmpeg_Expecter struct {
	mock *mock.Mock
}

func (_m *FFmpeg) EXPECT() *FFmpeg_Expecter {
	return &FFmpeg_Expecter{mock: &_m.Mock}
}

// ReadTag provides a mock function with given fields: path
func (_m *FFmpeg) ReadTag(path string) (*codec.AudioTag, *core.Error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for ReadTag")
	}

	var r0 *codec.AudioTag
	var r1 *core.Error
	if rf, ok := ret.Get(0).(func(string) (*codec.AudioTag, *core.Error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *codec.AudioTag); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codec.AudioTag)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *core.Error); ok {
		r1 = rf(path)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.Error)
		}
	}

	return r0, r1
}

// FFmpeg_ReadTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadTag'
type FFmpeg_ReadTag_Call struct {
	*mock.Call
}

// ReadTag is a helper method to define mock.On call
//   - path string
func (_e *FFmpeg_Expecter) ReadTag(path interface{}) *FFmpeg_ReadTag_Call {
	return &FFmpeg_ReadTag_Call{Call: _e.mock.On("ReadTag", path)}
}

func (_c *FFmpeg_ReadTag_Call) Run(run func(path string)) *FFmpeg_ReadTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FFmpeg_ReadTag_Call) Return(tag *codec.AudioTag, err *core.Error) *FFmpeg_ReadTag_Call {
	_c.Call.Return(tag, err)
	return _c
}

func (_c *FFmpeg_ReadTag_Call) RunAndReturn(run func(string) (*codec.AudioTag, *core.Error)) *FFmpeg_ReadTag_Call {
	_c.Call.Return(run)
	return _c
}

// NewFFmpeg creates a new instance of FFmpeg. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFFmpeg(t interface {
	mock.TestingT
	Cleanup(func())
}) *FFmpeg {
	mock := &FFmpeg{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}