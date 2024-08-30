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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	api "github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/spi/values"
)

// Code generated by code-generation. DO NOT EDIT.

func DataItemParse(ctx context.Context, theBytes []byte, dataType string, numberOfValues uint16) (api.PlcValue, error) {
	return DataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataType, numberOfValues)
}

func DataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataType string, numberOfValues uint16) (api.PlcValue, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	readBuffer.PullContext("DataItem")
	switch {
	case dataType == "BOOL" && numberOfValues == uint16(1): // BOOL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadBit("value")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBOOL(value), nil
	case dataType == "BOOL": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadBit("value")
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcBOOL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "BYTE" && numberOfValues == uint16(1): // BYTE
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBYTE(value), nil
	case dataType == "BYTE": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint8("value", 8)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUSINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "WORD" && numberOfValues == uint16(1): // WORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcWORD(value), nil
	case dataType == "WORD": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint16("value", 16)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "DWORD" && numberOfValues == uint16(1): // DWORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDWORD(value), nil
	case dataType == "DWORD": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUDINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "LWORD" && numberOfValues == uint16(1): // LWORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLWORD(value), nil
	case dataType == "LWORD": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcULINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "SINT" && numberOfValues == uint16(1): // SINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSINT(value), nil
	case dataType == "SINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadInt8("value", 8)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcSINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "INT" && numberOfValues == uint16(1): // INT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcINT(value), nil
	case dataType == "INT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadInt16("value", 16)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "DINT" && numberOfValues == uint16(1): // DINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDINT(value), nil
	case dataType == "DINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadInt32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcDINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "LINT" && numberOfValues == uint16(1): // LINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLINT(value), nil
	case dataType == "LINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadInt64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcLINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "USINT" && numberOfValues == uint16(1): // USINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUSINT(value), nil
	case dataType == "USINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint8("value", 8)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUSINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "UINT" && numberOfValues == uint16(1): // UINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUINT(value), nil
	case dataType == "UINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint16("value", 16)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "UDINT" && numberOfValues == uint16(1): // UDINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUDINT(value), nil
	case dataType == "UDINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcUDINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "ULINT" && numberOfValues == uint16(1): // ULINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcULINT(value), nil
	case dataType == "ULINT": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadUint64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcULINT(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "REAL" && numberOfValues == uint16(1): // REAL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcREAL(value), nil
	case dataType == "REAL": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadFloat32("value", 32)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcREAL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "LREAL" && numberOfValues == uint16(1): // LREAL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLREAL(value), nil
	case dataType == "LREAL": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadFloat64("value", 64)
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcLREAL(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "CHAR" && numberOfValues == uint16(1): // CHAR
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32(8), utils.WithEncoding("UTF-8"))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcCHAR(value), nil
	case dataType == "CHAR": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadString("value", uint32(8), utils.WithEncoding("UTF-8"))
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcSTRING(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "WCHAR" && numberOfValues == uint16(1): // WCHAR
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32(16), utils.WithEncoding("UTF-16"))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcWCHAR(value), nil
	case dataType == "WCHAR": // List
		// Array Field (value)
		var value []api.PlcValue
		for i := 0; i < int(numberOfValues); i++ {
			_item, _itemErr := readBuffer.ReadString("value", uint32(16), utils.WithEncoding("UTF-16"))
			if _itemErr != nil {
				return nil, errors.Wrap(_itemErr, "Error parsing 'value' field")
			}
			value = append(value, values.NewPlcSTRING(_item))
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcList(value), nil
	case dataType == "STRING": // STRING
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32(255), utils.WithEncoding("UTF-8"))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataType == "WSTRING": // STRING
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32(255), utils.WithEncoding("UTF-16"))
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	}
	// TODO: add more info which type it is actually
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(value api.PlcValue, dataType string, numberOfValues uint16) ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := DataItemSerializeWithWriteBuffer(context.Background(), wb, value, dataType, numberOfValues); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func DataItemSerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer, value api.PlcValue, dataType string, numberOfValues uint16) error {
	log := zerolog.Ctx(ctx)
	_ = log
	m := struct {
		DataType       string
		NumberOfValues uint16
	}{
		DataType:       dataType,
		NumberOfValues: numberOfValues,
	}
	_ = m
	writeBuffer.PushContext("DataItem")
	switch {
	case dataType == "BOOL" && numberOfValues == uint16(1): // BOOL
		// Simple Field (value)
		if _err := writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "BOOL": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteBit("", value.GetIndex(i).GetBool())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "BYTE" && numberOfValues == uint16(1): // BYTE
		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "BYTE": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint8("", 8, uint8(value.GetIndex(i).GetUint8()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "WORD" && numberOfValues == uint16(1): // WORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "WORD": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint16("", 16, uint16(value.GetIndex(i).GetUint16()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "DWORD" && numberOfValues == uint16(1): // DWORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "DWORD": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint32("", 32, uint32(value.GetIndex(i).GetUint32()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "LWORD" && numberOfValues == uint16(1): // LWORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "LWORD": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint64("", 64, uint64(value.GetIndex(i).GetUint64()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "SINT" && numberOfValues == uint16(1): // SINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt8("value", 8, int8(value.GetInt8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "SINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteInt8("", 8, int8(value.GetIndex(i).GetInt8()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "INT" && numberOfValues == uint16(1): // INT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt16("value", 16, int16(value.GetInt16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "INT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteInt16("", 16, int16(value.GetIndex(i).GetInt16()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "DINT" && numberOfValues == uint16(1): // DINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt32("value", 32, int32(value.GetInt32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "DINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteInt32("", 32, int32(value.GetIndex(i).GetInt32()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "LINT" && numberOfValues == uint16(1): // LINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt64("value", 64, int64(value.GetInt64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "LINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteInt64("", 64, int64(value.GetIndex(i).GetInt64()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "USINT" && numberOfValues == uint16(1): // USINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "USINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint8("", 8, uint8(value.GetIndex(i).GetUint8()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "UINT" && numberOfValues == uint16(1): // UINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "UINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint16("", 16, uint16(value.GetIndex(i).GetUint16()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "UDINT" && numberOfValues == uint16(1): // UDINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "UDINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint32("", 32, uint32(value.GetIndex(i).GetUint32()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "ULINT" && numberOfValues == uint16(1): // ULINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "ULINT": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteUint64("", 64, uint64(value.GetIndex(i).GetUint64()))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "REAL" && numberOfValues == uint16(1): // REAL
		// Simple Field (value)
		if _err := writeBuffer.WriteFloat32("value", 32, value.GetFloat32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "REAL": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteFloat32("", 32, value.GetIndex(i).GetFloat32())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "LREAL" && numberOfValues == uint16(1): // LREAL
		// Simple Field (value)
		if _err := writeBuffer.WriteFloat64("value", 64, value.GetFloat64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "LREAL": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteFloat64("", 64, value.GetIndex(i).GetFloat64())
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "CHAR" && numberOfValues == uint16(1): // CHAR
		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint32(8), value.GetString(), utils.WithEncoding("UTF-8)")); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "CHAR": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteString("", uint32(8), value.GetIndex(i).GetString(), utils.WithEncoding("UTF-8)"))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "WCHAR" && numberOfValues == uint16(1): // WCHAR
		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint32(16), value.GetString(), utils.WithEncoding("UTF-16)")); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "WCHAR": // List
		// Array Field (value)
		for i := uint32(0); i < uint32(m.NumberOfValues); i++ {
			_itemErr := writeBuffer.WriteString("", uint32(16), value.GetIndex(i).GetString(), utils.WithEncoding("UTF-16)"))
			if _itemErr != nil {
				return errors.Wrap(_itemErr, "Error serializing 'value' field")
			}
		}
	case dataType == "STRING": // STRING
		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint32(255), value.GetString(), utils.WithEncoding("UTF-8)")); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataType == "WSTRING": // STRING
		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint32(255), value.GetString(), utils.WithEncoding("UTF-16)")); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	default:
		// TODO: add more info which type it is actually
		return errors.New("unsupported type")
	}
	writeBuffer.PopContext("DataItem")
	return nil
}
