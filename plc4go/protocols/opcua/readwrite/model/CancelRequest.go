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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CancelRequest is the corresponding interface of CancelRequest
type CancelRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetRequestHandle returns RequestHandle (property field)
	GetRequestHandle() uint32
}

// CancelRequestExactly can be used when we want exactly this type and not a type which fulfills CancelRequest.
// This is useful for switch cases.
type CancelRequestExactly interface {
	CancelRequest
	isCancelRequest() bool
}

// _CancelRequest is the data-structure of this message
type _CancelRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader ExtensionObjectDefinition
	RequestHandle uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CancelRequest) GetIdentifier() string {
	return "479"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CancelRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CancelRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CancelRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CancelRequest) GetRequestHandle() uint32 {
	return m.RequestHandle
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCancelRequest factory function for _CancelRequest
func NewCancelRequest(requestHeader ExtensionObjectDefinition, requestHandle uint32) *_CancelRequest {
	_result := &_CancelRequest{
		RequestHeader:              requestHeader,
		RequestHandle:              requestHandle,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCancelRequest(structType any) CancelRequest {
	if casted, ok := structType.(CancelRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CancelRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CancelRequest) GetTypeName() string {
	return "CancelRequest"
}

func (m *_CancelRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (requestHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CancelRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CancelRequestParse(ctx context.Context, theBytes []byte, identifier string) (CancelRequest, error) {
	return CancelRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CancelRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CancelRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CancelRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CancelRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of CancelRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (requestHandle)
	_requestHandle, _requestHandleErr := /*TODO: migrate me*/ readBuffer.ReadUint32("requestHandle", 32)
	if _requestHandleErr != nil {
		return nil, errors.Wrap(_requestHandleErr, "Error parsing 'requestHandle' field of CancelRequest")
	}
	requestHandle := _requestHandle

	if closeErr := readBuffer.CloseContext("CancelRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CancelRequest")
	}

	// Create a partially initialized instance
	_child := &_CancelRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		RequestHandle:              requestHandle,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CancelRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CancelRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CancelRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CancelRequest")
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

		// Simple Field (requestHandle)
		requestHandle := uint32(m.GetRequestHandle())
		_requestHandleErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("requestHandle", 32, uint32((requestHandle)))
		if _requestHandleErr != nil {
			return errors.Wrap(_requestHandleErr, "Error serializing 'requestHandle' field")
		}

		if popErr := writeBuffer.PopContext("CancelRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CancelRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CancelRequest) isCancelRequest() bool {
	return true
}

func (m *_CancelRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
