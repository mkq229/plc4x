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

// AdsWriteControlRequest is the corresponding interface of AdsWriteControlRequest
type AdsWriteControlRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetAdsState returns AdsState (property field)
	GetAdsState() uint16
	// GetDeviceState returns DeviceState (property field)
	GetDeviceState() uint16
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsWriteControlRequestExactly can be used when we want exactly this type and not a type which fulfills AdsWriteControlRequest.
// This is useful for switch cases.
type AdsWriteControlRequestExactly interface {
	AdsWriteControlRequest
	isAdsWriteControlRequest() bool
}

// _AdsWriteControlRequest is the data-structure of this message
type _AdsWriteControlRequest struct {
	*_AmsPacket
	AdsState    uint16
	DeviceState uint16
	Data        []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsWriteControlRequest) GetCommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *_AdsWriteControlRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsWriteControlRequest) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsWriteControlRequest) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsWriteControlRequest) GetAdsState() uint16 {
	return m.AdsState
}

func (m *_AdsWriteControlRequest) GetDeviceState() uint16 {
	return m.DeviceState
}

func (m *_AdsWriteControlRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsWriteControlRequest factory function for _AdsWriteControlRequest
func NewAdsWriteControlRequest(adsState uint16, deviceState uint16, data []byte, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsWriteControlRequest {
	_result := &_AdsWriteControlRequest{
		AdsState:    adsState,
		DeviceState: deviceState,
		Data:        data,
		_AmsPacket:  NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsWriteControlRequest(structType any) AdsWriteControlRequest {
	if casted, ok := structType.(AdsWriteControlRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsWriteControlRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsWriteControlRequest) GetTypeName() string {
	return "AdsWriteControlRequest"
}

func (m *_AdsWriteControlRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsWriteControlRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsWriteControlRequestParse(ctx context.Context, theBytes []byte) (AdsWriteControlRequest, error) {
	return AdsWriteControlRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsWriteControlRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsWriteControlRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsWriteControlRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsWriteControlRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (adsState)
	_adsState, _adsStateErr := /*TODO: migrate me*/ readBuffer.ReadUint16("adsState", 16)
	if _adsStateErr != nil {
		return nil, errors.Wrap(_adsStateErr, "Error parsing 'adsState' field of AdsWriteControlRequest")
	}
	adsState := _adsState

	// Simple Field (deviceState)
	_deviceState, _deviceStateErr := /*TODO: migrate me*/ readBuffer.ReadUint16("deviceState", 16)
	if _deviceStateErr != nil {
		return nil, errors.Wrap(_deviceStateErr, "Error parsing 'deviceState' field of AdsWriteControlRequest")
	}
	deviceState := _deviceState

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AdsWriteControlRequest")
	}

	data, err := readBuffer.ReadByteArray("data", int(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsWriteControlRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsWriteControlRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsWriteControlRequest{
		_AmsPacket:  &_AmsPacket{},
		AdsState:    adsState,
		DeviceState: deviceState,
		Data:        data,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsWriteControlRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsWriteControlRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteControlRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsWriteControlRequest")
		}

		// Simple Field (adsState)
		adsState := uint16(m.GetAdsState())
		_adsStateErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("adsState", 16, uint16((adsState)))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.GetDeviceState())
		_deviceStateErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("deviceState", 16, uint16((deviceState)))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.GetData())))
		_lengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("length", 32, uint32((length)))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsWriteControlRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsWriteControlRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsWriteControlRequest) isAdsWriteControlRequest() bool {
	return true
}

func (m *_AdsWriteControlRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
