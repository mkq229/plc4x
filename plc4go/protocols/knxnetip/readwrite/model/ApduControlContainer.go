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

// ApduControlContainer is the corresponding interface of ApduControlContainer
type ApduControlContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Apdu
	// GetControlApdu returns ControlApdu (property field)
	GetControlApdu() ApduControl
}

// ApduControlContainerExactly can be used when we want exactly this type and not a type which fulfills ApduControlContainer.
// This is useful for switch cases.
type ApduControlContainerExactly interface {
	ApduControlContainer
	isApduControlContainer() bool
}

// _ApduControlContainer is the data-structure of this message
type _ApduControlContainer struct {
	*_Apdu
	ControlApdu ApduControl
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlContainer) GetControl() uint8 {
	return uint8(1)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlContainer) InitializeParent(parent Apdu, numbered bool, counter uint8) {
	m.Numbered = numbered
	m.Counter = counter
}

func (m *_ApduControlContainer) GetParent() Apdu {
	return m._Apdu
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduControlContainer) GetControlApdu() ApduControl {
	return m.ControlApdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduControlContainer factory function for _ApduControlContainer
func NewApduControlContainer(controlApdu ApduControl, numbered bool, counter uint8, dataLength uint8) *_ApduControlContainer {
	_result := &_ApduControlContainer{
		ControlApdu: controlApdu,
		_Apdu:       NewApdu(numbered, counter, dataLength),
	}
	_result._Apdu._ApduChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduControlContainer(structType any) ApduControlContainer {
	if casted, ok := structType.(ApduControlContainer); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlContainer); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlContainer) GetTypeName() string {
	return "ApduControlContainer"
}

func (m *_ApduControlContainer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (controlApdu)
	lengthInBits += m.ControlApdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ApduControlContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduControlContainerParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduControlContainer, error) {
	return ApduControlContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduControlContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduControlContainer, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduControlContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlContainer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (controlApdu)
	if pullErr := readBuffer.PullContext("controlApdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for controlApdu")
	}
	_controlApdu, _controlApduErr := ApduControlParseWithBuffer(ctx, readBuffer)
	if _controlApduErr != nil {
		return nil, errors.Wrap(_controlApduErr, "Error parsing 'controlApdu' field of ApduControlContainer")
	}
	controlApdu := _controlApdu.(ApduControl)
	if closeErr := readBuffer.CloseContext("controlApdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for controlApdu")
	}

	if closeErr := readBuffer.CloseContext("ApduControlContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlContainer")
	}

	// Create a partially initialized instance
	_child := &_ApduControlContainer{
		_Apdu: &_Apdu{
			DataLength: dataLength,
		},
		ControlApdu: controlApdu,
	}
	_child._Apdu._ApduChildRequirements = _child
	return _child, nil
}

func (m *_ApduControlContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlContainer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlContainer")
		}

		// Simple Field (controlApdu)
		if pushErr := writeBuffer.PushContext("controlApdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for controlApdu")
		}
		_controlApduErr := writeBuffer.WriteSerializable(ctx, m.GetControlApdu())
		if popErr := writeBuffer.PopContext("controlApdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for controlApdu")
		}
		if _controlApduErr != nil {
			return errors.Wrap(_controlApduErr, "Error serializing 'controlApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduControlContainer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlContainer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduControlContainer) isApduControlContainer() bool {
	return true
}

func (m *_ApduControlContainer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
