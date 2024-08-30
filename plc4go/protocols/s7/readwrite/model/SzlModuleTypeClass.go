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

// SzlModuleTypeClass is an enum
type SzlModuleTypeClass uint8

type ISzlModuleTypeClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	SzlModuleTypeClass_CPU SzlModuleTypeClass = 0x0
	SzlModuleTypeClass_IM  SzlModuleTypeClass = 0x4
	SzlModuleTypeClass_FM  SzlModuleTypeClass = 0x8
	SzlModuleTypeClass_CP  SzlModuleTypeClass = 0xC
)

var SzlModuleTypeClassValues []SzlModuleTypeClass

func init() {
	_ = errors.New
	SzlModuleTypeClassValues = []SzlModuleTypeClass{
		SzlModuleTypeClass_CPU,
		SzlModuleTypeClass_IM,
		SzlModuleTypeClass_FM,
		SzlModuleTypeClass_CP,
	}
}

func SzlModuleTypeClassByValue(value uint8) (enum SzlModuleTypeClass, ok bool) {
	switch value {
	case 0x0:
		return SzlModuleTypeClass_CPU, true
	case 0x4:
		return SzlModuleTypeClass_IM, true
	case 0x8:
		return SzlModuleTypeClass_FM, true
	case 0xC:
		return SzlModuleTypeClass_CP, true
	}
	return 0, false
}

func SzlModuleTypeClassByName(value string) (enum SzlModuleTypeClass, ok bool) {
	switch value {
	case "CPU":
		return SzlModuleTypeClass_CPU, true
	case "IM":
		return SzlModuleTypeClass_IM, true
	case "FM":
		return SzlModuleTypeClass_FM, true
	case "CP":
		return SzlModuleTypeClass_CP, true
	}
	return 0, false
}

func SzlModuleTypeClassKnows(value uint8) bool {
	for _, typeValue := range SzlModuleTypeClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSzlModuleTypeClass(structType any) SzlModuleTypeClass {
	castFunc := func(typ any) SzlModuleTypeClass {
		if sSzlModuleTypeClass, ok := typ.(SzlModuleTypeClass); ok {
			return sSzlModuleTypeClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m SzlModuleTypeClass) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m SzlModuleTypeClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SzlModuleTypeClassParse(ctx context.Context, theBytes []byte) (SzlModuleTypeClass, error) {
	return SzlModuleTypeClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SzlModuleTypeClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SzlModuleTypeClass, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("SzlModuleTypeClass", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SzlModuleTypeClass")
	}
	if enum, ok := SzlModuleTypeClassByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for SzlModuleTypeClass")
		return SzlModuleTypeClass(val), nil
	} else {
		return enum, nil
	}
}

func (e SzlModuleTypeClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e SzlModuleTypeClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("SzlModuleTypeClass", 4, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SzlModuleTypeClass) PLC4XEnumName() string {
	switch e {
	case SzlModuleTypeClass_CPU:
		return "CPU"
	case SzlModuleTypeClass_IM:
		return "IM"
	case SzlModuleTypeClass_FM:
		return "FM"
	case SzlModuleTypeClass_CP:
		return "CP"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e SzlModuleTypeClass) String() string {
	return e.PLC4XEnumName()
}
