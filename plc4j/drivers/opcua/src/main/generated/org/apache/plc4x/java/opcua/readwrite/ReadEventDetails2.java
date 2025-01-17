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

public class ReadEventDetails2 extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 32801;
  }

  // Properties.
  protected final long numValuesPerNode;
  protected final long startTime;
  protected final long endTime;
  protected final EventFilter filter;
  protected final boolean readModified;

  public ReadEventDetails2(
      long numValuesPerNode,
      long startTime,
      long endTime,
      EventFilter filter,
      boolean readModified) {
    super();
    this.numValuesPerNode = numValuesPerNode;
    this.startTime = startTime;
    this.endTime = endTime;
    this.filter = filter;
    this.readModified = readModified;
  }

  public long getNumValuesPerNode() {
    return numValuesPerNode;
  }

  public long getStartTime() {
    return startTime;
  }

  public long getEndTime() {
    return endTime;
  }

  public EventFilter getFilter() {
    return filter;
  }

  public boolean getReadModified() {
    return readModified;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ReadEventDetails2");

    // Simple Field (numValuesPerNode)
    writeSimpleField("numValuesPerNode", numValuesPerNode, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (endTime)
    writeSimpleField("endTime", endTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (filter)
    writeSimpleField("filter", filter, writeComplex(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (readModified)
    writeSimpleField("readModified", readModified, writeBoolean(writeBuffer));

    writeBuffer.popContext("ReadEventDetails2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ReadEventDetails2 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (numValuesPerNode)
    lengthInBits += 32;

    // Simple field (startTime)
    lengthInBits += 64;

    // Simple field (endTime)
    lengthInBits += 64;

    // Simple field (filter)
    lengthInBits += filter.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (readModified)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("ReadEventDetails2");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long numValuesPerNode = readSimpleField("numValuesPerNode", readUnsignedLong(readBuffer, 32));

    long startTime = readSimpleField("startTime", readSignedLong(readBuffer, 64));

    long endTime = readSimpleField("endTime", readSignedLong(readBuffer, 64));

    EventFilter filter =
        readSimpleField(
            "filter",
            readComplex(
                () -> (EventFilter) ExtensionObjectDefinition.staticParse(readBuffer, (int) (727)),
                readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean readModified = readSimpleField("readModified", readBoolean(readBuffer));

    readBuffer.closeContext("ReadEventDetails2");
    // Create the instance
    return new ReadEventDetails2BuilderImpl(
        numValuesPerNode, startTime, endTime, filter, readModified);
  }

  public static class ReadEventDetails2BuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long numValuesPerNode;
    private final long startTime;
    private final long endTime;
    private final EventFilter filter;
    private final boolean readModified;

    public ReadEventDetails2BuilderImpl(
        long numValuesPerNode,
        long startTime,
        long endTime,
        EventFilter filter,
        boolean readModified) {
      this.numValuesPerNode = numValuesPerNode;
      this.startTime = startTime;
      this.endTime = endTime;
      this.filter = filter;
      this.readModified = readModified;
    }

    public ReadEventDetails2 build() {
      ReadEventDetails2 readEventDetails2 =
          new ReadEventDetails2(numValuesPerNode, startTime, endTime, filter, readModified);
      return readEventDetails2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ReadEventDetails2)) {
      return false;
    }
    ReadEventDetails2 that = (ReadEventDetails2) o;
    return (getNumValuesPerNode() == that.getNumValuesPerNode())
        && (getStartTime() == that.getStartTime())
        && (getEndTime() == that.getEndTime())
        && (getFilter() == that.getFilter())
        && (getReadModified() == that.getReadModified())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNumValuesPerNode(),
        getStartTime(),
        getEndTime(),
        getFilter(),
        getReadModified());
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
