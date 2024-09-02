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

// CipService is the corresponding interface of CipService
type CipService interface {
	CipServiceContract
	CipServiceRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CipServiceContract provides a set of functions which can be overwritten by a sub struct
type CipServiceContract interface {
	// GetServiceLen() returns a parser argument
	GetServiceLen() uint16
}

// CipServiceRequirements provides a set of functions which need to be implemented by a sub struct
type CipServiceRequirements interface {
	// GetConnected returns Connected (discriminator field)
	GetConnected() bool
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetService returns Service (discriminator field)
	GetService() uint8
}

// CipServiceExactly can be used when we want exactly this type and not a type which fulfills CipService.
// This is useful for switch cases.
type CipServiceExactly interface {
	CipService
	isCipService() bool
}

// _CipService is the data-structure of this message
type _CipService struct {
	_CipServiceChildRequirements

	// Arguments.
	ServiceLen uint16
}

var _ CipServiceContract = (*_CipService)(nil)

type _CipServiceChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetConnected() bool
	GetResponse() bool
	GetService() uint8
}

type CipServiceChild interface {
	utils.Serializable

	GetParent() *CipService

	GetTypeName() string
	CipService
}

// NewCipService factory function for _CipService
func NewCipService(serviceLen uint16) *_CipService {
	return &_CipService{ServiceLen: serviceLen}
}

// Deprecated: use the interface for direct cast
func CastCipService(structType any) CipService {
	if casted, ok := structType.(CipService); ok {
		return casted
	}
	if casted, ok := structType.(*CipService); ok {
		return *casted
	}
	return nil
}

func (m *_CipService) GetTypeName() string {
	return "CipService"
}

func (m *_CipService) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (response)
	lengthInBits += 1
	// Discriminator Field (service)
	lengthInBits += 7

	return lengthInBits
}

func (m *_CipService) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipServiceParse[T CipService](ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (T, error) {
	return CipServiceParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipServiceParseWithBufferProducer[T CipService](connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CipServiceParseWithBuffer[T](ctx, readBuffer, connected, serviceLen)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CipServiceParseWithBuffer[T CipService](ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (T, error) {
	v, err := (&_CipService{}).parse(ctx, readBuffer, connected, serviceLen)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CipService) parse(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (__cipService CipService, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipService"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipService")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	response, err := ReadDiscriminatorField[bool](ctx, "response", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'response' field"))
	}

	service, err := ReadDiscriminatorField[uint8](ctx, "service", ReadUnsignedByte(readBuffer, uint8(7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'service' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CipService
	switch {
	case service == 0x01 && response == bool(false): // GetAttributeAllRequest
		if _child, err = (&_GetAttributeAllRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeAllRequest for type-switch of CipService")
		}
	case service == 0x01 && response == bool(true): // GetAttributeAllResponse
		if _child, err = (&_GetAttributeAllResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeAllResponse for type-switch of CipService")
		}
	case service == 0x02 && response == bool(false): // SetAttributeAllRequest
		if _child, err = (&_SetAttributeAllRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeAllRequest for type-switch of CipService")
		}
	case service == 0x02 && response == bool(true): // SetAttributeAllResponse
		if _child, err = (&_SetAttributeAllResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeAllResponse for type-switch of CipService")
		}
	case service == 0x03 && response == bool(false): // GetAttributeListRequest
		if _child, err = (&_GetAttributeListRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeListRequest for type-switch of CipService")
		}
	case service == 0x03 && response == bool(true): // GetAttributeListResponse
		if _child, err = (&_GetAttributeListResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeListResponse for type-switch of CipService")
		}
	case service == 0x04 && response == bool(false): // SetAttributeListRequest
		if _child, err = (&_SetAttributeListRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeListRequest for type-switch of CipService")
		}
	case service == 0x04 && response == bool(true): // SetAttributeListResponse
		if _child, err = (&_SetAttributeListResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeListResponse for type-switch of CipService")
		}
	case service == 0x0A && response == bool(false): // MultipleServiceRequest
		if _child, err = (&_MultipleServiceRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MultipleServiceRequest for type-switch of CipService")
		}
	case service == 0x0A && response == bool(true): // MultipleServiceResponse
		if _child, err = (&_MultipleServiceResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MultipleServiceResponse for type-switch of CipService")
		}
	case service == 0x0E && response == bool(false): // GetAttributeSingleRequest
		if _child, err = (&_GetAttributeSingleRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeSingleRequest for type-switch of CipService")
		}
	case service == 0x0E && response == bool(true): // GetAttributeSingleResponse
		if _child, err = (&_GetAttributeSingleResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type GetAttributeSingleResponse for type-switch of CipService")
		}
	case service == 0x10 && response == bool(false): // SetAttributeSingleRequest
		if _child, err = (&_SetAttributeSingleRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeSingleRequest for type-switch of CipService")
		}
	case service == 0x10 && response == bool(true): // SetAttributeSingleResponse
		if _child, err = (&_SetAttributeSingleResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SetAttributeSingleResponse for type-switch of CipService")
		}
	case service == 0x4C && response == bool(false): // CipReadRequest
		if _child, err = (&_CipReadRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipReadRequest for type-switch of CipService")
		}
	case service == 0x4C && response == bool(true): // CipReadResponse
		if _child, err = (&_CipReadResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipReadResponse for type-switch of CipService")
		}
	case service == 0x4D && response == bool(false): // CipWriteRequest
		if _child, err = (&_CipWriteRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipWriteRequest for type-switch of CipService")
		}
	case service == 0x4D && response == bool(true): // CipWriteResponse
		if _child, err = (&_CipWriteResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipWriteResponse for type-switch of CipService")
		}
	case service == 0x4E && response == bool(false): // CipConnectionManagerCloseRequest
		if _child, err = (&_CipConnectionManagerCloseRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerCloseRequest for type-switch of CipService")
		}
	case service == 0x4E && response == bool(true): // CipConnectionManagerCloseResponse
		if _child, err = (&_CipConnectionManagerCloseResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerCloseResponse for type-switch of CipService")
		}
	case service == 0x52 && response == bool(false) && connected == bool(false): // CipUnconnectedRequest
		if _child, err = (&_CipUnconnectedRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipUnconnectedRequest for type-switch of CipService")
		}
	case service == 0x52 && response == bool(false) && connected == bool(true): // CipConnectedRequest
		if _child, err = (&_CipConnectedRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectedRequest for type-switch of CipService")
		}
	case service == 0x52 && response == bool(true): // CipConnectedResponse
		if _child, err = (&_CipConnectedResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectedResponse for type-switch of CipService")
		}
	case service == 0x5B && response == bool(false): // CipConnectionManagerRequest
		if _child, err = (&_CipConnectionManagerRequest{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerRequest for type-switch of CipService")
		}
	case service == 0x5B && response == bool(true): // CipConnectionManagerResponse
		if _child, err = (&_CipConnectionManagerResponse{}).parse(ctx, readBuffer, m, connected, serviceLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipConnectionManagerResponse for type-switch of CipService")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [service=%v, response=%v, connected=%v]", service, response, connected)
	}

	if closeErr := readBuffer.CloseContext("CipService"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipService")
	}

	return _child, nil
}

func (pm *_CipService) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CipService, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CipService"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CipService")
	}

	if err := WriteDiscriminatorField(ctx, "response", m.GetResponse(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'response' field")
	}

	if err := WriteDiscriminatorField(ctx, "service", m.GetService(), WriteUnsignedByte(writeBuffer, 7)); err != nil {
		return errors.Wrap(err, "Error serializing 'service' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CipService"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CipService")
	}
	return nil
}

////
// Arguments Getter

func (m *_CipService) GetServiceLen() uint16 {
	return m.ServiceLen
}

//
////

func (m *_CipService) isCipService() bool {
	return true
}
