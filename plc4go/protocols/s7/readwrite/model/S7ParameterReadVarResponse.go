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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7ParameterReadVarResponse is the corresponding interface of S7ParameterReadVarResponse
type S7ParameterReadVarResponse interface {
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetNumItems returns NumItems (property field)
	GetNumItems() uint8
}

// S7ParameterReadVarResponseExactly can be used when we want exactly this type and not a type which fulfills S7ParameterReadVarResponse.
// This is useful for switch cases.
type S7ParameterReadVarResponseExactly interface {
	S7ParameterReadVarResponse
	isS7ParameterReadVarResponse() bool
}

// _S7ParameterReadVarResponse is the data-structure of this message
type _S7ParameterReadVarResponse struct {
	*_S7Parameter
	NumItems uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterReadVarResponse) GetParameterType() uint8 {
	return 0x04
}

func (m *_S7ParameterReadVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterReadVarResponse) InitializeParent(parent S7Parameter) {}

func (m *_S7ParameterReadVarResponse) GetParent() S7Parameter {
	return m._S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterReadVarResponse) GetNumItems() uint8 {
	return m.NumItems
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterReadVarResponse factory function for _S7ParameterReadVarResponse
func NewS7ParameterReadVarResponse(numItems uint8) *_S7ParameterReadVarResponse {
	_result := &_S7ParameterReadVarResponse{
		NumItems:     numItems,
		_S7Parameter: NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterReadVarResponse(structType interface{}) S7ParameterReadVarResponse {
	if casted, ok := structType.(S7ParameterReadVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterReadVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterReadVarResponse) GetTypeName() string {
	return "S7ParameterReadVarResponse"
}

func (m *_S7ParameterReadVarResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7ParameterReadVarResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numItems)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterReadVarResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterReadVarResponseParse(readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterReadVarResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterReadVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterReadVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numItems)
	_numItems, _numItemsErr := readBuffer.ReadUint8("numItems", 8)
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field of S7ParameterReadVarResponse")
	}
	numItems := _numItems

	if closeErr := readBuffer.CloseContext("S7ParameterReadVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterReadVarResponse")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterReadVarResponse{
		_S7Parameter: &_S7Parameter{},
		NumItems:     numItems,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterReadVarResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterReadVarResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterReadVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterReadVarResponse")
		}

		// Simple Field (numItems)
		numItems := uint8(m.GetNumItems())
		_numItemsErr := writeBuffer.WriteUint8("numItems", 8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterReadVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterReadVarResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7ParameterReadVarResponse) isS7ParameterReadVarResponse() bool {
	return true
}

func (m *_S7ParameterReadVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}