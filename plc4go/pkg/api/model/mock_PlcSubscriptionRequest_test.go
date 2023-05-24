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

package model

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockPlcSubscriptionRequest is an autogenerated mock type for the PlcSubscriptionRequest type
type MockPlcSubscriptionRequest struct {
	mock.Mock
}

type MockPlcSubscriptionRequest_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcSubscriptionRequest) EXPECT() *MockPlcSubscriptionRequest_Expecter {
	return &MockPlcSubscriptionRequest_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequest) Execute() <-chan PlcSubscriptionRequestResult {
	ret := _m.Called()

	var r0 <-chan PlcSubscriptionRequestResult
	if rf, ok := ret.Get(0).(func() <-chan PlcSubscriptionRequestResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcSubscriptionRequestResult)
		}
	}

	return r0
}

// MockPlcSubscriptionRequest_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPlcSubscriptionRequest_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequest_Expecter) Execute() *MockPlcSubscriptionRequest_Execute_Call {
	return &MockPlcSubscriptionRequest_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockPlcSubscriptionRequest_Execute_Call) Run(run func()) *MockPlcSubscriptionRequest_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequest_Execute_Call) Return(_a0 <-chan PlcSubscriptionRequestResult) *MockPlcSubscriptionRequest_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequest_Execute_Call) RunAndReturn(run func() <-chan PlcSubscriptionRequestResult) *MockPlcSubscriptionRequest_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteWithContext provides a mock function with given fields: ctx
func (_m *MockPlcSubscriptionRequest) ExecuteWithContext(ctx context.Context) <-chan PlcSubscriptionRequestResult {
	ret := _m.Called(ctx)

	var r0 <-chan PlcSubscriptionRequestResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan PlcSubscriptionRequestResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcSubscriptionRequestResult)
		}
	}

	return r0
}

// MockPlcSubscriptionRequest_ExecuteWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteWithContext'
type MockPlcSubscriptionRequest_ExecuteWithContext_Call struct {
	*mock.Call
}

// ExecuteWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPlcSubscriptionRequest_Expecter) ExecuteWithContext(ctx interface{}) *MockPlcSubscriptionRequest_ExecuteWithContext_Call {
	return &MockPlcSubscriptionRequest_ExecuteWithContext_Call{Call: _e.mock.On("ExecuteWithContext", ctx)}
}

func (_c *MockPlcSubscriptionRequest_ExecuteWithContext_Call) Run(run func(ctx context.Context)) *MockPlcSubscriptionRequest_ExecuteWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPlcSubscriptionRequest_ExecuteWithContext_Call) Return(_a0 <-chan PlcSubscriptionRequestResult) *MockPlcSubscriptionRequest_ExecuteWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequest_ExecuteWithContext_Call) RunAndReturn(run func(context.Context) <-chan PlcSubscriptionRequestResult) *MockPlcSubscriptionRequest_ExecuteWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetTag provides a mock function with given fields: tagName
func (_m *MockPlcSubscriptionRequest) GetTag(tagName string) PlcTag {
	ret := _m.Called(tagName)

	var r0 PlcTag
	if rf, ok := ret.Get(0).(func(string) PlcTag); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcTag)
		}
	}

	return r0
}

// MockPlcSubscriptionRequest_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockPlcSubscriptionRequest_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcSubscriptionRequest_Expecter) GetTag(tagName interface{}) *MockPlcSubscriptionRequest_GetTag_Call {
	return &MockPlcSubscriptionRequest_GetTag_Call{Call: _e.mock.On("GetTag", tagName)}
}

func (_c *MockPlcSubscriptionRequest_GetTag_Call) Run(run func(tagName string)) *MockPlcSubscriptionRequest_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionRequest_GetTag_Call) Return(_a0 PlcTag) *MockPlcSubscriptionRequest_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequest_GetTag_Call) RunAndReturn(run func(string) PlcTag) *MockPlcSubscriptionRequest_GetTag_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequest) GetTagNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockPlcSubscriptionRequest_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcSubscriptionRequest_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequest_Expecter) GetTagNames() *MockPlcSubscriptionRequest_GetTagNames_Call {
	return &MockPlcSubscriptionRequest_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcSubscriptionRequest_GetTagNames_Call) Run(run func()) *MockPlcSubscriptionRequest_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequest_GetTagNames_Call) Return(_a0 []string) *MockPlcSubscriptionRequest_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequest_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcSubscriptionRequest_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequest) IsAPlcMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcSubscriptionRequest_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcSubscriptionRequest_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequest_Expecter) IsAPlcMessage() *MockPlcSubscriptionRequest_IsAPlcMessage_Call {
	return &MockPlcSubscriptionRequest_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcSubscriptionRequest_IsAPlcMessage_Call) Run(run func()) *MockPlcSubscriptionRequest_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequest_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcSubscriptionRequest_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequest_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcSubscriptionRequest_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequest) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcSubscriptionRequest_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcSubscriptionRequest_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequest_Expecter) String() *MockPlcSubscriptionRequest_String_Call {
	return &MockPlcSubscriptionRequest_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcSubscriptionRequest_String_Call) Run(run func()) *MockPlcSubscriptionRequest_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequest_String_Call) Return(_a0 string) *MockPlcSubscriptionRequest_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequest_String_Call) RunAndReturn(run func() string) *MockPlcSubscriptionRequest_String_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcSubscriptionRequest interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcSubscriptionRequest creates a new instance of MockPlcSubscriptionRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcSubscriptionRequest(t mockConstructorTestingTNewMockPlcSubscriptionRequest) *MockPlcSubscriptionRequest {
	mock := &MockPlcSubscriptionRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
