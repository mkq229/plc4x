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

public class SubscribeCOVPropertyMultipleErrorFirstFailedSubscription implements Message {

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagObjectIdentifier monitoredObjectIdentifier;
  protected final BACnetPropertyReferenceEnclosed monitoredPropertyReference;
  protected final ErrorEnclosed errorType;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;

  public SubscribeCOVPropertyMultipleErrorFirstFailedSubscription(
      BACnetOpeningTag openingTag,
      BACnetContextTagObjectIdentifier monitoredObjectIdentifier,
      BACnetPropertyReferenceEnclosed monitoredPropertyReference,
      ErrorEnclosed errorType,
      BACnetClosingTag closingTag,
      Short tagNumber) {
    super();
    this.openingTag = openingTag;
    this.monitoredObjectIdentifier = monitoredObjectIdentifier;
    this.monitoredPropertyReference = monitoredPropertyReference;
    this.errorType = errorType;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagObjectIdentifier getMonitoredObjectIdentifier() {
    return monitoredObjectIdentifier;
  }

  public BACnetPropertyReferenceEnclosed getMonitoredPropertyReference() {
    return monitoredPropertyReference;
  }

  public ErrorEnclosed getErrorType() {
    return errorType;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (monitoredObjectIdentifier)
    writeSimpleField(
        "monitoredObjectIdentifier",
        monitoredObjectIdentifier,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (monitoredPropertyReference)
    writeSimpleField(
        "monitoredPropertyReference",
        monitoredPropertyReference,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (errorType)
    writeSimpleField("errorType", errorType, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SubscribeCOVPropertyMultipleErrorFirstFailedSubscription _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (monitoredObjectIdentifier)
    lengthInBits += monitoredObjectIdentifier.getLengthInBits();

    // Simple field (monitoredPropertyReference)
    lengthInBits += monitoredPropertyReference.getLengthInBits();

    // Simple field (errorType)
    lengthInBits += errorType.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static SubscribeCOVPropertyMultipleErrorFirstFailedSubscription staticParse(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    BACnetContextTagObjectIdentifier monitoredObjectIdentifier =
        readSimpleField(
            "monitoredObjectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetPropertyReferenceEnclosed monitoredPropertyReference =
        readSimpleField(
            "monitoredPropertyReference",
            new DataReaderComplexDefault<>(
                () -> BACnetPropertyReferenceEnclosed.staticParse(readBuffer, (short) (1)),
                readBuffer));

    ErrorEnclosed errorType =
        readSimpleField(
            "errorType",
            new DataReaderComplexDefault<>(
                () -> ErrorEnclosed.staticParse(readBuffer, (short) (2)), readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription");
    // Create the instance
    SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
        _subscribeCOVPropertyMultipleErrorFirstFailedSubscription;
    _subscribeCOVPropertyMultipleErrorFirstFailedSubscription =
        new SubscribeCOVPropertyMultipleErrorFirstFailedSubscription(
            openingTag,
            monitoredObjectIdentifier,
            monitoredPropertyReference,
            errorType,
            closingTag,
            tagNumber);
    return _subscribeCOVPropertyMultipleErrorFirstFailedSubscription;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SubscribeCOVPropertyMultipleErrorFirstFailedSubscription)) {
      return false;
    }
    SubscribeCOVPropertyMultipleErrorFirstFailedSubscription that =
        (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getMonitoredObjectIdentifier() == that.getMonitoredObjectIdentifier())
        && (getMonitoredPropertyReference() == that.getMonitoredPropertyReference())
        && (getErrorType() == that.getErrorType())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getOpeningTag(),
        getMonitoredObjectIdentifier(),
        getMonitoredPropertyReference(),
        getErrorType(),
        getClosingTag());
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
