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

// DF1UnprotectedReadRequest is the corresponding interface of DF1UnprotectedReadRequest
type DF1UnprotectedReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1Command
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetSize returns Size (property field)
	GetSize() uint8
}

// DF1UnprotectedReadRequestExactly can be used when we want exactly this type and not a type which fulfills DF1UnprotectedReadRequest.
// This is useful for switch cases.
type DF1UnprotectedReadRequestExactly interface {
	DF1UnprotectedReadRequest
	isDF1UnprotectedReadRequest() bool
}

// _DF1UnprotectedReadRequest is the data-structure of this message
type _DF1UnprotectedReadRequest struct {
	*_DF1Command
	Address uint16
	Size    uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1UnprotectedReadRequest) GetCommandCode() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1UnprotectedReadRequest) InitializeParent(parent DF1Command, status uint8, transactionCounter uint16) {
	m.Status = status
	m.TransactionCounter = transactionCounter
}

func (m *_DF1UnprotectedReadRequest) GetParent() DF1Command {
	return m._DF1Command
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1UnprotectedReadRequest) GetAddress() uint16 {
	return m.Address
}

func (m *_DF1UnprotectedReadRequest) GetSize() uint8 {
	return m.Size
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1UnprotectedReadRequest factory function for _DF1UnprotectedReadRequest
func NewDF1UnprotectedReadRequest(address uint16, size uint8, status uint8, transactionCounter uint16) *_DF1UnprotectedReadRequest {
	_result := &_DF1UnprotectedReadRequest{
		Address:     address,
		Size:        size,
		_DF1Command: NewDF1Command(status, transactionCounter),
	}
	_result._DF1Command._DF1CommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1UnprotectedReadRequest(structType any) DF1UnprotectedReadRequest {
	if casted, ok := structType.(DF1UnprotectedReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DF1UnprotectedReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DF1UnprotectedReadRequest) GetTypeName() string {
	return "DF1UnprotectedReadRequest"
}

func (m *_DF1UnprotectedReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16

	// Simple field (size)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1UnprotectedReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1UnprotectedReadRequestParse(ctx context.Context, theBytes []byte) (DF1UnprotectedReadRequest, error) {
	return DF1UnprotectedReadRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DF1UnprotectedReadRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1UnprotectedReadRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1UnprotectedReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := /*TODO: migrate me*/ readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of DF1UnprotectedReadRequest")
	}
	address := _address

	// Simple Field (size)
	_size, _sizeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("size", 8)
	if _sizeErr != nil {
		return nil, errors.Wrap(_sizeErr, "Error parsing 'size' field of DF1UnprotectedReadRequest")
	}
	size := _size

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1UnprotectedReadRequest")
	}

	// Create a partially initialized instance
	_child := &_DF1UnprotectedReadRequest{
		_DF1Command: &_DF1Command{},
		Address:     address,
		Size:        size,
	}
	_child._DF1Command._DF1CommandChildRequirements = _child
	return _child, nil
}

func (m *_DF1UnprotectedReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1UnprotectedReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1UnprotectedReadRequest")
		}

		// Simple Field (address)
		address := uint16(m.GetAddress())
		_addressErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("address", 16, uint16((address)))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (size)
		size := uint8(m.GetSize())
		_sizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("size", 8, uint8((size)))
		if _sizeErr != nil {
			return errors.Wrap(_sizeErr, "Error serializing 'size' field")
		}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1UnprotectedReadRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1UnprotectedReadRequest) isDF1UnprotectedReadRequest() bool {
	return true
}

func (m *_DF1UnprotectedReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
