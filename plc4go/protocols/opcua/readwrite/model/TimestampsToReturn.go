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

// TimestampsToReturn is an enum
type TimestampsToReturn uint32

type ITimestampsToReturn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TimestampsToReturn_timestampsToReturnSource  TimestampsToReturn = 0
	TimestampsToReturn_timestampsToReturnServer  TimestampsToReturn = 1
	TimestampsToReturn_timestampsToReturnBoth    TimestampsToReturn = 2
	TimestampsToReturn_timestampsToReturnNeither TimestampsToReturn = 3
	TimestampsToReturn_timestampsToReturnInvalid TimestampsToReturn = 4
)

var TimestampsToReturnValues []TimestampsToReturn

func init() {
	_ = errors.New
	TimestampsToReturnValues = []TimestampsToReturn{
		TimestampsToReturn_timestampsToReturnSource,
		TimestampsToReturn_timestampsToReturnServer,
		TimestampsToReturn_timestampsToReturnBoth,
		TimestampsToReturn_timestampsToReturnNeither,
		TimestampsToReturn_timestampsToReturnInvalid,
	}
}

func TimestampsToReturnByValue(value uint32) (enum TimestampsToReturn, ok bool) {
	switch value {
	case 0:
		return TimestampsToReturn_timestampsToReturnSource, true
	case 1:
		return TimestampsToReturn_timestampsToReturnServer, true
	case 2:
		return TimestampsToReturn_timestampsToReturnBoth, true
	case 3:
		return TimestampsToReturn_timestampsToReturnNeither, true
	case 4:
		return TimestampsToReturn_timestampsToReturnInvalid, true
	}
	return 0, false
}

func TimestampsToReturnByName(value string) (enum TimestampsToReturn, ok bool) {
	switch value {
	case "timestampsToReturnSource":
		return TimestampsToReturn_timestampsToReturnSource, true
	case "timestampsToReturnServer":
		return TimestampsToReturn_timestampsToReturnServer, true
	case "timestampsToReturnBoth":
		return TimestampsToReturn_timestampsToReturnBoth, true
	case "timestampsToReturnNeither":
		return TimestampsToReturn_timestampsToReturnNeither, true
	case "timestampsToReturnInvalid":
		return TimestampsToReturn_timestampsToReturnInvalid, true
	}
	return 0, false
}

func TimestampsToReturnKnows(value uint32) bool {
	for _, typeValue := range TimestampsToReturnValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTimestampsToReturn(structType any) TimestampsToReturn {
	castFunc := func(typ any) TimestampsToReturn {
		if sTimestampsToReturn, ok := typ.(TimestampsToReturn); ok {
			return sTimestampsToReturn
		}
		return 0
	}
	return castFunc(structType)
}

func (m TimestampsToReturn) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m TimestampsToReturn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TimestampsToReturnParse(ctx context.Context, theBytes []byte) (TimestampsToReturn, error) {
	return TimestampsToReturnParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TimestampsToReturnParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TimestampsToReturn, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("TimestampsToReturn", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TimestampsToReturn")
	}
	if enum, ok := TimestampsToReturnByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TimestampsToReturn")
		return TimestampsToReturn(val), nil
	} else {
		return enum, nil
	}
}

func (e TimestampsToReturn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TimestampsToReturn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("TimestampsToReturn", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TimestampsToReturn) PLC4XEnumName() string {
	switch e {
	case TimestampsToReturn_timestampsToReturnSource:
		return "timestampsToReturnSource"
	case TimestampsToReturn_timestampsToReturnServer:
		return "timestampsToReturnServer"
	case TimestampsToReturn_timestampsToReturnBoth:
		return "timestampsToReturnBoth"
	case TimestampsToReturn_timestampsToReturnNeither:
		return "timestampsToReturnNeither"
	case TimestampsToReturn_timestampsToReturnInvalid:
		return "timestampsToReturnInvalid"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e TimestampsToReturn) String() string {
	return e.PLC4XEnumName()
}
