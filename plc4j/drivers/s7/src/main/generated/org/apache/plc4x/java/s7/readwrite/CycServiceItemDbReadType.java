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
package org.apache.plc4x.java.s7.readwrite;

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

public class CycServiceItemDbReadType extends CycServiceItemType implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final short numberOfAreas;
  protected final List<SubItem> items;

  public CycServiceItemDbReadType(
      short byteLength, short syntaxId, short numberOfAreas, List<SubItem> items) {
    super(byteLength, syntaxId);
    this.numberOfAreas = numberOfAreas;
    this.items = items;
  }

  public short getNumberOfAreas() {
    return numberOfAreas;
  }

  public List<SubItem> getItems() {
    return items;
  }

  @Override
  protected void serializeCycServiceItemTypeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CycServiceItemDbReadType");

    // Simple Field (numberOfAreas)
    writeSimpleField("numberOfAreas", numberOfAreas, writeUnsignedShort(writeBuffer, 8));

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("CycServiceItemDbReadType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CycServiceItemDbReadType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (numberOfAreas)
    lengthInBits += 8;

    // Array field
    if (items != null) {
      int i = 0;
      for (SubItem element : items) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= items.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static CycServiceItemTypeBuilder staticParseCycServiceItemTypeBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CycServiceItemDbReadType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short numberOfAreas = readSimpleField("numberOfAreas", readUnsignedShort(readBuffer, 8));

    List<SubItem> items =
        readCountArrayField(
            "items", readComplex(() -> SubItem.staticParse(readBuffer), readBuffer), numberOfAreas);

    readBuffer.closeContext("CycServiceItemDbReadType");
    // Create the instance
    return new CycServiceItemDbReadTypeBuilderImpl(numberOfAreas, items);
  }

  public static class CycServiceItemDbReadTypeBuilderImpl
      implements CycServiceItemType.CycServiceItemTypeBuilder {
    private final short numberOfAreas;
    private final List<SubItem> items;

    public CycServiceItemDbReadTypeBuilderImpl(short numberOfAreas, List<SubItem> items) {
      this.numberOfAreas = numberOfAreas;
      this.items = items;
    }

    public CycServiceItemDbReadType build(short byteLength, short syntaxId) {
      CycServiceItemDbReadType cycServiceItemDbReadType =
          new CycServiceItemDbReadType(byteLength, syntaxId, numberOfAreas, items);
      return cycServiceItemDbReadType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CycServiceItemDbReadType)) {
      return false;
    }
    CycServiceItemDbReadType that = (CycServiceItemDbReadType) o;
    return (getNumberOfAreas() == that.getNumberOfAreas())
        && (getItems() == that.getItems())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNumberOfAreas(), getItems());
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
