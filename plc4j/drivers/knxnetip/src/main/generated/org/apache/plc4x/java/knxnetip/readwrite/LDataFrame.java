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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public abstract class LDataFrame implements Message {

  // Abstract accessors for discriminator values.
  public abstract Boolean getNotAckFrame();

  public abstract Boolean getPolling();

  // Properties.
  protected final boolean frameType;
  protected final boolean notRepeated;
  protected final CEMIPriority priority;
  protected final boolean acknowledgeRequested;
  protected final boolean errorFlag;

  public LDataFrame(
      boolean frameType,
      boolean notRepeated,
      CEMIPriority priority,
      boolean acknowledgeRequested,
      boolean errorFlag) {
    super();
    this.frameType = frameType;
    this.notRepeated = notRepeated;
    this.priority = priority;
    this.acknowledgeRequested = acknowledgeRequested;
    this.errorFlag = errorFlag;
  }

  public boolean getFrameType() {
    return frameType;
  }

  public boolean getNotRepeated() {
    return notRepeated;
  }

  public CEMIPriority getPriority() {
    return priority;
  }

  public boolean getAcknowledgeRequested() {
    return acknowledgeRequested;
  }

  public boolean getErrorFlag() {
    return errorFlag;
  }

  protected abstract void serializeLDataFrameChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("LDataFrame");

    // Simple Field (frameType)
    writeSimpleField("frameType", frameType, writeBoolean(writeBuffer));

    // Discriminator Field (polling) (Used as input to a switch field)
    writeDiscriminatorField("polling", getPolling(), writeBoolean(writeBuffer));

    // Simple Field (notRepeated)
    writeSimpleField("notRepeated", notRepeated, writeBoolean(writeBuffer));

    // Discriminator Field (notAckFrame) (Used as input to a switch field)
    writeDiscriminatorField("notAckFrame", getNotAckFrame(), writeBoolean(writeBuffer));

    // Simple Field (priority)
    writeSimpleEnumField(
        "priority",
        "CEMIPriority",
        priority,
        new DataWriterEnumDefault<>(
            CEMIPriority::getValue, CEMIPriority::name, writeUnsignedByte(writeBuffer, 2)));

    // Simple Field (acknowledgeRequested)
    writeSimpleField("acknowledgeRequested", acknowledgeRequested, writeBoolean(writeBuffer));

    // Simple Field (errorFlag)
    writeSimpleField("errorFlag", errorFlag, writeBoolean(writeBuffer));

    // Switch field (Serialize the sub-type)
    serializeLDataFrameChild(writeBuffer);

    writeBuffer.popContext("LDataFrame");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LDataFrame _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (frameType)
    lengthInBits += 1;

    // Discriminator Field (polling)
    lengthInBits += 1;

    // Simple field (notRepeated)
    lengthInBits += 1;

    // Discriminator Field (notAckFrame)
    lengthInBits += 1;

    // Simple field (priority)
    lengthInBits += 2;

    // Simple field (acknowledgeRequested)
    lengthInBits += 1;

    // Simple field (errorFlag)
    lengthInBits += 1;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static LDataFrame staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("LDataFrame");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean frameType = readSimpleField("frameType", readBoolean(readBuffer));

    boolean polling = readDiscriminatorField("polling", readBoolean(readBuffer));

    boolean notRepeated = readSimpleField("notRepeated", readBoolean(readBuffer));

    boolean notAckFrame = readDiscriminatorField("notAckFrame", readBoolean(readBuffer));

    CEMIPriority priority =
        readEnumField(
            "priority",
            "CEMIPriority",
            new DataReaderEnumDefault<>(
                CEMIPriority::enumForValue, readUnsignedByte(readBuffer, 2)));

    boolean acknowledgeRequested = readSimpleField("acknowledgeRequested", readBoolean(readBuffer));

    boolean errorFlag = readSimpleField("errorFlag", readBoolean(readBuffer));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    LDataFrameBuilder builder = null;
    if (EvaluationHelper.equals(notAckFrame, (boolean) true)
        && EvaluationHelper.equals(polling, (boolean) false)) {
      builder = LDataExtended.staticParseLDataFrameBuilder(readBuffer);
    } else if (EvaluationHelper.equals(notAckFrame, (boolean) true)
        && EvaluationHelper.equals(polling, (boolean) true)) {
      builder = LPollData.staticParseLDataFrameBuilder(readBuffer);
    } else if (EvaluationHelper.equals(notAckFrame, (boolean) false)) {
      builder = LDataFrameACK.staticParseLDataFrameBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "notAckFrame="
              + notAckFrame
              + " "
              + "polling="
              + polling
              + "]");
    }

    readBuffer.closeContext("LDataFrame");
    // Create the instance
    LDataFrame _lDataFrame =
        builder.build(frameType, notRepeated, priority, acknowledgeRequested, errorFlag);
    return _lDataFrame;
  }

  public interface LDataFrameBuilder {
    LDataFrame build(
        boolean frameType,
        boolean notRepeated,
        CEMIPriority priority,
        boolean acknowledgeRequested,
        boolean errorFlag);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LDataFrame)) {
      return false;
    }
    LDataFrame that = (LDataFrame) o;
    return (getFrameType() == that.getFrameType())
        && (getNotRepeated() == that.getNotRepeated())
        && (getPriority() == that.getPriority())
        && (getAcknowledgeRequested() == that.getAcknowledgeRequested())
        && (getErrorFlag() == that.getErrorFlag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getFrameType(), getNotRepeated(), getPriority(), getAcknowledgeRequested(), getErrorFlag());
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
