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

// BACnetConstructedDataSlaveAddressBinding is the corresponding interface of BACnetConstructedDataSlaveAddressBinding
type BACnetConstructedDataSlaveAddressBinding interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSlaveAddressBinding returns SlaveAddressBinding (property field)
	GetSlaveAddressBinding() []BACnetAddressBinding
}

// BACnetConstructedDataSlaveAddressBindingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSlaveAddressBinding.
// This is useful for switch cases.
type BACnetConstructedDataSlaveAddressBindingExactly interface {
	BACnetConstructedDataSlaveAddressBinding
	isBACnetConstructedDataSlaveAddressBinding() bool
}

// _BACnetConstructedDataSlaveAddressBinding is the data-structure of this message
type _BACnetConstructedDataSlaveAddressBinding struct {
	BACnetConstructedDataContract
	SlaveAddressBinding []BACnetAddressBinding
}

var _ BACnetConstructedDataSlaveAddressBinding = (*_BACnetConstructedDataSlaveAddressBinding)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSlaveAddressBinding) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSlaveAddressBinding) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SLAVE_ADDRESS_BINDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSlaveAddressBinding) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSlaveAddressBinding) GetSlaveAddressBinding() []BACnetAddressBinding {
	return m.SlaveAddressBinding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSlaveAddressBinding factory function for _BACnetConstructedDataSlaveAddressBinding
func NewBACnetConstructedDataSlaveAddressBinding(slaveAddressBinding []BACnetAddressBinding, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSlaveAddressBinding {
	_result := &_BACnetConstructedDataSlaveAddressBinding{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SlaveAddressBinding:           slaveAddressBinding,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSlaveAddressBinding(structType any) BACnetConstructedDataSlaveAddressBinding {
	if casted, ok := structType.(BACnetConstructedDataSlaveAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSlaveAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSlaveAddressBinding) GetTypeName() string {
	return "BACnetConstructedDataSlaveAddressBinding"
}

func (m *_BACnetConstructedDataSlaveAddressBinding) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.SlaveAddressBinding) > 0 {
		for _, element := range m.SlaveAddressBinding {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSlaveAddressBinding) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSlaveAddressBinding) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSlaveAddressBinding BACnetConstructedDataSlaveAddressBinding, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSlaveAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSlaveAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	slaveAddressBinding, err := ReadTerminatedArrayField[BACnetAddressBinding](ctx, "slaveAddressBinding", ReadComplex[BACnetAddressBinding](BACnetAddressBindingParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'slaveAddressBinding' field"))
	}
	m.SlaveAddressBinding = slaveAddressBinding

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSlaveAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSlaveAddressBinding")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSlaveAddressBinding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSlaveAddressBinding) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSlaveAddressBinding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSlaveAddressBinding")
		}

		if err := WriteComplexTypeArrayField(ctx, "slaveAddressBinding", m.GetSlaveAddressBinding(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'slaveAddressBinding' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSlaveAddressBinding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSlaveAddressBinding")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSlaveAddressBinding) isBACnetConstructedDataSlaveAddressBinding() bool {
	return true
}

func (m *_BACnetConstructedDataSlaveAddressBinding) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
