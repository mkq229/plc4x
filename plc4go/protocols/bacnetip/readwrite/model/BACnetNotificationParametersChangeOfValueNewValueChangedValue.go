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

// BACnetNotificationParametersChangeOfValueNewValueChangedValue is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValueChangedValue
type BACnetNotificationParametersChangeOfValueNewValueChangedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfValueNewValue
	// GetChangedValue returns ChangedValue (property field)
	GetChangedValue() BACnetContextTagReal
}

// BACnetNotificationParametersChangeOfValueNewValueChangedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfValueNewValueChangedValue.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfValueNewValueChangedValueExactly interface {
	BACnetNotificationParametersChangeOfValueNewValueChangedValue
	isBACnetNotificationParametersChangeOfValueNewValueChangedValue() bool
}

// _BACnetNotificationParametersChangeOfValueNewValueChangedValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfValueNewValueChangedValue struct {
	BACnetNotificationParametersChangeOfValueNewValueContract
	ChangedValue BACnetContextTagReal
}

var _ BACnetNotificationParametersChangeOfValueNewValueChangedValue = (*_BACnetNotificationParametersChangeOfValueNewValueChangedValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetParent() BACnetNotificationParametersChangeOfValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetChangedValue() BACnetContextTagReal {
	return m.ChangedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValueChangedValue factory function for _BACnetNotificationParametersChangeOfValueNewValueChangedValue
func NewBACnetNotificationParametersChangeOfValueNewValueChangedValue(changedValue BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfValueNewValueChangedValue {
	_result := &_BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		BACnetNotificationParametersChangeOfValueNewValueContract: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		ChangedValue: changedValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(structType any) BACnetNotificationParametersChangeOfValueNewValueChangedValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedValue"
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfValueNewValueContract.(*_BACnetNotificationParametersChangeOfValueNewValue).getLengthInBits(ctx))

	// Simple field (changedValue)
	lengthInBits += m.ChangedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfValueNewValue, peekedTagNumber uint8, tagNumber uint8) (__bACnetNotificationParametersChangeOfValueNewValueChangedValue BACnetNotificationParametersChangeOfValueNewValueChangedValue, err error) {
	m.BACnetNotificationParametersChangeOfValueNewValueContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	changedValue, err := ReadSimpleField[BACnetContextTagReal](ctx, "changedValue", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'changedValue' field"))
	}
	m.ChangedValue = changedValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "changedValue", m.GetChangedValue(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'changedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValueNewValueChangedValue")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfValueNewValueContract.(*_BACnetNotificationParametersChangeOfValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) isBACnetNotificationParametersChangeOfValueNewValueChangedValue() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
