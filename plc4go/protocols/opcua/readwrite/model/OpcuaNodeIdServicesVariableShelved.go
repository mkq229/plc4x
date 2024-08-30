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

// OpcuaNodeIdServicesVariableShelved is an enum
type OpcuaNodeIdServicesVariableShelved int32

type IOpcuaNodeIdServicesVariableShelved interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve_InputArguments                    OpcuaNodeIdServicesVariableShelved = 2991
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelved_StateNumber                         OpcuaNodeIdServicesVariableShelved = 6098
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelved_StateNumber                      OpcuaNodeIdServicesVariableShelved = 6100
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelved_StateNumber                    OpcuaNodeIdServicesVariableShelved = 6101
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelveTime                                  OpcuaNodeIdServicesVariableShelved = 9115
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber      OpcuaNodeIdServicesVariableShelved = 11322
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber    OpcuaNodeIdServicesVariableShelved = 11323
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber      OpcuaNodeIdServicesVariableShelved = 11324
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber OpcuaNodeIdServicesVariableShelved = 11325
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber    OpcuaNodeIdServicesVariableShelved = 11326
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber OpcuaNodeIdServicesVariableShelved = 11327
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve2_InputArguments                   OpcuaNodeIdServicesVariableShelved = 24757
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelve2_InputArguments                      OpcuaNodeIdServicesVariableShelved = 24759
	OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelve2_InputArguments                 OpcuaNodeIdServicesVariableShelved = 24761
)

var OpcuaNodeIdServicesVariableShelvedValues []OpcuaNodeIdServicesVariableShelved

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableShelvedValues = []OpcuaNodeIdServicesVariableShelved{
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve_InputArguments,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelved_StateNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelved_StateNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelved_StateNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelveTime,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve2_InputArguments,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelve2_InputArguments,
		OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelve2_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableShelvedByValue(value int32) (enum OpcuaNodeIdServicesVariableShelved, ok bool) {
	switch value {
	case 11322:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber, true
	case 11323:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber, true
	case 11324:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber, true
	case 11325:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber, true
	case 11326:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber, true
	case 11327:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber, true
	case 24757:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve2_InputArguments, true
	case 24759:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelve2_InputArguments, true
	case 24761:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelve2_InputArguments, true
	case 2991:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve_InputArguments, true
	case 6098:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelved_StateNumber, true
	case 6100:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelved_StateNumber, true
	case 6101:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelved_StateNumber, true
	case 9115:
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelveTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableShelvedByName(value string) (enum OpcuaNodeIdServicesVariableShelved, ok bool) {
	switch value {
	case "ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber, true
	case "ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber, true
	case "ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber, true
	case "ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber, true
	case "ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber, true
	case "ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber, true
	case "ShelvedStateMachineType_TimedShelve2_InputArguments":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve2_InputArguments, true
	case "ShelvedStateMachineType_Unshelve2_InputArguments":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelve2_InputArguments, true
	case "ShelvedStateMachineType_OneShotShelve2_InputArguments":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelve2_InputArguments, true
	case "ShelvedStateMachineType_TimedShelve_InputArguments":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve_InputArguments, true
	case "ShelvedStateMachineType_Unshelved_StateNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelved_StateNumber, true
	case "ShelvedStateMachineType_TimedShelved_StateNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelved_StateNumber, true
	case "ShelvedStateMachineType_OneShotShelved_StateNumber":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelved_StateNumber, true
	case "ShelvedStateMachineType_UnshelveTime":
		return OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelveTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableShelvedKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableShelvedValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableShelved(structType any) OpcuaNodeIdServicesVariableShelved {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableShelved {
		if sOpcuaNodeIdServicesVariableShelved, ok := typ.(OpcuaNodeIdServicesVariableShelved); ok {
			return sOpcuaNodeIdServicesVariableShelved
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableShelved) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableShelved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableShelvedParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableShelved, error) {
	return OpcuaNodeIdServicesVariableShelvedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableShelvedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableShelved, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableShelved", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableShelved")
	}
	if enum, ok := OpcuaNodeIdServicesVariableShelvedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableShelved")
		return OpcuaNodeIdServicesVariableShelved(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableShelved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableShelved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableShelved", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableShelved) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber:
		return "ShelvedStateMachineType_UnshelvedToTimedShelved_TransitionNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber:
		return "ShelvedStateMachineType_UnshelvedToOneShotShelved_TransitionNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber:
		return "ShelvedStateMachineType_TimedShelvedToUnshelved_TransitionNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber:
		return "ShelvedStateMachineType_TimedShelvedToOneShotShelved_TransitionNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber:
		return "ShelvedStateMachineType_OneShotShelvedToUnshelved_TransitionNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber:
		return "ShelvedStateMachineType_OneShotShelvedToTimedShelved_TransitionNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve2_InputArguments:
		return "ShelvedStateMachineType_TimedShelve2_InputArguments"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelve2_InputArguments:
		return "ShelvedStateMachineType_Unshelve2_InputArguments"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelve2_InputArguments:
		return "ShelvedStateMachineType_OneShotShelve2_InputArguments"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelve_InputArguments:
		return "ShelvedStateMachineType_TimedShelve_InputArguments"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_Unshelved_StateNumber:
		return "ShelvedStateMachineType_Unshelved_StateNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_TimedShelved_StateNumber:
		return "ShelvedStateMachineType_TimedShelved_StateNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_OneShotShelved_StateNumber:
		return "ShelvedStateMachineType_OneShotShelved_StateNumber"
	case OpcuaNodeIdServicesVariableShelved_ShelvedStateMachineType_UnshelveTime:
		return "ShelvedStateMachineType_UnshelveTime"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableShelved) String() string {
	return e.PLC4XEnumName()
}
