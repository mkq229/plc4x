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

// BACnetOptionalCharacterStringValue is the corresponding interface of BACnetOptionalCharacterStringValue
type BACnetOptionalCharacterStringValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalCharacterString
	// GetCharacterstring returns Characterstring (property field)
	GetCharacterstring() BACnetApplicationTagCharacterString
}

// BACnetOptionalCharacterStringValueExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalCharacterStringValue.
// This is useful for switch cases.
type BACnetOptionalCharacterStringValueExactly interface {
	BACnetOptionalCharacterStringValue
	isBACnetOptionalCharacterStringValue() bool
}

// _BACnetOptionalCharacterStringValue is the data-structure of this message
type _BACnetOptionalCharacterStringValue struct {
	BACnetOptionalCharacterStringContract
	Characterstring BACnetApplicationTagCharacterString
}

var _ BACnetOptionalCharacterStringValue = (*_BACnetOptionalCharacterStringValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalCharacterStringValue) GetParent() BACnetOptionalCharacterStringContract {
	return m.BACnetOptionalCharacterStringContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalCharacterStringValue) GetCharacterstring() BACnetApplicationTagCharacterString {
	return m.Characterstring
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalCharacterStringValue factory function for _BACnetOptionalCharacterStringValue
func NewBACnetOptionalCharacterStringValue(characterstring BACnetApplicationTagCharacterString, peekedTagHeader BACnetTagHeader) *_BACnetOptionalCharacterStringValue {
	_result := &_BACnetOptionalCharacterStringValue{
		BACnetOptionalCharacterStringContract: NewBACnetOptionalCharacterString(peekedTagHeader),
		Characterstring:                       characterstring,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalCharacterStringValue(structType any) BACnetOptionalCharacterStringValue {
	if casted, ok := structType.(BACnetOptionalCharacterStringValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalCharacterStringValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalCharacterStringValue) GetTypeName() string {
	return "BACnetOptionalCharacterStringValue"
}

func (m *_BACnetOptionalCharacterStringValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString).getLengthInBits(ctx))

	// Simple field (characterstring)
	lengthInBits += m.Characterstring.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalCharacterStringValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalCharacterStringValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalCharacterString) (__bACnetOptionalCharacterStringValue BACnetOptionalCharacterStringValue, err error) {
	m.BACnetOptionalCharacterStringContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalCharacterStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalCharacterStringValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterstring, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "characterstring", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterstring' field"))
	}
	m.Characterstring = characterstring

	if closeErr := readBuffer.CloseContext("BACnetOptionalCharacterStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalCharacterStringValue")
	}

	return m, nil
}

func (m *_BACnetOptionalCharacterStringValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalCharacterStringValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalCharacterStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalCharacterStringValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "characterstring", m.GetCharacterstring(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterstring' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalCharacterStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalCharacterStringValue")
		}
		return nil
	}
	return m.BACnetOptionalCharacterStringContract.(*_BACnetOptionalCharacterString).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalCharacterStringValue) isBACnetOptionalCharacterStringValue() bool {
	return true
}

func (m *_BACnetOptionalCharacterStringValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
