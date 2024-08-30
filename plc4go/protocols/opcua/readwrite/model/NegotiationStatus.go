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

// NegotiationStatus is an enum
type NegotiationStatus uint32

type INegotiationStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	NegotiationStatus_negotiationStatusInProgress    NegotiationStatus = 0
	NegotiationStatus_negotiationStatusComplete      NegotiationStatus = 1
	NegotiationStatus_negotiationStatusFailed        NegotiationStatus = 2
	NegotiationStatus_negotiationStatusUnknown       NegotiationStatus = 3
	NegotiationStatus_negotiationStatusNoNegotiation NegotiationStatus = 4
)

var NegotiationStatusValues []NegotiationStatus

func init() {
	_ = errors.New
	NegotiationStatusValues = []NegotiationStatus{
		NegotiationStatus_negotiationStatusInProgress,
		NegotiationStatus_negotiationStatusComplete,
		NegotiationStatus_negotiationStatusFailed,
		NegotiationStatus_negotiationStatusUnknown,
		NegotiationStatus_negotiationStatusNoNegotiation,
	}
}

func NegotiationStatusByValue(value uint32) (enum NegotiationStatus, ok bool) {
	switch value {
	case 0:
		return NegotiationStatus_negotiationStatusInProgress, true
	case 1:
		return NegotiationStatus_negotiationStatusComplete, true
	case 2:
		return NegotiationStatus_negotiationStatusFailed, true
	case 3:
		return NegotiationStatus_negotiationStatusUnknown, true
	case 4:
		return NegotiationStatus_negotiationStatusNoNegotiation, true
	}
	return 0, false
}

func NegotiationStatusByName(value string) (enum NegotiationStatus, ok bool) {
	switch value {
	case "negotiationStatusInProgress":
		return NegotiationStatus_negotiationStatusInProgress, true
	case "negotiationStatusComplete":
		return NegotiationStatus_negotiationStatusComplete, true
	case "negotiationStatusFailed":
		return NegotiationStatus_negotiationStatusFailed, true
	case "negotiationStatusUnknown":
		return NegotiationStatus_negotiationStatusUnknown, true
	case "negotiationStatusNoNegotiation":
		return NegotiationStatus_negotiationStatusNoNegotiation, true
	}
	return 0, false
}

func NegotiationStatusKnows(value uint32) bool {
	for _, typeValue := range NegotiationStatusValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNegotiationStatus(structType any) NegotiationStatus {
	castFunc := func(typ any) NegotiationStatus {
		if sNegotiationStatus, ok := typ.(NegotiationStatus); ok {
			return sNegotiationStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m NegotiationStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m NegotiationStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NegotiationStatusParse(ctx context.Context, theBytes []byte) (NegotiationStatus, error) {
	return NegotiationStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NegotiationStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NegotiationStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("NegotiationStatus", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NegotiationStatus")
	}
	if enum, ok := NegotiationStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for NegotiationStatus")
		return NegotiationStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e NegotiationStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NegotiationStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("NegotiationStatus", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NegotiationStatus) PLC4XEnumName() string {
	switch e {
	case NegotiationStatus_negotiationStatusInProgress:
		return "negotiationStatusInProgress"
	case NegotiationStatus_negotiationStatusComplete:
		return "negotiationStatusComplete"
	case NegotiationStatus_negotiationStatusFailed:
		return "negotiationStatusFailed"
	case NegotiationStatus_negotiationStatusUnknown:
		return "negotiationStatusUnknown"
	case NegotiationStatus_negotiationStatusNoNegotiation:
		return "negotiationStatusNoNegotiation"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e NegotiationStatus) String() string {
	return e.PLC4XEnumName()
}
