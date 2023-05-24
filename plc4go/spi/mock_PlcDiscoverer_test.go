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

package spi

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	options "github.com/apache/plc4x/plc4go/spi/options"
)

// MockPlcDiscoverer is an autogenerated mock type for the PlcDiscoverer type
type MockPlcDiscoverer struct {
	mock.Mock
}

type MockPlcDiscoverer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcDiscoverer) EXPECT() *MockPlcDiscoverer_Expecter {
	return &MockPlcDiscoverer_Expecter{mock: &_m.Mock}
}

// Discover provides a mock function with given fields: callback, discoveryOptions
func (_m *MockPlcDiscoverer) Discover(callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption) error {
	_va := make([]interface{}, len(discoveryOptions))
	for _i := range discoveryOptions {
		_va[_i] = discoveryOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, callback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error); ok {
		r0 = rf(callback, discoveryOptions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDiscoverer_Discover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Discover'
type MockPlcDiscoverer_Discover_Call struct {
	*mock.Call
}

// Discover is a helper method to define mock.On call
//   - callback func(model.PlcDiscoveryItem)
//   - discoveryOptions ...options.WithDiscoveryOption
func (_e *MockPlcDiscoverer_Expecter) Discover(callback interface{}, discoveryOptions ...interface{}) *MockPlcDiscoverer_Discover_Call {
	return &MockPlcDiscoverer_Discover_Call{Call: _e.mock.On("Discover",
		append([]interface{}{callback}, discoveryOptions...)...)}
}

func (_c *MockPlcDiscoverer_Discover_Call) Run(run func(callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption)) *MockPlcDiscoverer_Discover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.WithDiscoveryOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(options.WithDiscoveryOption)
			}
		}
		run(args[0].(func(model.PlcDiscoveryItem)), variadicArgs...)
	})
	return _c
}

func (_c *MockPlcDiscoverer_Discover_Call) Return(_a0 error) *MockPlcDiscoverer_Discover_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDiscoverer_Discover_Call) RunAndReturn(run func(func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error) *MockPlcDiscoverer_Discover_Call {
	_c.Call.Return(run)
	return _c
}

// DiscoverWithContext provides a mock function with given fields: ctx, callback, discoveryOptions
func (_m *MockPlcDiscoverer) DiscoverWithContext(ctx context.Context, callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption) error {
	_va := make([]interface{}, len(discoveryOptions))
	for _i := range discoveryOptions {
		_va[_i] = discoveryOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, callback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error); ok {
		r0 = rf(ctx, callback, discoveryOptions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDiscoverer_DiscoverWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DiscoverWithContext'
type MockPlcDiscoverer_DiscoverWithContext_Call struct {
	*mock.Call
}

// DiscoverWithContext is a helper method to define mock.On call
//   - ctx context.Context
//   - callback func(model.PlcDiscoveryItem)
//   - discoveryOptions ...options.WithDiscoveryOption
func (_e *MockPlcDiscoverer_Expecter) DiscoverWithContext(ctx interface{}, callback interface{}, discoveryOptions ...interface{}) *MockPlcDiscoverer_DiscoverWithContext_Call {
	return &MockPlcDiscoverer_DiscoverWithContext_Call{Call: _e.mock.On("DiscoverWithContext",
		append([]interface{}{ctx, callback}, discoveryOptions...)...)}
}

func (_c *MockPlcDiscoverer_DiscoverWithContext_Call) Run(run func(ctx context.Context, callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption)) *MockPlcDiscoverer_DiscoverWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.WithDiscoveryOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.WithDiscoveryOption)
			}
		}
		run(args[0].(context.Context), args[1].(func(model.PlcDiscoveryItem)), variadicArgs...)
	})
	return _c
}

func (_c *MockPlcDiscoverer_DiscoverWithContext_Call) Return(_a0 error) *MockPlcDiscoverer_DiscoverWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDiscoverer_DiscoverWithContext_Call) RunAndReturn(run func(context.Context, func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error) *MockPlcDiscoverer_DiscoverWithContext_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcDiscoverer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcDiscoverer creates a new instance of MockPlcDiscoverer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcDiscoverer(t mockConstructorTestingTNewMockPlcDiscoverer) *MockPlcDiscoverer {
	mock := &MockPlcDiscoverer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
