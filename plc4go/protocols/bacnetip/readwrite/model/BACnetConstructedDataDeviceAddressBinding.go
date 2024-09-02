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

// BACnetConstructedDataDeviceAddressBinding is the corresponding interface of BACnetConstructedDataDeviceAddressBinding
type BACnetConstructedDataDeviceAddressBinding interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDeviceAddressBinding returns DeviceAddressBinding (property field)
	GetDeviceAddressBinding() []BACnetAddressBinding
}

// BACnetConstructedDataDeviceAddressBindingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDeviceAddressBinding.
// This is useful for switch cases.
type BACnetConstructedDataDeviceAddressBindingExactly interface {
	BACnetConstructedDataDeviceAddressBinding
	isBACnetConstructedDataDeviceAddressBinding() bool
}

// _BACnetConstructedDataDeviceAddressBinding is the data-structure of this message
type _BACnetConstructedDataDeviceAddressBinding struct {
	BACnetConstructedDataContract
	DeviceAddressBinding []BACnetAddressBinding
}

var _ BACnetConstructedDataDeviceAddressBinding = (*_BACnetConstructedDataDeviceAddressBinding)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeviceAddressBinding) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEVICE_ADDRESS_BINDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeviceAddressBinding) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceAddressBinding) GetDeviceAddressBinding() []BACnetAddressBinding {
	return m.DeviceAddressBinding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeviceAddressBinding factory function for _BACnetConstructedDataDeviceAddressBinding
func NewBACnetConstructedDataDeviceAddressBinding(deviceAddressBinding []BACnetAddressBinding, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeviceAddressBinding {
	_result := &_BACnetConstructedDataDeviceAddressBinding{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DeviceAddressBinding:          deviceAddressBinding,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeviceAddressBinding(structType any) BACnetConstructedDataDeviceAddressBinding {
	if casted, ok := structType.(BACnetConstructedDataDeviceAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetTypeName() string {
	return "BACnetConstructedDataDeviceAddressBinding"
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.DeviceAddressBinding) > 0 {
		for _, element := range m.DeviceAddressBinding {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataDeviceAddressBinding) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDeviceAddressBinding) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDeviceAddressBinding BACnetConstructedDataDeviceAddressBinding, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceAddressBinding, err := ReadTerminatedArrayField[BACnetAddressBinding](ctx, "deviceAddressBinding", ReadComplex[BACnetAddressBinding](BACnetAddressBindingParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceAddressBinding' field"))
	}
	m.DeviceAddressBinding = deviceAddressBinding

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceAddressBinding")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDeviceAddressBinding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeviceAddressBinding) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceAddressBinding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceAddressBinding")
		}

		if err := WriteComplexTypeArrayField(ctx, "deviceAddressBinding", m.GetDeviceAddressBinding(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceAddressBinding' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceAddressBinding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceAddressBinding")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeviceAddressBinding) isBACnetConstructedDataDeviceAddressBinding() bool {
	return true
}

func (m *_BACnetConstructedDataDeviceAddressBinding) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
