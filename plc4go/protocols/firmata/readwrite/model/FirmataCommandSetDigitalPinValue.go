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

// FirmataCommandSetDigitalPinValue is the corresponding interface of FirmataCommandSetDigitalPinValue
type FirmataCommandSetDigitalPinValue interface {
	utils.LengthAware
	utils.Serializable
	FirmataCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetOn returns On (property field)
	GetOn() bool
}

// FirmataCommandSetDigitalPinValueExactly can be used when we want exactly this type and not a type which fulfills FirmataCommandSetDigitalPinValue.
// This is useful for switch cases.
type FirmataCommandSetDigitalPinValueExactly interface {
	FirmataCommandSetDigitalPinValue
	isFirmataCommandSetDigitalPinValue() bool
}

// _FirmataCommandSetDigitalPinValue is the data-structure of this message
type _FirmataCommandSetDigitalPinValue struct {
	*_FirmataCommand
	Pin uint8
	On  bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandSetDigitalPinValue) GetCommandCode() uint8 {
	return 0x5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandSetDigitalPinValue) InitializeParent(parent FirmataCommand) {}

func (m *_FirmataCommandSetDigitalPinValue) GetParent() FirmataCommand {
	return m._FirmataCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandSetDigitalPinValue) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataCommandSetDigitalPinValue) GetOn() bool {
	return m.On
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataCommandSetDigitalPinValue factory function for _FirmataCommandSetDigitalPinValue
func NewFirmataCommandSetDigitalPinValue(pin uint8, on bool, response bool) *_FirmataCommandSetDigitalPinValue {
	_result := &_FirmataCommandSetDigitalPinValue{
		Pin:             pin,
		On:              on,
		_FirmataCommand: NewFirmataCommand(response),
	}
	_result._FirmataCommand._FirmataCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataCommandSetDigitalPinValue(structType interface{}) FirmataCommandSetDigitalPinValue {
	if casted, ok := structType.(FirmataCommandSetDigitalPinValue); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandSetDigitalPinValue); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandSetDigitalPinValue) GetTypeName() string {
	return "FirmataCommandSetDigitalPinValue"
}

func (m *_FirmataCommandSetDigitalPinValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_FirmataCommandSetDigitalPinValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (on)
	lengthInBits += 1

	return lengthInBits
}

func (m *_FirmataCommandSetDigitalPinValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSetDigitalPinValueParse(readBuffer utils.ReadBuffer, response bool) (FirmataCommandSetDigitalPinValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandSetDigitalPinValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandSetDigitalPinValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of FirmataCommandSetDigitalPinValue")
	}
	pin := _pin

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of FirmataCommandSetDigitalPinValue")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (on)
	_on, _onErr := readBuffer.ReadBit("on")
	if _onErr != nil {
		return nil, errors.Wrap(_onErr, "Error parsing 'on' field of FirmataCommandSetDigitalPinValue")
	}
	on := _on

	if closeErr := readBuffer.CloseContext("FirmataCommandSetDigitalPinValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandSetDigitalPinValue")
	}

	// Create a partially initialized instance
	_child := &_FirmataCommandSetDigitalPinValue{
		_FirmataCommand: &_FirmataCommand{
			Response: response,
		},
		Pin:            pin,
		On:             on,
		reservedField0: reservedField0,
	}
	_child._FirmataCommand._FirmataCommandChildRequirements = _child
	return _child, nil
}

func (m *_FirmataCommandSetDigitalPinValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSetDigitalPinValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSetDigitalPinValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSetDigitalPinValue")
		}

		// Simple Field (pin)
		pin := uint8(m.GetPin())
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (on)
		on := bool(m.GetOn())
		_onErr := writeBuffer.WriteBit("on", (on))
		if _onErr != nil {
			return errors.Wrap(_onErr, "Error serializing 'on' field")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSetDigitalPinValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSetDigitalPinValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_FirmataCommandSetDigitalPinValue) isFirmataCommandSetDigitalPinValue() bool {
	return true
}

func (m *_FirmataCommandSetDigitalPinValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}