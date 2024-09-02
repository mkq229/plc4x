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

// BACnetConstructedDataEventMessageTextsConfig is the corresponding interface of BACnetConstructedDataEventMessageTextsConfig
type BACnetConstructedDataEventMessageTextsConfig interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetEventMessageTextsConfig returns EventMessageTextsConfig (property field)
	GetEventMessageTextsConfig() []BACnetOptionalCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormalTextConfig returns ToOffnormalTextConfig (virtual field)
	GetToOffnormalTextConfig() BACnetOptionalCharacterString
	// GetToFaultTextConfig returns ToFaultTextConfig (virtual field)
	GetToFaultTextConfig() BACnetOptionalCharacterString
	// GetToNormalTextConfig returns ToNormalTextConfig (virtual field)
	GetToNormalTextConfig() BACnetOptionalCharacterString
}

// BACnetConstructedDataEventMessageTextsConfigExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventMessageTextsConfig.
// This is useful for switch cases.
type BACnetConstructedDataEventMessageTextsConfigExactly interface {
	BACnetConstructedDataEventMessageTextsConfig
	isBACnetConstructedDataEventMessageTextsConfig() bool
}

// _BACnetConstructedDataEventMessageTextsConfig is the data-structure of this message
type _BACnetConstructedDataEventMessageTextsConfig struct {
	BACnetConstructedDataContract
	NumberOfDataElements    BACnetApplicationTagUnsignedInteger
	EventMessageTextsConfig []BACnetOptionalCharacterString
}

var _ BACnetConstructedDataEventMessageTextsConfig = (*_BACnetConstructedDataEventMessageTextsConfig)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_MESSAGE_TEXTS_CONFIG
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetEventMessageTextsConfig() []BACnetOptionalCharacterString {
	return m.EventMessageTextsConfig
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetToOffnormalTextConfig() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTextsConfig())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTextsConfig()[0]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetToFaultTextConfig() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTextsConfig())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTextsConfig()[1]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetToNormalTextConfig() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTextsConfig())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTextsConfig()[2]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventMessageTextsConfig factory function for _BACnetConstructedDataEventMessageTextsConfig
func NewBACnetConstructedDataEventMessageTextsConfig(numberOfDataElements BACnetApplicationTagUnsignedInteger, eventMessageTextsConfig []BACnetOptionalCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventMessageTextsConfig {
	_result := &_BACnetConstructedDataEventMessageTextsConfig{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		EventMessageTextsConfig:       eventMessageTextsConfig,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventMessageTextsConfig(structType any) BACnetConstructedDataEventMessageTextsConfig {
	if casted, ok := structType.(BACnetConstructedDataEventMessageTextsConfig); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventMessageTextsConfig); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetTypeName() string {
	return "BACnetConstructedDataEventMessageTextsConfig"
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.EventMessageTextsConfig) > 0 {
		for _, element := range m.EventMessageTextsConfig {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventMessageTextsConfig BACnetConstructedDataEventMessageTextsConfig, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventMessageTextsConfig"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventMessageTextsConfig")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	eventMessageTextsConfig, err := ReadTerminatedArrayField[BACnetOptionalCharacterString](ctx, "eventMessageTextsConfig", ReadComplex[BACnetOptionalCharacterString](BACnetOptionalCharacterStringParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventMessageTextsConfig' field"))
	}
	m.EventMessageTextsConfig = eventMessageTextsConfig

	toOffnormalTextConfig, err := ReadVirtualField[BACnetOptionalCharacterString](ctx, "toOffnormalTextConfig", (*BACnetOptionalCharacterString)(nil), CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTextsConfig)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTextsConfig[0]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormalTextConfig' field"))
	}
	_ = toOffnormalTextConfig

	toFaultTextConfig, err := ReadVirtualField[BACnetOptionalCharacterString](ctx, "toFaultTextConfig", (*BACnetOptionalCharacterString)(nil), CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTextsConfig)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTextsConfig[1]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFaultTextConfig' field"))
	}
	_ = toFaultTextConfig

	toNormalTextConfig, err := ReadVirtualField[BACnetOptionalCharacterString](ctx, "toNormalTextConfig", (*BACnetOptionalCharacterString)(nil), CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTextsConfig)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTextsConfig[2]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormalTextConfig' field"))
	}
	_ = toNormalTextConfig

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventMessageTextsConfig)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "eventMessageTextsConfig should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventMessageTextsConfig"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventMessageTextsConfig")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventMessageTextsConfig"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventMessageTextsConfig")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "eventMessageTextsConfig", m.GetEventMessageTextsConfig(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'eventMessageTextsConfig' field")
		}
		// Virtual field
		toOffnormalTextConfig := m.GetToOffnormalTextConfig()
		_ = toOffnormalTextConfig
		if _toOffnormalTextConfigErr := writeBuffer.WriteVirtual(ctx, "toOffnormalTextConfig", m.GetToOffnormalTextConfig()); _toOffnormalTextConfigErr != nil {
			return errors.Wrap(_toOffnormalTextConfigErr, "Error serializing 'toOffnormalTextConfig' field")
		}
		// Virtual field
		toFaultTextConfig := m.GetToFaultTextConfig()
		_ = toFaultTextConfig
		if _toFaultTextConfigErr := writeBuffer.WriteVirtual(ctx, "toFaultTextConfig", m.GetToFaultTextConfig()); _toFaultTextConfigErr != nil {
			return errors.Wrap(_toFaultTextConfigErr, "Error serializing 'toFaultTextConfig' field")
		}
		// Virtual field
		toNormalTextConfig := m.GetToNormalTextConfig()
		_ = toNormalTextConfig
		if _toNormalTextConfigErr := writeBuffer.WriteVirtual(ctx, "toNormalTextConfig", m.GetToNormalTextConfig()); _toNormalTextConfigErr != nil {
			return errors.Wrap(_toNormalTextConfigErr, "Error serializing 'toNormalTextConfig' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventMessageTextsConfig"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventMessageTextsConfig")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) isBACnetConstructedDataEventMessageTextsConfig() bool {
	return true
}

func (m *_BACnetConstructedDataEventMessageTextsConfig) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
