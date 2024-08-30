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

// ApplicationInstanceCertificate is the corresponding interface of ApplicationInstanceCertificate
type ApplicationInstanceCertificate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ApplicationInstanceCertificateExactly can be used when we want exactly this type and not a type which fulfills ApplicationInstanceCertificate.
// This is useful for switch cases.
type ApplicationInstanceCertificateExactly interface {
	ApplicationInstanceCertificate
	isApplicationInstanceCertificate() bool
}

// _ApplicationInstanceCertificate is the data-structure of this message
type _ApplicationInstanceCertificate struct {
}

// NewApplicationInstanceCertificate factory function for _ApplicationInstanceCertificate
func NewApplicationInstanceCertificate() *_ApplicationInstanceCertificate {
	return &_ApplicationInstanceCertificate{}
}

// Deprecated: use the interface for direct cast
func CastApplicationInstanceCertificate(structType any) ApplicationInstanceCertificate {
	if casted, ok := structType.(ApplicationInstanceCertificate); ok {
		return casted
	}
	if casted, ok := structType.(*ApplicationInstanceCertificate); ok {
		return *casted
	}
	return nil
}

func (m *_ApplicationInstanceCertificate) GetTypeName() string {
	return "ApplicationInstanceCertificate"
}

func (m *_ApplicationInstanceCertificate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ApplicationInstanceCertificate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApplicationInstanceCertificateParse(ctx context.Context, theBytes []byte) (ApplicationInstanceCertificate, error) {
	return ApplicationInstanceCertificateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ApplicationInstanceCertificateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApplicationInstanceCertificate, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApplicationInstanceCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApplicationInstanceCertificate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApplicationInstanceCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApplicationInstanceCertificate")
	}

	// Create the instance
	return &_ApplicationInstanceCertificate{}, nil
}

func (m *_ApplicationInstanceCertificate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApplicationInstanceCertificate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApplicationInstanceCertificate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApplicationInstanceCertificate")
	}

	if popErr := writeBuffer.PopContext("ApplicationInstanceCertificate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApplicationInstanceCertificate")
	}
	return nil
}

func (m *_ApplicationInstanceCertificate) isApplicationInstanceCertificate() bool {
	return true
}

func (m *_ApplicationInstanceCertificate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
