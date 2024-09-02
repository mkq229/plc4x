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

// ReplyOrConfirmationReply is the corresponding interface of ReplyOrConfirmationReply
type ReplyOrConfirmationReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ReplyOrConfirmation
	// GetReply returns Reply (property field)
	GetReply() Reply
	// GetTermination returns Termination (property field)
	GetTermination() ResponseTermination
}

// ReplyOrConfirmationReplyExactly can be used when we want exactly this type and not a type which fulfills ReplyOrConfirmationReply.
// This is useful for switch cases.
type ReplyOrConfirmationReplyExactly interface {
	ReplyOrConfirmationReply
	isReplyOrConfirmationReply() bool
}

// _ReplyOrConfirmationReply is the data-structure of this message
type _ReplyOrConfirmationReply struct {
	ReplyOrConfirmationContract
	Reply       Reply
	Termination ResponseTermination
}

var _ ReplyOrConfirmationReply = (*_ReplyOrConfirmationReply)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyOrConfirmationReply) GetParent() ReplyOrConfirmationContract {
	return m.ReplyOrConfirmationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmationReply) GetReply() Reply {
	return m.Reply
}

func (m *_ReplyOrConfirmationReply) GetTermination() ResponseTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyOrConfirmationReply factory function for _ReplyOrConfirmationReply
func NewReplyOrConfirmationReply(reply Reply, termination ResponseTermination, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmationReply {
	_result := &_ReplyOrConfirmationReply{
		ReplyOrConfirmationContract: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
		Reply:                       reply,
		Termination:                 termination,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmationReply(structType any) ReplyOrConfirmationReply {
	if casted, ok := structType.(ReplyOrConfirmationReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmationReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmationReply) GetTypeName() string {
	return "ReplyOrConfirmationReply"
}

func (m *_ReplyOrConfirmationReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).getLengthInBits(ctx))

	// Simple field (reply)
	lengthInBits += m.Reply.GetLengthInBits(ctx)

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReplyOrConfirmationReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReplyOrConfirmationReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ReplyOrConfirmation, cBusOptions CBusOptions, requestContext RequestContext) (__replyOrConfirmationReply ReplyOrConfirmationReply, err error) {
	m.ReplyOrConfirmationContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyOrConfirmationReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmationReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reply, err := ReadSimpleField[Reply](ctx, "reply", ReadComplex[Reply](ReplyParseWithBufferProducer[Reply]((CBusOptions)(cBusOptions), (RequestContext)(requestContext)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reply' field"))
	}
	m.Reply = reply

	termination, err := ReadSimpleField[ResponseTermination](ctx, "termination", ReadComplex[ResponseTermination](ResponseTerminationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'termination' field"))
	}
	m.Termination = termination

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmationReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmationReply")
	}

	return m, nil
}

func (m *_ReplyOrConfirmationReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyOrConfirmationReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyOrConfirmationReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmationReply")
		}

		if err := WriteSimpleField[Reply](ctx, "reply", m.GetReply(), WriteComplex[Reply](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reply' field")
		}

		if err := WriteSimpleField[ResponseTermination](ctx, "termination", m.GetTermination(), WriteComplex[ResponseTermination](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'termination' field")
		}

		if popErr := writeBuffer.PopContext("ReplyOrConfirmationReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyOrConfirmationReply")
		}
		return nil
	}
	return m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReplyOrConfirmationReply) isReplyOrConfirmationReply() bool {
	return true
}

func (m *_ReplyOrConfirmationReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
