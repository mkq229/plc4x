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

// BACnetConstructedDataStopWhenFull is the corresponding interface of BACnetConstructedDataStopWhenFull
type BACnetConstructedDataStopWhenFull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStopWhenFull returns StopWhenFull (property field)
	GetStopWhenFull() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataStopWhenFullExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataStopWhenFull.
// This is useful for switch cases.
type BACnetConstructedDataStopWhenFullExactly interface {
	BACnetConstructedDataStopWhenFull
	isBACnetConstructedDataStopWhenFull() bool
}

// _BACnetConstructedDataStopWhenFull is the data-structure of this message
type _BACnetConstructedDataStopWhenFull struct {
	BACnetConstructedDataContract
	StopWhenFull BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataStopWhenFull = (*_BACnetConstructedDataStopWhenFull)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStopWhenFull) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_WHEN_FULL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetStopWhenFull() BACnetApplicationTagBoolean {
	return m.StopWhenFull
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetStopWhenFull())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStopWhenFull factory function for _BACnetConstructedDataStopWhenFull
func NewBACnetConstructedDataStopWhenFull(stopWhenFull BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStopWhenFull {
	_result := &_BACnetConstructedDataStopWhenFull{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StopWhenFull:                  stopWhenFull,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStopWhenFull(structType any) BACnetConstructedDataStopWhenFull {
	if casted, ok := structType.(BACnetConstructedDataStopWhenFull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopWhenFull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStopWhenFull) GetTypeName() string {
	return "BACnetConstructedDataStopWhenFull"
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (stopWhenFull)
	lengthInBits += m.StopWhenFull.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStopWhenFull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStopWhenFull BACnetConstructedDataStopWhenFull, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopWhenFull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStopWhenFull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	stopWhenFull, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "stopWhenFull", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stopWhenFull' field"))
	}
	m.StopWhenFull = stopWhenFull

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), stopWhenFull)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopWhenFull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStopWhenFull")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStopWhenFull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStopWhenFull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopWhenFull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStopWhenFull")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "stopWhenFull", m.GetStopWhenFull(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'stopWhenFull' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopWhenFull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStopWhenFull")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStopWhenFull) isBACnetConstructedDataStopWhenFull() bool {
	return true
}

func (m *_BACnetConstructedDataStopWhenFull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
