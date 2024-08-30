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

// S7PayloadUserDataItemCyclicServicesUnsubscribeRequest is the corresponding interface of S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
type S7PayloadUserDataItemCyclicServicesUnsubscribeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetFunction returns Function (property field)
	GetFunction() uint8
	// GetJobId returns JobId (property field)
	GetJobId() uint8
}

// S7PayloadUserDataItemCyclicServicesUnsubscribeRequestExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCyclicServicesUnsubscribeRequest.
// This is useful for switch cases.
type S7PayloadUserDataItemCyclicServicesUnsubscribeRequestExactly interface {
	S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
	isS7PayloadUserDataItemCyclicServicesUnsubscribeRequest() bool
}

// _S7PayloadUserDataItemCyclicServicesUnsubscribeRequest is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesUnsubscribeRequest struct {
	*_S7PayloadUserDataItem
	Function uint8
	JobId    uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetCpuSubfunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetFunction() uint8 {
	return m.Function
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetJobId() uint8 {
	return m.JobId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequest factory function for _S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
func NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequest(function uint8, jobId uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest {
	_result := &_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest{
		Function:               function,
		JobId:                  jobId,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesUnsubscribeRequest(structType any) S7PayloadUserDataItemCyclicServicesUnsubscribeRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesUnsubscribeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesUnsubscribeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (function)
	lengthInBits += 8

	// Simple field (jobId)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCyclicServicesUnsubscribeRequestParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCyclicServicesUnsubscribeRequest, error) {
	return S7PayloadUserDataItemCyclicServicesUnsubscribeRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCyclicServicesUnsubscribeRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCyclicServicesUnsubscribeRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (function)
	_function, _functionErr := /*TODO: migrate me*/ readBuffer.ReadUint8("function", 8)
	if _functionErr != nil {
		return nil, errors.Wrap(_functionErr, "Error parsing 'function' field of S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
	}
	function := _function

	// Simple Field (jobId)
	_jobId, _jobIdErr := /*TODO: migrate me*/ readBuffer.ReadUint8("jobId", 8)
	if _jobIdErr != nil {
		return nil, errors.Wrap(_jobIdErr, "Error parsing 'jobId' field of S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
	}
	jobId := _jobId

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Function:               function,
		JobId:                  jobId,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
		}

		// Simple Field (function)
		function := uint8(m.GetFunction())
		_functionErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("function", 8, uint8((function)))
		if _functionErr != nil {
			return errors.Wrap(_functionErr, "Error serializing 'function' field")
		}

		// Simple Field (jobId)
		jobId := uint8(m.GetJobId())
		_jobIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("jobId", 8, uint8((jobId)))
		if _jobIdErr != nil {
			return errors.Wrap(_jobIdErr, "Error serializing 'jobId' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) isS7PayloadUserDataItemCyclicServicesUnsubscribeRequest() bool {
	return true
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
