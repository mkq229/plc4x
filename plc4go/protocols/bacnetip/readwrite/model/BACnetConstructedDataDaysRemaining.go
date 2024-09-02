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

// BACnetConstructedDataDaysRemaining is the corresponding interface of BACnetConstructedDataDaysRemaining
type BACnetConstructedDataDaysRemaining interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDaysRemaining returns DaysRemaining (property field)
	GetDaysRemaining() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataDaysRemainingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDaysRemaining.
// This is useful for switch cases.
type BACnetConstructedDataDaysRemainingExactly interface {
	BACnetConstructedDataDaysRemaining
	isBACnetConstructedDataDaysRemaining() bool
}

// _BACnetConstructedDataDaysRemaining is the data-structure of this message
type _BACnetConstructedDataDaysRemaining struct {
	BACnetConstructedDataContract
	DaysRemaining BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataDaysRemaining = (*_BACnetConstructedDataDaysRemaining)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDaysRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYS_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetDaysRemaining() BACnetApplicationTagSignedInteger {
	return m.DaysRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDaysRemaining) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetDaysRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDaysRemaining factory function for _BACnetConstructedDataDaysRemaining
func NewBACnetConstructedDataDaysRemaining(daysRemaining BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDaysRemaining {
	_result := &_BACnetConstructedDataDaysRemaining{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DaysRemaining:                 daysRemaining,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDaysRemaining(structType any) BACnetConstructedDataDaysRemaining {
	if casted, ok := structType.(BACnetConstructedDataDaysRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaysRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDaysRemaining) GetTypeName() string {
	return "BACnetConstructedDataDaysRemaining"
}

func (m *_BACnetConstructedDataDaysRemaining) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (daysRemaining)
	lengthInBits += m.DaysRemaining.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDaysRemaining) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDaysRemaining) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDaysRemaining BACnetConstructedDataDaysRemaining, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaysRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDaysRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	daysRemaining, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "daysRemaining", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daysRemaining' field"))
	}
	m.DaysRemaining = daysRemaining

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), daysRemaining)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaysRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDaysRemaining")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDaysRemaining) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDaysRemaining) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaysRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDaysRemaining")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "daysRemaining", m.GetDaysRemaining(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'daysRemaining' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaysRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDaysRemaining")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDaysRemaining) isBACnetConstructedDataDaysRemaining() bool {
	return true
}

func (m *_BACnetConstructedDataDaysRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
