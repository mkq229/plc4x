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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetConfirmedServiceRequestCreateObjectObjectSpecifier implements Message {

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagEnumerated rawObjectType;
  protected final BACnetContextTagObjectIdentifier objectIdentifier;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;

  public BACnetConfirmedServiceRequestCreateObjectObjectSpecifier(
      BACnetOpeningTag openingTag,
      BACnetContextTagEnumerated rawObjectType,
      BACnetContextTagObjectIdentifier objectIdentifier,
      BACnetClosingTag closingTag,
      Short tagNumber) {
    super();
    this.openingTag = openingTag;
    this.rawObjectType = rawObjectType;
    this.objectIdentifier = objectIdentifier;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagEnumerated getRawObjectType() {
    return rawObjectType;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public boolean getIsObjectType() {
    return (boolean) ((getRawObjectType()) != (null));
  }

  public BACnetObjectType getObjectType() {
    return (BACnetObjectType)
        (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.mapBACnetObjectType(
            getRawObjectType()));
  }

  public boolean getIsObjectIdentifier() {
    return (boolean) ((getObjectIdentifier()) != (null));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (rawObjectType) (Can be skipped, if the value is null)
    writeOptionalField("rawObjectType", rawObjectType, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isObjectType = getIsObjectType();
    writeBuffer.writeVirtual("isObjectType", isObjectType);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetObjectType objectType = getObjectType();
    writeBuffer.writeVirtual("objectType", objectType);

    // Optional Field (objectIdentifier) (Can be skipped, if the value is null)
    writeOptionalField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isObjectIdentifier = getIsObjectIdentifier();
    writeBuffer.writeVirtual("isObjectIdentifier", isObjectIdentifier);

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetConfirmedServiceRequestCreateObjectObjectSpecifier _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Optional Field (rawObjectType)
    if (rawObjectType != null) {
      lengthInBits += rawObjectType.getLengthInBits();
    }

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Optional Field (objectIdentifier)
    if (objectIdentifier != null) {
      lengthInBits += objectIdentifier.getLengthInBits();
    }

    // A virtual field doesn't have any in- or output.

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestCreateObjectObjectSpecifier staticParse(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    BACnetContextTagEnumerated rawObjectType =
        readOptionalField(
            "rawObjectType",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagEnumerated)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (0), (BACnetDataType) (BACnetDataType.ENUMERATED)),
                readBuffer));
    boolean isObjectType =
        readVirtualField("isObjectType", boolean.class, (rawObjectType) != (null));
    BACnetObjectType objectType =
        readVirtualField(
            "objectType",
            BACnetObjectType.class,
            org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.mapBACnetObjectType(
                rawObjectType));

    BACnetContextTagObjectIdentifier objectIdentifier =
        readOptionalField(
            "objectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));
    boolean isObjectIdentifier =
        readVirtualField("isObjectIdentifier", boolean.class, (objectIdentifier) != (null));
    // Validation
    if (!((isObjectType) || (isObjectIdentifier))) {
      throw new ParseValidationException("either we need a objectType or a objectIdentifier");
    }

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier");
    // Create the instance
    BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
        _bACnetConfirmedServiceRequestCreateObjectObjectSpecifier;
    _bACnetConfirmedServiceRequestCreateObjectObjectSpecifier =
        new BACnetConfirmedServiceRequestCreateObjectObjectSpecifier(
            openingTag, rawObjectType, objectIdentifier, closingTag, tagNumber);
    return _bACnetConfirmedServiceRequestCreateObjectObjectSpecifier;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestCreateObjectObjectSpecifier)) {
      return false;
    }
    BACnetConfirmedServiceRequestCreateObjectObjectSpecifier that =
        (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getRawObjectType() == that.getRawObjectType())
        && (getObjectIdentifier() == that.getObjectIdentifier())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getOpeningTag(), getRawObjectType(), getObjectIdentifier(), getClosingTag());
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
