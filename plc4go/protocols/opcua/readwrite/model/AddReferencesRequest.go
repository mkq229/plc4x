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

// AddReferencesRequest is the corresponding interface of AddReferencesRequest
type AddReferencesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfReferencesToAdd returns NoOfReferencesToAdd (property field)
	GetNoOfReferencesToAdd() int32
	// GetReferencesToAdd returns ReferencesToAdd (property field)
	GetReferencesToAdd() []ExtensionObjectDefinition
}

// AddReferencesRequestExactly can be used when we want exactly this type and not a type which fulfills AddReferencesRequest.
// This is useful for switch cases.
type AddReferencesRequestExactly interface {
	AddReferencesRequest
	isAddReferencesRequest() bool
}

// _AddReferencesRequest is the data-structure of this message
type _AddReferencesRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader       ExtensionObjectDefinition
	NoOfReferencesToAdd int32
	ReferencesToAdd     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddReferencesRequest) GetIdentifier() string {
	return "494"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddReferencesRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_AddReferencesRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddReferencesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_AddReferencesRequest) GetNoOfReferencesToAdd() int32 {
	return m.NoOfReferencesToAdd
}

func (m *_AddReferencesRequest) GetReferencesToAdd() []ExtensionObjectDefinition {
	return m.ReferencesToAdd
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAddReferencesRequest factory function for _AddReferencesRequest
func NewAddReferencesRequest(requestHeader ExtensionObjectDefinition, noOfReferencesToAdd int32, referencesToAdd []ExtensionObjectDefinition) *_AddReferencesRequest {
	_result := &_AddReferencesRequest{
		RequestHeader:              requestHeader,
		NoOfReferencesToAdd:        noOfReferencesToAdd,
		ReferencesToAdd:            referencesToAdd,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAddReferencesRequest(structType any) AddReferencesRequest {
	if casted, ok := structType.(AddReferencesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AddReferencesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AddReferencesRequest) GetTypeName() string {
	return "AddReferencesRequest"
}

func (m *_AddReferencesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfReferencesToAdd)
	lengthInBits += 32

	// Array field
	if len(m.ReferencesToAdd) > 0 {
		for _curItem, element := range m.ReferencesToAdd {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencesToAdd), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AddReferencesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AddReferencesRequestParse(ctx context.Context, theBytes []byte, identifier string) (AddReferencesRequest, error) {
	return AddReferencesRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AddReferencesRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (AddReferencesRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AddReferencesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddReferencesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of AddReferencesRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfReferencesToAdd)
	_noOfReferencesToAdd, _noOfReferencesToAddErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfReferencesToAdd", 32)
	if _noOfReferencesToAddErr != nil {
		return nil, errors.Wrap(_noOfReferencesToAddErr, "Error parsing 'noOfReferencesToAdd' field of AddReferencesRequest")
	}
	noOfReferencesToAdd := _noOfReferencesToAdd

	referencesToAdd, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "referencesToAdd", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("381"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfReferencesToAdd))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencesToAdd' field"))
	}

	if closeErr := readBuffer.CloseContext("AddReferencesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddReferencesRequest")
	}

	// Create a partially initialized instance
	_child := &_AddReferencesRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfReferencesToAdd:        noOfReferencesToAdd,
		ReferencesToAdd:            referencesToAdd,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_AddReferencesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddReferencesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddReferencesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddReferencesRequest")
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

		// Simple Field (noOfReferencesToAdd)
		noOfReferencesToAdd := int32(m.GetNoOfReferencesToAdd())
		_noOfReferencesToAddErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfReferencesToAdd", 32, int32((noOfReferencesToAdd)))
		if _noOfReferencesToAddErr != nil {
			return errors.Wrap(_noOfReferencesToAddErr, "Error serializing 'noOfReferencesToAdd' field")
		}

		// Array Field (referencesToAdd)
		if pushErr := writeBuffer.PushContext("referencesToAdd", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referencesToAdd")
		}
		for _curItem, _element := range m.GetReferencesToAdd() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReferencesToAdd()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'referencesToAdd' field")
			}
		}
		if popErr := writeBuffer.PopContext("referencesToAdd", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referencesToAdd")
		}

		if popErr := writeBuffer.PopContext("AddReferencesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddReferencesRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddReferencesRequest) isAddReferencesRequest() bool {
	return true
}

func (m *_AddReferencesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
