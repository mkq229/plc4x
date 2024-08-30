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

// HVACZoneList is the corresponding interface of HVACZoneList
type HVACZoneList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetExpansion returns Expansion (property field)
	GetExpansion() bool
	// GetZone6 returns Zone6 (property field)
	GetZone6() bool
	// GetZone5 returns Zone5 (property field)
	GetZone5() bool
	// GetZone4 returns Zone4 (property field)
	GetZone4() bool
	// GetZone3 returns Zone3 (property field)
	GetZone3() bool
	// GetZone2 returns Zone2 (property field)
	GetZone2() bool
	// GetZone1 returns Zone1 (property field)
	GetZone1() bool
	// GetZone0 returns Zone0 (property field)
	GetZone0() bool
	// GetUnswitchedZone returns UnswitchedZone (virtual field)
	GetUnswitchedZone() bool
}

// HVACZoneListExactly can be used when we want exactly this type and not a type which fulfills HVACZoneList.
// This is useful for switch cases.
type HVACZoneListExactly interface {
	HVACZoneList
	isHVACZoneList() bool
}

// _HVACZoneList is the data-structure of this message
type _HVACZoneList struct {
	Expansion bool
	Zone6     bool
	Zone5     bool
	Zone4     bool
	Zone3     bool
	Zone2     bool
	Zone1     bool
	Zone0     bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACZoneList) GetExpansion() bool {
	return m.Expansion
}

func (m *_HVACZoneList) GetZone6() bool {
	return m.Zone6
}

func (m *_HVACZoneList) GetZone5() bool {
	return m.Zone5
}

func (m *_HVACZoneList) GetZone4() bool {
	return m.Zone4
}

func (m *_HVACZoneList) GetZone3() bool {
	return m.Zone3
}

func (m *_HVACZoneList) GetZone2() bool {
	return m.Zone2
}

func (m *_HVACZoneList) GetZone1() bool {
	return m.Zone1
}

func (m *_HVACZoneList) GetZone0() bool {
	return m.Zone0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACZoneList) GetUnswitchedZone() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetZone0())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACZoneList factory function for _HVACZoneList
func NewHVACZoneList(expansion bool, zone6 bool, zone5 bool, zone4 bool, zone3 bool, zone2 bool, zone1 bool, zone0 bool) *_HVACZoneList {
	return &_HVACZoneList{Expansion: expansion, Zone6: zone6, Zone5: zone5, Zone4: zone4, Zone3: zone3, Zone2: zone2, Zone1: zone1, Zone0: zone0}
}

// Deprecated: use the interface for direct cast
func CastHVACZoneList(structType any) HVACZoneList {
	if casted, ok := structType.(HVACZoneList); ok {
		return casted
	}
	if casted, ok := structType.(*HVACZoneList); ok {
		return *casted
	}
	return nil
}

func (m *_HVACZoneList) GetTypeName() string {
	return "HVACZoneList"
}

func (m *_HVACZoneList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (expansion)
	lengthInBits += 1

	// Simple field (zone6)
	lengthInBits += 1

	// Simple field (zone5)
	lengthInBits += 1

	// Simple field (zone4)
	lengthInBits += 1

	// Simple field (zone3)
	lengthInBits += 1

	// Simple field (zone2)
	lengthInBits += 1

	// Simple field (zone1)
	lengthInBits += 1

	// Simple field (zone0)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACZoneList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACZoneListParse(ctx context.Context, theBytes []byte) (HVACZoneList, error) {
	return HVACZoneListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACZoneListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACZoneList, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HVACZoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACZoneList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (expansion)
	_expansion, _expansionErr := /*TODO: migrate me*/ readBuffer.ReadBit("expansion")
	if _expansionErr != nil {
		return nil, errors.Wrap(_expansionErr, "Error parsing 'expansion' field of HVACZoneList")
	}
	expansion := _expansion

	// Simple Field (zone6)
	_zone6, _zone6Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone6")
	if _zone6Err != nil {
		return nil, errors.Wrap(_zone6Err, "Error parsing 'zone6' field of HVACZoneList")
	}
	zone6 := _zone6

	// Simple Field (zone5)
	_zone5, _zone5Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone5")
	if _zone5Err != nil {
		return nil, errors.Wrap(_zone5Err, "Error parsing 'zone5' field of HVACZoneList")
	}
	zone5 := _zone5

	// Simple Field (zone4)
	_zone4, _zone4Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone4")
	if _zone4Err != nil {
		return nil, errors.Wrap(_zone4Err, "Error parsing 'zone4' field of HVACZoneList")
	}
	zone4 := _zone4

	// Simple Field (zone3)
	_zone3, _zone3Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone3")
	if _zone3Err != nil {
		return nil, errors.Wrap(_zone3Err, "Error parsing 'zone3' field of HVACZoneList")
	}
	zone3 := _zone3

	// Simple Field (zone2)
	_zone2, _zone2Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone2")
	if _zone2Err != nil {
		return nil, errors.Wrap(_zone2Err, "Error parsing 'zone2' field of HVACZoneList")
	}
	zone2 := _zone2

	// Simple Field (zone1)
	_zone1, _zone1Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone1")
	if _zone1Err != nil {
		return nil, errors.Wrap(_zone1Err, "Error parsing 'zone1' field of HVACZoneList")
	}
	zone1 := _zone1

	// Simple Field (zone0)
	_zone0, _zone0Err := /*TODO: migrate me*/ readBuffer.ReadBit("zone0")
	if _zone0Err != nil {
		return nil, errors.Wrap(_zone0Err, "Error parsing 'zone0' field of HVACZoneList")
	}
	zone0 := _zone0

	// Virtual field
	_unswitchedZone := zone0
	unswitchedZone := bool(_unswitchedZone)
	_ = unswitchedZone

	if closeErr := readBuffer.CloseContext("HVACZoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACZoneList")
	}

	// Create the instance
	return &_HVACZoneList{
		Expansion: expansion,
		Zone6:     zone6,
		Zone5:     zone5,
		Zone4:     zone4,
		Zone3:     zone3,
		Zone2:     zone2,
		Zone1:     zone1,
		Zone0:     zone0,
	}, nil
}

func (m *_HVACZoneList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACZoneList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACZoneList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACZoneList")
	}

	// Simple Field (expansion)
	expansion := bool(m.GetExpansion())
	_expansionErr := /*TODO: migrate me*/ writeBuffer.WriteBit("expansion", (expansion))
	if _expansionErr != nil {
		return errors.Wrap(_expansionErr, "Error serializing 'expansion' field")
	}

	// Simple Field (zone6)
	zone6 := bool(m.GetZone6())
	_zone6Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone6", (zone6))
	if _zone6Err != nil {
		return errors.Wrap(_zone6Err, "Error serializing 'zone6' field")
	}

	// Simple Field (zone5)
	zone5 := bool(m.GetZone5())
	_zone5Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone5", (zone5))
	if _zone5Err != nil {
		return errors.Wrap(_zone5Err, "Error serializing 'zone5' field")
	}

	// Simple Field (zone4)
	zone4 := bool(m.GetZone4())
	_zone4Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone4", (zone4))
	if _zone4Err != nil {
		return errors.Wrap(_zone4Err, "Error serializing 'zone4' field")
	}

	// Simple Field (zone3)
	zone3 := bool(m.GetZone3())
	_zone3Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone3", (zone3))
	if _zone3Err != nil {
		return errors.Wrap(_zone3Err, "Error serializing 'zone3' field")
	}

	// Simple Field (zone2)
	zone2 := bool(m.GetZone2())
	_zone2Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone2", (zone2))
	if _zone2Err != nil {
		return errors.Wrap(_zone2Err, "Error serializing 'zone2' field")
	}

	// Simple Field (zone1)
	zone1 := bool(m.GetZone1())
	_zone1Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone1", (zone1))
	if _zone1Err != nil {
		return errors.Wrap(_zone1Err, "Error serializing 'zone1' field")
	}

	// Simple Field (zone0)
	zone0 := bool(m.GetZone0())
	_zone0Err := /*TODO: migrate me*/ writeBuffer.WriteBit("zone0", (zone0))
	if _zone0Err != nil {
		return errors.Wrap(_zone0Err, "Error serializing 'zone0' field")
	}
	// Virtual field
	unswitchedZone := m.GetUnswitchedZone()
	_ = unswitchedZone
	if _unswitchedZoneErr := writeBuffer.WriteVirtual(ctx, "unswitchedZone", m.GetUnswitchedZone()); _unswitchedZoneErr != nil {
		return errors.Wrap(_unswitchedZoneErr, "Error serializing 'unswitchedZone' field")
	}

	if popErr := writeBuffer.PopContext("HVACZoneList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACZoneList")
	}
	return nil
}

func (m *_HVACZoneList) isHVACZoneList() bool {
	return true
}

func (m *_HVACZoneList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
