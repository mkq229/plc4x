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
	"context"

	"github.com/pkg/errors"

	"github.com/apache/plc4x/plc4go/protocols/eip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

type EipParserHelper struct {
}

func (m EipParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (any, error) {
	switch typeName {
	case "PathSegment":
		return model.PathSegmentParseWithBuffer(context.Background(), io)
	case "EipConstants":
		return model.EipConstantsParseWithBuffer(context.Background(), io)
	case "TransportType":
		return model.TransportTypeParseWithBuffer(context.Background(), io)
	case "PortSegmentType":
		return model.PortSegmentTypeParseWithBuffer(context.Background(), io)
	case "NetworkConnectionParameters":
		return model.NetworkConnectionParametersParseWithBuffer(context.Background(), io)
	case "TypeId":
		return model.TypeIdParseWithBuffer(context.Background(), io)
	case "InstanceSegment":
		return model.InstanceSegmentParseWithBuffer(context.Background(), io)
	case "CIPData":
		packetLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CIPDataParseWithBuffer(context.Background(), io, packetLength)
	case "ClassSegment":
		return model.ClassSegmentParseWithBuffer(context.Background(), io)
	case "EipPacket":
		response, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.EipPacketParseWithBuffer(context.Background(), io, response)
	case "CIPAttributes":
		packetLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CIPAttributesParseWithBuffer(context.Background(), io, packetLength)
	case "CipService":
		connected, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		serviceLen, err := utils.StrToUint16(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CipServiceParseWithBuffer(context.Background(), io, connected, serviceLen)
	case "CommandSpecificDataItem":
		return model.CommandSpecificDataItemParseWithBuffer(context.Background(), io)
	case "Services":
		servicesLen, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.ServicesParseWithBuffer(context.Background(), io, servicesLen)
	case "LogicalSegmentType":
		return model.LogicalSegmentTypeParseWithBuffer(context.Background(), io)
	case "DataSegmentType":
		return model.DataSegmentTypeParseWithBuffer(context.Background(), io)
	case "CIPDataConnected":
		return model.CIPDataConnectedParseWithBuffer(context.Background(), io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
