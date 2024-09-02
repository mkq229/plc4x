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

// BACnetConstructedDataLastStateChange is the corresponding interface of BACnetConstructedDataLastStateChange
type BACnetConstructedDataLastStateChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastStateChange returns LastStateChange (property field)
	GetLastStateChange() BACnetTimerTransitionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimerTransitionTagged
}

// BACnetConstructedDataLastStateChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastStateChange.
// This is useful for switch cases.
type BACnetConstructedDataLastStateChangeExactly interface {
	BACnetConstructedDataLastStateChange
	isBACnetConstructedDataLastStateChange() bool
}

// _BACnetConstructedDataLastStateChange is the data-structure of this message
type _BACnetConstructedDataLastStateChange struct {
	BACnetConstructedDataContract
	LastStateChange BACnetTimerTransitionTagged
}

var _ BACnetConstructedDataLastStateChange = (*_BACnetConstructedDataLastStateChange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastStateChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_STATE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetLastStateChange() BACnetTimerTransitionTagged {
	return m.LastStateChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetActualValue() BACnetTimerTransitionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimerTransitionTagged(m.GetLastStateChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastStateChange factory function for _BACnetConstructedDataLastStateChange
func NewBACnetConstructedDataLastStateChange(lastStateChange BACnetTimerTransitionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastStateChange {
	_result := &_BACnetConstructedDataLastStateChange{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastStateChange:               lastStateChange,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastStateChange(structType any) BACnetConstructedDataLastStateChange {
	if casted, ok := structType.(BACnetConstructedDataLastStateChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastStateChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastStateChange) GetTypeName() string {
	return "BACnetConstructedDataLastStateChange"
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastStateChange)
	lengthInBits += m.LastStateChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastStateChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastStateChange BACnetConstructedDataLastStateChange, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastStateChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastStateChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastStateChange, err := ReadSimpleField[BACnetTimerTransitionTagged](ctx, "lastStateChange", ReadComplex[BACnetTimerTransitionTagged](BACnetTimerTransitionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastStateChange' field"))
	}
	m.LastStateChange = lastStateChange

	actualValue, err := ReadVirtualField[BACnetTimerTransitionTagged](ctx, "actualValue", (*BACnetTimerTransitionTagged)(nil), lastStateChange)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastStateChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastStateChange")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastStateChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastStateChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastStateChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastStateChange")
		}

		if err := WriteSimpleField[BACnetTimerTransitionTagged](ctx, "lastStateChange", m.GetLastStateChange(), WriteComplex[BACnetTimerTransitionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastStateChange' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastStateChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastStateChange")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastStateChange) isBACnetConstructedDataLastStateChange() bool {
	return true
}

func (m *_BACnetConstructedDataLastStateChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
