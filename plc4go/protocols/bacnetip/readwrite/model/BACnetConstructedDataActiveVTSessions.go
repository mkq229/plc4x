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

// BACnetConstructedDataActiveVTSessions is the corresponding interface of BACnetConstructedDataActiveVTSessions
type BACnetConstructedDataActiveVTSessions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActiveVTSession returns ActiveVTSession (property field)
	GetActiveVTSession() []BACnetVTSession
}

// BACnetConstructedDataActiveVTSessionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActiveVTSessions.
// This is useful for switch cases.
type BACnetConstructedDataActiveVTSessionsExactly interface {
	BACnetConstructedDataActiveVTSessions
	isBACnetConstructedDataActiveVTSessions() bool
}

// _BACnetConstructedDataActiveVTSessions is the data-structure of this message
type _BACnetConstructedDataActiveVTSessions struct {
	BACnetConstructedDataContract
	ActiveVTSession []BACnetVTSession
}

var _ BACnetConstructedDataActiveVTSessions = (*_BACnetConstructedDataActiveVTSessions)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActiveVTSessions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActiveVTSessions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_VT_SESSIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActiveVTSessions) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActiveVTSessions) GetActiveVTSession() []BACnetVTSession {
	return m.ActiveVTSession
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveVTSessions factory function for _BACnetConstructedDataActiveVTSessions
func NewBACnetConstructedDataActiveVTSessions(activeVTSession []BACnetVTSession, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActiveVTSessions {
	_result := &_BACnetConstructedDataActiveVTSessions{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ActiveVTSession:               activeVTSession,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActiveVTSessions(structType any) BACnetConstructedDataActiveVTSessions {
	if casted, ok := structType.(BACnetConstructedDataActiveVTSessions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveVTSessions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActiveVTSessions) GetTypeName() string {
	return "BACnetConstructedDataActiveVTSessions"
}

func (m *_BACnetConstructedDataActiveVTSessions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ActiveVTSession) > 0 {
		for _, element := range m.ActiveVTSession {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActiveVTSessions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataActiveVTSessions) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataActiveVTSessions BACnetConstructedDataActiveVTSessions, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveVTSessions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveVTSessions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	activeVTSession, err := ReadTerminatedArrayField[BACnetVTSession](ctx, "activeVTSession", ReadComplex[BACnetVTSession](BACnetVTSessionParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'activeVTSession' field"))
	}
	m.ActiveVTSession = activeVTSession

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveVTSessions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveVTSessions")
	}

	return m, nil
}

func (m *_BACnetConstructedDataActiveVTSessions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActiveVTSessions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveVTSessions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveVTSessions")
		}

		if err := WriteComplexTypeArrayField(ctx, "activeVTSession", m.GetActiveVTSession(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'activeVTSession' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveVTSessions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveVTSessions")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActiveVTSessions) isBACnetConstructedDataActiveVTSessions() bool {
	return true
}

func (m *_BACnetConstructedDataActiveVTSessions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
