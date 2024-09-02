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

// ModifySubscriptionRequest is the corresponding interface of ModifySubscriptionRequest
type ModifySubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetRequestedPublishingInterval returns RequestedPublishingInterval (property field)
	GetRequestedPublishingInterval() float64
	// GetRequestedLifetimeCount returns RequestedLifetimeCount (property field)
	GetRequestedLifetimeCount() uint32
	// GetRequestedMaxKeepAliveCount returns RequestedMaxKeepAliveCount (property field)
	GetRequestedMaxKeepAliveCount() uint32
	// GetMaxNotificationsPerPublish returns MaxNotificationsPerPublish (property field)
	GetMaxNotificationsPerPublish() uint32
	// GetPriority returns Priority (property field)
	GetPriority() uint8
}

// ModifySubscriptionRequestExactly can be used when we want exactly this type and not a type which fulfills ModifySubscriptionRequest.
// This is useful for switch cases.
type ModifySubscriptionRequestExactly interface {
	ModifySubscriptionRequest
	isModifySubscriptionRequest() bool
}

// _ModifySubscriptionRequest is the data-structure of this message
type _ModifySubscriptionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader               ExtensionObjectDefinition
	SubscriptionId              uint32
	RequestedPublishingInterval float64
	RequestedLifetimeCount      uint32
	RequestedMaxKeepAliveCount  uint32
	MaxNotificationsPerPublish  uint32
	Priority                    uint8
}

var _ ModifySubscriptionRequest = (*_ModifySubscriptionRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModifySubscriptionRequest) GetIdentifier() string {
	return "793"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModifySubscriptionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModifySubscriptionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_ModifySubscriptionRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_ModifySubscriptionRequest) GetRequestedPublishingInterval() float64 {
	return m.RequestedPublishingInterval
}

func (m *_ModifySubscriptionRequest) GetRequestedLifetimeCount() uint32 {
	return m.RequestedLifetimeCount
}

func (m *_ModifySubscriptionRequest) GetRequestedMaxKeepAliveCount() uint32 {
	return m.RequestedMaxKeepAliveCount
}

func (m *_ModifySubscriptionRequest) GetMaxNotificationsPerPublish() uint32 {
	return m.MaxNotificationsPerPublish
}

func (m *_ModifySubscriptionRequest) GetPriority() uint8 {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModifySubscriptionRequest factory function for _ModifySubscriptionRequest
func NewModifySubscriptionRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, priority uint8) *_ModifySubscriptionRequest {
	_result := &_ModifySubscriptionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		RequestedPublishingInterval:       requestedPublishingInterval,
		RequestedLifetimeCount:            requestedLifetimeCount,
		RequestedMaxKeepAliveCount:        requestedMaxKeepAliveCount,
		MaxNotificationsPerPublish:        maxNotificationsPerPublish,
		Priority:                          priority,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModifySubscriptionRequest(structType any) ModifySubscriptionRequest {
	if casted, ok := structType.(ModifySubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModifySubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModifySubscriptionRequest) GetTypeName() string {
	return "ModifySubscriptionRequest"
}

func (m *_ModifySubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (requestedPublishingInterval)
	lengthInBits += 64

	// Simple field (requestedLifetimeCount)
	lengthInBits += 32

	// Simple field (requestedMaxKeepAliveCount)
	lengthInBits += 32

	// Simple field (maxNotificationsPerPublish)
	lengthInBits += 32

	// Simple field (priority)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModifySubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModifySubscriptionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__modifySubscriptionRequest ModifySubscriptionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModifySubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModifySubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	requestedPublishingInterval, err := ReadSimpleField(ctx, "requestedPublishingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedPublishingInterval' field"))
	}
	m.RequestedPublishingInterval = requestedPublishingInterval

	requestedLifetimeCount, err := ReadSimpleField(ctx, "requestedLifetimeCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedLifetimeCount' field"))
	}
	m.RequestedLifetimeCount = requestedLifetimeCount

	requestedMaxKeepAliveCount, err := ReadSimpleField(ctx, "requestedMaxKeepAliveCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedMaxKeepAliveCount' field"))
	}
	m.RequestedMaxKeepAliveCount = requestedMaxKeepAliveCount

	maxNotificationsPerPublish, err := ReadSimpleField(ctx, "maxNotificationsPerPublish", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationsPerPublish' field"))
	}
	m.MaxNotificationsPerPublish = maxNotificationsPerPublish

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	if closeErr := readBuffer.CloseContext("ModifySubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModifySubscriptionRequest")
	}

	return m, nil
}

func (m *_ModifySubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModifySubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModifySubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModifySubscriptionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[float64](ctx, "requestedPublishingInterval", m.GetRequestedPublishingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedPublishingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedLifetimeCount", m.GetRequestedLifetimeCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedLifetimeCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedMaxKeepAliveCount", m.GetRequestedMaxKeepAliveCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedMaxKeepAliveCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNotificationsPerPublish", m.GetMaxNotificationsPerPublish(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNotificationsPerPublish' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("ModifySubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModifySubscriptionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModifySubscriptionRequest) isModifySubscriptionRequest() bool {
	return true
}

func (m *_ModifySubscriptionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
