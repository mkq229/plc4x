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

// DataTypeSchemaHeader is the corresponding interface of DataTypeSchemaHeader
type DataTypeSchemaHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfNamespaces returns NoOfNamespaces (property field)
	GetNoOfNamespaces() int32
	// GetNamespaces returns Namespaces (property field)
	GetNamespaces() []PascalString
	// GetNoOfStructureDataTypes returns NoOfStructureDataTypes (property field)
	GetNoOfStructureDataTypes() int32
	// GetStructureDataTypes returns StructureDataTypes (property field)
	GetStructureDataTypes() []DataTypeDescription
	// GetNoOfEnumDataTypes returns NoOfEnumDataTypes (property field)
	GetNoOfEnumDataTypes() int32
	// GetEnumDataTypes returns EnumDataTypes (property field)
	GetEnumDataTypes() []DataTypeDescription
	// GetNoOfSimpleDataTypes returns NoOfSimpleDataTypes (property field)
	GetNoOfSimpleDataTypes() int32
	// GetSimpleDataTypes returns SimpleDataTypes (property field)
	GetSimpleDataTypes() []DataTypeDescription
}

// DataTypeSchemaHeaderExactly can be used when we want exactly this type and not a type which fulfills DataTypeSchemaHeader.
// This is useful for switch cases.
type DataTypeSchemaHeaderExactly interface {
	DataTypeSchemaHeader
	isDataTypeSchemaHeader() bool
}

// _DataTypeSchemaHeader is the data-structure of this message
type _DataTypeSchemaHeader struct {
	*_ExtensionObjectDefinition
	NoOfNamespaces         int32
	Namespaces             []PascalString
	NoOfStructureDataTypes int32
	StructureDataTypes     []DataTypeDescription
	NoOfEnumDataTypes      int32
	EnumDataTypes          []DataTypeDescription
	NoOfSimpleDataTypes    int32
	SimpleDataTypes        []DataTypeDescription
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataTypeSchemaHeader) GetIdentifier() string {
	return "15536"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataTypeSchemaHeader) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DataTypeSchemaHeader) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataTypeSchemaHeader) GetNoOfNamespaces() int32 {
	return m.NoOfNamespaces
}

func (m *_DataTypeSchemaHeader) GetNamespaces() []PascalString {
	return m.Namespaces
}

func (m *_DataTypeSchemaHeader) GetNoOfStructureDataTypes() int32 {
	return m.NoOfStructureDataTypes
}

func (m *_DataTypeSchemaHeader) GetStructureDataTypes() []DataTypeDescription {
	return m.StructureDataTypes
}

func (m *_DataTypeSchemaHeader) GetNoOfEnumDataTypes() int32 {
	return m.NoOfEnumDataTypes
}

func (m *_DataTypeSchemaHeader) GetEnumDataTypes() []DataTypeDescription {
	return m.EnumDataTypes
}

func (m *_DataTypeSchemaHeader) GetNoOfSimpleDataTypes() int32 {
	return m.NoOfSimpleDataTypes
}

func (m *_DataTypeSchemaHeader) GetSimpleDataTypes() []DataTypeDescription {
	return m.SimpleDataTypes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDataTypeSchemaHeader factory function for _DataTypeSchemaHeader
func NewDataTypeSchemaHeader(noOfNamespaces int32, namespaces []PascalString, noOfStructureDataTypes int32, structureDataTypes []DataTypeDescription, noOfEnumDataTypes int32, enumDataTypes []DataTypeDescription, noOfSimpleDataTypes int32, simpleDataTypes []DataTypeDescription) *_DataTypeSchemaHeader {
	_result := &_DataTypeSchemaHeader{
		NoOfNamespaces:             noOfNamespaces,
		Namespaces:                 namespaces,
		NoOfStructureDataTypes:     noOfStructureDataTypes,
		StructureDataTypes:         structureDataTypes,
		NoOfEnumDataTypes:          noOfEnumDataTypes,
		EnumDataTypes:              enumDataTypes,
		NoOfSimpleDataTypes:        noOfSimpleDataTypes,
		SimpleDataTypes:            simpleDataTypes,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataTypeSchemaHeader(structType any) DataTypeSchemaHeader {
	if casted, ok := structType.(DataTypeSchemaHeader); ok {
		return casted
	}
	if casted, ok := structType.(*DataTypeSchemaHeader); ok {
		return *casted
	}
	return nil
}

func (m *_DataTypeSchemaHeader) GetTypeName() string {
	return "DataTypeSchemaHeader"
}

func (m *_DataTypeSchemaHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfNamespaces)
	lengthInBits += 32

	// Array field
	if len(m.Namespaces) > 0 {
		for _curItem, element := range m.Namespaces {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Namespaces), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfStructureDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.StructureDataTypes) > 0 {
		for _curItem, element := range m.StructureDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StructureDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfEnumDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.EnumDataTypes) > 0 {
		for _curItem, element := range m.EnumDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EnumDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfSimpleDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.SimpleDataTypes) > 0 {
		for _curItem, element := range m.SimpleDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SimpleDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DataTypeSchemaHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataTypeSchemaHeaderParse(ctx context.Context, theBytes []byte, identifier string) (DataTypeSchemaHeader, error) {
	return DataTypeSchemaHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DataTypeSchemaHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DataTypeSchemaHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DataTypeSchemaHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataTypeSchemaHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfNamespaces)
	_noOfNamespaces, _noOfNamespacesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfNamespaces", 32)
	if _noOfNamespacesErr != nil {
		return nil, errors.Wrap(_noOfNamespacesErr, "Error parsing 'noOfNamespaces' field of DataTypeSchemaHeader")
	}
	noOfNamespaces := _noOfNamespaces

	namespaces, err := ReadCountArrayField[PascalString](ctx, "namespaces", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfNamespaces))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaces' field"))
	}

	// Simple Field (noOfStructureDataTypes)
	_noOfStructureDataTypes, _noOfStructureDataTypesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfStructureDataTypes", 32)
	if _noOfStructureDataTypesErr != nil {
		return nil, errors.Wrap(_noOfStructureDataTypesErr, "Error parsing 'noOfStructureDataTypes' field of DataTypeSchemaHeader")
	}
	noOfStructureDataTypes := _noOfStructureDataTypes

	structureDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "structureDataTypes", ReadComplex[DataTypeDescription](func(ctx context.Context, buffer utils.ReadBuffer) (DataTypeDescription, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("14525"))
		if err != nil {
			return nil, err
		}
		return v.(DataTypeDescription), nil
	}, readBuffer), uint64(noOfStructureDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureDataTypes' field"))
	}

	// Simple Field (noOfEnumDataTypes)
	_noOfEnumDataTypes, _noOfEnumDataTypesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfEnumDataTypes", 32)
	if _noOfEnumDataTypesErr != nil {
		return nil, errors.Wrap(_noOfEnumDataTypesErr, "Error parsing 'noOfEnumDataTypes' field of DataTypeSchemaHeader")
	}
	noOfEnumDataTypes := _noOfEnumDataTypes

	enumDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "enumDataTypes", ReadComplex[DataTypeDescription](func(ctx context.Context, buffer utils.ReadBuffer) (DataTypeDescription, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("14525"))
		if err != nil {
			return nil, err
		}
		return v.(DataTypeDescription), nil
	}, readBuffer), uint64(noOfEnumDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumDataTypes' field"))
	}

	// Simple Field (noOfSimpleDataTypes)
	_noOfSimpleDataTypes, _noOfSimpleDataTypesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfSimpleDataTypes", 32)
	if _noOfSimpleDataTypesErr != nil {
		return nil, errors.Wrap(_noOfSimpleDataTypesErr, "Error parsing 'noOfSimpleDataTypes' field of DataTypeSchemaHeader")
	}
	noOfSimpleDataTypes := _noOfSimpleDataTypes

	simpleDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "simpleDataTypes", ReadComplex[DataTypeDescription](func(ctx context.Context, buffer utils.ReadBuffer) (DataTypeDescription, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("14525"))
		if err != nil {
			return nil, err
		}
		return v.(DataTypeDescription), nil
	}, readBuffer), uint64(noOfSimpleDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleDataTypes' field"))
	}

	if closeErr := readBuffer.CloseContext("DataTypeSchemaHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataTypeSchemaHeader")
	}

	// Create a partially initialized instance
	_child := &_DataTypeSchemaHeader{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfNamespaces:             noOfNamespaces,
		Namespaces:                 namespaces,
		NoOfStructureDataTypes:     noOfStructureDataTypes,
		StructureDataTypes:         structureDataTypes,
		NoOfEnumDataTypes:          noOfEnumDataTypes,
		EnumDataTypes:              enumDataTypes,
		NoOfSimpleDataTypes:        noOfSimpleDataTypes,
		SimpleDataTypes:            simpleDataTypes,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DataTypeSchemaHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataTypeSchemaHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataTypeSchemaHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataTypeSchemaHeader")
		}

		// Simple Field (noOfNamespaces)
		noOfNamespaces := int32(m.GetNoOfNamespaces())
		_noOfNamespacesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfNamespaces", 32, int32((noOfNamespaces)))
		if _noOfNamespacesErr != nil {
			return errors.Wrap(_noOfNamespacesErr, "Error serializing 'noOfNamespaces' field")
		}

		// Array Field (namespaces)
		if pushErr := writeBuffer.PushContext("namespaces", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for namespaces")
		}
		for _curItem, _element := range m.GetNamespaces() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNamespaces()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'namespaces' field")
			}
		}
		if popErr := writeBuffer.PopContext("namespaces", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for namespaces")
		}

		// Simple Field (noOfStructureDataTypes)
		noOfStructureDataTypes := int32(m.GetNoOfStructureDataTypes())
		_noOfStructureDataTypesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfStructureDataTypes", 32, int32((noOfStructureDataTypes)))
		if _noOfStructureDataTypesErr != nil {
			return errors.Wrap(_noOfStructureDataTypesErr, "Error serializing 'noOfStructureDataTypes' field")
		}

		// Array Field (structureDataTypes)
		if pushErr := writeBuffer.PushContext("structureDataTypes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for structureDataTypes")
		}
		for _curItem, _element := range m.GetStructureDataTypes() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetStructureDataTypes()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'structureDataTypes' field")
			}
		}
		if popErr := writeBuffer.PopContext("structureDataTypes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for structureDataTypes")
		}

		// Simple Field (noOfEnumDataTypes)
		noOfEnumDataTypes := int32(m.GetNoOfEnumDataTypes())
		_noOfEnumDataTypesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfEnumDataTypes", 32, int32((noOfEnumDataTypes)))
		if _noOfEnumDataTypesErr != nil {
			return errors.Wrap(_noOfEnumDataTypesErr, "Error serializing 'noOfEnumDataTypes' field")
		}

		// Array Field (enumDataTypes)
		if pushErr := writeBuffer.PushContext("enumDataTypes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enumDataTypes")
		}
		for _curItem, _element := range m.GetEnumDataTypes() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetEnumDataTypes()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'enumDataTypes' field")
			}
		}
		if popErr := writeBuffer.PopContext("enumDataTypes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enumDataTypes")
		}

		// Simple Field (noOfSimpleDataTypes)
		noOfSimpleDataTypes := int32(m.GetNoOfSimpleDataTypes())
		_noOfSimpleDataTypesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfSimpleDataTypes", 32, int32((noOfSimpleDataTypes)))
		if _noOfSimpleDataTypesErr != nil {
			return errors.Wrap(_noOfSimpleDataTypesErr, "Error serializing 'noOfSimpleDataTypes' field")
		}

		// Array Field (simpleDataTypes)
		if pushErr := writeBuffer.PushContext("simpleDataTypes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for simpleDataTypes")
		}
		for _curItem, _element := range m.GetSimpleDataTypes() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSimpleDataTypes()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'simpleDataTypes' field")
			}
		}
		if popErr := writeBuffer.PopContext("simpleDataTypes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for simpleDataTypes")
		}

		if popErr := writeBuffer.PopContext("DataTypeSchemaHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataTypeSchemaHeader")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataTypeSchemaHeader) isDataTypeSchemaHeader() bool {
	return true
}

func (m *_DataTypeSchemaHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
