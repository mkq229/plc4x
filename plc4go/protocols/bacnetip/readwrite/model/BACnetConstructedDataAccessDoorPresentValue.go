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

// BACnetConstructedDataAccessDoorPresentValue is the corresponding interface of BACnetConstructedDataAccessDoorPresentValue
type BACnetConstructedDataAccessDoorPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetDoorValueTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorValueTagged
}

// BACnetConstructedDataAccessDoorPresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessDoorPresentValue.
// This is useful for switch cases.
type BACnetConstructedDataAccessDoorPresentValueExactly interface {
	BACnetConstructedDataAccessDoorPresentValue
	isBACnetConstructedDataAccessDoorPresentValue() bool
}

// _BACnetConstructedDataAccessDoorPresentValue is the data-structure of this message
type _BACnetConstructedDataAccessDoorPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetDoorValueTagged
}

var _ BACnetConstructedDataAccessDoorPresentValue = (*_BACnetConstructedDataAccessDoorPresentValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetPresentValue() BACnetDoorValueTagged {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetActualValue() BACnetDoorValueTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorValueTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessDoorPresentValue factory function for _BACnetConstructedDataAccessDoorPresentValue
func NewBACnetConstructedDataAccessDoorPresentValue(presentValue BACnetDoorValueTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorPresentValue {
	_result := &_BACnetConstructedDataAccessDoorPresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorPresentValue(structType any) BACnetConstructedDataAccessDoorPresentValue {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorPresentValue"
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessDoorPresentValue BACnetConstructedDataAccessDoorPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetDoorValueTagged](ctx, "presentValue", ReadComplex[BACnetDoorValueTagged](BACnetDoorValueTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetDoorValueTagged](ctx, "actualValue", (*BACnetDoorValueTagged)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorPresentValue")
		}

		if err := WriteSimpleField[BACnetDoorValueTagged](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetDoorValueTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) isBACnetConstructedDataAccessDoorPresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataAccessDoorPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
