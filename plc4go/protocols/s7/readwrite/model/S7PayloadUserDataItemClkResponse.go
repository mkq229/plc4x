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

// S7PayloadUserDataItemClkResponse is the corresponding interface of S7PayloadUserDataItemClkResponse
type S7PayloadUserDataItemClkResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetRes returns Res (property field)
	GetRes() uint8
	// GetYear1 returns Year1 (property field)
	GetYear1() uint8
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
}

// S7PayloadUserDataItemClkResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemClkResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemClkResponseExactly interface {
	S7PayloadUserDataItemClkResponse
	isS7PayloadUserDataItemClkResponse() bool
}

// _S7PayloadUserDataItemClkResponse is the data-structure of this message
type _S7PayloadUserDataItemClkResponse struct {
	*_S7PayloadUserDataItem
	Res       uint8
	Year1     uint8
	TimeStamp DateAndTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkResponse) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemClkResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemClkResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemClkResponse) GetRes() uint8 {
	return m.Res
}

func (m *_S7PayloadUserDataItemClkResponse) GetYear1() uint8 {
	return m.Year1
}

func (m *_S7PayloadUserDataItemClkResponse) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemClkResponse factory function for _S7PayloadUserDataItemClkResponse
func NewS7PayloadUserDataItemClkResponse(res uint8, year1 uint8, timeStamp DateAndTime, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemClkResponse {
	_result := &_S7PayloadUserDataItemClkResponse{
		Res:                    res,
		Year1:                  year1,
		TimeStamp:              timeStamp,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkResponse(structType any) S7PayloadUserDataItemClkResponse {
	if casted, ok := structType.(S7PayloadUserDataItemClkResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkResponse) GetTypeName() string {
	return "S7PayloadUserDataItemClkResponse"
}

func (m *_S7PayloadUserDataItemClkResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (res)
	lengthInBits += 8

	// Simple field (year1)
	lengthInBits += 8

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemClkResponseParse(ctx context.Context, theBytes []byte, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemClkResponse, error) {
	return S7PayloadUserDataItemClkResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemClkResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemClkResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (res)
	_res, _resErr := /*TODO: migrate me*/ readBuffer.ReadUint8("res", 8)
	if _resErr != nil {
		return nil, errors.Wrap(_resErr, "Error parsing 'res' field of S7PayloadUserDataItemClkResponse")
	}
	res := _res

	// Simple Field (year1)
	_year1, _year1Err := /*TODO: migrate me*/ readBuffer.ReadUint8("year1", 8)
	if _year1Err != nil {
		return nil, errors.Wrap(_year1Err, "Error parsing 'year1' field of S7PayloadUserDataItemClkResponse")
	}
	year1 := _year1

	// Simple Field (timeStamp)
	if pullErr := readBuffer.PullContext("timeStamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeStamp")
	}
	_timeStamp, _timeStampErr := DateAndTimeParseWithBuffer(ctx, readBuffer)
	if _timeStampErr != nil {
		return nil, errors.Wrap(_timeStampErr, "Error parsing 'timeStamp' field of S7PayloadUserDataItemClkResponse")
	}
	timeStamp := _timeStamp.(DateAndTime)
	if closeErr := readBuffer.CloseContext("timeStamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeStamp")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemClkResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Res:                    res,
		Year1:                  year1,
		TimeStamp:              timeStamp,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemClkResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkResponse")
		}

		// Simple Field (res)
		res := uint8(m.GetRes())
		_resErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("res", 8, uint8((res)))
		if _resErr != nil {
			return errors.Wrap(_resErr, "Error serializing 'res' field")
		}

		// Simple Field (year1)
		year1 := uint8(m.GetYear1())
		_year1Err := /*TODO: migrate me*/ writeBuffer.WriteUint8("year1", 8, uint8((year1)))
		if _year1Err != nil {
			return errors.Wrap(_year1Err, "Error serializing 'year1' field")
		}

		// Simple Field (timeStamp)
		if pushErr := writeBuffer.PushContext("timeStamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeStamp")
		}
		_timeStampErr := writeBuffer.WriteSerializable(ctx, m.GetTimeStamp())
		if popErr := writeBuffer.PopContext("timeStamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeStamp")
		}
		if _timeStampErr != nil {
			return errors.Wrap(_timeStampErr, "Error serializing 'timeStamp' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkResponse) isS7PayloadUserDataItemClkResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemClkResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
