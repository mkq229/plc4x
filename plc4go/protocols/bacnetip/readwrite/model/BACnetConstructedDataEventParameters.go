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

// BACnetConstructedDataEventParameters is the corresponding interface of BACnetConstructedDataEventParameters
type BACnetConstructedDataEventParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventParameter returns EventParameter (property field)
	GetEventParameter() BACnetEventParameter
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventParameter
}

// BACnetConstructedDataEventParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventParameters.
// This is useful for switch cases.
type BACnetConstructedDataEventParametersExactly interface {
	BACnetConstructedDataEventParameters
	isBACnetConstructedDataEventParameters() bool
}

// _BACnetConstructedDataEventParameters is the data-structure of this message
type _BACnetConstructedDataEventParameters struct {
	BACnetConstructedDataContract
	EventParameter BACnetEventParameter
}

var _ BACnetConstructedDataEventParameters = (*_BACnetConstructedDataEventParameters)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventParameters) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_PARAMETERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventParameters) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetEventParameter() BACnetEventParameter {
	return m.EventParameter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventParameters) GetActualValue() BACnetEventParameter {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventParameter(m.GetEventParameter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventParameters factory function for _BACnetConstructedDataEventParameters
func NewBACnetConstructedDataEventParameters(eventParameter BACnetEventParameter, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventParameters {
	_result := &_BACnetConstructedDataEventParameters{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventParameter:                eventParameter,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventParameters(structType any) BACnetConstructedDataEventParameters {
	if casted, ok := structType.(BACnetConstructedDataEventParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventParameters) GetTypeName() string {
	return "BACnetConstructedDataEventParameters"
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventParameter)
	lengthInBits += m.EventParameter.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventParameters BACnetConstructedDataEventParameters, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventParameter, err := ReadSimpleField[BACnetEventParameter](ctx, "eventParameter", ReadComplex[BACnetEventParameter](BACnetEventParameterParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventParameter' field"))
	}
	m.EventParameter = eventParameter

	actualValue, err := ReadVirtualField[BACnetEventParameter](ctx, "actualValue", (*BACnetEventParameter)(nil), eventParameter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventParameters")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventParameters")
		}

		if err := WriteSimpleField[BACnetEventParameter](ctx, "eventParameter", m.GetEventParameter(), WriteComplex[BACnetEventParameter](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventParameter' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventParameters")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventParameters) isBACnetConstructedDataEventParameters() bool {
	return true
}

func (m *_BACnetConstructedDataEventParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
