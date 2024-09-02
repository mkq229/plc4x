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

// Constant values.
const CEMIAdditionalInformationBusmonitorInfo_LEN uint8 = uint8(1)

// CEMIAdditionalInformationBusmonitorInfo is the corresponding interface of CEMIAdditionalInformationBusmonitorInfo
type CEMIAdditionalInformationBusmonitorInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMIAdditionalInformation
	// GetFrameErrorFlag returns FrameErrorFlag (property field)
	GetFrameErrorFlag() bool
	// GetBitErrorFlag returns BitErrorFlag (property field)
	GetBitErrorFlag() bool
	// GetParityErrorFlag returns ParityErrorFlag (property field)
	GetParityErrorFlag() bool
	// GetUnknownFlag returns UnknownFlag (property field)
	GetUnknownFlag() bool
	// GetLostFlag returns LostFlag (property field)
	GetLostFlag() bool
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
}

// CEMIAdditionalInformationBusmonitorInfoExactly can be used when we want exactly this type and not a type which fulfills CEMIAdditionalInformationBusmonitorInfo.
// This is useful for switch cases.
type CEMIAdditionalInformationBusmonitorInfoExactly interface {
	CEMIAdditionalInformationBusmonitorInfo
	isCEMIAdditionalInformationBusmonitorInfo() bool
}

// _CEMIAdditionalInformationBusmonitorInfo is the data-structure of this message
type _CEMIAdditionalInformationBusmonitorInfo struct {
	CEMIAdditionalInformationContract
	FrameErrorFlag  bool
	BitErrorFlag    bool
	ParityErrorFlag bool
	UnknownFlag     bool
	LostFlag        bool
	SequenceNumber  uint8
}

var _ CEMIAdditionalInformationBusmonitorInfo = (*_CEMIAdditionalInformationBusmonitorInfo)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetAdditionalInformationType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetParent() CEMIAdditionalInformationContract {
	return m.CEMIAdditionalInformationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetFrameErrorFlag() bool {
	return m.FrameErrorFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetBitErrorFlag() bool {
	return m.BitErrorFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetParityErrorFlag() bool {
	return m.ParityErrorFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetUnknownFlag() bool {
	return m.UnknownFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLostFlag() bool {
	return m.LostFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLen() uint8 {
	return CEMIAdditionalInformationBusmonitorInfo_LEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCEMIAdditionalInformationBusmonitorInfo factory function for _CEMIAdditionalInformationBusmonitorInfo
func NewCEMIAdditionalInformationBusmonitorInfo(frameErrorFlag bool, bitErrorFlag bool, parityErrorFlag bool, unknownFlag bool, lostFlag bool, sequenceNumber uint8) *_CEMIAdditionalInformationBusmonitorInfo {
	_result := &_CEMIAdditionalInformationBusmonitorInfo{
		CEMIAdditionalInformationContract: NewCEMIAdditionalInformation(),
		FrameErrorFlag:                    frameErrorFlag,
		BitErrorFlag:                      bitErrorFlag,
		ParityErrorFlag:                   parityErrorFlag,
		UnknownFlag:                       unknownFlag,
		LostFlag:                          lostFlag,
		SequenceNumber:                    sequenceNumber,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformationBusmonitorInfo(structType any) CEMIAdditionalInformationBusmonitorInfo {
	if casted, ok := structType.(CEMIAdditionalInformationBusmonitorInfo); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationBusmonitorInfo); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetTypeName() string {
	return "CEMIAdditionalInformationBusmonitorInfo"
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).getLengthInBits(ctx))

	// Const Field (len)
	lengthInBits += 8

	// Simple field (frameErrorFlag)
	lengthInBits += 1

	// Simple field (bitErrorFlag)
	lengthInBits += 1

	// Simple field (parityErrorFlag)
	lengthInBits += 1

	// Simple field (unknownFlag)
	lengthInBits += 1

	// Simple field (lostFlag)
	lengthInBits += 1

	// Simple field (sequenceNumber)
	lengthInBits += 3

	return lengthInBits
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMIAdditionalInformation) (__cEMIAdditionalInformationBusmonitorInfo CEMIAdditionalInformationBusmonitorInfo, err error) {
	m.CEMIAdditionalInformationContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationBusmonitorInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformationBusmonitorInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	len, err := ReadConstField[uint8](ctx, "len", ReadUnsignedByte(readBuffer, uint8(8)), CEMIAdditionalInformationBusmonitorInfo_LEN)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'len' field"))
	}
	_ = len

	frameErrorFlag, err := ReadSimpleField(ctx, "frameErrorFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'frameErrorFlag' field"))
	}
	m.FrameErrorFlag = frameErrorFlag

	bitErrorFlag, err := ReadSimpleField(ctx, "bitErrorFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitErrorFlag' field"))
	}
	m.BitErrorFlag = bitErrorFlag

	parityErrorFlag, err := ReadSimpleField(ctx, "parityErrorFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parityErrorFlag' field"))
	}
	m.ParityErrorFlag = parityErrorFlag

	unknownFlag, err := ReadSimpleField(ctx, "unknownFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownFlag' field"))
	}
	m.UnknownFlag = unknownFlag

	lostFlag, err := ReadSimpleField(ctx, "lostFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lostFlag' field"))
	}
	m.LostFlag = lostFlag

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationBusmonitorInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformationBusmonitorInfo")
	}

	return m, nil
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationBusmonitorInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformationBusmonitorInfo")
		}

		if err := WriteConstField(ctx, "len", CEMIAdditionalInformationBusmonitorInfo_LEN, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'len' field")
		}

		if err := WriteSimpleField[bool](ctx, "frameErrorFlag", m.GetFrameErrorFlag(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'frameErrorFlag' field")
		}

		if err := WriteSimpleField[bool](ctx, "bitErrorFlag", m.GetBitErrorFlag(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitErrorFlag' field")
		}

		if err := WriteSimpleField[bool](ctx, "parityErrorFlag", m.GetParityErrorFlag(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parityErrorFlag' field")
		}

		if err := WriteSimpleField[bool](ctx, "unknownFlag", m.GetUnknownFlag(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownFlag' field")
		}

		if err := WriteSimpleField[bool](ctx, "lostFlag", m.GetLostFlag(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lostFlag' field")
		}

		if err := WriteSimpleField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationBusmonitorInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformationBusmonitorInfo")
		}
		return nil
	}
	return m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) isCEMIAdditionalInformationBusmonitorInfo() bool {
	return true
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
