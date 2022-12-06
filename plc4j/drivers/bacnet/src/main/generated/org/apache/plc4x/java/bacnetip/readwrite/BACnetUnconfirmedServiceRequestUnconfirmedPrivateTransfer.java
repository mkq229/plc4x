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

public class BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
    extends BACnetUnconfirmedServiceRequest implements Message {

  // Accessors for discriminator values.
  public BACnetUnconfirmedServiceChoice getServiceChoice() {
    return BACnetUnconfirmedServiceChoice.UNCONFIRMED_PRIVATE_TRANSFER;
  }

  // Properties.
  protected final BACnetVendorIdTagged vendorId;
  protected final BACnetContextTagUnsignedInteger serviceNumber;
  protected final BACnetConstructedData serviceParameters;

  // Arguments.
  protected final Integer serviceRequestLength;

  public BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(
      BACnetVendorIdTagged vendorId,
      BACnetContextTagUnsignedInteger serviceNumber,
      BACnetConstructedData serviceParameters,
      Integer serviceRequestLength) {
    super(serviceRequestLength);
    this.vendorId = vendorId;
    this.serviceNumber = serviceNumber;
    this.serviceParameters = serviceParameters;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetVendorIdTagged getVendorId() {
    return vendorId;
  }

  public BACnetContextTagUnsignedInteger getServiceNumber() {
    return serviceNumber;
  }

  public BACnetConstructedData getServiceParameters() {
    return serviceParameters;
  }

  @Override
  protected void serializeBACnetUnconfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer");

    // Simple Field (vendorId)
    writeSimpleField("vendorId", vendorId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (serviceNumber)
    writeSimpleField("serviceNumber", serviceNumber, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (serviceParameters) (Can be skipped, if the value is null)
    writeOptionalField(
        "serviceParameters", serviceParameters, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer _value = this;

    // Simple field (vendorId)
    lengthInBits += vendorId.getLengthInBits();

    // Simple field (serviceNumber)
    lengthInBits += serviceNumber.getLengthInBits();

    // Optional Field (serviceParameters)
    if (serviceParameters != null) {
      lengthInBits += serviceParameters.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferBuilder staticParseBuilder(
      ReadBuffer readBuffer, Integer serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetVendorIdTagged vendorId =
        readSimpleField(
            "vendorId",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetVendorIdTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger serviceNumber =
        readSimpleField(
            "serviceNumber",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData serviceParameters =
        readOptionalField(
            "serviceParameters",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (2),
                        (BACnetObjectType) (BACnetObjectType.VENDOR_PROPRIETARY_VALUE),
                        (BACnetPropertyIdentifier)
                            (BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE),
                        (BACnetTagPayloadUnsignedInteger) (null)),
                readBuffer));

    readBuffer.closeContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer");
    // Create the instance
    return new BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferBuilder(
        vendorId, serviceNumber, serviceParameters, serviceRequestLength);
  }

  public static class BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferBuilder
      implements BACnetUnconfirmedServiceRequest.BACnetUnconfirmedServiceRequestBuilder {
    private final BACnetVendorIdTagged vendorId;
    private final BACnetContextTagUnsignedInteger serviceNumber;
    private final BACnetConstructedData serviceParameters;
    private final Integer serviceRequestLength;

    public BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferBuilder(
        BACnetVendorIdTagged vendorId,
        BACnetContextTagUnsignedInteger serviceNumber,
        BACnetConstructedData serviceParameters,
        Integer serviceRequestLength) {

      this.vendorId = vendorId;
      this.serviceNumber = serviceNumber;
      this.serviceParameters = serviceParameters;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer build(
        Integer serviceRequestLength) {
      BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
          bACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer =
              new BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(
                  vendorId, serviceNumber, serviceParameters, serviceRequestLength);
      return bACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer)) {
      return false;
    }
    BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer that =
        (BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) o;
    return (getVendorId() == that.getVendorId())
        && (getServiceNumber() == that.getServiceNumber())
        && (getServiceParameters() == that.getServiceParameters())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getVendorId(), getServiceNumber(), getServiceParameters());
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