// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	command "github.com/jperdior/recordari-kit/command"
	mock "github.com/stretchr/testify/mock"
)

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// Type provides a mock function with given fields:
func (_m *Command) Type() command.Type {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 command.Type
	if rf, ok := ret.Get(0).(func() command.Type); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(command.Type)
	}

	return r0
}

// NewCommand creates a new instance of Command. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommand(t interface {
	mock.TestingT
	Cleanup(func())
}) *Command {
	mock := &Command{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
