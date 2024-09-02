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

// BACnetLogDataLogDataTimeChange is the corresponding interface of BACnetLogDataLogDataTimeChange
type BACnetLogDataLogDataTimeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogData
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
}

// BACnetLogDataLogDataTimeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataTimeChange.
// This is useful for switch cases.
type BACnetLogDataLogDataTimeChangeExactly interface {
	BACnetLogDataLogDataTimeChange
	isBACnetLogDataLogDataTimeChange() bool
}

// _BACnetLogDataLogDataTimeChange is the data-structure of this message
type _BACnetLogDataLogDataTimeChange struct {
	BACnetLogDataContract
	TimeChange BACnetContextTagReal
}

var _ BACnetLogDataLogDataTimeChange = (*_BACnetLogDataLogDataTimeChange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataTimeChange) GetParent() BACnetLogDataContract {
	return m.BACnetLogDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataTimeChange factory function for _BACnetLogDataLogDataTimeChange
func NewBACnetLogDataLogDataTimeChange(timeChange BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogDataLogDataTimeChange {
	_result := &_BACnetLogDataLogDataTimeChange{
		BACnetLogDataContract: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
		TimeChange:            timeChange,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataTimeChange(structType any) BACnetLogDataLogDataTimeChange {
	if casted, ok := structType.(BACnetLogDataLogDataTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataTimeChange) GetTypeName() string {
	return "BACnetLogDataLogDataTimeChange"
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataContract.(*_BACnetLogData).getLengthInBits(ctx))

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataTimeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataTimeChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogData, tagNumber uint8) (__bACnetLogDataLogDataTimeChange BACnetLogDataLogDataTimeChange, err error) {
	m.BACnetLogDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeChange, err := ReadSimpleField[BACnetContextTagReal](ctx, "timeChange", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeChange' field"))
	}
	m.TimeChange = timeChange

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataTimeChange")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataTimeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataTimeChange")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "timeChange", m.GetTimeChange(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataTimeChange")
		}
		return nil
	}
	return m.BACnetLogDataContract.(*_BACnetLogData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataTimeChange) isBACnetLogDataLogDataTimeChange() bool {
	return true
}

func (m *_BACnetLogDataLogDataTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
