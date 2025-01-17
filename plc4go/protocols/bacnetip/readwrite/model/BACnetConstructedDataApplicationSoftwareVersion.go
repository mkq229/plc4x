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

// BACnetConstructedDataApplicationSoftwareVersion is the corresponding interface of BACnetConstructedDataApplicationSoftwareVersion
type BACnetConstructedDataApplicationSoftwareVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetApplicationSoftwareVersion returns ApplicationSoftwareVersion (property field)
	GetApplicationSoftwareVersion() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataApplicationSoftwareVersion is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataApplicationSoftwareVersion()
	// CreateBuilder creates a BACnetConstructedDataApplicationSoftwareVersionBuilder
	CreateBACnetConstructedDataApplicationSoftwareVersionBuilder() BACnetConstructedDataApplicationSoftwareVersionBuilder
}

// _BACnetConstructedDataApplicationSoftwareVersion is the data-structure of this message
type _BACnetConstructedDataApplicationSoftwareVersion struct {
	BACnetConstructedDataContract
	ApplicationSoftwareVersion BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataApplicationSoftwareVersion = (*_BACnetConstructedDataApplicationSoftwareVersion)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataApplicationSoftwareVersion)(nil)

// NewBACnetConstructedDataApplicationSoftwareVersion factory function for _BACnetConstructedDataApplicationSoftwareVersion
func NewBACnetConstructedDataApplicationSoftwareVersion(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, applicationSoftwareVersion BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataApplicationSoftwareVersion {
	if applicationSoftwareVersion == nil {
		panic("applicationSoftwareVersion of type BACnetApplicationTagCharacterString for BACnetConstructedDataApplicationSoftwareVersion must not be nil")
	}
	_result := &_BACnetConstructedDataApplicationSoftwareVersion{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ApplicationSoftwareVersion:    applicationSoftwareVersion,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataApplicationSoftwareVersionBuilder is a builder for BACnetConstructedDataApplicationSoftwareVersion
type BACnetConstructedDataApplicationSoftwareVersionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(applicationSoftwareVersion BACnetApplicationTagCharacterString) BACnetConstructedDataApplicationSoftwareVersionBuilder
	// WithApplicationSoftwareVersion adds ApplicationSoftwareVersion (property field)
	WithApplicationSoftwareVersion(BACnetApplicationTagCharacterString) BACnetConstructedDataApplicationSoftwareVersionBuilder
	// WithApplicationSoftwareVersionBuilder adds ApplicationSoftwareVersion (property field) which is build by the builder
	WithApplicationSoftwareVersionBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataApplicationSoftwareVersionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataApplicationSoftwareVersion or returns an error if something is wrong
	Build() (BACnetConstructedDataApplicationSoftwareVersion, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataApplicationSoftwareVersion
}

// NewBACnetConstructedDataApplicationSoftwareVersionBuilder() creates a BACnetConstructedDataApplicationSoftwareVersionBuilder
func NewBACnetConstructedDataApplicationSoftwareVersionBuilder() BACnetConstructedDataApplicationSoftwareVersionBuilder {
	return &_BACnetConstructedDataApplicationSoftwareVersionBuilder{_BACnetConstructedDataApplicationSoftwareVersion: new(_BACnetConstructedDataApplicationSoftwareVersion)}
}

type _BACnetConstructedDataApplicationSoftwareVersionBuilder struct {
	*_BACnetConstructedDataApplicationSoftwareVersion

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataApplicationSoftwareVersionBuilder) = (*_BACnetConstructedDataApplicationSoftwareVersionBuilder)(nil)

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataApplicationSoftwareVersion
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) WithMandatoryFields(applicationSoftwareVersion BACnetApplicationTagCharacterString) BACnetConstructedDataApplicationSoftwareVersionBuilder {
	return b.WithApplicationSoftwareVersion(applicationSoftwareVersion)
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) WithApplicationSoftwareVersion(applicationSoftwareVersion BACnetApplicationTagCharacterString) BACnetConstructedDataApplicationSoftwareVersionBuilder {
	b.ApplicationSoftwareVersion = applicationSoftwareVersion
	return b
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) WithApplicationSoftwareVersionBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataApplicationSoftwareVersionBuilder {
	builder := builderSupplier(b.ApplicationSoftwareVersion.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.ApplicationSoftwareVersion, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) Build() (BACnetConstructedDataApplicationSoftwareVersion, error) {
	if b.ApplicationSoftwareVersion == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'applicationSoftwareVersion' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataApplicationSoftwareVersion.deepCopy(), nil
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) MustBuild() BACnetConstructedDataApplicationSoftwareVersion {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataApplicationSoftwareVersionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataApplicationSoftwareVersionBuilder().(*_BACnetConstructedDataApplicationSoftwareVersionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataApplicationSoftwareVersionBuilder creates a BACnetConstructedDataApplicationSoftwareVersionBuilder
func (b *_BACnetConstructedDataApplicationSoftwareVersion) CreateBACnetConstructedDataApplicationSoftwareVersionBuilder() BACnetConstructedDataApplicationSoftwareVersionBuilder {
	if b == nil {
		return NewBACnetConstructedDataApplicationSoftwareVersionBuilder()
	}
	return &_BACnetConstructedDataApplicationSoftwareVersionBuilder{_BACnetConstructedDataApplicationSoftwareVersion: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APPLICATION_SOFTWARE_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetApplicationSoftwareVersion() BACnetApplicationTagCharacterString {
	return m.ApplicationSoftwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetApplicationSoftwareVersion())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataApplicationSoftwareVersion(structType any) BACnetConstructedDataApplicationSoftwareVersion {
	if casted, ok := structType.(BACnetConstructedDataApplicationSoftwareVersion); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataApplicationSoftwareVersion); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetTypeName() string {
	return "BACnetConstructedDataApplicationSoftwareVersion"
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (applicationSoftwareVersion)
	lengthInBits += m.ApplicationSoftwareVersion.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataApplicationSoftwareVersion BACnetConstructedDataApplicationSoftwareVersion, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataApplicationSoftwareVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataApplicationSoftwareVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	applicationSoftwareVersion, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "applicationSoftwareVersion", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationSoftwareVersion' field"))
	}
	m.ApplicationSoftwareVersion = applicationSoftwareVersion

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), applicationSoftwareVersion)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataApplicationSoftwareVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataApplicationSoftwareVersion")
	}

	return m, nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataApplicationSoftwareVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataApplicationSoftwareVersion")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "applicationSoftwareVersion", m.GetApplicationSoftwareVersion(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'applicationSoftwareVersion' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataApplicationSoftwareVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataApplicationSoftwareVersion")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) IsBACnetConstructedDataApplicationSoftwareVersion() {
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) deepCopy() *_BACnetConstructedDataApplicationSoftwareVersion {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataApplicationSoftwareVersionCopy := &_BACnetConstructedDataApplicationSoftwareVersion{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.ApplicationSoftwareVersion),
	}
	_BACnetConstructedDataApplicationSoftwareVersionCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataApplicationSoftwareVersionCopy
}

func (m *_BACnetConstructedDataApplicationSoftwareVersion) String() string {
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
