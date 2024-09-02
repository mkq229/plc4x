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

// BACnetDaysOfWeekTagged is the corresponding interface of BACnetDaysOfWeekTagged
type BACnetDaysOfWeekTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetMonday returns Monday (virtual field)
	GetMonday() bool
	// GetTuesday returns Tuesday (virtual field)
	GetTuesday() bool
	// GetWednesday returns Wednesday (virtual field)
	GetWednesday() bool
	// GetThursday returns Thursday (virtual field)
	GetThursday() bool
	// GetFriday returns Friday (virtual field)
	GetFriday() bool
	// GetSaturday returns Saturday (virtual field)
	GetSaturday() bool
	// GetSunday returns Sunday (virtual field)
	GetSunday() bool
}

// BACnetDaysOfWeekTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetDaysOfWeekTagged.
// This is useful for switch cases.
type BACnetDaysOfWeekTaggedExactly interface {
	BACnetDaysOfWeekTagged
	isBACnetDaysOfWeekTagged() bool
}

// _BACnetDaysOfWeekTagged is the data-structure of this message
type _BACnetDaysOfWeekTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetDaysOfWeekTagged = (*_BACnetDaysOfWeekTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDaysOfWeekTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetDaysOfWeekTagged) GetPayload() BACnetTagPayloadBitString {
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

func (m *_BACnetDaysOfWeekTagged) GetMonday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetDaysOfWeekTagged) GetTuesday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetDaysOfWeekTagged) GetWednesday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetDaysOfWeekTagged) GetThursday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (3))), func() any { return bool(m.GetPayload().GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetDaysOfWeekTagged) GetFriday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (4))), func() any { return bool(m.GetPayload().GetData()[4]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetDaysOfWeekTagged) GetSaturday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (5))), func() any { return bool(m.GetPayload().GetData()[5]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetDaysOfWeekTagged) GetSunday() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (6))), func() any { return bool(m.GetPayload().GetData()[6]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDaysOfWeekTagged factory function for _BACnetDaysOfWeekTagged
func NewBACnetDaysOfWeekTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetDaysOfWeekTagged {
	return &_BACnetDaysOfWeekTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetDaysOfWeekTagged(structType any) BACnetDaysOfWeekTagged {
	if casted, ok := structType.(BACnetDaysOfWeekTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDaysOfWeekTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDaysOfWeekTagged) GetTypeName() string {
	return "BACnetDaysOfWeekTagged"
}

func (m *_BACnetDaysOfWeekTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetDaysOfWeekTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDaysOfWeekTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetDaysOfWeekTagged, error) {
	return BACnetDaysOfWeekTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetDaysOfWeekTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDaysOfWeekTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDaysOfWeekTagged, error) {
		return BACnetDaysOfWeekTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetDaysOfWeekTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetDaysOfWeekTagged, error) {
	v, err := (&_BACnetDaysOfWeekTagged{}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetDaysOfWeekTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetDaysOfWeekTagged BACnetDaysOfWeekTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDaysOfWeekTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDaysOfWeekTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	monday, err := ReadVirtualField[bool](ctx, "monday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monday' field"))
	}
	_ = monday

	tuesday, err := ReadVirtualField[bool](ctx, "tuesday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tuesday' field"))
	}
	_ = tuesday

	wednesday, err := ReadVirtualField[bool](ctx, "wednesday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'wednesday' field"))
	}
	_ = wednesday

	thursday, err := ReadVirtualField[bool](ctx, "thursday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (3))), func() any { return bool(payload.GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'thursday' field"))
	}
	_ = thursday

	friday, err := ReadVirtualField[bool](ctx, "friday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (4))), func() any { return bool(payload.GetData()[4]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'friday' field"))
	}
	_ = friday

	saturday, err := ReadVirtualField[bool](ctx, "saturday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (5))), func() any { return bool(payload.GetData()[5]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'saturday' field"))
	}
	_ = saturday

	sunday, err := ReadVirtualField[bool](ctx, "sunday", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (6))), func() any { return bool(payload.GetData()[6]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sunday' field"))
	}
	_ = sunday

	if closeErr := readBuffer.CloseContext("BACnetDaysOfWeekTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDaysOfWeekTagged")
	}

	return m, nil
}

func (m *_BACnetDaysOfWeekTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDaysOfWeekTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDaysOfWeekTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDaysOfWeekTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	monday := m.GetMonday()
	_ = monday
	if _mondayErr := writeBuffer.WriteVirtual(ctx, "monday", m.GetMonday()); _mondayErr != nil {
		return errors.Wrap(_mondayErr, "Error serializing 'monday' field")
	}
	// Virtual field
	tuesday := m.GetTuesday()
	_ = tuesday
	if _tuesdayErr := writeBuffer.WriteVirtual(ctx, "tuesday", m.GetTuesday()); _tuesdayErr != nil {
		return errors.Wrap(_tuesdayErr, "Error serializing 'tuesday' field")
	}
	// Virtual field
	wednesday := m.GetWednesday()
	_ = wednesday
	if _wednesdayErr := writeBuffer.WriteVirtual(ctx, "wednesday", m.GetWednesday()); _wednesdayErr != nil {
		return errors.Wrap(_wednesdayErr, "Error serializing 'wednesday' field")
	}
	// Virtual field
	thursday := m.GetThursday()
	_ = thursday
	if _thursdayErr := writeBuffer.WriteVirtual(ctx, "thursday", m.GetThursday()); _thursdayErr != nil {
		return errors.Wrap(_thursdayErr, "Error serializing 'thursday' field")
	}
	// Virtual field
	friday := m.GetFriday()
	_ = friday
	if _fridayErr := writeBuffer.WriteVirtual(ctx, "friday", m.GetFriday()); _fridayErr != nil {
		return errors.Wrap(_fridayErr, "Error serializing 'friday' field")
	}
	// Virtual field
	saturday := m.GetSaturday()
	_ = saturday
	if _saturdayErr := writeBuffer.WriteVirtual(ctx, "saturday", m.GetSaturday()); _saturdayErr != nil {
		return errors.Wrap(_saturdayErr, "Error serializing 'saturday' field")
	}
	// Virtual field
	sunday := m.GetSunday()
	_ = sunday
	if _sundayErr := writeBuffer.WriteVirtual(ctx, "sunday", m.GetSunday()); _sundayErr != nil {
		return errors.Wrap(_sundayErr, "Error serializing 'sunday' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDaysOfWeekTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDaysOfWeekTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDaysOfWeekTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetDaysOfWeekTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetDaysOfWeekTagged) isBACnetDaysOfWeekTagged() bool {
	return true
}

func (m *_BACnetDaysOfWeekTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
