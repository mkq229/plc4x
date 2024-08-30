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

// OpenChannelMessageResponse is the corresponding interface of OpenChannelMessageResponse
type OpenChannelMessageResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	OpenChannelMessage
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() int32
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetSenderCertificate returns SenderCertificate (property field)
	GetSenderCertificate() PascalByteString
	// GetReceiverCertificateThumbprint returns ReceiverCertificateThumbprint (property field)
	GetReceiverCertificateThumbprint() PascalByteString
}

// OpenChannelMessageResponseExactly can be used when we want exactly this type and not a type which fulfills OpenChannelMessageResponse.
// This is useful for switch cases.
type OpenChannelMessageResponseExactly interface {
	OpenChannelMessageResponse
	isOpenChannelMessageResponse() bool
}

// _OpenChannelMessageResponse is the data-structure of this message
type _OpenChannelMessageResponse struct {
	*_OpenChannelMessage
	SecureChannelId               int32
	SecurityPolicyUri             PascalString
	SenderCertificate             PascalByteString
	ReceiverCertificateThumbprint PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpenChannelMessageResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpenChannelMessageResponse) InitializeParent(parent OpenChannelMessage) {}

func (m *_OpenChannelMessageResponse) GetParent() OpenChannelMessage {
	return m._OpenChannelMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpenChannelMessageResponse) GetSecureChannelId() int32 {
	return m.SecureChannelId
}

func (m *_OpenChannelMessageResponse) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_OpenChannelMessageResponse) GetSenderCertificate() PascalByteString {
	return m.SenderCertificate
}

func (m *_OpenChannelMessageResponse) GetReceiverCertificateThumbprint() PascalByteString {
	return m.ReceiverCertificateThumbprint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpenChannelMessageResponse factory function for _OpenChannelMessageResponse
func NewOpenChannelMessageResponse(secureChannelId int32, securityPolicyUri PascalString, senderCertificate PascalByteString, receiverCertificateThumbprint PascalByteString) *_OpenChannelMessageResponse {
	_result := &_OpenChannelMessageResponse{
		SecureChannelId:               secureChannelId,
		SecurityPolicyUri:             securityPolicyUri,
		SenderCertificate:             senderCertificate,
		ReceiverCertificateThumbprint: receiverCertificateThumbprint,
		_OpenChannelMessage:           NewOpenChannelMessage(),
	}
	_result._OpenChannelMessage._OpenChannelMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpenChannelMessageResponse(structType any) OpenChannelMessageResponse {
	if casted, ok := structType.(OpenChannelMessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpenChannelMessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpenChannelMessageResponse) GetTypeName() string {
	return "OpenChannelMessageResponse"
}

func (m *_OpenChannelMessageResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (senderCertificate)
	lengthInBits += m.SenderCertificate.GetLengthInBits(ctx)

	// Simple field (receiverCertificateThumbprint)
	lengthInBits += m.ReceiverCertificateThumbprint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpenChannelMessageResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpenChannelMessageResponseParse(ctx context.Context, theBytes []byte, response bool) (OpenChannelMessageResponse, error) {
	return OpenChannelMessageResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func OpenChannelMessageResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (OpenChannelMessageResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpenChannelMessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpenChannelMessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (secureChannelId)
	_secureChannelId, _secureChannelIdErr := /*TODO: migrate me*/ readBuffer.ReadInt32("secureChannelId", 32)
	if _secureChannelIdErr != nil {
		return nil, errors.Wrap(_secureChannelIdErr, "Error parsing 'secureChannelId' field of OpenChannelMessageResponse")
	}
	secureChannelId := _secureChannelId

	// Simple Field (securityPolicyUri)
	if pullErr := readBuffer.PullContext("securityPolicyUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityPolicyUri")
	}
	_securityPolicyUri, _securityPolicyUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _securityPolicyUriErr != nil {
		return nil, errors.Wrap(_securityPolicyUriErr, "Error parsing 'securityPolicyUri' field of OpenChannelMessageResponse")
	}
	securityPolicyUri := _securityPolicyUri.(PascalString)
	if closeErr := readBuffer.CloseContext("securityPolicyUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityPolicyUri")
	}

	// Simple Field (senderCertificate)
	if pullErr := readBuffer.PullContext("senderCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for senderCertificate")
	}
	_senderCertificate, _senderCertificateErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _senderCertificateErr != nil {
		return nil, errors.Wrap(_senderCertificateErr, "Error parsing 'senderCertificate' field of OpenChannelMessageResponse")
	}
	senderCertificate := _senderCertificate.(PascalByteString)
	if closeErr := readBuffer.CloseContext("senderCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for senderCertificate")
	}

	// Simple Field (receiverCertificateThumbprint)
	if pullErr := readBuffer.PullContext("receiverCertificateThumbprint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for receiverCertificateThumbprint")
	}
	_receiverCertificateThumbprint, _receiverCertificateThumbprintErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _receiverCertificateThumbprintErr != nil {
		return nil, errors.Wrap(_receiverCertificateThumbprintErr, "Error parsing 'receiverCertificateThumbprint' field of OpenChannelMessageResponse")
	}
	receiverCertificateThumbprint := _receiverCertificateThumbprint.(PascalByteString)
	if closeErr := readBuffer.CloseContext("receiverCertificateThumbprint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for receiverCertificateThumbprint")
	}

	if closeErr := readBuffer.CloseContext("OpenChannelMessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpenChannelMessageResponse")
	}

	// Create a partially initialized instance
	_child := &_OpenChannelMessageResponse{
		_OpenChannelMessage:           &_OpenChannelMessage{},
		SecureChannelId:               secureChannelId,
		SecurityPolicyUri:             securityPolicyUri,
		SenderCertificate:             senderCertificate,
		ReceiverCertificateThumbprint: receiverCertificateThumbprint,
	}
	_child._OpenChannelMessage._OpenChannelMessageChildRequirements = _child
	return _child, nil
}

func (m *_OpenChannelMessageResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpenChannelMessageResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpenChannelMessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpenChannelMessageResponse")
		}

		// Simple Field (secureChannelId)
		secureChannelId := int32(m.GetSecureChannelId())
		_secureChannelIdErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("secureChannelId", 32, int32((secureChannelId)))
		if _secureChannelIdErr != nil {
			return errors.Wrap(_secureChannelIdErr, "Error serializing 'secureChannelId' field")
		}

		// Simple Field (securityPolicyUri)
		if pushErr := writeBuffer.PushContext("securityPolicyUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityPolicyUri")
		}
		_securityPolicyUriErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityPolicyUri())
		if popErr := writeBuffer.PopContext("securityPolicyUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityPolicyUri")
		}
		if _securityPolicyUriErr != nil {
			return errors.Wrap(_securityPolicyUriErr, "Error serializing 'securityPolicyUri' field")
		}

		// Simple Field (senderCertificate)
		if pushErr := writeBuffer.PushContext("senderCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for senderCertificate")
		}
		_senderCertificateErr := writeBuffer.WriteSerializable(ctx, m.GetSenderCertificate())
		if popErr := writeBuffer.PopContext("senderCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for senderCertificate")
		}
		if _senderCertificateErr != nil {
			return errors.Wrap(_senderCertificateErr, "Error serializing 'senderCertificate' field")
		}

		// Simple Field (receiverCertificateThumbprint)
		if pushErr := writeBuffer.PushContext("receiverCertificateThumbprint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for receiverCertificateThumbprint")
		}
		_receiverCertificateThumbprintErr := writeBuffer.WriteSerializable(ctx, m.GetReceiverCertificateThumbprint())
		if popErr := writeBuffer.PopContext("receiverCertificateThumbprint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for receiverCertificateThumbprint")
		}
		if _receiverCertificateThumbprintErr != nil {
			return errors.Wrap(_receiverCertificateThumbprintErr, "Error serializing 'receiverCertificateThumbprint' field")
		}

		if popErr := writeBuffer.PopContext("OpenChannelMessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpenChannelMessageResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpenChannelMessageResponse) isOpenChannelMessageResponse() bool {
	return true
}

func (m *_OpenChannelMessageResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
