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

// BACnetServiceAckAtomicWriteFile is the corresponding interface of BACnetServiceAckAtomicWriteFile
type BACnetServiceAckAtomicWriteFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetContextTagSignedInteger
}

// BACnetServiceAckAtomicWriteFileExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAtomicWriteFile.
// This is useful for switch cases.
type BACnetServiceAckAtomicWriteFileExactly interface {
	BACnetServiceAckAtomicWriteFile
	isBACnetServiceAckAtomicWriteFile() bool
}

// _BACnetServiceAckAtomicWriteFile is the data-structure of this message
type _BACnetServiceAckAtomicWriteFile struct {
	*_BACnetServiceAck
	FileStartPosition BACnetContextTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckAtomicWriteFile) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetFileStartPosition() BACnetContextTagSignedInteger {
	return m.FileStartPosition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicWriteFile factory function for _BACnetServiceAckAtomicWriteFile
func NewBACnetServiceAckAtomicWriteFile(fileStartPosition BACnetContextTagSignedInteger, serviceAckLength uint32) *_BACnetServiceAckAtomicWriteFile {
	_result := &_BACnetServiceAckAtomicWriteFile{
		FileStartPosition: fileStartPosition,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicWriteFile(structType any) BACnetServiceAckAtomicWriteFile {
	if casted, ok := structType.(BACnetServiceAckAtomicWriteFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicWriteFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicWriteFile) GetTypeName() string {
	return "BACnetServiceAckAtomicWriteFile"
}

func (m *_BACnetServiceAckAtomicWriteFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicWriteFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckAtomicWriteFileParse(ctx context.Context, theBytes []byte, serviceAckLength uint32) (BACnetServiceAckAtomicWriteFile, error) {
	return BACnetServiceAckAtomicWriteFileParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckAtomicWriteFileParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckAtomicWriteFile, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicWriteFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicWriteFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileStartPosition")
	}
	_fileStartPosition, _fileStartPositionErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field of BACnetServiceAckAtomicWriteFile")
	}
	fileStartPosition := _fileStartPosition.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileStartPosition")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicWriteFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicWriteFile")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckAtomicWriteFile{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		FileStartPosition: fileStartPosition,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckAtomicWriteFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicWriteFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicWriteFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicWriteFile")
		}

		// Simple Field (fileStartPosition)
		if pushErr := writeBuffer.PushContext("fileStartPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileStartPosition")
		}
		_fileStartPositionErr := writeBuffer.WriteSerializable(ctx, m.GetFileStartPosition())
		if popErr := writeBuffer.PopContext("fileStartPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileStartPosition")
		}
		if _fileStartPositionErr != nil {
			return errors.Wrap(_fileStartPositionErr, "Error serializing 'fileStartPosition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicWriteFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicWriteFile")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicWriteFile) isBACnetServiceAckAtomicWriteFile() bool {
	return true
}

func (m *_BACnetServiceAckAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
