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

// BACnetPropertyStatesUnits is the corresponding interface of BACnetPropertyStatesUnits
type BACnetPropertyStatesUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
}

// BACnetPropertyStatesUnitsExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesUnits.
// This is useful for switch cases.
type BACnetPropertyStatesUnitsExactly interface {
	BACnetPropertyStatesUnits
	isBACnetPropertyStatesUnits() bool
}

// _BACnetPropertyStatesUnits is the data-structure of this message
type _BACnetPropertyStatesUnits struct {
	BACnetPropertyStatesContract
	Units BACnetEngineeringUnitsTagged
}

var _ BACnetPropertyStatesUnits = (*_BACnetPropertyStatesUnits)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesUnits) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesUnits factory function for _BACnetPropertyStatesUnits
func NewBACnetPropertyStatesUnits(units BACnetEngineeringUnitsTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesUnits {
	_result := &_BACnetPropertyStatesUnits{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Units:                        units,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesUnits(structType any) BACnetPropertyStatesUnits {
	if casted, ok := structType.(BACnetPropertyStatesUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesUnits) GetTypeName() string {
	return "BACnetPropertyStatesUnits"
}

func (m *_BACnetPropertyStatesUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesUnits BACnetPropertyStatesUnits, err error) {
	m.BACnetPropertyStatesContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesUnits")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesUnits")
		}

		if err := WriteSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", m.GetUnits(), WriteComplex[BACnetEngineeringUnitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesUnits")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesUnits) isBACnetPropertyStatesUnits() bool {
	return true
}

func (m *_BACnetPropertyStatesUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
