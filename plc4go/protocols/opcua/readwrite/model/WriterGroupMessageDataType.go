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

// WriterGroupMessageDataType is the corresponding interface of WriterGroupMessageDataType
type WriterGroupMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// WriterGroupMessageDataTypeExactly can be used when we want exactly this type and not a type which fulfills WriterGroupMessageDataType.
// This is useful for switch cases.
type WriterGroupMessageDataTypeExactly interface {
	WriterGroupMessageDataType
	isWriterGroupMessageDataType() bool
}

// _WriterGroupMessageDataType is the data-structure of this message
type _WriterGroupMessageDataType struct {
	ExtensionObjectDefinitionContract
}

var _ WriterGroupMessageDataType = (*_WriterGroupMessageDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriterGroupMessageDataType) GetIdentifier() string {
	return "15618"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriterGroupMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewWriterGroupMessageDataType factory function for _WriterGroupMessageDataType
func NewWriterGroupMessageDataType() *_WriterGroupMessageDataType {
	_result := &_WriterGroupMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastWriterGroupMessageDataType(structType any) WriterGroupMessageDataType {
	if casted, ok := structType.(WriterGroupMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*WriterGroupMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_WriterGroupMessageDataType) GetTypeName() string {
	return "WriterGroupMessageDataType"
}

func (m *_WriterGroupMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_WriterGroupMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_WriterGroupMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__writerGroupMessageDataType WriterGroupMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WriterGroupMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriterGroupMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("WriterGroupMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriterGroupMessageDataType")
	}

	return m, nil
}

func (m *_WriterGroupMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriterGroupMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriterGroupMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriterGroupMessageDataType")
		}

		if popErr := writeBuffer.PopContext("WriterGroupMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriterGroupMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriterGroupMessageDataType) isWriterGroupMessageDataType() bool {
	return true
}

func (m *_WriterGroupMessageDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
