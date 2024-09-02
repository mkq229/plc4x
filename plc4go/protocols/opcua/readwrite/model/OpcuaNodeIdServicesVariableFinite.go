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

// OpcuaNodeIdServicesVariableFinite is an enum
type OpcuaNodeIdServicesVariableFinite int32

type IOpcuaNodeIdServicesVariableFinite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableFinite_FiniteStateVariableType_Id                                    OpcuaNodeIdServicesVariableFinite = 2761
	OpcuaNodeIdServicesVariableFinite_FiniteTransitionVariableType_Id                               OpcuaNodeIdServicesVariableFinite = 2768
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState                           OpcuaNodeIdServicesVariableFinite = 2772
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition                         OpcuaNodeIdServicesVariableFinite = 2773
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Id                        OpcuaNodeIdServicesVariableFinite = 3728
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Name                      OpcuaNodeIdServicesVariableFinite = 3729
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Number                    OpcuaNodeIdServicesVariableFinite = 3730
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_EffectiveDisplayName      OpcuaNodeIdServicesVariableFinite = 3731
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Id                      OpcuaNodeIdServicesVariableFinite = 3732
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Name                    OpcuaNodeIdServicesVariableFinite = 3733
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Number                  OpcuaNodeIdServicesVariableFinite = 3734
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_TransitionTime          OpcuaNodeIdServicesVariableFinite = 3735
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_EffectiveTransitionTime OpcuaNodeIdServicesVariableFinite = 11459
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableStates                        OpcuaNodeIdServicesVariableFinite = 17635
	OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableTransitions                   OpcuaNodeIdServicesVariableFinite = 17636
)

var OpcuaNodeIdServicesVariableFiniteValues []OpcuaNodeIdServicesVariableFinite

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableFiniteValues = []OpcuaNodeIdServicesVariableFinite{
		OpcuaNodeIdServicesVariableFinite_FiniteStateVariableType_Id,
		OpcuaNodeIdServicesVariableFinite_FiniteTransitionVariableType_Id,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Id,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Name,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Number,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Id,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Name,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Number,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_TransitionTime,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableStates,
		OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableTransitions,
	}
}

func OpcuaNodeIdServicesVariableFiniteByValue(value int32) (enum OpcuaNodeIdServicesVariableFinite, ok bool) {
	switch value {
	case 11459:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_EffectiveTransitionTime, true
	case 17635:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableStates, true
	case 17636:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableTransitions, true
	case 2761:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateVariableType_Id, true
	case 2768:
		return OpcuaNodeIdServicesVariableFinite_FiniteTransitionVariableType_Id, true
	case 2772:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState, true
	case 2773:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition, true
	case 3728:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Id, true
	case 3729:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Name, true
	case 3730:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Number, true
	case 3731:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_EffectiveDisplayName, true
	case 3732:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Id, true
	case 3733:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Name, true
	case 3734:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Number, true
	case 3735:
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_TransitionTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableFiniteByName(value string) (enum OpcuaNodeIdServicesVariableFinite, ok bool) {
	switch value {
	case "FiniteStateMachineType_LastTransition_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_EffectiveTransitionTime, true
	case "FiniteStateMachineType_AvailableStates":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableStates, true
	case "FiniteStateMachineType_AvailableTransitions":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableTransitions, true
	case "FiniteStateVariableType_Id":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateVariableType_Id, true
	case "FiniteTransitionVariableType_Id":
		return OpcuaNodeIdServicesVariableFinite_FiniteTransitionVariableType_Id, true
	case "FiniteStateMachineType_CurrentState":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState, true
	case "FiniteStateMachineType_LastTransition":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition, true
	case "FiniteStateMachineType_CurrentState_Id":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Id, true
	case "FiniteStateMachineType_CurrentState_Name":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Name, true
	case "FiniteStateMachineType_CurrentState_Number":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Number, true
	case "FiniteStateMachineType_CurrentState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_EffectiveDisplayName, true
	case "FiniteStateMachineType_LastTransition_Id":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Id, true
	case "FiniteStateMachineType_LastTransition_Name":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Name, true
	case "FiniteStateMachineType_LastTransition_Number":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Number, true
	case "FiniteStateMachineType_LastTransition_TransitionTime":
		return OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_TransitionTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableFiniteKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableFiniteValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableFinite(structType any) OpcuaNodeIdServicesVariableFinite {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableFinite {
		if sOpcuaNodeIdServicesVariableFinite, ok := typ.(OpcuaNodeIdServicesVariableFinite); ok {
			return sOpcuaNodeIdServicesVariableFinite
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableFinite) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableFinite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableFiniteParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableFinite, error) {
	return OpcuaNodeIdServicesVariableFiniteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableFiniteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableFinite, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableFinite", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableFinite")
	}
	if enum, ok := OpcuaNodeIdServicesVariableFiniteByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableFinite")
		return OpcuaNodeIdServicesVariableFinite(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableFinite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableFinite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableFinite", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableFinite) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableFinite) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_EffectiveTransitionTime:
		return "FiniteStateMachineType_LastTransition_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableStates:
		return "FiniteStateMachineType_AvailableStates"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_AvailableTransitions:
		return "FiniteStateMachineType_AvailableTransitions"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateVariableType_Id:
		return "FiniteStateVariableType_Id"
	case OpcuaNodeIdServicesVariableFinite_FiniteTransitionVariableType_Id:
		return "FiniteTransitionVariableType_Id"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState:
		return "FiniteStateMachineType_CurrentState"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition:
		return "FiniteStateMachineType_LastTransition"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Id:
		return "FiniteStateMachineType_CurrentState_Id"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Name:
		return "FiniteStateMachineType_CurrentState_Name"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_Number:
		return "FiniteStateMachineType_CurrentState_Number"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_CurrentState_EffectiveDisplayName:
		return "FiniteStateMachineType_CurrentState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Id:
		return "FiniteStateMachineType_LastTransition_Id"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Name:
		return "FiniteStateMachineType_LastTransition_Name"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_Number:
		return "FiniteStateMachineType_LastTransition_Number"
	case OpcuaNodeIdServicesVariableFinite_FiniteStateMachineType_LastTransition_TransitionTime:
		return "FiniteStateMachineType_LastTransition_TransitionTime"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableFinite) String() string {
	return e.PLC4XEnumName()
}
