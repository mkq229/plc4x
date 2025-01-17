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

// VariantFloat is the corresponding interface of VariantFloat
type VariantFloat interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []float32
	// IsVariantFloat is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantFloat()
	// CreateBuilder creates a VariantFloatBuilder
	CreateVariantFloatBuilder() VariantFloatBuilder
}

// _VariantFloat is the data-structure of this message
type _VariantFloat struct {
	VariantContract
	ArrayLength *int32
	Value       []float32
}

var _ VariantFloat = (*_VariantFloat)(nil)
var _ VariantRequirements = (*_VariantFloat)(nil)

// NewVariantFloat factory function for _VariantFloat
func NewVariantFloat(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []float32) *_VariantFloat {
	_result := &_VariantFloat{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantFloatBuilder is a builder for VariantFloat
type VariantFloatBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []float32) VariantFloatBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantFloatBuilder
	// WithValue adds Value (property field)
	WithValue(...float32) VariantFloatBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantFloat or returns an error if something is wrong
	Build() (VariantFloat, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantFloat
}

// NewVariantFloatBuilder() creates a VariantFloatBuilder
func NewVariantFloatBuilder() VariantFloatBuilder {
	return &_VariantFloatBuilder{_VariantFloat: new(_VariantFloat)}
}

type _VariantFloatBuilder struct {
	*_VariantFloat

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantFloatBuilder) = (*_VariantFloatBuilder)(nil)

func (b *_VariantFloatBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
	contract.(*_Variant)._SubType = b._VariantFloat
}

func (b *_VariantFloatBuilder) WithMandatoryFields(value []float32) VariantFloatBuilder {
	return b.WithValue(value...)
}

func (b *_VariantFloatBuilder) WithOptionalArrayLength(arrayLength int32) VariantFloatBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantFloatBuilder) WithValue(value ...float32) VariantFloatBuilder {
	b.Value = value
	return b
}

func (b *_VariantFloatBuilder) Build() (VariantFloat, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantFloat.deepCopy(), nil
}

func (b *_VariantFloatBuilder) MustBuild() VariantFloat {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantFloatBuilder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantFloatBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantFloatBuilder) DeepCopy() any {
	_copy := b.CreateVariantFloatBuilder().(*_VariantFloatBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantFloatBuilder creates a VariantFloatBuilder
func (b *_VariantFloat) CreateVariantFloatBuilder() VariantFloatBuilder {
	if b == nil {
		return NewVariantFloatBuilder()
	}
	return &_VariantFloatBuilder{_VariantFloat: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantFloat) GetVariantType() uint8 {
	return uint8(10)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantFloat) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantFloat) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantFloat) GetValue() []float32 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantFloat(structType any) VariantFloat {
	if casted, ok := structType.(VariantFloat); ok {
		return casted
	}
	if casted, ok := structType.(*VariantFloat); ok {
		return *casted
	}
	return nil
}

func (m *_VariantFloat) GetTypeName() string {
	return "VariantFloat"
}

func (m *_VariantFloat) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 32 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantFloat) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantFloat) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantFloat VariantFloat, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantFloat"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantFloat")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[float32](ctx, "value", ReadFloat(readBuffer, uint8(32)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantFloat"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantFloat")
	}

	return m, nil
}

func (m *_VariantFloat) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantFloat) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantFloat"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantFloat")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantFloat"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantFloat")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantFloat) IsVariantFloat() {}

func (m *_VariantFloat) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantFloat) deepCopy() *_VariantFloat {
	if m == nil {
		return nil
	}
	_VariantFloatCopy := &_VariantFloat{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[float32, float32](m.Value),
	}
	_VariantFloatCopy.VariantContract.(*_Variant)._SubType = m
	return _VariantFloatCopy
}

func (m *_VariantFloat) String() string {
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
