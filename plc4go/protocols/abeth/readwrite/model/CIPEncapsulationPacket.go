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

// CIPEncapsulationPacket is the corresponding interface of CIPEncapsulationPacket
type CIPEncapsulationPacket interface {
	CIPEncapsulationPacketContract
	CIPEncapsulationPacketRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CIPEncapsulationPacketContract provides a set of functions which can be overwritten by a sub struct
type CIPEncapsulationPacketContract interface {
	// GetSessionHandle returns SessionHandle (property field)
	GetSessionHandle() uint32
	// GetStatus returns Status (property field)
	GetStatus() uint32
	// GetSenderContext returns SenderContext (property field)
	GetSenderContext() []uint8
	// GetOptions returns Options (property field)
	GetOptions() uint32
}

// CIPEncapsulationPacketRequirements provides a set of functions which need to be implemented by a sub struct
type CIPEncapsulationPacketRequirements interface {
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() uint16
}

// CIPEncapsulationPacketExactly can be used when we want exactly this type and not a type which fulfills CIPEncapsulationPacket.
// This is useful for switch cases.
type CIPEncapsulationPacketExactly interface {
	CIPEncapsulationPacket
	isCIPEncapsulationPacket() bool
}

// _CIPEncapsulationPacket is the data-structure of this message
type _CIPEncapsulationPacket struct {
	_CIPEncapsulationPacketChildRequirements
	SessionHandle uint32
	Status        uint32
	SenderContext []uint8
	Options       uint32
	// Reserved Fields
	reservedField0 *uint32
}

var _ CIPEncapsulationPacketContract = (*_CIPEncapsulationPacket)(nil)

type _CIPEncapsulationPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() uint16
}

type CIPEncapsulationPacketChild interface {
	utils.Serializable

	GetParent() *CIPEncapsulationPacket

	GetTypeName() string
	CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationPacket) GetSessionHandle() uint32 {
	return m.SessionHandle
}

func (m *_CIPEncapsulationPacket) GetStatus() uint32 {
	return m.Status
}

func (m *_CIPEncapsulationPacket) GetSenderContext() []uint8 {
	return m.SenderContext
}

func (m *_CIPEncapsulationPacket) GetOptions() uint32 {
	return m.Options
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPEncapsulationPacket factory function for _CIPEncapsulationPacket
func NewCIPEncapsulationPacket(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_CIPEncapsulationPacket {
	return &_CIPEncapsulationPacket{SessionHandle: sessionHandle, Status: status, SenderContext: senderContext, Options: options}
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationPacket(structType any) CIPEncapsulationPacket {
	if casted, ok := structType.(CIPEncapsulationPacket); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationPacket); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationPacket) GetTypeName() string {
	return "CIPEncapsulationPacket"
}

func (m *_CIPEncapsulationPacket) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandType)
	lengthInBits += 16

	// Implicit Field (packetLen)
	lengthInBits += 16

	// Simple field (sessionHandle)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += 32

	// Array field
	if len(m.SenderContext) > 0 {
		lengthInBits += 8 * uint16(len(m.SenderContext))
	}

	// Simple field (options)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CIPEncapsulationPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPEncapsulationPacketParse[T CIPEncapsulationPacket](ctx context.Context, theBytes []byte) (T, error) {
	return CIPEncapsulationPacketParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func CIPEncapsulationPacketParseWithBufferProducer[T CIPEncapsulationPacket]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CIPEncapsulationPacketParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CIPEncapsulationPacketParseWithBuffer[T CIPEncapsulationPacket](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_CIPEncapsulationPacket{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CIPEncapsulationPacket) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__cIPEncapsulationPacket CIPEncapsulationPacket, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	commandType, err := ReadDiscriminatorField[uint16](ctx, "commandType", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}

	packetLen, err := ReadImplicitField[uint16](ctx, "packetLen", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'packetLen' field"))
	}
	_ = packetLen

	sessionHandle, err := ReadSimpleField(ctx, "sessionHandle", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionHandle' field"))
	}
	m.SessionHandle = sessionHandle

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	senderContext, err := ReadCountArrayField[uint8](ctx, "senderContext", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'senderContext' field"))
	}
	m.SenderContext = senderContext

	options, err := ReadSimpleField(ctx, "options", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'options' field"))
	}
	m.Options = options

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedInt(readBuffer, uint8(32)), uint32(0x00000000), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CIPEncapsulationPacket
	switch {
	case commandType == 0x0101: // CIPEncapsulationConnectionRequest
		if _child, err = (&_CIPEncapsulationConnectionRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CIPEncapsulationConnectionRequest for type-switch of CIPEncapsulationPacket")
		}
	case commandType == 0x0201: // CIPEncapsulationConnectionResponse
		if _child, err = (&_CIPEncapsulationConnectionResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CIPEncapsulationConnectionResponse for type-switch of CIPEncapsulationPacket")
		}
	case commandType == 0x0107: // CIPEncapsulationReadRequest
		if _child, err = (&_CIPEncapsulationReadRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CIPEncapsulationReadRequest for type-switch of CIPEncapsulationPacket")
		}
	case commandType == 0x0207: // CIPEncapsulationReadResponse
		if _child, err = (&_CIPEncapsulationReadResponse{}).parse(ctx, readBuffer, m, packetLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CIPEncapsulationReadResponse for type-switch of CIPEncapsulationPacket")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationPacket")
	}

	return _child, nil
}

func (pm *_CIPEncapsulationPacket) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CIPEncapsulationPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPEncapsulationPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationPacket")
	}

	if err := WriteDiscriminatorField(ctx, "commandType", m.GetCommandType(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commandType' field")
	}
	packetLen := uint16(uint16(uint16(m.GetLengthInBytes(ctx))) - uint16(uint16(28)))
	if err := WriteImplicitField(ctx, "packetLen", packetLen, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'packetLen' field")
	}

	if err := WriteSimpleField[uint32](ctx, "sessionHandle", m.GetSessionHandle(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'sessionHandle' field")
	}

	if err := WriteSimpleField[uint32](ctx, "status", m.GetStatus(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "senderContext", m.GetSenderContext(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'senderContext' field")
	}

	if err := WriteSimpleField[uint32](ctx, "options", m.GetOptions(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'options' field")
	}

	if err := WriteReservedField[uint32](ctx, "reserved", uint32(0x00000000), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CIPEncapsulationPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPEncapsulationPacket")
	}
	return nil
}

func (m *_CIPEncapsulationPacket) isCIPEncapsulationPacket() bool {
	return true
}
