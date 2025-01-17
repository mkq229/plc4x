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

// BACnetConstructedDataDescription is the corresponding interface of BACnetConstructedDataDescription
type BACnetConstructedDataDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDescription returns Description (property field)
	GetDescription() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDescription()
	// CreateBuilder creates a BACnetConstructedDataDescriptionBuilder
	CreateBACnetConstructedDataDescriptionBuilder() BACnetConstructedDataDescriptionBuilder
}

// _BACnetConstructedDataDescription is the data-structure of this message
type _BACnetConstructedDataDescription struct {
	BACnetConstructedDataContract
	Description BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataDescription = (*_BACnetConstructedDataDescription)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDescription)(nil)

// NewBACnetConstructedDataDescription factory function for _BACnetConstructedDataDescription
func NewBACnetConstructedDataDescription(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, description BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDescription {
	if description == nil {
		panic("description of type BACnetApplicationTagCharacterString for BACnetConstructedDataDescription must not be nil")
	}
	_result := &_BACnetConstructedDataDescription{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Description:                   description,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDescriptionBuilder is a builder for BACnetConstructedDataDescription
type BACnetConstructedDataDescriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(description BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionBuilder
	// WithDescription adds Description (property field)
	WithDescription(BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataDescriptionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDescription or returns an error if something is wrong
	Build() (BACnetConstructedDataDescription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDescription
}

// NewBACnetConstructedDataDescriptionBuilder() creates a BACnetConstructedDataDescriptionBuilder
func NewBACnetConstructedDataDescriptionBuilder() BACnetConstructedDataDescriptionBuilder {
	return &_BACnetConstructedDataDescriptionBuilder{_BACnetConstructedDataDescription: new(_BACnetConstructedDataDescription)}
}

type _BACnetConstructedDataDescriptionBuilder struct {
	*_BACnetConstructedDataDescription

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDescriptionBuilder) = (*_BACnetConstructedDataDescriptionBuilder)(nil)

func (b *_BACnetConstructedDataDescriptionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDescription
}

func (b *_BACnetConstructedDataDescriptionBuilder) WithMandatoryFields(description BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionBuilder {
	return b.WithDescription(description)
}

func (b *_BACnetConstructedDataDescriptionBuilder) WithDescription(description BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionBuilder {
	b.Description = description
	return b
}

func (b *_BACnetConstructedDataDescriptionBuilder) WithDescriptionBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataDescriptionBuilder {
	builder := builderSupplier(b.Description.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDescriptionBuilder) Build() (BACnetConstructedDataDescription, error) {
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDescription.deepCopy(), nil
}

func (b *_BACnetConstructedDataDescriptionBuilder) MustBuild() BACnetConstructedDataDescription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDescriptionBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDescriptionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDescriptionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDescriptionBuilder().(*_BACnetConstructedDataDescriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDescriptionBuilder creates a BACnetConstructedDataDescriptionBuilder
func (b *_BACnetConstructedDataDescription) CreateBACnetConstructedDataDescriptionBuilder() BACnetConstructedDataDescriptionBuilder {
	if b == nil {
		return NewBACnetConstructedDataDescriptionBuilder()
	}
	return &_BACnetConstructedDataDescriptionBuilder{_BACnetConstructedDataDescription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDescription) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDescription) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DESCRIPTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDescription) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDescription) GetDescription() BACnetApplicationTagCharacterString {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDescription) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDescription())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDescription(structType any) BACnetConstructedDataDescription {
	if casted, ok := structType.(BACnetConstructedDataDescription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDescription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDescription) GetTypeName() string {
	return "BACnetConstructedDataDescription"
}

func (m *_BACnetConstructedDataDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDescription BACnetConstructedDataDescription, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	description, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "description", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), description)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDescription")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDescription")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "description", m.GetDescription(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDescription")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDescription) IsBACnetConstructedDataDescription() {}

func (m *_BACnetConstructedDataDescription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDescription) deepCopy() *_BACnetConstructedDataDescription {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDescriptionCopy := &_BACnetConstructedDataDescription{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.Description),
	}
	_BACnetConstructedDataDescriptionCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDescriptionCopy
}

func (m *_BACnetConstructedDataDescription) String() string {
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
