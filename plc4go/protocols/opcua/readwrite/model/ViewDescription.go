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

// ViewDescription is the corresponding interface of ViewDescription
type ViewDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetViewId returns ViewId (property field)
	GetViewId() NodeId
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() int64
	// GetViewVersion returns ViewVersion (property field)
	GetViewVersion() uint32
}

// ViewDescriptionExactly can be used when we want exactly this type and not a type which fulfills ViewDescription.
// This is useful for switch cases.
type ViewDescriptionExactly interface {
	ViewDescription
	isViewDescription() bool
}

// _ViewDescription is the data-structure of this message
type _ViewDescription struct {
	ExtensionObjectDefinitionContract
	ViewId      NodeId
	Timestamp   int64
	ViewVersion uint32
}

var _ ViewDescription = (*_ViewDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ViewDescription) GetIdentifier() string {
	return "513"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ViewDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ViewDescription) GetViewId() NodeId {
	return m.ViewId
}

func (m *_ViewDescription) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *_ViewDescription) GetViewVersion() uint32 {
	return m.ViewVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewViewDescription factory function for _ViewDescription
func NewViewDescription(viewId NodeId, timestamp int64, viewVersion uint32) *_ViewDescription {
	_result := &_ViewDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ViewId:                            viewId,
		Timestamp:                         timestamp,
		ViewVersion:                       viewVersion,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastViewDescription(structType any) ViewDescription {
	if casted, ok := structType.(ViewDescription); ok {
		return casted
	}
	if casted, ok := structType.(*ViewDescription); ok {
		return *casted
	}
	return nil
}

func (m *_ViewDescription) GetTypeName() string {
	return "ViewDescription"
}

func (m *_ViewDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (viewId)
	lengthInBits += m.ViewId.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (viewVersion)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ViewDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ViewDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__viewDescription ViewDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ViewDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ViewDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	viewId, err := ReadSimpleField[NodeId](ctx, "viewId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'viewId' field"))
	}
	m.ViewId = viewId

	timestamp, err := ReadSimpleField(ctx, "timestamp", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	viewVersion, err := ReadSimpleField(ctx, "viewVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'viewVersion' field"))
	}
	m.ViewVersion = viewVersion

	if closeErr := readBuffer.CloseContext("ViewDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ViewDescription")
	}

	return m, nil
}

func (m *_ViewDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ViewDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ViewDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ViewDescription")
		}

		if err := WriteSimpleField[NodeId](ctx, "viewId", m.GetViewId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'viewId' field")
		}

		if err := WriteSimpleField[int64](ctx, "timestamp", m.GetTimestamp(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[uint32](ctx, "viewVersion", m.GetViewVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'viewVersion' field")
		}

		if popErr := writeBuffer.PopContext("ViewDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ViewDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ViewDescription) isViewDescription() bool {
	return true
}

func (m *_ViewDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
