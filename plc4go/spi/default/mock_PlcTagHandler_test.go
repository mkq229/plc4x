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

package _default

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcTagHandler is an autogenerated mock type for the PlcTagHandler type
type MockPlcTagHandler struct {
	mock.Mock
}

type MockPlcTagHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcTagHandler) EXPECT() *MockPlcTagHandler_Expecter {
	return &MockPlcTagHandler_Expecter{mock: &_m.Mock}
}

// ParseQuery provides a mock function with given fields: query
func (_m *MockPlcTagHandler) ParseQuery(query string) (model.PlcQuery, error) {
	ret := _m.Called(query)

	var r0 model.PlcQuery
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.PlcQuery, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) model.PlcQuery); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcQuery)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcTagHandler_ParseQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseQuery'
type MockPlcTagHandler_ParseQuery_Call struct {
	*mock.Call
}

// ParseQuery is a helper method to define mock.On call
//   - query string
func (_e *MockPlcTagHandler_Expecter) ParseQuery(query interface{}) *MockPlcTagHandler_ParseQuery_Call {
	return &MockPlcTagHandler_ParseQuery_Call{Call: _e.mock.On("ParseQuery", query)}
}

func (_c *MockPlcTagHandler_ParseQuery_Call) Run(run func(query string)) *MockPlcTagHandler_ParseQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcTagHandler_ParseQuery_Call) Return(_a0 model.PlcQuery, _a1 error) *MockPlcTagHandler_ParseQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcTagHandler_ParseQuery_Call) RunAndReturn(run func(string) (model.PlcQuery, error)) *MockPlcTagHandler_ParseQuery_Call {
	_c.Call.Return(run)
	return _c
}

// ParseTag provides a mock function with given fields: tagAddress
func (_m *MockPlcTagHandler) ParseTag(tagAddress string) (model.PlcTag, error) {
	ret := _m.Called(tagAddress)

	var r0 model.PlcTag
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.PlcTag, error)); ok {
		return rf(tagAddress)
	}
	if rf, ok := ret.Get(0).(func(string) model.PlcTag); ok {
		r0 = rf(tagAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcTag)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tagAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcTagHandler_ParseTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseTag'
type MockPlcTagHandler_ParseTag_Call struct {
	*mock.Call
}

// ParseTag is a helper method to define mock.On call
//   - tagAddress string
func (_e *MockPlcTagHandler_Expecter) ParseTag(tagAddress interface{}) *MockPlcTagHandler_ParseTag_Call {
	return &MockPlcTagHandler_ParseTag_Call{Call: _e.mock.On("ParseTag", tagAddress)}
}

func (_c *MockPlcTagHandler_ParseTag_Call) Run(run func(tagAddress string)) *MockPlcTagHandler_ParseTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcTagHandler_ParseTag_Call) Return(_a0 model.PlcTag, _a1 error) *MockPlcTagHandler_ParseTag_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcTagHandler_ParseTag_Call) RunAndReturn(run func(string) (model.PlcTag, error)) *MockPlcTagHandler_ParseTag_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcTagHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcTagHandler creates a new instance of MockPlcTagHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcTagHandler(t mockConstructorTestingTNewMockPlcTagHandler) *MockPlcTagHandler {
	mock := &MockPlcTagHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
