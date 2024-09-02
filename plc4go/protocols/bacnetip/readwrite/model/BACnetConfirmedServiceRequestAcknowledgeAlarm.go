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

// BACnetConfirmedServiceRequestAcknowledgeAlarm is the corresponding interface of BACnetConfirmedServiceRequestAcknowledgeAlarm
type BACnetConfirmedServiceRequestAcknowledgeAlarm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetAcknowledgingProcessIdentifier returns AcknowledgingProcessIdentifier (property field)
	GetAcknowledgingProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetEventStateAcknowledged returns EventStateAcknowledged (property field)
	GetEventStateAcknowledged() BACnetEventStateTagged
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetAcknowledgmentSource returns AcknowledgmentSource (property field)
	GetAcknowledgmentSource() BACnetContextTagCharacterString
	// GetTimeOfAcknowledgment returns TimeOfAcknowledgment (property field)
	GetTimeOfAcknowledgment() BACnetTimeStampEnclosed
}

// BACnetConfirmedServiceRequestAcknowledgeAlarmExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestAcknowledgeAlarm.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestAcknowledgeAlarmExactly interface {
	BACnetConfirmedServiceRequestAcknowledgeAlarm
	isBACnetConfirmedServiceRequestAcknowledgeAlarm() bool
}

// _BACnetConfirmedServiceRequestAcknowledgeAlarm is the data-structure of this message
type _BACnetConfirmedServiceRequestAcknowledgeAlarm struct {
	BACnetConfirmedServiceRequestContract
	AcknowledgingProcessIdentifier BACnetContextTagUnsignedInteger
	EventObjectIdentifier          BACnetContextTagObjectIdentifier
	EventStateAcknowledged         BACnetEventStateTagged
	Timestamp                      BACnetTimeStampEnclosed
	AcknowledgmentSource           BACnetContextTagCharacterString
	TimeOfAcknowledgment           BACnetTimeStampEnclosed
}

var _ BACnetConfirmedServiceRequestAcknowledgeAlarm = (*_BACnetConfirmedServiceRequestAcknowledgeAlarm)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetAcknowledgingProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.AcknowledgingProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetEventObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetEventStateAcknowledged() BACnetEventStateTagged {
	return m.EventStateAcknowledged
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetAcknowledgmentSource() BACnetContextTagCharacterString {
	return m.AcknowledgmentSource
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTimeOfAcknowledgment() BACnetTimeStampEnclosed {
	return m.TimeOfAcknowledgment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAcknowledgeAlarm factory function for _BACnetConfirmedServiceRequestAcknowledgeAlarm
func NewBACnetConfirmedServiceRequestAcknowledgeAlarm(acknowledgingProcessIdentifier BACnetContextTagUnsignedInteger, eventObjectIdentifier BACnetContextTagObjectIdentifier, eventStateAcknowledged BACnetEventStateTagged, timestamp BACnetTimeStampEnclosed, acknowledgmentSource BACnetContextTagCharacterString, timeOfAcknowledgment BACnetTimeStampEnclosed, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAcknowledgeAlarm {
	_result := &_BACnetConfirmedServiceRequestAcknowledgeAlarm{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		AcknowledgingProcessIdentifier:        acknowledgingProcessIdentifier,
		EventObjectIdentifier:                 eventObjectIdentifier,
		EventStateAcknowledged:                eventStateAcknowledged,
		Timestamp:                             timestamp,
		AcknowledgmentSource:                  acknowledgmentSource,
		TimeOfAcknowledgment:                  timeOfAcknowledgment,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAcknowledgeAlarm(structType any) BACnetConfirmedServiceRequestAcknowledgeAlarm {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAcknowledgeAlarm"
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (acknowledgingProcessIdentifier)
	lengthInBits += m.AcknowledgingProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (eventStateAcknowledged)
	lengthInBits += m.EventStateAcknowledged.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (acknowledgmentSource)
	lengthInBits += m.AcknowledgmentSource.GetLengthInBits(ctx)

	// Simple field (timeOfAcknowledgment)
	lengthInBits += m.TimeOfAcknowledgment.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAcknowledgeAlarm BACnetConfirmedServiceRequestAcknowledgeAlarm, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAcknowledgeAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	acknowledgingProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "acknowledgingProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgingProcessIdentifier' field"))
	}
	m.AcknowledgingProcessIdentifier = acknowledgingProcessIdentifier

	eventObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventObjectIdentifier' field"))
	}
	m.EventObjectIdentifier = eventObjectIdentifier

	eventStateAcknowledged, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventStateAcknowledged", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventStateAcknowledged' field"))
	}
	m.EventStateAcknowledged = eventStateAcknowledged

	timestamp, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	acknowledgmentSource, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "acknowledgmentSource", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgmentSource' field"))
	}
	m.AcknowledgmentSource = acknowledgmentSource

	timeOfAcknowledgment, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timeOfAcknowledgment", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(5))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeOfAcknowledgment' field"))
	}
	m.TimeOfAcknowledgment = timeOfAcknowledgment

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAcknowledgeAlarm")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAcknowledgeAlarm")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "acknowledgingProcessIdentifier", m.GetAcknowledgingProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'acknowledgingProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", m.GetEventObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventObjectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "eventStateAcknowledged", m.GetEventStateAcknowledged(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventStateAcknowledged' field")
		}

		if err := WriteSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetTimeStampEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "acknowledgmentSource", m.GetAcknowledgmentSource(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'acknowledgmentSource' field")
		}

		if err := WriteSimpleField[BACnetTimeStampEnclosed](ctx, "timeOfAcknowledgment", m.GetTimeOfAcknowledgment(), WriteComplex[BACnetTimeStampEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeOfAcknowledgment' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAcknowledgeAlarm")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) isBACnetConfirmedServiceRequestAcknowledgeAlarm() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestAcknowledgeAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
