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

// SecurityDataOtherAlarmCleared is the corresponding interface of SecurityDataOtherAlarmCleared
type SecurityDataOtherAlarmCleared interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataOtherAlarmClearedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataOtherAlarmCleared.
// This is useful for switch cases.
type SecurityDataOtherAlarmClearedExactly interface {
	SecurityDataOtherAlarmCleared
	isSecurityDataOtherAlarmCleared() bool
}

// _SecurityDataOtherAlarmCleared is the data-structure of this message
type _SecurityDataOtherAlarmCleared struct {
	SecurityDataContract
}

var _ SecurityDataOtherAlarmCleared = (*_SecurityDataOtherAlarmCleared)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataOtherAlarmCleared) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataOtherAlarmCleared factory function for _SecurityDataOtherAlarmCleared
func NewSecurityDataOtherAlarmCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataOtherAlarmCleared {
	_result := &_SecurityDataOtherAlarmCleared{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataOtherAlarmCleared(structType any) SecurityDataOtherAlarmCleared {
	if casted, ok := structType.(SecurityDataOtherAlarmCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataOtherAlarmCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataOtherAlarmCleared) GetTypeName() string {
	return "SecurityDataOtherAlarmCleared"
}

func (m *_SecurityDataOtherAlarmCleared) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataOtherAlarmCleared) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataOtherAlarmCleared) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataOtherAlarmCleared SecurityDataOtherAlarmCleared, err error) {
	m.SecurityDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataOtherAlarmCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataOtherAlarmCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataOtherAlarmCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataOtherAlarmCleared")
	}

	return m, nil
}

func (m *_SecurityDataOtherAlarmCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataOtherAlarmCleared) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataOtherAlarmCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataOtherAlarmCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataOtherAlarmCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataOtherAlarmCleared")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataOtherAlarmCleared) isSecurityDataOtherAlarmCleared() bool {
	return true
}

func (m *_SecurityDataOtherAlarmCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
