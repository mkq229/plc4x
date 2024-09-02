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

// BACnetConstructedDataDeployedProfileLocation is the corresponding interface of BACnetConstructedDataDeployedProfileLocation
type BACnetConstructedDataDeployedProfileLocation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDeployedProfileLocation returns DeployedProfileLocation (property field)
	GetDeployedProfileLocation() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataDeployedProfileLocationExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDeployedProfileLocation.
// This is useful for switch cases.
type BACnetConstructedDataDeployedProfileLocationExactly interface {
	BACnetConstructedDataDeployedProfileLocation
	isBACnetConstructedDataDeployedProfileLocation() bool
}

// _BACnetConstructedDataDeployedProfileLocation is the data-structure of this message
type _BACnetConstructedDataDeployedProfileLocation struct {
	BACnetConstructedDataContract
	DeployedProfileLocation BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataDeployedProfileLocation = (*_BACnetConstructedDataDeployedProfileLocation)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEPLOYED_PROFILE_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetDeployedProfileLocation() BACnetApplicationTagCharacterString {
	return m.DeployedProfileLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDeployedProfileLocation())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeployedProfileLocation factory function for _BACnetConstructedDataDeployedProfileLocation
func NewBACnetConstructedDataDeployedProfileLocation(deployedProfileLocation BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeployedProfileLocation {
	_result := &_BACnetConstructedDataDeployedProfileLocation{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DeployedProfileLocation:       deployedProfileLocation,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeployedProfileLocation(structType any) BACnetConstructedDataDeployedProfileLocation {
	if casted, ok := structType.(BACnetConstructedDataDeployedProfileLocation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeployedProfileLocation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetTypeName() string {
	return "BACnetConstructedDataDeployedProfileLocation"
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (deployedProfileLocation)
	lengthInBits += m.DeployedProfileLocation.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDeployedProfileLocation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDeployedProfileLocation BACnetConstructedDataDeployedProfileLocation, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeployedProfileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeployedProfileLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deployedProfileLocation, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "deployedProfileLocation", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deployedProfileLocation' field"))
	}
	m.DeployedProfileLocation = deployedProfileLocation

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), deployedProfileLocation)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeployedProfileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeployedProfileLocation")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDeployedProfileLocation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeployedProfileLocation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeployedProfileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeployedProfileLocation")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "deployedProfileLocation", m.GetDeployedProfileLocation(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deployedProfileLocation' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeployedProfileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeployedProfileLocation")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeployedProfileLocation) isBACnetConstructedDataDeployedProfileLocation() bool {
	return true
}

func (m *_BACnetConstructedDataDeployedProfileLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
