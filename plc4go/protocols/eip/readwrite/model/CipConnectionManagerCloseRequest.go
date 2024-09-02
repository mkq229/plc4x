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

// CipConnectionManagerCloseRequest is the corresponding interface of CipConnectionManagerCloseRequest
type CipConnectionManagerCloseRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetRequestPathSize returns RequestPathSize (property field)
	GetRequestPathSize() uint8
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetTickTime returns TickTime (property field)
	GetTickTime() uint8
	// GetTimeoutTicks returns TimeoutTicks (property field)
	GetTimeoutTicks() uint8
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetConnectionPathSize returns ConnectionPathSize (property field)
	GetConnectionPathSize() uint8
	// GetConnectionPaths returns ConnectionPaths (property field)
	GetConnectionPaths() []PathSegment
}

// CipConnectionManagerCloseRequestExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerCloseRequest.
// This is useful for switch cases.
type CipConnectionManagerCloseRequestExactly interface {
	CipConnectionManagerCloseRequest
	isCipConnectionManagerCloseRequest() bool
}

// _CipConnectionManagerCloseRequest is the data-structure of this message
type _CipConnectionManagerCloseRequest struct {
	CipServiceContract
	RequestPathSize        uint8
	ClassSegment           PathSegment
	InstanceSegment        PathSegment
	Priority               uint8
	TickTime               uint8
	TimeoutTicks           uint8
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	ConnectionPathSize     uint8
	ConnectionPaths        []PathSegment
	// Reserved Fields
	reservedField0 *byte
}

var _ CipConnectionManagerCloseRequest = (*_CipConnectionManagerCloseRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerCloseRequest) GetService() uint8 {
	return 0x4E
}

func (m *_CipConnectionManagerCloseRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipConnectionManagerCloseRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerCloseRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerCloseRequest) GetRequestPathSize() uint8 {
	return m.RequestPathSize
}

func (m *_CipConnectionManagerCloseRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_CipConnectionManagerCloseRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

func (m *_CipConnectionManagerCloseRequest) GetPriority() uint8 {
	return m.Priority
}

func (m *_CipConnectionManagerCloseRequest) GetTickTime() uint8 {
	return m.TickTime
}

func (m *_CipConnectionManagerCloseRequest) GetTimeoutTicks() uint8 {
	return m.TimeoutTicks
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerCloseRequest) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerCloseRequest) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionPathSize() uint8 {
	return m.ConnectionPathSize
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionPaths() []PathSegment {
	return m.ConnectionPaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerCloseRequest factory function for _CipConnectionManagerCloseRequest
func NewCipConnectionManagerCloseRequest(requestPathSize uint8, classSegment PathSegment, instanceSegment PathSegment, priority uint8, tickTime uint8, timeoutTicks uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, connectionPathSize uint8, connectionPaths []PathSegment, serviceLen uint16) *_CipConnectionManagerCloseRequest {
	_result := &_CipConnectionManagerCloseRequest{
		CipServiceContract:     NewCipService(serviceLen),
		RequestPathSize:        requestPathSize,
		ClassSegment:           classSegment,
		InstanceSegment:        instanceSegment,
		Priority:               priority,
		TickTime:               tickTime,
		TimeoutTicks:           timeoutTicks,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ConnectionPathSize:     connectionPathSize,
		ConnectionPaths:        connectionPaths,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerCloseRequest(structType any) CipConnectionManagerCloseRequest {
	if casted, ok := structType.(CipConnectionManagerCloseRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerCloseRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerCloseRequest) GetTypeName() string {
	return "CipConnectionManagerCloseRequest"
}

func (m *_CipConnectionManagerCloseRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Simple field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += 4

	// Simple field (tickTime)
	lengthInBits += 4

	// Simple field (timeoutTicks)
	lengthInBits += 8

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (connectionPathSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Array field
	if len(m.ConnectionPaths) > 0 {
		for _, element := range m.ConnectionPaths {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_CipConnectionManagerCloseRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectionManagerCloseRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectionManagerCloseRequest CipConnectionManagerCloseRequest, err error) {
	m.CipServiceContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerCloseRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerCloseRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadSimpleField(ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	m.RequestPathSize = requestPathSize

	classSegment, err := ReadSimpleField[PathSegment](ctx, "classSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}
	m.ClassSegment = classSegment

	instanceSegment, err := ReadSimpleField[PathSegment](ctx, "instanceSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceSegment' field"))
	}
	m.InstanceSegment = instanceSegment

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	tickTime, err := ReadSimpleField(ctx, "tickTime", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tickTime' field"))
	}
	m.TickTime = tickTime

	timeoutTicks, err := ReadSimpleField(ctx, "timeoutTicks", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeoutTicks' field"))
	}
	m.TimeoutTicks = timeoutTicks

	connectionSerialNumber, err := ReadSimpleField(ctx, "connectionSerialNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSerialNumber' field"))
	}
	m.ConnectionSerialNumber = connectionSerialNumber

	originatorVendorId, err := ReadSimpleField(ctx, "originatorVendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorVendorId' field"))
	}
	m.OriginatorVendorId = originatorVendorId

	originatorSerialNumber, err := ReadSimpleField(ctx, "originatorSerialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorSerialNumber' field"))
	}
	m.OriginatorSerialNumber = originatorSerialNumber

	connectionPathSize, err := ReadSimpleField(ctx, "connectionPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionPathSize' field"))
	}
	m.ConnectionPathSize = connectionPathSize

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	connectionPaths, err := ReadTerminatedArrayField[PathSegment](ctx, "connectionPaths", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer), NoMorePathSegments(ctx, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionPaths' field"))
	}
	m.ConnectionPaths = connectionPaths

	if closeErr := readBuffer.CloseContext("CipConnectionManagerCloseRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerCloseRequest")
	}

	return m, nil
}

func (m *_CipConnectionManagerCloseRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerCloseRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerCloseRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerCloseRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "requestPathSize", m.GetRequestPathSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "classSegment", m.GetClassSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'classSegment' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "instanceSegment", m.GetInstanceSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instanceSegment' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if err := WriteSimpleField[uint8](ctx, "tickTime", m.GetTickTime(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'tickTime' field")
		}

		if err := WriteSimpleField[uint8](ctx, "timeoutTicks", m.GetTimeoutTicks(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeoutTicks' field")
		}

		if err := WriteSimpleField[uint16](ctx, "connectionSerialNumber", m.GetConnectionSerialNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionSerialNumber' field")
		}

		if err := WriteSimpleField[uint16](ctx, "originatorVendorId", m.GetOriginatorVendorId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorVendorId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "originatorSerialNumber", m.GetOriginatorSerialNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorSerialNumber' field")
		}

		if err := WriteSimpleField[uint8](ctx, "connectionPathSize", m.GetConnectionPathSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionPathSize' field")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x00), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteComplexTypeArrayField(ctx, "connectionPaths", m.GetConnectionPaths(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionPaths' field")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerCloseRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerCloseRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerCloseRequest) isCipConnectionManagerCloseRequest() bool {
	return true
}

func (m *_CipConnectionManagerCloseRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
