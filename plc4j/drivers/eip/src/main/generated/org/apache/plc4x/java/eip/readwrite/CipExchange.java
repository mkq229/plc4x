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

public class CipExchange implements Message {

  // Constant values.
  public static final Integer ITEMCOUNT = 0x02;
  public static final Long NULLPTR = 0x0L;
  public static final Integer UNCONNECTEDDATA = 0x00B2;

  // Properties.
  protected final CipService service;

  public CipExchange(CipService service) {
    super();
    this.service = service;
  }

  public CipService getService() {
    return service;
  }

  public int getItemCount() {
    return ITEMCOUNT;
  }

  public long getNullPtr() {
    return NULLPTR;
  }

  public int getUnconnectedData() {
    return UNCONNECTEDDATA;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CipExchange");

    // Const Field (itemCount)
    writeConstField("itemCount", ITEMCOUNT, writeUnsignedInt(writeBuffer, 16));

    // Const Field (nullPtr)
    writeConstField("nullPtr", NULLPTR, writeUnsignedLong(writeBuffer, 32));

    // Const Field (unconnectedData)
    writeConstField("unconnectedData", UNCONNECTEDDATA, writeUnsignedInt(writeBuffer, 16));

    // Implicit Field (size) (Used for parsing, but its value is not stored as it's implicitly given
    // by the objects content)
    int size = (int) (((getLengthInBytes()) - (8)) - (2));
    writeImplicitField("size", size, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (service)
    writeSimpleField("service", service, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CipExchange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CipExchange _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (itemCount)
    lengthInBits += 16;

    // Const Field (nullPtr)
    lengthInBits += 32;

    // Const Field (unconnectedData)
    lengthInBits += 16;

    // Implicit Field (size)
    lengthInBits += 16;

    // Simple field (service)
    lengthInBits += service.getLengthInBits();

    return lengthInBits;
  }

  public static CipExchange staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Integer exchangeLen;
    if (args[0] instanceof Integer) {
      exchangeLen = (Integer) args[0];
    } else if (args[0] instanceof String) {
      exchangeLen = Integer.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Integer or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, exchangeLen);
  }

  public static CipExchange staticParse(ReadBuffer readBuffer, Integer exchangeLen)
      throws ParseException {
    readBuffer.pullContext("CipExchange");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int itemCount =
        readConstField("itemCount", readUnsignedInt(readBuffer, 16), CipExchange.ITEMCOUNT);

    long nullPtr = readConstField("nullPtr", readUnsignedLong(readBuffer, 32), CipExchange.NULLPTR);

    int unconnectedData =
        readConstField(
            "unconnectedData", readUnsignedInt(readBuffer, 16), CipExchange.UNCONNECTEDDATA);

    int size = readImplicitField("size", readUnsignedInt(readBuffer, 16));

    CipService service =
        readSimpleField(
            "service",
            new DataReaderComplexDefault<>(
                () -> CipService.staticParse(readBuffer, (int) ((exchangeLen) - (10))),
                readBuffer));

    readBuffer.closeContext("CipExchange");
    // Create the instance
    CipExchange _cipExchange;
    _cipExchange = new CipExchange(service);
    return _cipExchange;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipExchange)) {
      return false;
    }
    CipExchange that = (CipExchange) o;
    return (getService() == that.getService()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getService());
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