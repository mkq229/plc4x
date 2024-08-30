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

// IdentifyReplyCommandType is the corresponding interface of IdentifyReplyCommandType
type IdentifyReplyCommandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetUnitType returns UnitType (property field)
	GetUnitType() string
}

// IdentifyReplyCommandTypeExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandType.
// This is useful for switch cases.
type IdentifyReplyCommandTypeExactly interface {
	IdentifyReplyCommandType
	isIdentifyReplyCommandType() bool
}

// _IdentifyReplyCommandType is the data-structure of this message
type _IdentifyReplyCommandType struct {
	*_IdentifyReplyCommand
	UnitType string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandType) GetAttribute() Attribute {
	return Attribute_Type
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandType) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandType) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandType) GetUnitType() string {
	return m.UnitType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandType factory function for _IdentifyReplyCommandType
func NewIdentifyReplyCommandType(unitType string, numBytes uint8) *_IdentifyReplyCommandType {
	_result := &_IdentifyReplyCommandType{
		UnitType:              unitType,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandType(structType any) IdentifyReplyCommandType {
	if casted, ok := structType.(IdentifyReplyCommandType); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandType); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandType) GetTypeName() string {
	return "IdentifyReplyCommandType"
}

func (m *_IdentifyReplyCommandType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unitType)
	lengthInBits += 64

	return lengthInBits
}

func (m *_IdentifyReplyCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandTypeParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandType, error) {
	return IdentifyReplyCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unitType)
	_unitType, _unitTypeErr := readBuffer.ReadString("unitType", uint32(64), utils.WithEncoding("UTF-8"))
	if _unitTypeErr != nil {
		return nil, errors.Wrap(_unitTypeErr, "Error parsing 'unitType' field of IdentifyReplyCommandType")
	}
	unitType := _unitType

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandType")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandType{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		UnitType: unitType,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandType")
		}

		// Simple Field (unitType)
		unitType := string(m.GetUnitType())
		_unitTypeErr := writeBuffer.WriteString("unitType", uint32(64), (unitType), utils.WithEncoding("UTF-8)"))
		if _unitTypeErr != nil {
			return errors.Wrap(_unitTypeErr, "Error serializing 'unitType' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandType) isIdentifyReplyCommandType() bool {
	return true
}

func (m *_IdentifyReplyCommandType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
