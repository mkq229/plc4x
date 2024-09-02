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

// GetEndpointsResponse is the corresponding interface of GetEndpointsResponse
type GetEndpointsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfEndpoints returns NoOfEndpoints (property field)
	GetNoOfEndpoints() int32
	// GetEndpoints returns Endpoints (property field)
	GetEndpoints() []ExtensionObjectDefinition
}

// GetEndpointsResponseExactly can be used when we want exactly this type and not a type which fulfills GetEndpointsResponse.
// This is useful for switch cases.
type GetEndpointsResponseExactly interface {
	GetEndpointsResponse
	isGetEndpointsResponse() bool
}

// _GetEndpointsResponse is the data-structure of this message
type _GetEndpointsResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
	NoOfEndpoints  int32
	Endpoints      []ExtensionObjectDefinition
}

var _ GetEndpointsResponse = (*_GetEndpointsResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetEndpointsResponse) GetIdentifier() string {
	return "431"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetEndpointsResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetEndpointsResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_GetEndpointsResponse) GetNoOfEndpoints() int32 {
	return m.NoOfEndpoints
}

func (m *_GetEndpointsResponse) GetEndpoints() []ExtensionObjectDefinition {
	return m.Endpoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGetEndpointsResponse factory function for _GetEndpointsResponse
func NewGetEndpointsResponse(responseHeader ExtensionObjectDefinition, noOfEndpoints int32, endpoints []ExtensionObjectDefinition) *_GetEndpointsResponse {
	_result := &_GetEndpointsResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfEndpoints:                     noOfEndpoints,
		Endpoints:                         endpoints,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetEndpointsResponse(structType any) GetEndpointsResponse {
	if casted, ok := structType.(GetEndpointsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*GetEndpointsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_GetEndpointsResponse) GetTypeName() string {
	return "GetEndpointsResponse"
}

func (m *_GetEndpointsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfEndpoints)
	lengthInBits += 32

	// Array field
	if len(m.Endpoints) > 0 {
		for _curItem, element := range m.Endpoints {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Endpoints), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_GetEndpointsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetEndpointsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__getEndpointsResponse GetEndpointsResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetEndpointsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetEndpointsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfEndpoints, err := ReadSimpleField(ctx, "noOfEndpoints", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEndpoints' field"))
	}
	m.NoOfEndpoints = noOfEndpoints

	endpoints, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "endpoints", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("314")), readBuffer), uint64(noOfEndpoints))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpoints' field"))
	}
	m.Endpoints = endpoints

	if closeErr := readBuffer.CloseContext("GetEndpointsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetEndpointsResponse")
	}

	return m, nil
}

func (m *_GetEndpointsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetEndpointsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetEndpointsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetEndpointsResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfEndpoints", m.GetNoOfEndpoints(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEndpoints' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "endpoints", m.GetEndpoints(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'endpoints' field")
		}

		if popErr := writeBuffer.PopContext("GetEndpointsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetEndpointsResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetEndpointsResponse) isGetEndpointsResponse() bool {
	return true
}

func (m *_GetEndpointsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
