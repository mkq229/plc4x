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

// KnxNetIpTunneling is the corresponding interface of KnxNetIpTunneling
type KnxNetIpTunneling interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpTunnelingExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpTunneling.
// This is useful for switch cases.
type KnxNetIpTunnelingExactly interface {
	KnxNetIpTunneling
	isKnxNetIpTunneling() bool
}

// _KnxNetIpTunneling is the data-structure of this message
type _KnxNetIpTunneling struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetIpTunneling = (*_KnxNetIpTunneling)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpTunneling) GetServiceType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpTunneling) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpTunneling) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpTunneling factory function for _KnxNetIpTunneling
func NewKnxNetIpTunneling(version uint8) *_KnxNetIpTunneling {
	_result := &_KnxNetIpTunneling{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpTunneling(structType any) KnxNetIpTunneling {
	if casted, ok := structType.(KnxNetIpTunneling); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpTunneling); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpTunneling) GetTypeName() string {
	return "KnxNetIpTunneling"
}

func (m *_KnxNetIpTunneling) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpTunneling) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetIpTunneling) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetIpTunneling KnxNetIpTunneling, err error) {
	m.ServiceIdContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpTunneling"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpTunneling")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetIpTunneling"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpTunneling")
	}

	return m, nil
}

func (m *_KnxNetIpTunneling) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetIpTunneling) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpTunneling"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpTunneling")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpTunneling"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpTunneling")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetIpTunneling) isKnxNetIpTunneling() bool {
	return true
}

func (m *_KnxNetIpTunneling) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
