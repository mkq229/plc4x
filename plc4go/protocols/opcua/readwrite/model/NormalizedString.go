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

// NormalizedString is the corresponding interface of NormalizedString
type NormalizedString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// NormalizedStringExactly can be used when we want exactly this type and not a type which fulfills NormalizedString.
// This is useful for switch cases.
type NormalizedStringExactly interface {
	NormalizedString
	isNormalizedString() bool
}

// _NormalizedString is the data-structure of this message
type _NormalizedString struct {
}

var _ NormalizedString = (*_NormalizedString)(nil)

// NewNormalizedString factory function for _NormalizedString
func NewNormalizedString() *_NormalizedString {
	return &_NormalizedString{}
}

// Deprecated: use the interface for direct cast
func CastNormalizedString(structType any) NormalizedString {
	if casted, ok := structType.(NormalizedString); ok {
		return casted
	}
	if casted, ok := structType.(*NormalizedString); ok {
		return *casted
	}
	return nil
}

func (m *_NormalizedString) GetTypeName() string {
	return "NormalizedString"
}

func (m *_NormalizedString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_NormalizedString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NormalizedStringParse(ctx context.Context, theBytes []byte) (NormalizedString, error) {
	return NormalizedStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NormalizedStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (NormalizedString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NormalizedString, error) {
		return NormalizedStringParseWithBuffer(ctx, readBuffer)
	}
}

func NormalizedStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NormalizedString, error) {
	v, err := (&_NormalizedString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_NormalizedString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__normalizedString NormalizedString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NormalizedString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NormalizedString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NormalizedString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NormalizedString")
	}

	return m, nil
}

func (m *_NormalizedString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NormalizedString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NormalizedString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NormalizedString")
	}

	if popErr := writeBuffer.PopContext("NormalizedString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NormalizedString")
	}
	return nil
}

func (m *_NormalizedString) isNormalizedString() bool {
	return true
}

func (m *_NormalizedString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
