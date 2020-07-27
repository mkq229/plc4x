/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "modbus_serial_adu.h"


// Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_serial_adu_parse(plc4c_spi_read_buffer* buf, bool response, plc4c_modbus_read_write_modbus_serial_adu** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(buf);
  uint16_t curPos;
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_modbus_read_write_modbus_serial_adu));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (transactionId)
  uint16_t transactionId = 0;
  _res = plc4c_spi_read_unsigned_short(buf, 16, (uint16_t*) &transactionId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->transaction_id = transactionId;

  // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
  {
    uint16_t _reserved = 0;
    _res = plc4c_spi_read_unsigned_short(buf, 16, (uint16_t*) &_reserved);
    if(_res != OK) {
      return _res;
    }
    if(_reserved != 0x0000) {
      printf("Expected constant value '%d' but got '%d' for reserved field.", 0x0000, _reserved);
    }
  }

  // Simple Field (length)
  uint16_t length = 0;
  _res = plc4c_spi_read_unsigned_short(buf, 16, (uint16_t*) &length);
  if(_res != OK) {
    return _res;
  }
  (*_message)->length = length;

  // Simple Field (address)
  uint8_t address = 0;
  _res = plc4c_spi_read_unsigned_byte(buf, 8, (uint8_t*) &address);
  if(_res != OK) {
    return _res;
  }
  (*_message)->address = address;

  // Simple Field (pdu)
  plc4c_modbus_read_write_modbus_pdu* pdu;
  _res = plc4c_modbus_read_write_modbus_pdu_parse(buf, response, (void*) &pdu);
  if(_res != OK) {
    return _res;
  }
  (*_message)->pdu = pdu;

  return OK;
}

plc4c_return_code plc4c_modbus_read_write_modbus_serial_adu_serialize(plc4c_spi_write_buffer* buf, plc4c_modbus_read_write_modbus_serial_adu* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (transactionId)
  _res = plc4c_spi_write_unsigned_short(buf, 16, _message->transaction_id);
  if(_res != OK) {
    return _res;
  }

  // Reserved Field
  _res = plc4c_spi_write_unsigned_short(buf, 16, 0x0000);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (length)
  _res = plc4c_spi_write_unsigned_short(buf, 16, _message->length);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (address)
  _res = plc4c_spi_write_unsigned_byte(buf, 8, _message->address);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (pdu)
  _res = plc4c_modbus_read_write_modbus_pdu_serialize(buf, _message->pdu);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_modbus_read_write_modbus_serial_adu_length_in_bytes(plc4c_modbus_read_write_modbus_serial_adu* _message) {
  return plc4c_modbus_read_write_modbus_serial_adu_length_in_bits(_message) / 8;
}

uint16_t plc4c_modbus_read_write_modbus_serial_adu_length_in_bits(plc4c_modbus_read_write_modbus_serial_adu* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (transactionId)
  lengthInBits += 16;

  // Reserved Field (reserved)
  lengthInBits += 16;

  // Simple field (length)
  lengthInBits += 16;

  // Simple field (address)
  lengthInBits += 8;

  // Simple field (pdu)
  lengthInBits += plc4c_modbus_read_write_modbus_pdu_length_in_bits(_message->pdu);

  return lengthInBits;
}

