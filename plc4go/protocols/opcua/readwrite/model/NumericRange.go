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

// NumericRange is the corresponding interface of NumericRange
type NumericRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// NumericRangeExactly can be used when we want exactly this type and not a type which fulfills NumericRange.
// This is useful for switch cases.
type NumericRangeExactly interface {
	NumericRange
	isNumericRange() bool
}

// _NumericRange is the data-structure of this message
type _NumericRange struct {
}

// NewNumericRange factory function for _NumericRange
func NewNumericRange() *_NumericRange {
	return &_NumericRange{}
}

// Deprecated: use the interface for direct cast
func CastNumericRange(structType any) NumericRange {
	if casted, ok := structType.(NumericRange); ok {
		return casted
	}
	if casted, ok := structType.(*NumericRange); ok {
		return *casted
	}
	return nil
}

func (m *_NumericRange) GetTypeName() string {
	return "NumericRange"
}

func (m *_NumericRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_NumericRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NumericRangeParse(ctx context.Context, theBytes []byte) (NumericRange, error) {
	return NumericRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NumericRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NumericRange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NumericRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NumericRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NumericRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NumericRange")
	}

	// Create the instance
	return &_NumericRange{}, nil
}

func (m *_NumericRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NumericRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NumericRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NumericRange")
	}

	if popErr := writeBuffer.PopContext("NumericRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NumericRange")
	}
	return nil
}

func (m *_NumericRange) isNumericRange() bool {
	return true
}

func (m *_NumericRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
