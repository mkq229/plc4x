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

// KnxGroupAddress2Level is the corresponding interface of KnxGroupAddress2Level
type KnxGroupAddress2Level interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxGroupAddress
	// GetMainGroup returns MainGroup (property field)
	GetMainGroup() uint8
	// GetSubGroup returns SubGroup (property field)
	GetSubGroup() uint16
}

// KnxGroupAddress2LevelExactly can be used when we want exactly this type and not a type which fulfills KnxGroupAddress2Level.
// This is useful for switch cases.
type KnxGroupAddress2LevelExactly interface {
	KnxGroupAddress2Level
	isKnxGroupAddress2Level() bool
}

// _KnxGroupAddress2Level is the data-structure of this message
type _KnxGroupAddress2Level struct {
	KnxGroupAddressContract
	MainGroup uint8
	SubGroup  uint16
}

var _ KnxGroupAddress2Level = (*_KnxGroupAddress2Level)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxGroupAddress2Level) GetNumLevels() uint8 {
	return uint8(2)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxGroupAddress2Level) GetParent() KnxGroupAddressContract {
	return m.KnxGroupAddressContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxGroupAddress2Level) GetMainGroup() uint8 {
	return m.MainGroup
}

func (m *_KnxGroupAddress2Level) GetSubGroup() uint16 {
	return m.SubGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxGroupAddress2Level factory function for _KnxGroupAddress2Level
func NewKnxGroupAddress2Level(mainGroup uint8, subGroup uint16) *_KnxGroupAddress2Level {
	_result := &_KnxGroupAddress2Level{
		KnxGroupAddressContract: NewKnxGroupAddress(),
		MainGroup:               mainGroup,
		SubGroup:                subGroup,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxGroupAddress2Level(structType any) KnxGroupAddress2Level {
	if casted, ok := structType.(KnxGroupAddress2Level); ok {
		return casted
	}
	if casted, ok := structType.(*KnxGroupAddress2Level); ok {
		return *casted
	}
	return nil
}

func (m *_KnxGroupAddress2Level) GetTypeName() string {
	return "KnxGroupAddress2Level"
}

func (m *_KnxGroupAddress2Level) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxGroupAddressContract.(*_KnxGroupAddress).getLengthInBits(ctx))

	// Simple field (mainGroup)
	lengthInBits += 5

	// Simple field (subGroup)
	lengthInBits += 11

	return lengthInBits
}

func (m *_KnxGroupAddress2Level) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxGroupAddress2Level) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxGroupAddress, numLevels uint8) (__knxGroupAddress2Level KnxGroupAddress2Level, err error) {
	m.KnxGroupAddressContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxGroupAddress2Level"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxGroupAddress2Level")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	mainGroup, err := ReadSimpleField(ctx, "mainGroup", ReadUnsignedByte(readBuffer, uint8(5)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mainGroup' field"))
	}
	m.MainGroup = mainGroup

	subGroup, err := ReadSimpleField(ctx, "subGroup", ReadUnsignedShort(readBuffer, uint8(11)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subGroup' field"))
	}
	m.SubGroup = subGroup

	if closeErr := readBuffer.CloseContext("KnxGroupAddress2Level"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxGroupAddress2Level")
	}

	return m, nil
}

func (m *_KnxGroupAddress2Level) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxGroupAddress2Level) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddress2Level"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxGroupAddress2Level")
		}

		if err := WriteSimpleField[uint8](ctx, "mainGroup", m.GetMainGroup(), WriteUnsignedByte(writeBuffer, 5)); err != nil {
			return errors.Wrap(err, "Error serializing 'mainGroup' field")
		}

		if err := WriteSimpleField[uint16](ctx, "subGroup", m.GetSubGroup(), WriteUnsignedShort(writeBuffer, 11)); err != nil {
			return errors.Wrap(err, "Error serializing 'subGroup' field")
		}

		if popErr := writeBuffer.PopContext("KnxGroupAddress2Level"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxGroupAddress2Level")
		}
		return nil
	}
	return m.KnxGroupAddressContract.(*_KnxGroupAddress).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxGroupAddress2Level) isKnxGroupAddress2Level() bool {
	return true
}

func (m *_KnxGroupAddress2Level) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
