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

// CALDataReset is the corresponding interface of CALDataReset
type CALDataReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
}

// CALDataResetExactly can be used when we want exactly this type and not a type which fulfills CALDataReset.
// This is useful for switch cases.
type CALDataResetExactly interface {
	CALDataReset
	isCALDataReset() bool
}

// _CALDataReset is the data-structure of this message
type _CALDataReset struct {
	CALDataContract
}

var _ CALDataReset = (*_CALDataReset)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataReset) GetParent() CALDataContract {
	return m.CALDataContract
}

// NewCALDataReset factory function for _CALDataReset
func NewCALDataReset(commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataReset {
	_result := &_CALDataReset{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataReset(structType any) CALDataReset {
	if casted, ok := structType.(CALDataReset); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataReset); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataReset) GetTypeName() string {
	return "CALDataReset"
}

func (m *_CALDataReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_CALDataReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataReset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, requestContext RequestContext) (__cALDataReset CALDataReset, err error) {
	m.CALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALDataReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReset")
	}

	return m, nil
}

func (m *_CALDataReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReset")
		}

		if popErr := writeBuffer.PopContext("CALDataReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReset")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataReset) isCALDataReset() bool {
	return true
}

func (m *_CALDataReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
