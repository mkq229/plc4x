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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/apache/plc4x/plc4go/protocols/ads/readwrite/model"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by code-generation. DO NOT EDIT.

type AdsXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m AdsXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "AmsSerialFrame":
		return model.AmsSerialFrameParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DataItem":
		// TODO: find a way to parse the sub types
		var dataFormatName string
		parsedInt1, err := strconv.ParseInt(parserArguments[1], 10, 32)
		if err != nil {
			return nil, err
		}
		stringLength := int32(parsedInt1)
		return model.DataItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), dataFormatName, stringLength)
	case "AdsMultiRequestItem":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		indexGroup := uint32(parsedUint0)
		return model.AdsMultiRequestItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), indexGroup)
	case "AmsSerialAcknowledgeFrame":
		return model.AmsSerialAcknowledgeFrameParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsData":
		commandId, _ := model.CommandIdByName(parserArguments[0])
		response := parserArguments[1] == "true"
		return model.AdsDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), commandId, response)
	case "AmsNetId":
		return model.AmsNetIdParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsStampHeader":
		return model.AdsStampHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsSerialResetFrame":
		return model.AmsSerialResetFrameParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsConstants":
		return model.AdsConstantsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsNotificationSample":
		return model.AdsNotificationSampleParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsTCPPacket":
		return model.AmsTCPPacketParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "State":
		return model.StateParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsPacket":
		return model.AmsPacketParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
