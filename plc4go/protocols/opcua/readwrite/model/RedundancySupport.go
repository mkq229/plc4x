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

// RedundancySupport is an enum
type RedundancySupport uint32

type IRedundancySupport interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	RedundancySupport_redundancySupportNone           RedundancySupport = 0
	RedundancySupport_redundancySupportCold           RedundancySupport = 1
	RedundancySupport_redundancySupportWarm           RedundancySupport = 2
	RedundancySupport_redundancySupportHot            RedundancySupport = 3
	RedundancySupport_redundancySupportTransparent    RedundancySupport = 4
	RedundancySupport_redundancySupportHotAndMirrored RedundancySupport = 5
)

var RedundancySupportValues []RedundancySupport

func init() {
	_ = errors.New
	RedundancySupportValues = []RedundancySupport{
		RedundancySupport_redundancySupportNone,
		RedundancySupport_redundancySupportCold,
		RedundancySupport_redundancySupportWarm,
		RedundancySupport_redundancySupportHot,
		RedundancySupport_redundancySupportTransparent,
		RedundancySupport_redundancySupportHotAndMirrored,
	}
}

func RedundancySupportByValue(value uint32) (enum RedundancySupport, ok bool) {
	switch value {
	case 0:
		return RedundancySupport_redundancySupportNone, true
	case 1:
		return RedundancySupport_redundancySupportCold, true
	case 2:
		return RedundancySupport_redundancySupportWarm, true
	case 3:
		return RedundancySupport_redundancySupportHot, true
	case 4:
		return RedundancySupport_redundancySupportTransparent, true
	case 5:
		return RedundancySupport_redundancySupportHotAndMirrored, true
	}
	return 0, false
}

func RedundancySupportByName(value string) (enum RedundancySupport, ok bool) {
	switch value {
	case "redundancySupportNone":
		return RedundancySupport_redundancySupportNone, true
	case "redundancySupportCold":
		return RedundancySupport_redundancySupportCold, true
	case "redundancySupportWarm":
		return RedundancySupport_redundancySupportWarm, true
	case "redundancySupportHot":
		return RedundancySupport_redundancySupportHot, true
	case "redundancySupportTransparent":
		return RedundancySupport_redundancySupportTransparent, true
	case "redundancySupportHotAndMirrored":
		return RedundancySupport_redundancySupportHotAndMirrored, true
	}
	return 0, false
}

func RedundancySupportKnows(value uint32) bool {
	for _, typeValue := range RedundancySupportValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastRedundancySupport(structType any) RedundancySupport {
	castFunc := func(typ any) RedundancySupport {
		if sRedundancySupport, ok := typ.(RedundancySupport); ok {
			return sRedundancySupport
		}
		return 0
	}
	return castFunc(structType)
}

func (m RedundancySupport) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m RedundancySupport) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RedundancySupportParse(ctx context.Context, theBytes []byte) (RedundancySupport, error) {
	return RedundancySupportParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RedundancySupportParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RedundancySupport, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("RedundancySupport", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading RedundancySupport")
	}
	if enum, ok := RedundancySupportByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for RedundancySupport")
		return RedundancySupport(val), nil
	} else {
		return enum, nil
	}
}

func (e RedundancySupport) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e RedundancySupport) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("RedundancySupport", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e RedundancySupport) PLC4XEnumName() string {
	switch e {
	case RedundancySupport_redundancySupportNone:
		return "redundancySupportNone"
	case RedundancySupport_redundancySupportCold:
		return "redundancySupportCold"
	case RedundancySupport_redundancySupportWarm:
		return "redundancySupportWarm"
	case RedundancySupport_redundancySupportHot:
		return "redundancySupportHot"
	case RedundancySupport_redundancySupportTransparent:
		return "redundancySupportTransparent"
	case RedundancySupport_redundancySupportHotAndMirrored:
		return "redundancySupportHotAndMirrored"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e RedundancySupport) String() string {
	return e.PLC4XEnumName()
}
