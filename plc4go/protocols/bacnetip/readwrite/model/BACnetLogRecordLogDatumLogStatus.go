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

// BACnetLogRecordLogDatumLogStatus is the corresponding interface of BACnetLogRecordLogDatumLogStatus
type BACnetLogRecordLogDatumLogStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
}

// BACnetLogRecordLogDatumLogStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumLogStatus.
// This is useful for switch cases.
type BACnetLogRecordLogDatumLogStatusExactly interface {
	BACnetLogRecordLogDatumLogStatus
	isBACnetLogRecordLogDatumLogStatus() bool
}

// _BACnetLogRecordLogDatumLogStatus is the data-structure of this message
type _BACnetLogRecordLogDatumLogStatus struct {
	BACnetLogRecordLogDatumContract
	LogStatus BACnetLogStatusTagged
}

var _ BACnetLogRecordLogDatumLogStatus = (*_BACnetLogRecordLogDatumLogStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumLogStatus) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumLogStatus factory function for _BACnetLogRecordLogDatumLogStatus
func NewBACnetLogRecordLogDatumLogStatus(logStatus BACnetLogStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumLogStatus {
	_result := &_BACnetLogRecordLogDatumLogStatus{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		LogStatus:                       logStatus,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumLogStatus(structType any) BACnetLogRecordLogDatumLogStatus {
	if casted, ok := structType.(BACnetLogRecordLogDatumLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetTypeName() string {
	return "BACnetLogRecordLogDatumLogStatus"
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumLogStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumLogStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumLogStatus BACnetLogRecordLogDatumLogStatus, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logStatus, err := ReadSimpleField[BACnetLogStatusTagged](ctx, "logStatus", ReadComplex[BACnetLogStatusTagged](BACnetLogStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logStatus' field"))
	}
	m.LogStatus = logStatus

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumLogStatus")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumLogStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumLogStatus")
		}

		if err := WriteSimpleField[BACnetLogStatusTagged](ctx, "logStatus", m.GetLogStatus(), WriteComplex[BACnetLogStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumLogStatus")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumLogStatus) isBACnetLogRecordLogDatumLogStatus() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumLogStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
