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

// DestinationAddressType is an enum
type DestinationAddressType uint8

type IDestinationAddressType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DestinationAddressType_PointToPointToMultiPoint DestinationAddressType = 0x03
	DestinationAddressType_PointToMultiPoint        DestinationAddressType = 0x05
	DestinationAddressType_PointToPoint             DestinationAddressType = 0x06
)

var DestinationAddressTypeValues []DestinationAddressType

func init() {
	_ = errors.New
	DestinationAddressTypeValues = []DestinationAddressType{
		DestinationAddressType_PointToPointToMultiPoint,
		DestinationAddressType_PointToMultiPoint,
		DestinationAddressType_PointToPoint,
	}
}

func DestinationAddressTypeByValue(value uint8) (enum DestinationAddressType, ok bool) {
	switch value {
	case 0x03:
		return DestinationAddressType_PointToPointToMultiPoint, true
	case 0x05:
		return DestinationAddressType_PointToMultiPoint, true
	case 0x06:
		return DestinationAddressType_PointToPoint, true
	}
	return 0, false
}

func DestinationAddressTypeByName(value string) (enum DestinationAddressType, ok bool) {
	switch value {
	case "PointToPointToMultiPoint":
		return DestinationAddressType_PointToPointToMultiPoint, true
	case "PointToMultiPoint":
		return DestinationAddressType_PointToMultiPoint, true
	case "PointToPoint":
		return DestinationAddressType_PointToPoint, true
	}
	return 0, false
}

func DestinationAddressTypeKnows(value uint8) bool {
	for _, typeValue := range DestinationAddressTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDestinationAddressType(structType any) DestinationAddressType {
	castFunc := func(typ any) DestinationAddressType {
		if sDestinationAddressType, ok := typ.(DestinationAddressType); ok {
			return sDestinationAddressType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DestinationAddressType) GetLengthInBits(ctx context.Context) uint16 {
	return 3
}

func (m DestinationAddressType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DestinationAddressTypeParse(ctx context.Context, theBytes []byte) (DestinationAddressType, error) {
	return DestinationAddressTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DestinationAddressTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DestinationAddressType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("DestinationAddressType", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DestinationAddressType")
	}
	if enum, ok := DestinationAddressTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DestinationAddressType")
		return DestinationAddressType(val), nil
	} else {
		return enum, nil
	}
}

func (e DestinationAddressType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DestinationAddressType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("DestinationAddressType", 3, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DestinationAddressType) PLC4XEnumName() string {
	switch e {
	case DestinationAddressType_PointToPointToMultiPoint:
		return "PointToPointToMultiPoint"
	case DestinationAddressType_PointToMultiPoint:
		return "PointToMultiPoint"
	case DestinationAddressType_PointToPoint:
		return "PointToPoint"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e DestinationAddressType) String() string {
	return e.PLC4XEnumName()
}
