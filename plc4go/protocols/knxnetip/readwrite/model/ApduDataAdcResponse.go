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

// ApduDataAdcResponse is the corresponding interface of ApduDataAdcResponse
type ApduDataAdcResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataAdcResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataAdcResponse.
// This is useful for switch cases.
type ApduDataAdcResponseExactly interface {
	ApduDataAdcResponse
	isApduDataAdcResponse() bool
}

// _ApduDataAdcResponse is the data-structure of this message
type _ApduDataAdcResponse struct {
	ApduDataContract
}

var _ ApduDataAdcResponse = (*_ApduDataAdcResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataAdcResponse) GetApciType() uint8 {
	return 0x7
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataAdcResponse) GetParent() ApduDataContract {
	return m.ApduDataContract
}

// NewApduDataAdcResponse factory function for _ApduDataAdcResponse
func NewApduDataAdcResponse(dataLength uint8) *_ApduDataAdcResponse {
	_result := &_ApduDataAdcResponse{
		ApduDataContract: NewApduData(dataLength),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataAdcResponse(structType any) ApduDataAdcResponse {
	if casted, ok := structType.(ApduDataAdcResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataAdcResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataAdcResponse) GetTypeName() string {
	return "ApduDataAdcResponse"
}

func (m *_ApduDataAdcResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataAdcResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataAdcResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataAdcResponse ApduDataAdcResponse, err error) {
	m.ApduDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataAdcResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataAdcResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataAdcResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataAdcResponse")
	}

	return m, nil
}

func (m *_ApduDataAdcResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataAdcResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataAdcResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataAdcResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataAdcResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataAdcResponse")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataAdcResponse) isApduDataAdcResponse() bool {
	return true
}

func (m *_ApduDataAdcResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
