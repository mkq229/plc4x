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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTrendLogMultipleAll is the corresponding interface of BACnetConstructedDataTrendLogMultipleAll
type BACnetConstructedDataTrendLogMultipleAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataTrendLogMultipleAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTrendLogMultipleAll.
// This is useful for switch cases.
type BACnetConstructedDataTrendLogMultipleAllExactly interface {
	BACnetConstructedDataTrendLogMultipleAll
	isBACnetConstructedDataTrendLogMultipleAll() bool
}

// _BACnetConstructedDataTrendLogMultipleAll is the data-structure of this message
type _BACnetConstructedDataTrendLogMultipleAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataTrendLogMultipleAll = (*_BACnetConstructedDataTrendLogMultipleAll)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG_MULTIPLE
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogMultipleAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// NewBACnetConstructedDataTrendLogMultipleAll factory function for _BACnetConstructedDataTrendLogMultipleAll
func NewBACnetConstructedDataTrendLogMultipleAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogMultipleAll {
	_result := &_BACnetConstructedDataTrendLogMultipleAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogMultipleAll(structType any) BACnetConstructedDataTrendLogMultipleAll {
	if casted, ok := structType.(BACnetConstructedDataTrendLogMultipleAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogMultipleAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) GetTypeName() string {
	return "BACnetConstructedDataTrendLogMultipleAll"
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTrendLogMultipleAll BACnetConstructedDataTrendLogMultipleAll, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogMultipleAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogMultipleAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogMultipleAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogMultipleAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogMultipleAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogMultipleAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogMultipleAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogMultipleAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) isBACnetConstructedDataTrendLogMultipleAll() bool {
	return true
}

func (m *_BACnetConstructedDataTrendLogMultipleAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
