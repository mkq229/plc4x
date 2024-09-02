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

// BACnetConstructedDataMultiStateOutputFeedbackValue is the corresponding interface of BACnetConstructedDataMultiStateOutputFeedbackValue
type BACnetConstructedDataMultiStateOutputFeedbackValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataMultiStateOutputFeedbackValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMultiStateOutputFeedbackValue.
// This is useful for switch cases.
type BACnetConstructedDataMultiStateOutputFeedbackValueExactly interface {
	BACnetConstructedDataMultiStateOutputFeedbackValue
	isBACnetConstructedDataMultiStateOutputFeedbackValue() bool
}

// _BACnetConstructedDataMultiStateOutputFeedbackValue is the data-structure of this message
type _BACnetConstructedDataMultiStateOutputFeedbackValue struct {
	BACnetConstructedDataContract
	FeedbackValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMultiStateOutputFeedbackValue = (*_BACnetConstructedDataMultiStateOutputFeedbackValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_OUTPUT
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FEEDBACK_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetFeedbackValue() BACnetApplicationTagUnsignedInteger {
	return m.FeedbackValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFeedbackValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMultiStateOutputFeedbackValue factory function for _BACnetConstructedDataMultiStateOutputFeedbackValue
func NewBACnetConstructedDataMultiStateOutputFeedbackValue(feedbackValue BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateOutputFeedbackValue {
	_result := &_BACnetConstructedDataMultiStateOutputFeedbackValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FeedbackValue:                 feedbackValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateOutputFeedbackValue(structType any) BACnetConstructedDataMultiStateOutputFeedbackValue {
	if casted, ok := structType.(BACnetConstructedDataMultiStateOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateOutputFeedbackValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataMultiStateOutputFeedbackValue"
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateOutputFeedbackValue BACnetConstructedDataMultiStateOutputFeedbackValue, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	feedbackValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "feedbackValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'feedbackValue' field"))
	}
	m.FeedbackValue = feedbackValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), feedbackValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateOutputFeedbackValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateOutputFeedbackValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "feedbackValue", m.GetFeedbackValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'feedbackValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateOutputFeedbackValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) isBACnetConstructedDataMultiStateOutputFeedbackValue() bool {
	return true
}

func (m *_BACnetConstructedDataMultiStateOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
