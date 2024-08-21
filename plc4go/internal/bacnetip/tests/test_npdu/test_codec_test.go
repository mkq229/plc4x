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

package test_npdu

import (
	"testing"

	"github.com/apache/plc4x/plc4go/internal/bacnetip"
	"github.com/apache/plc4x/plc4go/internal/bacnetip/tests"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/testutils"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

func WhoIsRouterToNetwork(net uint16) *bacnetip.WhoIsRouterToNetwork {
	network, err := bacnetip.NewWhoIsRouterToNetwork(bacnetip.WithWhoIsRouterToNetworkNet(net))
	if err != nil {
		panic(err)
	}
	return network
}

func IAmRouterToNetwork(netList ...uint16) *bacnetip.IAmRouterToNetwork {
	network, err := bacnetip.NewIAmRouterToNetwork(bacnetip.WithIAmRouterToNetworkNetworkList(netList...))
	if err != nil {
		panic(err)
	}
	return network
}

func ICouldBeRouterToNetwork(net uint16, perf uint8) *bacnetip.ICouldBeRouterToNetwork {
	network, err := bacnetip.NewICouldBeRouterToNetwork(bacnetip.WithICouldBeRouterToNetworkNetwork(net), bacnetip.WithICouldBeRouterToNetworkPerformanceIndex(perf))
	if err != nil {
		panic(err)
	}
	return network
}

func RejectMessageToNetwork(reason uint8, dnet uint16) *bacnetip.RejectMessageToNetwork {
	network, err := bacnetip.NewRejectMessageToNetwork(bacnetip.WithRejectMessageToNetworkRejectionReason(readWriteModel.NLMRejectMessageToNetworkRejectReason(reason)), bacnetip.WithRejectMessageToNetworkDnet(dnet))
	if err != nil {
		panic(err)
	}
	return network
}

type TestNPDUCodecSuite struct {
	suite.Suite

	client *tests.TrappedClient
	codec  *NPDUCodec
	server *tests.TrappedServer

	log zerolog.Logger
}

func (suite *TestNPDUCodecSuite) SetupSuite() {
	t := suite.T()
	suite.log = testutils.ProduceTestingLogger(t)
}

func (suite *TestNPDUCodecSuite) SetupTest() {
	var err error
	suite.codec, err = NewNPDUCodec(suite.log)
	suite.Require().NoError(err)
	suite.client, err = tests.NewTrappedClient(suite.log)
	suite.Require().NoError(err)
	suite.server, err = tests.NewTrappedServer(suite.log)
	suite.Require().NoError(err)
	err = bacnetip.Bind(suite.log, suite.client, suite.codec, suite.server)
	suite.Require().NoError(err)
}

// Pass the PDU to the client to send down the stack.
func (suite *TestNPDUCodecSuite) Request(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Request")

	return suite.client.Request(args, kwargs)
}

// Check what the server received.
func (suite *TestNPDUCodecSuite) Indication(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Indication")

	var pduType any
	if len(args) > 0 {
		pduType = args[0].(any)
	}
	pduAttrs := kwargs
	suite.Assert().True(tests.MatchPdu(suite.log, suite.server.GetIndicationReceived(), pduType, pduAttrs))
	return nil
}

// Check what the server received.
func (suite *TestNPDUCodecSuite) Response(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Response")

	return suite.server.Response(args, kwargs)
}

// Check what the server received.
func (suite *TestNPDUCodecSuite) Confirmation(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Confirmation")

	pduType := args[0].(any)
	pduAttrs := kwargs
	suite.Assert().True(tests.MatchPdu(suite.log, suite.client.GetConfirmationReceived(), pduType, pduAttrs))
	return nil
}

func (suite *TestNPDUCodecSuite) TestWhoIsRouterToNetwork() { // Test the Result encoding and decoding.
	// Request successful
	pduBytes, err := bacnetip.Xtob(
		"01.80" + // version, network layer message
			"00 0001", // message type and network
	)
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.NPDUParse(testutils.TestContext(suite.T()), pduBytes, uint16(len(pduBytes)))
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(WhoIsRouterToNetwork(1)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs(&bacnetip.WhoIsRouterToNetwork{}), bacnetip.NewKWArgs(bacnetip.KWWirtnNetwork, uint16(1)))
}

func (suite *TestNPDUCodecSuite) TestIAMRouterToNetworkEmpty() { // Test the Result encoding and decoding.
	// Request successful
	networkList := []uint16{}
	pduBytes, err := bacnetip.Xtob(
		"01.80" + // version, network layer message
			"01", // message type, no network
	)
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.NPDUParse(testutils.TestContext(suite.T()), pduBytes, uint16(len(pduBytes)))
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(IAmRouterToNetwork(networkList...)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs(&bacnetip.IAmRouterToNetwork{}), bacnetip.NewKWArgs(bacnetip.KWIartnNetworkList, networkList))
}

func (suite *TestNPDUCodecSuite) TestIAMRouterToNetworks() { // Test the Result encoding and decoding.
	// Request successful
	networkList := []uint16{1, 2, 3}
	pduBytes, err := bacnetip.Xtob(
		"01.80" + // version, network layer message
			"01 0001 0002 0003", // message type and network list
	)
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.NPDUParse(testutils.TestContext(suite.T()), pduBytes, uint16(len(pduBytes)))
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(IAmRouterToNetwork(networkList...)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs(&bacnetip.IAmRouterToNetwork{}), bacnetip.NewKWArgs(bacnetip.KWIartnNetworkList, networkList))
}

func (suite *TestNPDUCodecSuite) TestICouldBeRouterToNetworks() { // Test the Result encoding and decoding.
	// Request successful
	pduBytes, err := bacnetip.Xtob(
		"01.80" + // version, network layer message
			"02 0001 02", // message type, network, performance
	)
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.NPDUParse(testutils.TestContext(suite.T()), pduBytes, uint16(len(pduBytes)))
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(ICouldBeRouterToNetwork(1, 2)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs(&bacnetip.ICouldBeRouterToNetwork{}), bacnetip.NewKWArgs(bacnetip.KWIcbrtnNetwork, uint16(1), bacnetip.KWIcbrtnPerformanceIndex, uint8(2)))
}

func (suite *TestNPDUCodecSuite) TestRejectMessageToNetworks() { // Test the Result encoding and decoding.
	// Request successful
	pduBytes, err := bacnetip.Xtob(
		"01.80" + // version, network layer message
			"03 01 0002", // message type, reason, performance
	)
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.NPDUParse(testutils.TestContext(suite.T()), pduBytes, uint16(len(pduBytes)))
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(RejectMessageToNetwork(1, 2)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs(&bacnetip.RejectMessageToNetwork{}), bacnetip.NewKWArgs(bacnetip.KWRmtnRejectionReason, readWriteModel.NLMRejectMessageToNetworkRejectReason(1), bacnetip.KWRmtnDNET, uint16(2)))
}

func TestNPDUCodec(t *testing.T) {
	suite.Run(t, new(TestNPDUCodecSuite))
}
