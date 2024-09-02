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

// OptionSet is the corresponding interface of OptionSet
type OptionSet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetValue returns Value (property field)
	GetValue() PascalByteString
	// GetValidBits returns ValidBits (property field)
	GetValidBits() PascalByteString
}

// OptionSetExactly can be used when we want exactly this type and not a type which fulfills OptionSet.
// This is useful for switch cases.
type OptionSetExactly interface {
	OptionSet
	isOptionSet() bool
}

// _OptionSet is the data-structure of this message
type _OptionSet struct {
	ExtensionObjectDefinitionContract
	Value     PascalByteString
	ValidBits PascalByteString
}

var _ OptionSet = (*_OptionSet)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OptionSet) GetIdentifier() string {
	return "12757"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OptionSet) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OptionSet) GetValue() PascalByteString {
	return m.Value
}

func (m *_OptionSet) GetValidBits() PascalByteString {
	return m.ValidBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOptionSet factory function for _OptionSet
func NewOptionSet(value PascalByteString, validBits PascalByteString) *_OptionSet {
	_result := &_OptionSet{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Value:                             value,
		ValidBits:                         validBits,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastOptionSet(structType any) OptionSet {
	if casted, ok := structType.(OptionSet); ok {
		return casted
	}
	if casted, ok := structType.(*OptionSet); ok {
		return *casted
	}
	return nil
}

func (m *_OptionSet) GetTypeName() string {
	return "OptionSet"
}

func (m *_OptionSet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Simple field (validBits)
	lengthInBits += m.ValidBits.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OptionSet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OptionSet) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__optionSet OptionSet, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OptionSet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OptionSet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField[PascalByteString](ctx, "value", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	validBits, err := ReadSimpleField[PascalByteString](ctx, "validBits", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'validBits' field"))
	}
	m.ValidBits = validBits

	if closeErr := readBuffer.CloseContext("OptionSet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OptionSet")
	}

	return m, nil
}

func (m *_OptionSet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OptionSet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OptionSet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OptionSet")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "value", m.GetValue(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "validBits", m.GetValidBits(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'validBits' field")
		}

		if popErr := writeBuffer.PopContext("OptionSet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OptionSet")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OptionSet) isOptionSet() bool {
	return true
}

func (m *_OptionSet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
