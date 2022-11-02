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

// BACnetPropertyStatesNetworkPortCommand is the corresponding interface of BACnetPropertyStatesNetworkPortCommand
type BACnetPropertyStatesNetworkPortCommand interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetNetworkPortCommand returns NetworkPortCommand (property field)
	GetNetworkPortCommand() BACnetNetworkPortCommandTagged
}

// BACnetPropertyStatesNetworkPortCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesNetworkPortCommand.
// This is useful for switch cases.
type BACnetPropertyStatesNetworkPortCommandExactly interface {
	BACnetPropertyStatesNetworkPortCommand
	isBACnetPropertyStatesNetworkPortCommand() bool
}

// _BACnetPropertyStatesNetworkPortCommand is the data-structure of this message
type _BACnetPropertyStatesNetworkPortCommand struct {
	*_BACnetPropertyStates
	NetworkPortCommand BACnetNetworkPortCommandTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNetworkPortCommand) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNetworkPortCommand) GetNetworkPortCommand() BACnetNetworkPortCommandTagged {
	return m.NetworkPortCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNetworkPortCommand factory function for _BACnetPropertyStatesNetworkPortCommand
func NewBACnetPropertyStatesNetworkPortCommand(networkPortCommand BACnetNetworkPortCommandTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesNetworkPortCommand {
	_result := &_BACnetPropertyStatesNetworkPortCommand{
		NetworkPortCommand:    networkPortCommand,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNetworkPortCommand(structType interface{}) BACnetPropertyStatesNetworkPortCommand {
	if casted, ok := structType.(BACnetPropertyStatesNetworkPortCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkPortCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetTypeName() string {
	return "BACnetPropertyStatesNetworkPortCommand"
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkPortCommand)
	lengthInBits += m.NetworkPortCommand.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNetworkPortCommandParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesNetworkPortCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkPortCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkPortCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkPortCommand)
	if pullErr := readBuffer.PullContext("networkPortCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkPortCommand")
	}
	_networkPortCommand, _networkPortCommandErr := BACnetNetworkPortCommandTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _networkPortCommandErr != nil {
		return nil, errors.Wrap(_networkPortCommandErr, "Error parsing 'networkPortCommand' field of BACnetPropertyStatesNetworkPortCommand")
	}
	networkPortCommand := _networkPortCommand.(BACnetNetworkPortCommandTagged)
	if closeErr := readBuffer.CloseContext("networkPortCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkPortCommand")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkPortCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkPortCommand")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesNetworkPortCommand{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		NetworkPortCommand:    networkPortCommand,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesNetworkPortCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNetworkPortCommand) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkPortCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkPortCommand")
		}

		// Simple Field (networkPortCommand)
		if pushErr := writeBuffer.PushContext("networkPortCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkPortCommand")
		}
		_networkPortCommandErr := writeBuffer.WriteSerializable(m.GetNetworkPortCommand())
		if popErr := writeBuffer.PopContext("networkPortCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkPortCommand")
		}
		if _networkPortCommandErr != nil {
			return errors.Wrap(_networkPortCommandErr, "Error serializing 'networkPortCommand' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkPortCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkPortCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNetworkPortCommand) isBACnetPropertyStatesNetworkPortCommand() bool {
	return true
}

func (m *_BACnetPropertyStatesNetworkPortCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}