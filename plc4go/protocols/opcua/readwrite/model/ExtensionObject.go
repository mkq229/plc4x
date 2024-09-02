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

// ExtensionObject is the corresponding interface of ExtensionObject
type ExtensionObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTypeId returns TypeId (property field)
	GetTypeId() ExpandedNodeId
	// GetEncodingMask returns EncodingMask (property field)
	GetEncodingMask() ExtensionObjectEncodingMask
	// GetBody returns Body (property field)
	GetBody() ExtensionObjectDefinition
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// ExtensionObjectExactly can be used when we want exactly this type and not a type which fulfills ExtensionObject.
// This is useful for switch cases.
type ExtensionObjectExactly interface {
	ExtensionObject
	isExtensionObject() bool
}

// _ExtensionObject is the data-structure of this message
type _ExtensionObject struct {
	TypeId       ExpandedNodeId
	EncodingMask ExtensionObjectEncodingMask
	Body         ExtensionObjectDefinition

	// Arguments.
	IncludeEncodingMask bool
}

var _ ExtensionObject = (*_ExtensionObject)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensionObject) GetTypeId() ExpandedNodeId {
	return m.TypeId
}

func (m *_ExtensionObject) GetEncodingMask() ExtensionObjectEncodingMask {
	return m.EncodingMask
}

func (m *_ExtensionObject) GetBody() ExtensionObjectDefinition {
	return m.Body
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ExtensionObject) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	encodingMask := m.EncodingMask
	_ = encodingMask
	return fmt.Sprintf("%v", m.GetTypeId().GetIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExtensionObject factory function for _ExtensionObject
func NewExtensionObject(typeId ExpandedNodeId, encodingMask ExtensionObjectEncodingMask, body ExtensionObjectDefinition, includeEncodingMask bool) *_ExtensionObject {
	return &_ExtensionObject{TypeId: typeId, EncodingMask: encodingMask, Body: body, IncludeEncodingMask: includeEncodingMask}
}

// Deprecated: use the interface for direct cast
func CastExtensionObject(structType any) ExtensionObject {
	if casted, ok := structType.(ExtensionObject); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensionObject); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensionObject) GetTypeName() string {
	return "ExtensionObject"
}

func (m *_ExtensionObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeId)
	lengthInBits += m.TypeId.GetLengthInBits(ctx)

	// Optional Field (encodingMask)
	if m.EncodingMask != nil {
		lengthInBits += m.EncodingMask.GetLengthInBits(ctx)
	}

	// A virtual field doesn't have any in- or output.

	// Simple field (body)
	lengthInBits += m.Body.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ExtensionObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ExtensionObjectParse(ctx context.Context, theBytes []byte, includeEncodingMask bool) (ExtensionObject, error) {
	return ExtensionObjectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), includeEncodingMask)
}

func ExtensionObjectParseWithBufferProducer(includeEncodingMask bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (ExtensionObject, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ExtensionObject, error) {
		return ExtensionObjectParseWithBuffer(ctx, readBuffer, includeEncodingMask)
	}
}

func ExtensionObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, includeEncodingMask bool) (ExtensionObject, error) {
	v, err := (&_ExtensionObject{}).parse(ctx, readBuffer, includeEncodingMask)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ExtensionObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, includeEncodingMask bool) (__extensionObject ExtensionObject, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtensionObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensionObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	typeId, err := ReadSimpleField[ExpandedNodeId](ctx, "typeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeId' field"))
	}
	m.TypeId = typeId

	var encodingMask ExtensionObjectEncodingMask
	_encodingMask, err := ReadOptionalField[ExtensionObjectEncodingMask](ctx, "encodingMask", ReadComplex[ExtensionObjectEncodingMask](ExtensionObjectEncodingMaskParseWithBuffer, readBuffer), includeEncodingMask)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encodingMask' field"))
	}
	if _encodingMask != nil {
		encodingMask = *_encodingMask
		m.EncodingMask = encodingMask
	}

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), typeId.GetIdentifier())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	body, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "body", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)(identifier)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'body' field"))
	}
	m.Body = body

	if closeErr := readBuffer.CloseContext("ExtensionObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensionObject")
	}

	return m, nil
}

func (m *_ExtensionObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExtensionObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ExtensionObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExtensionObject")
	}

	if err := WriteSimpleField[ExpandedNodeId](ctx, "typeId", m.GetTypeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'typeId' field")
	}

	if err := WriteOptionalField[ExtensionObjectEncodingMask](ctx, "encodingMask", GetRef(m.GetEncodingMask()), WriteComplex[ExtensionObjectEncodingMask](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'encodingMask' field")
	}
	// Virtual field
	identifier := m.GetIdentifier()
	_ = identifier
	if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
		return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
	}

	if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "body", m.GetBody(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'body' field")
	}

	if popErr := writeBuffer.PopContext("ExtensionObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExtensionObject")
	}
	return nil
}

////
// Arguments Getter

func (m *_ExtensionObject) GetIncludeEncodingMask() bool {
	return m.IncludeEncodingMask
}

//
////

func (m *_ExtensionObject) isExtensionObject() bool {
	return true
}

func (m *_ExtensionObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
