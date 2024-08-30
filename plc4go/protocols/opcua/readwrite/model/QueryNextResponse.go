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

// QueryNextResponse is the corresponding interface of QueryNextResponse
type QueryNextResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfQueryDataSets returns NoOfQueryDataSets (property field)
	GetNoOfQueryDataSets() int32
	// GetQueryDataSets returns QueryDataSets (property field)
	GetQueryDataSets() []ExtensionObjectDefinition
	// GetRevisedContinuationPoint returns RevisedContinuationPoint (property field)
	GetRevisedContinuationPoint() PascalByteString
}

// QueryNextResponseExactly can be used when we want exactly this type and not a type which fulfills QueryNextResponse.
// This is useful for switch cases.
type QueryNextResponseExactly interface {
	QueryNextResponse
	isQueryNextResponse() bool
}

// _QueryNextResponse is the data-structure of this message
type _QueryNextResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader           ExtensionObjectDefinition
	NoOfQueryDataSets        int32
	QueryDataSets            []ExtensionObjectDefinition
	RevisedContinuationPoint PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryNextResponse) GetIdentifier() string {
	return "624"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryNextResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_QueryNextResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryNextResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_QueryNextResponse) GetNoOfQueryDataSets() int32 {
	return m.NoOfQueryDataSets
}

func (m *_QueryNextResponse) GetQueryDataSets() []ExtensionObjectDefinition {
	return m.QueryDataSets
}

func (m *_QueryNextResponse) GetRevisedContinuationPoint() PascalByteString {
	return m.RevisedContinuationPoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQueryNextResponse factory function for _QueryNextResponse
func NewQueryNextResponse(responseHeader ExtensionObjectDefinition, noOfQueryDataSets int32, queryDataSets []ExtensionObjectDefinition, revisedContinuationPoint PascalByteString) *_QueryNextResponse {
	_result := &_QueryNextResponse{
		ResponseHeader:             responseHeader,
		NoOfQueryDataSets:          noOfQueryDataSets,
		QueryDataSets:              queryDataSets,
		RevisedContinuationPoint:   revisedContinuationPoint,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryNextResponse(structType any) QueryNextResponse {
	if casted, ok := structType.(QueryNextResponse); ok {
		return casted
	}
	if casted, ok := structType.(*QueryNextResponse); ok {
		return *casted
	}
	return nil
}

func (m *_QueryNextResponse) GetTypeName() string {
	return "QueryNextResponse"
}

func (m *_QueryNextResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfQueryDataSets)
	lengthInBits += 32

	// Array field
	if len(m.QueryDataSets) > 0 {
		for _curItem, element := range m.QueryDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.QueryDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (revisedContinuationPoint)
	lengthInBits += m.RevisedContinuationPoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QueryNextResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QueryNextResponseParse(ctx context.Context, theBytes []byte, identifier string) (QueryNextResponse, error) {
	return QueryNextResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QueryNextResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QueryNextResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QueryNextResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryNextResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of QueryNextResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfQueryDataSets)
	_noOfQueryDataSets, _noOfQueryDataSetsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfQueryDataSets", 32)
	if _noOfQueryDataSetsErr != nil {
		return nil, errors.Wrap(_noOfQueryDataSetsErr, "Error parsing 'noOfQueryDataSets' field of QueryNextResponse")
	}
	noOfQueryDataSets := _noOfQueryDataSets

	queryDataSets, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "queryDataSets", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("579"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfQueryDataSets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queryDataSets' field"))
	}

	// Simple Field (revisedContinuationPoint)
	if pullErr := readBuffer.PullContext("revisedContinuationPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for revisedContinuationPoint")
	}
	_revisedContinuationPoint, _revisedContinuationPointErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _revisedContinuationPointErr != nil {
		return nil, errors.Wrap(_revisedContinuationPointErr, "Error parsing 'revisedContinuationPoint' field of QueryNextResponse")
	}
	revisedContinuationPoint := _revisedContinuationPoint.(PascalByteString)
	if closeErr := readBuffer.CloseContext("revisedContinuationPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for revisedContinuationPoint")
	}

	if closeErr := readBuffer.CloseContext("QueryNextResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryNextResponse")
	}

	// Create a partially initialized instance
	_child := &_QueryNextResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfQueryDataSets:          noOfQueryDataSets,
		QueryDataSets:              queryDataSets,
		RevisedContinuationPoint:   revisedContinuationPoint,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QueryNextResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryNextResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryNextResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryNextResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (noOfQueryDataSets)
		noOfQueryDataSets := int32(m.GetNoOfQueryDataSets())
		_noOfQueryDataSetsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfQueryDataSets", 32, int32((noOfQueryDataSets)))
		if _noOfQueryDataSetsErr != nil {
			return errors.Wrap(_noOfQueryDataSetsErr, "Error serializing 'noOfQueryDataSets' field")
		}

		// Array Field (queryDataSets)
		if pushErr := writeBuffer.PushContext("queryDataSets", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for queryDataSets")
		}
		for _curItem, _element := range m.GetQueryDataSets() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetQueryDataSets()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'queryDataSets' field")
			}
		}
		if popErr := writeBuffer.PopContext("queryDataSets", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for queryDataSets")
		}

		// Simple Field (revisedContinuationPoint)
		if pushErr := writeBuffer.PushContext("revisedContinuationPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for revisedContinuationPoint")
		}
		_revisedContinuationPointErr := writeBuffer.WriteSerializable(ctx, m.GetRevisedContinuationPoint())
		if popErr := writeBuffer.PopContext("revisedContinuationPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for revisedContinuationPoint")
		}
		if _revisedContinuationPointErr != nil {
			return errors.Wrap(_revisedContinuationPointErr, "Error serializing 'revisedContinuationPoint' field")
		}

		if popErr := writeBuffer.PopContext("QueryNextResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryNextResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryNextResponse) isQueryNextResponse() bool {
	return true
}

func (m *_QueryNextResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
