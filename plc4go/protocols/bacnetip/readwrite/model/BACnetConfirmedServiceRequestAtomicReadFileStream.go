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

// BACnetConfirmedServiceRequestAtomicReadFileStream is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileStream
type BACnetConfirmedServiceRequestAtomicReadFileStream interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetApplicationTagSignedInteger
	// GetRequestOctetCount returns RequestOctetCount (property field)
	GetRequestOctetCount() BACnetApplicationTagUnsignedInteger
}

// BACnetConfirmedServiceRequestAtomicReadFileStreamExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestAtomicReadFileStream.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestAtomicReadFileStreamExactly interface {
	BACnetConfirmedServiceRequestAtomicReadFileStream
	isBACnetConfirmedServiceRequestAtomicReadFileStream() bool
}

// _BACnetConfirmedServiceRequestAtomicReadFileStream is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileStream struct {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
	FileStartPosition BACnetApplicationTagSignedInteger
	RequestOctetCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestAtomicReadFileStream = (*_BACnetConfirmedServiceRequestAtomicReadFileStream)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetParent() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract {
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetFileStartPosition() BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetRequestOctetCount() BACnetApplicationTagUnsignedInteger {
	return m.RequestOctetCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicReadFileStream factory function for _BACnetConfirmedServiceRequestAtomicReadFileStream
func NewBACnetConfirmedServiceRequestAtomicReadFileStream(fileStartPosition BACnetApplicationTagSignedInteger, requestOctetCount BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestAtomicReadFileStream {
	_result := &_BACnetConfirmedServiceRequestAtomicReadFileStream{
		BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract: NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
		FileStartPosition: fileStartPosition,
		RequestOctetCount: requestOctetCount,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFileStream(structType any) BACnetConfirmedServiceRequestAtomicReadFileStream {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileStream); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileStream); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileStream"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).getLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	// Simple field (requestOctetCount)
	lengthInBits += m.RequestOctetCount.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) (__bACnetConfirmedServiceRequestAtomicReadFileStream BACnetConfirmedServiceRequestAtomicReadFileStream, err error) {
	m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileStream")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartPosition, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	requestOctetCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestOctetCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestOctetCount' field"))
	}
	m.RequestOctetCount = requestOctetCount

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileStream")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileStream")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestOctetCount", m.GetRequestOctetCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestOctetCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileStream")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) isBACnetConfirmedServiceRequestAtomicReadFileStream() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
