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

// S7ParameterWriteVarRequest is the corresponding interface of S7ParameterWriteVarRequest
type S7ParameterWriteVarRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7VarRequestParameterItem
}

// S7ParameterWriteVarRequestExactly can be used when we want exactly this type and not a type which fulfills S7ParameterWriteVarRequest.
// This is useful for switch cases.
type S7ParameterWriteVarRequestExactly interface {
	S7ParameterWriteVarRequest
	isS7ParameterWriteVarRequest() bool
}

// _S7ParameterWriteVarRequest is the data-structure of this message
type _S7ParameterWriteVarRequest struct {
	*_S7Parameter
	Items []S7VarRequestParameterItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterWriteVarRequest) GetParameterType() uint8 {
	return 0x05
}

func (m *_S7ParameterWriteVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterWriteVarRequest) InitializeParent(parent S7Parameter) {}

func (m *_S7ParameterWriteVarRequest) GetParent() S7Parameter {
	return m._S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterWriteVarRequest) GetItems() []S7VarRequestParameterItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterWriteVarRequest factory function for _S7ParameterWriteVarRequest
func NewS7ParameterWriteVarRequest(items []S7VarRequestParameterItem) *_S7ParameterWriteVarRequest {
	_result := &_S7ParameterWriteVarRequest{
		Items:        items,
		_S7Parameter: NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterWriteVarRequest(structType any) S7ParameterWriteVarRequest {
	if casted, ok := structType.(S7ParameterWriteVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterWriteVarRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterWriteVarRequest) GetTypeName() string {
	return "S7ParameterWriteVarRequest"
}

func (m *_S7ParameterWriteVarRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (numItems)
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

func (m *_S7ParameterWriteVarRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7ParameterWriteVarRequestParse(ctx context.Context, theBytes []byte, messageType uint8) (S7ParameterWriteVarRequest, error) {
	return S7ParameterWriteVarRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), messageType)
}

func S7ParameterWriteVarRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterWriteVarRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterWriteVarRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numItems, err := ReadImplicitField[uint8](ctx, "numItems", ReadUnsignedByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numItems' field"))
	}
	_ = numItems

	items, err := ReadCountArrayField[S7VarRequestParameterItem](ctx, "items", ReadComplex[S7VarRequestParameterItem](S7VarRequestParameterItemParseWithBuffer, readBuffer), uint64(numItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterWriteVarRequest")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterWriteVarRequest{
		_S7Parameter: &_S7Parameter{},
		Items:        items,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterWriteVarRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterWriteVarRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterWriteVarRequest")
		}

		// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numItems := uint8(uint8(len(m.GetItems())))
		_numItemsErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("numItems", 8, uint8((numItems)))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
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

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterWriteVarRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterWriteVarRequest) isS7ParameterWriteVarRequest() bool {
	return true
}

func (m *_S7ParameterWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
