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

// ModbusPDUDiagnosticResponse is the corresponding interface of ModbusPDUDiagnosticResponse
type ModbusPDUDiagnosticResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
}

// ModbusPDUDiagnosticResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUDiagnosticResponse.
// This is useful for switch cases.
type ModbusPDUDiagnosticResponseExactly interface {
	ModbusPDUDiagnosticResponse
	isModbusPDUDiagnosticResponse() bool
}

// _ModbusPDUDiagnosticResponse is the data-structure of this message
type _ModbusPDUDiagnosticResponse struct {
	ModbusPDUContract
	SubFunction uint16
	Data        uint16
}

var _ ModbusPDUDiagnosticResponse = (*_ModbusPDUDiagnosticResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUDiagnosticResponse) GetFunctionFlag() uint8 {
	return 0x08
}

func (m *_ModbusPDUDiagnosticResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *_ModbusPDUDiagnosticResponse) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUDiagnosticResponse factory function for _ModbusPDUDiagnosticResponse
func NewModbusPDUDiagnosticResponse(subFunction uint16, data uint16) *_ModbusPDUDiagnosticResponse {
	_result := &_ModbusPDUDiagnosticResponse{
		ModbusPDUContract: NewModbusPDU(),
		SubFunction:       subFunction,
		Data:              data,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUDiagnosticResponse(structType any) ModbusPDUDiagnosticResponse {
	if casted, ok := structType.(ModbusPDUDiagnosticResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUDiagnosticResponse) GetTypeName() string {
	return "ModbusPDUDiagnosticResponse"
}

func (m *_ModbusPDUDiagnosticResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (subFunction)
	lengthInBits += 16

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUDiagnosticResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUDiagnosticResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUDiagnosticResponse ModbusPDUDiagnosticResponse, err error) {
	m.ModbusPDUContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUDiagnosticResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subFunction, err := ReadSimpleField(ctx, "subFunction", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subFunction' field"))
	}
	m.SubFunction = subFunction

	data, err := ReadSimpleField(ctx, "data", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUDiagnosticResponse")
	}

	return m, nil
}

func (m *_ModbusPDUDiagnosticResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUDiagnosticResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUDiagnosticResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "subFunction", m.GetSubFunction(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'subFunction' field")
		}

		if err := WriteSimpleField[uint16](ctx, "data", m.GetData(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUDiagnosticResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUDiagnosticResponse) isModbusPDUDiagnosticResponse() bool {
	return true
}

func (m *_ModbusPDUDiagnosticResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
