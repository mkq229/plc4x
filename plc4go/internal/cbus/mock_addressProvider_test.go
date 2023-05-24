/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.28.1. DO NOT EDIT.

package cbus

import (
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// mockAddressProvider is an autogenerated mock type for the addressProvider type
type mockAddressProvider struct {
	mock.Mock
}

type mockAddressProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockAddressProvider) EXPECT() *mockAddressProvider_Expecter {
	return &mockAddressProvider_Expecter{mock: &_m.Mock}
}

// Addrs provides a mock function with given fields:
func (_m *mockAddressProvider) Addrs() ([]net.Addr, error) {
	ret := _m.Called()

	var r0 []net.Addr
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]net.Addr, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]net.Addr)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAddressProvider_Addrs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Addrs'
type mockAddressProvider_Addrs_Call struct {
	*mock.Call
}

// Addrs is a helper method to define mock.On call
func (_e *mockAddressProvider_Expecter) Addrs() *mockAddressProvider_Addrs_Call {
	return &mockAddressProvider_Addrs_Call{Call: _e.mock.On("Addrs")}
}

func (_c *mockAddressProvider_Addrs_Call) Run(run func()) *mockAddressProvider_Addrs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockAddressProvider_Addrs_Call) Return(_a0 []net.Addr, _a1 error) *mockAddressProvider_Addrs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAddressProvider_Addrs_Call) RunAndReturn(run func() ([]net.Addr, error)) *mockAddressProvider_Addrs_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *mockAddressProvider) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// mockAddressProvider_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type mockAddressProvider_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *mockAddressProvider_Expecter) String() *mockAddressProvider_String_Call {
	return &mockAddressProvider_String_Call{Call: _e.mock.On("String")}
}

func (_c *mockAddressProvider_String_Call) Run(run func()) *mockAddressProvider_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockAddressProvider_String_Call) Return(_a0 string) *mockAddressProvider_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAddressProvider_String_Call) RunAndReturn(run func() string) *mockAddressProvider_String_Call {
	_c.Call.Return(run)
	return _c
}

// containedInterface provides a mock function with given fields:
func (_m *mockAddressProvider) containedInterface() net.Interface {
	ret := _m.Called()

	var r0 net.Interface
	if rf, ok := ret.Get(0).(func() net.Interface); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(net.Interface)
	}

	return r0
}

// mockAddressProvider_containedInterface_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'containedInterface'
type mockAddressProvider_containedInterface_Call struct {
	*mock.Call
}

// containedInterface is a helper method to define mock.On call
func (_e *mockAddressProvider_Expecter) containedInterface() *mockAddressProvider_containedInterface_Call {
	return &mockAddressProvider_containedInterface_Call{Call: _e.mock.On("containedInterface")}
}

func (_c *mockAddressProvider_containedInterface_Call) Run(run func()) *mockAddressProvider_containedInterface_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockAddressProvider_containedInterface_Call) Return(_a0 net.Interface) *mockAddressProvider_containedInterface_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAddressProvider_containedInterface_Call) RunAndReturn(run func() net.Interface) *mockAddressProvider_containedInterface_Call {
	_c.Call.Return(run)
	return _c
}

// name provides a mock function with given fields:
func (_m *mockAddressProvider) name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// mockAddressProvider_name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'name'
type mockAddressProvider_name_Call struct {
	*mock.Call
}

// name is a helper method to define mock.On call
func (_e *mockAddressProvider_Expecter) name() *mockAddressProvider_name_Call {
	return &mockAddressProvider_name_Call{Call: _e.mock.On("name")}
}

func (_c *mockAddressProvider_name_Call) Run(run func()) *mockAddressProvider_name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockAddressProvider_name_Call) Return(_a0 string) *mockAddressProvider_name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAddressProvider_name_Call) RunAndReturn(run func() string) *mockAddressProvider_name_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTnewMockAddressProvider interface {
	mock.TestingT
	Cleanup(func())
}

// newMockAddressProvider creates a new instance of mockAddressProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockAddressProvider(t mockConstructorTestingTnewMockAddressProvider) *mockAddressProvider {
	mock := &mockAddressProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
