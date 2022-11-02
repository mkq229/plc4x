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

// BACnetLiftFault is an enum
type BACnetLiftFault uint16

type IBACnetLiftFault interface {
	utils.Serializable
}

const (
	BACnetLiftFault_CONTROLLER_FAULT                 BACnetLiftFault = 0
	BACnetLiftFault_DRIVE_AND_MOTOR_FAULT            BACnetLiftFault = 1
	BACnetLiftFault_GOVERNOR_AND_SAFETY_GEAR_FAULT   BACnetLiftFault = 2
	BACnetLiftFault_LIFT_SHAFT_DEVICE_FAULT          BACnetLiftFault = 3
	BACnetLiftFault_POWER_SUPPLY_FAULT               BACnetLiftFault = 4
	BACnetLiftFault_SAFETY_INTERLOCK_FAULT           BACnetLiftFault = 5
	BACnetLiftFault_DOOR_CLOSING_FAULT               BACnetLiftFault = 6
	BACnetLiftFault_DOOR_OPENING_FAULT               BACnetLiftFault = 7
	BACnetLiftFault_CAR_STOPPED_OUTSIDE_LANDING_ZONE BACnetLiftFault = 8
	BACnetLiftFault_CALL_BUTTON_STUCK                BACnetLiftFault = 9
	BACnetLiftFault_START_FAILURE                    BACnetLiftFault = 10
	BACnetLiftFault_CONTROLLER_SUPPLY_FAULT          BACnetLiftFault = 11
	BACnetLiftFault_SELF_TEST_FAILURE                BACnetLiftFault = 12
	BACnetLiftFault_RUNTIME_LIMIT_EXCEEDED           BACnetLiftFault = 13
	BACnetLiftFault_POSITION_LOST                    BACnetLiftFault = 14
	BACnetLiftFault_DRIVE_TEMPERATURE_EXCEEDED       BACnetLiftFault = 15
	BACnetLiftFault_LOAD_MEASUREMENT_FAULT           BACnetLiftFault = 16
	BACnetLiftFault_VENDOR_PROPRIETARY_VALUE         BACnetLiftFault = 0xFFFF
)

var BACnetLiftFaultValues []BACnetLiftFault

func init() {
	_ = errors.New
	BACnetLiftFaultValues = []BACnetLiftFault{
		BACnetLiftFault_CONTROLLER_FAULT,
		BACnetLiftFault_DRIVE_AND_MOTOR_FAULT,
		BACnetLiftFault_GOVERNOR_AND_SAFETY_GEAR_FAULT,
		BACnetLiftFault_LIFT_SHAFT_DEVICE_FAULT,
		BACnetLiftFault_POWER_SUPPLY_FAULT,
		BACnetLiftFault_SAFETY_INTERLOCK_FAULT,
		BACnetLiftFault_DOOR_CLOSING_FAULT,
		BACnetLiftFault_DOOR_OPENING_FAULT,
		BACnetLiftFault_CAR_STOPPED_OUTSIDE_LANDING_ZONE,
		BACnetLiftFault_CALL_BUTTON_STUCK,
		BACnetLiftFault_START_FAILURE,
		BACnetLiftFault_CONTROLLER_SUPPLY_FAULT,
		BACnetLiftFault_SELF_TEST_FAILURE,
		BACnetLiftFault_RUNTIME_LIMIT_EXCEEDED,
		BACnetLiftFault_POSITION_LOST,
		BACnetLiftFault_DRIVE_TEMPERATURE_EXCEEDED,
		BACnetLiftFault_LOAD_MEASUREMENT_FAULT,
		BACnetLiftFault_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLiftFaultByValue(value uint16) (enum BACnetLiftFault, ok bool) {
	switch value {
	case 0:
		return BACnetLiftFault_CONTROLLER_FAULT, true
	case 0xFFFF:
		return BACnetLiftFault_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLiftFault_DRIVE_AND_MOTOR_FAULT, true
	case 10:
		return BACnetLiftFault_START_FAILURE, true
	case 11:
		return BACnetLiftFault_CONTROLLER_SUPPLY_FAULT, true
	case 12:
		return BACnetLiftFault_SELF_TEST_FAILURE, true
	case 13:
		return BACnetLiftFault_RUNTIME_LIMIT_EXCEEDED, true
	case 14:
		return BACnetLiftFault_POSITION_LOST, true
	case 15:
		return BACnetLiftFault_DRIVE_TEMPERATURE_EXCEEDED, true
	case 16:
		return BACnetLiftFault_LOAD_MEASUREMENT_FAULT, true
	case 2:
		return BACnetLiftFault_GOVERNOR_AND_SAFETY_GEAR_FAULT, true
	case 3:
		return BACnetLiftFault_LIFT_SHAFT_DEVICE_FAULT, true
	case 4:
		return BACnetLiftFault_POWER_SUPPLY_FAULT, true
	case 5:
		return BACnetLiftFault_SAFETY_INTERLOCK_FAULT, true
	case 6:
		return BACnetLiftFault_DOOR_CLOSING_FAULT, true
	case 7:
		return BACnetLiftFault_DOOR_OPENING_FAULT, true
	case 8:
		return BACnetLiftFault_CAR_STOPPED_OUTSIDE_LANDING_ZONE, true
	case 9:
		return BACnetLiftFault_CALL_BUTTON_STUCK, true
	}
	return 0, false
}

func BACnetLiftFaultByName(value string) (enum BACnetLiftFault, ok bool) {
	switch value {
	case "CONTROLLER_FAULT":
		return BACnetLiftFault_CONTROLLER_FAULT, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLiftFault_VENDOR_PROPRIETARY_VALUE, true
	case "DRIVE_AND_MOTOR_FAULT":
		return BACnetLiftFault_DRIVE_AND_MOTOR_FAULT, true
	case "START_FAILURE":
		return BACnetLiftFault_START_FAILURE, true
	case "CONTROLLER_SUPPLY_FAULT":
		return BACnetLiftFault_CONTROLLER_SUPPLY_FAULT, true
	case "SELF_TEST_FAILURE":
		return BACnetLiftFault_SELF_TEST_FAILURE, true
	case "RUNTIME_LIMIT_EXCEEDED":
		return BACnetLiftFault_RUNTIME_LIMIT_EXCEEDED, true
	case "POSITION_LOST":
		return BACnetLiftFault_POSITION_LOST, true
	case "DRIVE_TEMPERATURE_EXCEEDED":
		return BACnetLiftFault_DRIVE_TEMPERATURE_EXCEEDED, true
	case "LOAD_MEASUREMENT_FAULT":
		return BACnetLiftFault_LOAD_MEASUREMENT_FAULT, true
	case "GOVERNOR_AND_SAFETY_GEAR_FAULT":
		return BACnetLiftFault_GOVERNOR_AND_SAFETY_GEAR_FAULT, true
	case "LIFT_SHAFT_DEVICE_FAULT":
		return BACnetLiftFault_LIFT_SHAFT_DEVICE_FAULT, true
	case "POWER_SUPPLY_FAULT":
		return BACnetLiftFault_POWER_SUPPLY_FAULT, true
	case "SAFETY_INTERLOCK_FAULT":
		return BACnetLiftFault_SAFETY_INTERLOCK_FAULT, true
	case "DOOR_CLOSING_FAULT":
		return BACnetLiftFault_DOOR_CLOSING_FAULT, true
	case "DOOR_OPENING_FAULT":
		return BACnetLiftFault_DOOR_OPENING_FAULT, true
	case "CAR_STOPPED_OUTSIDE_LANDING_ZONE":
		return BACnetLiftFault_CAR_STOPPED_OUTSIDE_LANDING_ZONE, true
	case "CALL_BUTTON_STUCK":
		return BACnetLiftFault_CALL_BUTTON_STUCK, true
	}
	return 0, false
}

func BACnetLiftFaultKnows(value uint16) bool {
	for _, typeValue := range BACnetLiftFaultValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftFault(structType interface{}) BACnetLiftFault {
	castFunc := func(typ interface{}) BACnetLiftFault {
		if sBACnetLiftFault, ok := typ.(BACnetLiftFault); ok {
			return sBACnetLiftFault
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftFault) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLiftFault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftFaultParse(readBuffer utils.ReadBuffer) (BACnetLiftFault, error) {
	val, err := readBuffer.ReadUint16("BACnetLiftFault", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLiftFault")
	}
	if enum, ok := BACnetLiftFaultByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLiftFault(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLiftFault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLiftFault) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLiftFault", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftFault) PLC4XEnumName() string {
	switch e {
	case BACnetLiftFault_CONTROLLER_FAULT:
		return "CONTROLLER_FAULT"
	case BACnetLiftFault_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLiftFault_DRIVE_AND_MOTOR_FAULT:
		return "DRIVE_AND_MOTOR_FAULT"
	case BACnetLiftFault_START_FAILURE:
		return "START_FAILURE"
	case BACnetLiftFault_CONTROLLER_SUPPLY_FAULT:
		return "CONTROLLER_SUPPLY_FAULT"
	case BACnetLiftFault_SELF_TEST_FAILURE:
		return "SELF_TEST_FAILURE"
	case BACnetLiftFault_RUNTIME_LIMIT_EXCEEDED:
		return "RUNTIME_LIMIT_EXCEEDED"
	case BACnetLiftFault_POSITION_LOST:
		return "POSITION_LOST"
	case BACnetLiftFault_DRIVE_TEMPERATURE_EXCEEDED:
		return "DRIVE_TEMPERATURE_EXCEEDED"
	case BACnetLiftFault_LOAD_MEASUREMENT_FAULT:
		return "LOAD_MEASUREMENT_FAULT"
	case BACnetLiftFault_GOVERNOR_AND_SAFETY_GEAR_FAULT:
		return "GOVERNOR_AND_SAFETY_GEAR_FAULT"
	case BACnetLiftFault_LIFT_SHAFT_DEVICE_FAULT:
		return "LIFT_SHAFT_DEVICE_FAULT"
	case BACnetLiftFault_POWER_SUPPLY_FAULT:
		return "POWER_SUPPLY_FAULT"
	case BACnetLiftFault_SAFETY_INTERLOCK_FAULT:
		return "SAFETY_INTERLOCK_FAULT"
	case BACnetLiftFault_DOOR_CLOSING_FAULT:
		return "DOOR_CLOSING_FAULT"
	case BACnetLiftFault_DOOR_OPENING_FAULT:
		return "DOOR_OPENING_FAULT"
	case BACnetLiftFault_CAR_STOPPED_OUTSIDE_LANDING_ZONE:
		return "CAR_STOPPED_OUTSIDE_LANDING_ZONE"
	case BACnetLiftFault_CALL_BUTTON_STUCK:
		return "CALL_BUTTON_STUCK"
	}
	return ""
}

func (e BACnetLiftFault) String() string {
	return e.PLC4XEnumName()
}