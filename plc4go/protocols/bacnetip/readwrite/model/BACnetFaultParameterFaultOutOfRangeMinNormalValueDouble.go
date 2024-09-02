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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
type BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleExactly interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
	isBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble struct {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(doubleValue BACnetApplicationTagDouble, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble {
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble{
		BACnetFaultParameterFaultOutOfRangeMinNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		DoubleValue: doubleValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMinNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMinNormalValueDouble BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) isBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
