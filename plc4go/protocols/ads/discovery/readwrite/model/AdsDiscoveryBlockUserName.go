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

// AdsDiscoveryBlockUserName is the corresponding interface of AdsDiscoveryBlockUserName
type AdsDiscoveryBlockUserName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetUserName returns UserName (property field)
	GetUserName() AmsString
}

// AdsDiscoveryBlockUserNameExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockUserName.
// This is useful for switch cases.
type AdsDiscoveryBlockUserNameExactly interface {
	AdsDiscoveryBlockUserName
	isAdsDiscoveryBlockUserName() bool
}

// _AdsDiscoveryBlockUserName is the data-structure of this message
type _AdsDiscoveryBlockUserName struct {
	AdsDiscoveryBlockContract
	UserName AmsString
}

var _ AdsDiscoveryBlockUserName = (*_AdsDiscoveryBlockUserName)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockUserName) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_USER_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockUserName) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockUserName) GetUserName() AmsString {
	return m.UserName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockUserName factory function for _AdsDiscoveryBlockUserName
func NewAdsDiscoveryBlockUserName(userName AmsString) *_AdsDiscoveryBlockUserName {
	_result := &_AdsDiscoveryBlockUserName{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		UserName:                  userName,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockUserName(structType any) AdsDiscoveryBlockUserName {
	if casted, ok := structType.(AdsDiscoveryBlockUserName); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockUserName); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockUserName) GetTypeName() string {
	return "AdsDiscoveryBlockUserName"
}

func (m *_AdsDiscoveryBlockUserName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockUserName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockUserName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockUserName AdsDiscoveryBlockUserName, err error) {
	m.AdsDiscoveryBlockContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockUserName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockUserName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	userName, err := ReadSimpleField[AmsString](ctx, "userName", ReadComplex[AmsString](AmsStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userName' field"))
	}
	m.UserName = userName

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockUserName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockUserName")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockUserName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockUserName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockUserName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockUserName")
		}

		if err := WriteSimpleField[AmsString](ctx, "userName", m.GetUserName(), WriteComplex[AmsString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userName' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockUserName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockUserName")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockUserName) isAdsDiscoveryBlockUserName() bool {
	return true
}

func (m *_AdsDiscoveryBlockUserName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
