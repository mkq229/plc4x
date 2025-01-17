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

// MediaTransportControlDataSetSelection is the corresponding interface of MediaTransportControlDataSetSelection
type MediaTransportControlDataSetSelection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetSelectionHi returns SelectionHi (property field)
	GetSelectionHi() byte
	// GetSelectionLo returns SelectionLo (property field)
	GetSelectionLo() byte
	// IsMediaTransportControlDataSetSelection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataSetSelection()
	// CreateBuilder creates a MediaTransportControlDataSetSelectionBuilder
	CreateMediaTransportControlDataSetSelectionBuilder() MediaTransportControlDataSetSelectionBuilder
}

// _MediaTransportControlDataSetSelection is the data-structure of this message
type _MediaTransportControlDataSetSelection struct {
	MediaTransportControlDataContract
	SelectionHi byte
	SelectionLo byte
}

var _ MediaTransportControlDataSetSelection = (*_MediaTransportControlDataSetSelection)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataSetSelection)(nil)

// NewMediaTransportControlDataSetSelection factory function for _MediaTransportControlDataSetSelection
func NewMediaTransportControlDataSetSelection(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, selectionHi byte, selectionLo byte) *_MediaTransportControlDataSetSelection {
	_result := &_MediaTransportControlDataSetSelection{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		SelectionHi:                       selectionHi,
		SelectionLo:                       selectionLo,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataSetSelectionBuilder is a builder for MediaTransportControlDataSetSelection
type MediaTransportControlDataSetSelectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(selectionHi byte, selectionLo byte) MediaTransportControlDataSetSelectionBuilder
	// WithSelectionHi adds SelectionHi (property field)
	WithSelectionHi(byte) MediaTransportControlDataSetSelectionBuilder
	// WithSelectionLo adds SelectionLo (property field)
	WithSelectionLo(byte) MediaTransportControlDataSetSelectionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MediaTransportControlDataBuilder
	// Build builds the MediaTransportControlDataSetSelection or returns an error if something is wrong
	Build() (MediaTransportControlDataSetSelection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataSetSelection
}

// NewMediaTransportControlDataSetSelectionBuilder() creates a MediaTransportControlDataSetSelectionBuilder
func NewMediaTransportControlDataSetSelectionBuilder() MediaTransportControlDataSetSelectionBuilder {
	return &_MediaTransportControlDataSetSelectionBuilder{_MediaTransportControlDataSetSelection: new(_MediaTransportControlDataSetSelection)}
}

type _MediaTransportControlDataSetSelectionBuilder struct {
	*_MediaTransportControlDataSetSelection

	parentBuilder *_MediaTransportControlDataBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataSetSelectionBuilder) = (*_MediaTransportControlDataSetSelectionBuilder)(nil)

func (b *_MediaTransportControlDataSetSelectionBuilder) setParent(contract MediaTransportControlDataContract) {
	b.MediaTransportControlDataContract = contract
	contract.(*_MediaTransportControlData)._SubType = b._MediaTransportControlDataSetSelection
}

func (b *_MediaTransportControlDataSetSelectionBuilder) WithMandatoryFields(selectionHi byte, selectionLo byte) MediaTransportControlDataSetSelectionBuilder {
	return b.WithSelectionHi(selectionHi).WithSelectionLo(selectionLo)
}

func (b *_MediaTransportControlDataSetSelectionBuilder) WithSelectionHi(selectionHi byte) MediaTransportControlDataSetSelectionBuilder {
	b.SelectionHi = selectionHi
	return b
}

func (b *_MediaTransportControlDataSetSelectionBuilder) WithSelectionLo(selectionLo byte) MediaTransportControlDataSetSelectionBuilder {
	b.SelectionLo = selectionLo
	return b
}

func (b *_MediaTransportControlDataSetSelectionBuilder) Build() (MediaTransportControlDataSetSelection, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlDataSetSelection.deepCopy(), nil
}

func (b *_MediaTransportControlDataSetSelectionBuilder) MustBuild() MediaTransportControlDataSetSelection {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataSetSelectionBuilder) Done() MediaTransportControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MediaTransportControlDataSetSelectionBuilder) buildForMediaTransportControlData() (MediaTransportControlData, error) {
	return b.Build()
}

func (b *_MediaTransportControlDataSetSelectionBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataSetSelectionBuilder().(*_MediaTransportControlDataSetSelectionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataSetSelectionBuilder creates a MediaTransportControlDataSetSelectionBuilder
func (b *_MediaTransportControlDataSetSelection) CreateMediaTransportControlDataSetSelectionBuilder() MediaTransportControlDataSetSelectionBuilder {
	if b == nil {
		return NewMediaTransportControlDataSetSelectionBuilder()
	}
	return &_MediaTransportControlDataSetSelectionBuilder{_MediaTransportControlDataSetSelection: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSetSelection) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSetSelection) GetSelectionHi() byte {
	return m.SelectionHi
}

func (m *_MediaTransportControlDataSetSelection) GetSelectionLo() byte {
	return m.SelectionLo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSetSelection(structType any) MediaTransportControlDataSetSelection {
	if casted, ok := structType.(MediaTransportControlDataSetSelection); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSetSelection); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSetSelection) GetTypeName() string {
	return "MediaTransportControlDataSetSelection"
}

func (m *_MediaTransportControlDataSetSelection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (selectionHi)
	lengthInBits += 8

	// Simple field (selectionLo)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlDataSetSelection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataSetSelection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataSetSelection MediaTransportControlDataSetSelection, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSetSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSetSelection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	selectionHi, err := ReadSimpleField(ctx, "selectionHi", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'selectionHi' field"))
	}
	m.SelectionHi = selectionHi

	selectionLo, err := ReadSimpleField(ctx, "selectionLo", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'selectionLo' field"))
	}
	m.SelectionLo = selectionLo

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSetSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSetSelection")
	}

	return m, nil
}

func (m *_MediaTransportControlDataSetSelection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSetSelection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSetSelection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSetSelection")
		}

		if err := WriteSimpleField[byte](ctx, "selectionHi", m.GetSelectionHi(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'selectionHi' field")
		}

		if err := WriteSimpleField[byte](ctx, "selectionLo", m.GetSelectionLo(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'selectionLo' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSetSelection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSetSelection")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataSetSelection) IsMediaTransportControlDataSetSelection() {}

func (m *_MediaTransportControlDataSetSelection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataSetSelection) deepCopy() *_MediaTransportControlDataSetSelection {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataSetSelectionCopy := &_MediaTransportControlDataSetSelection{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.SelectionHi,
		m.SelectionLo,
	}
	_MediaTransportControlDataSetSelectionCopy.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataSetSelectionCopy
}

func (m *_MediaTransportControlDataSetSelection) String() string {
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
