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

// RedundantServerDataType is the corresponding interface of RedundantServerDataType
type RedundantServerDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetServerId returns ServerId (property field)
	GetServerId() PascalString
	// GetServiceLevel returns ServiceLevel (property field)
	GetServiceLevel() uint8
	// GetServerState returns ServerState (property field)
	GetServerState() ServerState
}

// RedundantServerDataTypeExactly can be used when we want exactly this type and not a type which fulfills RedundantServerDataType.
// This is useful for switch cases.
type RedundantServerDataTypeExactly interface {
	RedundantServerDataType
	isRedundantServerDataType() bool
}

// _RedundantServerDataType is the data-structure of this message
type _RedundantServerDataType struct {
	*_ExtensionObjectDefinition
	ServerId     PascalString
	ServiceLevel uint8
	ServerState  ServerState
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RedundantServerDataType) GetIdentifier() string {
	return "855"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RedundantServerDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RedundantServerDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RedundantServerDataType) GetServerId() PascalString {
	return m.ServerId
}

func (m *_RedundantServerDataType) GetServiceLevel() uint8 {
	return m.ServiceLevel
}

func (m *_RedundantServerDataType) GetServerState() ServerState {
	return m.ServerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRedundantServerDataType factory function for _RedundantServerDataType
func NewRedundantServerDataType(serverId PascalString, serviceLevel uint8, serverState ServerState) *_RedundantServerDataType {
	_result := &_RedundantServerDataType{
		ServerId:                   serverId,
		ServiceLevel:               serviceLevel,
		ServerState:                serverState,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRedundantServerDataType(structType any) RedundantServerDataType {
	if casted, ok := structType.(RedundantServerDataType); ok {
		return casted
	}
	if casted, ok := structType.(*RedundantServerDataType); ok {
		return *casted
	}
	return nil
}

func (m *_RedundantServerDataType) GetTypeName() string {
	return "RedundantServerDataType"
}

func (m *_RedundantServerDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (serverId)
	lengthInBits += m.ServerId.GetLengthInBits(ctx)

	// Simple field (serviceLevel)
	lengthInBits += 8

	// Simple field (serverState)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RedundantServerDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RedundantServerDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (RedundantServerDataType, error) {
	return RedundantServerDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RedundantServerDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RedundantServerDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RedundantServerDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RedundantServerDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (serverId)
	if pullErr := readBuffer.PullContext("serverId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverId")
	}
	_serverId, _serverIdErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _serverIdErr != nil {
		return nil, errors.Wrap(_serverIdErr, "Error parsing 'serverId' field of RedundantServerDataType")
	}
	serverId := _serverId.(PascalString)
	if closeErr := readBuffer.CloseContext("serverId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverId")
	}

	// Simple Field (serviceLevel)
	_serviceLevel, _serviceLevelErr := /*TODO: migrate me*/ readBuffer.ReadUint8("serviceLevel", 8)
	if _serviceLevelErr != nil {
		return nil, errors.Wrap(_serviceLevelErr, "Error parsing 'serviceLevel' field of RedundantServerDataType")
	}
	serviceLevel := _serviceLevel

	// Simple Field (serverState)
	if pullErr := readBuffer.PullContext("serverState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverState")
	}
	_serverState, _serverStateErr := ServerStateParseWithBuffer(ctx, readBuffer)
	if _serverStateErr != nil {
		return nil, errors.Wrap(_serverStateErr, "Error parsing 'serverState' field of RedundantServerDataType")
	}
	serverState := _serverState
	if closeErr := readBuffer.CloseContext("serverState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverState")
	}

	if closeErr := readBuffer.CloseContext("RedundantServerDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RedundantServerDataType")
	}

	// Create a partially initialized instance
	_child := &_RedundantServerDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ServerId:                   serverId,
		ServiceLevel:               serviceLevel,
		ServerState:                serverState,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RedundantServerDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RedundantServerDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RedundantServerDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RedundantServerDataType")
		}

		// Simple Field (serverId)
		if pushErr := writeBuffer.PushContext("serverId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverId")
		}
		_serverIdErr := writeBuffer.WriteSerializable(ctx, m.GetServerId())
		if popErr := writeBuffer.PopContext("serverId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverId")
		}
		if _serverIdErr != nil {
			return errors.Wrap(_serverIdErr, "Error serializing 'serverId' field")
		}

		// Simple Field (serviceLevel)
		serviceLevel := uint8(m.GetServiceLevel())
		_serviceLevelErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("serviceLevel", 8, uint8((serviceLevel)))
		if _serviceLevelErr != nil {
			return errors.Wrap(_serviceLevelErr, "Error serializing 'serviceLevel' field")
		}

		// Simple Field (serverState)
		if pushErr := writeBuffer.PushContext("serverState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverState")
		}
		_serverStateErr := writeBuffer.WriteSerializable(ctx, m.GetServerState())
		if popErr := writeBuffer.PopContext("serverState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverState")
		}
		if _serverStateErr != nil {
			return errors.Wrap(_serverStateErr, "Error serializing 'serverState' field")
		}

		if popErr := writeBuffer.PopContext("RedundantServerDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RedundantServerDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RedundantServerDataType) isRedundantServerDataType() bool {
	return true
}

func (m *_RedundantServerDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
