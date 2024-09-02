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

// BACnetConstructedDataRestoreCompletionTime is the corresponding interface of BACnetConstructedDataRestoreCompletionTime
type BACnetConstructedDataRestoreCompletionTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCompletionTime returns CompletionTime (property field)
	GetCompletionTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataRestoreCompletionTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRestoreCompletionTime.
// This is useful for switch cases.
type BACnetConstructedDataRestoreCompletionTimeExactly interface {
	BACnetConstructedDataRestoreCompletionTime
	isBACnetConstructedDataRestoreCompletionTime() bool
}

// _BACnetConstructedDataRestoreCompletionTime is the data-structure of this message
type _BACnetConstructedDataRestoreCompletionTime struct {
	BACnetConstructedDataContract
	CompletionTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataRestoreCompletionTime = (*_BACnetConstructedDataRestoreCompletionTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRestoreCompletionTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRestoreCompletionTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESTORE_COMPLETION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRestoreCompletionTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRestoreCompletionTime) GetCompletionTime() BACnetApplicationTagUnsignedInteger {
	return m.CompletionTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRestoreCompletionTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCompletionTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRestoreCompletionTime factory function for _BACnetConstructedDataRestoreCompletionTime
func NewBACnetConstructedDataRestoreCompletionTime(completionTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRestoreCompletionTime {
	_result := &_BACnetConstructedDataRestoreCompletionTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CompletionTime:                completionTime,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRestoreCompletionTime(structType any) BACnetConstructedDataRestoreCompletionTime {
	if casted, ok := structType.(BACnetConstructedDataRestoreCompletionTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestoreCompletionTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRestoreCompletionTime) GetTypeName() string {
	return "BACnetConstructedDataRestoreCompletionTime"
}

func (m *_BACnetConstructedDataRestoreCompletionTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (completionTime)
	lengthInBits += m.CompletionTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRestoreCompletionTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRestoreCompletionTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRestoreCompletionTime BACnetConstructedDataRestoreCompletionTime, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRestoreCompletionTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRestoreCompletionTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	completionTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "completionTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'completionTime' field"))
	}
	m.CompletionTime = completionTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), completionTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestoreCompletionTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRestoreCompletionTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRestoreCompletionTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRestoreCompletionTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRestoreCompletionTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRestoreCompletionTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "completionTime", m.GetCompletionTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'completionTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestoreCompletionTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRestoreCompletionTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRestoreCompletionTime) isBACnetConstructedDataRestoreCompletionTime() bool {
	return true
}

func (m *_BACnetConstructedDataRestoreCompletionTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
