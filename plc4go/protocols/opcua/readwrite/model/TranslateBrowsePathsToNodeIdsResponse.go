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

// TranslateBrowsePathsToNodeIdsResponse is the corresponding interface of TranslateBrowsePathsToNodeIdsResponse
type TranslateBrowsePathsToNodeIdsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// TranslateBrowsePathsToNodeIdsResponseExactly can be used when we want exactly this type and not a type which fulfills TranslateBrowsePathsToNodeIdsResponse.
// This is useful for switch cases.
type TranslateBrowsePathsToNodeIdsResponseExactly interface {
	TranslateBrowsePathsToNodeIdsResponse
	isTranslateBrowsePathsToNodeIdsResponse() bool
}

// _TranslateBrowsePathsToNodeIdsResponse is the data-structure of this message
type _TranslateBrowsePathsToNodeIdsResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetIdentifier() string {
	return "557"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TranslateBrowsePathsToNodeIdsResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetResults() []ExtensionObjectDefinition {
	return m.Results
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTranslateBrowsePathsToNodeIdsResponse factory function for _TranslateBrowsePathsToNodeIdsResponse
func NewTranslateBrowsePathsToNodeIdsResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_TranslateBrowsePathsToNodeIdsResponse {
	_result := &_TranslateBrowsePathsToNodeIdsResponse{
		ResponseHeader:             responseHeader,
		NoOfResults:                noOfResults,
		Results:                    results,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTranslateBrowsePathsToNodeIdsResponse(structType any) TranslateBrowsePathsToNodeIdsResponse {
	if casted, ok := structType.(TranslateBrowsePathsToNodeIdsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*TranslateBrowsePathsToNodeIdsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetTypeName() string {
	return "TranslateBrowsePathsToNodeIdsResponse"
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TranslateBrowsePathsToNodeIdsResponseParse(ctx context.Context, theBytes []byte, identifier string) (TranslateBrowsePathsToNodeIdsResponse, error) {
	return TranslateBrowsePathsToNodeIdsResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TranslateBrowsePathsToNodeIdsResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TranslateBrowsePathsToNodeIdsResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TranslateBrowsePathsToNodeIdsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TranslateBrowsePathsToNodeIdsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of TranslateBrowsePathsToNodeIdsResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfResults)
	_noOfResults, _noOfResultsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfResults", 32)
	if _noOfResultsErr != nil {
		return nil, errors.Wrap(_noOfResultsErr, "Error parsing 'noOfResults' field of TranslateBrowsePathsToNodeIdsResponse")
	}
	noOfResults := _noOfResults

	results, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "results", ReadComplex[ExtensionObjectDefinition](func(ctx context.Context, buffer utils.ReadBuffer) (ExtensionObjectDefinition, error) {
		v, err := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, (string)("551"))
		if err != nil {
			return nil, err
		}
		return v.(ExtensionObjectDefinition), nil
	}, readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}

	// Simple Field (noOfDiagnosticInfos)
	_noOfDiagnosticInfos, _noOfDiagnosticInfosErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfDiagnosticInfos", 32)
	if _noOfDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfDiagnosticInfosErr, "Error parsing 'noOfDiagnosticInfos' field of TranslateBrowsePathsToNodeIdsResponse")
	}
	noOfDiagnosticInfos := _noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}

	if closeErr := readBuffer.CloseContext("TranslateBrowsePathsToNodeIdsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TranslateBrowsePathsToNodeIdsResponse")
	}

	// Create a partially initialized instance
	_child := &_TranslateBrowsePathsToNodeIdsResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfResults:                noOfResults,
		Results:                    results,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TranslateBrowsePathsToNodeIdsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TranslateBrowsePathsToNodeIdsResponse")
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

		// Simple Field (noOfResults)
		noOfResults := int32(m.GetNoOfResults())
		_noOfResultsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfResults", 32, int32((noOfResults)))
		if _noOfResultsErr != nil {
			return errors.Wrap(_noOfResultsErr, "Error serializing 'noOfResults' field")
		}

		// Array Field (results)
		if pushErr := writeBuffer.PushContext("results", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for results")
		}
		for _curItem, _element := range m.GetResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'results' field")
			}
		}
		if popErr := writeBuffer.PopContext("results", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for results")
		}

		// Simple Field (noOfDiagnosticInfos)
		noOfDiagnosticInfos := int32(m.GetNoOfDiagnosticInfos())
		_noOfDiagnosticInfosErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfDiagnosticInfos", 32, int32((noOfDiagnosticInfos)))
		if _noOfDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfDiagnosticInfosErr, "Error serializing 'noOfDiagnosticInfos' field")
		}

		// Array Field (diagnosticInfos)
		if pushErr := writeBuffer.PushContext("diagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for diagnosticInfos")
		}
		for _curItem, _element := range m.GetDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'diagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("diagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for diagnosticInfos")
		}

		if popErr := writeBuffer.PopContext("TranslateBrowsePathsToNodeIdsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TranslateBrowsePathsToNodeIdsResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) isTranslateBrowsePathsToNodeIdsResponse() bool {
	return true
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
