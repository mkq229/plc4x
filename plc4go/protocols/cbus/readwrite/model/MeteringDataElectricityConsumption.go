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

// MeteringDataElectricityConsumption is the corresponding interface of MeteringDataElectricityConsumption
type MeteringDataElectricityConsumption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeteringData
	// GetKWhr returns KWhr (property field)
	GetKWhr() uint32
}

// MeteringDataElectricityConsumptionExactly can be used when we want exactly this type and not a type which fulfills MeteringDataElectricityConsumption.
// This is useful for switch cases.
type MeteringDataElectricityConsumptionExactly interface {
	MeteringDataElectricityConsumption
	isMeteringDataElectricityConsumption() bool
}

// _MeteringDataElectricityConsumption is the data-structure of this message
type _MeteringDataElectricityConsumption struct {
	MeteringDataContract
	KWhr uint32
}

var _ MeteringDataElectricityConsumption = (*_MeteringDataElectricityConsumption)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataElectricityConsumption) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataElectricityConsumption) GetKWhr() uint32 {
	return m.KWhr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeteringDataElectricityConsumption factory function for _MeteringDataElectricityConsumption
func NewMeteringDataElectricityConsumption(kWhr uint32, commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataElectricityConsumption {
	_result := &_MeteringDataElectricityConsumption{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
		KWhr:                 kWhr,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataElectricityConsumption(structType any) MeteringDataElectricityConsumption {
	if casted, ok := structType.(MeteringDataElectricityConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataElectricityConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataElectricityConsumption) GetTypeName() string {
	return "MeteringDataElectricityConsumption"
}

func (m *_MeteringDataElectricityConsumption) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	// Simple field (kWhr)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MeteringDataElectricityConsumption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataElectricityConsumption) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataElectricityConsumption MeteringDataElectricityConsumption, err error) {
	m.MeteringDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataElectricityConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataElectricityConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	kWhr, err := ReadSimpleField(ctx, "kWhr", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'kWhr' field"))
	}
	m.KWhr = kWhr

	if closeErr := readBuffer.CloseContext("MeteringDataElectricityConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataElectricityConsumption")
	}

	return m, nil
}

func (m *_MeteringDataElectricityConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataElectricityConsumption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataElectricityConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataElectricityConsumption")
		}

		if err := WriteSimpleField[uint32](ctx, "kWhr", m.GetKWhr(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'kWhr' field")
		}

		if popErr := writeBuffer.PopContext("MeteringDataElectricityConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataElectricityConsumption")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataElectricityConsumption) isMeteringDataElectricityConsumption() bool {
	return true
}

func (m *_MeteringDataElectricityConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
