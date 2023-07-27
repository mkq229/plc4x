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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// PublishedDataSetSourceDataType is the corresponding interface of PublishedDataSetSourceDataType
type PublishedDataSetSourceDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// PublishedDataSetSourceDataTypeExactly can be used when we want exactly this type and not a type which fulfills PublishedDataSetSourceDataType.
// This is useful for switch cases.
type PublishedDataSetSourceDataTypeExactly interface {
	PublishedDataSetSourceDataType
	isPublishedDataSetSourceDataType() bool
}

// _PublishedDataSetSourceDataType is the data-structure of this message
type _PublishedDataSetSourceDataType struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishedDataSetSourceDataType) GetIdentifier() string {
	return "15582"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishedDataSetSourceDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PublishedDataSetSourceDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewPublishedDataSetSourceDataType factory function for _PublishedDataSetSourceDataType
func NewPublishedDataSetSourceDataType() *_PublishedDataSetSourceDataType {
	_result := &_PublishedDataSetSourceDataType{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPublishedDataSetSourceDataType(structType any) PublishedDataSetSourceDataType {
	if casted, ok := structType.(PublishedDataSetSourceDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PublishedDataSetSourceDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PublishedDataSetSourceDataType) GetTypeName() string {
	return "PublishedDataSetSourceDataType"
}

func (m *_PublishedDataSetSourceDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_PublishedDataSetSourceDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PublishedDataSetSourceDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (PublishedDataSetSourceDataType, error) {
	return PublishedDataSetSourceDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PublishedDataSetSourceDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PublishedDataSetSourceDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PublishedDataSetSourceDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishedDataSetSourceDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("PublishedDataSetSourceDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishedDataSetSourceDataType")
	}

	// Create a partially initialized instance
	_child := &_PublishedDataSetSourceDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PublishedDataSetSourceDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishedDataSetSourceDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishedDataSetSourceDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishedDataSetSourceDataType")
		}

		if popErr := writeBuffer.PopContext("PublishedDataSetSourceDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishedDataSetSourceDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishedDataSetSourceDataType) isPublishedDataSetSourceDataType() bool {
	return true
}

func (m *_PublishedDataSetSourceDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}