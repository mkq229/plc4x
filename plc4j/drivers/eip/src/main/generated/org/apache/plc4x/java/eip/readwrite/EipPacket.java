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

public abstract class EipPacket implements Message {

  // Abstract accessors for discriminator values.
  public abstract Integer getCommand();

  public abstract Integer getPacketLength();

  public abstract Boolean getResponse();

  // Properties.
  protected final long sessionHandle;
  protected final long status;
  protected final byte[] senderContext;
  protected final long options;

  public EipPacket(long sessionHandle, long status, byte[] senderContext, long options) {
    super();
    this.sessionHandle = sessionHandle;
    this.status = status;
    this.senderContext = senderContext;
    this.options = options;
  }

  public long getSessionHandle() {
    return sessionHandle;
  }

  public long getStatus() {
    return status;
  }

  public byte[] getSenderContext() {
    return senderContext;
  }

  public long getOptions() {
    return options;
  }

  protected abstract void serializeEipPacketChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EipPacket");

    // Discriminator Field (command) (Used as input to a switch field)
    writeDiscriminatorField("command", getCommand(), writeUnsignedInt(writeBuffer, 16));

    // Implicit Field (packetLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int packetLength = (int) ((getLengthInBytes()) - (24));
    writeImplicitField("packetLength", packetLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (sessionHandle)
    writeSimpleField("sessionHandle", sessionHandle, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedLong(writeBuffer, 32));

    // Array Field (senderContext)
    writeByteArrayField("senderContext", senderContext, writeByteArray(writeBuffer, 8));

    // Simple Field (options)
    writeSimpleField("options", options, writeUnsignedLong(writeBuffer, 32));

    // Switch field (Serialize the sub-type)
    serializeEipPacketChild(writeBuffer);

    writeBuffer.popContext("EipPacket");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    EipPacket _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (command)
    lengthInBits += 16;

    // Implicit Field (packetLength)
    lengthInBits += 16;

    // Simple field (sessionHandle)
    lengthInBits += 32;

    // Simple field (status)
    lengthInBits += 32;

    // Array field
    if (senderContext != null) {
      lengthInBits += 8 * senderContext.length;
    }

    // Simple field (options)
    lengthInBits += 32;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static EipPacket staticParse(ReadBuffer readBuffer, Boolean response)
      throws ParseException {
    readBuffer.pullContext("EipPacket");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int command = readDiscriminatorField("command", readUnsignedInt(readBuffer, 16));

    int packetLength = readImplicitField("packetLength", readUnsignedInt(readBuffer, 16));

    long sessionHandle = readSimpleField("sessionHandle", readUnsignedLong(readBuffer, 32));

    long status = readSimpleField("status", readUnsignedLong(readBuffer, 32));

    byte[] senderContext = readBuffer.readByteArray("senderContext", Math.toIntExact(8));

    long options = readSimpleField("options", readUnsignedLong(readBuffer, 32));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    EipPacketBuilder builder = null;
    if (EvaluationHelper.equals(command, (int) 0x0001)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = NullCommandRequest.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0001)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = NullCommandResponse.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0004)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = ListServicesRequest.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0004)
        && EvaluationHelper.equals(response, (boolean) true)
        && EvaluationHelper.equals(packetLength, (int) 0)) {
      builder = NullListServicesResponse.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0004)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = ListServicesResponse.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0063)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = EipListIdentityRequest.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0063)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = EipListIdentityResponse.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0065)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = EipConnectionRequest.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0065)
        && EvaluationHelper.equals(response, (boolean) true)
        && EvaluationHelper.equals(packetLength, (int) 0)) {
      builder = NullEipConnectionResponse.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0065)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = EipConnectionResponse.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0066)) {
      builder = EipDisconnectRequest.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x006F)) {
      builder = CipRRData.staticParseEipPacketBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(command, (int) 0x0070)) {
      builder = SendUnitData.staticParseEipPacketBuilder(readBuffer, response);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "command="
              + command
              + " "
              + "response="
              + response
              + " "
              + "packetLength="
              + packetLength
              + "]");
    }

    readBuffer.closeContext("EipPacket");
    // Create the instance
    EipPacket _eipPacket = builder.build(sessionHandle, status, senderContext, options);
    return _eipPacket;
  }

  public interface EipPacketBuilder {
    EipPacket build(long sessionHandle, long status, byte[] senderContext, long options);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EipPacket)) {
      return false;
    }
    EipPacket that = (EipPacket) o;
    return (getSessionHandle() == that.getSessionHandle())
        && (getStatus() == that.getStatus())
        && (getSenderContext() == that.getSenderContext())
        && (getOptions() == that.getOptions())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getSessionHandle(), getStatus(), getSenderContext(), getOptions());
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
