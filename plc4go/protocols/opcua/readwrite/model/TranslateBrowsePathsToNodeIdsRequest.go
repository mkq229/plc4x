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

// TranslateBrowsePathsToNodeIdsRequest is the corresponding interface of TranslateBrowsePathsToNodeIdsRequest
type TranslateBrowsePathsToNodeIdsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfBrowsePaths returns NoOfBrowsePaths (property field)
	GetNoOfBrowsePaths() int32
	// GetBrowsePaths returns BrowsePaths (property field)
	GetBrowsePaths() []ExtensionObjectDefinition
}

// TranslateBrowsePathsToNodeIdsRequestExactly can be used when we want exactly this type and not a type which fulfills TranslateBrowsePathsToNodeIdsRequest.
// This is useful for switch cases.
type TranslateBrowsePathsToNodeIdsRequestExactly interface {
	TranslateBrowsePathsToNodeIdsRequest
	isTranslateBrowsePathsToNodeIdsRequest() bool
}

// _TranslateBrowsePathsToNodeIdsRequest is the data-structure of this message
type _TranslateBrowsePathsToNodeIdsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader   ExtensionObjectDefinition
	NoOfBrowsePaths int32
	BrowsePaths     []ExtensionObjectDefinition
}

var _ TranslateBrowsePathsToNodeIdsRequest = (*_TranslateBrowsePathsToNodeIdsRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetIdentifier() string {
	return "554"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetNoOfBrowsePaths() int32 {
	return m.NoOfBrowsePaths
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetBrowsePaths() []ExtensionObjectDefinition {
	return m.BrowsePaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTranslateBrowsePathsToNodeIdsRequest factory function for _TranslateBrowsePathsToNodeIdsRequest
func NewTranslateBrowsePathsToNodeIdsRequest(requestHeader ExtensionObjectDefinition, noOfBrowsePaths int32, browsePaths []ExtensionObjectDefinition) *_TranslateBrowsePathsToNodeIdsRequest {
	_result := &_TranslateBrowsePathsToNodeIdsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfBrowsePaths:                   noOfBrowsePaths,
		BrowsePaths:                       browsePaths,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastTranslateBrowsePathsToNodeIdsRequest(structType any) TranslateBrowsePathsToNodeIdsRequest {
	if casted, ok := structType.(TranslateBrowsePathsToNodeIdsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TranslateBrowsePathsToNodeIdsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetTypeName() string {
	return "TranslateBrowsePathsToNodeIdsRequest"
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfBrowsePaths)
	lengthInBits += 32

	// Array field
	if len(m.BrowsePaths) > 0 {
		for _curItem, element := range m.BrowsePaths {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.BrowsePaths), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__translateBrowsePathsToNodeIdsRequest TranslateBrowsePathsToNodeIdsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TranslateBrowsePathsToNodeIdsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TranslateBrowsePathsToNodeIdsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfBrowsePaths, err := ReadSimpleField(ctx, "noOfBrowsePaths", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfBrowsePaths' field"))
	}
	m.NoOfBrowsePaths = noOfBrowsePaths

	browsePaths, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "browsePaths", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("545")), readBuffer), uint64(noOfBrowsePaths))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'browsePaths' field"))
	}
	m.BrowsePaths = browsePaths

	if closeErr := readBuffer.CloseContext("TranslateBrowsePathsToNodeIdsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TranslateBrowsePathsToNodeIdsRequest")
	}

	return m, nil
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TranslateBrowsePathsToNodeIdsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TranslateBrowsePathsToNodeIdsRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfBrowsePaths", m.GetNoOfBrowsePaths(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfBrowsePaths' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "browsePaths", m.GetBrowsePaths(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'browsePaths' field")
		}

		if popErr := writeBuffer.PopContext("TranslateBrowsePathsToNodeIdsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TranslateBrowsePathsToNodeIdsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) isTranslateBrowsePathsToNodeIdsRequest() bool {
	return true
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
