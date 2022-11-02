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

// SysexCommandReportFirmwareRequest is the corresponding interface of SysexCommandReportFirmwareRequest
type SysexCommandReportFirmwareRequest interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandReportFirmwareRequestExactly can be used when we want exactly this type and not a type which fulfills SysexCommandReportFirmwareRequest.
// This is useful for switch cases.
type SysexCommandReportFirmwareRequestExactly interface {
	SysexCommandReportFirmwareRequest
	isSysexCommandReportFirmwareRequest() bool
}

// _SysexCommandReportFirmwareRequest is the data-structure of this message
type _SysexCommandReportFirmwareRequest struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandReportFirmwareRequest) GetCommandType() uint8 {
	return 0x79
}

func (m *_SysexCommandReportFirmwareRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandReportFirmwareRequest) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandReportFirmwareRequest) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandReportFirmwareRequest factory function for _SysexCommandReportFirmwareRequest
func NewSysexCommandReportFirmwareRequest() *_SysexCommandReportFirmwareRequest {
	_result := &_SysexCommandReportFirmwareRequest{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandReportFirmwareRequest(structType interface{}) SysexCommandReportFirmwareRequest {
	if casted, ok := structType.(SysexCommandReportFirmwareRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandReportFirmwareRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandReportFirmwareRequest) GetTypeName() string {
	return "SysexCommandReportFirmwareRequest"
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandReportFirmwareRequestParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandReportFirmwareRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandReportFirmwareRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandReportFirmwareRequest")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandReportFirmwareRequest{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandReportFirmwareRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandReportFirmwareRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandReportFirmwareRequest")
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandReportFirmwareRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SysexCommandReportFirmwareRequest) isSysexCommandReportFirmwareRequest() bool {
	return true
}

func (m *_SysexCommandReportFirmwareRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}