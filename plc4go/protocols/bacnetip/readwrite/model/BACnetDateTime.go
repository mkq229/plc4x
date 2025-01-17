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

// BACnetDateTime is the corresponding interface of BACnetDateTime
type BACnetDateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// IsBACnetDateTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDateTime()
	// CreateBuilder creates a BACnetDateTimeBuilder
	CreateBACnetDateTimeBuilder() BACnetDateTimeBuilder
}

// _BACnetDateTime is the data-structure of this message
type _BACnetDateTime struct {
	DateValue BACnetApplicationTagDate
	TimeValue BACnetApplicationTagTime
}

var _ BACnetDateTime = (*_BACnetDateTime)(nil)

// NewBACnetDateTime factory function for _BACnetDateTime
func NewBACnetDateTime(dateValue BACnetApplicationTagDate, timeValue BACnetApplicationTagTime) *_BACnetDateTime {
	if dateValue == nil {
		panic("dateValue of type BACnetApplicationTagDate for BACnetDateTime must not be nil")
	}
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetDateTime must not be nil")
	}
	return &_BACnetDateTime{DateValue: dateValue, TimeValue: timeValue}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetDateTimeBuilder is a builder for BACnetDateTime
type BACnetDateTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateValue BACnetApplicationTagDate, timeValue BACnetApplicationTagTime) BACnetDateTimeBuilder
	// WithDateValue adds DateValue (property field)
	WithDateValue(BACnetApplicationTagDate) BACnetDateTimeBuilder
	// WithDateValueBuilder adds DateValue (property field) which is build by the builder
	WithDateValueBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetDateTimeBuilder
	// WithTimeValue adds TimeValue (property field)
	WithTimeValue(BACnetApplicationTagTime) BACnetDateTimeBuilder
	// WithTimeValueBuilder adds TimeValue (property field) which is build by the builder
	WithTimeValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetDateTimeBuilder
	// Build builds the BACnetDateTime or returns an error if something is wrong
	Build() (BACnetDateTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetDateTime
}

// NewBACnetDateTimeBuilder() creates a BACnetDateTimeBuilder
func NewBACnetDateTimeBuilder() BACnetDateTimeBuilder {
	return &_BACnetDateTimeBuilder{_BACnetDateTime: new(_BACnetDateTime)}
}

type _BACnetDateTimeBuilder struct {
	*_BACnetDateTime

	err *utils.MultiError
}

var _ (BACnetDateTimeBuilder) = (*_BACnetDateTimeBuilder)(nil)

func (b *_BACnetDateTimeBuilder) WithMandatoryFields(dateValue BACnetApplicationTagDate, timeValue BACnetApplicationTagTime) BACnetDateTimeBuilder {
	return b.WithDateValue(dateValue).WithTimeValue(timeValue)
}

func (b *_BACnetDateTimeBuilder) WithDateValue(dateValue BACnetApplicationTagDate) BACnetDateTimeBuilder {
	b.DateValue = dateValue
	return b
}

func (b *_BACnetDateTimeBuilder) WithDateValueBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetDateTimeBuilder {
	builder := builderSupplier(b.DateValue.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.DateValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetDateTimeBuilder) WithTimeValue(timeValue BACnetApplicationTagTime) BACnetDateTimeBuilder {
	b.TimeValue = timeValue
	return b
}

func (b *_BACnetDateTimeBuilder) WithTimeValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetDateTimeBuilder {
	builder := builderSupplier(b.TimeValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.TimeValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetDateTimeBuilder) Build() (BACnetDateTime, error) {
	if b.DateValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dateValue' not set"))
	}
	if b.TimeValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetDateTime.deepCopy(), nil
}

func (b *_BACnetDateTimeBuilder) MustBuild() BACnetDateTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetDateTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetDateTimeBuilder().(*_BACnetDateTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetDateTimeBuilder creates a BACnetDateTimeBuilder
func (b *_BACnetDateTime) CreateBACnetDateTimeBuilder() BACnetDateTimeBuilder {
	if b == nil {
		return NewBACnetDateTimeBuilder()
	}
	return &_BACnetDateTimeBuilder{_BACnetDateTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateTime) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

func (m *_BACnetDateTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetDateTime(structType any) BACnetDateTime {
	if casted, ok := structType.(BACnetDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateTime) GetTypeName() string {
	return "BACnetDateTime"
}

func (m *_BACnetDateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateTimeParse(ctx context.Context, theBytes []byte) (BACnetDateTime, error) {
	return BACnetDateTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDateTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateTime, error) {
		return BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDateTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateTime, error) {
	v, err := (&_BACnetDateTime{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetDateTime) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetDateTime BACnetDateTime, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	m.DateValue = dateValue

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	if closeErr := readBuffer.CloseContext("BACnetDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateTime")
	}

	return m, nil
}

func (m *_BACnetDateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDateTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateTime")
	}

	if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "dateValue", m.GetDateValue(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'dateValue' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateTime")
	}
	return nil
}

func (m *_BACnetDateTime) IsBACnetDateTime() {}

func (m *_BACnetDateTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetDateTime) deepCopy() *_BACnetDateTime {
	if m == nil {
		return nil
	}
	_BACnetDateTimeCopy := &_BACnetDateTime{
		utils.DeepCopy[BACnetApplicationTagDate](m.DateValue),
		utils.DeepCopy[BACnetApplicationTagTime](m.TimeValue),
	}
	return _BACnetDateTimeCopy
}

func (m *_BACnetDateTime) String() string {
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
