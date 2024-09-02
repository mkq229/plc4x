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

// BACnetConstructedDataMemberStatusFlags is the corresponding interface of BACnetConstructedDataMemberStatusFlags
type BACnetConstructedDataMemberStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetStatusFlagsTagged
}

// BACnetConstructedDataMemberStatusFlagsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMemberStatusFlags.
// This is useful for switch cases.
type BACnetConstructedDataMemberStatusFlagsExactly interface {
	BACnetConstructedDataMemberStatusFlags
	isBACnetConstructedDataMemberStatusFlags() bool
}

// _BACnetConstructedDataMemberStatusFlags is the data-structure of this message
type _BACnetConstructedDataMemberStatusFlags struct {
	BACnetConstructedDataContract
	StatusFlags BACnetStatusFlagsTagged
}

var _ BACnetConstructedDataMemberStatusFlags = (*_BACnetConstructedDataMemberStatusFlags)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMemberStatusFlags) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMemberStatusFlags) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MEMBER_STATUS_FLAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMemberStatusFlags) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMemberStatusFlags) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMemberStatusFlags) GetActualValue() BACnetStatusFlagsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetStatusFlagsTagged(m.GetStatusFlags())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMemberStatusFlags factory function for _BACnetConstructedDataMemberStatusFlags
func NewBACnetConstructedDataMemberStatusFlags(statusFlags BACnetStatusFlagsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMemberStatusFlags {
	_result := &_BACnetConstructedDataMemberStatusFlags{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StatusFlags:                   statusFlags,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMemberStatusFlags(structType any) BACnetConstructedDataMemberStatusFlags {
	if casted, ok := structType.(BACnetConstructedDataMemberStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMemberStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMemberStatusFlags) GetTypeName() string {
	return "BACnetConstructedDataMemberStatusFlags"
}

func (m *_BACnetConstructedDataMemberStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMemberStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMemberStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMemberStatusFlags BACnetConstructedDataMemberStatusFlags, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMemberStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMemberStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	actualValue, err := ReadVirtualField[BACnetStatusFlagsTagged](ctx, "actualValue", (*BACnetStatusFlagsTagged)(nil), statusFlags)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMemberStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMemberStatusFlags")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMemberStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMemberStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMemberStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMemberStatusFlags")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMemberStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMemberStatusFlags")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMemberStatusFlags) isBACnetConstructedDataMemberStatusFlags() bool {
	return true
}

func (m *_BACnetConstructedDataMemberStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
