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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataLifeSafetyZoneFaultValues is the corresponding interface of BACnetConstructedDataLifeSafetyZoneFaultValues
type BACnetConstructedDataLifeSafetyZoneFaultValues interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultValues returns FaultValues (property field)
	GetFaultValues() []BACnetLifeSafetyStateTagged
}

// BACnetConstructedDataLifeSafetyZoneFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLifeSafetyZoneFaultValues.
// This is useful for switch cases.
type BACnetConstructedDataLifeSafetyZoneFaultValuesExactly interface {
	BACnetConstructedDataLifeSafetyZoneFaultValues
	isBACnetConstructedDataLifeSafetyZoneFaultValues() bool
}

// _BACnetConstructedDataLifeSafetyZoneFaultValues is the data-structure of this message
type _BACnetConstructedDataLifeSafetyZoneFaultValues struct {
	*_BACnetConstructedData
	FaultValues []BACnetLifeSafetyStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetFaultValues() []BACnetLifeSafetyStateTagged {
	return m.FaultValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyZoneFaultValues factory function for _BACnetConstructedDataLifeSafetyZoneFaultValues
func NewBACnetConstructedDataLifeSafetyZoneFaultValues(faultValues []BACnetLifeSafetyStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyZoneFaultValues {
	_result := &_BACnetConstructedDataLifeSafetyZoneFaultValues{
		FaultValues:            faultValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyZoneFaultValues(structType interface{}) BACnetConstructedDataLifeSafetyZoneFaultValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneFaultValues"
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.FaultValues) > 0 {
		for _, element := range m.FaultValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyZoneFaultValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLifeSafetyZoneFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (faultValues)
	if pullErr := readBuffer.PullContext("faultValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultValues")
	}
	// Terminated array
	var faultValues []BACnetLifeSafetyStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'faultValues' field of BACnetConstructedDataLifeSafetyZoneFaultValues")
			}
			faultValues = append(faultValues, _item.(BACnetLifeSafetyStateTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("faultValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneFaultValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLifeSafetyZoneFaultValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultValues: faultValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneFaultValues")
		}

		// Array Field (faultValues)
		if pushErr := writeBuffer.PushContext("faultValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultValues")
		}
		for _, _element := range m.GetFaultValues() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'faultValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("faultValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneFaultValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) isBACnetConstructedDataLifeSafetyZoneFaultValues() bool {
	return true
}

func (m *_BACnetConstructedDataLifeSafetyZoneFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}