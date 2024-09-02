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

// ErrorReportingSystemCategoryTypeSupportUnits is the corresponding interface of ErrorReportingSystemCategoryTypeSupportUnits
type ErrorReportingSystemCategoryTypeSupportUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForSupportUnits
}

// ErrorReportingSystemCategoryTypeSupportUnitsExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategoryTypeSupportUnits.
// This is useful for switch cases.
type ErrorReportingSystemCategoryTypeSupportUnitsExactly interface {
	ErrorReportingSystemCategoryTypeSupportUnits
	isErrorReportingSystemCategoryTypeSupportUnits() bool
}

// _ErrorReportingSystemCategoryTypeSupportUnits is the data-structure of this message
type _ErrorReportingSystemCategoryTypeSupportUnits struct {
	ErrorReportingSystemCategoryTypeContract
	CategoryForType ErrorReportingSystemCategoryTypeForSupportUnits
}

var _ ErrorReportingSystemCategoryTypeSupportUnits = (*_ErrorReportingSystemCategoryTypeSupportUnits)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return ErrorReportingSystemCategoryClass_SUPPORT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetParent() ErrorReportingSystemCategoryTypeContract {
	return m.ErrorReportingSystemCategoryTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetCategoryForType() ErrorReportingSystemCategoryTypeForSupportUnits {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingSystemCategoryTypeSupportUnits factory function for _ErrorReportingSystemCategoryTypeSupportUnits
func NewErrorReportingSystemCategoryTypeSupportUnits(categoryForType ErrorReportingSystemCategoryTypeForSupportUnits) *_ErrorReportingSystemCategoryTypeSupportUnits {
	_result := &_ErrorReportingSystemCategoryTypeSupportUnits{
		ErrorReportingSystemCategoryTypeContract: NewErrorReportingSystemCategoryType(),
		CategoryForType:                          categoryForType,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeSupportUnits(structType any) ErrorReportingSystemCategoryTypeSupportUnits {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeSupportUnits); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeSupportUnits); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeSupportUnits"
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).getLengthInBits(ctx))

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ErrorReportingSystemCategoryType, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryTypeSupportUnits ErrorReportingSystemCategoryTypeSupportUnits, err error) {
	m.ErrorReportingSystemCategoryTypeContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeSupportUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeSupportUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryForType, err := ReadEnumField[ErrorReportingSystemCategoryTypeForSupportUnits](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForSupportUnits", ReadEnum(ErrorReportingSystemCategoryTypeForSupportUnitsByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryForType' field"))
	}
	m.CategoryForType = categoryForType

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeSupportUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeSupportUnits")
	}

	return m, nil
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeSupportUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeSupportUnits")
		}

		if err := WriteSimpleEnumField[ErrorReportingSystemCategoryTypeForSupportUnits](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForSupportUnits", m.GetCategoryForType(), WriteEnum[ErrorReportingSystemCategoryTypeForSupportUnits, uint8](ErrorReportingSystemCategoryTypeForSupportUnits.GetValue, ErrorReportingSystemCategoryTypeForSupportUnits.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeSupportUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeSupportUnits")
		}
		return nil
	}
	return m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) isErrorReportingSystemCategoryTypeSupportUnits() bool {
	return true
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
