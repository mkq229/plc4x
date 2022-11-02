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

// LightingCompatible is an enum
type LightingCompatible uint8

type ILightingCompatible interface {
	utils.Serializable
}

const (
	LightingCompatible_NO                   LightingCompatible = 0x0
	LightingCompatible_YES                  LightingCompatible = 0x1
	LightingCompatible_YES_BUT_RESTRICTIONS LightingCompatible = 0x2
	LightingCompatible_NA                   LightingCompatible = 0x3
)

var LightingCompatibleValues []LightingCompatible

func init() {
	_ = errors.New
	LightingCompatibleValues = []LightingCompatible{
		LightingCompatible_NO,
		LightingCompatible_YES,
		LightingCompatible_YES_BUT_RESTRICTIONS,
		LightingCompatible_NA,
	}
}

func LightingCompatibleByValue(value uint8) (enum LightingCompatible, ok bool) {
	switch value {
	case 0x0:
		return LightingCompatible_NO, true
	case 0x1:
		return LightingCompatible_YES, true
	case 0x2:
		return LightingCompatible_YES_BUT_RESTRICTIONS, true
	case 0x3:
		return LightingCompatible_NA, true
	}
	return 0, false
}

func LightingCompatibleByName(value string) (enum LightingCompatible, ok bool) {
	switch value {
	case "NO":
		return LightingCompatible_NO, true
	case "YES":
		return LightingCompatible_YES, true
	case "YES_BUT_RESTRICTIONS":
		return LightingCompatible_YES_BUT_RESTRICTIONS, true
	case "NA":
		return LightingCompatible_NA, true
	}
	return 0, false
}

func LightingCompatibleKnows(value uint8) bool {
	for _, typeValue := range LightingCompatibleValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLightingCompatible(structType interface{}) LightingCompatible {
	castFunc := func(typ interface{}) LightingCompatible {
		if sLightingCompatible, ok := typ.(LightingCompatible); ok {
			return sLightingCompatible
		}
		return 0
	}
	return castFunc(structType)
}

func (m LightingCompatible) GetLengthInBits() uint16 {
	return 4
}

func (m LightingCompatible) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LightingCompatibleParse(readBuffer utils.ReadBuffer) (LightingCompatible, error) {
	val, err := readBuffer.ReadUint8("LightingCompatible", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LightingCompatible")
	}
	if enum, ok := LightingCompatibleByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return LightingCompatible(val), nil
	} else {
		return enum, nil
	}
}

func (e LightingCompatible) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LightingCompatible) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("LightingCompatible", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LightingCompatible) PLC4XEnumName() string {
	switch e {
	case LightingCompatible_NO:
		return "NO"
	case LightingCompatible_YES:
		return "YES"
	case LightingCompatible_YES_BUT_RESTRICTIONS:
		return "YES_BUT_RESTRICTIONS"
	case LightingCompatible_NA:
		return "NA"
	}
	return ""
}

func (e LightingCompatible) String() string {
	return e.PLC4XEnumName()
}