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
package org.apache.plc4x.java.cbus.readwrite;

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

public class NetworkProtocolControlInformation implements Message {

  // Properties.
  protected final byte stackCounter;
  protected final byte stackDepth;

  // Reserved Fields
  private Byte reservedField0;

  public NetworkProtocolControlInformation(byte stackCounter, byte stackDepth) {
    super();
    this.stackCounter = stackCounter;
    this.stackDepth = stackDepth;
  }

  public byte getStackCounter() {
    return stackCounter;
  }

  public byte getStackDepth() {
    return stackDepth;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NetworkProtocolControlInformation");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0x0,
        writeUnsignedByte(writeBuffer, 2));

    // Simple Field (stackCounter)
    writeSimpleField("stackCounter", stackCounter, writeUnsignedByte(writeBuffer, 3));

    // Simple Field (stackDepth)
    writeSimpleField("stackDepth", stackDepth, writeUnsignedByte(writeBuffer, 3));

    writeBuffer.popContext("NetworkProtocolControlInformation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    NetworkProtocolControlInformation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 2;

    // Simple field (stackCounter)
    lengthInBits += 3;

    // Simple field (stackDepth)
    lengthInBits += 3;

    return lengthInBits;
  }

  public static NetworkProtocolControlInformation staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("NetworkProtocolControlInformation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 2), (byte) 0x0);

    byte stackCounter = readSimpleField("stackCounter", readUnsignedByte(readBuffer, 3));

    byte stackDepth = readSimpleField("stackDepth", readUnsignedByte(readBuffer, 3));

    readBuffer.closeContext("NetworkProtocolControlInformation");
    // Create the instance
    NetworkProtocolControlInformation _networkProtocolControlInformation;
    _networkProtocolControlInformation =
        new NetworkProtocolControlInformation(stackCounter, stackDepth);
    _networkProtocolControlInformation.reservedField0 = reservedField0;
    return _networkProtocolControlInformation;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NetworkProtocolControlInformation)) {
      return false;
    }
    NetworkProtocolControlInformation that = (NetworkProtocolControlInformation) o;
    return (getStackCounter() == that.getStackCounter())
        && (getStackDepth() == that.getStackDepth())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getStackCounter(), getStackDepth());
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
