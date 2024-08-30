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

// OpcuaNodeIdServicesVariableBuild is an enum
type OpcuaNodeIdServicesVariableBuild int32

type IOpcuaNodeIdServicesVariableBuild interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductUri       OpcuaNodeIdServicesVariableBuild = 3052
	OpcuaNodeIdServicesVariableBuild_BuildInfoType_ManufacturerName OpcuaNodeIdServicesVariableBuild = 3053
	OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductName      OpcuaNodeIdServicesVariableBuild = 3054
	OpcuaNodeIdServicesVariableBuild_BuildInfoType_SoftwareVersion  OpcuaNodeIdServicesVariableBuild = 3055
	OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildNumber      OpcuaNodeIdServicesVariableBuild = 3056
	OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildDate        OpcuaNodeIdServicesVariableBuild = 3057
)

var OpcuaNodeIdServicesVariableBuildValues []OpcuaNodeIdServicesVariableBuild

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableBuildValues = []OpcuaNodeIdServicesVariableBuild{
		OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductUri,
		OpcuaNodeIdServicesVariableBuild_BuildInfoType_ManufacturerName,
		OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductName,
		OpcuaNodeIdServicesVariableBuild_BuildInfoType_SoftwareVersion,
		OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildNumber,
		OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildDate,
	}
}

func OpcuaNodeIdServicesVariableBuildByValue(value int32) (enum OpcuaNodeIdServicesVariableBuild, ok bool) {
	switch value {
	case 3052:
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductUri, true
	case 3053:
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_ManufacturerName, true
	case 3054:
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductName, true
	case 3055:
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_SoftwareVersion, true
	case 3056:
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildNumber, true
	case 3057:
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildDate, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBuildByName(value string) (enum OpcuaNodeIdServicesVariableBuild, ok bool) {
	switch value {
	case "BuildInfoType_ProductUri":
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductUri, true
	case "BuildInfoType_ManufacturerName":
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_ManufacturerName, true
	case "BuildInfoType_ProductName":
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductName, true
	case "BuildInfoType_SoftwareVersion":
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_SoftwareVersion, true
	case "BuildInfoType_BuildNumber":
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildNumber, true
	case "BuildInfoType_BuildDate":
		return OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildDate, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBuildKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableBuildValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableBuild(structType any) OpcuaNodeIdServicesVariableBuild {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableBuild {
		if sOpcuaNodeIdServicesVariableBuild, ok := typ.(OpcuaNodeIdServicesVariableBuild); ok {
			return sOpcuaNodeIdServicesVariableBuild
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableBuild) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableBuild) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableBuildParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableBuild, error) {
	return OpcuaNodeIdServicesVariableBuildParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableBuildParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableBuild, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableBuild", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableBuild")
	}
	if enum, ok := OpcuaNodeIdServicesVariableBuildByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableBuild")
		return OpcuaNodeIdServicesVariableBuild(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableBuild) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableBuild) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableBuild", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableBuild) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductUri:
		return "BuildInfoType_ProductUri"
	case OpcuaNodeIdServicesVariableBuild_BuildInfoType_ManufacturerName:
		return "BuildInfoType_ManufacturerName"
	case OpcuaNodeIdServicesVariableBuild_BuildInfoType_ProductName:
		return "BuildInfoType_ProductName"
	case OpcuaNodeIdServicesVariableBuild_BuildInfoType_SoftwareVersion:
		return "BuildInfoType_SoftwareVersion"
	case OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildNumber:
		return "BuildInfoType_BuildNumber"
	case OpcuaNodeIdServicesVariableBuild_BuildInfoType_BuildDate:
		return "BuildInfoType_BuildDate"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableBuild) String() string {
	return e.PLC4XEnumName()
}
