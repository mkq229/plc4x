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
package org.apache.plc4x.java.eip.readwrite;

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

public class NetworkConnectionParameters implements Message {

  // Properties.
  protected final int connectionSize;
  protected final boolean owner;
  protected final byte connectionType;
  protected final byte priority;
  protected final boolean connectionSizeType;

  public NetworkConnectionParameters(
      int connectionSize,
      boolean owner,
      byte connectionType,
      byte priority,
      boolean connectionSizeType) {
    super();
    this.connectionSize = connectionSize;
    this.owner = owner;
    this.connectionType = connectionType;
    this.priority = priority;
    this.connectionSizeType = connectionSizeType;
  }

  public int getConnectionSize() {
    return connectionSize;
  }

  public boolean getOwner() {
    return owner;
  }

  public byte getConnectionType() {
    return connectionType;
  }

  public byte getPriority() {
    return priority;
  }

  public boolean getConnectionSizeType() {
    return connectionSizeType;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NetworkConnectionParameters");

    // Simple Field (connectionSize)
    writeSimpleField("connectionSize", connectionSize, writeUnsignedInt(writeBuffer, 16));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (owner)
    writeSimpleField("owner", owner, writeBoolean(writeBuffer));

    // Simple Field (connectionType)
    writeSimpleField("connectionType", connectionType, writeUnsignedByte(writeBuffer, 2));

    // Reserved Field (reserved)
    writeReservedField("reserved", (boolean) false, writeBoolean(writeBuffer));

    // Simple Field (priority)
    writeSimpleField("priority", priority, writeUnsignedByte(writeBuffer, 2));

    // Simple Field (connectionSizeType)
    writeSimpleField("connectionSizeType", connectionSizeType, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (boolean) false, writeBoolean(writeBuffer));

    writeBuffer.popContext("NetworkConnectionParameters");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    NetworkConnectionParameters _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (connectionSize)
    lengthInBits += 16;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (owner)
    lengthInBits += 1;

    // Simple field (connectionType)
    lengthInBits += 2;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (priority)
    lengthInBits += 2;

    // Simple field (connectionSizeType)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static NetworkConnectionParameters staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("NetworkConnectionParameters");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int connectionSize = readSimpleField("connectionSize", readUnsignedInt(readBuffer, 16));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    boolean owner = readSimpleField("owner", readBoolean(readBuffer));

    byte connectionType = readSimpleField("connectionType", readUnsignedByte(readBuffer, 2));

    Boolean reservedField1 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    byte priority = readSimpleField("priority", readUnsignedByte(readBuffer, 2));

    boolean connectionSizeType = readSimpleField("connectionSizeType", readBoolean(readBuffer));

    Boolean reservedField2 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    readBuffer.closeContext("NetworkConnectionParameters");
    // Create the instance
    NetworkConnectionParameters _networkConnectionParameters;
    _networkConnectionParameters =
        new NetworkConnectionParameters(
            connectionSize, owner, connectionType, priority, connectionSizeType);
    return _networkConnectionParameters;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NetworkConnectionParameters)) {
      return false;
    }
    NetworkConnectionParameters that = (NetworkConnectionParameters) o;
    return (getConnectionSize() == that.getConnectionSize())
        && (getOwner() == that.getOwner())
        && (getConnectionType() == that.getConnectionType())
        && (getPriority() == that.getPriority())
        && (getConnectionSizeType() == that.getConnectionSizeType())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getConnectionSize(),
        getOwner(),
        getConnectionType(),
        getPriority(),
        getConnectionSizeType());
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
