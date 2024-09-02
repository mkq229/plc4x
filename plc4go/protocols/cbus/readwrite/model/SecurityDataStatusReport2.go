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

// SecurityDataStatusReport2 is the corresponding interface of SecurityDataStatusReport2
type SecurityDataStatusReport2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneStatus returns ZoneStatus (property field)
	GetZoneStatus() []ZoneStatus
}

// SecurityDataStatusReport2Exactly can be used when we want exactly this type and not a type which fulfills SecurityDataStatusReport2.
// This is useful for switch cases.
type SecurityDataStatusReport2Exactly interface {
	SecurityDataStatusReport2
	isSecurityDataStatusReport2() bool
}

// _SecurityDataStatusReport2 is the data-structure of this message
type _SecurityDataStatusReport2 struct {
	SecurityDataContract
	ZoneStatus []ZoneStatus
}

var _ SecurityDataStatusReport2 = (*_SecurityDataStatusReport2)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataStatusReport2) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataStatusReport2) GetZoneStatus() []ZoneStatus {
	return m.ZoneStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataStatusReport2 factory function for _SecurityDataStatusReport2
func NewSecurityDataStatusReport2(zoneStatus []ZoneStatus, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatusReport2 {
	_result := &_SecurityDataStatusReport2{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneStatus:           zoneStatus,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatusReport2(structType any) SecurityDataStatusReport2 {
	if casted, ok := structType.(SecurityDataStatusReport2); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatusReport2); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatusReport2) GetTypeName() string {
	return "SecurityDataStatusReport2"
}

func (m *_SecurityDataStatusReport2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Array field
	if len(m.ZoneStatus) > 0 {
		for _curItem, element := range m.ZoneStatus {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ZoneStatus), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SecurityDataStatusReport2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataStatusReport2) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataStatusReport2 SecurityDataStatusReport2, err error) {
	m.SecurityDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataStatusReport2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatusReport2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneStatus, err := ReadCountArrayField[ZoneStatus](ctx, "zoneStatus", ReadComplex[ZoneStatus](ZoneStatusParseWithBuffer, readBuffer), uint64(int32(48)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneStatus' field"))
	}
	m.ZoneStatus = zoneStatus

	if closeErr := readBuffer.CloseContext("SecurityDataStatusReport2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatusReport2")
	}

	return m, nil
}

func (m *_SecurityDataStatusReport2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatusReport2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatusReport2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatusReport2")
		}

		if err := WriteComplexTypeArrayField(ctx, "zoneStatus", m.GetZoneStatus(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneStatus' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatusReport2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatusReport2")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatusReport2) isSecurityDataStatusReport2() bool {
	return true
}

func (m *_SecurityDataStatusReport2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
