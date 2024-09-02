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

// AirConditioningDataSetHvacSetbackLimit is the corresponding interface of AirConditioningDataSetHvacSetbackLimit
type AirConditioningDataSetHvacSetbackLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACTemperature
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
}

// AirConditioningDataSetHvacSetbackLimitExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetHvacSetbackLimit.
// This is useful for switch cases.
type AirConditioningDataSetHvacSetbackLimitExactly interface {
	AirConditioningDataSetHvacSetbackLimit
	isAirConditioningDataSetHvacSetbackLimit() bool
}

// _AirConditioningDataSetHvacSetbackLimit is the data-structure of this message
type _AirConditioningDataSetHvacSetbackLimit struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACTemperature
	HvacModeAndFlags HVACModeAndFlags
}

var _ AirConditioningDataSetHvacSetbackLimit = (*_AirConditioningDataSetHvacSetbackLimit)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHvacSetbackLimit) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHvacSetbackLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetLimit() HVACTemperature {
	return m.Limit
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetHvacSetbackLimit factory function for _AirConditioningDataSetHvacSetbackLimit
func NewAirConditioningDataSetHvacSetbackLimit(zoneGroup byte, zoneList HVACZoneList, limit HVACTemperature, hvacModeAndFlags HVACModeAndFlags, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetHvacSetbackLimit {
	_result := &_AirConditioningDataSetHvacSetbackLimit{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Limit:                       limit,
		HvacModeAndFlags:            hvacModeAndFlags,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHvacSetbackLimit(structType any) AirConditioningDataSetHvacSetbackLimit {
	if casted, ok := structType.(AirConditioningDataSetHvacSetbackLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHvacSetbackLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetTypeName() string {
	return "AirConditioningDataSetHvacSetbackLimit"
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetHvacSetbackLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetHvacSetbackLimit AirConditioningDataSetHvacSetbackLimit, err error) {
	m.AirConditioningDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHvacSetbackLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHvacSetbackLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	limit, err := ReadSimpleField[HVACTemperature](ctx, "limit", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limit' field"))
	}
	m.Limit = limit

	hvacModeAndFlags, err := ReadSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACModeAndFlags](HVACModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHvacSetbackLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHvacSetbackLimit")
	}

	return m, nil
}

func (m *_AirConditioningDataSetHvacSetbackLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHvacSetbackLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHvacSetbackLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHvacSetbackLimit")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACTemperature](ctx, "limit", m.GetLimit(), WriteComplex[HVACTemperature](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limit' field")
		}

		if err := WriteSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHvacSetbackLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHvacSetbackLimit")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHvacSetbackLimit) isAirConditioningDataSetHvacSetbackLimit() bool {
	return true
}

func (m *_AirConditioningDataSetHvacSetbackLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
