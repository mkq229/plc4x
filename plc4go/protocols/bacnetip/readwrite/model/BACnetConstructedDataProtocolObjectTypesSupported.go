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

// BACnetConstructedDataProtocolObjectTypesSupported is the corresponding interface of BACnetConstructedDataProtocolObjectTypesSupported
type BACnetConstructedDataProtocolObjectTypesSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProtocolObjectTypesSupported returns ProtocolObjectTypesSupported (property field)
	GetProtocolObjectTypesSupported() BACnetObjectTypesSupportedTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectTypesSupportedTagged
}

// BACnetConstructedDataProtocolObjectTypesSupportedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProtocolObjectTypesSupported.
// This is useful for switch cases.
type BACnetConstructedDataProtocolObjectTypesSupportedExactly interface {
	BACnetConstructedDataProtocolObjectTypesSupported
	isBACnetConstructedDataProtocolObjectTypesSupported() bool
}

// _BACnetConstructedDataProtocolObjectTypesSupported is the data-structure of this message
type _BACnetConstructedDataProtocolObjectTypesSupported struct {
	BACnetConstructedDataContract
	ProtocolObjectTypesSupported BACnetObjectTypesSupportedTagged
}

var _ BACnetConstructedDataProtocolObjectTypesSupported = (*_BACnetConstructedDataProtocolObjectTypesSupported)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_OBJECT_TYPES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetProtocolObjectTypesSupported() BACnetObjectTypesSupportedTagged {
	return m.ProtocolObjectTypesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetActualValue() BACnetObjectTypesSupportedTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectTypesSupportedTagged(m.GetProtocolObjectTypesSupported())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolObjectTypesSupported factory function for _BACnetConstructedDataProtocolObjectTypesSupported
func NewBACnetConstructedDataProtocolObjectTypesSupported(protocolObjectTypesSupported BACnetObjectTypesSupportedTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolObjectTypesSupported {
	_result := &_BACnetConstructedDataProtocolObjectTypesSupported{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProtocolObjectTypesSupported:  protocolObjectTypesSupported,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolObjectTypesSupported(structType any) BACnetConstructedDataProtocolObjectTypesSupported {
	if casted, ok := structType.(BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetTypeName() string {
	return "BACnetConstructedDataProtocolObjectTypesSupported"
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (protocolObjectTypesSupported)
	lengthInBits += m.ProtocolObjectTypesSupported.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProtocolObjectTypesSupported BACnetConstructedDataProtocolObjectTypesSupported, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolObjectTypesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolObjectTypesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolObjectTypesSupported, err := ReadSimpleField[BACnetObjectTypesSupportedTagged](ctx, "protocolObjectTypesSupported", ReadComplex[BACnetObjectTypesSupportedTagged](BACnetObjectTypesSupportedTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolObjectTypesSupported' field"))
	}
	m.ProtocolObjectTypesSupported = protocolObjectTypesSupported

	actualValue, err := ReadVirtualField[BACnetObjectTypesSupportedTagged](ctx, "actualValue", (*BACnetObjectTypesSupportedTagged)(nil), protocolObjectTypesSupported)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolObjectTypesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolObjectTypesSupported")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolObjectTypesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolObjectTypesSupported")
		}

		if err := WriteSimpleField[BACnetObjectTypesSupportedTagged](ctx, "protocolObjectTypesSupported", m.GetProtocolObjectTypesSupported(), WriteComplex[BACnetObjectTypesSupportedTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolObjectTypesSupported' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolObjectTypesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolObjectTypesSupported")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) isBACnetConstructedDataProtocolObjectTypesSupported() bool {
	return true
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
