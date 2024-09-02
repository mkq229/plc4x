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

// BACnetRecipient is the corresponding interface of BACnetRecipient
type BACnetRecipient interface {
	BACnetRecipientContract
	BACnetRecipientRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetRecipientContract provides a set of functions which can be overwritten by a sub struct
type BACnetRecipientContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetRecipientRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetRecipientRequirements interface {
}

// BACnetRecipientExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipient.
// This is useful for switch cases.
type BACnetRecipientExactly interface {
	BACnetRecipient
	isBACnetRecipient() bool
}

// _BACnetRecipient is the data-structure of this message
type _BACnetRecipient struct {
	_BACnetRecipientChildRequirements
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetRecipientContract = (*_BACnetRecipient)(nil)

type _BACnetRecipientChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetRecipientChild interface {
	utils.Serializable

	GetParent() *BACnetRecipient

	GetTypeName() string
	BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipient) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetRecipient) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipient factory function for _BACnetRecipient
func NewBACnetRecipient(peekedTagHeader BACnetTagHeader) *_BACnetRecipient {
	return &_BACnetRecipient{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipient(structType any) BACnetRecipient {
	if casted, ok := structType.(BACnetRecipient); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipient); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipient) GetTypeName() string {
	return "BACnetRecipient"
}

func (m *_BACnetRecipient) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetRecipient) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientParse[T BACnetRecipient](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetRecipientParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRecipientParseWithBufferProducer[T BACnetRecipient]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetRecipientParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetRecipientParseWithBuffer[T BACnetRecipient](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetRecipient{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetRecipient) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetRecipient BACnetRecipient, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipient")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetRecipient
	switch {
	case peekedTagNumber == uint8(0): // BACnetRecipientDevice
		if _child, err = (&_BACnetRecipientDevice{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetRecipientDevice for type-switch of BACnetRecipient")
		}
	case peekedTagNumber == uint8(1): // BACnetRecipientAddress
		if _child, err = (&_BACnetRecipientAddress{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetRecipientAddress for type-switch of BACnetRecipient")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipient")
	}

	return _child, nil
}

func (pm *_BACnetRecipient) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetRecipient, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRecipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipient")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipient"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipient")
	}
	return nil
}

func (m *_BACnetRecipient) isBACnetRecipient() bool {
	return true
}
