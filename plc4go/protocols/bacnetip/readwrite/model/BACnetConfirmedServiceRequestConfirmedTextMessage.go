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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestConfirmedTextMessage is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessage
type BACnetConfirmedServiceRequestConfirmedTextMessage interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetTextMessageSourceDevice returns TextMessageSourceDevice (property field)
	GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier
	// GetMessageClass returns MessageClass (property field)
	GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetMessagePriority returns MessagePriority (property field)
	GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	// GetMessage returns Message (property field)
	GetMessage() BACnetContextTagCharacterString
}

// BACnetConfirmedServiceRequestConfirmedTextMessageExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedTextMessage.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedTextMessageExactly interface {
	BACnetConfirmedServiceRequestConfirmedTextMessage
	isBACnetConfirmedServiceRequestConfirmedTextMessage() bool
}

// _BACnetConfirmedServiceRequestConfirmedTextMessage is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessage struct {
	*_BACnetConfirmedServiceRequest
	TextMessageSourceDevice BACnetContextTagObjectIdentifier
	MessageClass            BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	MessagePriority         BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	Message                 BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier {
	return m.TextMessageSourceDevice
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m.MessageClass
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	return m.MessagePriority
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetMessage() BACnetContextTagCharacterString {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedTextMessage factory function for _BACnetConfirmedServiceRequestConfirmedTextMessage
func NewBACnetConfirmedServiceRequestConfirmedTextMessage(textMessageSourceDevice BACnetContextTagObjectIdentifier, messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, message BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedTextMessage {
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessage{
		TextMessageSourceDevice:        textMessageSourceDevice,
		MessageClass:                   messageClass,
		MessagePriority:                messagePriority,
		Message:                        message,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessage(structType interface{}) BACnetConfirmedServiceRequestConfirmedTextMessage {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessage); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessage); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessage"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (textMessageSourceDevice)
	lengthInBits += m.TextMessageSourceDevice.GetLengthInBits()

	// Optional Field (messageClass)
	if m.MessageClass != nil {
		lengthInBits += m.MessageClass.GetLengthInBits()
	}

	// Simple field (messagePriority)
	lengthInBits += m.MessagePriority.GetLengthInBits()

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageParse(readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedTextMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (textMessageSourceDevice)
	if pullErr := readBuffer.PullContext("textMessageSourceDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for textMessageSourceDevice")
	}
	_textMessageSourceDevice, _textMessageSourceDeviceErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _textMessageSourceDeviceErr != nil {
		return nil, errors.Wrap(_textMessageSourceDeviceErr, "Error parsing 'textMessageSourceDevice' field of BACnetConfirmedServiceRequestConfirmedTextMessage")
	}
	textMessageSourceDevice := _textMessageSourceDevice.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("textMessageSourceDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for textMessageSourceDevice")
	}

	// Optional Field (messageClass) (Can be skipped, if a given expression evaluates to false)
	var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("messageClass"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for messageClass")
		}
		_val, _err := BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParse(readBuffer, uint8(1))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'messageClass' field of BACnetConfirmedServiceRequestConfirmedTextMessage")
		default:
			messageClass = _val.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)
			if closeErr := readBuffer.CloseContext("messageClass"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for messageClass")
			}
		}
	}

	// Simple Field (messagePriority)
	if pullErr := readBuffer.PullContext("messagePriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messagePriority")
	}
	_messagePriority, _messagePriorityErr := BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParse(readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _messagePriorityErr != nil {
		return nil, errors.Wrap(_messagePriorityErr, "Error parsing 'messagePriority' field of BACnetConfirmedServiceRequestConfirmedTextMessage")
	}
	messagePriority := _messagePriority.(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged)
	if closeErr := readBuffer.CloseContext("messagePriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messagePriority")
	}

	// Simple Field (message)
	if pullErr := readBuffer.PullContext("message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for message")
	}
	_message, _messageErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field of BACnetConfirmedServiceRequestConfirmedTextMessage")
	}
	message := _message.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for message")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessage")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedTextMessage{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		TextMessageSourceDevice: textMessageSourceDevice,
		MessageClass:            messageClass,
		MessagePriority:         messagePriority,
		Message:                 message,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessage")
		}

		// Simple Field (textMessageSourceDevice)
		if pushErr := writeBuffer.PushContext("textMessageSourceDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for textMessageSourceDevice")
		}
		_textMessageSourceDeviceErr := writeBuffer.WriteSerializable(m.GetTextMessageSourceDevice())
		if popErr := writeBuffer.PopContext("textMessageSourceDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for textMessageSourceDevice")
		}
		if _textMessageSourceDeviceErr != nil {
			return errors.Wrap(_textMessageSourceDeviceErr, "Error serializing 'textMessageSourceDevice' field")
		}

		// Optional Field (messageClass) (Can be skipped, if the value is null)
		var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass = nil
		if m.GetMessageClass() != nil {
			if pushErr := writeBuffer.PushContext("messageClass"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for messageClass")
			}
			messageClass = m.GetMessageClass()
			_messageClassErr := writeBuffer.WriteSerializable(messageClass)
			if popErr := writeBuffer.PopContext("messageClass"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for messageClass")
			}
			if _messageClassErr != nil {
				return errors.Wrap(_messageClassErr, "Error serializing 'messageClass' field")
			}
		}

		// Simple Field (messagePriority)
		if pushErr := writeBuffer.PushContext("messagePriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messagePriority")
		}
		_messagePriorityErr := writeBuffer.WriteSerializable(m.GetMessagePriority())
		if popErr := writeBuffer.PopContext("messagePriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messagePriority")
		}
		if _messagePriorityErr != nil {
			return errors.Wrap(_messagePriorityErr, "Error serializing 'messagePriority' field")
		}

		// Simple Field (message)
		if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for message")
		}
		_messageErr := writeBuffer.WriteSerializable(m.GetMessage())
		if popErr := writeBuffer.PopContext("message"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for message")
		}
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessage")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) isBACnetConfirmedServiceRequestConfirmedTextMessage() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}