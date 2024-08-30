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

// ApduDataIndividualAddressRead is the corresponding interface of ApduDataIndividualAddressRead
type ApduDataIndividualAddressRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataIndividualAddressReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataIndividualAddressRead.
// This is useful for switch cases.
type ApduDataIndividualAddressReadExactly interface {
	ApduDataIndividualAddressRead
	isApduDataIndividualAddressRead() bool
}

// _ApduDataIndividualAddressRead is the data-structure of this message
type _ApduDataIndividualAddressRead struct {
	*_ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataIndividualAddressRead) GetApciType() uint8 {
	return 0x4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataIndividualAddressRead) InitializeParent(parent ApduData) {}

func (m *_ApduDataIndividualAddressRead) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataIndividualAddressRead factory function for _ApduDataIndividualAddressRead
func NewApduDataIndividualAddressRead(dataLength uint8) *_ApduDataIndividualAddressRead {
	_result := &_ApduDataIndividualAddressRead{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataIndividualAddressRead(structType any) ApduDataIndividualAddressRead {
	if casted, ok := structType.(ApduDataIndividualAddressRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataIndividualAddressRead) GetTypeName() string {
	return "ApduDataIndividualAddressRead"
}

func (m *_ApduDataIndividualAddressRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataIndividualAddressRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataIndividualAddressReadParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduDataIndividualAddressRead, error) {
	return ApduDataIndividualAddressReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataIndividualAddressReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataIndividualAddressRead, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataIndividualAddressRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataIndividualAddressRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataIndividualAddressRead{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataIndividualAddressRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataIndividualAddressRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataIndividualAddressRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataIndividualAddressRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataIndividualAddressRead) isApduDataIndividualAddressRead() bool {
	return true
}

func (m *_ApduDataIndividualAddressRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
