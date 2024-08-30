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

// DeleteSubscriptionsRequest is the corresponding interface of DeleteSubscriptionsRequest
type DeleteSubscriptionsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfSubscriptionIds returns NoOfSubscriptionIds (property field)
	GetNoOfSubscriptionIds() int32
	// GetSubscriptionIds returns SubscriptionIds (property field)
	GetSubscriptionIds() []uint32
}

// DeleteSubscriptionsRequestExactly can be used when we want exactly this type and not a type which fulfills DeleteSubscriptionsRequest.
// This is useful for switch cases.
type DeleteSubscriptionsRequestExactly interface {
	DeleteSubscriptionsRequest
	isDeleteSubscriptionsRequest() bool
}

// _DeleteSubscriptionsRequest is the data-structure of this message
type _DeleteSubscriptionsRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader       ExtensionObjectDefinition
	NoOfSubscriptionIds int32
	SubscriptionIds     []uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteSubscriptionsRequest) GetIdentifier() string {
	return "847"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteSubscriptionsRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteSubscriptionsRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteSubscriptionsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteSubscriptionsRequest) GetNoOfSubscriptionIds() int32 {
	return m.NoOfSubscriptionIds
}

func (m *_DeleteSubscriptionsRequest) GetSubscriptionIds() []uint32 {
	return m.SubscriptionIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteSubscriptionsRequest factory function for _DeleteSubscriptionsRequest
func NewDeleteSubscriptionsRequest(requestHeader ExtensionObjectDefinition, noOfSubscriptionIds int32, subscriptionIds []uint32) *_DeleteSubscriptionsRequest {
	_result := &_DeleteSubscriptionsRequest{
		RequestHeader:              requestHeader,
		NoOfSubscriptionIds:        noOfSubscriptionIds,
		SubscriptionIds:            subscriptionIds,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteSubscriptionsRequest(structType any) DeleteSubscriptionsRequest {
	if casted, ok := structType.(DeleteSubscriptionsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteSubscriptionsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteSubscriptionsRequest) GetTypeName() string {
	return "DeleteSubscriptionsRequest"
}

func (m *_DeleteSubscriptionsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfSubscriptionIds)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionIds) > 0 {
		lengthInBits += 32 * uint16(len(m.SubscriptionIds))
	}

	return lengthInBits
}

func (m *_DeleteSubscriptionsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteSubscriptionsRequestParse(ctx context.Context, theBytes []byte, identifier string) (DeleteSubscriptionsRequest, error) {
	return DeleteSubscriptionsRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteSubscriptionsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteSubscriptionsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeleteSubscriptionsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteSubscriptionsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of DeleteSubscriptionsRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfSubscriptionIds)
	_noOfSubscriptionIds, _noOfSubscriptionIdsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfSubscriptionIds", 32)
	if _noOfSubscriptionIdsErr != nil {
		return nil, errors.Wrap(_noOfSubscriptionIdsErr, "Error parsing 'noOfSubscriptionIds' field of DeleteSubscriptionsRequest")
	}
	noOfSubscriptionIds := _noOfSubscriptionIds

	subscriptionIds, err := ReadCountArrayField[uint32](ctx, "subscriptionIds", ReadUnsignedInt(readBuffer, 32), uint64(noOfSubscriptionIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionIds' field"))
	}

	if closeErr := readBuffer.CloseContext("DeleteSubscriptionsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteSubscriptionsRequest")
	}

	// Create a partially initialized instance
	_child := &_DeleteSubscriptionsRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfSubscriptionIds:        noOfSubscriptionIds,
		SubscriptionIds:            subscriptionIds,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteSubscriptionsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteSubscriptionsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteSubscriptionsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteSubscriptionsRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (noOfSubscriptionIds)
		noOfSubscriptionIds := int32(m.GetNoOfSubscriptionIds())
		_noOfSubscriptionIdsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfSubscriptionIds", 32, int32((noOfSubscriptionIds)))
		if _noOfSubscriptionIdsErr != nil {
			return errors.Wrap(_noOfSubscriptionIdsErr, "Error serializing 'noOfSubscriptionIds' field")
		}

		// Array Field (subscriptionIds)
		if pushErr := writeBuffer.PushContext("subscriptionIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriptionIds")
		}
		for _curItem, _element := range m.GetSubscriptionIds() {
			_ = _curItem
			_elementErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("", 32, uint32(_element))
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'subscriptionIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("subscriptionIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriptionIds")
		}

		if popErr := writeBuffer.PopContext("DeleteSubscriptionsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteSubscriptionsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteSubscriptionsRequest) isDeleteSubscriptionsRequest() bool {
	return true
}

func (m *_DeleteSubscriptionsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
