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

// BACnetVTSession is the corresponding interface of BACnetVTSession
type BACnetVTSession interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLocalVtSessionId returns LocalVtSessionId (property field)
	GetLocalVtSessionId() BACnetApplicationTagUnsignedInteger
	// GetRemoveVtSessionId returns RemoveVtSessionId (property field)
	GetRemoveVtSessionId() BACnetApplicationTagUnsignedInteger
	// GetRemoteVtAddress returns RemoteVtAddress (property field)
	GetRemoteVtAddress() BACnetAddress
}

// BACnetVTSessionExactly can be used when we want exactly this type and not a type which fulfills BACnetVTSession.
// This is useful for switch cases.
type BACnetVTSessionExactly interface {
	BACnetVTSession
	isBACnetVTSession() bool
}

// _BACnetVTSession is the data-structure of this message
type _BACnetVTSession struct {
	LocalVtSessionId  BACnetApplicationTagUnsignedInteger
	RemoveVtSessionId BACnetApplicationTagUnsignedInteger
	RemoteVtAddress   BACnetAddress
}

var _ BACnetVTSession = (*_BACnetVTSession)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetVTSession) GetLocalVtSessionId() BACnetApplicationTagUnsignedInteger {
	return m.LocalVtSessionId
}

func (m *_BACnetVTSession) GetRemoveVtSessionId() BACnetApplicationTagUnsignedInteger {
	return m.RemoveVtSessionId
}

func (m *_BACnetVTSession) GetRemoteVtAddress() BACnetAddress {
	return m.RemoteVtAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetVTSession factory function for _BACnetVTSession
func NewBACnetVTSession(localVtSessionId BACnetApplicationTagUnsignedInteger, removeVtSessionId BACnetApplicationTagUnsignedInteger, remoteVtAddress BACnetAddress) *_BACnetVTSession {
	return &_BACnetVTSession{LocalVtSessionId: localVtSessionId, RemoveVtSessionId: removeVtSessionId, RemoteVtAddress: remoteVtAddress}
}

// Deprecated: use the interface for direct cast
func CastBACnetVTSession(structType any) BACnetVTSession {
	if casted, ok := structType.(BACnetVTSession); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetVTSession); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetVTSession) GetTypeName() string {
	return "BACnetVTSession"
}

func (m *_BACnetVTSession) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (localVtSessionId)
	lengthInBits += m.LocalVtSessionId.GetLengthInBits(ctx)

	// Simple field (removeVtSessionId)
	lengthInBits += m.RemoveVtSessionId.GetLengthInBits(ctx)

	// Simple field (remoteVtAddress)
	lengthInBits += m.RemoteVtAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetVTSession) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetVTSessionParse(ctx context.Context, theBytes []byte) (BACnetVTSession, error) {
	return BACnetVTSessionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetVTSessionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVTSession, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVTSession, error) {
		return BACnetVTSessionParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetVTSessionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVTSession, error) {
	v, err := (&_BACnetVTSession{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetVTSession) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetVTSession BACnetVTSession, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetVTSession"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetVTSession")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	localVtSessionId, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "localVtSessionId", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localVtSessionId' field"))
	}
	m.LocalVtSessionId = localVtSessionId

	removeVtSessionId, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "removeVtSessionId", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'removeVtSessionId' field"))
	}
	m.RemoveVtSessionId = removeVtSessionId

	remoteVtAddress, err := ReadSimpleField[BACnetAddress](ctx, "remoteVtAddress", ReadComplex[BACnetAddress](BACnetAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remoteVtAddress' field"))
	}
	m.RemoteVtAddress = remoteVtAddress

	if closeErr := readBuffer.CloseContext("BACnetVTSession"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetVTSession")
	}

	return m, nil
}

func (m *_BACnetVTSession) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetVTSession) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetVTSession"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetVTSession")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "localVtSessionId", m.GetLocalVtSessionId(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localVtSessionId' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "removeVtSessionId", m.GetRemoveVtSessionId(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'removeVtSessionId' field")
	}

	if err := WriteSimpleField[BACnetAddress](ctx, "remoteVtAddress", m.GetRemoteVtAddress(), WriteComplex[BACnetAddress](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'remoteVtAddress' field")
	}

	if popErr := writeBuffer.PopContext("BACnetVTSession"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetVTSession")
	}
	return nil
}

func (m *_BACnetVTSession) isBACnetVTSession() bool {
	return true
}

func (m *_BACnetVTSession) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
