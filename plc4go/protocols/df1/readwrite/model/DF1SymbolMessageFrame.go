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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const DF1SymbolMessageFrame_MESSAGEEND uint8 = 0x10
const DF1SymbolMessageFrame_ENDTRANSACTION uint8 = 0x03

// DF1SymbolMessageFrame is the corresponding interface of DF1SymbolMessageFrame
type DF1SymbolMessageFrame interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1Symbol
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() uint8
	// GetCommand returns Command (property field)
	GetCommand() DF1Command
}

// DF1SymbolMessageFrameExactly can be used when we want exactly this type and not a type which fulfills DF1SymbolMessageFrame.
// This is useful for switch cases.
type DF1SymbolMessageFrameExactly interface {
	DF1SymbolMessageFrame
	isDF1SymbolMessageFrame() bool
}

// _DF1SymbolMessageFrame is the data-structure of this message
type _DF1SymbolMessageFrame struct {
	*_DF1Symbol
	DestinationAddress uint8
	SourceAddress      uint8
	Command            DF1Command
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetSymbolType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrame) InitializeParent(parent DF1Symbol) {}

func (m *_DF1SymbolMessageFrame) GetParent() DF1Symbol {
	return m._DF1Symbol
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetDestinationAddress() uint8 {
	return m.DestinationAddress
}

func (m *_DF1SymbolMessageFrame) GetSourceAddress() uint8 {
	return m.SourceAddress
}

func (m *_DF1SymbolMessageFrame) GetCommand() DF1Command {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetMessageEnd() uint8 {
	return DF1SymbolMessageFrame_MESSAGEEND
}

func (m *_DF1SymbolMessageFrame) GetEndTransaction() uint8 {
	return DF1SymbolMessageFrame_ENDTRANSACTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1SymbolMessageFrame factory function for _DF1SymbolMessageFrame
func NewDF1SymbolMessageFrame(destinationAddress uint8, sourceAddress uint8, command DF1Command) *_DF1SymbolMessageFrame {
	_result := &_DF1SymbolMessageFrame{
		DestinationAddress: destinationAddress,
		SourceAddress:      sourceAddress,
		Command:            command,
		_DF1Symbol:         NewDF1Symbol(),
	}
	_result._DF1Symbol._DF1SymbolChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrame(structType any) DF1SymbolMessageFrame {
	if casted, ok := structType.(DF1SymbolMessageFrame); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrame); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrame) GetTypeName() string {
	return "DF1SymbolMessageFrame"
}

func (m *_DF1SymbolMessageFrame) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	// Const Field (messageEnd)
	lengthInBits += 8

	// Const Field (endTransaction)
	lengthInBits += 8

	// Checksum Field (checksum)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1SymbolMessageFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1SymbolMessageFrameParse(ctx context.Context, theBytes []byte) (DF1SymbolMessageFrame, error) {
	return DF1SymbolMessageFrameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DF1SymbolMessageFrameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1SymbolMessageFrame, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationAddress)
	_destinationAddress, _destinationAddressErr := /*TODO: migrate me*/ readBuffer.ReadUint8("destinationAddress", 8)
	if _destinationAddressErr != nil {
		return nil, errors.Wrap(_destinationAddressErr, "Error parsing 'destinationAddress' field of DF1SymbolMessageFrame")
	}
	destinationAddress := _destinationAddress

	// Simple Field (sourceAddress)
	_sourceAddress, _sourceAddressErr := /*TODO: migrate me*/ readBuffer.ReadUint8("sourceAddress", 8)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field of DF1SymbolMessageFrame")
	}
	sourceAddress := _sourceAddress

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
	_command, _commandErr := DF1CommandParseWithBuffer(ctx, readBuffer)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of DF1SymbolMessageFrame")
	}
	command := _command.(DF1Command)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	messageEnd, err := ReadConstField[uint8](ctx, "messageEnd", ReadUnsignedByte(readBuffer, 8), DF1SymbolMessageFrame_MESSAGEEND, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageEnd' field"))
	}
	_ = messageEnd

	endTransaction, err := ReadConstField[uint8](ctx, "endTransaction", ReadUnsignedByte(readBuffer, 8), DF1SymbolMessageFrame_ENDTRANSACTION, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endTransaction' field"))
	}
	_ = endTransaction

	crc, err := ReadChecksumField[uint16](ctx, "crc", ReadUnsignedShort(readBuffer, 16), CrcCheck(ctx, destinationAddress, sourceAddress, command), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	_ = crc

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrame")
	}

	// Create a partially initialized instance
	_child := &_DF1SymbolMessageFrame{
		_DF1Symbol:         &_DF1Symbol{},
		DestinationAddress: destinationAddress,
		SourceAddress:      sourceAddress,
		Command:            command,
	}
	_child._DF1Symbol._DF1SymbolChildRequirements = _child
	return _child, nil
}

func (m *_DF1SymbolMessageFrame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrame) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrame"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrame")
		}

		// Simple Field (destinationAddress)
		destinationAddress := uint8(m.GetDestinationAddress())
		_destinationAddressErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("destinationAddress", 8, uint8((destinationAddress)))
		if _destinationAddressErr != nil {
			return errors.Wrap(_destinationAddressErr, "Error serializing 'destinationAddress' field")
		}

		// Simple Field (sourceAddress)
		sourceAddress := uint8(m.GetSourceAddress())
		_sourceAddressErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("sourceAddress", 8, uint8((sourceAddress)))
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for command")
		}
		_commandErr := writeBuffer.WriteSerializable(ctx, m.GetCommand())
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for command")
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		// Const Field (messageEnd)
		_messageEndErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint8("messageEnd", 8, uint8(0x10))
		if _messageEndErr != nil {
			return errors.Wrap(_messageEndErr, "Error serializing 'messageEnd' field")
		}

		// Const Field (endTransaction)
		_endTransactionErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint8("endTransaction", 8, uint8(0x03))
		if _endTransactionErr != nil {
			return errors.Wrap(_endTransactionErr, "Error serializing 'endTransaction' field")
		}

		// Checksum Field (checksum) (Calculated)
		if err := WriteChecksumField[uint16](ctx, "crc", CrcCheck(ctx, m.GetDestinationAddress(), m.GetSourceAddress(), m.GetCommand()), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrame"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrame")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1SymbolMessageFrame) isDF1SymbolMessageFrame() bool {
	return true
}

func (m *_DF1SymbolMessageFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
