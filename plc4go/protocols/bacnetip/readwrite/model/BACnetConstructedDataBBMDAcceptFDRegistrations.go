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

// BACnetConstructedDataBBMDAcceptFDRegistrations is the corresponding interface of BACnetConstructedDataBBMDAcceptFDRegistrations
type BACnetConstructedDataBBMDAcceptFDRegistrations interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBbmdAcceptFDRegistrations returns BbmdAcceptFDRegistrations (property field)
	GetBbmdAcceptFDRegistrations() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataBBMDAcceptFDRegistrationsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBBMDAcceptFDRegistrations.
// This is useful for switch cases.
type BACnetConstructedDataBBMDAcceptFDRegistrationsExactly interface {
	BACnetConstructedDataBBMDAcceptFDRegistrations
	isBACnetConstructedDataBBMDAcceptFDRegistrations() bool
}

// _BACnetConstructedDataBBMDAcceptFDRegistrations is the data-structure of this message
type _BACnetConstructedDataBBMDAcceptFDRegistrations struct {
	BACnetConstructedDataContract
	BbmdAcceptFDRegistrations BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataBBMDAcceptFDRegistrations = (*_BACnetConstructedDataBBMDAcceptFDRegistrations)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_ACCEPT_FD_REGISTRATIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetBbmdAcceptFDRegistrations() BACnetApplicationTagBoolean {
	return m.BbmdAcceptFDRegistrations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetBbmdAcceptFDRegistrations())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBBMDAcceptFDRegistrations factory function for _BACnetConstructedDataBBMDAcceptFDRegistrations
func NewBACnetConstructedDataBBMDAcceptFDRegistrations(bbmdAcceptFDRegistrations BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBBMDAcceptFDRegistrations {
	_result := &_BACnetConstructedDataBBMDAcceptFDRegistrations{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BbmdAcceptFDRegistrations:     bbmdAcceptFDRegistrations,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBBMDAcceptFDRegistrations(structType any) BACnetConstructedDataBBMDAcceptFDRegistrations {
	if casted, ok := structType.(BACnetConstructedDataBBMDAcceptFDRegistrations); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDAcceptFDRegistrations); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetTypeName() string {
	return "BACnetConstructedDataBBMDAcceptFDRegistrations"
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (bbmdAcceptFDRegistrations)
	lengthInBits += m.BbmdAcceptFDRegistrations.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBBMDAcceptFDRegistrations BACnetConstructedDataBBMDAcceptFDRegistrations, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDAcceptFDRegistrations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bbmdAcceptFDRegistrations, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "bbmdAcceptFDRegistrations", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bbmdAcceptFDRegistrations' field"))
	}
	m.BbmdAcceptFDRegistrations = bbmdAcceptFDRegistrations

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), bbmdAcceptFDRegistrations)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDAcceptFDRegistrations")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDAcceptFDRegistrations")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "bbmdAcceptFDRegistrations", m.GetBbmdAcceptFDRegistrations(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bbmdAcceptFDRegistrations' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDAcceptFDRegistrations")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) isBACnetConstructedDataBBMDAcceptFDRegistrations() bool {
	return true
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
