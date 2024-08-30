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

// LightingDataRampToLevel is the corresponding interface of LightingDataRampToLevel
type LightingDataRampToLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// GetLevel returns Level (property field)
	GetLevel() byte
}

// LightingDataRampToLevelExactly can be used when we want exactly this type and not a type which fulfills LightingDataRampToLevel.
// This is useful for switch cases.
type LightingDataRampToLevelExactly interface {
	LightingDataRampToLevel
	isLightingDataRampToLevel() bool
}

// _LightingDataRampToLevel is the data-structure of this message
type _LightingDataRampToLevel struct {
	*_LightingData
	Group byte
	Level byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataRampToLevel) InitializeParent(parent LightingData, commandTypeContainer LightingCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_LightingDataRampToLevel) GetParent() LightingData {
	return m._LightingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataRampToLevel) GetGroup() byte {
	return m.Group
}

func (m *_LightingDataRampToLevel) GetLevel() byte {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataRampToLevel factory function for _LightingDataRampToLevel
func NewLightingDataRampToLevel(group byte, level byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataRampToLevel {
	_result := &_LightingDataRampToLevel{
		Group:         group,
		Level:         level,
		_LightingData: NewLightingData(commandTypeContainer),
	}
	_result._LightingData._LightingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataRampToLevel(structType any) LightingDataRampToLevel {
	if casted, ok := structType.(LightingDataRampToLevel); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataRampToLevel); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataRampToLevel) GetTypeName() string {
	return "LightingDataRampToLevel"
}

func (m *_LightingDataRampToLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataRampToLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingDataRampToLevelParse(ctx context.Context, theBytes []byte) (LightingDataRampToLevel, error) {
	return LightingDataRampToLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LightingDataRampToLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LightingDataRampToLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("LightingDataRampToLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataRampToLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := /*TODO: migrate me*/ readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field of LightingDataRampToLevel")
	}
	group := _group

	// Simple Field (level)
	_level, _levelErr := /*TODO: migrate me*/ readBuffer.ReadByte("level")
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of LightingDataRampToLevel")
	}
	level := _level

	if closeErr := readBuffer.CloseContext("LightingDataRampToLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataRampToLevel")
	}

	// Create a partially initialized instance
	_child := &_LightingDataRampToLevel{
		_LightingData: &_LightingData{},
		Group:         group,
		Level:         level,
	}
	_child._LightingData._LightingDataChildRequirements = _child
	return _child, nil
}

func (m *_LightingDataRampToLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataRampToLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataRampToLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataRampToLevel")
		}

		// Simple Field (group)
		group := byte(m.GetGroup())
		_groupErr := /*TODO: migrate me*/ writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		// Simple Field (level)
		level := byte(m.GetLevel())
		_levelErr := /*TODO: migrate me*/ writeBuffer.WriteByte("level", (level))
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataRampToLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataRampToLevel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataRampToLevel) isLightingDataRampToLevel() bool {
	return true
}

func (m *_LightingDataRampToLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
