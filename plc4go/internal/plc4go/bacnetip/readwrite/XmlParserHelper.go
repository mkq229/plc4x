/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by code-generation. DO NOT EDIT.

type BacnetipXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m BacnetipXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "BACnetContextTag":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumberArgument := uint8(parsedUint0)
		dataType := model.BACnetDataTypeByName(parserArguments[1])
		return model.BACnetContextTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumberArgument, dataType)
	case "BACnetStatusFlags":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetStatusFlagsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetTagPayloadReal":
		return model.BACnetTagPayloadRealParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NLM":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		apduLength := uint16(parsedUint0)
		return model.NLMParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), apduLength)
	case "BACnetActionCommand":
		return model.BACnetActionCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadDate":
		return model.BACnetTagPayloadDateParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetNotificationParametersExtendedParameters":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersExtendedParametersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetNotificationParametersChangeOfValueNewValue":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersChangeOfValueNewValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetTagPayloadEnumerated":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadEnumeratedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetTagPayloadOctetString":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadOctetStringParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetServiceAckAtomicReadFileStreamOrRecord":
		return model.BACnetServiceAckAtomicReadFileStreamOrRecordParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NPDUControl":
		return model.NPDUControlParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyStates":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetPropertyStatesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetReadAccessSpecification":
		return model.BACnetReadAccessSpecificationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConstructedData":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		// TODO: find a way to parse the sub types
		var propertyIdentifierArgument model.BACnetContextTagPropertyIdentifier
		return model.BACnetConstructedDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType, &propertyIdentifierArgument)
	case "BACnetSegmentation":
		return model.BACnetSegmentationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadTime":
		return model.BACnetTagPayloadTimeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConfirmedServiceACK":
		return model.BACnetConfirmedServiceACKParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetTagPayloadSignedInteger":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadSignedIntegerParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetUnconfirmedServiceRequest":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		len := uint16(parsedUint0)
		return model.BACnetUnconfirmedServiceRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), len)
	case "BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord":
		return model.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BVLC":
		return model.BVLCParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadObjectIdentifier":
		return model.BACnetTagPayloadObjectIdentifierParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyWriteDefinition":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		return model.BACnetPropertyWriteDefinitionParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType)
	case "BACnetDateTime":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetDateTimeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "APDU":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		apduLength := uint16(parsedUint0)
		return model.APDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), apduLength)
	case "BACnetTagPayloadCharacterString":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadCharacterStringParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetError":
		return model.BACnetErrorParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTimeStamp":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetTimeStampParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetNotificationParameters":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		return model.BACnetNotificationParametersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType)
	case "BACnetConfirmedServiceRequest":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		len := uint16(parsedUint0)
		return model.BACnetConfirmedServiceRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), len)
	case "BACnetAddress":
		return model.BACnetAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadUnsignedInteger":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadUnsignedIntegerParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetApplicationTag":
		return model.BACnetApplicationTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadBitString":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadBitStringParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetDeviceObjectPropertyReference":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetDeviceObjectPropertyReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConstructedDataElement":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		// TODO: find a way to parse the sub types
		var propertyIdentifier model.BACnetContextTagPropertyIdentifier
		return model.BACnetConstructedDataElementParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType, &propertyIdentifier)
	case "BACnetPropertyValues":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		return model.BACnetPropertyValuesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType)
	case "BACnetTagHeader":
		return model.BACnetTagHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadBoolean":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadBooleanParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetTagPayloadDouble":
		return model.BACnetTagPayloadDoubleParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyValue":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		return model.BACnetPropertyValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType)
	case "NLMInitalizeRoutingTablePortMapping":
		return model.NLMInitalizeRoutingTablePortMappingParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetWriteAccessSpecification":
		return model.BACnetWriteAccessSpecificationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetServiceAck":
		return model.BACnetServiceAckParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetBinaryPV":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetBinaryPVParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetAction":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetActionParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "NPDU":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		npduLength := uint16(parsedUint0)
		return model.NPDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), npduLength)
	case "BACnetPropertyReference":
		return model.BACnetPropertyReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BVLCWriteBroadcastDistributionTableEntry":
		return model.BVLCWriteBroadcastDistributionTableEntryParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
