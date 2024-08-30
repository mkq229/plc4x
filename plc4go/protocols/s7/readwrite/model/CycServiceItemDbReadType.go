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

// CycServiceItemDbReadType is the corresponding interface of CycServiceItemDbReadType
type CycServiceItemDbReadType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CycServiceItemType
	// GetNumberOfAreas returns NumberOfAreas (property field)
	GetNumberOfAreas() uint8
	// GetItems returns Items (property field)
	GetItems() []SubItem
}

// CycServiceItemDbReadTypeExactly can be used when we want exactly this type and not a type which fulfills CycServiceItemDbReadType.
// This is useful for switch cases.
type CycServiceItemDbReadTypeExactly interface {
	CycServiceItemDbReadType
	isCycServiceItemDbReadType() bool
}

// _CycServiceItemDbReadType is the data-structure of this message
type _CycServiceItemDbReadType struct {
	*_CycServiceItemType
	NumberOfAreas uint8
	Items         []SubItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CycServiceItemDbReadType) InitializeParent(parent CycServiceItemType, byteLength uint8, syntaxId uint8) {
	m.ByteLength = byteLength
	m.SyntaxId = syntaxId
}

func (m *_CycServiceItemDbReadType) GetParent() CycServiceItemType {
	return m._CycServiceItemType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CycServiceItemDbReadType) GetNumberOfAreas() uint8 {
	return m.NumberOfAreas
}

func (m *_CycServiceItemDbReadType) GetItems() []SubItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCycServiceItemDbReadType factory function for _CycServiceItemDbReadType
func NewCycServiceItemDbReadType(numberOfAreas uint8, items []SubItem, byteLength uint8, syntaxId uint8) *_CycServiceItemDbReadType {
	_result := &_CycServiceItemDbReadType{
		NumberOfAreas:       numberOfAreas,
		Items:               items,
		_CycServiceItemType: NewCycServiceItemType(byteLength, syntaxId),
	}
	_result._CycServiceItemType._CycServiceItemTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCycServiceItemDbReadType(structType any) CycServiceItemDbReadType {
	if casted, ok := structType.(CycServiceItemDbReadType); ok {
		return casted
	}
	if casted, ok := structType.(*CycServiceItemDbReadType); ok {
		return *casted
	}
	return nil
}

func (m *_CycServiceItemDbReadType) GetTypeName() string {
	return "CycServiceItemDbReadType"
}

func (m *_CycServiceItemDbReadType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numberOfAreas)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CycServiceItemDbReadType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CycServiceItemDbReadTypeParse(ctx context.Context, theBytes []byte) (CycServiceItemDbReadType, error) {
	return CycServiceItemDbReadTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CycServiceItemDbReadTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CycServiceItemDbReadType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CycServiceItemDbReadType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CycServiceItemDbReadType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfAreas)
	_numberOfAreas, _numberOfAreasErr := /*TODO: migrate me*/ readBuffer.ReadUint8("numberOfAreas", 8)
	if _numberOfAreasErr != nil {
		return nil, errors.Wrap(_numberOfAreasErr, "Error parsing 'numberOfAreas' field of CycServiceItemDbReadType")
	}
	numberOfAreas := _numberOfAreas

	items, err := ReadCountArrayField[SubItem](ctx, "items", ReadComplex[SubItem](SubItemParseWithBuffer, readBuffer), uint64(numberOfAreas))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}

	if closeErr := readBuffer.CloseContext("CycServiceItemDbReadType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CycServiceItemDbReadType")
	}

	// Create a partially initialized instance
	_child := &_CycServiceItemDbReadType{
		_CycServiceItemType: &_CycServiceItemType{},
		NumberOfAreas:       numberOfAreas,
		Items:               items,
	}
	_child._CycServiceItemType._CycServiceItemTypeChildRequirements = _child
	return _child, nil
}

func (m *_CycServiceItemDbReadType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CycServiceItemDbReadType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CycServiceItemDbReadType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CycServiceItemDbReadType")
		}

		// Simple Field (numberOfAreas)
		numberOfAreas := uint8(m.GetNumberOfAreas())
		_numberOfAreasErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("numberOfAreas", 8, uint8((numberOfAreas)))
		if _numberOfAreasErr != nil {
			return errors.Wrap(_numberOfAreasErr, "Error serializing 'numberOfAreas' field")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _curItem, _element := range m.GetItems() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("CycServiceItemDbReadType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CycServiceItemDbReadType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CycServiceItemDbReadType) isCycServiceItemDbReadType() bool {
	return true
}

func (m *_CycServiceItemDbReadType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
