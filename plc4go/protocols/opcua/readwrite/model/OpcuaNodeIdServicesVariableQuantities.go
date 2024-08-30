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

// OpcuaNodeIdServicesVariableQuantities is an enum
type OpcuaNodeIdServicesVariableQuantities int32

type IOpcuaNodeIdServicesVariableQuantities interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Symbol            OpcuaNodeIdServicesVariableQuantities = 32504
	OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Description       OpcuaNodeIdServicesVariableQuantities = 32505
	OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Annotation        OpcuaNodeIdServicesVariableQuantities = 32506
	OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_ConversionService OpcuaNodeIdServicesVariableQuantities = 32507
	OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Dimension         OpcuaNodeIdServicesVariableQuantities = 32508
)

var OpcuaNodeIdServicesVariableQuantitiesValues []OpcuaNodeIdServicesVariableQuantities

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableQuantitiesValues = []OpcuaNodeIdServicesVariableQuantities{
		OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Symbol,
		OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Description,
		OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Annotation,
		OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_ConversionService,
		OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Dimension,
	}
}

func OpcuaNodeIdServicesVariableQuantitiesByValue(value int32) (enum OpcuaNodeIdServicesVariableQuantities, ok bool) {
	switch value {
	case 32504:
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Symbol, true
	case 32505:
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Description, true
	case 32506:
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Annotation, true
	case 32507:
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_ConversionService, true
	case 32508:
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Dimension, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableQuantitiesByName(value string) (enum OpcuaNodeIdServicesVariableQuantities, ok bool) {
	switch value {
	case "QuantitiesFolderType_Quantity_Placeholder_Symbol":
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Symbol, true
	case "QuantitiesFolderType_Quantity_Placeholder_Description":
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Description, true
	case "QuantitiesFolderType_Quantity_Placeholder_Annotation":
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Annotation, true
	case "QuantitiesFolderType_Quantity_Placeholder_ConversionService":
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_ConversionService, true
	case "QuantitiesFolderType_Quantity_Placeholder_Dimension":
		return OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Dimension, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableQuantitiesKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableQuantitiesValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableQuantities(structType any) OpcuaNodeIdServicesVariableQuantities {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableQuantities {
		if sOpcuaNodeIdServicesVariableQuantities, ok := typ.(OpcuaNodeIdServicesVariableQuantities); ok {
			return sOpcuaNodeIdServicesVariableQuantities
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableQuantities) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableQuantities) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableQuantitiesParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableQuantities, error) {
	return OpcuaNodeIdServicesVariableQuantitiesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableQuantitiesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableQuantities, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableQuantities", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableQuantities")
	}
	if enum, ok := OpcuaNodeIdServicesVariableQuantitiesByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableQuantities")
		return OpcuaNodeIdServicesVariableQuantities(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableQuantities) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableQuantities) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableQuantities", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableQuantities) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Symbol:
		return "QuantitiesFolderType_Quantity_Placeholder_Symbol"
	case OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Description:
		return "QuantitiesFolderType_Quantity_Placeholder_Description"
	case OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Annotation:
		return "QuantitiesFolderType_Quantity_Placeholder_Annotation"
	case OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_ConversionService:
		return "QuantitiesFolderType_Quantity_Placeholder_ConversionService"
	case OpcuaNodeIdServicesVariableQuantities_QuantitiesFolderType_Quantity_Placeholder_Dimension:
		return "QuantitiesFolderType_Quantity_Placeholder_Dimension"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableQuantities) String() string {
	return e.PLC4XEnumName()
}
