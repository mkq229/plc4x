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

// PriorityMappingEntryType is the corresponding interface of PriorityMappingEntryType
type PriorityMappingEntryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMappingUri returns MappingUri (property field)
	GetMappingUri() PascalString
	// GetPriorityLabel returns PriorityLabel (property field)
	GetPriorityLabel() PascalString
	// GetPriorityValue_PCP returns PriorityValue_PCP (property field)
	GetPriorityValue_PCP() uint8
	// GetPriorityValue_DSCP returns PriorityValue_DSCP (property field)
	GetPriorityValue_DSCP() uint32
}

// PriorityMappingEntryTypeExactly can be used when we want exactly this type and not a type which fulfills PriorityMappingEntryType.
// This is useful for switch cases.
type PriorityMappingEntryTypeExactly interface {
	PriorityMappingEntryType
	isPriorityMappingEntryType() bool
}

// _PriorityMappingEntryType is the data-structure of this message
type _PriorityMappingEntryType struct {
	ExtensionObjectDefinitionContract
	MappingUri         PascalString
	PriorityLabel      PascalString
	PriorityValue_PCP  uint8
	PriorityValue_DSCP uint32
}

var _ PriorityMappingEntryType = (*_PriorityMappingEntryType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PriorityMappingEntryType) GetIdentifier() string {
	return "25222"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PriorityMappingEntryType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PriorityMappingEntryType) GetMappingUri() PascalString {
	return m.MappingUri
}

func (m *_PriorityMappingEntryType) GetPriorityLabel() PascalString {
	return m.PriorityLabel
}

func (m *_PriorityMappingEntryType) GetPriorityValue_PCP() uint8 {
	return m.PriorityValue_PCP
}

func (m *_PriorityMappingEntryType) GetPriorityValue_DSCP() uint32 {
	return m.PriorityValue_DSCP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPriorityMappingEntryType factory function for _PriorityMappingEntryType
func NewPriorityMappingEntryType(mappingUri PascalString, priorityLabel PascalString, priorityValue_PCP uint8, priorityValue_DSCP uint32) *_PriorityMappingEntryType {
	_result := &_PriorityMappingEntryType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		MappingUri:                        mappingUri,
		PriorityLabel:                     priorityLabel,
		PriorityValue_PCP:                 priorityValue_PCP,
		PriorityValue_DSCP:                priorityValue_DSCP,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastPriorityMappingEntryType(structType any) PriorityMappingEntryType {
	if casted, ok := structType.(PriorityMappingEntryType); ok {
		return casted
	}
	if casted, ok := structType.(*PriorityMappingEntryType); ok {
		return *casted
	}
	return nil
}

func (m *_PriorityMappingEntryType) GetTypeName() string {
	return "PriorityMappingEntryType"
}

func (m *_PriorityMappingEntryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (mappingUri)
	lengthInBits += m.MappingUri.GetLengthInBits(ctx)

	// Simple field (priorityLabel)
	lengthInBits += m.PriorityLabel.GetLengthInBits(ctx)

	// Simple field (priorityValue_PCP)
	lengthInBits += 8

	// Simple field (priorityValue_DSCP)
	lengthInBits += 32

	return lengthInBits
}

func (m *_PriorityMappingEntryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PriorityMappingEntryType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__priorityMappingEntryType PriorityMappingEntryType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PriorityMappingEntryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PriorityMappingEntryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	mappingUri, err := ReadSimpleField[PascalString](ctx, "mappingUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mappingUri' field"))
	}
	m.MappingUri = mappingUri

	priorityLabel, err := ReadSimpleField[PascalString](ctx, "priorityLabel", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityLabel' field"))
	}
	m.PriorityLabel = priorityLabel

	priorityValue_PCP, err := ReadSimpleField(ctx, "priorityValue_PCP", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityValue_PCP' field"))
	}
	m.PriorityValue_PCP = priorityValue_PCP

	priorityValue_DSCP, err := ReadSimpleField(ctx, "priorityValue_DSCP", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityValue_DSCP' field"))
	}
	m.PriorityValue_DSCP = priorityValue_DSCP

	if closeErr := readBuffer.CloseContext("PriorityMappingEntryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PriorityMappingEntryType")
	}

	return m, nil
}

func (m *_PriorityMappingEntryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PriorityMappingEntryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PriorityMappingEntryType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PriorityMappingEntryType")
		}

		if err := WriteSimpleField[PascalString](ctx, "mappingUri", m.GetMappingUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'mappingUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "priorityLabel", m.GetPriorityLabel(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityLabel' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priorityValue_PCP", m.GetPriorityValue_PCP(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityValue_PCP' field")
		}

		if err := WriteSimpleField[uint32](ctx, "priorityValue_DSCP", m.GetPriorityValue_DSCP(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityValue_DSCP' field")
		}

		if popErr := writeBuffer.PopContext("PriorityMappingEntryType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PriorityMappingEntryType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PriorityMappingEntryType) isPriorityMappingEntryType() bool {
	return true
}

func (m *_PriorityMappingEntryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
