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

// AdsMultiRequestItemRead is the corresponding interface of AdsMultiRequestItemRead
type AdsMultiRequestItemRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsMultiRequestItem
	// GetItemIndexGroup returns ItemIndexGroup (property field)
	GetItemIndexGroup() uint32
	// GetItemIndexOffset returns ItemIndexOffset (property field)
	GetItemIndexOffset() uint32
	// GetItemReadLength returns ItemReadLength (property field)
	GetItemReadLength() uint32
}

// AdsMultiRequestItemReadExactly can be used when we want exactly this type and not a type which fulfills AdsMultiRequestItemRead.
// This is useful for switch cases.
type AdsMultiRequestItemReadExactly interface {
	AdsMultiRequestItemRead
	isAdsMultiRequestItemRead() bool
}

// _AdsMultiRequestItemRead is the data-structure of this message
type _AdsMultiRequestItemRead struct {
	AdsMultiRequestItemContract
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
}

var _ AdsMultiRequestItemRead = (*_AdsMultiRequestItemRead)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsMultiRequestItemRead) GetIndexGroup() uint32 {
	return uint32(61568)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsMultiRequestItemRead) GetParent() AdsMultiRequestItemContract {
	return m.AdsMultiRequestItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsMultiRequestItemRead) GetItemIndexGroup() uint32 {
	return m.ItemIndexGroup
}

func (m *_AdsMultiRequestItemRead) GetItemIndexOffset() uint32 {
	return m.ItemIndexOffset
}

func (m *_AdsMultiRequestItemRead) GetItemReadLength() uint32 {
	return m.ItemReadLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsMultiRequestItemRead factory function for _AdsMultiRequestItemRead
func NewAdsMultiRequestItemRead(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32) *_AdsMultiRequestItemRead {
	_result := &_AdsMultiRequestItemRead{
		AdsMultiRequestItemContract: NewAdsMultiRequestItem(),
		ItemIndexGroup:              itemIndexGroup,
		ItemIndexOffset:             itemIndexOffset,
		ItemReadLength:              itemReadLength,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItemRead(structType any) AdsMultiRequestItemRead {
	if casted, ok := structType.(AdsMultiRequestItemRead); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItemRead); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItemRead) GetTypeName() string {
	return "AdsMultiRequestItemRead"
}

func (m *_AdsMultiRequestItemRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).getLengthInBits(ctx))

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsMultiRequestItemRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsMultiRequestItemRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsMultiRequestItem, indexGroup uint32) (__adsMultiRequestItemRead AdsMultiRequestItemRead, err error) {
	m.AdsMultiRequestItemContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItemRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemIndexGroup, err := ReadSimpleField(ctx, "itemIndexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexGroup' field"))
	}
	m.ItemIndexGroup = itemIndexGroup

	itemIndexOffset, err := ReadSimpleField(ctx, "itemIndexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndexOffset' field"))
	}
	m.ItemIndexOffset = itemIndexOffset

	itemReadLength, err := ReadSimpleField(ctx, "itemReadLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemReadLength' field"))
	}
	m.ItemReadLength = itemReadLength

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItemRead")
	}

	return m, nil
}

func (m *_AdsMultiRequestItemRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsMultiRequestItemRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItemRead")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexGroup", m.GetItemIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemIndexOffset", m.GetItemIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemIndexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "itemReadLength", m.GetItemReadLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemReadLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsMultiRequestItemRead")
		}
		return nil
	}
	return m.AdsMultiRequestItemContract.(*_AdsMultiRequestItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsMultiRequestItemRead) isAdsMultiRequestItemRead() bool {
	return true
}

func (m *_AdsMultiRequestItemRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
