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

// ClassID is the corresponding interface of ClassID
type ClassID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LogicalSegmentType
	// GetFormat returns Format (property field)
	GetFormat() uint8
	// GetSegmentClass returns SegmentClass (property field)
	GetSegmentClass() uint8
}

// ClassIDExactly can be used when we want exactly this type and not a type which fulfills ClassID.
// This is useful for switch cases.
type ClassIDExactly interface {
	ClassID
	isClassID() bool
}

// _ClassID is the data-structure of this message
type _ClassID struct {
	*_LogicalSegmentType
	Format       uint8
	SegmentClass uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ClassID) GetLogicalSegmentType() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClassID) InitializeParent(parent LogicalSegmentType) {}

func (m *_ClassID) GetParent() LogicalSegmentType {
	return m._LogicalSegmentType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClassID) GetFormat() uint8 {
	return m.Format
}

func (m *_ClassID) GetSegmentClass() uint8 {
	return m.SegmentClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewClassID factory function for _ClassID
func NewClassID(format uint8, segmentClass uint8) *_ClassID {
	_result := &_ClassID{
		Format:              format,
		SegmentClass:        segmentClass,
		_LogicalSegmentType: NewLogicalSegmentType(),
	}
	_result._LogicalSegmentType._LogicalSegmentTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastClassID(structType any) ClassID {
	if casted, ok := structType.(ClassID); ok {
		return casted
	}
	if casted, ok := structType.(*ClassID); ok {
		return *casted
	}
	return nil
}

func (m *_ClassID) GetTypeName() string {
	return "ClassID"
}

func (m *_ClassID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (format)
	lengthInBits += 2

	// Simple field (segmentClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClassID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClassIDParse(ctx context.Context, theBytes []byte) (ClassID, error) {
	return ClassIDParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClassIDParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClassID, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ClassID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClassID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (format)
	_format, _formatErr := /*TODO: migrate me*/ readBuffer.ReadUint8("format", 2)
	if _formatErr != nil {
		return nil, errors.Wrap(_formatErr, "Error parsing 'format' field of ClassID")
	}
	format := _format

	// Simple Field (segmentClass)
	_segmentClass, _segmentClassErr := /*TODO: migrate me*/ readBuffer.ReadUint8("segmentClass", 8)
	if _segmentClassErr != nil {
		return nil, errors.Wrap(_segmentClassErr, "Error parsing 'segmentClass' field of ClassID")
	}
	segmentClass := _segmentClass

	if closeErr := readBuffer.CloseContext("ClassID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClassID")
	}

	// Create a partially initialized instance
	_child := &_ClassID{
		_LogicalSegmentType: &_LogicalSegmentType{},
		Format:              format,
		SegmentClass:        segmentClass,
	}
	_child._LogicalSegmentType._LogicalSegmentTypeChildRequirements = _child
	return _child, nil
}

func (m *_ClassID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClassID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClassID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClassID")
		}

		// Simple Field (format)
		format := uint8(m.GetFormat())
		_formatErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("format", 2, uint8((format)))
		if _formatErr != nil {
			return errors.Wrap(_formatErr, "Error serializing 'format' field")
		}

		// Simple Field (segmentClass)
		segmentClass := uint8(m.GetSegmentClass())
		_segmentClassErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("segmentClass", 8, uint8((segmentClass)))
		if _segmentClassErr != nil {
			return errors.Wrap(_segmentClassErr, "Error serializing 'segmentClass' field")
		}

		if popErr := writeBuffer.PopContext("ClassID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClassID")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClassID) isClassID() bool {
	return true
}

func (m *_ClassID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
