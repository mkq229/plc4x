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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNetworkSecurityPolicy is the corresponding interface of BACnetNetworkSecurityPolicy
type BACnetNetworkSecurityPolicy interface {
	utils.LengthAware
	utils.Serializable
	// GetPortId returns PortId (property field)
	GetPortId() BACnetContextTagUnsignedInteger
	// GetSecurityLevel returns SecurityLevel (property field)
	GetSecurityLevel() BACnetSecurityPolicyTagged
}

// BACnetNetworkSecurityPolicyExactly can be used when we want exactly this type and not a type which fulfills BACnetNetworkSecurityPolicy.
// This is useful for switch cases.
type BACnetNetworkSecurityPolicyExactly interface {
	BACnetNetworkSecurityPolicy
	isBACnetNetworkSecurityPolicy() bool
}

// _BACnetNetworkSecurityPolicy is the data-structure of this message
type _BACnetNetworkSecurityPolicy struct {
	PortId        BACnetContextTagUnsignedInteger
	SecurityLevel BACnetSecurityPolicyTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNetworkSecurityPolicy) GetPortId() BACnetContextTagUnsignedInteger {
	return m.PortId
}

func (m *_BACnetNetworkSecurityPolicy) GetSecurityLevel() BACnetSecurityPolicyTagged {
	return m.SecurityLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNetworkSecurityPolicy factory function for _BACnetNetworkSecurityPolicy
func NewBACnetNetworkSecurityPolicy(portId BACnetContextTagUnsignedInteger, securityLevel BACnetSecurityPolicyTagged) *_BACnetNetworkSecurityPolicy {
	return &_BACnetNetworkSecurityPolicy{PortId: portId, SecurityLevel: securityLevel}
}

// Deprecated: use the interface for direct cast
func CastBACnetNetworkSecurityPolicy(structType interface{}) BACnetNetworkSecurityPolicy {
	if casted, ok := structType.(BACnetNetworkSecurityPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNetworkSecurityPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNetworkSecurityPolicy) GetTypeName() string {
	return "BACnetNetworkSecurityPolicy"
}

func (m *_BACnetNetworkSecurityPolicy) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNetworkSecurityPolicy) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (portId)
	lengthInBits += m.PortId.GetLengthInBits()

	// Simple field (securityLevel)
	lengthInBits += m.SecurityLevel.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNetworkSecurityPolicy) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNetworkSecurityPolicyParse(readBuffer utils.ReadBuffer) (BACnetNetworkSecurityPolicy, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNetworkSecurityPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNetworkSecurityPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (portId)
	if pullErr := readBuffer.PullContext("portId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for portId")
	}
	_portId, _portIdErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _portIdErr != nil {
		return nil, errors.Wrap(_portIdErr, "Error parsing 'portId' field of BACnetNetworkSecurityPolicy")
	}
	portId := _portId.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("portId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for portId")
	}

	// Simple Field (securityLevel)
	if pullErr := readBuffer.PullContext("securityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityLevel")
	}
	_securityLevel, _securityLevelErr := BACnetSecurityPolicyTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _securityLevelErr != nil {
		return nil, errors.Wrap(_securityLevelErr, "Error parsing 'securityLevel' field of BACnetNetworkSecurityPolicy")
	}
	securityLevel := _securityLevel.(BACnetSecurityPolicyTagged)
	if closeErr := readBuffer.CloseContext("securityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetNetworkSecurityPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNetworkSecurityPolicy")
	}

	// Create the instance
	return &_BACnetNetworkSecurityPolicy{
		PortId:        portId,
		SecurityLevel: securityLevel,
	}, nil
}

func (m *_BACnetNetworkSecurityPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNetworkSecurityPolicy) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNetworkSecurityPolicy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNetworkSecurityPolicy")
	}

	// Simple Field (portId)
	if pushErr := writeBuffer.PushContext("portId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for portId")
	}
	_portIdErr := writeBuffer.WriteSerializable(m.GetPortId())
	if popErr := writeBuffer.PopContext("portId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for portId")
	}
	if _portIdErr != nil {
		return errors.Wrap(_portIdErr, "Error serializing 'portId' field")
	}

	// Simple Field (securityLevel)
	if pushErr := writeBuffer.PushContext("securityLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for securityLevel")
	}
	_securityLevelErr := writeBuffer.WriteSerializable(m.GetSecurityLevel())
	if popErr := writeBuffer.PopContext("securityLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for securityLevel")
	}
	if _securityLevelErr != nil {
		return errors.Wrap(_securityLevelErr, "Error serializing 'securityLevel' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNetworkSecurityPolicy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNetworkSecurityPolicy")
	}
	return nil
}

func (m *_BACnetNetworkSecurityPolicy) isBACnetNetworkSecurityPolicy() bool {
	return true
}

func (m *_BACnetNetworkSecurityPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}