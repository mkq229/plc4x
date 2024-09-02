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

// SysexCommandStringData is the corresponding interface of SysexCommandStringData
type SysexCommandStringData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandStringDataExactly can be used when we want exactly this type and not a type which fulfills SysexCommandStringData.
// This is useful for switch cases.
type SysexCommandStringDataExactly interface {
	SysexCommandStringData
	isSysexCommandStringData() bool
}

// _SysexCommandStringData is the data-structure of this message
type _SysexCommandStringData struct {
	SysexCommandContract
}

var _ SysexCommandStringData = (*_SysexCommandStringData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandStringData) GetCommandType() uint8 {
	return 0x71
}

func (m *_SysexCommandStringData) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandStringData) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// NewSysexCommandStringData factory function for _SysexCommandStringData
func NewSysexCommandStringData() *_SysexCommandStringData {
	_result := &_SysexCommandStringData{
		SysexCommandContract: NewSysexCommand(),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandStringData(structType any) SysexCommandStringData {
	if casted, ok := structType.(SysexCommandStringData); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandStringData); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandStringData) GetTypeName() string {
	return "SysexCommandStringData"
}

func (m *_SysexCommandStringData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandStringData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandStringData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandStringData SysexCommandStringData, err error) {
	m.SysexCommandContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandStringData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandStringData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandStringData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandStringData")
	}

	return m, nil
}

func (m *_SysexCommandStringData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandStringData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandStringData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandStringData")
		}

		if popErr := writeBuffer.PopContext("SysexCommandStringData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandStringData")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandStringData) isSysexCommandStringData() bool {
	return true
}

func (m *_SysexCommandStringData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
