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

// OpcuaHelloRequest is the corresponding interface of OpcuaHelloRequest
type OpcuaHelloRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetVersion returns Version (property field)
	GetVersion() uint32
	// GetLimits returns Limits (property field)
	GetLimits() OpcuaProtocolLimits
	// GetEndpoint returns Endpoint (property field)
	GetEndpoint() PascalString
}

// OpcuaHelloRequestExactly can be used when we want exactly this type and not a type which fulfills OpcuaHelloRequest.
// This is useful for switch cases.
type OpcuaHelloRequestExactly interface {
	OpcuaHelloRequest
	isOpcuaHelloRequest() bool
}

// _OpcuaHelloRequest is the data-structure of this message
type _OpcuaHelloRequest struct {
	MessagePDUContract
	Version  uint32
	Limits   OpcuaProtocolLimits
	Endpoint PascalString
}

var _ OpcuaHelloRequest = (*_OpcuaHelloRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaHelloRequest) GetMessageType() string {
	return "HEL"
}

func (m *_OpcuaHelloRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaHelloRequest) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaHelloRequest) GetVersion() uint32 {
	return m.Version
}

func (m *_OpcuaHelloRequest) GetLimits() OpcuaProtocolLimits {
	return m.Limits
}

func (m *_OpcuaHelloRequest) GetEndpoint() PascalString {
	return m.Endpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaHelloRequest factory function for _OpcuaHelloRequest
func NewOpcuaHelloRequest(version uint32, limits OpcuaProtocolLimits, endpoint PascalString, chunk ChunkType) *_OpcuaHelloRequest {
	_result := &_OpcuaHelloRequest{
		MessagePDUContract: NewMessagePDU(chunk),
		Version:            version,
		Limits:             limits,
		Endpoint:           endpoint,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaHelloRequest(structType any) OpcuaHelloRequest {
	if casted, ok := structType.(OpcuaHelloRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaHelloRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaHelloRequest) GetTypeName() string {
	return "OpcuaHelloRequest"
}

func (m *_OpcuaHelloRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 32

	// Simple field (limits)
	lengthInBits += m.Limits.GetLengthInBits(ctx)

	// Simple field (endpoint)
	lengthInBits += m.Endpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaHelloRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaHelloRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, response bool) (__opcuaHelloRequest OpcuaHelloRequest, err error) {
	m.MessagePDUContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaHelloRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaHelloRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	limits, err := ReadSimpleField[OpcuaProtocolLimits](ctx, "limits", ReadComplex[OpcuaProtocolLimits](OpcuaProtocolLimitsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limits' field"))
	}
	m.Limits = limits

	endpoint, err := ReadSimpleField[PascalString](ctx, "endpoint", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpoint' field"))
	}
	m.Endpoint = endpoint

	if closeErr := readBuffer.CloseContext("OpcuaHelloRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaHelloRequest")
	}

	return m, nil
}

func (m *_OpcuaHelloRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaHelloRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaHelloRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaHelloRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "version", m.GetVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if err := WriteSimpleField[OpcuaProtocolLimits](ctx, "limits", m.GetLimits(), WriteComplex[OpcuaProtocolLimits](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limits' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpoint", m.GetEndpoint(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpoint' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaHelloRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaHelloRequest")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaHelloRequest) isOpcuaHelloRequest() bool {
	return true
}

func (m *_OpcuaHelloRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
