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

package bacnetip

import mock "github.com/stretchr/testify/mock"

// mock_IOCB is an autogenerated mock type for the _IOCB type
type mock_IOCB struct {
	mock.Mock
}

type mock_IOCB_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_IOCB) EXPECT() *mock_IOCB_Expecter {
	return &mock_IOCB_Expecter{mock: &_m.Mock}
}

// Abort provides a mock function with given fields: err
func (_m *mock_IOCB) Abort(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_IOCB_Abort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Abort'
type mock_IOCB_Abort_Call struct {
	*mock.Call
}

// Abort is a helper method to define mock.On call
//   - err error
func (_e *mock_IOCB_Expecter) Abort(err interface{}) *mock_IOCB_Abort_Call {
	return &mock_IOCB_Abort_Call{Call: _e.mock.On("Abort", err)}
}

func (_c *mock_IOCB_Abort_Call) Run(run func(err error)) *mock_IOCB_Abort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *mock_IOCB_Abort_Call) Return(_a0 error) *mock_IOCB_Abort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOCB_Abort_Call) RunAndReturn(run func(error) error) *mock_IOCB_Abort_Call {
	_c.Call.Return(run)
	return _c
}

// Trigger provides a mock function with given fields:
func (_m *mock_IOCB) Trigger() {
	_m.Called()
}

// mock_IOCB_Trigger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trigger'
type mock_IOCB_Trigger_Call struct {
	*mock.Call
}

// Trigger is a helper method to define mock.On call
func (_e *mock_IOCB_Expecter) Trigger() *mock_IOCB_Trigger_Call {
	return &mock_IOCB_Trigger_Call{Call: _e.mock.On("Trigger")}
}

func (_c *mock_IOCB_Trigger_Call) Run(run func()) *mock_IOCB_Trigger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_IOCB_Trigger_Call) Return() *mock_IOCB_Trigger_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_IOCB_Trigger_Call) RunAndReturn(run func()) *mock_IOCB_Trigger_Call {
	_c.Call.Return(run)
	return _c
}

// clearQueue provides a mock function with given fields:
func (_m *mock_IOCB) clearQueue() {
	_m.Called()
}

// mock_IOCB_clearQueue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'clearQueue'
type mock_IOCB_clearQueue_Call struct {
	*mock.Call
}

// clearQueue is a helper method to define mock.On call
func (_e *mock_IOCB_Expecter) clearQueue() *mock_IOCB_clearQueue_Call {
	return &mock_IOCB_clearQueue_Call{Call: _e.mock.On("clearQueue")}
}

func (_c *mock_IOCB_clearQueue_Call) Run(run func()) *mock_IOCB_clearQueue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_IOCB_clearQueue_Call) Return() *mock_IOCB_clearQueue_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_IOCB_clearQueue_Call) RunAndReturn(run func()) *mock_IOCB_clearQueue_Call {
	_c.Call.Return(run)
	return _c
}

// getDestination provides a mock function with given fields:
func (_m *mock_IOCB) getDestination() *Address {
	ret := _m.Called()

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// mock_IOCB_getDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getDestination'
type mock_IOCB_getDestination_Call struct {
	*mock.Call
}

// getDestination is a helper method to define mock.On call
func (_e *mock_IOCB_Expecter) getDestination() *mock_IOCB_getDestination_Call {
	return &mock_IOCB_getDestination_Call{Call: _e.mock.On("getDestination")}
}

func (_c *mock_IOCB_getDestination_Call) Run(run func()) *mock_IOCB_getDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_IOCB_getDestination_Call) Return(_a0 *Address) *mock_IOCB_getDestination_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOCB_getDestination_Call) RunAndReturn(run func() *Address) *mock_IOCB_getDestination_Call {
	_c.Call.Return(run)
	return _c
}

// getIOState provides a mock function with given fields:
func (_m *mock_IOCB) getIOState() IOCBState {
	ret := _m.Called()

	var r0 IOCBState
	if rf, ok := ret.Get(0).(func() IOCBState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(IOCBState)
	}

	return r0
}

// mock_IOCB_getIOState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getIOState'
type mock_IOCB_getIOState_Call struct {
	*mock.Call
}

// getIOState is a helper method to define mock.On call
func (_e *mock_IOCB_Expecter) getIOState() *mock_IOCB_getIOState_Call {
	return &mock_IOCB_getIOState_Call{Call: _e.mock.On("getIOState")}
}

func (_c *mock_IOCB_getIOState_Call) Run(run func()) *mock_IOCB_getIOState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_IOCB_getIOState_Call) Return(_a0 IOCBState) *mock_IOCB_getIOState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOCB_getIOState_Call) RunAndReturn(run func() IOCBState) *mock_IOCB_getIOState_Call {
	_c.Call.Return(run)
	return _c
}

// getPriority provides a mock function with given fields:
func (_m *mock_IOCB) getPriority() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// mock_IOCB_getPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getPriority'
type mock_IOCB_getPriority_Call struct {
	*mock.Call
}

// getPriority is a helper method to define mock.On call
func (_e *mock_IOCB_Expecter) getPriority() *mock_IOCB_getPriority_Call {
	return &mock_IOCB_getPriority_Call{Call: _e.mock.On("getPriority")}
}

func (_c *mock_IOCB_getPriority_Call) Run(run func()) *mock_IOCB_getPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_IOCB_getPriority_Call) Return(_a0 int) *mock_IOCB_getPriority_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOCB_getPriority_Call) RunAndReturn(run func() int) *mock_IOCB_getPriority_Call {
	_c.Call.Return(run)
	return _c
}

// getRequest provides a mock function with given fields:
func (_m *mock_IOCB) getRequest() _PDU {
	ret := _m.Called()

	var r0 _PDU
	if rf, ok := ret.Get(0).(func() _PDU); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(_PDU)
		}
	}

	return r0
}

// mock_IOCB_getRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getRequest'
type mock_IOCB_getRequest_Call struct {
	*mock.Call
}

// getRequest is a helper method to define mock.On call
func (_e *mock_IOCB_Expecter) getRequest() *mock_IOCB_getRequest_Call {
	return &mock_IOCB_getRequest_Call{Call: _e.mock.On("getRequest")}
}

func (_c *mock_IOCB_getRequest_Call) Run(run func()) *mock_IOCB_getRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_IOCB_getRequest_Call) Return(_a0 _PDU) *mock_IOCB_getRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_IOCB_getRequest_Call) RunAndReturn(run func() _PDU) *mock_IOCB_getRequest_Call {
	_c.Call.Return(run)
	return _c
}

// setIOController provides a mock function with given fields: ioController
func (_m *mock_IOCB) setIOController(ioController _IOController) {
	_m.Called(ioController)
}

// mock_IOCB_setIOController_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setIOController'
type mock_IOCB_setIOController_Call struct {
	*mock.Call
}

// setIOController is a helper method to define mock.On call
//   - ioController _IOController
func (_e *mock_IOCB_Expecter) setIOController(ioController interface{}) *mock_IOCB_setIOController_Call {
	return &mock_IOCB_setIOController_Call{Call: _e.mock.On("setIOController", ioController)}
}

func (_c *mock_IOCB_setIOController_Call) Run(run func(ioController _IOController)) *mock_IOCB_setIOController_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOController))
	})
	return _c
}

func (_c *mock_IOCB_setIOController_Call) Return() *mock_IOCB_setIOController_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_IOCB_setIOController_Call) RunAndReturn(run func(_IOController)) *mock_IOCB_setIOController_Call {
	_c.Call.Return(run)
	return _c
}

// setIOError provides a mock function with given fields: err
func (_m *mock_IOCB) setIOError(err error) {
	_m.Called(err)
}

// mock_IOCB_setIOError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setIOError'
type mock_IOCB_setIOError_Call struct {
	*mock.Call
}

// setIOError is a helper method to define mock.On call
//   - err error
func (_e *mock_IOCB_Expecter) setIOError(err interface{}) *mock_IOCB_setIOError_Call {
	return &mock_IOCB_setIOError_Call{Call: _e.mock.On("setIOError", err)}
}

func (_c *mock_IOCB_setIOError_Call) Run(run func(err error)) *mock_IOCB_setIOError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *mock_IOCB_setIOError_Call) Return() *mock_IOCB_setIOError_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_IOCB_setIOError_Call) RunAndReturn(run func(error)) *mock_IOCB_setIOError_Call {
	_c.Call.Return(run)
	return _c
}

// setIOResponse provides a mock function with given fields: msg
func (_m *mock_IOCB) setIOResponse(msg _PDU) {
	_m.Called(msg)
}

// mock_IOCB_setIOResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setIOResponse'
type mock_IOCB_setIOResponse_Call struct {
	*mock.Call
}

// setIOResponse is a helper method to define mock.On call
//   - msg _PDU
func (_e *mock_IOCB_Expecter) setIOResponse(msg interface{}) *mock_IOCB_setIOResponse_Call {
	return &mock_IOCB_setIOResponse_Call{Call: _e.mock.On("setIOResponse", msg)}
}

func (_c *mock_IOCB_setIOResponse_Call) Run(run func(msg _PDU)) *mock_IOCB_setIOResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_IOCB_setIOResponse_Call) Return() *mock_IOCB_setIOResponse_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_IOCB_setIOResponse_Call) RunAndReturn(run func(_PDU)) *mock_IOCB_setIOResponse_Call {
	_c.Call.Return(run)
	return _c
}

// setIOState provides a mock function with given fields: newState
func (_m *mock_IOCB) setIOState(newState IOCBState) {
	_m.Called(newState)
}

// mock_IOCB_setIOState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setIOState'
type mock_IOCB_setIOState_Call struct {
	*mock.Call
}

// setIOState is a helper method to define mock.On call
//   - newState IOCBState
func (_e *mock_IOCB_Expecter) setIOState(newState interface{}) *mock_IOCB_setIOState_Call {
	return &mock_IOCB_setIOState_Call{Call: _e.mock.On("setIOState", newState)}
}

func (_c *mock_IOCB_setIOState_Call) Run(run func(newState IOCBState)) *mock_IOCB_setIOState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(IOCBState))
	})
	return _c
}

func (_c *mock_IOCB_setIOState_Call) Return() *mock_IOCB_setIOState_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_IOCB_setIOState_Call) RunAndReturn(run func(IOCBState)) *mock_IOCB_setIOState_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTnewMock_IOCB interface {
	mock.TestingT
	Cleanup(func())
}

// newMock_IOCB creates a new instance of mock_IOCB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMock_IOCB(t mockConstructorTestingTnewMock_IOCB) *mock_IOCB {
	mock := &mock_IOCB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
