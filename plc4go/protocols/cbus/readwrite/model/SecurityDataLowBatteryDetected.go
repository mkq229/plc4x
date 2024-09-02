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

// SecurityDataLowBatteryDetected is the corresponding interface of SecurityDataLowBatteryDetected
type SecurityDataLowBatteryDetected interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataLowBatteryDetectedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataLowBatteryDetected.
// This is useful for switch cases.
type SecurityDataLowBatteryDetectedExactly interface {
	SecurityDataLowBatteryDetected
	isSecurityDataLowBatteryDetected() bool
}

// _SecurityDataLowBatteryDetected is the data-structure of this message
type _SecurityDataLowBatteryDetected struct {
	SecurityDataContract
}

var _ SecurityDataLowBatteryDetected = (*_SecurityDataLowBatteryDetected)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataLowBatteryDetected) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataLowBatteryDetected factory function for _SecurityDataLowBatteryDetected
func NewSecurityDataLowBatteryDetected(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLowBatteryDetected {
	_result := &_SecurityDataLowBatteryDetected{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLowBatteryDetected(structType any) SecurityDataLowBatteryDetected {
	if casted, ok := structType.(SecurityDataLowBatteryDetected); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLowBatteryDetected); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLowBatteryDetected) GetTypeName() string {
	return "SecurityDataLowBatteryDetected"
}

func (m *_SecurityDataLowBatteryDetected) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataLowBatteryDetected) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataLowBatteryDetected) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataLowBatteryDetected SecurityDataLowBatteryDetected, err error) {
	m.SecurityDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLowBatteryDetected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLowBatteryDetected")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataLowBatteryDetected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLowBatteryDetected")
	}

	return m, nil
}

func (m *_SecurityDataLowBatteryDetected) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLowBatteryDetected) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLowBatteryDetected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLowBatteryDetected")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLowBatteryDetected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLowBatteryDetected")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLowBatteryDetected) isSecurityDataLowBatteryDetected() bool {
	return true
}

func (m *_SecurityDataLowBatteryDetected) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
