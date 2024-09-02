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

// EUInformation is the corresponding interface of EUInformation
type EUInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNamespaceUri returns NamespaceUri (property field)
	GetNamespaceUri() PascalString
	// GetUnitId returns UnitId (property field)
	GetUnitId() int32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
}

// EUInformationExactly can be used when we want exactly this type and not a type which fulfills EUInformation.
// This is useful for switch cases.
type EUInformationExactly interface {
	EUInformation
	isEUInformation() bool
}

// _EUInformation is the data-structure of this message
type _EUInformation struct {
	ExtensionObjectDefinitionContract
	NamespaceUri PascalString
	UnitId       int32
	DisplayName  LocalizedText
	Description  LocalizedText
}

var _ EUInformation = (*_EUInformation)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EUInformation) GetIdentifier() string {
	return "889"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EUInformation) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EUInformation) GetNamespaceUri() PascalString {
	return m.NamespaceUri
}

func (m *_EUInformation) GetUnitId() int32 {
	return m.UnitId
}

func (m *_EUInformation) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_EUInformation) GetDescription() LocalizedText {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEUInformation factory function for _EUInformation
func NewEUInformation(namespaceUri PascalString, unitId int32, displayName LocalizedText, description LocalizedText) *_EUInformation {
	_result := &_EUInformation{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NamespaceUri:                      namespaceUri,
		UnitId:                            unitId,
		DisplayName:                       displayName,
		Description:                       description,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastEUInformation(structType any) EUInformation {
	if casted, ok := structType.(EUInformation); ok {
		return casted
	}
	if casted, ok := structType.(*EUInformation); ok {
		return *casted
	}
	return nil
}

func (m *_EUInformation) GetTypeName() string {
	return "EUInformation"
}

func (m *_EUInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (namespaceUri)
	lengthInBits += m.NamespaceUri.GetLengthInBits(ctx)

	// Simple field (unitId)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EUInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EUInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__eUInformation EUInformation, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EUInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EUInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceUri, err := ReadSimpleField[PascalString](ctx, "namespaceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceUri' field"))
	}
	m.NamespaceUri = namespaceUri

	unitId, err := ReadSimpleField(ctx, "unitId", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitId' field"))
	}
	m.UnitId = unitId

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	if closeErr := readBuffer.CloseContext("EUInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EUInformation")
	}

	return m, nil
}

func (m *_EUInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EUInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EUInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EUInformation")
		}

		if err := WriteSimpleField[PascalString](ctx, "namespaceUri", m.GetNamespaceUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceUri' field")
		}

		if err := WriteSimpleField[int32](ctx, "unitId", m.GetUnitId(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitId' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("EUInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EUInformation")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EUInformation) isEUInformation() bool {
	return true
}

func (m *_EUInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
