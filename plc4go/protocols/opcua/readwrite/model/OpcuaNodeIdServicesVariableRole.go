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

// OpcuaNodeIdServicesVariableRole is an enum
type OpcuaNodeIdServicesVariableRole int32

type IOpcuaNodeIdServicesVariableRole interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_ApplicationsExclude              OpcuaNodeIdServicesVariableRole = 15408
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_EndpointsExclude                 OpcuaNodeIdServicesVariableRole = 15409
	OpcuaNodeIdServicesVariableRole_RoleType_ApplicationsExclude                                      OpcuaNodeIdServicesVariableRole = 15410
	OpcuaNodeIdServicesVariableRole_RoleType_EndpointsExclude                                         OpcuaNodeIdServicesVariableRole = 15411
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments       OpcuaNodeIdServicesVariableRole = 15613
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments    OpcuaNodeIdServicesVariableRole = 15615
	OpcuaNodeIdServicesVariableRole_RoleType_AddIdentity_InputArguments                               OpcuaNodeIdServicesVariableRole = 15625
	OpcuaNodeIdServicesVariableRole_RoleType_RemoveIdentity_InputArguments                            OpcuaNodeIdServicesVariableRole = 15627
	OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_InputArguments                                OpcuaNodeIdServicesVariableRole = 15998
	OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_OutputArguments                               OpcuaNodeIdServicesVariableRole = 15999
	OpcuaNodeIdServicesVariableRole_RoleSetType_RemoveRole_InputArguments                             OpcuaNodeIdServicesVariableRole = 16001
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Identities                       OpcuaNodeIdServicesVariableRole = 16162
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Applications                     OpcuaNodeIdServicesVariableRole = 16163
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Endpoints                        OpcuaNodeIdServicesVariableRole = 16164
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddApplication_InputArguments    OpcuaNodeIdServicesVariableRole = 16166
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments OpcuaNodeIdServicesVariableRole = 16168
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments       OpcuaNodeIdServicesVariableRole = 16170
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments    OpcuaNodeIdServicesVariableRole = 16172
	OpcuaNodeIdServicesVariableRole_RoleType_Identities                                               OpcuaNodeIdServicesVariableRole = 16173
	OpcuaNodeIdServicesVariableRole_RoleType_Applications                                             OpcuaNodeIdServicesVariableRole = 16174
	OpcuaNodeIdServicesVariableRole_RoleType_Endpoints                                                OpcuaNodeIdServicesVariableRole = 16175
	OpcuaNodeIdServicesVariableRole_RoleType_AddApplication_InputArguments                            OpcuaNodeIdServicesVariableRole = 16177
	OpcuaNodeIdServicesVariableRole_RoleType_RemoveApplication_InputArguments                         OpcuaNodeIdServicesVariableRole = 16179
	OpcuaNodeIdServicesVariableRole_RoleType_AddEndpoint_InputArguments                               OpcuaNodeIdServicesVariableRole = 16181
	OpcuaNodeIdServicesVariableRole_RoleType_RemoveEndpoint_InputArguments                            OpcuaNodeIdServicesVariableRole = 16183
	OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_CustomConfiguration              OpcuaNodeIdServicesVariableRole = 24138
	OpcuaNodeIdServicesVariableRole_RoleType_CustomConfiguration                                      OpcuaNodeIdServicesVariableRole = 24139
)

var OpcuaNodeIdServicesVariableRoleValues []OpcuaNodeIdServicesVariableRole

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableRoleValues = []OpcuaNodeIdServicesVariableRole{
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_ApplicationsExclude,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_EndpointsExclude,
		OpcuaNodeIdServicesVariableRole_RoleType_ApplicationsExclude,
		OpcuaNodeIdServicesVariableRole_RoleType_EndpointsExclude,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleType_AddIdentity_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleType_RemoveIdentity_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_OutputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RemoveRole_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Identities,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Applications,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Endpoints,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddApplication_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleType_Identities,
		OpcuaNodeIdServicesVariableRole_RoleType_Applications,
		OpcuaNodeIdServicesVariableRole_RoleType_Endpoints,
		OpcuaNodeIdServicesVariableRole_RoleType_AddApplication_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleType_RemoveApplication_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleType_AddEndpoint_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleType_RemoveEndpoint_InputArguments,
		OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_CustomConfiguration,
		OpcuaNodeIdServicesVariableRole_RoleType_CustomConfiguration,
	}
}

func OpcuaNodeIdServicesVariableRoleByValue(value int32) (enum OpcuaNodeIdServicesVariableRole, ok bool) {
	switch value {
	case 15408:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_ApplicationsExclude, true
	case 15409:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_EndpointsExclude, true
	case 15410:
		return OpcuaNodeIdServicesVariableRole_RoleType_ApplicationsExclude, true
	case 15411:
		return OpcuaNodeIdServicesVariableRole_RoleType_EndpointsExclude, true
	case 15613:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments, true
	case 15615:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments, true
	case 15625:
		return OpcuaNodeIdServicesVariableRole_RoleType_AddIdentity_InputArguments, true
	case 15627:
		return OpcuaNodeIdServicesVariableRole_RoleType_RemoveIdentity_InputArguments, true
	case 15998:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_InputArguments, true
	case 15999:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_OutputArguments, true
	case 16001:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RemoveRole_InputArguments, true
	case 16162:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Identities, true
	case 16163:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Applications, true
	case 16164:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Endpoints, true
	case 16166:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddApplication_InputArguments, true
	case 16168:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments, true
	case 16170:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments, true
	case 16172:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments, true
	case 16173:
		return OpcuaNodeIdServicesVariableRole_RoleType_Identities, true
	case 16174:
		return OpcuaNodeIdServicesVariableRole_RoleType_Applications, true
	case 16175:
		return OpcuaNodeIdServicesVariableRole_RoleType_Endpoints, true
	case 16177:
		return OpcuaNodeIdServicesVariableRole_RoleType_AddApplication_InputArguments, true
	case 16179:
		return OpcuaNodeIdServicesVariableRole_RoleType_RemoveApplication_InputArguments, true
	case 16181:
		return OpcuaNodeIdServicesVariableRole_RoleType_AddEndpoint_InputArguments, true
	case 16183:
		return OpcuaNodeIdServicesVariableRole_RoleType_RemoveEndpoint_InputArguments, true
	case 24138:
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_CustomConfiguration, true
	case 24139:
		return OpcuaNodeIdServicesVariableRole_RoleType_CustomConfiguration, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRoleByName(value string) (enum OpcuaNodeIdServicesVariableRole, ok bool) {
	switch value {
	case "RoleSetType_RoleName_Placeholder_ApplicationsExclude":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_ApplicationsExclude, true
	case "RoleSetType_RoleName_Placeholder_EndpointsExclude":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_EndpointsExclude, true
	case "RoleType_ApplicationsExclude":
		return OpcuaNodeIdServicesVariableRole_RoleType_ApplicationsExclude, true
	case "RoleType_EndpointsExclude":
		return OpcuaNodeIdServicesVariableRole_RoleType_EndpointsExclude, true
	case "RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments, true
	case "RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments, true
	case "RoleType_AddIdentity_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleType_AddIdentity_InputArguments, true
	case "RoleType_RemoveIdentity_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleType_RemoveIdentity_InputArguments, true
	case "RoleSetType_AddRole_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_InputArguments, true
	case "RoleSetType_AddRole_OutputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_OutputArguments, true
	case "RoleSetType_RemoveRole_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RemoveRole_InputArguments, true
	case "RoleSetType_RoleName_Placeholder_Identities":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Identities, true
	case "RoleSetType_RoleName_Placeholder_Applications":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Applications, true
	case "RoleSetType_RoleName_Placeholder_Endpoints":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Endpoints, true
	case "RoleSetType_RoleName_Placeholder_AddApplication_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddApplication_InputArguments, true
	case "RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments, true
	case "RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments, true
	case "RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments, true
	case "RoleType_Identities":
		return OpcuaNodeIdServicesVariableRole_RoleType_Identities, true
	case "RoleType_Applications":
		return OpcuaNodeIdServicesVariableRole_RoleType_Applications, true
	case "RoleType_Endpoints":
		return OpcuaNodeIdServicesVariableRole_RoleType_Endpoints, true
	case "RoleType_AddApplication_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleType_AddApplication_InputArguments, true
	case "RoleType_RemoveApplication_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleType_RemoveApplication_InputArguments, true
	case "RoleType_AddEndpoint_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleType_AddEndpoint_InputArguments, true
	case "RoleType_RemoveEndpoint_InputArguments":
		return OpcuaNodeIdServicesVariableRole_RoleType_RemoveEndpoint_InputArguments, true
	case "RoleSetType_RoleName_Placeholder_CustomConfiguration":
		return OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_CustomConfiguration, true
	case "RoleType_CustomConfiguration":
		return OpcuaNodeIdServicesVariableRole_RoleType_CustomConfiguration, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRoleKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableRoleValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableRole(structType any) OpcuaNodeIdServicesVariableRole {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableRole {
		if sOpcuaNodeIdServicesVariableRole, ok := typ.(OpcuaNodeIdServicesVariableRole); ok {
			return sOpcuaNodeIdServicesVariableRole
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableRole) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableRole) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableRoleParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableRole, error) {
	return OpcuaNodeIdServicesVariableRoleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableRoleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableRole, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableRole", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableRole")
	}
	if enum, ok := OpcuaNodeIdServicesVariableRoleByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableRole")
		return OpcuaNodeIdServicesVariableRole(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableRole) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableRole) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableRole", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableRole) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableRole) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_ApplicationsExclude:
		return "RoleSetType_RoleName_Placeholder_ApplicationsExclude"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_EndpointsExclude:
		return "RoleSetType_RoleName_Placeholder_EndpointsExclude"
	case OpcuaNodeIdServicesVariableRole_RoleType_ApplicationsExclude:
		return "RoleType_ApplicationsExclude"
	case OpcuaNodeIdServicesVariableRole_RoleType_EndpointsExclude:
		return "RoleType_EndpointsExclude"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments:
		return "RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments:
		return "RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleType_AddIdentity_InputArguments:
		return "RoleType_AddIdentity_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleType_RemoveIdentity_InputArguments:
		return "RoleType_RemoveIdentity_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_InputArguments:
		return "RoleSetType_AddRole_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_AddRole_OutputArguments:
		return "RoleSetType_AddRole_OutputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RemoveRole_InputArguments:
		return "RoleSetType_RemoveRole_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Identities:
		return "RoleSetType_RoleName_Placeholder_Identities"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Applications:
		return "RoleSetType_RoleName_Placeholder_Applications"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_Endpoints:
		return "RoleSetType_RoleName_Placeholder_Endpoints"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddApplication_InputArguments:
		return "RoleSetType_RoleName_Placeholder_AddApplication_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments:
		return "RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments:
		return "RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments:
		return "RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleType_Identities:
		return "RoleType_Identities"
	case OpcuaNodeIdServicesVariableRole_RoleType_Applications:
		return "RoleType_Applications"
	case OpcuaNodeIdServicesVariableRole_RoleType_Endpoints:
		return "RoleType_Endpoints"
	case OpcuaNodeIdServicesVariableRole_RoleType_AddApplication_InputArguments:
		return "RoleType_AddApplication_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleType_RemoveApplication_InputArguments:
		return "RoleType_RemoveApplication_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleType_AddEndpoint_InputArguments:
		return "RoleType_AddEndpoint_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleType_RemoveEndpoint_InputArguments:
		return "RoleType_RemoveEndpoint_InputArguments"
	case OpcuaNodeIdServicesVariableRole_RoleSetType_RoleName_Placeholder_CustomConfiguration:
		return "RoleSetType_RoleName_Placeholder_CustomConfiguration"
	case OpcuaNodeIdServicesVariableRole_RoleType_CustomConfiguration:
		return "RoleType_CustomConfiguration"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableRole) String() string {
	return e.PLC4XEnumName()
}
