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

// BACnetConstructedDataAPDUTimeout is the corresponding interface of BACnetConstructedDataAPDUTimeout
type BACnetConstructedDataAPDUTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApduTimeout returns ApduTimeout (property field)
	GetApduTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAPDUTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAPDUTimeout.
// This is useful for switch cases.
type BACnetConstructedDataAPDUTimeoutExactly interface {
	BACnetConstructedDataAPDUTimeout
	isBACnetConstructedDataAPDUTimeout() bool
}

// _BACnetConstructedDataAPDUTimeout is the data-structure of this message
type _BACnetConstructedDataAPDUTimeout struct {
	BACnetConstructedDataContract
	ApduTimeout BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAPDUTimeout = (*_BACnetConstructedDataAPDUTimeout)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAPDUTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APDU_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetApduTimeout() BACnetApplicationTagUnsignedInteger {
	return m.ApduTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetApduTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAPDUTimeout factory function for _BACnetConstructedDataAPDUTimeout
func NewBACnetConstructedDataAPDUTimeout(apduTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAPDUTimeout {
	_result := &_BACnetConstructedDataAPDUTimeout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ApduTimeout:                   apduTimeout,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAPDUTimeout(structType any) BACnetConstructedDataAPDUTimeout {
	if casted, ok := structType.(BACnetConstructedDataAPDUTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDUTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAPDUTimeout) GetTypeName() string {
	return "BACnetConstructedDataAPDUTimeout"
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (apduTimeout)
	lengthInBits += m.ApduTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAPDUTimeout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAPDUTimeout BACnetConstructedDataAPDUTimeout, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDUTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDUTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	apduTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "apduTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduTimeout' field"))
	}
	m.ApduTimeout = apduTimeout

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), apduTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDUTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDUTimeout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAPDUTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAPDUTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDUTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDUTimeout")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "apduTimeout", m.GetApduTimeout(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'apduTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDUTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDUTimeout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAPDUTimeout) isBACnetConstructedDataAPDUTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataAPDUTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
