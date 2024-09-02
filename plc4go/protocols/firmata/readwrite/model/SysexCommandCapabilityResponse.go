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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SysexCommandCapabilityResponse is the corresponding interface of SysexCommandCapabilityResponse
type SysexCommandCapabilityResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandCapabilityResponseExactly can be used when we want exactly this type and not a type which fulfills SysexCommandCapabilityResponse.
// This is useful for switch cases.
type SysexCommandCapabilityResponseExactly interface {
	SysexCommandCapabilityResponse
	isSysexCommandCapabilityResponse() bool
}

// _SysexCommandCapabilityResponse is the data-structure of this message
type _SysexCommandCapabilityResponse struct {
	SysexCommandContract
}

var _ SysexCommandCapabilityResponse = (*_SysexCommandCapabilityResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandCapabilityResponse) GetCommandType() uint8 {
	return 0x6C
}

func (m *_SysexCommandCapabilityResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandCapabilityResponse) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// NewSysexCommandCapabilityResponse factory function for _SysexCommandCapabilityResponse
func NewSysexCommandCapabilityResponse() *_SysexCommandCapabilityResponse {
	_result := &_SysexCommandCapabilityResponse{
		SysexCommandContract: NewSysexCommand(),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandCapabilityResponse(structType any) SysexCommandCapabilityResponse {
	if casted, ok := structType.(SysexCommandCapabilityResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandCapabilityResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandCapabilityResponse) GetTypeName() string {
	return "SysexCommandCapabilityResponse"
}

func (m *_SysexCommandCapabilityResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandCapabilityResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandCapabilityResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandCapabilityResponse SysexCommandCapabilityResponse, err error) {
	m.SysexCommandContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandCapabilityResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandCapabilityResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandCapabilityResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandCapabilityResponse")
	}

	return m, nil
}

func (m *_SysexCommandCapabilityResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandCapabilityResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandCapabilityResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandCapabilityResponse")
		}

		if popErr := writeBuffer.PopContext("SysexCommandCapabilityResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandCapabilityResponse")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandCapabilityResponse) isSysexCommandCapabilityResponse() bool {
	return true
}

func (m *_SysexCommandCapabilityResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
