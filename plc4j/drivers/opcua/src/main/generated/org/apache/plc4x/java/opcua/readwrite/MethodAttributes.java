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

public class MethodAttributes extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 360;
  }

  // Properties.
  protected final long specifiedAttributes;
  protected final LocalizedText displayName;
  protected final LocalizedText description;
  protected final long writeMask;
  protected final long userWriteMask;
  protected final boolean userExecutable;
  protected final boolean executable;

  public MethodAttributes(
      long specifiedAttributes,
      LocalizedText displayName,
      LocalizedText description,
      long writeMask,
      long userWriteMask,
      boolean userExecutable,
      boolean executable) {
    super();
    this.specifiedAttributes = specifiedAttributes;
    this.displayName = displayName;
    this.description = description;
    this.writeMask = writeMask;
    this.userWriteMask = userWriteMask;
    this.userExecutable = userExecutable;
    this.executable = executable;
  }

  public long getSpecifiedAttributes() {
    return specifiedAttributes;
  }

  public LocalizedText getDisplayName() {
    return displayName;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public long getWriteMask() {
    return writeMask;
  }

  public long getUserWriteMask() {
    return userWriteMask;
  }

  public boolean getUserExecutable() {
    return userExecutable;
  }

  public boolean getExecutable() {
    return executable;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MethodAttributes");

    // Simple Field (specifiedAttributes)
    writeSimpleField(
        "specifiedAttributes", specifiedAttributes, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (displayName)
    writeSimpleField("displayName", displayName, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (writeMask)
    writeSimpleField("writeMask", writeMask, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (userWriteMask)
    writeSimpleField("userWriteMask", userWriteMask, writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 6));

    // Simple Field (userExecutable)
    writeSimpleField("userExecutable", userExecutable, writeBoolean(writeBuffer));

    // Simple Field (executable)
    writeSimpleField("executable", executable, writeBoolean(writeBuffer));

    writeBuffer.popContext("MethodAttributes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MethodAttributes _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (specifiedAttributes)
    lengthInBits += 32;

    // Simple field (displayName)
    lengthInBits += displayName.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (writeMask)
    lengthInBits += 32;

    // Simple field (userWriteMask)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 6;

    // Simple field (userExecutable)
    lengthInBits += 1;

    // Simple field (executable)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("MethodAttributes");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long specifiedAttributes =
        readSimpleField("specifiedAttributes", readUnsignedLong(readBuffer, 32));

    LocalizedText displayName =
        readSimpleField(
            "displayName", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    long writeMask = readSimpleField("writeMask", readUnsignedLong(readBuffer, 32));

    long userWriteMask = readSimpleField("userWriteMask", readUnsignedLong(readBuffer, 32));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 6), (byte) 0x00);

    boolean userExecutable = readSimpleField("userExecutable", readBoolean(readBuffer));

    boolean executable = readSimpleField("executable", readBoolean(readBuffer));

    readBuffer.closeContext("MethodAttributes");
    // Create the instance
    return new MethodAttributesBuilderImpl(
        specifiedAttributes,
        displayName,
        description,
        writeMask,
        userWriteMask,
        userExecutable,
        executable);
  }

  public static class MethodAttributesBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long specifiedAttributes;
    private final LocalizedText displayName;
    private final LocalizedText description;
    private final long writeMask;
    private final long userWriteMask;
    private final boolean userExecutable;
    private final boolean executable;

    public MethodAttributesBuilderImpl(
        long specifiedAttributes,
        LocalizedText displayName,
        LocalizedText description,
        long writeMask,
        long userWriteMask,
        boolean userExecutable,
        boolean executable) {
      this.specifiedAttributes = specifiedAttributes;
      this.displayName = displayName;
      this.description = description;
      this.writeMask = writeMask;
      this.userWriteMask = userWriteMask;
      this.userExecutable = userExecutable;
      this.executable = executable;
    }

    public MethodAttributes build() {
      MethodAttributes methodAttributes =
          new MethodAttributes(
              specifiedAttributes,
              displayName,
              description,
              writeMask,
              userWriteMask,
              userExecutable,
              executable);
      return methodAttributes;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MethodAttributes)) {
      return false;
    }
    MethodAttributes that = (MethodAttributes) o;
    return (getSpecifiedAttributes() == that.getSpecifiedAttributes())
        && (getDisplayName() == that.getDisplayName())
        && (getDescription() == that.getDescription())
        && (getWriteMask() == that.getWriteMask())
        && (getUserWriteMask() == that.getUserWriteMask())
        && (getUserExecutable() == that.getUserExecutable())
        && (getExecutable() == that.getExecutable())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSpecifiedAttributes(),
        getDisplayName(),
        getDescription(),
        getWriteMask(),
        getUserWriteMask(),
        getUserExecutable(),
        getExecutable());
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