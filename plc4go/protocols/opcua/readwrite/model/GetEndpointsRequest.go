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

// GetEndpointsRequest is the corresponding interface of GetEndpointsRequest
type GetEndpointsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetNoOfLocaleIds returns NoOfLocaleIds (property field)
	GetNoOfLocaleIds() int32
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetNoOfProfileUris returns NoOfProfileUris (property field)
	GetNoOfProfileUris() int32
	// GetProfileUris returns ProfileUris (property field)
	GetProfileUris() []PascalString
}

// GetEndpointsRequestExactly can be used when we want exactly this type and not a type which fulfills GetEndpointsRequest.
// This is useful for switch cases.
type GetEndpointsRequestExactly interface {
	GetEndpointsRequest
	isGetEndpointsRequest() bool
}

// _GetEndpointsRequest is the data-structure of this message
type _GetEndpointsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader   ExtensionObjectDefinition
	EndpointUrl     PascalString
	NoOfLocaleIds   int32
	LocaleIds       []PascalString
	NoOfProfileUris int32
	ProfileUris     []PascalString
}

var _ GetEndpointsRequest = (*_GetEndpointsRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetEndpointsRequest) GetIdentifier() string {
	return "428"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetEndpointsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetEndpointsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_GetEndpointsRequest) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_GetEndpointsRequest) GetNoOfLocaleIds() int32 {
	return m.NoOfLocaleIds
}

func (m *_GetEndpointsRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_GetEndpointsRequest) GetNoOfProfileUris() int32 {
	return m.NoOfProfileUris
}

func (m *_GetEndpointsRequest) GetProfileUris() []PascalString {
	return m.ProfileUris
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGetEndpointsRequest factory function for _GetEndpointsRequest
func NewGetEndpointsRequest(requestHeader ExtensionObjectDefinition, endpointUrl PascalString, noOfLocaleIds int32, localeIds []PascalString, noOfProfileUris int32, profileUris []PascalString) *_GetEndpointsRequest {
	_result := &_GetEndpointsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		EndpointUrl:                       endpointUrl,
		NoOfLocaleIds:                     noOfLocaleIds,
		LocaleIds:                         localeIds,
		NoOfProfileUris:                   noOfProfileUris,
		ProfileUris:                       profileUris,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetEndpointsRequest(structType any) GetEndpointsRequest {
	if casted, ok := structType.(GetEndpointsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetEndpointsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetEndpointsRequest) GetTypeName() string {
	return "GetEndpointsRequest"
}

func (m *_GetEndpointsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfProfileUris)
	lengthInBits += 32

	// Array field
	if len(m.ProfileUris) > 0 {
		for _curItem, element := range m.ProfileUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ProfileUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_GetEndpointsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetEndpointsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__getEndpointsRequest GetEndpointsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetEndpointsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetEndpointsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	endpointUrl, err := ReadSimpleField[PascalString](ctx, "endpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpointUrl' field"))
	}
	m.EndpointUrl = endpointUrl

	noOfLocaleIds, err := ReadSimpleField(ctx, "noOfLocaleIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLocaleIds' field"))
	}
	m.NoOfLocaleIds = noOfLocaleIds

	localeIds, err := ReadCountArrayField[PascalString](ctx, "localeIds", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfLocaleIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeIds' field"))
	}
	m.LocaleIds = localeIds

	noOfProfileUris, err := ReadSimpleField(ctx, "noOfProfileUris", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfProfileUris' field"))
	}
	m.NoOfProfileUris = noOfProfileUris

	profileUris, err := ReadCountArrayField[PascalString](ctx, "profileUris", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfProfileUris))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'profileUris' field"))
	}
	m.ProfileUris = profileUris

	if closeErr := readBuffer.CloseContext("GetEndpointsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetEndpointsRequest")
	}

	return m, nil
}

func (m *_GetEndpointsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetEndpointsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetEndpointsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetEndpointsRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpointUrl", m.GetEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpointUrl' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLocaleIds", m.GetNoOfLocaleIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfProfileUris", m.GetNoOfProfileUris(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfProfileUris' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "profileUris", m.GetProfileUris(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'profileUris' field")
		}

		if popErr := writeBuffer.PopContext("GetEndpointsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetEndpointsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetEndpointsRequest) isGetEndpointsRequest() bool {
	return true
}

func (m *_GetEndpointsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
