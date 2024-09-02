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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MessagePDU is the corresponding interface of MessagePDU
type MessagePDU interface {
	MessagePDUContract
	MessagePDURequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// MessagePDUContract provides a set of functions which can be overwritten by a sub struct
type MessagePDUContract interface {
	// GetChunk returns Chunk (property field)
	GetChunk() ChunkType
}

// MessagePDURequirements provides a set of functions which need to be implemented by a sub struct
type MessagePDURequirements interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() string
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// MessagePDUExactly can be used when we want exactly this type and not a type which fulfills MessagePDU.
// This is useful for switch cases.
type MessagePDUExactly interface {
	MessagePDU
	isMessagePDU() bool
}

// _MessagePDU is the data-structure of this message
type _MessagePDU struct {
	_MessagePDUChildRequirements
	Chunk ChunkType
}

var _ MessagePDUContract = (*_MessagePDU)(nil)

type _MessagePDUChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() string
	GetResponse() bool
}

type MessagePDUChild interface {
	utils.Serializable

	GetParent() *MessagePDU

	GetTypeName() string
	MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MessagePDU) GetChunk() ChunkType {
	return m.Chunk
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMessagePDU factory function for _MessagePDU
func NewMessagePDU(chunk ChunkType) *_MessagePDU {
	return &_MessagePDU{Chunk: chunk}
}

// Deprecated: use the interface for direct cast
func CastMessagePDU(structType any) MessagePDU {
	if casted, ok := structType.(MessagePDU); ok {
		return casted
	}
	if casted, ok := structType.(*MessagePDU); ok {
		return *casted
	}
	return nil
}

func (m *_MessagePDU) GetTypeName() string {
	return "MessagePDU"
}

func (m *_MessagePDU) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 24

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (totalLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MessagePDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MessagePDUParse[T MessagePDU](ctx context.Context, theBytes []byte, response bool) (T, error) {
	return MessagePDUParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func MessagePDUParseWithBufferProducer[T MessagePDU](response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := MessagePDUParseWithBuffer[T](ctx, readBuffer, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func MessagePDUParseWithBuffer[T MessagePDU](ctx context.Context, readBuffer utils.ReadBuffer, response bool) (T, error) {
	v, err := (&_MessagePDU{}).parse(ctx, readBuffer, response)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_MessagePDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (__messagePDU MessagePDU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MessagePDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MessagePDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageType, err := ReadDiscriminatorField[string](ctx, "messageType", ReadString(readBuffer, uint32(24)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	chunk, err := ReadEnumField[ChunkType](ctx, "chunk", "ChunkType", ReadEnum(ChunkTypeByValue, ReadString(readBuffer, uint32(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chunk' field"))
	}
	m.Chunk = chunk

	totalLength, err := ReadImplicitField[uint32](ctx, "totalLength", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'totalLength' field"))
	}
	_ = totalLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child MessagePDU
	switch {
	case messageType == "HEL" && response == bool(false): // OpcuaHelloRequest
		if _child, err = (&_OpcuaHelloRequest{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaHelloRequest for type-switch of MessagePDU")
		}
	case messageType == "ACK" && response == bool(true): // OpcuaAcknowledgeResponse
		if _child, err = (&_OpcuaAcknowledgeResponse{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaAcknowledgeResponse for type-switch of MessagePDU")
		}
	case messageType == "OPN" && response == bool(false): // OpcuaOpenRequest
		if _child, err = (&_OpcuaOpenRequest{}).parse(ctx, readBuffer, m, totalLength, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaOpenRequest for type-switch of MessagePDU")
		}
	case messageType == "OPN" && response == bool(true): // OpcuaOpenResponse
		if _child, err = (&_OpcuaOpenResponse{}).parse(ctx, readBuffer, m, totalLength, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaOpenResponse for type-switch of MessagePDU")
		}
	case messageType == "CLO" && response == bool(false): // OpcuaCloseRequest
		if _child, err = (&_OpcuaCloseRequest{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaCloseRequest for type-switch of MessagePDU")
		}
	case messageType == "MSG" && response == bool(false): // OpcuaMessageRequest
		if _child, err = (&_OpcuaMessageRequest{}).parse(ctx, readBuffer, m, totalLength, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaMessageRequest for type-switch of MessagePDU")
		}
	case messageType == "MSG" && response == bool(true): // OpcuaMessageResponse
		if _child, err = (&_OpcuaMessageResponse{}).parse(ctx, readBuffer, m, totalLength, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaMessageResponse for type-switch of MessagePDU")
		}
	case messageType == "ERR" && response == bool(true): // OpcuaMessageError
		if _child, err = (&_OpcuaMessageError{}).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type OpcuaMessageError for type-switch of MessagePDU")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageType=%v, response=%v]", messageType, response)
	}

	if closeErr := readBuffer.CloseContext("MessagePDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MessagePDU")
	}

	return _child, nil
}

func (pm *_MessagePDU) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MessagePDU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MessagePDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MessagePDU")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteString(writeBuffer, 24)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}

	if err := WriteSimpleEnumField[ChunkType](ctx, "chunk", "ChunkType", m.GetChunk(), WriteEnum[ChunkType, string](ChunkType.GetValue, ChunkType.PLC4XEnumName, WriteString(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'chunk' field")
	}
	totalLength := uint32(uint32(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "totalLength", totalLength, WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MessagePDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MessagePDU")
	}
	return nil
}

func (m *_MessagePDU) isMessagePDU() bool {
	return true
}
