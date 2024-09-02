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

// BACnetWeekNDay is the corresponding interface of BACnetWeekNDay
type BACnetWeekNDay interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetWeekNDayExactly can be used when we want exactly this type and not a type which fulfills BACnetWeekNDay.
// This is useful for switch cases.
type BACnetWeekNDayExactly interface {
	BACnetWeekNDay
	isBACnetWeekNDay() bool
}

// _BACnetWeekNDay is the data-structure of this message
type _BACnetWeekNDay struct {
}

var _ BACnetWeekNDay = (*_BACnetWeekNDay)(nil)

// NewBACnetWeekNDay factory function for _BACnetWeekNDay
func NewBACnetWeekNDay() *_BACnetWeekNDay {
	return &_BACnetWeekNDay{}
}

// Deprecated: use the interface for direct cast
func CastBACnetWeekNDay(structType any) BACnetWeekNDay {
	if casted, ok := structType.(BACnetWeekNDay); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetWeekNDay); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetWeekNDay) GetTypeName() string {
	return "BACnetWeekNDay"
}

func (m *_BACnetWeekNDay) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_BACnetWeekNDay) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetWeekNDayParse(ctx context.Context, theBytes []byte) (BACnetWeekNDay, error) {
	return BACnetWeekNDayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetWeekNDayParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWeekNDay, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWeekNDay, error) {
		return BACnetWeekNDayParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetWeekNDayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWeekNDay, error) {
	v, err := (&_BACnetWeekNDay{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetWeekNDay) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetWeekNDay BACnetWeekNDay, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetWeekNDay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetWeekNDay")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "Unusable type. Exits only for consistency. Use BACnetWeekNDayTagged"})
	}

	if closeErr := readBuffer.CloseContext("BACnetWeekNDay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetWeekNDay")
	}

	return m, nil
}

func (m *_BACnetWeekNDay) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetWeekNDay) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetWeekNDay"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetWeekNDay")
	}

	if popErr := writeBuffer.PopContext("BACnetWeekNDay"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetWeekNDay")
	}
	return nil
}

func (m *_BACnetWeekNDay) isBACnetWeekNDay() bool {
	return true
}

func (m *_BACnetWeekNDay) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
