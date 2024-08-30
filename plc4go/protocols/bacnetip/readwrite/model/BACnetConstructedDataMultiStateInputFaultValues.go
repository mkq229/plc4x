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

// BACnetConstructedDataMultiStateInputFaultValues is the corresponding interface of BACnetConstructedDataMultiStateInputFaultValues
type BACnetConstructedDataMultiStateInputFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultValues returns FaultValues (property field)
	GetFaultValues() []BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataMultiStateInputFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMultiStateInputFaultValues.
// This is useful for switch cases.
type BACnetConstructedDataMultiStateInputFaultValuesExactly interface {
	BACnetConstructedDataMultiStateInputFaultValues
	isBACnetConstructedDataMultiStateInputFaultValues() bool
}

// _BACnetConstructedDataMultiStateInputFaultValues is the data-structure of this message
type _BACnetConstructedDataMultiStateInputFaultValues struct {
	*_BACnetConstructedData
	FaultValues []BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_INPUT
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateInputFaultValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetFaultValues() []BACnetApplicationTagUnsignedInteger {
	return m.FaultValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMultiStateInputFaultValues factory function for _BACnetConstructedDataMultiStateInputFaultValues
func NewBACnetConstructedDataMultiStateInputFaultValues(faultValues []BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateInputFaultValues {
	_result := &_BACnetConstructedDataMultiStateInputFaultValues{
		FaultValues:            faultValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateInputFaultValues(structType any) BACnetConstructedDataMultiStateInputFaultValues {
	if casted, ok := structType.(BACnetConstructedDataMultiStateInputFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateInputFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetTypeName() string {
	return "BACnetConstructedDataMultiStateInputFaultValues"
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.FaultValues) > 0 {
		for _, element := range m.FaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataMultiStateInputFaultValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMultiStateInputFaultValues, error) {
	return BACnetConstructedDataMultiStateInputFaultValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataMultiStateInputFaultValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMultiStateInputFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateInputFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateInputFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultValues, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "faultValues", ReadComplex[BACnetApplicationTagUnsignedInteger](func(ctx context.Context, buffer utils.ReadBuffer) (BACnetApplicationTagUnsignedInteger, error) {
		v, err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		if err != nil {
			return nil, err
		}
		return v.(BACnetApplicationTagUnsignedInteger), nil
	}, readBuffer), func() bool { return IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber) })
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultValues' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateInputFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateInputFaultValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMultiStateInputFaultValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultValues: faultValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateInputFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateInputFaultValues")
		}

		// Array Field (faultValues)
		if pushErr := writeBuffer.PushContext("faultValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultValues")
		}
		for _curItem, _element := range m.GetFaultValues() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetFaultValues()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'faultValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("faultValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateInputFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateInputFaultValues")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) isBACnetConstructedDataMultiStateInputFaultValues() bool {
	return true
}

func (m *_BACnetConstructedDataMultiStateInputFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
