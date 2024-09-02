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

// SALDataHvacActuator is the corresponding interface of SALDataHvacActuator
type SALDataHvacActuator interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetHvacActuatorData returns HvacActuatorData (property field)
	GetHvacActuatorData() LightingData
}

// SALDataHvacActuatorExactly can be used when we want exactly this type and not a type which fulfills SALDataHvacActuator.
// This is useful for switch cases.
type SALDataHvacActuatorExactly interface {
	SALDataHvacActuator
	isSALDataHvacActuator() bool
}

// _SALDataHvacActuator is the data-structure of this message
type _SALDataHvacActuator struct {
	SALDataContract
	HvacActuatorData LightingData
}

var _ SALDataHvacActuator = (*_SALDataHvacActuator)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataHvacActuator) GetApplicationId() ApplicationId {
	return ApplicationId_HVAC_ACTUATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataHvacActuator) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataHvacActuator) GetHvacActuatorData() LightingData {
	return m.HvacActuatorData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataHvacActuator factory function for _SALDataHvacActuator
func NewSALDataHvacActuator(hvacActuatorData LightingData, salData SALData) *_SALDataHvacActuator {
	_result := &_SALDataHvacActuator{
		SALDataContract:  NewSALData(salData),
		HvacActuatorData: hvacActuatorData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataHvacActuator(structType any) SALDataHvacActuator {
	if casted, ok := structType.(SALDataHvacActuator); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataHvacActuator); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataHvacActuator) GetTypeName() string {
	return "SALDataHvacActuator"
}

func (m *_SALDataHvacActuator) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (hvacActuatorData)
	lengthInBits += m.HvacActuatorData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataHvacActuator) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataHvacActuator) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataHvacActuator SALDataHvacActuator, err error) {
	m.SALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataHvacActuator"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataHvacActuator")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hvacActuatorData, err := ReadSimpleField[LightingData](ctx, "hvacActuatorData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacActuatorData' field"))
	}
	m.HvacActuatorData = hvacActuatorData

	if closeErr := readBuffer.CloseContext("SALDataHvacActuator"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataHvacActuator")
	}

	return m, nil
}

func (m *_SALDataHvacActuator) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataHvacActuator) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataHvacActuator"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataHvacActuator")
		}

		if err := WriteSimpleField[LightingData](ctx, "hvacActuatorData", m.GetHvacActuatorData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacActuatorData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataHvacActuator"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataHvacActuator")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataHvacActuator) isSALDataHvacActuator() bool {
	return true
}

func (m *_SALDataHvacActuator) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
