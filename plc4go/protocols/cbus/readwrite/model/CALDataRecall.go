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

// CALDataRecall is the corresponding interface of CALDataRecall
type CALDataRecall interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCount returns Count (property field)
	GetCount() uint8
}

// CALDataRecallExactly can be used when we want exactly this type and not a type which fulfills CALDataRecall.
// This is useful for switch cases.
type CALDataRecallExactly interface {
	CALDataRecall
	isCALDataRecall() bool
}

// _CALDataRecall is the data-structure of this message
type _CALDataRecall struct {
	CALDataContract
	ParamNo Parameter
	Count   uint8
}

var _ CALDataRecall = (*_CALDataRecall)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataRecall) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataRecall) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataRecall) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataRecall factory function for _CALDataRecall
func NewCALDataRecall(paramNo Parameter, count uint8, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataRecall {
	_result := &_CALDataRecall{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		ParamNo:         paramNo,
		Count:           count,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataRecall(structType any) CALDataRecall {
	if casted, ok := structType.(CALDataRecall); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataRecall); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataRecall) GetTypeName() string {
	return "CALDataRecall"
}

func (m *_CALDataRecall) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataRecall) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataRecall) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, requestContext RequestContext) (__cALDataRecall CALDataRecall, err error) {
	m.CALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataRecall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataRecall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}
	m.ParamNo = paramNo

	count, err := ReadSimpleField(ctx, "count", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}
	m.Count = count

	if closeErr := readBuffer.CloseContext("CALDataRecall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataRecall")
	}

	return m, nil
}

func (m *_CALDataRecall) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataRecall) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRecall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataRecall")
		}

		if err := WriteSimpleEnumField[Parameter](ctx, "paramNo", "Parameter", m.GetParamNo(), WriteEnum[Parameter, uint8](Parameter.GetValue, Parameter.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'paramNo' field")
		}

		if err := WriteSimpleField[uint8](ctx, "count", m.GetCount(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRecall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataRecall")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataRecall) isCALDataRecall() bool {
	return true
}

func (m *_CALDataRecall) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
