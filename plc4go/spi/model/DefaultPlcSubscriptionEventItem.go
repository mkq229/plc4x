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
	"encoding/binary"
	"fmt"
	"time"

	"github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

type DefaultPlcSubscriptionEventItem struct {
	code             model.PlcResponseCode
	tag              model.PlcTag
	subscriptionType SubscriptionType
	interval         time.Duration
	value            values.PlcValue
}

func NewSubscriptionEventItem(code model.PlcResponseCode, tag model.PlcTag, subscriptionType SubscriptionType, interval time.Duration, value values.PlcValue) *DefaultPlcSubscriptionEventItem {
	return &DefaultPlcSubscriptionEventItem{
		code:             code,
		tag:              tag,
		subscriptionType: subscriptionType,
		interval:         interval,
		value:            value,
	}
}

func (r *DefaultPlcSubscriptionEventItem) GetCode() model.PlcResponseCode {
	return r.code
}

func (r *DefaultPlcSubscriptionEventItem) GetTag() model.PlcTag {
	return r.tag
}

func (r *DefaultPlcSubscriptionEventItem) GetSubscriptionType() SubscriptionType {
	return r.subscriptionType
}

func (r *DefaultPlcSubscriptionEventItem) GetInterval() time.Duration {
	return r.interval
}

func (r *DefaultPlcSubscriptionEventItem) GetValue() values.PlcValue {
	return r.value
}

func (d *DefaultPlcSubscriptionEventItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcSubscriptionEventItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcSubscriptionEventItem"); err != nil {
		return err
	}

	{
		stringValue := fmt.Sprintf("%v", d.code)
		if err := writeBuffer.WriteString("code", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
			return err
		}
	}

	if d.tag != nil {
		if serializableField, ok := d.tag.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("tag"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("tag"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.tag)
			if err := writeBuffer.WriteString("tag", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
	_value := fmt.Sprintf("%v", d.subscriptionType)

	if err := writeBuffer.WriteString("subscriptionType", uint32(len(_value)*8), "UTF-8", _value); err != nil {
		return err
	}

	{
		stringValue := fmt.Sprintf("%v", d.interval)
		if err := writeBuffer.WriteString("interval", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
			return err
		}
	}

	if d.value != nil {
		if serializableField, ok := d.value.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("value"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("value"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.value)
			if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("PlcSubscriptionEventItem"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcSubscriptionEventItem) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}