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
package org.apache.plc4x.java.opcua.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class MonitoredItemNotification extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "808";
  }

  // Properties.
  protected final long clientHandle;
  protected final DataValue value;

  public MonitoredItemNotification(long clientHandle, DataValue value) {
    super();
    this.clientHandle = clientHandle;
    this.value = value;
  }

  public long getClientHandle() {
    return clientHandle;
  }

  public DataValue getValue() {
    return value;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MonitoredItemNotification");

    // Simple Field (clientHandle)
    writeSimpleField("clientHandle", clientHandle, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (value)
    writeSimpleField("value", value, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("MonitoredItemNotification");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MonitoredItemNotification _value = this;

    // Simple field (clientHandle)
    lengthInBits += 32;

    // Simple field (value)
    lengthInBits += value.getLengthInBits();

    return lengthInBits;
  }

  public static MonitoredItemNotificationBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("MonitoredItemNotification");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    long clientHandle = readSimpleField("clientHandle", readUnsignedLong(readBuffer, 32));

    DataValue value =
        readSimpleField(
            "value",
            new DataReaderComplexDefault<>(() -> DataValue.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("MonitoredItemNotification");
    // Create the instance
    return new MonitoredItemNotificationBuilder(clientHandle, value);
  }

  public static class MonitoredItemNotificationBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long clientHandle;
    private final DataValue value;

    public MonitoredItemNotificationBuilder(long clientHandle, DataValue value) {

      this.clientHandle = clientHandle;
      this.value = value;
    }

    public MonitoredItemNotification build() {
      MonitoredItemNotification monitoredItemNotification =
          new MonitoredItemNotification(clientHandle, value);
      return monitoredItemNotification;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoredItemNotification)) {
      return false;
    }
    MonitoredItemNotification that = (MonitoredItemNotification) o;
    return (getClientHandle() == that.getClientHandle())
        && (getValue() == that.getValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getClientHandle(), getValue());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}