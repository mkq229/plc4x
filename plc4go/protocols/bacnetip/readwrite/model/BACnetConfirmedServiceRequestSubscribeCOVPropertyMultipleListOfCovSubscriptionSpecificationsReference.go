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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference interface {
	utils.LengthAware
	utils.Serializable
	// GetMonitoredProperty returns MonitoredProperty (property field)
	GetMonitoredProperty() BACnetPropertyReferenceEnclosed
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
	// GetTimestamped returns Timestamped (property field)
	GetTimestamped() BACnetContextTagBoolean
}

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReferenceExactly interface {
	BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
	isBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference() bool
}

// _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference struct {
	MonitoredProperty BACnetPropertyReferenceEnclosed
	CovIncrement      BACnetContextTagReal
	Timestamped       BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetMonitoredProperty() BACnetPropertyReferenceEnclosed {
	return m.MonitoredProperty
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetTimestamped() BACnetContextTagBoolean {
	return m.Timestamped
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference factory function for _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(monitoredProperty BACnetPropertyReferenceEnclosed, covIncrement BACnetContextTagReal, timestamped BACnetContextTagBoolean) *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference {
	return &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference{MonitoredProperty: monitoredProperty, CovIncrement: covIncrement, Timestamped: timestamped}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(structType interface{}) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredProperty)
	lengthInBits += m.MonitoredProperty.GetLengthInBits()

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits()
	}

	// Simple field (timestamped)
	lengthInBits += m.Timestamped.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReferenceParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (monitoredProperty)
	if pullErr := readBuffer.PullContext("monitoredProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredProperty")
	}
	_monitoredProperty, _monitoredPropertyErr := BACnetPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _monitoredPropertyErr != nil {
		return nil, errors.Wrap(_monitoredPropertyErr, "Error parsing 'monitoredProperty' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
	}
	monitoredProperty := _monitoredProperty.(BACnetPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("monitoredProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredProperty")
	}

	// Optional Field (covIncrement) (Can be skipped, if a given expression evaluates to false)
	var covIncrement BACnetContextTagReal = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_REAL)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'covIncrement' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
		default:
			covIncrement = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
			}
		}
	}

	// Simple Field (timestamped)
	if pullErr := readBuffer.PullContext("timestamped"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamped")
	}
	_timestamped, _timestampedErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _timestampedErr != nil {
		return nil, errors.Wrap(_timestampedErr, "Error parsing 'timestamped' field of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
	}
	timestamped := _timestamped.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("timestamped"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamped")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
	}

	// Create the instance
	return &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference{
		MonitoredProperty: monitoredProperty,
		CovIncrement:      covIncrement,
		Timestamped:       timestamped,
	}, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
	}

	// Simple Field (monitoredProperty)
	if pushErr := writeBuffer.PushContext("monitoredProperty"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredProperty")
	}
	_monitoredPropertyErr := writeBuffer.WriteSerializable(m.GetMonitoredProperty())
	if popErr := writeBuffer.PopContext("monitoredProperty"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredProperty")
	}
	if _monitoredPropertyErr != nil {
		return errors.Wrap(_monitoredPropertyErr, "Error serializing 'monitoredProperty' field")
	}

	// Optional Field (covIncrement) (Can be skipped, if the value is null)
	var covIncrement BACnetContextTagReal = nil
	if m.GetCovIncrement() != nil {
		if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covIncrement")
		}
		covIncrement = m.GetCovIncrement()
		_covIncrementErr := writeBuffer.WriteSerializable(covIncrement)
		if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covIncrement")
		}
		if _covIncrementErr != nil {
			return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
		}
	}

	// Simple Field (timestamped)
	if pushErr := writeBuffer.PushContext("timestamped"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamped")
	}
	_timestampedErr := writeBuffer.WriteSerializable(m.GetTimestamped())
	if popErr := writeBuffer.PopContext("timestamped"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamped")
	}
	if _timestampedErr != nil {
		return errors.Wrap(_timestampedErr, "Error serializing 'timestamped' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference")
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) isBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}