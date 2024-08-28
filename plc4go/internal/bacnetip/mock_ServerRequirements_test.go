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

// Code generated by mockery v2.42.2. DO NOT EDIT.

package bacnetip

import mock "github.com/stretchr/testify/mock"

// MockServerRequirements is an autogenerated mock type for the ServerRequirements type
type MockServerRequirements struct {
	mock.Mock
}

type MockServerRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServerRequirements) EXPECT() *MockServerRequirements_Expecter {
	return &MockServerRequirements_Expecter{mock: &_m.Mock}
}

// NewMockServerRequirements creates a new instance of MockServerRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServerRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServerRequirements {
	mock := &MockServerRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}