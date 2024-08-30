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

// Code generated by "plc4xgenerator -type=DefaultPlcDiscoveryItem"; DO NOT EDIT.

package model

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DefaultPlcDiscoveryItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcDiscoveryItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcDiscoveryItem"); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("protocolCode", uint32(len(d.ProtocolCode)*8), d.ProtocolCode); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("transportCode", uint32(len(d.TransportCode)*8), d.TransportCode); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("options", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _name, elem := range d.Options {
		name := _name
		_value := fmt.Sprintf("%v", elem)

		if err := writeBuffer.WriteString(name, uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("options", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("name", uint32(len(d.Name)*8), d.Name); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("attributes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _name, elem := range d.Attributes {
		name := _name

		var elem any = elem
		if serializable, ok := elem.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(name); err != nil {
				return err
			}
			if err := serializable.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(name); err != nil {
				return err
			}
		} else {
			elemAsString := fmt.Sprintf("%v", elem)
			if err := writeBuffer.WriteString(name, uint32(len(elemAsString)*8), elemAsString); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("attributes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("PlcDiscoveryItem"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcDiscoveryItem) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
