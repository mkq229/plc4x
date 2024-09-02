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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CBusMessage is the corresponding interface of CBusMessage
type CBusMessage interface {
	CBusMessageContract
	CBusMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CBusMessageContract provides a set of functions which can be overwritten by a sub struct
type CBusMessageContract interface {
	// GetRequestContext() returns a parser argument
	GetRequestContext() RequestContext
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
}

// CBusMessageRequirements provides a set of functions which need to be implemented by a sub struct
type CBusMessageRequirements interface {
	// GetIsResponse returns IsResponse (discriminator field)
	GetIsResponse() bool
}

// CBusMessageExactly can be used when we want exactly this type and not a type which fulfills CBusMessage.
// This is useful for switch cases.
type CBusMessageExactly interface {
	CBusMessage
	isCBusMessage() bool
}

// _CBusMessage is the data-structure of this message
type _CBusMessage struct {
	_CBusMessageChildRequirements

	// Arguments.
	RequestContext RequestContext
	CBusOptions    CBusOptions
}

var _ CBusMessageContract = (*_CBusMessage)(nil)

type _CBusMessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetIsResponse() bool
}

type CBusMessageChild interface {
	utils.Serializable

	GetParent() *CBusMessage

	GetTypeName() string
	CBusMessage
}

// NewCBusMessage factory function for _CBusMessage
func NewCBusMessage(requestContext RequestContext, cBusOptions CBusOptions) *_CBusMessage {
	return &_CBusMessage{RequestContext: requestContext, CBusOptions: cBusOptions}
}

// Deprecated: use the interface for direct cast
func CastCBusMessage(structType any) CBusMessage {
	if casted, ok := structType.(CBusMessage); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessage); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessage) GetTypeName() string {
	return "CBusMessage"
}

func (m *_CBusMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_CBusMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusMessageParse[T CBusMessage](ctx context.Context, theBytes []byte, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (T, error) {
	return CBusMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), isResponse, requestContext, cBusOptions)
}

func CBusMessageParseWithBufferProducer[T CBusMessage](isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CBusMessageParseWithBuffer[T](ctx, readBuffer, isResponse, requestContext, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CBusMessageParseWithBuffer[T CBusMessage](ctx context.Context, readBuffer utils.ReadBuffer, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (T, error) {
	v, err := (&_CBusMessage{}).parse(ctx, readBuffer, isResponse, requestContext, cBusOptions)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CBusMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (__cBusMessage CBusMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((requestContext) != (nil))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "requestContext required"})
	}

	// Validation
	if !(bool((cBusOptions) != (nil))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "cBusOptions required"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CBusMessage
	switch {
	case isResponse == bool(false): // CBusMessageToServer
		if _child, err = (&_CBusMessageToServer{}).parse(ctx, readBuffer, m, isResponse, requestContext, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusMessageToServer for type-switch of CBusMessage")
		}
	case isResponse == bool(true): // CBusMessageToClient
		if _child, err = (&_CBusMessageToClient{}).parse(ctx, readBuffer, m, isResponse, requestContext, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusMessageToClient for type-switch of CBusMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [isResponse=%v]", isResponse)
	}

	if closeErr := readBuffer.CloseContext("CBusMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessage")
	}

	return _child, nil
}

func (pm *_CBusMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusMessage")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusMessage")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusMessage) GetRequestContext() RequestContext {
	return m.RequestContext
}
func (m *_CBusMessage) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusMessage) isCBusMessage() bool {
	return true
}
