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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7ParameterSetupCommunication is the corresponding interface of S7ParameterSetupCommunication
type S7ParameterSetupCommunication interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetMaxAmqCaller returns MaxAmqCaller (property field)
	GetMaxAmqCaller() uint16
	// GetMaxAmqCallee returns MaxAmqCallee (property field)
	GetMaxAmqCallee() uint16
	// GetPduLength returns PduLength (property field)
	GetPduLength() uint16
}

// S7ParameterSetupCommunicationExactly can be used when we want exactly this type and not a type which fulfills S7ParameterSetupCommunication.
// This is useful for switch cases.
type S7ParameterSetupCommunicationExactly interface {
	S7ParameterSetupCommunication
	isS7ParameterSetupCommunication() bool
}

// _S7ParameterSetupCommunication is the data-structure of this message
type _S7ParameterSetupCommunication struct {
	S7ParameterContract
	MaxAmqCaller uint16
	MaxAmqCallee uint16
	PduLength    uint16
	// Reserved Fields
	reservedField0 *uint8
}

var _ S7ParameterSetupCommunication = (*_S7ParameterSetupCommunication)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterSetupCommunication) GetParameterType() uint8 {
	return 0xF0
}

func (m *_S7ParameterSetupCommunication) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterSetupCommunication) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterSetupCommunication) GetMaxAmqCaller() uint16 {
	return m.MaxAmqCaller
}

func (m *_S7ParameterSetupCommunication) GetMaxAmqCallee() uint16 {
	return m.MaxAmqCallee
}

func (m *_S7ParameterSetupCommunication) GetPduLength() uint16 {
	return m.PduLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterSetupCommunication factory function for _S7ParameterSetupCommunication
func NewS7ParameterSetupCommunication(maxAmqCaller uint16, maxAmqCallee uint16, pduLength uint16) *_S7ParameterSetupCommunication {
	_result := &_S7ParameterSetupCommunication{
		S7ParameterContract: NewS7Parameter(),
		MaxAmqCaller:        maxAmqCaller,
		MaxAmqCallee:        maxAmqCallee,
		PduLength:           pduLength,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterSetupCommunication(structType any) S7ParameterSetupCommunication {
	if casted, ok := structType.(S7ParameterSetupCommunication); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterSetupCommunication); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterSetupCommunication) GetTypeName() string {
	return "S7ParameterSetupCommunication"
}

func (m *_S7ParameterSetupCommunication) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (maxAmqCaller)
	lengthInBits += 16

	// Simple field (maxAmqCallee)
	lengthInBits += 16

	// Simple field (pduLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_S7ParameterSetupCommunication) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterSetupCommunication) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterSetupCommunication S7ParameterSetupCommunication, err error) {
	m.S7ParameterContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterSetupCommunication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterSetupCommunication")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	maxAmqCaller, err := ReadSimpleField(ctx, "maxAmqCaller", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxAmqCaller' field"))
	}
	m.MaxAmqCaller = maxAmqCaller

	maxAmqCallee, err := ReadSimpleField(ctx, "maxAmqCallee", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxAmqCallee' field"))
	}
	m.MaxAmqCallee = maxAmqCallee

	pduLength, err := ReadSimpleField(ctx, "pduLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pduLength' field"))
	}
	m.PduLength = pduLength

	if closeErr := readBuffer.CloseContext("S7ParameterSetupCommunication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterSetupCommunication")
	}

	return m, nil
}

func (m *_S7ParameterSetupCommunication) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterSetupCommunication) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterSetupCommunication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterSetupCommunication")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint16](ctx, "maxAmqCaller", m.GetMaxAmqCaller(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxAmqCaller' field")
		}

		if err := WriteSimpleField[uint16](ctx, "maxAmqCallee", m.GetMaxAmqCallee(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxAmqCallee' field")
		}

		if err := WriteSimpleField[uint16](ctx, "pduLength", m.GetPduLength(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'pduLength' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterSetupCommunication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterSetupCommunication")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterSetupCommunication) isS7ParameterSetupCommunication() bool {
	return true
}

func (m *_S7ParameterSetupCommunication) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
