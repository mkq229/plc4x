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

// BACnetConstructedDataAnalogValueMaxPresValue is the corresponding interface of BACnetConstructedDataAnalogValueMaxPresValue
type BACnetConstructedDataAnalogValueMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataAnalogValueMaxPresValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAnalogValueMaxPresValue.
// This is useful for switch cases.
type BACnetConstructedDataAnalogValueMaxPresValueExactly interface {
	BACnetConstructedDataAnalogValueMaxPresValue
	isBACnetConstructedDataAnalogValueMaxPresValue() bool
}

// _BACnetConstructedDataAnalogValueMaxPresValue is the data-structure of this message
type _BACnetConstructedDataAnalogValueMaxPresValue struct {
	BACnetConstructedDataContract
	MaxPresValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataAnalogValueMaxPresValue = (*_BACnetConstructedDataAnalogValueMaxPresValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetMaxPresValue() BACnetApplicationTagReal {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAnalogValueMaxPresValue factory function for _BACnetConstructedDataAnalogValueMaxPresValue
func NewBACnetConstructedDataAnalogValueMaxPresValue(maxPresValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogValueMaxPresValue {
	_result := &_BACnetConstructedDataAnalogValueMaxPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxPresValue:                  maxPresValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogValueMaxPresValue(structType any) BACnetConstructedDataAnalogValueMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogValueMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValueMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogValueMaxPresValue"
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogValueMaxPresValue BACnetConstructedDataAnalogValueMaxPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValueMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValueMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxPresValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "maxPresValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPresValue' field"))
	}
	m.MaxPresValue = maxPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), maxPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValueMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValueMaxPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValueMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValueMaxPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "maxPresValue", m.GetMaxPresValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValueMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValueMaxPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) isBACnetConstructedDataAnalogValueMaxPresValue() bool {
	return true
}

func (m *_BACnetConstructedDataAnalogValueMaxPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
