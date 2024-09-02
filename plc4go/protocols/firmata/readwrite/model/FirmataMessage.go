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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessage is the corresponding interface of FirmataMessage
type FirmataMessage interface {
	FirmataMessageContract
	FirmataMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// FirmataMessageContract provides a set of functions which can be overwritten by a sub struct
type FirmataMessageContract interface {
	// GetResponse() returns a parser argument
	GetResponse() bool
}

// FirmataMessageRequirements provides a set of functions which need to be implemented by a sub struct
type FirmataMessageRequirements interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
}

// FirmataMessageExactly can be used when we want exactly this type and not a type which fulfills FirmataMessage.
// This is useful for switch cases.
type FirmataMessageExactly interface {
	FirmataMessage
	isFirmataMessage() bool
}

// _FirmataMessage is the data-structure of this message
type _FirmataMessage struct {
	_FirmataMessageChildRequirements

	// Arguments.
	Response bool
}

var _ FirmataMessageContract = (*_FirmataMessage)(nil)

type _FirmataMessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() uint8
}

type FirmataMessageChild interface {
	utils.Serializable

	GetParent() *FirmataMessage

	GetTypeName() string
	FirmataMessage
}

// NewFirmataMessage factory function for _FirmataMessage
func NewFirmataMessage(response bool) *_FirmataMessage {
	return &_FirmataMessage{Response: response}
}

// Deprecated: use the interface for direct cast
func CastFirmataMessage(structType any) FirmataMessage {
	if casted, ok := structType.(FirmataMessage); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessage); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessage) GetTypeName() string {
	return "FirmataMessage"
}

func (m *_FirmataMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_FirmataMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataMessageParse[T FirmataMessage](ctx context.Context, theBytes []byte, response bool) (T, error) {
	return FirmataMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response)
}

func FirmataMessageParseWithBufferProducer[T FirmataMessage](response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := FirmataMessageParseWithBuffer[T](ctx, readBuffer, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func FirmataMessageParseWithBuffer[T FirmataMessage](ctx context.Context, readBuffer utils.ReadBuffer, response bool) (T, error) {
	v, err := (&_FirmataMessage{}).parse(ctx, readBuffer, response)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_FirmataMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (__firmataMessage FirmataMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageType, err := ReadDiscriminatorField[uint8](ctx, "messageType", ReadUnsignedByte(readBuffer, uint8(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child FirmataMessage
	switch {
	case messageType == 0xE: // FirmataMessageAnalogIO
		if _child, err = (&_FirmataMessageAnalogIO{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type FirmataMessageAnalogIO for type-switch of FirmataMessage")
		}
	case messageType == 0x9: // FirmataMessageDigitalIO
		if _child, err = (&_FirmataMessageDigitalIO{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type FirmataMessageDigitalIO for type-switch of FirmataMessage")
		}
	case messageType == 0xC: // FirmataMessageSubscribeAnalogPinValue
		if _child, err = (&_FirmataMessageSubscribeAnalogPinValue{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type FirmataMessageSubscribeAnalogPinValue for type-switch of FirmataMessage")
		}
	case messageType == 0xD: // FirmataMessageSubscribeDigitalPinValue
		if _child, err = (&_FirmataMessageSubscribeDigitalPinValue{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type FirmataMessageSubscribeDigitalPinValue for type-switch of FirmataMessage")
		}
	case messageType == 0xF: // FirmataMessageCommand
		if _child, err = (&_FirmataMessageCommand{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type FirmataMessageCommand for type-switch of FirmataMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageType=%v]", messageType)
	}

	if closeErr := readBuffer.CloseContext("FirmataMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessage")
	}

	return _child, nil
}

func (pm *_FirmataMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child FirmataMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("FirmataMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for FirmataMessage")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteUnsignedByte(writeBuffer, 4), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("FirmataMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for FirmataMessage")
	}
	return nil
}

////
// Arguments Getter

func (m *_FirmataMessage) GetResponse() bool {
	return m.Response
}

//
////

func (m *_FirmataMessage) isFirmataMessage() bool {
	return true
}
