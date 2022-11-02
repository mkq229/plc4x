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

// KnxNetRemoteLogging is the corresponding interface of KnxNetRemoteLogging
type KnxNetRemoteLogging interface {
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetRemoteLoggingExactly can be used when we want exactly this type and not a type which fulfills KnxNetRemoteLogging.
// This is useful for switch cases.
type KnxNetRemoteLoggingExactly interface {
	KnxNetRemoteLogging
	isKnxNetRemoteLogging() bool
}

// _KnxNetRemoteLogging is the data-structure of this message
type _KnxNetRemoteLogging struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetRemoteLogging) GetServiceType() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetRemoteLogging) InitializeParent(parent ServiceId) {}

func (m *_KnxNetRemoteLogging) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetRemoteLogging) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetRemoteLogging factory function for _KnxNetRemoteLogging
func NewKnxNetRemoteLogging(version uint8) *_KnxNetRemoteLogging {
	_result := &_KnxNetRemoteLogging{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetRemoteLogging(structType interface{}) KnxNetRemoteLogging {
	if casted, ok := structType.(KnxNetRemoteLogging); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetRemoteLogging); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetRemoteLogging) GetTypeName() string {
	return "KnxNetRemoteLogging"
}

func (m *_KnxNetRemoteLogging) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxNetRemoteLogging) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetRemoteLogging) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetRemoteLoggingParse(readBuffer utils.ReadBuffer) (KnxNetRemoteLogging, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetRemoteLogging"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetRemoteLogging")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of KnxNetRemoteLogging")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetRemoteLogging"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetRemoteLogging")
	}

	// Create a partially initialized instance
	_child := &_KnxNetRemoteLogging{
		_ServiceId: &_ServiceId{},
		Version:    version,
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetRemoteLogging) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetRemoteLogging) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetRemoteLogging"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetRemoteLogging")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetRemoteLogging"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetRemoteLogging")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxNetRemoteLogging) isKnxNetRemoteLogging() bool {
	return true
}

func (m *_KnxNetRemoteLogging) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}