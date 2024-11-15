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

// PubSubConfigurationValueDataType is the corresponding interface of PubSubConfigurationValueDataType
type PubSubConfigurationValueDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetConfigurationElement returns ConfigurationElement (property field)
	GetConfigurationElement() PubSubConfigurationRefDataType
	// GetName returns Name (property field)
	GetName() PascalString
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() Variant
	// IsPubSubConfigurationValueDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubConfigurationValueDataType()
	// CreateBuilder creates a PubSubConfigurationValueDataTypeBuilder
	CreatePubSubConfigurationValueDataTypeBuilder() PubSubConfigurationValueDataTypeBuilder
}

// _PubSubConfigurationValueDataType is the data-structure of this message
type _PubSubConfigurationValueDataType struct {
	ExtensionObjectDefinitionContract
	ConfigurationElement PubSubConfigurationRefDataType
	Name                 PascalString
	Identifier           Variant
}

var _ PubSubConfigurationValueDataType = (*_PubSubConfigurationValueDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubConfigurationValueDataType)(nil)

// NewPubSubConfigurationValueDataType factory function for _PubSubConfigurationValueDataType
func NewPubSubConfigurationValueDataType(configurationElement PubSubConfigurationRefDataType, name PascalString, identifier Variant) *_PubSubConfigurationValueDataType {
	if configurationElement == nil {
		panic("configurationElement of type PubSubConfigurationRefDataType for PubSubConfigurationValueDataType must not be nil")
	}
	if name == nil {
		panic("name of type PascalString for PubSubConfigurationValueDataType must not be nil")
	}
	if identifier == nil {
		panic("identifier of type Variant for PubSubConfigurationValueDataType must not be nil")
	}
	_result := &_PubSubConfigurationValueDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ConfigurationElement:              configurationElement,
		Name:                              name,
		Identifier:                        identifier,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PubSubConfigurationValueDataTypeBuilder is a builder for PubSubConfigurationValueDataType
type PubSubConfigurationValueDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(configurationElement PubSubConfigurationRefDataType, name PascalString, identifier Variant) PubSubConfigurationValueDataTypeBuilder
	// WithConfigurationElement adds ConfigurationElement (property field)
	WithConfigurationElement(PubSubConfigurationRefDataType) PubSubConfigurationValueDataTypeBuilder
	// WithConfigurationElementBuilder adds ConfigurationElement (property field) which is build by the builder
	WithConfigurationElementBuilder(func(PubSubConfigurationRefDataTypeBuilder) PubSubConfigurationRefDataTypeBuilder) PubSubConfigurationValueDataTypeBuilder
	// WithName adds Name (property field)
	WithName(PascalString) PubSubConfigurationValueDataTypeBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) PubSubConfigurationValueDataTypeBuilder
	// WithIdentifier adds Identifier (property field)
	WithIdentifier(Variant) PubSubConfigurationValueDataTypeBuilder
	// WithIdentifierBuilder adds Identifier (property field) which is build by the builder
	WithIdentifierBuilder(func(VariantBuilder) VariantBuilder) PubSubConfigurationValueDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the PubSubConfigurationValueDataType or returns an error if something is wrong
	Build() (PubSubConfigurationValueDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PubSubConfigurationValueDataType
}

// NewPubSubConfigurationValueDataTypeBuilder() creates a PubSubConfigurationValueDataTypeBuilder
func NewPubSubConfigurationValueDataTypeBuilder() PubSubConfigurationValueDataTypeBuilder {
	return &_PubSubConfigurationValueDataTypeBuilder{_PubSubConfigurationValueDataType: new(_PubSubConfigurationValueDataType)}
}

type _PubSubConfigurationValueDataTypeBuilder struct {
	*_PubSubConfigurationValueDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PubSubConfigurationValueDataTypeBuilder) = (*_PubSubConfigurationValueDataTypeBuilder)(nil)

func (b *_PubSubConfigurationValueDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._PubSubConfigurationValueDataType
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithMandatoryFields(configurationElement PubSubConfigurationRefDataType, name PascalString, identifier Variant) PubSubConfigurationValueDataTypeBuilder {
	return b.WithConfigurationElement(configurationElement).WithName(name).WithIdentifier(identifier)
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithConfigurationElement(configurationElement PubSubConfigurationRefDataType) PubSubConfigurationValueDataTypeBuilder {
	b.ConfigurationElement = configurationElement
	return b
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithConfigurationElementBuilder(builderSupplier func(PubSubConfigurationRefDataTypeBuilder) PubSubConfigurationRefDataTypeBuilder) PubSubConfigurationValueDataTypeBuilder {
	builder := builderSupplier(b.ConfigurationElement.CreatePubSubConfigurationRefDataTypeBuilder())
	var err error
	b.ConfigurationElement, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PubSubConfigurationRefDataTypeBuilder failed"))
	}
	return b
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithName(name PascalString) PubSubConfigurationValueDataTypeBuilder {
	b.Name = name
	return b
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) PubSubConfigurationValueDataTypeBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithIdentifier(identifier Variant) PubSubConfigurationValueDataTypeBuilder {
	b.Identifier = identifier
	return b
}

func (b *_PubSubConfigurationValueDataTypeBuilder) WithIdentifierBuilder(builderSupplier func(VariantBuilder) VariantBuilder) PubSubConfigurationValueDataTypeBuilder {
	builder := builderSupplier(b.Identifier.CreateVariantBuilder())
	var err error
	b.Identifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "VariantBuilder failed"))
	}
	return b
}

func (b *_PubSubConfigurationValueDataTypeBuilder) Build() (PubSubConfigurationValueDataType, error) {
	if b.ConfigurationElement == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'configurationElement' not set"))
	}
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.Identifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'identifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PubSubConfigurationValueDataType.deepCopy(), nil
}

func (b *_PubSubConfigurationValueDataTypeBuilder) MustBuild() PubSubConfigurationValueDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PubSubConfigurationValueDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_PubSubConfigurationValueDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PubSubConfigurationValueDataTypeBuilder) DeepCopy() any {
	_copy := b.CreatePubSubConfigurationValueDataTypeBuilder().(*_PubSubConfigurationValueDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePubSubConfigurationValueDataTypeBuilder creates a PubSubConfigurationValueDataTypeBuilder
func (b *_PubSubConfigurationValueDataType) CreatePubSubConfigurationValueDataTypeBuilder() PubSubConfigurationValueDataTypeBuilder {
	if b == nil {
		return NewPubSubConfigurationValueDataTypeBuilder()
	}
	return &_PubSubConfigurationValueDataTypeBuilder{_PubSubConfigurationValueDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConfigurationValueDataType) GetExtensionId() int32 {
	return int32(25522)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConfigurationValueDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConfigurationValueDataType) GetConfigurationElement() PubSubConfigurationRefDataType {
	return m.ConfigurationElement
}

func (m *_PubSubConfigurationValueDataType) GetName() PascalString {
	return m.Name
}

func (m *_PubSubConfigurationValueDataType) GetIdentifier() Variant {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPubSubConfigurationValueDataType(structType any) PubSubConfigurationValueDataType {
	if casted, ok := structType.(PubSubConfigurationValueDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConfigurationValueDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConfigurationValueDataType) GetTypeName() string {
	return "PubSubConfigurationValueDataType"
}

func (m *_PubSubConfigurationValueDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (configurationElement)
	lengthInBits += m.ConfigurationElement.GetLengthInBits(ctx)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (identifier)
	lengthInBits += m.Identifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PubSubConfigurationValueDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PubSubConfigurationValueDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__pubSubConfigurationValueDataType PubSubConfigurationValueDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PubSubConfigurationValueDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConfigurationValueDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	configurationElement, err := ReadSimpleField[PubSubConfigurationRefDataType](ctx, "configurationElement", ReadComplex[PubSubConfigurationRefDataType](ExtensionObjectDefinitionParseWithBufferProducer[PubSubConfigurationRefDataType]((int32)(int32(25521))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationElement' field"))
	}
	m.ConfigurationElement = configurationElement

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	identifier, err := ReadSimpleField[Variant](ctx, "identifier", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	m.Identifier = identifier

	if closeErr := readBuffer.CloseContext("PubSubConfigurationValueDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConfigurationValueDataType")
	}

	return m, nil
}

func (m *_PubSubConfigurationValueDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConfigurationValueDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConfigurationValueDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConfigurationValueDataType")
		}

		if err := WriteSimpleField[PubSubConfigurationRefDataType](ctx, "configurationElement", m.GetConfigurationElement(), WriteComplex[PubSubConfigurationRefDataType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationElement' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[Variant](ctx, "identifier", m.GetIdentifier(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConfigurationValueDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConfigurationValueDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConfigurationValueDataType) IsPubSubConfigurationValueDataType() {}

func (m *_PubSubConfigurationValueDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PubSubConfigurationValueDataType) deepCopy() *_PubSubConfigurationValueDataType {
	if m == nil {
		return nil
	}
	_PubSubConfigurationValueDataTypeCopy := &_PubSubConfigurationValueDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PubSubConfigurationRefDataType](m.ConfigurationElement),
		utils.DeepCopy[PascalString](m.Name),
		utils.DeepCopy[Variant](m.Identifier),
	}
	_PubSubConfigurationValueDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PubSubConfigurationValueDataTypeCopy
}

func (m *_PubSubConfigurationValueDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}