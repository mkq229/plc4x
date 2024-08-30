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

// Constant values.
const PowerUp_POWERUPINDICATOR1 byte = 0x2B
const PowerUp_POWERUPINDICATOR2 byte = 0x2B

// PowerUp is the corresponding interface of PowerUp
type PowerUp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// PowerUpExactly can be used when we want exactly this type and not a type which fulfills PowerUp.
// This is useful for switch cases.
type PowerUpExactly interface {
	PowerUp
	isPowerUp() bool
}

// _PowerUp is the data-structure of this message
type _PowerUp struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_PowerUp) GetPowerUpIndicator1() byte {
	return PowerUp_POWERUPINDICATOR1
}

func (m *_PowerUp) GetPowerUpIndicator2() byte {
	return PowerUp_POWERUPINDICATOR2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPowerUp factory function for _PowerUp
func NewPowerUp() *_PowerUp {
	return &_PowerUp{}
}

// Deprecated: use the interface for direct cast
func CastPowerUp(structType any) PowerUp {
	if casted, ok := structType.(PowerUp); ok {
		return casted
	}
	if casted, ok := structType.(*PowerUp); ok {
		return *casted
	}
	return nil
}

func (m *_PowerUp) GetTypeName() string {
	return "PowerUp"
}

func (m *_PowerUp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (powerUpIndicator1)
	lengthInBits += 8

	// Const Field (powerUpIndicator2)
	lengthInBits += 8

	return lengthInBits
}

func (m *_PowerUp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PowerUpParse(ctx context.Context, theBytes []byte) (PowerUp, error) {
	return PowerUpParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PowerUpParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PowerUp, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PowerUp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PowerUp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	powerUpIndicator1, err := ReadConstField[byte](ctx, "powerUpIndicator1", ReadByte(readBuffer, 8), PowerUp_POWERUPINDICATOR1)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'powerUpIndicator1' field"))
	}
	_ = powerUpIndicator1

	powerUpIndicator2, err := ReadConstField[byte](ctx, "powerUpIndicator2", ReadByte(readBuffer, 8), PowerUp_POWERUPINDICATOR2)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'powerUpIndicator2' field"))
	}
	_ = powerUpIndicator2

	if closeErr := readBuffer.CloseContext("PowerUp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PowerUp")
	}

	// Create the instance
	return &_PowerUp{}, nil
}

func (m *_PowerUp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PowerUp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PowerUp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PowerUp")
	}

	// Const Field (powerUpIndicator1)
	_powerUpIndicator1Err := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteByte("powerUpIndicator1", 0x2B)
	if _powerUpIndicator1Err != nil {
		return errors.Wrap(_powerUpIndicator1Err, "Error serializing 'powerUpIndicator1' field")
	}

	// Const Field (powerUpIndicator2)
	_powerUpIndicator2Err := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteByte("powerUpIndicator2", 0x2B)
	if _powerUpIndicator2Err != nil {
		return errors.Wrap(_powerUpIndicator2Err, "Error serializing 'powerUpIndicator2' field")
	}

	if popErr := writeBuffer.PopContext("PowerUp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PowerUp")
	}
	return nil
}

func (m *_PowerUp) isPowerUp() bool {
	return true
}

func (m *_PowerUp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
