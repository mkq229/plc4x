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

// BACnetConstructedDataMultiStateInputInterfaceValue is the corresponding interface of BACnetConstructedDataMultiStateInputInterfaceValue
type BACnetConstructedDataMultiStateInputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalBinaryPV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalBinaryPV
}

// BACnetConstructedDataMultiStateInputInterfaceValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMultiStateInputInterfaceValue.
// This is useful for switch cases.
type BACnetConstructedDataMultiStateInputInterfaceValueExactly interface {
	BACnetConstructedDataMultiStateInputInterfaceValue
	isBACnetConstructedDataMultiStateInputInterfaceValue() bool
}

// _BACnetConstructedDataMultiStateInputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataMultiStateInputInterfaceValue struct {
	BACnetConstructedDataContract
	InterfaceValue BACnetOptionalBinaryPV
}

var _ BACnetConstructedDataMultiStateInputInterfaceValue = (*_BACnetConstructedDataMultiStateInputInterfaceValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_INPUT
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetInterfaceValue() BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetActualValue() BACnetOptionalBinaryPV {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalBinaryPV(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMultiStateInputInterfaceValue factory function for _BACnetConstructedDataMultiStateInputInterfaceValue
func NewBACnetConstructedDataMultiStateInputInterfaceValue(interfaceValue BACnetOptionalBinaryPV, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateInputInterfaceValue {
	_result := &_BACnetConstructedDataMultiStateInputInterfaceValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InterfaceValue:                interfaceValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateInputInterfaceValue(structType any) BACnetConstructedDataMultiStateInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataMultiStateInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateInputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataMultiStateInputInterfaceValue"
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateInputInterfaceValue BACnetConstructedDataMultiStateInputInterfaceValue, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateInputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateInputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceValue, err := ReadSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", ReadComplex[BACnetOptionalBinaryPV](BACnetOptionalBinaryPVParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceValue' field"))
	}
	m.InterfaceValue = interfaceValue

	actualValue, err := ReadVirtualField[BACnetOptionalBinaryPV](ctx, "actualValue", (*BACnetOptionalBinaryPV)(nil), interfaceValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateInputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateInputInterfaceValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateInputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateInputInterfaceValue")
		}

		if err := WriteSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", m.GetInterfaceValue(), WriteComplex[BACnetOptionalBinaryPV](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateInputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateInputInterfaceValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) isBACnetConstructedDataMultiStateInputInterfaceValue() bool {
	return true
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
