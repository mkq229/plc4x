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

// APDUSegmentAck is the corresponding interface of APDUSegmentAck
type APDUSegmentAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetNegativeAck returns NegativeAck (property field)
	GetNegativeAck() bool
	// GetServer returns Server (property field)
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetActualWindowSize returns ActualWindowSize (property field)
	GetActualWindowSize() uint8
}

// APDUSegmentAckExactly can be used when we want exactly this type and not a type which fulfills APDUSegmentAck.
// This is useful for switch cases.
type APDUSegmentAckExactly interface {
	APDUSegmentAck
	isAPDUSegmentAck() bool
}

// _APDUSegmentAck is the data-structure of this message
type _APDUSegmentAck struct {
	*_APDU
	NegativeAck      bool
	Server           bool
	OriginalInvokeId uint8
	SequenceNumber   uint8
	ActualWindowSize uint8
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUSegmentAck) GetApduType() ApduType {
	return ApduType_SEGMENT_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUSegmentAck) InitializeParent(parent APDU) {}

func (m *_APDUSegmentAck) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUSegmentAck) GetNegativeAck() bool {
	return m.NegativeAck
}

func (m *_APDUSegmentAck) GetServer() bool {
	return m.Server
}

func (m *_APDUSegmentAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUSegmentAck) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *_APDUSegmentAck) GetActualWindowSize() uint8 {
	return m.ActualWindowSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUSegmentAck factory function for _APDUSegmentAck
func NewAPDUSegmentAck(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, actualWindowSize uint8, apduLength uint16) *_APDUSegmentAck {
	_result := &_APDUSegmentAck{
		NegativeAck:      negativeAck,
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		SequenceNumber:   sequenceNumber,
		ActualWindowSize: actualWindowSize,
		_APDU:            NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUSegmentAck(structType any) APDUSegmentAck {
	if casted, ok := structType.(APDUSegmentAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUSegmentAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUSegmentAck) GetTypeName() string {
	return "APDUSegmentAck"
}

func (m *_APDUSegmentAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (negativeAck)
	lengthInBits += 1

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Simple field (actualWindowSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *_APDUSegmentAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUSegmentAckParse(ctx context.Context, theBytes []byte, apduLength uint16) (APDUSegmentAck, error) {
	return APDUSegmentAckParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUSegmentAckParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUSegmentAck, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("APDUSegmentAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUSegmentAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of APDUSegmentAck")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (negativeAck)
	_negativeAck, _negativeAckErr := /*TODO: migrate me*/ readBuffer.ReadBit("negativeAck")
	if _negativeAckErr != nil {
		return nil, errors.Wrap(_negativeAckErr, "Error parsing 'negativeAck' field of APDUSegmentAck")
	}
	negativeAck := _negativeAck

	// Simple Field (server)
	_server, _serverErr := /*TODO: migrate me*/ readBuffer.ReadBit("server")
	if _serverErr != nil {
		return nil, errors.Wrap(_serverErr, "Error parsing 'server' field of APDUSegmentAck")
	}
	server := _server

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := /*TODO: migrate me*/ readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field of APDUSegmentAck")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := /*TODO: migrate me*/ readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of APDUSegmentAck")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (actualWindowSize)
	_actualWindowSize, _actualWindowSizeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("actualWindowSize", 8)
	if _actualWindowSizeErr != nil {
		return nil, errors.Wrap(_actualWindowSizeErr, "Error parsing 'actualWindowSize' field of APDUSegmentAck")
	}
	actualWindowSize := _actualWindowSize

	if closeErr := readBuffer.CloseContext("APDUSegmentAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUSegmentAck")
	}

	// Create a partially initialized instance
	_child := &_APDUSegmentAck{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		NegativeAck:      negativeAck,
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		SequenceNumber:   sequenceNumber,
		ActualWindowSize: actualWindowSize,
		reservedField0:   reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUSegmentAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUSegmentAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSegmentAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUSegmentAck")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 2, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (negativeAck)
		negativeAck := bool(m.GetNegativeAck())
		_negativeAckErr := /*TODO: migrate me*/ writeBuffer.WriteBit("negativeAck", (negativeAck))
		if _negativeAckErr != nil {
			return errors.Wrap(_negativeAckErr, "Error serializing 'negativeAck' field")
		}

		// Simple Field (server)
		server := bool(m.GetServer())
		_serverErr := /*TODO: migrate me*/ writeBuffer.WriteBit("server", (server))
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.GetOriginalInvokeId())
		_originalInvokeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("originalInvokeId", 8, uint8((originalInvokeId)))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.GetSequenceNumber())
		_sequenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("sequenceNumber", 8, uint8((sequenceNumber)))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (actualWindowSize)
		actualWindowSize := uint8(m.GetActualWindowSize())
		_actualWindowSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("actualWindowSize", 8, uint8((actualWindowSize)))
		if _actualWindowSizeErr != nil {
			return errors.Wrap(_actualWindowSizeErr, "Error serializing 'actualWindowSize' field")
		}

		if popErr := writeBuffer.PopContext("APDUSegmentAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUSegmentAck")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUSegmentAck) isAPDUSegmentAck() bool {
	return true
}

func (m *_APDUSegmentAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
