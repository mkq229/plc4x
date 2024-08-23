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

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetPriorityArray implements Message {

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger numberOfDataElements;
  protected final List<BACnetPriorityValue> data;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetPriorityArray(
      BACnetApplicationTagUnsignedInteger numberOfDataElements,
      List<BACnetPriorityValue> data,
      BACnetObjectType objectTypeArgument,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super();
    this.numberOfDataElements = numberOfDataElements;
    this.data = data;
    this.objectTypeArgument = objectTypeArgument;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetApplicationTagUnsignedInteger getNumberOfDataElements() {
    return numberOfDataElements;
  }

  public List<BACnetPriorityValue> getData() {
    return data;
  }

  public BigInteger getZero() {
    Object o = 0L;
    if (o instanceof BigInteger) return (BigInteger) o;
    return BigInteger.valueOf(((Number) o).longValue());
  }

  public BACnetPriorityValue getPriorityValue01() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (0)) ? getData().get(0) : null));
  }

  public BACnetPriorityValue getPriorityValue02() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (1)) ? getData().get(1) : null));
  }

  public BACnetPriorityValue getPriorityValue03() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (2)) ? getData().get(2) : null));
  }

  public BACnetPriorityValue getPriorityValue04() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (3)) ? getData().get(3) : null));
  }

  public BACnetPriorityValue getPriorityValue05() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (4)) ? getData().get(4) : null));
  }

  public BACnetPriorityValue getPriorityValue06() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (5)) ? getData().get(5) : null));
  }

  public BACnetPriorityValue getPriorityValue07() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (6)) ? getData().get(6) : null));
  }

  public BACnetPriorityValue getPriorityValue08() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (7)) ? getData().get(7) : null));
  }

  public BACnetPriorityValue getPriorityValue09() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (8)) ? getData().get(8) : null));
  }

  public BACnetPriorityValue getPriorityValue10() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (9)) ? getData().get(9) : null));
  }

  public BACnetPriorityValue getPriorityValue11() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (10)) ? getData().get(10) : null));
  }

  public BACnetPriorityValue getPriorityValue12() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (11)) ? getData().get(11) : null));
  }

  public BACnetPriorityValue getPriorityValue13() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (12)) ? getData().get(12) : null));
  }

  public BACnetPriorityValue getPriorityValue14() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (13)) ? getData().get(13) : null));
  }

  public BACnetPriorityValue getPriorityValue15() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (14)) ? getData().get(14) : null));
  }

  public BACnetPriorityValue getPriorityValue16() {
    return (BACnetPriorityValue) ((((COUNT(getData())) > (15)) ? getData().get(15) : null));
  }

  public boolean getIsIndexedAccess() {
    return (boolean) ((COUNT(getData())) == (1));
  }

  public BACnetPriorityValue getIndexEntry() {
    return (BACnetPriorityValue) (getPriorityValue01());
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPriorityArray");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BigInteger zero = getZero();
    writeBuffer.writeVirtual("zero", zero);

    // Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
    writeOptionalField(
        "numberOfDataElements",
        numberOfDataElements,
        new DataWriterComplexDefault<>(writeBuffer),
        ((arrayIndexArgument) != (null)) && ((arrayIndexArgument.getActualValue()) == (getZero())));

    // Array Field (data)
    writeComplexTypeArrayField("data", data, writeBuffer);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue01 = getPriorityValue01();
    writeBuffer.writeVirtual("priorityValue01", priorityValue01);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue02 = getPriorityValue02();
    writeBuffer.writeVirtual("priorityValue02", priorityValue02);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue03 = getPriorityValue03();
    writeBuffer.writeVirtual("priorityValue03", priorityValue03);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue04 = getPriorityValue04();
    writeBuffer.writeVirtual("priorityValue04", priorityValue04);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue05 = getPriorityValue05();
    writeBuffer.writeVirtual("priorityValue05", priorityValue05);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue06 = getPriorityValue06();
    writeBuffer.writeVirtual("priorityValue06", priorityValue06);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue07 = getPriorityValue07();
    writeBuffer.writeVirtual("priorityValue07", priorityValue07);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue08 = getPriorityValue08();
    writeBuffer.writeVirtual("priorityValue08", priorityValue08);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue09 = getPriorityValue09();
    writeBuffer.writeVirtual("priorityValue09", priorityValue09);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue10 = getPriorityValue10();
    writeBuffer.writeVirtual("priorityValue10", priorityValue10);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue11 = getPriorityValue11();
    writeBuffer.writeVirtual("priorityValue11", priorityValue11);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue12 = getPriorityValue12();
    writeBuffer.writeVirtual("priorityValue12", priorityValue12);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue13 = getPriorityValue13();
    writeBuffer.writeVirtual("priorityValue13", priorityValue13);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue14 = getPriorityValue14();
    writeBuffer.writeVirtual("priorityValue14", priorityValue14);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue15 = getPriorityValue15();
    writeBuffer.writeVirtual("priorityValue15", priorityValue15);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue priorityValue16 = getPriorityValue16();
    writeBuffer.writeVirtual("priorityValue16", priorityValue16);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isIndexedAccess = getIsIndexedAccess();
    writeBuffer.writeVirtual("isIndexedAccess", isIndexedAccess);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetPriorityValue indexEntry = getIndexEntry();
    writeBuffer.writeVirtual("indexEntry", indexEntry);

    writeBuffer.popContext("BACnetPriorityArray");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetPriorityArray _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Optional Field (numberOfDataElements)
    if (numberOfDataElements != null) {
      lengthInBits += numberOfDataElements.getLengthInBits();
    }

    // Array field
    if (data != null) {
      for (Message element : data) {
        lengthInBits += element.getLengthInBits();
      }
    }

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetPriorityArray staticParse(
      ReadBuffer readBuffer,
      BACnetObjectType objectTypeArgument,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetPriorityArray");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    BigInteger zero = readVirtualField("zero", BigInteger.class, 0L);

    BACnetApplicationTagUnsignedInteger numberOfDataElements =
        readOptionalField(
            "numberOfDataElements",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer),
            ((arrayIndexArgument) != (null)) && ((arrayIndexArgument.getActualValue()) == (zero)));

    List<BACnetPriorityValue> data =
        readTerminatedArrayField(
            "data",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPriorityValue.staticParse(
                        readBuffer, (BACnetObjectType) (objectTypeArgument)),
                readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .isBACnetConstructedDataClosingTag(readBuffer, false, tagNumber))));
    BACnetPriorityValue priorityValue01 =
        readVirtualField(
            "priorityValue01",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (0)) ? data.get(0) : null));
    BACnetPriorityValue priorityValue02 =
        readVirtualField(
            "priorityValue02",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (1)) ? data.get(1) : null));
    BACnetPriorityValue priorityValue03 =
        readVirtualField(
            "priorityValue03",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (2)) ? data.get(2) : null));
    BACnetPriorityValue priorityValue04 =
        readVirtualField(
            "priorityValue04",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (3)) ? data.get(3) : null));
    BACnetPriorityValue priorityValue05 =
        readVirtualField(
            "priorityValue05",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (4)) ? data.get(4) : null));
    BACnetPriorityValue priorityValue06 =
        readVirtualField(
            "priorityValue06",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (5)) ? data.get(5) : null));
    BACnetPriorityValue priorityValue07 =
        readVirtualField(
            "priorityValue07",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (6)) ? data.get(6) : null));
    BACnetPriorityValue priorityValue08 =
        readVirtualField(
            "priorityValue08",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (7)) ? data.get(7) : null));
    BACnetPriorityValue priorityValue09 =
        readVirtualField(
            "priorityValue09",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (8)) ? data.get(8) : null));
    BACnetPriorityValue priorityValue10 =
        readVirtualField(
            "priorityValue10",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (9)) ? data.get(9) : null));
    BACnetPriorityValue priorityValue11 =
        readVirtualField(
            "priorityValue11",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (10)) ? data.get(10) : null));
    BACnetPriorityValue priorityValue12 =
        readVirtualField(
            "priorityValue12",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (11)) ? data.get(11) : null));
    BACnetPriorityValue priorityValue13 =
        readVirtualField(
            "priorityValue13",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (12)) ? data.get(12) : null));
    BACnetPriorityValue priorityValue14 =
        readVirtualField(
            "priorityValue14",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (13)) ? data.get(13) : null));
    BACnetPriorityValue priorityValue15 =
        readVirtualField(
            "priorityValue15",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (14)) ? data.get(14) : null));
    BACnetPriorityValue priorityValue16 =
        readVirtualField(
            "priorityValue16",
            BACnetPriorityValue.class,
            (((COUNT(data)) > (15)) ? data.get(15) : null));
    // Validation
    if (!(((arrayIndexArgument) != (null)) || ((COUNT(data)) == (16)))) {
      throw new ParseValidationException("Either indexed access or lenght 16 expected");
    }
    boolean isIndexedAccess =
        readVirtualField("isIndexedAccess", boolean.class, (COUNT(data)) == (1));
    BACnetPriorityValue indexEntry =
        readVirtualField("indexEntry", BACnetPriorityValue.class, priorityValue01);

    readBuffer.closeContext("BACnetPriorityArray");
    // Create the instance
    BACnetPriorityArray _bACnetPriorityArray;
    _bACnetPriorityArray =
        new BACnetPriorityArray(
            numberOfDataElements, data, objectTypeArgument, tagNumber, arrayIndexArgument);
    return _bACnetPriorityArray;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPriorityArray)) {
      return false;
    }
    BACnetPriorityArray that = (BACnetPriorityArray) o;
    return (getNumberOfDataElements() == that.getNumberOfDataElements())
        && (getData() == that.getData())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNumberOfDataElements(), getData());
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
