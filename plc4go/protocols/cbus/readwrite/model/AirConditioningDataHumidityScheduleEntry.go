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
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AirConditioningDataHumidityScheduleEntry is the corresponding interface of AirConditioningDataHumidityScheduleEntry
type AirConditioningDataHumidityScheduleEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetEntry returns Entry (property field)
	GetEntry() uint8
	// GetFormat returns Format (property field)
	GetFormat() byte
	// GetHumidityModeAndFlags returns HumidityModeAndFlags (property field)
	GetHumidityModeAndFlags() HVACHumidityModeAndFlags
	// GetStartTime returns StartTime (property field)
	GetStartTime() HVACStartTime
	// GetLevel returns Level (property field)
	GetLevel() HVACHumidity
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
}

// AirConditioningDataHumidityScheduleEntryExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataHumidityScheduleEntry.
// This is useful for switch cases.
type AirConditioningDataHumidityScheduleEntryExactly interface {
	AirConditioningDataHumidityScheduleEntry
	isAirConditioningDataHumidityScheduleEntry() bool
}

// _AirConditioningDataHumidityScheduleEntry is the data-structure of this message
type _AirConditioningDataHumidityScheduleEntry struct {
	*_AirConditioningData
	ZoneGroup            byte
	ZoneList             HVACZoneList
	Entry                uint8
	Format               byte
	HumidityModeAndFlags HVACHumidityModeAndFlags
	StartTime            HVACStartTime
	Level                HVACHumidity
	RawLevel             HVACRawLevels
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataHumidityScheduleEntry) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataHumidityScheduleEntry) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetEntry() uint8 {
	return m.Entry
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetFormat() byte {
	return m.Format
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetHumidityModeAndFlags() HVACHumidityModeAndFlags {
	return m.HumidityModeAndFlags
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetStartTime() HVACStartTime {
	return m.StartTime
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetLevel() HVACHumidity {
	return m.Level
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataHumidityScheduleEntry factory function for _AirConditioningDataHumidityScheduleEntry
func NewAirConditioningDataHumidityScheduleEntry(zoneGroup byte, zoneList HVACZoneList, entry uint8, format byte, humidityModeAndFlags HVACHumidityModeAndFlags, startTime HVACStartTime, level HVACHumidity, rawLevel HVACRawLevels, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataHumidityScheduleEntry {
	_result := &_AirConditioningDataHumidityScheduleEntry{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Entry:                entry,
		Format:               format,
		HumidityModeAndFlags: humidityModeAndFlags,
		StartTime:            startTime,
		Level:                level,
		RawLevel:             rawLevel,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataHumidityScheduleEntry(structType any) AirConditioningDataHumidityScheduleEntry {
	if casted, ok := structType.(AirConditioningDataHumidityScheduleEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataHumidityScheduleEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetTypeName() string {
	return "AirConditioningDataHumidityScheduleEntry"
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (entry)
	lengthInBits += 8

	// Simple field (format)
	lengthInBits += 8

	// Simple field (humidityModeAndFlags)
	lengthInBits += m.HumidityModeAndFlags.GetLengthInBits(ctx)

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits(ctx)

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataHumidityScheduleEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataHumidityScheduleEntryParse(ctx context.Context, theBytes []byte) (AirConditioningDataHumidityScheduleEntry, error) {
	return AirConditioningDataHumidityScheduleEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataHumidityScheduleEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataHumidityScheduleEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AirConditioningDataHumidityScheduleEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataHumidityScheduleEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := /*TODO: migrate me*/ readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataHumidityScheduleEntry")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParseWithBuffer(ctx, readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataHumidityScheduleEntry")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (entry)
	_entry, _entryErr := /*TODO: migrate me*/ readBuffer.ReadUint8("entry", 8)
	if _entryErr != nil {
		return nil, errors.Wrap(_entryErr, "Error parsing 'entry' field of AirConditioningDataHumidityScheduleEntry")
	}
	entry := _entry

	// Simple Field (format)
	_format, _formatErr := /*TODO: migrate me*/ readBuffer.ReadByte("format")
	if _formatErr != nil {
		return nil, errors.Wrap(_formatErr, "Error parsing 'format' field of AirConditioningDataHumidityScheduleEntry")
	}
	format := _format

	// Simple Field (humidityModeAndFlags)
	if pullErr := readBuffer.PullContext("humidityModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidityModeAndFlags")
	}
	_humidityModeAndFlags, _humidityModeAndFlagsErr := HVACHumidityModeAndFlagsParseWithBuffer(ctx, readBuffer)
	if _humidityModeAndFlagsErr != nil {
		return nil, errors.Wrap(_humidityModeAndFlagsErr, "Error parsing 'humidityModeAndFlags' field of AirConditioningDataHumidityScheduleEntry")
	}
	humidityModeAndFlags := _humidityModeAndFlags.(HVACHumidityModeAndFlags)
	if closeErr := readBuffer.CloseContext("humidityModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidityModeAndFlags")
	}

	// Simple Field (startTime)
	if pullErr := readBuffer.PullContext("startTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for startTime")
	}
	_startTime, _startTimeErr := HVACStartTimeParseWithBuffer(ctx, readBuffer)
	if _startTimeErr != nil {
		return nil, errors.Wrap(_startTimeErr, "Error parsing 'startTime' field of AirConditioningDataHumidityScheduleEntry")
	}
	startTime := _startTime.(HVACStartTime)
	if closeErr := readBuffer.CloseContext("startTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for startTime")
	}

	// Optional Field (level) (Can be skipped, if a given expression evaluates to false)
	var level HVACHumidity = nil
	if humidityModeAndFlags.GetIsLevelHumidity() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("level"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for level")
		}
		_val, _err := HVACHumidityParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'level' field of AirConditioningDataHumidityScheduleEntry")
		default:
			level = _val.(HVACHumidity)
			if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for level")
			}
		}
	}

	// Optional Field (rawLevel) (Can be skipped, if a given expression evaluates to false)
	var rawLevel HVACRawLevels = nil
	if humidityModeAndFlags.GetIsLevelRaw() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("rawLevel"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for rawLevel")
		}
		_val, _err := HVACRawLevelsParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawLevel' field of AirConditioningDataHumidityScheduleEntry")
		default:
			rawLevel = _val.(HVACRawLevels)
			if closeErr := readBuffer.CloseContext("rawLevel"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for rawLevel")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataHumidityScheduleEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataHumidityScheduleEntry")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataHumidityScheduleEntry{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Entry:                entry,
		Format:               format,
		HumidityModeAndFlags: humidityModeAndFlags,
		StartTime:            startTime,
		Level:                level,
		RawLevel:             rawLevel,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataHumidityScheduleEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataHumidityScheduleEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataHumidityScheduleEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataHumidityScheduleEntry")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := /*TODO: migrate me*/ writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(ctx, m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (entry)
		entry := uint8(m.GetEntry())
		_entryErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("entry", 8, uint8((entry)))
		if _entryErr != nil {
			return errors.Wrap(_entryErr, "Error serializing 'entry' field")
		}

		// Simple Field (format)
		format := byte(m.GetFormat())
		_formatErr := /*TODO: migrate me*/ writeBuffer.WriteByte("format", (format))
		if _formatErr != nil {
			return errors.Wrap(_formatErr, "Error serializing 'format' field")
		}

		// Simple Field (humidityModeAndFlags)
		if pushErr := writeBuffer.PushContext("humidityModeAndFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidityModeAndFlags")
		}
		_humidityModeAndFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetHumidityModeAndFlags())
		if popErr := writeBuffer.PopContext("humidityModeAndFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidityModeAndFlags")
		}
		if _humidityModeAndFlagsErr != nil {
			return errors.Wrap(_humidityModeAndFlagsErr, "Error serializing 'humidityModeAndFlags' field")
		}

		// Simple Field (startTime)
		if pushErr := writeBuffer.PushContext("startTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for startTime")
		}
		_startTimeErr := writeBuffer.WriteSerializable(ctx, m.GetStartTime())
		if popErr := writeBuffer.PopContext("startTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for startTime")
		}
		if _startTimeErr != nil {
			return errors.Wrap(_startTimeErr, "Error serializing 'startTime' field")
		}

		// Optional Field (level) (Can be skipped, if the value is null)
		var level HVACHumidity = nil
		if m.GetLevel() != nil {
			if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for level")
			}
			level = m.GetLevel()
			_levelErr := writeBuffer.WriteSerializable(ctx, level)
			if popErr := writeBuffer.PopContext("level"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for level")
			}
			if _levelErr != nil {
				return errors.Wrap(_levelErr, "Error serializing 'level' field")
			}
		}

		// Optional Field (rawLevel) (Can be skipped, if the value is null)
		var rawLevel HVACRawLevels = nil
		if m.GetRawLevel() != nil {
			if pushErr := writeBuffer.PushContext("rawLevel"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for rawLevel")
			}
			rawLevel = m.GetRawLevel()
			_rawLevelErr := writeBuffer.WriteSerializable(ctx, rawLevel)
			if popErr := writeBuffer.PopContext("rawLevel"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for rawLevel")
			}
			if _rawLevelErr != nil {
				return errors.Wrap(_rawLevelErr, "Error serializing 'rawLevel' field")
			}
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataHumidityScheduleEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataHumidityScheduleEntry")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataHumidityScheduleEntry) isAirConditioningDataHumidityScheduleEntry() bool {
	return true
}

func (m *_AirConditioningDataHumidityScheduleEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
