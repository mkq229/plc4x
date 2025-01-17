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

public class CipIdentity extends CommandSpecificDataItem implements Message {

  // Accessors for discriminator values.
  public Integer getItemType() {
    return (int) 0x000C;
  }

  // Constant values.
  public static final Long ZEROES1 = 0x00000000L;
  public static final Long ZEROES2 = 0x00000000L;

  // Properties.
  protected final int encapsulationProtocolVersion;
  protected final int socketAddressFamily;
  protected final int socketAddressPort;
  protected final List<Short> socketAddressAddress;
  protected final int vendorId;
  protected final int deviceType;
  protected final int productCode;
  protected final short revisionMajor;
  protected final short revisionMinor;
  protected final int status;
  protected final long serialNumber;
  protected final String productName;
  protected final short state;

  public CipIdentity(
      int encapsulationProtocolVersion,
      int socketAddressFamily,
      int socketAddressPort,
      List<Short> socketAddressAddress,
      int vendorId,
      int deviceType,
      int productCode,
      short revisionMajor,
      short revisionMinor,
      int status,
      long serialNumber,
      String productName,
      short state) {
    super();
    this.encapsulationProtocolVersion = encapsulationProtocolVersion;
    this.socketAddressFamily = socketAddressFamily;
    this.socketAddressPort = socketAddressPort;
    this.socketAddressAddress = socketAddressAddress;
    this.vendorId = vendorId;
    this.deviceType = deviceType;
    this.productCode = productCode;
    this.revisionMajor = revisionMajor;
    this.revisionMinor = revisionMinor;
    this.status = status;
    this.serialNumber = serialNumber;
    this.productName = productName;
    this.state = state;
  }

  public int getEncapsulationProtocolVersion() {
    return encapsulationProtocolVersion;
  }

  public int getSocketAddressFamily() {
    return socketAddressFamily;
  }

  public int getSocketAddressPort() {
    return socketAddressPort;
  }

  public List<Short> getSocketAddressAddress() {
    return socketAddressAddress;
  }

  public int getVendorId() {
    return vendorId;
  }

  public int getDeviceType() {
    return deviceType;
  }

  public int getProductCode() {
    return productCode;
  }

  public short getRevisionMajor() {
    return revisionMajor;
  }

  public short getRevisionMinor() {
    return revisionMinor;
  }

  public int getStatus() {
    return status;
  }

  public long getSerialNumber() {
    return serialNumber;
  }

  public String getProductName() {
    return productName;
  }

  public short getState() {
    return state;
  }

  public long getZeroes1() {
    return ZEROES1;
  }

  public long getZeroes2() {
    return ZEROES2;
  }

  @Override
  protected void serializeCommandSpecificDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CipIdentity");

    // Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int itemLength = (int) ((34) + (STR_LEN(getProductName())));
    writeImplicitField("itemLength", itemLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (encapsulationProtocolVersion)
    writeSimpleField(
        "encapsulationProtocolVersion",
        encapsulationProtocolVersion,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (socketAddressFamily)
    writeSimpleField(
        "socketAddressFamily",
        socketAddressFamily,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (socketAddressPort)
    writeSimpleField(
        "socketAddressPort",
        socketAddressPort,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (socketAddressAddress)
    writeSimpleTypeArrayField(
        "socketAddressAddress", socketAddressAddress, writeUnsignedShort(writeBuffer, 8));

    // Const Field (zeroes1)
    writeConstField("zeroes1", ZEROES1, writeUnsignedLong(writeBuffer, 32));

    // Const Field (zeroes2)
    writeConstField("zeroes2", ZEROES2, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (vendorId)
    writeSimpleField("vendorId", vendorId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (deviceType)
    writeSimpleField("deviceType", deviceType, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (productCode)
    writeSimpleField("productCode", productCode, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (revisionMajor)
    writeSimpleField("revisionMajor", revisionMajor, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (revisionMinor)
    writeSimpleField("revisionMinor", revisionMinor, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (serialNumber)
    writeSimpleField("serialNumber", serialNumber, writeUnsignedLong(writeBuffer, 32));

    // Implicit Field (productNameLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short productNameLength = (short) (STR_LEN(getProductName()));
    writeImplicitField("productNameLength", productNameLength, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (productName)
    writeSimpleField(
        "productName", productName, writeString(writeBuffer, (productNameLength) * (8)));

    // Simple Field (state)
    writeSimpleField("state", state, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("CipIdentity");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipIdentity _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (itemLength)
    lengthInBits += 16;

    // Simple field (encapsulationProtocolVersion)
    lengthInBits += 16;

    // Simple field (socketAddressFamily)
    lengthInBits += 16;

    // Simple field (socketAddressPort)
    lengthInBits += 16;

    // Array field
    if (socketAddressAddress != null) {
      lengthInBits += 8 * socketAddressAddress.size();
    }

    // Const Field (zeroes1)
    lengthInBits += 32;

    // Const Field (zeroes2)
    lengthInBits += 32;

    // Simple field (vendorId)
    lengthInBits += 16;

    // Simple field (deviceType)
    lengthInBits += 16;

    // Simple field (productCode)
    lengthInBits += 16;

    // Simple field (revisionMajor)
    lengthInBits += 8;

    // Simple field (revisionMinor)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 16;

    // Simple field (serialNumber)
    lengthInBits += 32;

    // Implicit Field (productNameLength)
    lengthInBits += 8;

    // Simple field (productName)
    lengthInBits += (STR_LEN(getProductName())) * (8);

    // Simple field (state)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static CommandSpecificDataItemBuilder staticParseCommandSpecificDataItemBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CipIdentity");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int itemLength = readImplicitField("itemLength", readUnsignedInt(readBuffer, 16));

    int encapsulationProtocolVersion =
        readSimpleField(
            "encapsulationProtocolVersion",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int socketAddressFamily =
        readSimpleField(
            "socketAddressFamily",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int socketAddressPort =
        readSimpleField(
            "socketAddressPort",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<Short> socketAddressAddress =
        readCountArrayField("socketAddressAddress", readUnsignedShort(readBuffer, 8), 4);

    long zeroes1 = readConstField("zeroes1", readUnsignedLong(readBuffer, 32), CipIdentity.ZEROES1);

    long zeroes2 = readConstField("zeroes2", readUnsignedLong(readBuffer, 32), CipIdentity.ZEROES2);

    int vendorId = readSimpleField("vendorId", readUnsignedInt(readBuffer, 16));

    int deviceType = readSimpleField("deviceType", readUnsignedInt(readBuffer, 16));

    int productCode = readSimpleField("productCode", readUnsignedInt(readBuffer, 16));

    short revisionMajor = readSimpleField("revisionMajor", readUnsignedShort(readBuffer, 8));

    short revisionMinor = readSimpleField("revisionMinor", readUnsignedShort(readBuffer, 8));

    int status = readSimpleField("status", readUnsignedInt(readBuffer, 16));

    long serialNumber = readSimpleField("serialNumber", readUnsignedLong(readBuffer, 32));

    short productNameLength =
        readImplicitField("productNameLength", readUnsignedShort(readBuffer, 8));

    String productName =
        readSimpleField("productName", readString(readBuffer, (productNameLength) * (8)));

    short state = readSimpleField("state", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("CipIdentity");
    // Create the instance
    return new CipIdentityBuilderImpl(
        encapsulationProtocolVersion,
        socketAddressFamily,
        socketAddressPort,
        socketAddressAddress,
        vendorId,
        deviceType,
        productCode,
        revisionMajor,
        revisionMinor,
        status,
        serialNumber,
        productName,
        state);
  }

  public static class CipIdentityBuilderImpl
      implements CommandSpecificDataItem.CommandSpecificDataItemBuilder {
    private final int encapsulationProtocolVersion;
    private final int socketAddressFamily;
    private final int socketAddressPort;
    private final List<Short> socketAddressAddress;
    private final int vendorId;
    private final int deviceType;
    private final int productCode;
    private final short revisionMajor;
    private final short revisionMinor;
    private final int status;
    private final long serialNumber;
    private final String productName;
    private final short state;

    public CipIdentityBuilderImpl(
        int encapsulationProtocolVersion,
        int socketAddressFamily,
        int socketAddressPort,
        List<Short> socketAddressAddress,
        int vendorId,
        int deviceType,
        int productCode,
        short revisionMajor,
        short revisionMinor,
        int status,
        long serialNumber,
        String productName,
        short state) {
      this.encapsulationProtocolVersion = encapsulationProtocolVersion;
      this.socketAddressFamily = socketAddressFamily;
      this.socketAddressPort = socketAddressPort;
      this.socketAddressAddress = socketAddressAddress;
      this.vendorId = vendorId;
      this.deviceType = deviceType;
      this.productCode = productCode;
      this.revisionMajor = revisionMajor;
      this.revisionMinor = revisionMinor;
      this.status = status;
      this.serialNumber = serialNumber;
      this.productName = productName;
      this.state = state;
    }

    public CipIdentity build() {
      CipIdentity cipIdentity =
          new CipIdentity(
              encapsulationProtocolVersion,
              socketAddressFamily,
              socketAddressPort,
              socketAddressAddress,
              vendorId,
              deviceType,
              productCode,
              revisionMajor,
              revisionMinor,
              status,
              serialNumber,
              productName,
              state);
      return cipIdentity;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipIdentity)) {
      return false;
    }
    CipIdentity that = (CipIdentity) o;
    return (getEncapsulationProtocolVersion() == that.getEncapsulationProtocolVersion())
        && (getSocketAddressFamily() == that.getSocketAddressFamily())
        && (getSocketAddressPort() == that.getSocketAddressPort())
        && (getSocketAddressAddress() == that.getSocketAddressAddress())
        && (getVendorId() == that.getVendorId())
        && (getDeviceType() == that.getDeviceType())
        && (getProductCode() == that.getProductCode())
        && (getRevisionMajor() == that.getRevisionMajor())
        && (getRevisionMinor() == that.getRevisionMinor())
        && (getStatus() == that.getStatus())
        && (getSerialNumber() == that.getSerialNumber())
        && (getProductName() == that.getProductName())
        && (getState() == that.getState())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getEncapsulationProtocolVersion(),
        getSocketAddressFamily(),
        getSocketAddressPort(),
        getSocketAddressAddress(),
        getVendorId(),
        getDeviceType(),
        getProductCode(),
        getRevisionMajor(),
        getRevisionMinor(),
        getStatus(),
        getSerialNumber(),
        getProductName(),
        getState());
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
