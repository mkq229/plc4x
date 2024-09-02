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

// CIPEncapsulationReadResponse is the corresponding interface of CIPEncapsulationReadResponse
type CIPEncapsulationReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CIPEncapsulationPacket
	// GetResponse returns Response (property field)
	GetResponse() DF1ResponseMessage
}

// CIPEncapsulationReadResponseExactly can be used when we want exactly this type and not a type which fulfills CIPEncapsulationReadResponse.
// This is useful for switch cases.
type CIPEncapsulationReadResponseExactly interface {
	CIPEncapsulationReadResponse
	isCIPEncapsulationReadResponse() bool
}

// _CIPEncapsulationReadResponse is the data-structure of this message
type _CIPEncapsulationReadResponse struct {
	CIPEncapsulationPacketContract
	Response DF1ResponseMessage

	// Arguments.
	PacketLen uint16
}

var _ CIPEncapsulationReadResponse = (*_CIPEncapsulationReadResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationReadResponse) GetCommandType() uint16 {
	return 0x0207
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationReadResponse) GetParent() CIPEncapsulationPacketContract {
	return m.CIPEncapsulationPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationReadResponse) GetResponse() DF1ResponseMessage {
	return m.Response
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPEncapsulationReadResponse factory function for _CIPEncapsulationReadResponse
func NewCIPEncapsulationReadResponse(response DF1ResponseMessage, sessionHandle uint32, status uint32, senderContext []uint8, options uint32, packetLen uint16) *_CIPEncapsulationReadResponse {
	_result := &_CIPEncapsulationReadResponse{
		CIPEncapsulationPacketContract: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
		Response:                       response,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationReadResponse(structType any) CIPEncapsulationReadResponse {
	if casted, ok := structType.(CIPEncapsulationReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationReadResponse) GetTypeName() string {
	return "CIPEncapsulationReadResponse"
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).getLengthInBits(ctx))

	// Simple field (response)
	lengthInBits += m.Response.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CIPEncapsulationReadResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CIPEncapsulationPacket, packetLen uint16) (__cIPEncapsulationReadResponse CIPEncapsulationReadResponse, err error) {
	m.CIPEncapsulationPacketContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	response, err := ReadSimpleField[DF1ResponseMessage](ctx, "response", ReadComplex[DF1ResponseMessage](DF1ResponseMessageParseWithBufferProducer[DF1ResponseMessage]((uint16)(packetLen)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'response' field"))
	}
	m.Response = response

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationReadResponse")
	}

	return m, nil
}

func (m *_CIPEncapsulationReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPEncapsulationReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationReadResponse")
		}

		if err := WriteSimpleField[DF1ResponseMessage](ctx, "response", m.GetResponse(), WriteComplex[DF1ResponseMessage](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'response' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationReadResponse")
		}
		return nil
	}
	return m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_CIPEncapsulationReadResponse) GetPacketLen() uint16 {
	return m.PacketLen
}

//
////

func (m *_CIPEncapsulationReadResponse) isCIPEncapsulationReadResponse() bool {
	return true
}

func (m *_CIPEncapsulationReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
