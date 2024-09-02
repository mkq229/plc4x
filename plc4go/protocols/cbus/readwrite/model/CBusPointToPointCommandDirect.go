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

// CBusPointToPointCommandDirect is the corresponding interface of CBusPointToPointCommandDirect
type CBusPointToPointCommandDirect interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusPointToPointCommand
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
}

// CBusPointToPointCommandDirectExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointCommandDirect.
// This is useful for switch cases.
type CBusPointToPointCommandDirectExactly interface {
	CBusPointToPointCommandDirect
	isCBusPointToPointCommandDirect() bool
}

// _CBusPointToPointCommandDirect is the data-structure of this message
type _CBusPointToPointCommandDirect struct {
	CBusPointToPointCommandContract
	UnitAddress UnitAddress
	// Reserved Fields
	reservedField0 *uint8
}

var _ CBusPointToPointCommandDirect = (*_CBusPointToPointCommandDirect)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToPointCommandDirect) GetParent() CBusPointToPointCommandContract {
	return m.CBusPointToPointCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommandDirect) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointCommandDirect factory function for _CBusPointToPointCommandDirect
func NewCBusPointToPointCommandDirect(unitAddress UnitAddress, bridgeAddressCountPeek uint16, calData CALData, cBusOptions CBusOptions) *_CBusPointToPointCommandDirect {
	_result := &_CBusPointToPointCommandDirect{
		CBusPointToPointCommandContract: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, cBusOptions),
		UnitAddress:                     unitAddress,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommandDirect(structType any) CBusPointToPointCommandDirect {
	if casted, ok := structType.(CBusPointToPointCommandDirect); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommandDirect); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommandDirect) GetTypeName() string {
	return "CBusPointToPointCommandDirect"
}

func (m *_CBusPointToPointCommandDirect) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).getLengthInBits(ctx))

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CBusPointToPointCommandDirect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusPointToPointCommandDirect) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusPointToPointCommand, cBusOptions CBusOptions) (__cBusPointToPointCommandDirect CBusPointToPointCommandDirect, err error) {
	m.CBusPointToPointCommandContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandDirect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommandDirect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unitAddress, err := ReadSimpleField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	m.UnitAddress = unitAddress

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandDirect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommandDirect")
	}

	return m, nil
}

func (m *_CBusPointToPointCommandDirect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointCommandDirect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandDirect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommandDirect")
		}

		if err := WriteSimpleField[UnitAddress](ctx, "unitAddress", m.GetUnitAddress(), WriteComplex[UnitAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandDirect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointCommandDirect")
		}
		return nil
	}
	return m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointCommandDirect) isCBusPointToPointCommandDirect() bool {
	return true
}

func (m *_CBusPointToPointCommandDirect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
