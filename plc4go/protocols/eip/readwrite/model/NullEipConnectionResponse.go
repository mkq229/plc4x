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

// NullEipConnectionResponse is the corresponding interface of NullEipConnectionResponse
type NullEipConnectionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// NullEipConnectionResponseExactly can be used when we want exactly this type and not a type which fulfills NullEipConnectionResponse.
// This is useful for switch cases.
type NullEipConnectionResponseExactly interface {
	NullEipConnectionResponse
	isNullEipConnectionResponse() bool
}

// _NullEipConnectionResponse is the data-structure of this message
type _NullEipConnectionResponse struct {
	*_EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullEipConnectionResponse) GetCommand() uint16 {
	return 0x0065
}

func (m *_NullEipConnectionResponse) GetResponse() bool {
	return bool(true)
}

func (m *_NullEipConnectionResponse) GetPacketLength() uint16 {
	return uint16(0)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullEipConnectionResponse) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_NullEipConnectionResponse) GetParent() EipPacket {
	return m._EipPacket
}

// NewNullEipConnectionResponse factory function for _NullEipConnectionResponse
func NewNullEipConnectionResponse(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_NullEipConnectionResponse {
	_result := &_NullEipConnectionResponse{
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNullEipConnectionResponse(structType any) NullEipConnectionResponse {
	if casted, ok := structType.(NullEipConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NullEipConnectionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NullEipConnectionResponse) GetTypeName() string {
	return "NullEipConnectionResponse"
}

func (m *_NullEipConnectionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullEipConnectionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NullEipConnectionResponseParse(ctx context.Context, theBytes []byte, response bool) (NullEipConnectionResponse, error) {
	return NullEipConnectionResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func NullEipConnectionResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (NullEipConnectionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NullEipConnectionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullEipConnectionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullEipConnectionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullEipConnectionResponse")
	}

	// Create a partially initialized instance
	_child := &_NullEipConnectionResponse{
		_EipPacket: &_EipPacket{},
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_NullEipConnectionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullEipConnectionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullEipConnectionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullEipConnectionResponse")
		}

		if popErr := writeBuffer.PopContext("NullEipConnectionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullEipConnectionResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullEipConnectionResponse) isNullEipConnectionResponse() bool {
	return true
}

func (m *_NullEipConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
