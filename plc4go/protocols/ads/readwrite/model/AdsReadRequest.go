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

// AdsReadRequest is the corresponding interface of AdsReadRequest
type AdsReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetLength returns Length (property field)
	GetLength() uint32
}

// AdsReadRequestExactly can be used when we want exactly this type and not a type which fulfills AdsReadRequest.
// This is useful for switch cases.
type AdsReadRequestExactly interface {
	AdsReadRequest
	isAdsReadRequest() bool
}

// _AdsReadRequest is the data-structure of this message
type _AdsReadRequest struct {
	AmsPacketContract
	IndexGroup  uint32
	IndexOffset uint32
	Length      uint32
}

var _ AdsReadRequest = (*_AdsReadRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ
}

func (m *_AdsReadRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsReadRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsReadRequest) GetLength() uint32 {
	return m.Length
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadRequest factory function for _AdsReadRequest
func NewAdsReadRequest(indexGroup uint32, indexOffset uint32, length uint32, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsReadRequest {
	_result := &_AdsReadRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		IndexGroup:        indexGroup,
		IndexOffset:       indexOffset,
		Length:            length,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadRequest(structType any) AdsReadRequest {
	if casted, ok := structType.(AdsReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadRequest) GetTypeName() string {
	return "AdsReadRequest"
}

func (m *_AdsReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadRequest AdsReadRequest, err error) {
	m.AmsPacketContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	indexGroup, err := ReadSimpleField(ctx, "indexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexGroup' field"))
	}
	m.IndexGroup = indexGroup

	indexOffset, err := ReadSimpleField(ctx, "indexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexOffset' field"))
	}
	m.IndexOffset = indexOffset

	length, err := ReadSimpleField(ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	m.Length = length

	if closeErr := readBuffer.CloseContext("AdsReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadRequest")
	}

	return m, nil
}

func (m *_AdsReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "indexGroup", m.GetIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "indexOffset", m.GetIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "length", m.GetLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadRequest) isAdsReadRequest() bool {
	return true
}

func (m *_AdsReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
