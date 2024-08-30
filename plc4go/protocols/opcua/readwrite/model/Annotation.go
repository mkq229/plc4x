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

// Annotation is the corresponding interface of Annotation
type Annotation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMessage returns Message (property field)
	GetMessage() PascalString
	// GetUserName returns UserName (property field)
	GetUserName() PascalString
	// GetAnnotationTime returns AnnotationTime (property field)
	GetAnnotationTime() int64
}

// AnnotationExactly can be used when we want exactly this type and not a type which fulfills Annotation.
// This is useful for switch cases.
type AnnotationExactly interface {
	Annotation
	isAnnotation() bool
}

// _Annotation is the data-structure of this message
type _Annotation struct {
	*_ExtensionObjectDefinition
	Message        PascalString
	UserName       PascalString
	AnnotationTime int64
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Annotation) GetIdentifier() string {
	return "893"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Annotation) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_Annotation) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Annotation) GetMessage() PascalString {
	return m.Message
}

func (m *_Annotation) GetUserName() PascalString {
	return m.UserName
}

func (m *_Annotation) GetAnnotationTime() int64 {
	return m.AnnotationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAnnotation factory function for _Annotation
func NewAnnotation(message PascalString, userName PascalString, annotationTime int64) *_Annotation {
	_result := &_Annotation{
		Message:                    message,
		UserName:                   userName,
		AnnotationTime:             annotationTime,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAnnotation(structType any) Annotation {
	if casted, ok := structType.(Annotation); ok {
		return casted
	}
	if casted, ok := structType.(*Annotation); ok {
		return *casted
	}
	return nil
}

func (m *_Annotation) GetTypeName() string {
	return "Annotation"
}

func (m *_Annotation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	// Simple field (annotationTime)
	lengthInBits += 64

	return lengthInBits
}

func (m *_Annotation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AnnotationParse(ctx context.Context, theBytes []byte, identifier string) (Annotation, error) {
	return AnnotationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AnnotationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (Annotation, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Annotation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Annotation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (message)
	if pullErr := readBuffer.PullContext("message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for message")
	}
	_message, _messageErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field of Annotation")
	}
	message := _message.(PascalString)
	if closeErr := readBuffer.CloseContext("message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for message")
	}

	// Simple Field (userName)
	if pullErr := readBuffer.PullContext("userName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userName")
	}
	_userName, _userNameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _userNameErr != nil {
		return nil, errors.Wrap(_userNameErr, "Error parsing 'userName' field of Annotation")
	}
	userName := _userName.(PascalString)
	if closeErr := readBuffer.CloseContext("userName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userName")
	}

	// Simple Field (annotationTime)
	_annotationTime, _annotationTimeErr := /*TODO: migrate me*/ readBuffer.ReadInt64("annotationTime", 64)
	if _annotationTimeErr != nil {
		return nil, errors.Wrap(_annotationTimeErr, "Error parsing 'annotationTime' field of Annotation")
	}
	annotationTime := _annotationTime

	if closeErr := readBuffer.CloseContext("Annotation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Annotation")
	}

	// Create a partially initialized instance
	_child := &_Annotation{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Message:                    message,
		UserName:                   userName,
		AnnotationTime:             annotationTime,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_Annotation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Annotation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Annotation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Annotation")
		}

		// Simple Field (message)
		if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for message")
		}
		_messageErr := writeBuffer.WriteSerializable(ctx, m.GetMessage())
		if popErr := writeBuffer.PopContext("message"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for message")
		}
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		// Simple Field (userName)
		if pushErr := writeBuffer.PushContext("userName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userName")
		}
		_userNameErr := writeBuffer.WriteSerializable(ctx, m.GetUserName())
		if popErr := writeBuffer.PopContext("userName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userName")
		}
		if _userNameErr != nil {
			return errors.Wrap(_userNameErr, "Error serializing 'userName' field")
		}

		// Simple Field (annotationTime)
		annotationTime := int64(m.GetAnnotationTime())
		_annotationTimeErr := /*TODO: migrate me*/ writeBuffer.WriteInt64("annotationTime", 64, int64((annotationTime)))
		if _annotationTimeErr != nil {
			return errors.Wrap(_annotationTimeErr, "Error serializing 'annotationTime' field")
		}

		if popErr := writeBuffer.PopContext("Annotation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Annotation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Annotation) isAnnotation() bool {
	return true
}

func (m *_Annotation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
