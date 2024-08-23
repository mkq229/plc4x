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
package org.apache.plc4x.java.canopen.readwrite;

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

public class CANOpenTime implements Message {

  // Properties.
  protected final long millis;
  protected final int days;

  public CANOpenTime(long millis, int days) {
    super();
    this.millis = millis;
    this.days = days;
  }

  public long getMillis() {
    return millis;
  }

  public int getDays() {
    return days;
  }

  public int getCleanMillis() {
    return (int) ((getMillis()) & (0x0FFFFFFF));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CANOpenTime");

    // Simple Field (millis)
    writeSimpleField("millis", millis, writeUnsignedLong(writeBuffer, 32));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    int cleanMillis = getCleanMillis();
    writeBuffer.writeVirtual("cleanMillis", cleanMillis);

    // Simple Field (days)
    writeSimpleField("days", days, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("CANOpenTime");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CANOpenTime _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (millis)
    lengthInBits += 32;

    // A virtual field doesn't have any in- or output.

    // Simple field (days)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static CANOpenTime staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CANOpenTime");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long millis = readSimpleField("millis", readUnsignedLong(readBuffer, 32));
    int cleanMillis = readVirtualField("cleanMillis", int.class, (millis) & (0x0FFFFFFF));

    int days = readSimpleField("days", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("CANOpenTime");
    // Create the instance
    CANOpenTime _cANOpenTime;
    _cANOpenTime = new CANOpenTime(millis, days);
    return _cANOpenTime;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CANOpenTime)) {
      return false;
    }
    CANOpenTime that = (CANOpenTime) o;
    return (getMillis() == that.getMillis()) && (getDays() == that.getDays()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getMillis(), getDays());
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
