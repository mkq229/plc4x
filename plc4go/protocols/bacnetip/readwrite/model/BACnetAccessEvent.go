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

// BACnetAccessEvent is an enum
type BACnetAccessEvent uint16

type IBACnetAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAccessEvent_NONE                                   BACnetAccessEvent = 0
	BACnetAccessEvent_GRANTED                                BACnetAccessEvent = 1
	BACnetAccessEvent_MUSTER                                 BACnetAccessEvent = 2
	BACnetAccessEvent_PASSBACK_DETECTED                      BACnetAccessEvent = 3
	BACnetAccessEvent_DURESS                                 BACnetAccessEvent = 4
	BACnetAccessEvent_TRACE                                  BACnetAccessEvent = 5
	BACnetAccessEvent_LOCKOUT_MAX_ATTEMPTS                   BACnetAccessEvent = 6
	BACnetAccessEvent_LOCKOUT_OTHER                          BACnetAccessEvent = 7
	BACnetAccessEvent_LOCKOUT_RELINQUISHED                   BACnetAccessEvent = 8
	BACnetAccessEvent_LOCKED_BY_HIGHER_PRIORITY              BACnetAccessEvent = 9
	BACnetAccessEvent_OUT_OF_SERVICE                         BACnetAccessEvent = 10
	BACnetAccessEvent_OUT_OF_SERVICE_RELINQUISHED            BACnetAccessEvent = 11
	BACnetAccessEvent_ACCOMPANIMENT_BY                       BACnetAccessEvent = 12
	BACnetAccessEvent_AUTHENTICATION_FACTOR_READ             BACnetAccessEvent = 13
	BACnetAccessEvent_AUTHORIZATION_DELAYED                  BACnetAccessEvent = 14
	BACnetAccessEvent_VERIFICATION_REQUIRED                  BACnetAccessEvent = 15
	BACnetAccessEvent_NO_ENTRY_AFTER_GRANTED                 BACnetAccessEvent = 16
	BACnetAccessEvent_DENIED_DENY_ALL                        BACnetAccessEvent = 128
	BACnetAccessEvent_DENIED_UNKNOWN_CREDENTIAL              BACnetAccessEvent = 129
	BACnetAccessEvent_DENIED_AUTHENTICATION_UNAVAILABLE      BACnetAccessEvent = 130
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_TIMEOUT   BACnetAccessEvent = 131
	BACnetAccessEvent_DENIED_INCORRECT_AUTHENTICATION_FACTOR BACnetAccessEvent = 132
	BACnetAccessEvent_DENIED_ZONE_NO_ACCESS_RIGHTS           BACnetAccessEvent = 133
	BACnetAccessEvent_DENIED_POINT_NO_ACCESS_RIGHTS          BACnetAccessEvent = 134
	BACnetAccessEvent_DENIED_NO_ACCESS_RIGHTS                BACnetAccessEvent = 135
	BACnetAccessEvent_DENIED_OUT_OF_TIME_RANGE               BACnetAccessEvent = 136
	BACnetAccessEvent_DENIED_THREAT_LEVEL                    BACnetAccessEvent = 137
	BACnetAccessEvent_DENIED_PASSBACK                        BACnetAccessEvent = 138
	BACnetAccessEvent_DENIED_UNEXPECTED_LOCATION_USAGE       BACnetAccessEvent = 139
	BACnetAccessEvent_DENIED_MAX_ATTEMPTS                    BACnetAccessEvent = 140
	BACnetAccessEvent_DENIED_LOWER_OCCUPANCY_LIMIT           BACnetAccessEvent = 141
	BACnetAccessEvent_DENIED_UPPER_OCCUPANCY_LIMIT           BACnetAccessEvent = 142
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_LOST      BACnetAccessEvent = 143
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_STOLEN    BACnetAccessEvent = 144
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DAMAGED   BACnetAccessEvent = 145
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DESTROYED BACnetAccessEvent = 146
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DISABLED  BACnetAccessEvent = 147
	BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_ERROR     BACnetAccessEvent = 148
	BACnetAccessEvent_DENIED_CREDENTIAL_UNASSIGNED           BACnetAccessEvent = 149
	BACnetAccessEvent_DENIED_CREDENTIAL_NOT_PROVISIONED      BACnetAccessEvent = 150
	BACnetAccessEvent_DENIED_CREDENTIAL_NOT_YET_ACTIVE       BACnetAccessEvent = 151
	BACnetAccessEvent_DENIED_CREDENTIAL_EXPIRED              BACnetAccessEvent = 152
	BACnetAccessEvent_DENIED_CREDENTIAL_MANUAL_DISABLE       BACnetAccessEvent = 153
	BACnetAccessEvent_DENIED_CREDENTIAL_LOCKOUT              BACnetAccessEvent = 154
	BACnetAccessEvent_DENIED_CREDENTIAL_MAX_DAYS             BACnetAccessEvent = 155
	BACnetAccessEvent_DENIED_CREDENTIAL_MAX_USES             BACnetAccessEvent = 156
	BACnetAccessEvent_DENIED_CREDENTIAL_INACTIVITY           BACnetAccessEvent = 157
	BACnetAccessEvent_DENIED_CREDENTIAL_DISABLED             BACnetAccessEvent = 158
	BACnetAccessEvent_DENIED_NO_ACCOMPANIMENT                BACnetAccessEvent = 159
	BACnetAccessEvent_DENIED_INCORRECT_ACCOMPANIMENT         BACnetAccessEvent = 160
	BACnetAccessEvent_DENIED_LOCKOUT                         BACnetAccessEvent = 161
	BACnetAccessEvent_DENIED_VERIFICATION_FAILED             BACnetAccessEvent = 162
	BACnetAccessEvent_DENIED_VERIFICATION_TIMEOUT            BACnetAccessEvent = 163
	BACnetAccessEvent_DENIED_OTHER                           BACnetAccessEvent = 164
	BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE               BACnetAccessEvent = 0xFFFF
)

var BACnetAccessEventValues []BACnetAccessEvent

func init() {
	_ = errors.New
	BACnetAccessEventValues = []BACnetAccessEvent{
		BACnetAccessEvent_NONE,
		BACnetAccessEvent_GRANTED,
		BACnetAccessEvent_MUSTER,
		BACnetAccessEvent_PASSBACK_DETECTED,
		BACnetAccessEvent_DURESS,
		BACnetAccessEvent_TRACE,
		BACnetAccessEvent_LOCKOUT_MAX_ATTEMPTS,
		BACnetAccessEvent_LOCKOUT_OTHER,
		BACnetAccessEvent_LOCKOUT_RELINQUISHED,
		BACnetAccessEvent_LOCKED_BY_HIGHER_PRIORITY,
		BACnetAccessEvent_OUT_OF_SERVICE,
		BACnetAccessEvent_OUT_OF_SERVICE_RELINQUISHED,
		BACnetAccessEvent_ACCOMPANIMENT_BY,
		BACnetAccessEvent_AUTHENTICATION_FACTOR_READ,
		BACnetAccessEvent_AUTHORIZATION_DELAYED,
		BACnetAccessEvent_VERIFICATION_REQUIRED,
		BACnetAccessEvent_NO_ENTRY_AFTER_GRANTED,
		BACnetAccessEvent_DENIED_DENY_ALL,
		BACnetAccessEvent_DENIED_UNKNOWN_CREDENTIAL,
		BACnetAccessEvent_DENIED_AUTHENTICATION_UNAVAILABLE,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_TIMEOUT,
		BACnetAccessEvent_DENIED_INCORRECT_AUTHENTICATION_FACTOR,
		BACnetAccessEvent_DENIED_ZONE_NO_ACCESS_RIGHTS,
		BACnetAccessEvent_DENIED_POINT_NO_ACCESS_RIGHTS,
		BACnetAccessEvent_DENIED_NO_ACCESS_RIGHTS,
		BACnetAccessEvent_DENIED_OUT_OF_TIME_RANGE,
		BACnetAccessEvent_DENIED_THREAT_LEVEL,
		BACnetAccessEvent_DENIED_PASSBACK,
		BACnetAccessEvent_DENIED_UNEXPECTED_LOCATION_USAGE,
		BACnetAccessEvent_DENIED_MAX_ATTEMPTS,
		BACnetAccessEvent_DENIED_LOWER_OCCUPANCY_LIMIT,
		BACnetAccessEvent_DENIED_UPPER_OCCUPANCY_LIMIT,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_LOST,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_STOLEN,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DAMAGED,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DESTROYED,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DISABLED,
		BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_ERROR,
		BACnetAccessEvent_DENIED_CREDENTIAL_UNASSIGNED,
		BACnetAccessEvent_DENIED_CREDENTIAL_NOT_PROVISIONED,
		BACnetAccessEvent_DENIED_CREDENTIAL_NOT_YET_ACTIVE,
		BACnetAccessEvent_DENIED_CREDENTIAL_EXPIRED,
		BACnetAccessEvent_DENIED_CREDENTIAL_MANUAL_DISABLE,
		BACnetAccessEvent_DENIED_CREDENTIAL_LOCKOUT,
		BACnetAccessEvent_DENIED_CREDENTIAL_MAX_DAYS,
		BACnetAccessEvent_DENIED_CREDENTIAL_MAX_USES,
		BACnetAccessEvent_DENIED_CREDENTIAL_INACTIVITY,
		BACnetAccessEvent_DENIED_CREDENTIAL_DISABLED,
		BACnetAccessEvent_DENIED_NO_ACCOMPANIMENT,
		BACnetAccessEvent_DENIED_INCORRECT_ACCOMPANIMENT,
		BACnetAccessEvent_DENIED_LOCKOUT,
		BACnetAccessEvent_DENIED_VERIFICATION_FAILED,
		BACnetAccessEvent_DENIED_VERIFICATION_TIMEOUT,
		BACnetAccessEvent_DENIED_OTHER,
		BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessEventByValue(value uint16) (enum BACnetAccessEvent, ok bool) {
	switch value {
	case 0:
		return BACnetAccessEvent_NONE, true
	case 0xFFFF:
		return BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAccessEvent_GRANTED, true
	case 10:
		return BACnetAccessEvent_OUT_OF_SERVICE, true
	case 11:
		return BACnetAccessEvent_OUT_OF_SERVICE_RELINQUISHED, true
	case 12:
		return BACnetAccessEvent_ACCOMPANIMENT_BY, true
	case 128:
		return BACnetAccessEvent_DENIED_DENY_ALL, true
	case 129:
		return BACnetAccessEvent_DENIED_UNKNOWN_CREDENTIAL, true
	case 13:
		return BACnetAccessEvent_AUTHENTICATION_FACTOR_READ, true
	case 130:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_UNAVAILABLE, true
	case 131:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_TIMEOUT, true
	case 132:
		return BACnetAccessEvent_DENIED_INCORRECT_AUTHENTICATION_FACTOR, true
	case 133:
		return BACnetAccessEvent_DENIED_ZONE_NO_ACCESS_RIGHTS, true
	case 134:
		return BACnetAccessEvent_DENIED_POINT_NO_ACCESS_RIGHTS, true
	case 135:
		return BACnetAccessEvent_DENIED_NO_ACCESS_RIGHTS, true
	case 136:
		return BACnetAccessEvent_DENIED_OUT_OF_TIME_RANGE, true
	case 137:
		return BACnetAccessEvent_DENIED_THREAT_LEVEL, true
	case 138:
		return BACnetAccessEvent_DENIED_PASSBACK, true
	case 139:
		return BACnetAccessEvent_DENIED_UNEXPECTED_LOCATION_USAGE, true
	case 14:
		return BACnetAccessEvent_AUTHORIZATION_DELAYED, true
	case 140:
		return BACnetAccessEvent_DENIED_MAX_ATTEMPTS, true
	case 141:
		return BACnetAccessEvent_DENIED_LOWER_OCCUPANCY_LIMIT, true
	case 142:
		return BACnetAccessEvent_DENIED_UPPER_OCCUPANCY_LIMIT, true
	case 143:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_LOST, true
	case 144:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_STOLEN, true
	case 145:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DAMAGED, true
	case 146:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DESTROYED, true
	case 147:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DISABLED, true
	case 148:
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_ERROR, true
	case 149:
		return BACnetAccessEvent_DENIED_CREDENTIAL_UNASSIGNED, true
	case 15:
		return BACnetAccessEvent_VERIFICATION_REQUIRED, true
	case 150:
		return BACnetAccessEvent_DENIED_CREDENTIAL_NOT_PROVISIONED, true
	case 151:
		return BACnetAccessEvent_DENIED_CREDENTIAL_NOT_YET_ACTIVE, true
	case 152:
		return BACnetAccessEvent_DENIED_CREDENTIAL_EXPIRED, true
	case 153:
		return BACnetAccessEvent_DENIED_CREDENTIAL_MANUAL_DISABLE, true
	case 154:
		return BACnetAccessEvent_DENIED_CREDENTIAL_LOCKOUT, true
	case 155:
		return BACnetAccessEvent_DENIED_CREDENTIAL_MAX_DAYS, true
	case 156:
		return BACnetAccessEvent_DENIED_CREDENTIAL_MAX_USES, true
	case 157:
		return BACnetAccessEvent_DENIED_CREDENTIAL_INACTIVITY, true
	case 158:
		return BACnetAccessEvent_DENIED_CREDENTIAL_DISABLED, true
	case 159:
		return BACnetAccessEvent_DENIED_NO_ACCOMPANIMENT, true
	case 16:
		return BACnetAccessEvent_NO_ENTRY_AFTER_GRANTED, true
	case 160:
		return BACnetAccessEvent_DENIED_INCORRECT_ACCOMPANIMENT, true
	case 161:
		return BACnetAccessEvent_DENIED_LOCKOUT, true
	case 162:
		return BACnetAccessEvent_DENIED_VERIFICATION_FAILED, true
	case 163:
		return BACnetAccessEvent_DENIED_VERIFICATION_TIMEOUT, true
	case 164:
		return BACnetAccessEvent_DENIED_OTHER, true
	case 2:
		return BACnetAccessEvent_MUSTER, true
	case 3:
		return BACnetAccessEvent_PASSBACK_DETECTED, true
	case 4:
		return BACnetAccessEvent_DURESS, true
	case 5:
		return BACnetAccessEvent_TRACE, true
	case 6:
		return BACnetAccessEvent_LOCKOUT_MAX_ATTEMPTS, true
	case 7:
		return BACnetAccessEvent_LOCKOUT_OTHER, true
	case 8:
		return BACnetAccessEvent_LOCKOUT_RELINQUISHED, true
	case 9:
		return BACnetAccessEvent_LOCKED_BY_HIGHER_PRIORITY, true
	}
	return 0, false
}

func BACnetAccessEventByName(value string) (enum BACnetAccessEvent, ok bool) {
	switch value {
	case "NONE":
		return BACnetAccessEvent_NONE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE, true
	case "GRANTED":
		return BACnetAccessEvent_GRANTED, true
	case "OUT_OF_SERVICE":
		return BACnetAccessEvent_OUT_OF_SERVICE, true
	case "OUT_OF_SERVICE_RELINQUISHED":
		return BACnetAccessEvent_OUT_OF_SERVICE_RELINQUISHED, true
	case "ACCOMPANIMENT_BY":
		return BACnetAccessEvent_ACCOMPANIMENT_BY, true
	case "DENIED_DENY_ALL":
		return BACnetAccessEvent_DENIED_DENY_ALL, true
	case "DENIED_UNKNOWN_CREDENTIAL":
		return BACnetAccessEvent_DENIED_UNKNOWN_CREDENTIAL, true
	case "AUTHENTICATION_FACTOR_READ":
		return BACnetAccessEvent_AUTHENTICATION_FACTOR_READ, true
	case "DENIED_AUTHENTICATION_UNAVAILABLE":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_UNAVAILABLE, true
	case "DENIED_AUTHENTICATION_FACTOR_TIMEOUT":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_TIMEOUT, true
	case "DENIED_INCORRECT_AUTHENTICATION_FACTOR":
		return BACnetAccessEvent_DENIED_INCORRECT_AUTHENTICATION_FACTOR, true
	case "DENIED_ZONE_NO_ACCESS_RIGHTS":
		return BACnetAccessEvent_DENIED_ZONE_NO_ACCESS_RIGHTS, true
	case "DENIED_POINT_NO_ACCESS_RIGHTS":
		return BACnetAccessEvent_DENIED_POINT_NO_ACCESS_RIGHTS, true
	case "DENIED_NO_ACCESS_RIGHTS":
		return BACnetAccessEvent_DENIED_NO_ACCESS_RIGHTS, true
	case "DENIED_OUT_OF_TIME_RANGE":
		return BACnetAccessEvent_DENIED_OUT_OF_TIME_RANGE, true
	case "DENIED_THREAT_LEVEL":
		return BACnetAccessEvent_DENIED_THREAT_LEVEL, true
	case "DENIED_PASSBACK":
		return BACnetAccessEvent_DENIED_PASSBACK, true
	case "DENIED_UNEXPECTED_LOCATION_USAGE":
		return BACnetAccessEvent_DENIED_UNEXPECTED_LOCATION_USAGE, true
	case "AUTHORIZATION_DELAYED":
		return BACnetAccessEvent_AUTHORIZATION_DELAYED, true
	case "DENIED_MAX_ATTEMPTS":
		return BACnetAccessEvent_DENIED_MAX_ATTEMPTS, true
	case "DENIED_LOWER_OCCUPANCY_LIMIT":
		return BACnetAccessEvent_DENIED_LOWER_OCCUPANCY_LIMIT, true
	case "DENIED_UPPER_OCCUPANCY_LIMIT":
		return BACnetAccessEvent_DENIED_UPPER_OCCUPANCY_LIMIT, true
	case "DENIED_AUTHENTICATION_FACTOR_LOST":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_LOST, true
	case "DENIED_AUTHENTICATION_FACTOR_STOLEN":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_STOLEN, true
	case "DENIED_AUTHENTICATION_FACTOR_DAMAGED":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DAMAGED, true
	case "DENIED_AUTHENTICATION_FACTOR_DESTROYED":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DESTROYED, true
	case "DENIED_AUTHENTICATION_FACTOR_DISABLED":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DISABLED, true
	case "DENIED_AUTHENTICATION_FACTOR_ERROR":
		return BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_ERROR, true
	case "DENIED_CREDENTIAL_UNASSIGNED":
		return BACnetAccessEvent_DENIED_CREDENTIAL_UNASSIGNED, true
	case "VERIFICATION_REQUIRED":
		return BACnetAccessEvent_VERIFICATION_REQUIRED, true
	case "DENIED_CREDENTIAL_NOT_PROVISIONED":
		return BACnetAccessEvent_DENIED_CREDENTIAL_NOT_PROVISIONED, true
	case "DENIED_CREDENTIAL_NOT_YET_ACTIVE":
		return BACnetAccessEvent_DENIED_CREDENTIAL_NOT_YET_ACTIVE, true
	case "DENIED_CREDENTIAL_EXPIRED":
		return BACnetAccessEvent_DENIED_CREDENTIAL_EXPIRED, true
	case "DENIED_CREDENTIAL_MANUAL_DISABLE":
		return BACnetAccessEvent_DENIED_CREDENTIAL_MANUAL_DISABLE, true
	case "DENIED_CREDENTIAL_LOCKOUT":
		return BACnetAccessEvent_DENIED_CREDENTIAL_LOCKOUT, true
	case "DENIED_CREDENTIAL_MAX_DAYS":
		return BACnetAccessEvent_DENIED_CREDENTIAL_MAX_DAYS, true
	case "DENIED_CREDENTIAL_MAX_USES":
		return BACnetAccessEvent_DENIED_CREDENTIAL_MAX_USES, true
	case "DENIED_CREDENTIAL_INACTIVITY":
		return BACnetAccessEvent_DENIED_CREDENTIAL_INACTIVITY, true
	case "DENIED_CREDENTIAL_DISABLED":
		return BACnetAccessEvent_DENIED_CREDENTIAL_DISABLED, true
	case "DENIED_NO_ACCOMPANIMENT":
		return BACnetAccessEvent_DENIED_NO_ACCOMPANIMENT, true
	case "NO_ENTRY_AFTER_GRANTED":
		return BACnetAccessEvent_NO_ENTRY_AFTER_GRANTED, true
	case "DENIED_INCORRECT_ACCOMPANIMENT":
		return BACnetAccessEvent_DENIED_INCORRECT_ACCOMPANIMENT, true
	case "DENIED_LOCKOUT":
		return BACnetAccessEvent_DENIED_LOCKOUT, true
	case "DENIED_VERIFICATION_FAILED":
		return BACnetAccessEvent_DENIED_VERIFICATION_FAILED, true
	case "DENIED_VERIFICATION_TIMEOUT":
		return BACnetAccessEvent_DENIED_VERIFICATION_TIMEOUT, true
	case "DENIED_OTHER":
		return BACnetAccessEvent_DENIED_OTHER, true
	case "MUSTER":
		return BACnetAccessEvent_MUSTER, true
	case "PASSBACK_DETECTED":
		return BACnetAccessEvent_PASSBACK_DETECTED, true
	case "DURESS":
		return BACnetAccessEvent_DURESS, true
	case "TRACE":
		return BACnetAccessEvent_TRACE, true
	case "LOCKOUT_MAX_ATTEMPTS":
		return BACnetAccessEvent_LOCKOUT_MAX_ATTEMPTS, true
	case "LOCKOUT_OTHER":
		return BACnetAccessEvent_LOCKOUT_OTHER, true
	case "LOCKOUT_RELINQUISHED":
		return BACnetAccessEvent_LOCKOUT_RELINQUISHED, true
	case "LOCKED_BY_HIGHER_PRIORITY":
		return BACnetAccessEvent_LOCKED_BY_HIGHER_PRIORITY, true
	}
	return 0, false
}

func BACnetAccessEventKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessEventValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessEvent(structType any) BACnetAccessEvent {
	castFunc := func(typ any) BACnetAccessEvent {
		if sBACnetAccessEvent, ok := typ.(BACnetAccessEvent); ok {
			return sBACnetAccessEvent
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessEventParse(ctx context.Context, theBytes []byte) (BACnetAccessEvent, error) {
	return BACnetAccessEventParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessEvent, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("BACnetAccessEvent", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessEvent")
	}
	if enum, ok := BACnetAccessEventByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAccessEvent")
		return BACnetAccessEvent(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("BACnetAccessEvent", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessEvent) PLC4XEnumName() string {
	switch e {
	case BACnetAccessEvent_NONE:
		return "NONE"
	case BACnetAccessEvent_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessEvent_GRANTED:
		return "GRANTED"
	case BACnetAccessEvent_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	case BACnetAccessEvent_OUT_OF_SERVICE_RELINQUISHED:
		return "OUT_OF_SERVICE_RELINQUISHED"
	case BACnetAccessEvent_ACCOMPANIMENT_BY:
		return "ACCOMPANIMENT_BY"
	case BACnetAccessEvent_DENIED_DENY_ALL:
		return "DENIED_DENY_ALL"
	case BACnetAccessEvent_DENIED_UNKNOWN_CREDENTIAL:
		return "DENIED_UNKNOWN_CREDENTIAL"
	case BACnetAccessEvent_AUTHENTICATION_FACTOR_READ:
		return "AUTHENTICATION_FACTOR_READ"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_UNAVAILABLE:
		return "DENIED_AUTHENTICATION_UNAVAILABLE"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_TIMEOUT:
		return "DENIED_AUTHENTICATION_FACTOR_TIMEOUT"
	case BACnetAccessEvent_DENIED_INCORRECT_AUTHENTICATION_FACTOR:
		return "DENIED_INCORRECT_AUTHENTICATION_FACTOR"
	case BACnetAccessEvent_DENIED_ZONE_NO_ACCESS_RIGHTS:
		return "DENIED_ZONE_NO_ACCESS_RIGHTS"
	case BACnetAccessEvent_DENIED_POINT_NO_ACCESS_RIGHTS:
		return "DENIED_POINT_NO_ACCESS_RIGHTS"
	case BACnetAccessEvent_DENIED_NO_ACCESS_RIGHTS:
		return "DENIED_NO_ACCESS_RIGHTS"
	case BACnetAccessEvent_DENIED_OUT_OF_TIME_RANGE:
		return "DENIED_OUT_OF_TIME_RANGE"
	case BACnetAccessEvent_DENIED_THREAT_LEVEL:
		return "DENIED_THREAT_LEVEL"
	case BACnetAccessEvent_DENIED_PASSBACK:
		return "DENIED_PASSBACK"
	case BACnetAccessEvent_DENIED_UNEXPECTED_LOCATION_USAGE:
		return "DENIED_UNEXPECTED_LOCATION_USAGE"
	case BACnetAccessEvent_AUTHORIZATION_DELAYED:
		return "AUTHORIZATION_DELAYED"
	case BACnetAccessEvent_DENIED_MAX_ATTEMPTS:
		return "DENIED_MAX_ATTEMPTS"
	case BACnetAccessEvent_DENIED_LOWER_OCCUPANCY_LIMIT:
		return "DENIED_LOWER_OCCUPANCY_LIMIT"
	case BACnetAccessEvent_DENIED_UPPER_OCCUPANCY_LIMIT:
		return "DENIED_UPPER_OCCUPANCY_LIMIT"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_LOST:
		return "DENIED_AUTHENTICATION_FACTOR_LOST"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_STOLEN:
		return "DENIED_AUTHENTICATION_FACTOR_STOLEN"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DAMAGED:
		return "DENIED_AUTHENTICATION_FACTOR_DAMAGED"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DESTROYED:
		return "DENIED_AUTHENTICATION_FACTOR_DESTROYED"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_DISABLED:
		return "DENIED_AUTHENTICATION_FACTOR_DISABLED"
	case BACnetAccessEvent_DENIED_AUTHENTICATION_FACTOR_ERROR:
		return "DENIED_AUTHENTICATION_FACTOR_ERROR"
	case BACnetAccessEvent_DENIED_CREDENTIAL_UNASSIGNED:
		return "DENIED_CREDENTIAL_UNASSIGNED"
	case BACnetAccessEvent_VERIFICATION_REQUIRED:
		return "VERIFICATION_REQUIRED"
	case BACnetAccessEvent_DENIED_CREDENTIAL_NOT_PROVISIONED:
		return "DENIED_CREDENTIAL_NOT_PROVISIONED"
	case BACnetAccessEvent_DENIED_CREDENTIAL_NOT_YET_ACTIVE:
		return "DENIED_CREDENTIAL_NOT_YET_ACTIVE"
	case BACnetAccessEvent_DENIED_CREDENTIAL_EXPIRED:
		return "DENIED_CREDENTIAL_EXPIRED"
	case BACnetAccessEvent_DENIED_CREDENTIAL_MANUAL_DISABLE:
		return "DENIED_CREDENTIAL_MANUAL_DISABLE"
	case BACnetAccessEvent_DENIED_CREDENTIAL_LOCKOUT:
		return "DENIED_CREDENTIAL_LOCKOUT"
	case BACnetAccessEvent_DENIED_CREDENTIAL_MAX_DAYS:
		return "DENIED_CREDENTIAL_MAX_DAYS"
	case BACnetAccessEvent_DENIED_CREDENTIAL_MAX_USES:
		return "DENIED_CREDENTIAL_MAX_USES"
	case BACnetAccessEvent_DENIED_CREDENTIAL_INACTIVITY:
		return "DENIED_CREDENTIAL_INACTIVITY"
	case BACnetAccessEvent_DENIED_CREDENTIAL_DISABLED:
		return "DENIED_CREDENTIAL_DISABLED"
	case BACnetAccessEvent_DENIED_NO_ACCOMPANIMENT:
		return "DENIED_NO_ACCOMPANIMENT"
	case BACnetAccessEvent_NO_ENTRY_AFTER_GRANTED:
		return "NO_ENTRY_AFTER_GRANTED"
	case BACnetAccessEvent_DENIED_INCORRECT_ACCOMPANIMENT:
		return "DENIED_INCORRECT_ACCOMPANIMENT"
	case BACnetAccessEvent_DENIED_LOCKOUT:
		return "DENIED_LOCKOUT"
	case BACnetAccessEvent_DENIED_VERIFICATION_FAILED:
		return "DENIED_VERIFICATION_FAILED"
	case BACnetAccessEvent_DENIED_VERIFICATION_TIMEOUT:
		return "DENIED_VERIFICATION_TIMEOUT"
	case BACnetAccessEvent_DENIED_OTHER:
		return "DENIED_OTHER"
	case BACnetAccessEvent_MUSTER:
		return "MUSTER"
	case BACnetAccessEvent_PASSBACK_DETECTED:
		return "PASSBACK_DETECTED"
	case BACnetAccessEvent_DURESS:
		return "DURESS"
	case BACnetAccessEvent_TRACE:
		return "TRACE"
	case BACnetAccessEvent_LOCKOUT_MAX_ATTEMPTS:
		return "LOCKOUT_MAX_ATTEMPTS"
	case BACnetAccessEvent_LOCKOUT_OTHER:
		return "LOCKOUT_OTHER"
	case BACnetAccessEvent_LOCKOUT_RELINQUISHED:
		return "LOCKOUT_RELINQUISHED"
	case BACnetAccessEvent_LOCKED_BY_HIGHER_PRIORITY:
		return "LOCKED_BY_HIGHER_PRIORITY"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetAccessEvent) String() string {
	return e.PLC4XEnumName()
}
