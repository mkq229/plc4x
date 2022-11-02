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

// BACnetLightingInProgress is an enum
type BACnetLightingInProgress uint8

type IBACnetLightingInProgress interface {
	utils.Serializable
}

const (
	BACnetLightingInProgress_IDLE           BACnetLightingInProgress = 0
	BACnetLightingInProgress_FADE_ACTIVE    BACnetLightingInProgress = 1
	BACnetLightingInProgress_RAMP_ACTIVE    BACnetLightingInProgress = 2
	BACnetLightingInProgress_NOT_CONTROLLED BACnetLightingInProgress = 3
	BACnetLightingInProgress_OTHER          BACnetLightingInProgress = 4
)

var BACnetLightingInProgressValues []BACnetLightingInProgress

func init() {
	_ = errors.New
	BACnetLightingInProgressValues = []BACnetLightingInProgress{
		BACnetLightingInProgress_IDLE,
		BACnetLightingInProgress_FADE_ACTIVE,
		BACnetLightingInProgress_RAMP_ACTIVE,
		BACnetLightingInProgress_NOT_CONTROLLED,
		BACnetLightingInProgress_OTHER,
	}
}

func BACnetLightingInProgressByValue(value uint8) (enum BACnetLightingInProgress, ok bool) {
	switch value {
	case 0:
		return BACnetLightingInProgress_IDLE, true
	case 1:
		return BACnetLightingInProgress_FADE_ACTIVE, true
	case 2:
		return BACnetLightingInProgress_RAMP_ACTIVE, true
	case 3:
		return BACnetLightingInProgress_NOT_CONTROLLED, true
	case 4:
		return BACnetLightingInProgress_OTHER, true
	}
	return 0, false
}

func BACnetLightingInProgressByName(value string) (enum BACnetLightingInProgress, ok bool) {
	switch value {
	case "IDLE":
		return BACnetLightingInProgress_IDLE, true
	case "FADE_ACTIVE":
		return BACnetLightingInProgress_FADE_ACTIVE, true
	case "RAMP_ACTIVE":
		return BACnetLightingInProgress_RAMP_ACTIVE, true
	case "NOT_CONTROLLED":
		return BACnetLightingInProgress_NOT_CONTROLLED, true
	case "OTHER":
		return BACnetLightingInProgress_OTHER, true
	}
	return 0, false
}

func BACnetLightingInProgressKnows(value uint8) bool {
	for _, typeValue := range BACnetLightingInProgressValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLightingInProgress(structType interface{}) BACnetLightingInProgress {
	castFunc := func(typ interface{}) BACnetLightingInProgress {
		if sBACnetLightingInProgress, ok := typ.(BACnetLightingInProgress); ok {
			return sBACnetLightingInProgress
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLightingInProgress) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetLightingInProgress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLightingInProgressParse(readBuffer utils.ReadBuffer) (BACnetLightingInProgress, error) {
	val, err := readBuffer.ReadUint8("BACnetLightingInProgress", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLightingInProgress")
	}
	if enum, ok := BACnetLightingInProgressByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLightingInProgress(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLightingInProgress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLightingInProgress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetLightingInProgress", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLightingInProgress) PLC4XEnumName() string {
	switch e {
	case BACnetLightingInProgress_IDLE:
		return "IDLE"
	case BACnetLightingInProgress_FADE_ACTIVE:
		return "FADE_ACTIVE"
	case BACnetLightingInProgress_RAMP_ACTIVE:
		return "RAMP_ACTIVE"
	case BACnetLightingInProgress_NOT_CONTROLLED:
		return "NOT_CONTROLLED"
	case BACnetLightingInProgress_OTHER:
		return "OTHER"
	}
	return ""
}

func (e BACnetLightingInProgress) String() string {
	return e.PLC4XEnumName()
}