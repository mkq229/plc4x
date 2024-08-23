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

public class BACnetPropertyAccessResult implements Message {

  // Properties.
  protected final BACnetContextTagObjectIdentifier objectIdentifier;
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger propertyArrayIndex;
  protected final BACnetContextTagObjectIdentifier deviceIdentifier;
  protected final BACnetPropertyAccessResultAccessResult accessResult;

  public BACnetPropertyAccessResult(
      BACnetContextTagObjectIdentifier objectIdentifier,
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger propertyArrayIndex,
      BACnetContextTagObjectIdentifier deviceIdentifier,
      BACnetPropertyAccessResultAccessResult accessResult) {
    super();
    this.objectIdentifier = objectIdentifier;
    this.propertyIdentifier = propertyIdentifier;
    this.propertyArrayIndex = propertyArrayIndex;
    this.deviceIdentifier = deviceIdentifier;
    this.accessResult = accessResult;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getPropertyArrayIndex() {
    return propertyArrayIndex;
  }

  public BACnetContextTagObjectIdentifier getDeviceIdentifier() {
    return deviceIdentifier;
  }

  public BACnetPropertyAccessResultAccessResult getAccessResult() {
    return accessResult;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPropertyAccessResult");

    // Simple Field (objectIdentifier)
    writeSimpleField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (propertyIdentifier)
    writeSimpleField(
        "propertyIdentifier", propertyIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
    writeOptionalField(
        "propertyArrayIndex", propertyArrayIndex, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
    writeOptionalField(
        "deviceIdentifier", deviceIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (accessResult)
    writeSimpleField("accessResult", accessResult, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyAccessResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetPropertyAccessResult _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (propertyArrayIndex)
    if (propertyArrayIndex != null) {
      lengthInBits += propertyArrayIndex.getLengthInBits();
    }

    // Optional Field (deviceIdentifier)
    if (deviceIdentifier != null) {
      lengthInBits += deviceIdentifier.getLengthInBits();
    }

    // Simple field (accessResult)
    lengthInBits += accessResult.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyAccessResult staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetPropertyAccessResult");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetPropertyIdentifierTagged propertyIdentifier =
        readSimpleField(
            "propertyIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPropertyIdentifierTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger propertyArrayIndex =
        readOptionalField(
            "propertyArrayIndex",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagObjectIdentifier deviceIdentifier =
        readOptionalField(
            "deviceIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetPropertyAccessResultAccessResult accessResult =
        readSimpleField(
            "accessResult",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPropertyAccessResultAccessResult.staticParse(
                        readBuffer,
                        (BACnetObjectType) (objectIdentifier.getObjectType()),
                        (BACnetPropertyIdentifier) (propertyIdentifier.getValue()),
                        (BACnetTagPayloadUnsignedInteger)
                            (((((propertyArrayIndex) != (null))
                                ? propertyArrayIndex.getPayload()
                                : null)))),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyAccessResult");
    // Create the instance
    BACnetPropertyAccessResult _bACnetPropertyAccessResult;
    _bACnetPropertyAccessResult =
        new BACnetPropertyAccessResult(
            objectIdentifier,
            propertyIdentifier,
            propertyArrayIndex,
            deviceIdentifier,
            accessResult);
    return _bACnetPropertyAccessResult;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyAccessResult)) {
      return false;
    }
    BACnetPropertyAccessResult that = (BACnetPropertyAccessResult) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getPropertyArrayIndex() == that.getPropertyArrayIndex())
        && (getDeviceIdentifier() == that.getDeviceIdentifier())
        && (getAccessResult() == that.getAccessResult())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getObjectIdentifier(),
        getPropertyIdentifier(),
        getPropertyArrayIndex(),
        getDeviceIdentifier(),
        getAccessResult());
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
