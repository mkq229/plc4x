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

// RelativePath is the corresponding interface of RelativePath
type RelativePath interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfElements returns NoOfElements (property field)
	GetNoOfElements() int32
	// GetElements returns Elements (property field)
	GetElements() []ExtensionObjectDefinition
}

// RelativePathExactly can be used when we want exactly this type and not a type which fulfills RelativePath.
// This is useful for switch cases.
type RelativePathExactly interface {
	RelativePath
	isRelativePath() bool
}

// _RelativePath is the data-structure of this message
type _RelativePath struct {
	ExtensionObjectDefinitionContract
	NoOfElements int32
	Elements     []ExtensionObjectDefinition
}

var _ RelativePath = (*_RelativePath)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RelativePath) GetIdentifier() string {
	return "542"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RelativePath) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RelativePath) GetNoOfElements() int32 {
	return m.NoOfElements
}

func (m *_RelativePath) GetElements() []ExtensionObjectDefinition {
	return m.Elements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRelativePath factory function for _RelativePath
func NewRelativePath(noOfElements int32, elements []ExtensionObjectDefinition) *_RelativePath {
	_result := &_RelativePath{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfElements:                      noOfElements,
		Elements:                          elements,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastRelativePath(structType any) RelativePath {
	if casted, ok := structType.(RelativePath); ok {
		return casted
	}
	if casted, ok := structType.(*RelativePath); ok {
		return *casted
	}
	return nil
}

func (m *_RelativePath) GetTypeName() string {
	return "RelativePath"
}

func (m *_RelativePath) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfElements)
	lengthInBits += 32

	// Array field
	if len(m.Elements) > 0 {
		for _curItem, element := range m.Elements {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Elements), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RelativePath) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RelativePath) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__relativePath RelativePath, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RelativePath"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RelativePath")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfElements, err := ReadSimpleField(ctx, "noOfElements", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElements' field"))
	}
	m.NoOfElements = noOfElements

	elements, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "elements", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("539")), readBuffer), uint64(noOfElements))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elements' field"))
	}
	m.Elements = elements

	if closeErr := readBuffer.CloseContext("RelativePath"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RelativePath")
	}

	return m, nil
}

func (m *_RelativePath) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RelativePath) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RelativePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RelativePath")
		}

		if err := WriteSimpleField[int32](ctx, "noOfElements", m.GetNoOfElements(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "elements", m.GetElements(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'elements' field")
		}

		if popErr := writeBuffer.PopContext("RelativePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RelativePath")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RelativePath) isRelativePath() bool {
	return true
}

func (m *_RelativePath) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
