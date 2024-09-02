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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerExactly interface {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
	isBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger() bool
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger struct {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(integerValue BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger {
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger{
		BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		IntegerValue: integerValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) isBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
