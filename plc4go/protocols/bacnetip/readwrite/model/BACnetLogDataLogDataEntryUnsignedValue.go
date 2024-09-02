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

// BACnetLogDataLogDataEntryUnsignedValue is the corresponding interface of BACnetLogDataLogDataEntryUnsignedValue
type BACnetLogDataLogDataEntryUnsignedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetContextTagUnsignedInteger
}

// BACnetLogDataLogDataEntryUnsignedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryUnsignedValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryUnsignedValueExactly interface {
	BACnetLogDataLogDataEntryUnsignedValue
	isBACnetLogDataLogDataEntryUnsignedValue() bool
}

// _BACnetLogDataLogDataEntryUnsignedValue is the data-structure of this message
type _BACnetLogDataLogDataEntryUnsignedValue struct {
	BACnetLogDataLogDataEntryContract
	UnsignedValue BACnetContextTagUnsignedInteger
}

var _ BACnetLogDataLogDataEntryUnsignedValue = (*_BACnetLogDataLogDataEntryUnsignedValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetUnsignedValue() BACnetContextTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryUnsignedValue factory function for _BACnetLogDataLogDataEntryUnsignedValue
func NewBACnetLogDataLogDataEntryUnsignedValue(unsignedValue BACnetContextTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryUnsignedValue {
	_result := &_BACnetLogDataLogDataEntryUnsignedValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		UnsignedValue:                     unsignedValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryUnsignedValue(structType any) BACnetLogDataLogDataEntryUnsignedValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryUnsignedValue"
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryUnsignedValue BACnetLogDataLogDataEntryUnsignedValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryUnsignedValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryUnsignedValue")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryUnsignedValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) isBACnetLogDataLogDataEntryUnsignedValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
