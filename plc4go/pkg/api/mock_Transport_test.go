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

package plc4go

import (
	url "net/url"

	transports "github.com/apache/plc4x/plc4go/spi/transports"
	mock "github.com/stretchr/testify/mock"
)

// MockTransport is an autogenerated mock type for the Transport type
type MockTransport struct {
	mock.Mock
}

type MockTransport_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransport) EXPECT() *MockTransport_Expecter {
	return &MockTransport_Expecter{mock: &_m.Mock}
}

// CreateTransportInstance provides a mock function with given fields: transportUrl, options
func (_m *MockTransport) CreateTransportInstance(transportUrl url.URL, options map[string][]string) (transports.TransportInstance, error) {
	ret := _m.Called(transportUrl, options)

	var r0 transports.TransportInstance
	var r1 error
	if rf, ok := ret.Get(0).(func(url.URL, map[string][]string) (transports.TransportInstance, error)); ok {
		return rf(transportUrl, options)
	}
	if rf, ok := ret.Get(0).(func(url.URL, map[string][]string) transports.TransportInstance); ok {
		r0 = rf(transportUrl, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transports.TransportInstance)
		}
	}

	if rf, ok := ret.Get(1).(func(url.URL, map[string][]string) error); ok {
		r1 = rf(transportUrl, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransport_CreateTransportInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransportInstance'
type MockTransport_CreateTransportInstance_Call struct {
	*mock.Call
}

// CreateTransportInstance is a helper method to define mock.On call
//   - transportUrl url.URL
//   - options map[string][]string
func (_e *MockTransport_Expecter) CreateTransportInstance(transportUrl interface{}, options interface{}) *MockTransport_CreateTransportInstance_Call {
	return &MockTransport_CreateTransportInstance_Call{Call: _e.mock.On("CreateTransportInstance", transportUrl, options)}
}

func (_c *MockTransport_CreateTransportInstance_Call) Run(run func(transportUrl url.URL, options map[string][]string)) *MockTransport_CreateTransportInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(url.URL), args[1].(map[string][]string))
	})
	return _c
}

func (_c *MockTransport_CreateTransportInstance_Call) Return(_a0 transports.TransportInstance, _a1 error) *MockTransport_CreateTransportInstance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransport_CreateTransportInstance_Call) RunAndReturn(run func(url.URL, map[string][]string) (transports.TransportInstance, error)) *MockTransport_CreateTransportInstance_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransportCode provides a mock function with given fields:
func (_m *MockTransport) GetTransportCode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTransport_GetTransportCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransportCode'
type MockTransport_GetTransportCode_Call struct {
	*mock.Call
}

// GetTransportCode is a helper method to define mock.On call
func (_e *MockTransport_Expecter) GetTransportCode() *MockTransport_GetTransportCode_Call {
	return &MockTransport_GetTransportCode_Call{Call: _e.mock.On("GetTransportCode")}
}

func (_c *MockTransport_GetTransportCode_Call) Run(run func()) *MockTransport_GetTransportCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransport_GetTransportCode_Call) Return(_a0 string) *MockTransport_GetTransportCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransport_GetTransportCode_Call) RunAndReturn(run func() string) *MockTransport_GetTransportCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransportName provides a mock function with given fields:
func (_m *MockTransport) GetTransportName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTransport_GetTransportName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransportName'
type MockTransport_GetTransportName_Call struct {
	*mock.Call
}

// GetTransportName is a helper method to define mock.On call
func (_e *MockTransport_Expecter) GetTransportName() *MockTransport_GetTransportName_Call {
	return &MockTransport_GetTransportName_Call{Call: _e.mock.On("GetTransportName")}
}

func (_c *MockTransport_GetTransportName_Call) Run(run func()) *MockTransport_GetTransportName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransport_GetTransportName_Call) Return(_a0 string) *MockTransport_GetTransportName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransport_GetTransportName_Call) RunAndReturn(run func() string) *MockTransport_GetTransportName_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockTransport interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTransport creates a new instance of MockTransport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTransport(t mockConstructorTestingTNewMockTransport) *MockTransport {
	mock := &MockTransport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
