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

// BACnetApplicationTagDouble is the corresponding interface of BACnetApplicationTagDouble
type BACnetApplicationTagDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float64
}

// BACnetApplicationTagDoubleExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagDouble.
// This is useful for switch cases.
type BACnetApplicationTagDoubleExactly interface {
	BACnetApplicationTagDouble
	isBACnetApplicationTagDouble() bool
}

// _BACnetApplicationTagDouble is the data-structure of this message
type _BACnetApplicationTagDouble struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadDouble
}

var _ BACnetApplicationTagDouble = (*_BACnetApplicationTagDouble)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagDouble) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagDouble) GetPayload() BACnetTagPayloadDouble {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagDouble) GetActualValue() float64 {
	ctx := context.Background()
	_ = ctx
	return float64(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagDouble factory function for _BACnetApplicationTagDouble
func NewBACnetApplicationTagDouble(payload BACnetTagPayloadDouble, header BACnetTagHeader) *_BACnetApplicationTagDouble {
	_result := &_BACnetApplicationTagDouble{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagDouble(structType any) BACnetApplicationTagDouble {
	if casted, ok := structType.(BACnetApplicationTagDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagDouble) GetTypeName() string {
	return "BACnetApplicationTagDouble"
}

func (m *_BACnetApplicationTagDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag) (__bACnetApplicationTagDouble BACnetApplicationTagDouble, err error) {
	m.BACnetApplicationTagContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadDouble](ctx, "payload", ReadComplex[BACnetTagPayloadDouble](BACnetTagPayloadDoubleParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[float64](ctx, "actualValue", (*float64)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagDouble")
	}

	return m, nil
}

func (m *_BACnetApplicationTagDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagDouble")
		}

		if err := WriteSimpleField[BACnetTagPayloadDouble](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagDouble")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagDouble) isBACnetApplicationTagDouble() bool {
	return true
}

func (m *_BACnetApplicationTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
