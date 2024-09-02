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

// SALDataIrrigationControl is the corresponding interface of SALDataIrrigationControl
type SALDataIrrigationControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetIrrigationControlData returns IrrigationControlData (property field)
	GetIrrigationControlData() LightingData
}

// SALDataIrrigationControlExactly can be used when we want exactly this type and not a type which fulfills SALDataIrrigationControl.
// This is useful for switch cases.
type SALDataIrrigationControlExactly interface {
	SALDataIrrigationControl
	isSALDataIrrigationControl() bool
}

// _SALDataIrrigationControl is the data-structure of this message
type _SALDataIrrigationControl struct {
	SALDataContract
	IrrigationControlData LightingData
}

var _ SALDataIrrigationControl = (*_SALDataIrrigationControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataIrrigationControl) GetApplicationId() ApplicationId {
	return ApplicationId_IRRIGATION_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataIrrigationControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataIrrigationControl) GetIrrigationControlData() LightingData {
	return m.IrrigationControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataIrrigationControl factory function for _SALDataIrrigationControl
func NewSALDataIrrigationControl(irrigationControlData LightingData, salData SALData) *_SALDataIrrigationControl {
	_result := &_SALDataIrrigationControl{
		SALDataContract:       NewSALData(salData),
		IrrigationControlData: irrigationControlData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataIrrigationControl(structType any) SALDataIrrigationControl {
	if casted, ok := structType.(SALDataIrrigationControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataIrrigationControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataIrrigationControl) GetTypeName() string {
	return "SALDataIrrigationControl"
}

func (m *_SALDataIrrigationControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (irrigationControlData)
	lengthInBits += m.IrrigationControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataIrrigationControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataIrrigationControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataIrrigationControl SALDataIrrigationControl, err error) {
	m.SALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataIrrigationControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataIrrigationControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	irrigationControlData, err := ReadSimpleField[LightingData](ctx, "irrigationControlData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'irrigationControlData' field"))
	}
	m.IrrigationControlData = irrigationControlData

	if closeErr := readBuffer.CloseContext("SALDataIrrigationControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataIrrigationControl")
	}

	return m, nil
}

func (m *_SALDataIrrigationControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataIrrigationControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataIrrigationControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataIrrigationControl")
		}

		if err := WriteSimpleField[LightingData](ctx, "irrigationControlData", m.GetIrrigationControlData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'irrigationControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataIrrigationControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataIrrigationControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataIrrigationControl) isSALDataIrrigationControl() bool {
	return true
}

func (m *_SALDataIrrigationControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
