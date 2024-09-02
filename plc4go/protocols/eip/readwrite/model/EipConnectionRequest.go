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

// Constant values.
const EipConnectionRequest_PROTOCOLVERSION uint16 = 0x01
const EipConnectionRequest_FLAGS uint16 = 0x00

// EipConnectionRequest is the corresponding interface of EipConnectionRequest
type EipConnectionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// EipConnectionRequestExactly can be used when we want exactly this type and not a type which fulfills EipConnectionRequest.
// This is useful for switch cases.
type EipConnectionRequestExactly interface {
	EipConnectionRequest
	isEipConnectionRequest() bool
}

// _EipConnectionRequest is the data-structure of this message
type _EipConnectionRequest struct {
	EipPacketContract
}

var _ EipConnectionRequest = (*_EipConnectionRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipConnectionRequest) GetCommand() uint16 {
	return 0x0065
}

func (m *_EipConnectionRequest) GetResponse() bool {
	return bool(false)
}

func (m *_EipConnectionRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipConnectionRequest) GetParent() EipPacketContract {
	return m.EipPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_EipConnectionRequest) GetProtocolVersion() uint16 {
	return EipConnectionRequest_PROTOCOLVERSION
}

func (m *_EipConnectionRequest) GetFlags() uint16 {
	return EipConnectionRequest_FLAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEipConnectionRequest factory function for _EipConnectionRequest
func NewEipConnectionRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipConnectionRequest {
	_result := &_EipConnectionRequest{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastEipConnectionRequest(structType any) EipConnectionRequest {
	if casted, ok := structType.(EipConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipConnectionRequest) GetTypeName() string {
	return "EipConnectionRequest"
}

func (m *_EipConnectionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	// Const Field (protocolVersion)
	lengthInBits += 16

	// Const Field (flags)
	lengthInBits += 16

	return lengthInBits
}

func (m *_EipConnectionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EipConnectionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__eipConnectionRequest EipConnectionRequest, err error) {
	m.EipPacketContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolVersion, err := ReadConstField[uint16](ctx, "protocolVersion", ReadUnsignedShort(readBuffer, uint8(16)), EipConnectionRequest_PROTOCOLVERSION)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersion' field"))
	}
	_ = protocolVersion

	flags, err := ReadConstField[uint16](ctx, "flags", ReadUnsignedShort(readBuffer, uint8(16)), EipConnectionRequest_FLAGS)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flags' field"))
	}
	_ = flags

	if closeErr := readBuffer.CloseContext("EipConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipConnectionRequest")
	}

	return m, nil
}

func (m *_EipConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipConnectionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipConnectionRequest")
		}

		if err := WriteConstField(ctx, "protocolVersion", EipConnectionRequest_PROTOCOLVERSION, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolVersion' field")
		}

		if err := WriteConstField(ctx, "flags", EipConnectionRequest_FLAGS, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'flags' field")
		}

		if popErr := writeBuffer.PopContext("EipConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipConnectionRequest")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipConnectionRequest) isEipConnectionRequest() bool {
	return true
}

func (m *_EipConnectionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
