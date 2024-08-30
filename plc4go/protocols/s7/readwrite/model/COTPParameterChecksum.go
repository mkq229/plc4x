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

// COTPParameterChecksum is the corresponding interface of COTPParameterChecksum
type COTPParameterChecksum interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	COTPParameter
	// GetCrc returns Crc (property field)
	GetCrc() uint8
}

// COTPParameterChecksumExactly can be used when we want exactly this type and not a type which fulfills COTPParameterChecksum.
// This is useful for switch cases.
type COTPParameterChecksumExactly interface {
	COTPParameterChecksum
	isCOTPParameterChecksum() bool
}

// _COTPParameterChecksum is the data-structure of this message
type _COTPParameterChecksum struct {
	*_COTPParameter
	Crc uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterChecksum) GetParameterType() uint8 {
	return 0xC3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterChecksum) InitializeParent(parent COTPParameter) {}

func (m *_COTPParameterChecksum) GetParent() COTPParameter {
	return m._COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterChecksum) GetCrc() uint8 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterChecksum factory function for _COTPParameterChecksum
func NewCOTPParameterChecksum(crc uint8, rest uint8) *_COTPParameterChecksum {
	_result := &_COTPParameterChecksum{
		Crc:            crc,
		_COTPParameter: NewCOTPParameter(rest),
	}
	_result._COTPParameter._COTPParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPParameterChecksum(structType any) COTPParameterChecksum {
	if casted, ok := structType.(COTPParameterChecksum); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterChecksum); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterChecksum) GetTypeName() string {
	return "COTPParameterChecksum"
}

func (m *_COTPParameterChecksum) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (crc)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameterChecksum) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPParameterChecksumParse(ctx context.Context, theBytes []byte, rest uint8) (COTPParameterChecksum, error) {
	return COTPParameterChecksumParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), rest)
}

func COTPParameterChecksumParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, rest uint8) (COTPParameterChecksum, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("COTPParameterChecksum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterChecksum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (crc)
	_crc, _crcErr := /*TODO: migrate me*/ readBuffer.ReadUint8("crc", 8)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field of COTPParameterChecksum")
	}
	crc := _crc

	if closeErr := readBuffer.CloseContext("COTPParameterChecksum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterChecksum")
	}

	// Create a partially initialized instance
	_child := &_COTPParameterChecksum{
		_COTPParameter: &_COTPParameter{
			Rest: rest,
		},
		Crc: crc,
	}
	_child._COTPParameter._COTPParameterChildRequirements = _child
	return _child, nil
}

func (m *_COTPParameterChecksum) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterChecksum) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterChecksum"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterChecksum")
		}

		// Simple Field (crc)
		crc := uint8(m.GetCrc())
		_crcErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("crc", 8, uint8((crc)))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterChecksum"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterChecksum")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterChecksum) isCOTPParameterChecksum() bool {
	return true
}

func (m *_COTPParameterChecksum) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
