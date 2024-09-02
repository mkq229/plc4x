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

// BACnetContextTagDate is the corresponding interface of BACnetContextTagDate
type BACnetContextTagDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDate
}

// BACnetContextTagDateExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagDate.
// This is useful for switch cases.
type BACnetContextTagDateExactly interface {
	BACnetContextTagDate
	isBACnetContextTagDate() bool
}

// _BACnetContextTagDate is the data-structure of this message
type _BACnetContextTagDate struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadDate
}

var _ BACnetContextTagDate = (*_BACnetContextTagDate)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagDate) GetDataType() BACnetDataType {
	return BACnetDataType_DATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagDate) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagDate) GetPayload() BACnetTagPayloadDate {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagDate factory function for _BACnetContextTagDate
func NewBACnetContextTagDate(payload BACnetTagPayloadDate, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagDate {
	_result := &_BACnetContextTagDate{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagDate(structType any) BACnetContextTagDate {
	if casted, ok := structType.(BACnetContextTagDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagDate) GetTypeName() string {
	return "BACnetContextTagDate"
}

func (m *_BACnetContextTagDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetContextTagDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagDate BACnetContextTagDate, err error) {
	m.BACnetContextTagContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadDate](ctx, "payload", ReadComplex[BACnetTagPayloadDate](BACnetTagPayloadDateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetContextTagDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagDate")
	}

	return m, nil
}

func (m *_BACnetContextTagDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagDate")
		}

		if err := WriteSimpleField[BACnetTagPayloadDate](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagDate")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagDate) isBACnetContextTagDate() bool {
	return true
}

func (m *_BACnetContextTagDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
