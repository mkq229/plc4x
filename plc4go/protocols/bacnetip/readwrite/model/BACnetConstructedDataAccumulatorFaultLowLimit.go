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

// BACnetConstructedDataAccumulatorFaultLowLimit is the corresponding interface of BACnetConstructedDataAccumulatorFaultLowLimit
type BACnetConstructedDataAccumulatorFaultLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAccumulatorFaultLowLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccumulatorFaultLowLimit.
// This is useful for switch cases.
type BACnetConstructedDataAccumulatorFaultLowLimitExactly interface {
	BACnetConstructedDataAccumulatorFaultLowLimit
	isBACnetConstructedDataAccumulatorFaultLowLimit() bool
}

// _BACnetConstructedDataAccumulatorFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataAccumulatorFaultLowLimit struct {
	BACnetConstructedDataContract
	FaultLowLimit BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAccumulatorFaultLowLimit = (*_BACnetConstructedDataAccumulatorFaultLowLimit)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagUnsignedInteger {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorFaultLowLimit factory function for _BACnetConstructedDataAccumulatorFaultLowLimit
func NewBACnetConstructedDataAccumulatorFaultLowLimit(faultLowLimit BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorFaultLowLimit {
	_result := &_BACnetConstructedDataAccumulatorFaultLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultLowLimit:                 faultLowLimit,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorFaultLowLimit(structType any) BACnetConstructedDataAccumulatorFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorFaultLowLimit"
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccumulatorFaultLowLimit BACnetConstructedDataAccumulatorFaultLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultLowLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "faultLowLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultLowLimit' field"))
	}
	m.FaultLowLimit = faultLowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), faultLowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorFaultLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorFaultLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "faultLowLimit", m.GetFaultLowLimit(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorFaultLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) isBACnetConstructedDataAccumulatorFaultLowLimit() bool {
	return true
}

func (m *_BACnetConstructedDataAccumulatorFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
