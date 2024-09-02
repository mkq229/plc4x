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

// BACnetAuthenticationPolicy is the corresponding interface of BACnetAuthenticationPolicy
type BACnetAuthenticationPolicy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPolicy returns Policy (property field)
	GetPolicy() BACnetAuthenticationPolicyList
	// GetOrderEnforced returns OrderEnforced (property field)
	GetOrderEnforced() BACnetContextTagBoolean
	// GetTimeout returns Timeout (property field)
	GetTimeout() BACnetContextTagUnsignedInteger
}

// BACnetAuthenticationPolicyExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationPolicy.
// This is useful for switch cases.
type BACnetAuthenticationPolicyExactly interface {
	BACnetAuthenticationPolicy
	isBACnetAuthenticationPolicy() bool
}

// _BACnetAuthenticationPolicy is the data-structure of this message
type _BACnetAuthenticationPolicy struct {
	Policy        BACnetAuthenticationPolicyList
	OrderEnforced BACnetContextTagBoolean
	Timeout       BACnetContextTagUnsignedInteger
}

var _ BACnetAuthenticationPolicy = (*_BACnetAuthenticationPolicy)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationPolicy) GetPolicy() BACnetAuthenticationPolicyList {
	return m.Policy
}

func (m *_BACnetAuthenticationPolicy) GetOrderEnforced() BACnetContextTagBoolean {
	return m.OrderEnforced
}

func (m *_BACnetAuthenticationPolicy) GetTimeout() BACnetContextTagUnsignedInteger {
	return m.Timeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationPolicy factory function for _BACnetAuthenticationPolicy
func NewBACnetAuthenticationPolicy(policy BACnetAuthenticationPolicyList, orderEnforced BACnetContextTagBoolean, timeout BACnetContextTagUnsignedInteger) *_BACnetAuthenticationPolicy {
	return &_BACnetAuthenticationPolicy{Policy: policy, OrderEnforced: orderEnforced, Timeout: timeout}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationPolicy(structType any) BACnetAuthenticationPolicy {
	if casted, ok := structType.(BACnetAuthenticationPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationPolicy) GetTypeName() string {
	return "BACnetAuthenticationPolicy"
}

func (m *_BACnetAuthenticationPolicy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (policy)
	lengthInBits += m.Policy.GetLengthInBits(ctx)

	// Simple field (orderEnforced)
	lengthInBits += m.OrderEnforced.GetLengthInBits(ctx)

	// Simple field (timeout)
	lengthInBits += m.Timeout.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationPolicy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationPolicyParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationPolicy, error) {
	return BACnetAuthenticationPolicyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationPolicyParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicy, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicy, error) {
		return BACnetAuthenticationPolicyParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAuthenticationPolicyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicy, error) {
	v, err := (&_BACnetAuthenticationPolicy{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetAuthenticationPolicy) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAuthenticationPolicy BACnetAuthenticationPolicy, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policy, err := ReadSimpleField[BACnetAuthenticationPolicyList](ctx, "policy", ReadComplex[BACnetAuthenticationPolicyList](BACnetAuthenticationPolicyListParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policy' field"))
	}
	m.Policy = policy

	orderEnforced, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "orderEnforced", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'orderEnforced' field"))
	}
	m.OrderEnforced = orderEnforced

	timeout, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeout", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeout' field"))
	}
	m.Timeout = timeout

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicy")
	}

	return m, nil
}

func (m *_BACnetAuthenticationPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationPolicy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicy")
	}

	if err := WriteSimpleField[BACnetAuthenticationPolicyList](ctx, "policy", m.GetPolicy(), WriteComplex[BACnetAuthenticationPolicyList](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'policy' field")
	}

	if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "orderEnforced", m.GetOrderEnforced(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'orderEnforced' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeout", m.GetTimeout(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeout' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationPolicy")
	}
	return nil
}

func (m *_BACnetAuthenticationPolicy) isBACnetAuthenticationPolicy() bool {
	return true
}

func (m *_BACnetAuthenticationPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
