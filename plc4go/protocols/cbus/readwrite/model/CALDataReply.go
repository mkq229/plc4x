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

// CALDataReply is the corresponding interface of CALDataReply
type CALDataReply interface {
	fmt.Stringer
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
	CALDataContract
	ParamNo        Parameter
	ParameterValue ParameterValue
}

var _ CALDataReply = (*_CALDataReply)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataReply) GetParent() CALDataContract {
	return m.CALDataContract
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
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		ParamNo:         paramNo,
		ParameterValue:  parameterValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataReply(structType any) CALDataReply {
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

func (m *_CALDataReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (parameterValue)
	lengthInBits += m.ParameterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALDataReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (__cALDataReply CALDataReply, err error) {
	m.CALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	parameterValue, err := ReadSimpleField[ParameterValue](ctx, "parameterValue", ReadComplex[ParameterValue](ParameterValueParseWithBufferProducer[ParameterValue]((ParameterType)(paramNo.ParameterType()), (uint8)(uint8(commandTypeContainer.NumBytes())-uint8(uint8(1)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterValue' field"))
	}
	m.ParameterValue = parameterValue

	if closeErr := readBuffer.CloseContext("CALDataReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReply")
	}

	return m, nil
}

func (m *_CALDataReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReply")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteSimpleField[ParameterValue](ctx, "parameterValue", m.GetParameterValue(), WriteComplex[ParameterValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameterValue' field")
		}

		if popErr := writeBuffer.PopContext("CALDataReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReply")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataReply) isCALDataReply() bool {
	return true
}

func (m *_CALDataReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
