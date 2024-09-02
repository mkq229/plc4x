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

// BACnetConstructedDataDaylightSavingsStatus is the corresponding interface of BACnetConstructedDataDaylightSavingsStatus
type BACnetConstructedDataDaylightSavingsStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDaylightSavingsStatus returns DaylightSavingsStatus (property field)
	GetDaylightSavingsStatus() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataDaylightSavingsStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDaylightSavingsStatus.
// This is useful for switch cases.
type BACnetConstructedDataDaylightSavingsStatusExactly interface {
	BACnetConstructedDataDaylightSavingsStatus
	isBACnetConstructedDataDaylightSavingsStatus() bool
}

// _BACnetConstructedDataDaylightSavingsStatus is the data-structure of this message
type _BACnetConstructedDataDaylightSavingsStatus struct {
	BACnetConstructedDataContract
	DaylightSavingsStatus BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataDaylightSavingsStatus = (*_BACnetConstructedDataDaylightSavingsStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DAYLIGHT_SAVINGS_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetDaylightSavingsStatus() BACnetApplicationTagBoolean {
	return m.DaylightSavingsStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetDaylightSavingsStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDaylightSavingsStatus factory function for _BACnetConstructedDataDaylightSavingsStatus
func NewBACnetConstructedDataDaylightSavingsStatus(daylightSavingsStatus BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDaylightSavingsStatus {
	_result := &_BACnetConstructedDataDaylightSavingsStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DaylightSavingsStatus:         daylightSavingsStatus,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDaylightSavingsStatus(structType any) BACnetConstructedDataDaylightSavingsStatus {
	if casted, ok := structType.(BACnetConstructedDataDaylightSavingsStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDaylightSavingsStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetTypeName() string {
	return "BACnetConstructedDataDaylightSavingsStatus"
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (daylightSavingsStatus)
	lengthInBits += m.DaylightSavingsStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDaylightSavingsStatus BACnetConstructedDataDaylightSavingsStatus, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDaylightSavingsStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDaylightSavingsStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	daylightSavingsStatus, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "daylightSavingsStatus", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daylightSavingsStatus' field"))
	}
	m.DaylightSavingsStatus = daylightSavingsStatus

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), daylightSavingsStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDaylightSavingsStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDaylightSavingsStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDaylightSavingsStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDaylightSavingsStatus")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "daylightSavingsStatus", m.GetDaylightSavingsStatus(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'daylightSavingsStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDaylightSavingsStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDaylightSavingsStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) isBACnetConstructedDataDaylightSavingsStatus() bool {
	return true
}

func (m *_BACnetConstructedDataDaylightSavingsStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
