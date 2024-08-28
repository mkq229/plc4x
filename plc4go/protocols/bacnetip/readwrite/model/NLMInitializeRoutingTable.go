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

// NLMInitializeRoutingTable is the corresponding interface of NLMInitializeRoutingTable
type NLMInitializeRoutingTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetNumberOfPorts returns NumberOfPorts (property field)
	GetNumberOfPorts() uint8
	// GetPortMappings returns PortMappings (property field)
	GetPortMappings() []NLMInitializeRoutingTablePortMapping
}

// NLMInitializeRoutingTableExactly can be used when we want exactly this type and not a type which fulfills NLMInitializeRoutingTable.
// This is useful for switch cases.
type NLMInitializeRoutingTableExactly interface {
	NLMInitializeRoutingTable
	isNLMInitializeRoutingTable() bool
}

// _NLMInitializeRoutingTable is the data-structure of this message
type _NLMInitializeRoutingTable struct {
	*_NLM
	NumberOfPorts uint8
	PortMappings  []NLMInitializeRoutingTablePortMapping
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMInitializeRoutingTable) GetMessageType() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMInitializeRoutingTable) InitializeParent(parent NLM) {}

func (m *_NLMInitializeRoutingTable) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMInitializeRoutingTable) GetNumberOfPorts() uint8 {
	return m.NumberOfPorts
}

func (m *_NLMInitializeRoutingTable) GetPortMappings() []NLMInitializeRoutingTablePortMapping {
	return m.PortMappings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMInitializeRoutingTable factory function for _NLMInitializeRoutingTable
func NewNLMInitializeRoutingTable(numberOfPorts uint8, portMappings []NLMInitializeRoutingTablePortMapping, apduLength uint16) *_NLMInitializeRoutingTable {
	_result := &_NLMInitializeRoutingTable{
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
		_NLM:          NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMInitializeRoutingTable(structType any) NLMInitializeRoutingTable {
	if casted, ok := structType.(NLMInitializeRoutingTable); ok {
		return casted
	}
	if casted, ok := structType.(*NLMInitializeRoutingTable); ok {
		return *casted
	}
	return nil
}

func (m *_NLMInitializeRoutingTable) GetTypeName() string {
	return "NLMInitializeRoutingTable"
}

func (m *_NLMInitializeRoutingTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numberOfPorts)
	lengthInBits += 8

	// Array field
	if len(m.PortMappings) > 0 {
		for _curItem, element := range m.PortMappings {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PortMappings), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NLMInitializeRoutingTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMInitializeRoutingTableParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMInitializeRoutingTable, error) {
	return NLMInitializeRoutingTableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMInitializeRoutingTableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMInitializeRoutingTable, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMInitializeRoutingTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMInitializeRoutingTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfPorts)
	_numberOfPorts, _numberOfPortsErr := readBuffer.ReadUint8("numberOfPorts", 8)
	if _numberOfPortsErr != nil {
		return nil, errors.Wrap(_numberOfPortsErr, "Error parsing 'numberOfPorts' field of NLMInitializeRoutingTable")
	}
	numberOfPorts := _numberOfPorts

	// Array field (portMappings)
	if pullErr := readBuffer.PullContext("portMappings", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for portMappings")
	}
	// Count array
	portMappings := make([]NLMInitializeRoutingTablePortMapping, max(numberOfPorts, 0))
	// This happens when the size is set conditional to 0
	if len(portMappings) == 0 {
		portMappings = nil
	}
	{
		_numItems := uint16(max(numberOfPorts, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := NLMInitializeRoutingTablePortMappingParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'portMappings' field of NLMInitializeRoutingTable")
			}
			portMappings[_curItem] = _item.(NLMInitializeRoutingTablePortMapping)
		}
	}
	if closeErr := readBuffer.CloseContext("portMappings", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for portMappings")
	}

	if closeErr := readBuffer.CloseContext("NLMInitializeRoutingTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMInitializeRoutingTable")
	}

	// Create a partially initialized instance
	_child := &_NLMInitializeRoutingTable{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMInitializeRoutingTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMInitializeRoutingTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMInitializeRoutingTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMInitializeRoutingTable")
		}

		// Simple Field (numberOfPorts)
		numberOfPorts := uint8(m.GetNumberOfPorts())
		_numberOfPortsErr := writeBuffer.WriteUint8("numberOfPorts", 8, uint8((numberOfPorts)))
		if _numberOfPortsErr != nil {
			return errors.Wrap(_numberOfPortsErr, "Error serializing 'numberOfPorts' field")
		}

		// Array Field (portMappings)
		if pushErr := writeBuffer.PushContext("portMappings", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for portMappings")
		}
		for _curItem, _element := range m.GetPortMappings() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetPortMappings()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'portMappings' field")
			}
		}
		if popErr := writeBuffer.PopContext("portMappings", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for portMappings")
		}

		if popErr := writeBuffer.PopContext("NLMInitializeRoutingTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMInitializeRoutingTable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMInitializeRoutingTable) isNLMInitializeRoutingTable() bool {
	return true
}

func (m *_NLMInitializeRoutingTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}