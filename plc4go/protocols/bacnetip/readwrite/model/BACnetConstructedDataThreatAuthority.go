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

// BACnetConstructedDataThreatAuthority is the corresponding interface of BACnetConstructedDataThreatAuthority
type BACnetConstructedDataThreatAuthority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetThreatAuthority returns ThreatAuthority (property field)
	GetThreatAuthority() BACnetAccessThreatLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessThreatLevel
}

// BACnetConstructedDataThreatAuthorityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataThreatAuthority.
// This is useful for switch cases.
type BACnetConstructedDataThreatAuthorityExactly interface {
	BACnetConstructedDataThreatAuthority
	isBACnetConstructedDataThreatAuthority() bool
}

// _BACnetConstructedDataThreatAuthority is the data-structure of this message
type _BACnetConstructedDataThreatAuthority struct {
	BACnetConstructedDataContract
	ThreatAuthority BACnetAccessThreatLevel
}

var _ BACnetConstructedDataThreatAuthority = (*_BACnetConstructedDataThreatAuthority)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataThreatAuthority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_THREAT_AUTHORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetThreatAuthority() BACnetAccessThreatLevel {
	return m.ThreatAuthority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataThreatAuthority) GetActualValue() BACnetAccessThreatLevel {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessThreatLevel(m.GetThreatAuthority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataThreatAuthority factory function for _BACnetConstructedDataThreatAuthority
func NewBACnetConstructedDataThreatAuthority(threatAuthority BACnetAccessThreatLevel, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataThreatAuthority {
	_result := &_BACnetConstructedDataThreatAuthority{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ThreatAuthority:               threatAuthority,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataThreatAuthority(structType any) BACnetConstructedDataThreatAuthority {
	if casted, ok := structType.(BACnetConstructedDataThreatAuthority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataThreatAuthority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataThreatAuthority) GetTypeName() string {
	return "BACnetConstructedDataThreatAuthority"
}

func (m *_BACnetConstructedDataThreatAuthority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (threatAuthority)
	lengthInBits += m.ThreatAuthority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataThreatAuthority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataThreatAuthority) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataThreatAuthority BACnetConstructedDataThreatAuthority, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataThreatAuthority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataThreatAuthority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	threatAuthority, err := ReadSimpleField[BACnetAccessThreatLevel](ctx, "threatAuthority", ReadComplex[BACnetAccessThreatLevel](BACnetAccessThreatLevelParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'threatAuthority' field"))
	}
	m.ThreatAuthority = threatAuthority

	actualValue, err := ReadVirtualField[BACnetAccessThreatLevel](ctx, "actualValue", (*BACnetAccessThreatLevel)(nil), threatAuthority)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataThreatAuthority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataThreatAuthority")
	}

	return m, nil
}

func (m *_BACnetConstructedDataThreatAuthority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataThreatAuthority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataThreatAuthority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataThreatAuthority")
		}

		if err := WriteSimpleField[BACnetAccessThreatLevel](ctx, "threatAuthority", m.GetThreatAuthority(), WriteComplex[BACnetAccessThreatLevel](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'threatAuthority' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataThreatAuthority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataThreatAuthority")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataThreatAuthority) isBACnetConstructedDataThreatAuthority() bool {
	return true
}

func (m *_BACnetConstructedDataThreatAuthority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
