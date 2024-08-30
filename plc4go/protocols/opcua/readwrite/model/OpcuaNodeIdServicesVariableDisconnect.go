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

// OpcuaNodeIdServicesVariableDisconnect is an enum
type OpcuaNodeIdServicesVariableDisconnect int32

type IOpcuaNodeIdServicesVariableDisconnect interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_InputArguments  OpcuaNodeIdServicesVariableDisconnect = 25335
	OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_OutputArguments OpcuaNodeIdServicesVariableDisconnect = 25336
)

var OpcuaNodeIdServicesVariableDisconnectValues []OpcuaNodeIdServicesVariableDisconnect

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableDisconnectValues = []OpcuaNodeIdServicesVariableDisconnect{
		OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_InputArguments,
		OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableDisconnectByValue(value int32) (enum OpcuaNodeIdServicesVariableDisconnect, ok bool) {
	switch value {
	case 25335:
		return OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_InputArguments, true
	case 25336:
		return OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDisconnectByName(value string) (enum OpcuaNodeIdServicesVariableDisconnect, ok bool) {
	switch value {
	case "DisconnectSecurityGroupsMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_InputArguments, true
	case "DisconnectSecurityGroupsMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDisconnectKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableDisconnectValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableDisconnect(structType any) OpcuaNodeIdServicesVariableDisconnect {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableDisconnect {
		if sOpcuaNodeIdServicesVariableDisconnect, ok := typ.(OpcuaNodeIdServicesVariableDisconnect); ok {
			return sOpcuaNodeIdServicesVariableDisconnect
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableDisconnect) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableDisconnect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableDisconnectParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableDisconnect, error) {
	return OpcuaNodeIdServicesVariableDisconnectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableDisconnectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableDisconnect, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableDisconnect", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableDisconnect")
	}
	if enum, ok := OpcuaNodeIdServicesVariableDisconnectByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableDisconnect")
		return OpcuaNodeIdServicesVariableDisconnect(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableDisconnect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableDisconnect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableDisconnect", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableDisconnect) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_InputArguments:
		return "DisconnectSecurityGroupsMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableDisconnect_DisconnectSecurityGroupsMethodType_OutputArguments:
		return "DisconnectSecurityGroupsMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableDisconnect) String() string {
	return e.PLC4XEnumName()
}
