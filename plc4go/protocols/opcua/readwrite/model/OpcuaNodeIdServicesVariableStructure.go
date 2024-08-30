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

// OpcuaNodeIdServicesVariableStructure is an enum
type OpcuaNodeIdServicesVariableStructure int32

type IOpcuaNodeIdServicesVariableStructure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableStructure_StructureType_EnumStrings OpcuaNodeIdServicesVariableStructure = 14528
)

var OpcuaNodeIdServicesVariableStructureValues []OpcuaNodeIdServicesVariableStructure

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableStructureValues = []OpcuaNodeIdServicesVariableStructure{
		OpcuaNodeIdServicesVariableStructure_StructureType_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableStructureByValue(value int32) (enum OpcuaNodeIdServicesVariableStructure, ok bool) {
	switch value {
	case 14528:
		return OpcuaNodeIdServicesVariableStructure_StructureType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableStructureByName(value string) (enum OpcuaNodeIdServicesVariableStructure, ok bool) {
	switch value {
	case "StructureType_EnumStrings":
		return OpcuaNodeIdServicesVariableStructure_StructureType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableStructureKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableStructureValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableStructure(structType any) OpcuaNodeIdServicesVariableStructure {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableStructure {
		if sOpcuaNodeIdServicesVariableStructure, ok := typ.(OpcuaNodeIdServicesVariableStructure); ok {
			return sOpcuaNodeIdServicesVariableStructure
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableStructure) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableStructure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableStructureParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableStructure, error) {
	return OpcuaNodeIdServicesVariableStructureParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableStructureParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableStructure, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableStructure", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableStructure")
	}
	if enum, ok := OpcuaNodeIdServicesVariableStructureByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableStructure")
		return OpcuaNodeIdServicesVariableStructure(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableStructure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableStructure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableStructure", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableStructure) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableStructure_StructureType_EnumStrings:
		return "StructureType_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableStructure) String() string {
	return e.PLC4XEnumName()
}
