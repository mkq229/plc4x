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
const AmsPacket_INITCOMMAND bool = bool(false)
const AmsPacket_UPDCOMMAND bool = bool(false)
const AmsPacket_TIMESTAMPADDED bool = bool(false)
const AmsPacket_HIGHPRIORITYCOMMAND bool = bool(false)
const AmsPacket_SYSTEMCOMMAND bool = bool(false)
const AmsPacket_ADSCOMMAND bool = bool(true)
const AmsPacket_NORETURN bool = bool(false)
const AmsPacket_BROADCAST bool = bool(false)

// AmsPacket is the corresponding interface of AmsPacket
type AmsPacket interface {
	AmsPacketContract
	AmsPacketRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// AmsPacketContract provides a set of functions which can be overwritten by a sub struct
type AmsPacketContract interface {
	// GetTargetAmsNetId returns TargetAmsNetId (property field)
	GetTargetAmsNetId() AmsNetId
	// GetTargetAmsPort returns TargetAmsPort (property field)
	GetTargetAmsPort() uint16
	// GetSourceAmsNetId returns SourceAmsNetId (property field)
	GetSourceAmsNetId() AmsNetId
	// GetSourceAmsPort returns SourceAmsPort (property field)
	GetSourceAmsPort() uint16
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() uint32
	// GetInvokeId returns InvokeId (property field)
	GetInvokeId() uint32
}

// AmsPacketRequirements provides a set of functions which need to be implemented by a sub struct
type AmsPacketRequirements interface {
	// GetCommandId returns CommandId (discriminator field)
	GetCommandId() CommandId
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// AmsPacketExactly can be used when we want exactly this type and not a type which fulfills AmsPacket.
// This is useful for switch cases.
type AmsPacketExactly interface {
	AmsPacket
	isAmsPacket() bool
}

// _AmsPacket is the data-structure of this message
type _AmsPacket struct {
	_AmsPacketChildRequirements
	TargetAmsNetId AmsNetId
	TargetAmsPort  uint16
	SourceAmsNetId AmsNetId
	SourceAmsPort  uint16
	ErrorCode      uint32
	InvokeId       uint32
	// Reserved Fields
	reservedField0 *int8
}

var _ AmsPacketContract = (*_AmsPacket)(nil)

type _AmsPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandId() CommandId
	GetErrorCode() uint32
	GetResponse() bool
}

type AmsPacketChild interface {
	utils.Serializable

	GetParent() *AmsPacket

	GetTypeName() string
	AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsPacket) GetTargetAmsNetId() AmsNetId {
	return m.TargetAmsNetId
}

func (m *_AmsPacket) GetTargetAmsPort() uint16 {
	return m.TargetAmsPort
}

func (m *_AmsPacket) GetSourceAmsNetId() AmsNetId {
	return m.SourceAmsNetId
}

func (m *_AmsPacket) GetSourceAmsPort() uint16 {
	return m.SourceAmsPort
}

func (m *_AmsPacket) GetErrorCode() uint32 {
	return m.ErrorCode
}

func (m *_AmsPacket) GetInvokeId() uint32 {
	return m.InvokeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AmsPacket) GetInitCommand() bool {
	return AmsPacket_INITCOMMAND
}

func (m *_AmsPacket) GetUpdCommand() bool {
	return AmsPacket_UPDCOMMAND
}

func (m *_AmsPacket) GetTimestampAdded() bool {
	return AmsPacket_TIMESTAMPADDED
}

func (m *_AmsPacket) GetHighPriorityCommand() bool {
	return AmsPacket_HIGHPRIORITYCOMMAND
}

func (m *_AmsPacket) GetSystemCommand() bool {
	return AmsPacket_SYSTEMCOMMAND
}

func (m *_AmsPacket) GetAdsCommand() bool {
	return AmsPacket_ADSCOMMAND
}

func (m *_AmsPacket) GetNoReturn() bool {
	return AmsPacket_NORETURN
}

func (m *_AmsPacket) GetBroadcast() bool {
	return AmsPacket_BROADCAST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAmsPacket factory function for _AmsPacket
func NewAmsPacket(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AmsPacket {
	return &_AmsPacket{TargetAmsNetId: targetAmsNetId, TargetAmsPort: targetAmsPort, SourceAmsNetId: sourceAmsNetId, SourceAmsPort: sourceAmsPort, ErrorCode: errorCode, InvokeId: invokeId}
}

// Deprecated: use the interface for direct cast
func CastAmsPacket(structType any) AmsPacket {
	if casted, ok := structType.(AmsPacket); ok {
		return casted
	}
	if casted, ok := structType.(*AmsPacket); ok {
		return *casted
	}
	return nil
}

func (m *_AmsPacket) GetTypeName() string {
	return "AmsPacket"
}

func (m *_AmsPacket) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (targetAmsNetId)
	lengthInBits += m.TargetAmsNetId.GetLengthInBits(ctx)

	// Simple field (targetAmsPort)
	lengthInBits += 16

	// Simple field (sourceAmsNetId)
	lengthInBits += m.SourceAmsNetId.GetLengthInBits(ctx)

	// Simple field (sourceAmsPort)
	lengthInBits += 16
	// Discriminator Field (commandId)
	lengthInBits += 16

	// Const Field (initCommand)
	lengthInBits += 1

	// Const Field (updCommand)
	lengthInBits += 1

	// Const Field (timestampAdded)
	lengthInBits += 1

	// Const Field (highPriorityCommand)
	lengthInBits += 1

	// Const Field (systemCommand)
	lengthInBits += 1

	// Const Field (adsCommand)
	lengthInBits += 1

	// Const Field (noReturn)
	lengthInBits += 1
	// Discriminator Field (response)
	lengthInBits += 1

	// Const Field (broadcast)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 7

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (errorCode)
	lengthInBits += 32

	// Simple field (invokeId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AmsPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsPacketParse[T AmsPacket](ctx context.Context, theBytes []byte) (T, error) {
	return AmsPacketParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func AmsPacketParseWithBufferProducer[T AmsPacket]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := AmsPacketParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func AmsPacketParseWithBuffer[T AmsPacket](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_AmsPacket{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_AmsPacket) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__amsPacket AmsPacket, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	targetAmsNetId, err := ReadSimpleField[AmsNetId](ctx, "targetAmsNetId", ReadComplex[AmsNetId](AmsNetIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetAmsNetId' field"))
	}
	m.TargetAmsNetId = targetAmsNetId

	targetAmsPort, err := ReadSimpleField(ctx, "targetAmsPort", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetAmsPort' field"))
	}
	m.TargetAmsPort = targetAmsPort

	sourceAmsNetId, err := ReadSimpleField[AmsNetId](ctx, "sourceAmsNetId", ReadComplex[AmsNetId](AmsNetIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAmsNetId' field"))
	}
	m.SourceAmsNetId = sourceAmsNetId

	sourceAmsPort, err := ReadSimpleField(ctx, "sourceAmsPort", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAmsPort' field"))
	}
	m.SourceAmsPort = sourceAmsPort

	commandId, err := ReadDiscriminatorEnumField[CommandId](ctx, "commandId", "CommandId", ReadEnum(CommandIdByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandId' field"))
	}

	initCommand, err := ReadConstField[bool](ctx, "initCommand", ReadBoolean(readBuffer), AmsPacket_INITCOMMAND)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initCommand' field"))
	}
	_ = initCommand

	updCommand, err := ReadConstField[bool](ctx, "updCommand", ReadBoolean(readBuffer), AmsPacket_UPDCOMMAND)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updCommand' field"))
	}
	_ = updCommand

	timestampAdded, err := ReadConstField[bool](ctx, "timestampAdded", ReadBoolean(readBuffer), AmsPacket_TIMESTAMPADDED)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampAdded' field"))
	}
	_ = timestampAdded

	highPriorityCommand, err := ReadConstField[bool](ctx, "highPriorityCommand", ReadBoolean(readBuffer), AmsPacket_HIGHPRIORITYCOMMAND)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highPriorityCommand' field"))
	}
	_ = highPriorityCommand

	systemCommand, err := ReadConstField[bool](ctx, "systemCommand", ReadBoolean(readBuffer), AmsPacket_SYSTEMCOMMAND)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemCommand' field"))
	}
	_ = systemCommand

	adsCommand, err := ReadConstField[bool](ctx, "adsCommand", ReadBoolean(readBuffer), AmsPacket_ADSCOMMAND)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsCommand' field"))
	}
	_ = adsCommand

	noReturn, err := ReadConstField[bool](ctx, "noReturn", ReadBoolean(readBuffer), AmsPacket_NORETURN)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noReturn' field"))
	}
	_ = noReturn

	response, err := ReadDiscriminatorField[bool](ctx, "response", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'response' field"))
	}

	broadcast, err := ReadConstField[bool](ctx, "broadcast", ReadBoolean(readBuffer), AmsPacket_BROADCAST)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'broadcast' field"))
	}
	_ = broadcast

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadSignedByte(readBuffer, uint8(7)), int8(0x0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	length, err := ReadImplicitField[uint32](ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	errorCode, err := ReadSimpleField(ctx, "errorCode", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCode' field"))
	}
	m.ErrorCode = errorCode

	invokeId, err := ReadSimpleField(ctx, "invokeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'invokeId' field"))
	}
	m.InvokeId = invokeId

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child AmsPacket
	switch {
	case errorCode == 0x00000000 && commandId == CommandId_INVALID && response == bool(false): // AdsInvalidRequest
		if _child, err = (&_AdsInvalidRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsInvalidRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_INVALID && response == bool(true): // AdsInvalidResponse
		if _child, err = (&_AdsInvalidResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsInvalidResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(false): // AdsReadDeviceInfoRequest
		if _child, err = (&_AdsReadDeviceInfoRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadDeviceInfoRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(true): // AdsReadDeviceInfoResponse
		if _child, err = (&_AdsReadDeviceInfoResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadDeviceInfoResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ && response == bool(false): // AdsReadRequest
		if _child, err = (&_AdsReadRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ && response == bool(true): // AdsReadResponse
		if _child, err = (&_AdsReadResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE && response == bool(false): // AdsWriteRequest
		if _child, err = (&_AdsWriteRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsWriteRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE && response == bool(true): // AdsWriteResponse
		if _child, err = (&_AdsWriteResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsWriteResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_STATE && response == bool(false): // AdsReadStateRequest
		if _child, err = (&_AdsReadStateRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadStateRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_STATE && response == bool(true): // AdsReadStateResponse
		if _child, err = (&_AdsReadStateResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadStateResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE_CONTROL && response == bool(false): // AdsWriteControlRequest
		if _child, err = (&_AdsWriteControlRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsWriteControlRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE_CONTROL && response == bool(true): // AdsWriteControlResponse
		if _child, err = (&_AdsWriteControlResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsWriteControlResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(false): // AdsAddDeviceNotificationRequest
		if _child, err = (&_AdsAddDeviceNotificationRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsAddDeviceNotificationRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(true): // AdsAddDeviceNotificationResponse
		if _child, err = (&_AdsAddDeviceNotificationResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsAddDeviceNotificationResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(false): // AdsDeleteDeviceNotificationRequest
		if _child, err = (&_AdsDeleteDeviceNotificationRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDeleteDeviceNotificationRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(true): // AdsDeleteDeviceNotificationResponse
		if _child, err = (&_AdsDeleteDeviceNotificationResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDeleteDeviceNotificationResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(false): // AdsDeviceNotificationRequest
		if _child, err = (&_AdsDeviceNotificationRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDeviceNotificationRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(true): // AdsDeviceNotificationResponse
		if _child, err = (&_AdsDeviceNotificationResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsDeviceNotificationResponse for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_WRITE && response == bool(false): // AdsReadWriteRequest
		if _child, err = (&_AdsReadWriteRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadWriteRequest for type-switch of AmsPacket")
		}
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_WRITE && response == bool(true): // AdsReadWriteResponse
		if _child, err = (&_AdsReadWriteResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AdsReadWriteResponse for type-switch of AmsPacket")
		}
	case true: // ErrorResponse
		if _child, err = (&_ErrorResponse{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorResponse for type-switch of AmsPacket")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [errorCode=%v, commandId=%v, response=%v]", errorCode, commandId, response)
	}

	if closeErr := readBuffer.CloseContext("AmsPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsPacket")
	}

	return _child, nil
}

func (pm *_AmsPacket) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AmsPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AmsPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsPacket")
	}

	if err := WriteSimpleField[AmsNetId](ctx, "targetAmsNetId", m.GetTargetAmsNetId(), WriteComplex[AmsNetId](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'targetAmsNetId' field")
	}

	if err := WriteSimpleField[uint16](ctx, "targetAmsPort", m.GetTargetAmsPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'targetAmsPort' field")
	}

	if err := WriteSimpleField[AmsNetId](ctx, "sourceAmsNetId", m.GetSourceAmsNetId(), WriteComplex[AmsNetId](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceAmsNetId' field")
	}

	if err := WriteSimpleField[uint16](ctx, "sourceAmsPort", m.GetSourceAmsPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceAmsPort' field")
	}

	if err := WriteDiscriminatorEnumField(ctx, "commandId", "CommandId", m.GetCommandId(), WriteEnum[CommandId, uint16](CommandId.GetValue, CommandId.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandId' field")
	}

	if err := WriteConstField(ctx, "initCommand", AmsPacket_INITCOMMAND, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'initCommand' field")
	}

	if err := WriteConstField(ctx, "updCommand", AmsPacket_UPDCOMMAND, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'updCommand' field")
	}

	if err := WriteConstField(ctx, "timestampAdded", AmsPacket_TIMESTAMPADDED, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestampAdded' field")
	}

	if err := WriteConstField(ctx, "highPriorityCommand", AmsPacket_HIGHPRIORITYCOMMAND, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'highPriorityCommand' field")
	}

	if err := WriteConstField(ctx, "systemCommand", AmsPacket_SYSTEMCOMMAND, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'systemCommand' field")
	}

	if err := WriteConstField(ctx, "adsCommand", AmsPacket_ADSCOMMAND, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'adsCommand' field")
	}

	if err := WriteConstField(ctx, "noReturn", AmsPacket_NORETURN, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'noReturn' field")
	}

	if err := WriteDiscriminatorField(ctx, "response", m.GetResponse(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'response' field")
	}

	if err := WriteConstField(ctx, "broadcast", AmsPacket_BROADCAST, WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'broadcast' field")
	}

	if err := WriteReservedField[int8](ctx, "reserved", int8(0x0), WriteSignedByte(writeBuffer, 7)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}
	length := uint32(uint32(uint32(m.GetLengthInBytes(ctx))) - uint32(uint32(32)))
	if err := WriteImplicitField(ctx, "length", length, WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'length' field")
	}

	if err := WriteSimpleField[uint32](ctx, "errorCode", m.GetErrorCode(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorCode' field")
	}

	if err := WriteSimpleField[uint32](ctx, "invokeId", m.GetInvokeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'invokeId' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AmsPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsPacket")
	}
	return nil
}

func (m *_AmsPacket) isAmsPacket() bool {
	return true
}
