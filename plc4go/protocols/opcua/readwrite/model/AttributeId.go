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

// AttributeId is an enum
type AttributeId uint32

type IAttributeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	AttributeId_NodeId                  AttributeId = 1
	AttributeId_NodeClass               AttributeId = 2
	AttributeId_BrowseName              AttributeId = 3
	AttributeId_DisplayName             AttributeId = 4
	AttributeId_Description             AttributeId = 5
	AttributeId_WriteMask               AttributeId = 6
	AttributeId_UserWriteMask           AttributeId = 7
	AttributeId_IsAbstract              AttributeId = 8
	AttributeId_Symmetric               AttributeId = 9
	AttributeId_InverseName             AttributeId = 10
	AttributeId_ContainsNoLoops         AttributeId = 11
	AttributeId_EventNotifier           AttributeId = 12
	AttributeId_Value                   AttributeId = 13
	AttributeId_DataType                AttributeId = 14
	AttributeId_ValueRank               AttributeId = 15
	AttributeId_ArrayDimensions         AttributeId = 16
	AttributeId_AccessLevel             AttributeId = 17
	AttributeId_UserAccessLevel         AttributeId = 18
	AttributeId_MinimumSamplingInterval AttributeId = 19
	AttributeId_Historizing             AttributeId = 20
	AttributeId_Executable              AttributeId = 21
	AttributeId_UserExecutable          AttributeId = 22
	AttributeId_DataTypeDefinition      AttributeId = 23
	AttributeId_RolePermissions         AttributeId = 24
	AttributeId_UserRolePermissions     AttributeId = 25
	AttributeId_AccessRestrictions      AttributeId = 26
	AttributeId_AccessLevelEx           AttributeId = 27
)

var AttributeIdValues []AttributeId

func init() {
	_ = errors.New
	AttributeIdValues = []AttributeId{
		AttributeId_NodeId,
		AttributeId_NodeClass,
		AttributeId_BrowseName,
		AttributeId_DisplayName,
		AttributeId_Description,
		AttributeId_WriteMask,
		AttributeId_UserWriteMask,
		AttributeId_IsAbstract,
		AttributeId_Symmetric,
		AttributeId_InverseName,
		AttributeId_ContainsNoLoops,
		AttributeId_EventNotifier,
		AttributeId_Value,
		AttributeId_DataType,
		AttributeId_ValueRank,
		AttributeId_ArrayDimensions,
		AttributeId_AccessLevel,
		AttributeId_UserAccessLevel,
		AttributeId_MinimumSamplingInterval,
		AttributeId_Historizing,
		AttributeId_Executable,
		AttributeId_UserExecutable,
		AttributeId_DataTypeDefinition,
		AttributeId_RolePermissions,
		AttributeId_UserRolePermissions,
		AttributeId_AccessRestrictions,
		AttributeId_AccessLevelEx,
	}
}

func AttributeIdByValue(value uint32) (enum AttributeId, ok bool) {
	switch value {
	case 1:
		return AttributeId_NodeId, true
	case 10:
		return AttributeId_InverseName, true
	case 11:
		return AttributeId_ContainsNoLoops, true
	case 12:
		return AttributeId_EventNotifier, true
	case 13:
		return AttributeId_Value, true
	case 14:
		return AttributeId_DataType, true
	case 15:
		return AttributeId_ValueRank, true
	case 16:
		return AttributeId_ArrayDimensions, true
	case 17:
		return AttributeId_AccessLevel, true
	case 18:
		return AttributeId_UserAccessLevel, true
	case 19:
		return AttributeId_MinimumSamplingInterval, true
	case 2:
		return AttributeId_NodeClass, true
	case 20:
		return AttributeId_Historizing, true
	case 21:
		return AttributeId_Executable, true
	case 22:
		return AttributeId_UserExecutable, true
	case 23:
		return AttributeId_DataTypeDefinition, true
	case 24:
		return AttributeId_RolePermissions, true
	case 25:
		return AttributeId_UserRolePermissions, true
	case 26:
		return AttributeId_AccessRestrictions, true
	case 27:
		return AttributeId_AccessLevelEx, true
	case 3:
		return AttributeId_BrowseName, true
	case 4:
		return AttributeId_DisplayName, true
	case 5:
		return AttributeId_Description, true
	case 6:
		return AttributeId_WriteMask, true
	case 7:
		return AttributeId_UserWriteMask, true
	case 8:
		return AttributeId_IsAbstract, true
	case 9:
		return AttributeId_Symmetric, true
	}
	return 0, false
}

func AttributeIdByName(value string) (enum AttributeId, ok bool) {
	switch value {
	case "NodeId":
		return AttributeId_NodeId, true
	case "InverseName":
		return AttributeId_InverseName, true
	case "ContainsNoLoops":
		return AttributeId_ContainsNoLoops, true
	case "EventNotifier":
		return AttributeId_EventNotifier, true
	case "Value":
		return AttributeId_Value, true
	case "DataType":
		return AttributeId_DataType, true
	case "ValueRank":
		return AttributeId_ValueRank, true
	case "ArrayDimensions":
		return AttributeId_ArrayDimensions, true
	case "AccessLevel":
		return AttributeId_AccessLevel, true
	case "UserAccessLevel":
		return AttributeId_UserAccessLevel, true
	case "MinimumSamplingInterval":
		return AttributeId_MinimumSamplingInterval, true
	case "NodeClass":
		return AttributeId_NodeClass, true
	case "Historizing":
		return AttributeId_Historizing, true
	case "Executable":
		return AttributeId_Executable, true
	case "UserExecutable":
		return AttributeId_UserExecutable, true
	case "DataTypeDefinition":
		return AttributeId_DataTypeDefinition, true
	case "RolePermissions":
		return AttributeId_RolePermissions, true
	case "UserRolePermissions":
		return AttributeId_UserRolePermissions, true
	case "AccessRestrictions":
		return AttributeId_AccessRestrictions, true
	case "AccessLevelEx":
		return AttributeId_AccessLevelEx, true
	case "BrowseName":
		return AttributeId_BrowseName, true
	case "DisplayName":
		return AttributeId_DisplayName, true
	case "Description":
		return AttributeId_Description, true
	case "WriteMask":
		return AttributeId_WriteMask, true
	case "UserWriteMask":
		return AttributeId_UserWriteMask, true
	case "IsAbstract":
		return AttributeId_IsAbstract, true
	case "Symmetric":
		return AttributeId_Symmetric, true
	}
	return 0, false
}

func AttributeIdKnows(value uint32) bool {
	for _, typeValue := range AttributeIdValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAttributeId(structType any) AttributeId {
	castFunc := func(typ any) AttributeId {
		if sAttributeId, ok := typ.(AttributeId); ok {
			return sAttributeId
		}
		return 0
	}
	return castFunc(structType)
}

func (m AttributeId) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m AttributeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AttributeIdParse(ctx context.Context, theBytes []byte) (AttributeId, error) {
	return AttributeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AttributeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AttributeId, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("AttributeId", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AttributeId")
	}
	if enum, ok := AttributeIdByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AttributeId")
		return AttributeId(val), nil
	} else {
		return enum, nil
	}
}

func (e AttributeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AttributeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("AttributeId", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e AttributeId) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AttributeId) PLC4XEnumName() string {
	switch e {
	case AttributeId_NodeId:
		return "NodeId"
	case AttributeId_InverseName:
		return "InverseName"
	case AttributeId_ContainsNoLoops:
		return "ContainsNoLoops"
	case AttributeId_EventNotifier:
		return "EventNotifier"
	case AttributeId_Value:
		return "Value"
	case AttributeId_DataType:
		return "DataType"
	case AttributeId_ValueRank:
		return "ValueRank"
	case AttributeId_ArrayDimensions:
		return "ArrayDimensions"
	case AttributeId_AccessLevel:
		return "AccessLevel"
	case AttributeId_UserAccessLevel:
		return "UserAccessLevel"
	case AttributeId_MinimumSamplingInterval:
		return "MinimumSamplingInterval"
	case AttributeId_NodeClass:
		return "NodeClass"
	case AttributeId_Historizing:
		return "Historizing"
	case AttributeId_Executable:
		return "Executable"
	case AttributeId_UserExecutable:
		return "UserExecutable"
	case AttributeId_DataTypeDefinition:
		return "DataTypeDefinition"
	case AttributeId_RolePermissions:
		return "RolePermissions"
	case AttributeId_UserRolePermissions:
		return "UserRolePermissions"
	case AttributeId_AccessRestrictions:
		return "AccessRestrictions"
	case AttributeId_AccessLevelEx:
		return "AccessLevelEx"
	case AttributeId_BrowseName:
		return "BrowseName"
	case AttributeId_DisplayName:
		return "DisplayName"
	case AttributeId_Description:
		return "Description"
	case AttributeId_WriteMask:
		return "WriteMask"
	case AttributeId_UserWriteMask:
		return "UserWriteMask"
	case AttributeId_IsAbstract:
		return "IsAbstract"
	case AttributeId_Symmetric:
		return "Symmetric"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e AttributeId) String() string {
	return e.PLC4XEnumName()
}