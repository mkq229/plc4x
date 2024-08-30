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

// NLMEstablishConnectionToNetwork is the corresponding interface of NLMEstablishConnectionToNetwork
type NLMEstablishConnectionToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetTerminationTime returns TerminationTime (property field)
	GetTerminationTime() uint8
}

// NLMEstablishConnectionToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMEstablishConnectionToNetwork.
// This is useful for switch cases.
type NLMEstablishConnectionToNetworkExactly interface {
	NLMEstablishConnectionToNetwork
	isNLMEstablishConnectionToNetwork() bool
}

// _NLMEstablishConnectionToNetwork is the data-structure of this message
type _NLMEstablishConnectionToNetwork struct {
	*_NLM
	DestinationNetworkAddress uint16
	TerminationTime           uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMEstablishConnectionToNetwork) GetMessageType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMEstablishConnectionToNetwork) InitializeParent(parent NLM) {}

func (m *_NLMEstablishConnectionToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMEstablishConnectionToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NLMEstablishConnectionToNetwork) GetTerminationTime() uint8 {
	return m.TerminationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMEstablishConnectionToNetwork factory function for _NLMEstablishConnectionToNetwork
func NewNLMEstablishConnectionToNetwork(destinationNetworkAddress uint16, terminationTime uint8, apduLength uint16) *_NLMEstablishConnectionToNetwork {
	_result := &_NLMEstablishConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		TerminationTime:           terminationTime,
		_NLM:                      NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMEstablishConnectionToNetwork(structType any) NLMEstablishConnectionToNetwork {
	if casted, ok := structType.(NLMEstablishConnectionToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMEstablishConnectionToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMEstablishConnectionToNetwork) GetTypeName() string {
	return "NLMEstablishConnectionToNetwork"
}

func (m *_NLMEstablishConnectionToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (terminationTime)
	lengthInBits += 8

	return lengthInBits
}

func (m *_NLMEstablishConnectionToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMEstablishConnectionToNetworkParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMEstablishConnectionToNetwork, error) {
	return NLMEstablishConnectionToNetworkParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMEstablishConnectionToNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMEstablishConnectionToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMEstablishConnectionToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMEstablishConnectionToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := /*TODO: migrate me*/ readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field of NLMEstablishConnectionToNetwork")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (terminationTime)
	_terminationTime, _terminationTimeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("terminationTime", 8)
	if _terminationTimeErr != nil {
		return nil, errors.Wrap(_terminationTimeErr, "Error parsing 'terminationTime' field of NLMEstablishConnectionToNetwork")
	}
	terminationTime := _terminationTime

	if closeErr := readBuffer.CloseContext("NLMEstablishConnectionToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMEstablishConnectionToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMEstablishConnectionToNetwork{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		DestinationNetworkAddress: destinationNetworkAddress,
		TerminationTime:           terminationTime,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMEstablishConnectionToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMEstablishConnectionToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMEstablishConnectionToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMEstablishConnectionToNetwork")
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.GetDestinationNetworkAddress())
		_destinationNetworkAddressErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("destinationNetworkAddress", 16, uint16((destinationNetworkAddress)))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		// Simple Field (terminationTime)
		terminationTime := uint8(m.GetTerminationTime())
		_terminationTimeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("terminationTime", 8, uint8((terminationTime)))
		if _terminationTimeErr != nil {
			return errors.Wrap(_terminationTimeErr, "Error serializing 'terminationTime' field")
		}

		if popErr := writeBuffer.PopContext("NLMEstablishConnectionToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMEstablishConnectionToNetwork")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMEstablishConnectionToNetwork) isNLMEstablishConnectionToNetwork() bool {
	return true
}

func (m *_NLMEstablishConnectionToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
