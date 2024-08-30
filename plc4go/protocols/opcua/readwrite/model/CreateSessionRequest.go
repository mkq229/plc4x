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

// CreateSessionRequest is the corresponding interface of CreateSessionRequest
type CreateSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetClientDescription returns ClientDescription (property field)
	GetClientDescription() ExtensionObjectDefinition
	// GetServerUri returns ServerUri (property field)
	GetServerUri() PascalString
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetSessionName returns SessionName (property field)
	GetSessionName() PascalString
	// GetClientNonce returns ClientNonce (property field)
	GetClientNonce() PascalByteString
	// GetClientCertificate returns ClientCertificate (property field)
	GetClientCertificate() PascalByteString
	// GetRequestedSessionTimeout returns RequestedSessionTimeout (property field)
	GetRequestedSessionTimeout() float64
	// GetMaxResponseMessageSize returns MaxResponseMessageSize (property field)
	GetMaxResponseMessageSize() uint32
}

// CreateSessionRequestExactly can be used when we want exactly this type and not a type which fulfills CreateSessionRequest.
// This is useful for switch cases.
type CreateSessionRequestExactly interface {
	CreateSessionRequest
	isCreateSessionRequest() bool
}

// _CreateSessionRequest is the data-structure of this message
type _CreateSessionRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader           ExtensionObjectDefinition
	ClientDescription       ExtensionObjectDefinition
	ServerUri               PascalString
	EndpointUrl             PascalString
	SessionName             PascalString
	ClientNonce             PascalByteString
	ClientCertificate       PascalByteString
	RequestedSessionTimeout float64
	MaxResponseMessageSize  uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSessionRequest) GetIdentifier() string {
	return "461"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSessionRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CreateSessionRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSessionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateSessionRequest) GetClientDescription() ExtensionObjectDefinition {
	return m.ClientDescription
}

func (m *_CreateSessionRequest) GetServerUri() PascalString {
	return m.ServerUri
}

func (m *_CreateSessionRequest) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_CreateSessionRequest) GetSessionName() PascalString {
	return m.SessionName
}

func (m *_CreateSessionRequest) GetClientNonce() PascalByteString {
	return m.ClientNonce
}

func (m *_CreateSessionRequest) GetClientCertificate() PascalByteString {
	return m.ClientCertificate
}

func (m *_CreateSessionRequest) GetRequestedSessionTimeout() float64 {
	return m.RequestedSessionTimeout
}

func (m *_CreateSessionRequest) GetMaxResponseMessageSize() uint32 {
	return m.MaxResponseMessageSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateSessionRequest factory function for _CreateSessionRequest
func NewCreateSessionRequest(requestHeader ExtensionObjectDefinition, clientDescription ExtensionObjectDefinition, serverUri PascalString, endpointUrl PascalString, sessionName PascalString, clientNonce PascalByteString, clientCertificate PascalByteString, requestedSessionTimeout float64, maxResponseMessageSize uint32) *_CreateSessionRequest {
	_result := &_CreateSessionRequest{
		RequestHeader:              requestHeader,
		ClientDescription:          clientDescription,
		ServerUri:                  serverUri,
		EndpointUrl:                endpointUrl,
		SessionName:                sessionName,
		ClientNonce:                clientNonce,
		ClientCertificate:          clientCertificate,
		RequestedSessionTimeout:    requestedSessionTimeout,
		MaxResponseMessageSize:     maxResponseMessageSize,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateSessionRequest(structType any) CreateSessionRequest {
	if casted, ok := structType.(CreateSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSessionRequest) GetTypeName() string {
	return "CreateSessionRequest"
}

func (m *_CreateSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientDescription)
	lengthInBits += m.ClientDescription.GetLengthInBits(ctx)

	// Simple field (serverUri)
	lengthInBits += m.ServerUri.GetLengthInBits(ctx)

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (sessionName)
	lengthInBits += m.SessionName.GetLengthInBits(ctx)

	// Simple field (clientNonce)
	lengthInBits += m.ClientNonce.GetLengthInBits(ctx)

	// Simple field (clientCertificate)
	lengthInBits += m.ClientCertificate.GetLengthInBits(ctx)

	// Simple field (requestedSessionTimeout)
	lengthInBits += 64

	// Simple field (maxResponseMessageSize)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CreateSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CreateSessionRequestParse(ctx context.Context, theBytes []byte, identifier string) (CreateSessionRequest, error) {
	return CreateSessionRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CreateSessionRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CreateSessionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CreateSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of CreateSessionRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (clientDescription)
	if pullErr := readBuffer.PullContext("clientDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientDescription")
	}
	_clientDescription, _clientDescriptionErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("310"))
	if _clientDescriptionErr != nil {
		return nil, errors.Wrap(_clientDescriptionErr, "Error parsing 'clientDescription' field of CreateSessionRequest")
	}
	clientDescription := _clientDescription.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("clientDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientDescription")
	}

	// Simple Field (serverUri)
	if pullErr := readBuffer.PullContext("serverUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverUri")
	}
	_serverUri, _serverUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _serverUriErr != nil {
		return nil, errors.Wrap(_serverUriErr, "Error parsing 'serverUri' field of CreateSessionRequest")
	}
	serverUri := _serverUri.(PascalString)
	if closeErr := readBuffer.CloseContext("serverUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverUri")
	}

	// Simple Field (endpointUrl)
	if pullErr := readBuffer.PullContext("endpointUrl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endpointUrl")
	}
	_endpointUrl, _endpointUrlErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _endpointUrlErr != nil {
		return nil, errors.Wrap(_endpointUrlErr, "Error parsing 'endpointUrl' field of CreateSessionRequest")
	}
	endpointUrl := _endpointUrl.(PascalString)
	if closeErr := readBuffer.CloseContext("endpointUrl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endpointUrl")
	}

	// Simple Field (sessionName)
	if pullErr := readBuffer.PullContext("sessionName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sessionName")
	}
	_sessionName, _sessionNameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _sessionNameErr != nil {
		return nil, errors.Wrap(_sessionNameErr, "Error parsing 'sessionName' field of CreateSessionRequest")
	}
	sessionName := _sessionName.(PascalString)
	if closeErr := readBuffer.CloseContext("sessionName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sessionName")
	}

	// Simple Field (clientNonce)
	if pullErr := readBuffer.PullContext("clientNonce"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientNonce")
	}
	_clientNonce, _clientNonceErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _clientNonceErr != nil {
		return nil, errors.Wrap(_clientNonceErr, "Error parsing 'clientNonce' field of CreateSessionRequest")
	}
	clientNonce := _clientNonce.(PascalByteString)
	if closeErr := readBuffer.CloseContext("clientNonce"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientNonce")
	}

	// Simple Field (clientCertificate)
	if pullErr := readBuffer.PullContext("clientCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientCertificate")
	}
	_clientCertificate, _clientCertificateErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _clientCertificateErr != nil {
		return nil, errors.Wrap(_clientCertificateErr, "Error parsing 'clientCertificate' field of CreateSessionRequest")
	}
	clientCertificate := _clientCertificate.(PascalByteString)
	if closeErr := readBuffer.CloseContext("clientCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientCertificate")
	}

	// Simple Field (requestedSessionTimeout)
	_requestedSessionTimeout, _requestedSessionTimeoutErr := /*TODO: migrate me*/ readBuffer.ReadFloat64("requestedSessionTimeout", 64)
	if _requestedSessionTimeoutErr != nil {
		return nil, errors.Wrap(_requestedSessionTimeoutErr, "Error parsing 'requestedSessionTimeout' field of CreateSessionRequest")
	}
	requestedSessionTimeout := _requestedSessionTimeout

	// Simple Field (maxResponseMessageSize)
	_maxResponseMessageSize, _maxResponseMessageSizeErr := /*TODO: migrate me*/ readBuffer.ReadUint32("maxResponseMessageSize", 32)
	if _maxResponseMessageSizeErr != nil {
		return nil, errors.Wrap(_maxResponseMessageSizeErr, "Error parsing 'maxResponseMessageSize' field of CreateSessionRequest")
	}
	maxResponseMessageSize := _maxResponseMessageSize

	if closeErr := readBuffer.CloseContext("CreateSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSessionRequest")
	}

	// Create a partially initialized instance
	_child := &_CreateSessionRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		ClientDescription:          clientDescription,
		ServerUri:                  serverUri,
		EndpointUrl:                endpointUrl,
		SessionName:                sessionName,
		ClientNonce:                clientNonce,
		ClientCertificate:          clientCertificate,
		RequestedSessionTimeout:    requestedSessionTimeout,
		MaxResponseMessageSize:     maxResponseMessageSize,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CreateSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSessionRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (clientDescription)
		if pushErr := writeBuffer.PushContext("clientDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientDescription")
		}
		_clientDescriptionErr := writeBuffer.WriteSerializable(ctx, m.GetClientDescription())
		if popErr := writeBuffer.PopContext("clientDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientDescription")
		}
		if _clientDescriptionErr != nil {
			return errors.Wrap(_clientDescriptionErr, "Error serializing 'clientDescription' field")
		}

		// Simple Field (serverUri)
		if pushErr := writeBuffer.PushContext("serverUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverUri")
		}
		_serverUriErr := writeBuffer.WriteSerializable(ctx, m.GetServerUri())
		if popErr := writeBuffer.PopContext("serverUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverUri")
		}
		if _serverUriErr != nil {
			return errors.Wrap(_serverUriErr, "Error serializing 'serverUri' field")
		}

		// Simple Field (endpointUrl)
		if pushErr := writeBuffer.PushContext("endpointUrl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endpointUrl")
		}
		_endpointUrlErr := writeBuffer.WriteSerializable(ctx, m.GetEndpointUrl())
		if popErr := writeBuffer.PopContext("endpointUrl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endpointUrl")
		}
		if _endpointUrlErr != nil {
			return errors.Wrap(_endpointUrlErr, "Error serializing 'endpointUrl' field")
		}

		// Simple Field (sessionName)
		if pushErr := writeBuffer.PushContext("sessionName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sessionName")
		}
		_sessionNameErr := writeBuffer.WriteSerializable(ctx, m.GetSessionName())
		if popErr := writeBuffer.PopContext("sessionName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sessionName")
		}
		if _sessionNameErr != nil {
			return errors.Wrap(_sessionNameErr, "Error serializing 'sessionName' field")
		}

		// Simple Field (clientNonce)
		if pushErr := writeBuffer.PushContext("clientNonce"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientNonce")
		}
		_clientNonceErr := writeBuffer.WriteSerializable(ctx, m.GetClientNonce())
		if popErr := writeBuffer.PopContext("clientNonce"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientNonce")
		}
		if _clientNonceErr != nil {
			return errors.Wrap(_clientNonceErr, "Error serializing 'clientNonce' field")
		}

		// Simple Field (clientCertificate)
		if pushErr := writeBuffer.PushContext("clientCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientCertificate")
		}
		_clientCertificateErr := writeBuffer.WriteSerializable(ctx, m.GetClientCertificate())
		if popErr := writeBuffer.PopContext("clientCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientCertificate")
		}
		if _clientCertificateErr != nil {
			return errors.Wrap(_clientCertificateErr, "Error serializing 'clientCertificate' field")
		}

		// Simple Field (requestedSessionTimeout)
		requestedSessionTimeout := float64(m.GetRequestedSessionTimeout())
		_requestedSessionTimeoutErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("requestedSessionTimeout", 64, (requestedSessionTimeout))
		if _requestedSessionTimeoutErr != nil {
			return errors.Wrap(_requestedSessionTimeoutErr, "Error serializing 'requestedSessionTimeout' field")
		}

		// Simple Field (maxResponseMessageSize)
		maxResponseMessageSize := uint32(m.GetMaxResponseMessageSize())
		_maxResponseMessageSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("maxResponseMessageSize", 32, uint32((maxResponseMessageSize)))
		if _maxResponseMessageSizeErr != nil {
			return errors.Wrap(_maxResponseMessageSizeErr, "Error serializing 'maxResponseMessageSize' field")
		}

		if popErr := writeBuffer.PopContext("CreateSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSessionRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSessionRequest) isCreateSessionRequest() bool {
	return true
}

func (m *_CreateSessionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
