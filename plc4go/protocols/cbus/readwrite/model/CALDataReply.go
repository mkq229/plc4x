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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataReply is the corresponding interface of CALDataReply
type CALDataReply interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetParameterValue returns ParameterValue (property field)
	GetParameterValue() ParameterValue
}

// CALDataReplyExactly can be used when we want exactly this type and not a type which fulfills CALDataReply.
// This is useful for switch cases.
type CALDataReplyExactly interface {
	CALDataReply
	isCALDataReply() bool
}

// _CALDataReply is the data-structure of this message
type _CALDataReply struct {
	*_CALData
	ParamNo        Parameter
	ParameterValue ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataReply) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataReply) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataReply) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataReply) GetParameterValue() ParameterValue {
	return m.ParameterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataReply factory function for _CALDataReply
func NewCALDataReply(paramNo Parameter, parameterValue ParameterValue, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataReply {
	_result := &_CALDataReply{
		ParamNo:        paramNo,
		ParameterValue: parameterValue,
		_CALData:       NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataReply(structType interface{}) CALDataReply {
	if casted, ok := structType.(CALDataReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataReply) GetTypeName() string {
	return "CALDataReply"
}

func (m *_CALDataReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (parameterValue)
	lengthInBits += m.ParameterValue.GetLengthInBits()

	return lengthInBits
}

func (m *_CALDataReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataReplyParse(readBuffer utils.ReadBuffer, requestContext RequestContext, commandTypeContainer CALCommandTypeContainer) (CALDataReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	if pullErr := readBuffer.PullContext("paramNo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for paramNo")
	}
	_paramNo, _paramNoErr := ParameterParse(readBuffer)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field of CALDataReply")
	}
	paramNo := _paramNo
	if closeErr := readBuffer.CloseContext("paramNo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for paramNo")
	}

	// Simple Field (parameterValue)
	if pullErr := readBuffer.PullContext("parameterValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameterValue")
	}
	_parameterValue, _parameterValueErr := ParameterValueParse(readBuffer, ParameterType(paramNo.ParameterType()), uint8(uint8(commandTypeContainer.NumBytes())-uint8(uint8(1))))
	if _parameterValueErr != nil {
		return nil, errors.Wrap(_parameterValueErr, "Error parsing 'parameterValue' field of CALDataReply")
	}
	parameterValue := _parameterValue.(ParameterValue)
	if closeErr := readBuffer.CloseContext("parameterValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameterValue")
	}

	if closeErr := readBuffer.CloseContext("CALDataReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReply")
	}

	// Create a partially initialized instance
	_child := &_CALDataReply{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		ParamNo:        paramNo,
		ParameterValue: parameterValue,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataReply) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReply")
		}

		// Simple Field (paramNo)
		if pushErr := writeBuffer.PushContext("paramNo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for paramNo")
		}
		_paramNoErr := writeBuffer.WriteSerializable(m.GetParamNo())
		if popErr := writeBuffer.PopContext("paramNo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for paramNo")
		}
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (parameterValue)
		if pushErr := writeBuffer.PushContext("parameterValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameterValue")
		}
		_parameterValueErr := writeBuffer.WriteSerializable(m.GetParameterValue())
		if popErr := writeBuffer.PopContext("parameterValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameterValue")
		}
		if _parameterValueErr != nil {
			return errors.Wrap(_parameterValueErr, "Error serializing 'parameterValue' field")
		}

		if popErr := writeBuffer.PopContext("CALDataReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataReply) isCALDataReply() bool {
	return true
}

func (m *_CALDataReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}