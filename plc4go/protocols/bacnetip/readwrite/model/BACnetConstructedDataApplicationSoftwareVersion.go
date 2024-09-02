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

// BACnetConstructedDataApplicationSoftwareVersion is the corresponding interface of BACnetConstructedDataApplicationSoftwareVersion
type BACnetConstructedDataApplicationSoftwareVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApplicationSoftwareVersion returns ApplicationSoftwareVersion (property field)
	GetApplicationSoftwareVersion() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataApplicationSoftwareVersionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataApplicationSoftwareVersion.
// This is useful for switch cases.
type BACnetConstructedDataApplicationSoftwareVersionExactly interface {
	BACnetConstructedDataApplicationSoftwareVersion
	isBACnetConstructedDataApplicationSoftwareVersion() bool
}

// _BACnetConstructedDataApplicationSoftwareVersion is the data-structure of this message
type _BACnetConstructedDataApplicationSoftwareVersion struct {
	BACnetConstructedDataContract
	ApplicationSoftwareVersion BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataApplicationSoftwareVersion = (*_BACnetConstructedDataApplicationSoftwareVersion)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APPLICATION_SOFTWARE_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetApplicationSoftwareVersion() BACnetApplicationTagCharacterString {
	return m.ApplicationSoftwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetApplicationSoftwareVersion())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataApplicationSoftwareVersion factory function for _BACnetConstructedDataApplicationSoftwareVersion
func NewBACnetConstructedDataApplicationSoftwareVersion(applicationSoftwareVersion BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataApplicationSoftwareVersion {
	_result := &_BACnetConstructedDataApplicationSoftwareVersion{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ApplicationSoftwareVersion:    applicationSoftwareVersion,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataApplicationSoftwareVersion(structType any) BACnetConstructedDataApplicationSoftwareVersion {
	if casted, ok := structType.(BACnetConstructedDataApplicationSoftwareVersion); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataApplicationSoftwareVersion); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetTypeName() string {
	return "BACnetConstructedDataApplicationSoftwareVersion"
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (applicationSoftwareVersion)
	lengthInBits += m.ApplicationSoftwareVersion.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataApplicationSoftwareVersion BACnetConstructedDataApplicationSoftwareVersion, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataApplicationSoftwareVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataApplicationSoftwareVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	applicationSoftwareVersion, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "applicationSoftwareVersion", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationSoftwareVersion' field"))
	}
	m.ApplicationSoftwareVersion = applicationSoftwareVersion

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), applicationSoftwareVersion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataApplicationSoftwareVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataApplicationSoftwareVersion")
	}

	return m, nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataApplicationSoftwareVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataApplicationSoftwareVersion")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "applicationSoftwareVersion", m.GetApplicationSoftwareVersion(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'applicationSoftwareVersion' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataApplicationSoftwareVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataApplicationSoftwareVersion")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) isBACnetConstructedDataApplicationSoftwareVersion() bool {
	return true
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
