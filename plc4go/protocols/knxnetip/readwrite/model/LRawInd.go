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

// LRawInd is the corresponding interface of LRawInd
type LRawInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// LRawIndExactly can be used when we want exactly this type and not a type which fulfills LRawInd.
// This is useful for switch cases.
type LRawIndExactly interface {
	LRawInd
	isLRawInd() bool
}

// _LRawInd is the data-structure of this message
type _LRawInd struct {
	CEMIContract
}

var _ LRawInd = (*_LRawInd)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LRawInd) GetMessageCode() uint8 {
	return 0x2D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LRawInd) GetParent() CEMIContract {
	return m.CEMIContract
}

// NewLRawInd factory function for _LRawInd
func NewLRawInd(size uint16) *_LRawInd {
	_result := &_LRawInd{
		CEMIContract: NewCEMI(size),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastLRawInd(structType any) LRawInd {
	if casted, ok := structType.(LRawInd); ok {
		return casted
	}
	if casted, ok := structType.(*LRawInd); ok {
		return *casted
	}
	return nil
}

func (m *_LRawInd) GetTypeName() string {
	return "LRawInd"
}

func (m *_LRawInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LRawInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LRawInd) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lRawInd LRawInd, err error) {
	m.CEMIContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LRawInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LRawInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LRawInd")
	}

	return m, nil
}

func (m *_LRawInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LRawInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LRawInd")
		}

		if popErr := writeBuffer.PopContext("LRawInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LRawInd")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LRawInd) isLRawInd() bool {
	return true
}

func (m *_LRawInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
