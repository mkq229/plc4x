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

// BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged is the corresponding interface of BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
type BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable
}

// BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedExactly interface {
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	isBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged() bool
}

// _BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged is the data-structure of this message
type _BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged struct {
	Header BACnetTagHeader
	Value  BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged = (*_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetValue() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged factory function for _BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
func NewBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged(header BACnetTagHeader, value BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, tagNumber uint8, tagClass TagClass) *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged {
	return &_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged(structType any) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged"
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, error) {
	return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, error) {
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, error) {
	v, err := (&_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged{}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged")
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

	value, err := ReadManualField[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable](ctx, "value", readBuffer, EnsureType[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) isBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
