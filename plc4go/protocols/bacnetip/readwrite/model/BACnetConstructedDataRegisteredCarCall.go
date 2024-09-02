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

// BACnetConstructedDataRegisteredCarCall is the corresponding interface of BACnetConstructedDataRegisteredCarCall
type BACnetConstructedDataRegisteredCarCall interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetRegisteredCarCall returns RegisteredCarCall (property field)
	GetRegisteredCarCall() []BACnetLiftCarCallList
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataRegisteredCarCallExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRegisteredCarCall.
// This is useful for switch cases.
type BACnetConstructedDataRegisteredCarCallExactly interface {
	BACnetConstructedDataRegisteredCarCall
	isBACnetConstructedDataRegisteredCarCall() bool
}

// _BACnetConstructedDataRegisteredCarCall is the data-structure of this message
type _BACnetConstructedDataRegisteredCarCall struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	RegisteredCarCall    []BACnetLiftCarCallList
}

var _ BACnetConstructedDataRegisteredCarCall = (*_BACnetConstructedDataRegisteredCarCall)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REGISTERED_CAR_CALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetRegisteredCarCall() []BACnetLiftCarCallList {
	return m.RegisteredCarCall
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRegisteredCarCall) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRegisteredCarCall factory function for _BACnetConstructedDataRegisteredCarCall
func NewBACnetConstructedDataRegisteredCarCall(numberOfDataElements BACnetApplicationTagUnsignedInteger, registeredCarCall []BACnetLiftCarCallList, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRegisteredCarCall {
	_result := &_BACnetConstructedDataRegisteredCarCall{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		RegisteredCarCall:             registeredCarCall,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRegisteredCarCall(structType any) BACnetConstructedDataRegisteredCarCall {
	if casted, ok := structType.(BACnetConstructedDataRegisteredCarCall); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRegisteredCarCall); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetTypeName() string {
	return "BACnetConstructedDataRegisteredCarCall"
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.RegisteredCarCall) > 0 {
		for _, element := range m.RegisteredCarCall {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataRegisteredCarCall) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRegisteredCarCall) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRegisteredCarCall BACnetConstructedDataRegisteredCarCall, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRegisteredCarCall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRegisteredCarCall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	registeredCarCall, err := ReadTerminatedArrayField[BACnetLiftCarCallList](ctx, "registeredCarCall", ReadComplex[BACnetLiftCarCallList](BACnetLiftCarCallListParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'registeredCarCall' field"))
	}
	m.RegisteredCarCall = registeredCarCall

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRegisteredCarCall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRegisteredCarCall")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRegisteredCarCall) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRegisteredCarCall) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRegisteredCarCall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRegisteredCarCall")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "registeredCarCall", m.GetRegisteredCarCall(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'registeredCarCall' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRegisteredCarCall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRegisteredCarCall")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRegisteredCarCall) isBACnetConstructedDataRegisteredCarCall() bool {
	return true
}

func (m *_BACnetConstructedDataRegisteredCarCall) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
