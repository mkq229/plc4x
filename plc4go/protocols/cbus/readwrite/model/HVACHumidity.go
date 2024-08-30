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

// HVACHumidity is the corresponding interface of HVACHumidity
type HVACHumidity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHumidityValue returns HumidityValue (property field)
	GetHumidityValue() uint16
	// GetHumidityInPercent returns HumidityInPercent (virtual field)
	GetHumidityInPercent() float32
}

// HVACHumidityExactly can be used when we want exactly this type and not a type which fulfills HVACHumidity.
// This is useful for switch cases.
type HVACHumidityExactly interface {
	HVACHumidity
	isHVACHumidity() bool
}

// _HVACHumidity is the data-structure of this message
type _HVACHumidity struct {
	HumidityValue uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACHumidity) GetHumidityValue() uint16 {
	return m.HumidityValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACHumidity) GetHumidityInPercent() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetHumidityValue()) / float32(float32(65535)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACHumidity factory function for _HVACHumidity
func NewHVACHumidity(humidityValue uint16) *_HVACHumidity {
	return &_HVACHumidity{HumidityValue: humidityValue}
}

// Deprecated: use the interface for direct cast
func CastHVACHumidity(structType any) HVACHumidity {
	if casted, ok := structType.(HVACHumidity); ok {
		return casted
	}
	if casted, ok := structType.(*HVACHumidity); ok {
		return *casted
	}
	return nil
}

func (m *_HVACHumidity) GetTypeName() string {
	return "HVACHumidity"
}

func (m *_HVACHumidity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (humidityValue)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACHumidity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACHumidityParse(ctx context.Context, theBytes []byte) (HVACHumidity, error) {
	return HVACHumidityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACHumidityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidity, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HVACHumidity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACHumidity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (humidityValue)
	_humidityValue, _humidityValueErr := /*TODO: migrate me*/ readBuffer.ReadUint16("humidityValue", 16)
	if _humidityValueErr != nil {
		return nil, errors.Wrap(_humidityValueErr, "Error parsing 'humidityValue' field of HVACHumidity")
	}
	humidityValue := _humidityValue

	// Virtual field
	_humidityInPercent := float32(humidityValue) / float32(float32(65535))
	humidityInPercent := float32(_humidityInPercent)
	_ = humidityInPercent

	if closeErr := readBuffer.CloseContext("HVACHumidity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACHumidity")
	}

	// Create the instance
	return &_HVACHumidity{
		HumidityValue: humidityValue,
	}, nil
}

func (m *_HVACHumidity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACHumidity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACHumidity"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACHumidity")
	}

	// Simple Field (humidityValue)
	humidityValue := uint16(m.GetHumidityValue())
	_humidityValueErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("humidityValue", 16, uint16((humidityValue)))
	if _humidityValueErr != nil {
		return errors.Wrap(_humidityValueErr, "Error serializing 'humidityValue' field")
	}
	// Virtual field
	humidityInPercent := m.GetHumidityInPercent()
	_ = humidityInPercent
	if _humidityInPercentErr := writeBuffer.WriteVirtual(ctx, "humidityInPercent", m.GetHumidityInPercent()); _humidityInPercentErr != nil {
		return errors.Wrap(_humidityInPercentErr, "Error serializing 'humidityInPercent' field")
	}

	if popErr := writeBuffer.PopContext("HVACHumidity"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACHumidity")
	}
	return nil
}

func (m *_HVACHumidity) isHVACHumidity() bool {
	return true
}

func (m *_HVACHumidity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
