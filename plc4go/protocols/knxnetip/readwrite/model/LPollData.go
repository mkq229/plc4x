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

// LPollData is the corresponding interface of LPollData
type LPollData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LDataFrame
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() KnxAddress
	// GetTargetAddress returns TargetAddress (property field)
	GetTargetAddress() []byte
	// GetNumberExpectedPollData returns NumberExpectedPollData (property field)
	GetNumberExpectedPollData() uint8
}

// LPollDataExactly can be used when we want exactly this type and not a type which fulfills LPollData.
// This is useful for switch cases.
type LPollDataExactly interface {
	LPollData
	isLPollData() bool
}

// _LPollData is the data-structure of this message
type _LPollData struct {
	LDataFrameContract
	SourceAddress          KnxAddress
	TargetAddress          []byte
	NumberExpectedPollData uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ LPollData = (*_LPollData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LPollData) GetNotAckFrame() bool {
	return bool(true)
}

func (m *_LPollData) GetPolling() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LPollData) GetParent() LDataFrameContract {
	return m.LDataFrameContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LPollData) GetSourceAddress() KnxAddress {
	return m.SourceAddress
}

func (m *_LPollData) GetTargetAddress() []byte {
	return m.TargetAddress
}

func (m *_LPollData) GetNumberExpectedPollData() uint8 {
	return m.NumberExpectedPollData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLPollData factory function for _LPollData
func NewLPollData(sourceAddress KnxAddress, targetAddress []byte, numberExpectedPollData uint8, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *_LPollData {
	_result := &_LPollData{
		LDataFrameContract:     NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
		SourceAddress:          sourceAddress,
		TargetAddress:          targetAddress,
		NumberExpectedPollData: numberExpectedPollData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastLPollData(structType any) LPollData {
	if casted, ok := structType.(LPollData); ok {
		return casted
	}
	if casted, ok := structType.(*LPollData); ok {
		return *casted
	}
	return nil
}

func (m *_LPollData) GetTypeName() string {
	return "LPollData"
}

func (m *_LPollData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LDataFrameContract.(*_LDataFrame).getLengthInBits(ctx))

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.GetLengthInBits(ctx)

	// Array field
	if len(m.TargetAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.TargetAddress))
	}

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (numberExpectedPollData)
	lengthInBits += 6

	return lengthInBits
}

func (m *_LPollData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LPollData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LDataFrame) (__lPollData LPollData, err error) {
	m.LDataFrameContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LPollData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LPollData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sourceAddress, err := ReadSimpleField[KnxAddress](ctx, "sourceAddress", ReadComplex[KnxAddress](KnxAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAddress' field"))
	}
	m.SourceAddress = sourceAddress

	targetAddress, err := readBuffer.ReadByteArray("targetAddress", int(int32(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetAddress' field"))
	}
	m.TargetAddress = targetAddress

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	numberExpectedPollData, err := ReadSimpleField(ctx, "numberExpectedPollData", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberExpectedPollData' field"))
	}
	m.NumberExpectedPollData = numberExpectedPollData

	if closeErr := readBuffer.CloseContext("LPollData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LPollData")
	}

	return m, nil
}

func (m *_LPollData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LPollData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LPollData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LPollData")
		}

		if err := WriteSimpleField[KnxAddress](ctx, "sourceAddress", m.GetSourceAddress(), WriteComplex[KnxAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceAddress' field")
		}

		if err := WriteByteArrayField(ctx, "targetAddress", m.GetTargetAddress(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetAddress' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "numberExpectedPollData", m.GetNumberExpectedPollData(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberExpectedPollData' field")
		}

		if popErr := writeBuffer.PopContext("LPollData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LPollData")
		}
		return nil
	}
	return m.LDataFrameContract.(*_LDataFrame).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LPollData) isLPollData() bool {
	return true
}

func (m *_LPollData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
