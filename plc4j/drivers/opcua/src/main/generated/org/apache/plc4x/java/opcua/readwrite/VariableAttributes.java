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

public class VariableAttributes extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 357;
  }

  // Properties.
  protected final long specifiedAttributes;
  protected final LocalizedText displayName;
  protected final LocalizedText description;
  protected final long writeMask;
  protected final long userWriteMask;
  protected final Variant value;
  protected final NodeId dataType;
  protected final int valueRank;
  protected final List<Long> arrayDimensions;
  protected final short accessLevel;
  protected final short userAccessLevel;
  protected final double minimumSamplingInterval;
  protected final boolean historizing;

  public VariableAttributes(
      long specifiedAttributes,
      LocalizedText displayName,
      LocalizedText description,
      long writeMask,
      long userWriteMask,
      Variant value,
      NodeId dataType,
      int valueRank,
      List<Long> arrayDimensions,
      short accessLevel,
      short userAccessLevel,
      double minimumSamplingInterval,
      boolean historizing) {
    super();
    this.specifiedAttributes = specifiedAttributes;
    this.displayName = displayName;
    this.description = description;
    this.writeMask = writeMask;
    this.userWriteMask = userWriteMask;
    this.value = value;
    this.dataType = dataType;
    this.valueRank = valueRank;
    this.arrayDimensions = arrayDimensions;
    this.accessLevel = accessLevel;
    this.userAccessLevel = userAccessLevel;
    this.minimumSamplingInterval = minimumSamplingInterval;
    this.historizing = historizing;
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

  public Variant getValue() {
    return value;
  }

  public NodeId getDataType() {
    return dataType;
  }

  public int getValueRank() {
    return valueRank;
  }

  public List<Long> getArrayDimensions() {
    return arrayDimensions;
  }

  public short getAccessLevel() {
    return accessLevel;
  }

  public short getUserAccessLevel() {
    return userAccessLevel;
  }

  public double getMinimumSamplingInterval() {
    return minimumSamplingInterval;
  }

  public boolean getHistorizing() {
    return historizing;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("VariableAttributes");

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

    // Simple Field (value)
    writeSimpleField("value", value, writeComplex(writeBuffer));

    // Simple Field (dataType)
    writeSimpleField("dataType", dataType, writeComplex(writeBuffer));

    // Simple Field (valueRank)
    writeSimpleField("valueRank", valueRank, writeSignedInt(writeBuffer, 32));

    // Implicit Field (noOfArrayDimensions) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfArrayDimensions =
        (int) ((((getArrayDimensions()) == (null)) ? -(1) : COUNT(getArrayDimensions())));
    writeImplicitField("noOfArrayDimensions", noOfArrayDimensions, writeSignedInt(writeBuffer, 32));

    // Array Field (arrayDimensions)
    writeSimpleTypeArrayField(
        "arrayDimensions", arrayDimensions, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (accessLevel)
    writeSimpleField("accessLevel", accessLevel, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (userAccessLevel)
    writeSimpleField("userAccessLevel", userAccessLevel, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (minimumSamplingInterval)
    writeSimpleField(
        "minimumSamplingInterval", minimumSamplingInterval, writeDouble(writeBuffer, 64));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (historizing)
    writeSimpleField("historizing", historizing, writeBoolean(writeBuffer));

    writeBuffer.popContext("VariableAttributes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    VariableAttributes _value = this;
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

    // Simple field (value)
    lengthInBits += value.getLengthInBits();

    // Simple field (dataType)
    lengthInBits += dataType.getLengthInBits();

    // Simple field (valueRank)
    lengthInBits += 32;

    // Implicit Field (noOfArrayDimensions)
    lengthInBits += 32;

    // Array field
    if (arrayDimensions != null) {
      lengthInBits += 32 * arrayDimensions.size();
    }

    // Simple field (accessLevel)
    lengthInBits += 8;

    // Simple field (userAccessLevel)
    lengthInBits += 8;

    // Simple field (minimumSamplingInterval)
    lengthInBits += 64;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (historizing)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("VariableAttributes");
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

    Variant value =
        readSimpleField("value", readComplex(() -> Variant.staticParse(readBuffer), readBuffer));

    NodeId dataType =
        readSimpleField("dataType", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    int valueRank = readSimpleField("valueRank", readSignedInt(readBuffer, 32));

    int noOfArrayDimensions =
        readImplicitField("noOfArrayDimensions", readSignedInt(readBuffer, 32));

    List<Long> arrayDimensions =
        readCountArrayField(
            "arrayDimensions", readUnsignedLong(readBuffer, 32), noOfArrayDimensions);

    short accessLevel = readSimpleField("accessLevel", readUnsignedShort(readBuffer, 8));

    short userAccessLevel = readSimpleField("userAccessLevel", readUnsignedShort(readBuffer, 8));

    double minimumSamplingInterval =
        readSimpleField("minimumSamplingInterval", readDouble(readBuffer, 64));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean historizing = readSimpleField("historizing", readBoolean(readBuffer));

    readBuffer.closeContext("VariableAttributes");
    // Create the instance
    return new VariableAttributesBuilderImpl(
        specifiedAttributes,
        displayName,
        description,
        writeMask,
        userWriteMask,
        value,
        dataType,
        valueRank,
        arrayDimensions,
        accessLevel,
        userAccessLevel,
        minimumSamplingInterval,
        historizing);
  }

  public static class VariableAttributesBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long specifiedAttributes;
    private final LocalizedText displayName;
    private final LocalizedText description;
    private final long writeMask;
    private final long userWriteMask;
    private final Variant value;
    private final NodeId dataType;
    private final int valueRank;
    private final List<Long> arrayDimensions;
    private final short accessLevel;
    private final short userAccessLevel;
    private final double minimumSamplingInterval;
    private final boolean historizing;

    public VariableAttributesBuilderImpl(
        long specifiedAttributes,
        LocalizedText displayName,
        LocalizedText description,
        long writeMask,
        long userWriteMask,
        Variant value,
        NodeId dataType,
        int valueRank,
        List<Long> arrayDimensions,
        short accessLevel,
        short userAccessLevel,
        double minimumSamplingInterval,
        boolean historizing) {
      this.specifiedAttributes = specifiedAttributes;
      this.displayName = displayName;
      this.description = description;
      this.writeMask = writeMask;
      this.userWriteMask = userWriteMask;
      this.value = value;
      this.dataType = dataType;
      this.valueRank = valueRank;
      this.arrayDimensions = arrayDimensions;
      this.accessLevel = accessLevel;
      this.userAccessLevel = userAccessLevel;
      this.minimumSamplingInterval = minimumSamplingInterval;
      this.historizing = historizing;
    }

    public VariableAttributes build() {
      VariableAttributes variableAttributes =
          new VariableAttributes(
              specifiedAttributes,
              displayName,
              description,
              writeMask,
              userWriteMask,
              value,
              dataType,
              valueRank,
              arrayDimensions,
              accessLevel,
              userAccessLevel,
              minimumSamplingInterval,
              historizing);
      return variableAttributes;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof VariableAttributes)) {
      return false;
    }
    VariableAttributes that = (VariableAttributes) o;
    return (getSpecifiedAttributes() == that.getSpecifiedAttributes())
        && (getDisplayName() == that.getDisplayName())
        && (getDescription() == that.getDescription())
        && (getWriteMask() == that.getWriteMask())
        && (getUserWriteMask() == that.getUserWriteMask())
        && (getValue() == that.getValue())
        && (getDataType() == that.getDataType())
        && (getValueRank() == that.getValueRank())
        && (getArrayDimensions() == that.getArrayDimensions())
        && (getAccessLevel() == that.getAccessLevel())
        && (getUserAccessLevel() == that.getUserAccessLevel())
        && (getMinimumSamplingInterval() == that.getMinimumSamplingInterval())
        && (getHistorizing() == that.getHistorizing())
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
        getValue(),
        getDataType(),
        getValueRank(),
        getArrayDimensions(),
        getAccessLevel(),
        getUserAccessLevel(),
        getMinimumSamplingInterval(),
        getHistorizing());
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