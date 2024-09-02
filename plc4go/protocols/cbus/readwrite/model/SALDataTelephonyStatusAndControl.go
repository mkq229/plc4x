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

// SALDataTelephonyStatusAndControl is the corresponding interface of SALDataTelephonyStatusAndControl
type SALDataTelephonyStatusAndControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetTelephonyData returns TelephonyData (property field)
	GetTelephonyData() TelephonyData
}

// SALDataTelephonyStatusAndControlExactly can be used when we want exactly this type and not a type which fulfills SALDataTelephonyStatusAndControl.
// This is useful for switch cases.
type SALDataTelephonyStatusAndControlExactly interface {
	SALDataTelephonyStatusAndControl
	isSALDataTelephonyStatusAndControl() bool
}

// _SALDataTelephonyStatusAndControl is the data-structure of this message
type _SALDataTelephonyStatusAndControl struct {
	SALDataContract
	TelephonyData TelephonyData
}

var _ SALDataTelephonyStatusAndControl = (*_SALDataTelephonyStatusAndControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetApplicationId() ApplicationId {
	return ApplicationId_TELEPHONY_STATUS_AND_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTelephonyStatusAndControl) GetTelephonyData() TelephonyData {
	return m.TelephonyData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataTelephonyStatusAndControl factory function for _SALDataTelephonyStatusAndControl
func NewSALDataTelephonyStatusAndControl(telephonyData TelephonyData, salData SALData) *_SALDataTelephonyStatusAndControl {
	_result := &_SALDataTelephonyStatusAndControl{
		SALDataContract: NewSALData(salData),
		TelephonyData:   telephonyData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTelephonyStatusAndControl(structType any) SALDataTelephonyStatusAndControl {
	if casted, ok := structType.(SALDataTelephonyStatusAndControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTelephonyStatusAndControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTelephonyStatusAndControl) GetTypeName() string {
	return "SALDataTelephonyStatusAndControl"
}

func (m *_SALDataTelephonyStatusAndControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (telephonyData)
	lengthInBits += m.TelephonyData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTelephonyStatusAndControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataTelephonyStatusAndControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataTelephonyStatusAndControl SALDataTelephonyStatusAndControl, err error) {
	m.SALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTelephonyStatusAndControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTelephonyStatusAndControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	telephonyData, err := ReadSimpleField[TelephonyData](ctx, "telephonyData", ReadComplex[TelephonyData](TelephonyDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'telephonyData' field"))
	}
	m.TelephonyData = telephonyData

	if closeErr := readBuffer.CloseContext("SALDataTelephonyStatusAndControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTelephonyStatusAndControl")
	}

	return m, nil
}

func (m *_SALDataTelephonyStatusAndControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTelephonyStatusAndControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTelephonyStatusAndControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTelephonyStatusAndControl")
		}

		if err := WriteSimpleField[TelephonyData](ctx, "telephonyData", m.GetTelephonyData(), WriteComplex[TelephonyData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'telephonyData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTelephonyStatusAndControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTelephonyStatusAndControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTelephonyStatusAndControl) isSALDataTelephonyStatusAndControl() bool {
	return true
}

func (m *_SALDataTelephonyStatusAndControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
