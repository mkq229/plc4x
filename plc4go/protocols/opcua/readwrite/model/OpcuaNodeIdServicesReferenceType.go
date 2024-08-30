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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesReferenceType is an enum
type OpcuaNodeIdServicesReferenceType int32

type IOpcuaNodeIdServicesReferenceType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesReferenceType_References                          OpcuaNodeIdServicesReferenceType = 31
	OpcuaNodeIdServicesReferenceType_NonHierarchicalReferences           OpcuaNodeIdServicesReferenceType = 32
	OpcuaNodeIdServicesReferenceType_HierarchicalReferences              OpcuaNodeIdServicesReferenceType = 33
	OpcuaNodeIdServicesReferenceType_HasChild                            OpcuaNodeIdServicesReferenceType = 34
	OpcuaNodeIdServicesReferenceType_Organizes                           OpcuaNodeIdServicesReferenceType = 35
	OpcuaNodeIdServicesReferenceType_HasEventSource                      OpcuaNodeIdServicesReferenceType = 36
	OpcuaNodeIdServicesReferenceType_HasModellingRule                    OpcuaNodeIdServicesReferenceType = 37
	OpcuaNodeIdServicesReferenceType_HasEncoding                         OpcuaNodeIdServicesReferenceType = 38
	OpcuaNodeIdServicesReferenceType_HasDescription                      OpcuaNodeIdServicesReferenceType = 39
	OpcuaNodeIdServicesReferenceType_HasTypeDefinition                   OpcuaNodeIdServicesReferenceType = 40
	OpcuaNodeIdServicesReferenceType_GeneratesEvent                      OpcuaNodeIdServicesReferenceType = 41
	OpcuaNodeIdServicesReferenceType_Aggregates                          OpcuaNodeIdServicesReferenceType = 44
	OpcuaNodeIdServicesReferenceType_HasSubtype                          OpcuaNodeIdServicesReferenceType = 45
	OpcuaNodeIdServicesReferenceType_HasProperty                         OpcuaNodeIdServicesReferenceType = 46
	OpcuaNodeIdServicesReferenceType_HasComponent                        OpcuaNodeIdServicesReferenceType = 47
	OpcuaNodeIdServicesReferenceType_HasNotifier                         OpcuaNodeIdServicesReferenceType = 48
	OpcuaNodeIdServicesReferenceType_HasOrderedComponent                 OpcuaNodeIdServicesReferenceType = 49
	OpcuaNodeIdServicesReferenceType_FromState                           OpcuaNodeIdServicesReferenceType = 51
	OpcuaNodeIdServicesReferenceType_ToState                             OpcuaNodeIdServicesReferenceType = 52
	OpcuaNodeIdServicesReferenceType_HasCause                            OpcuaNodeIdServicesReferenceType = 53
	OpcuaNodeIdServicesReferenceType_HasEffect                           OpcuaNodeIdServicesReferenceType = 54
	OpcuaNodeIdServicesReferenceType_HasHistoricalConfiguration          OpcuaNodeIdServicesReferenceType = 56
	OpcuaNodeIdServicesReferenceType_HasSubStateMachine                  OpcuaNodeIdServicesReferenceType = 117
	OpcuaNodeIdServicesReferenceType_HasArgumentDescription              OpcuaNodeIdServicesReferenceType = 129
	OpcuaNodeIdServicesReferenceType_HasOptionalInputArgumentDescription OpcuaNodeIdServicesReferenceType = 131
	OpcuaNodeIdServicesReferenceType_AlwaysGeneratesEvent                OpcuaNodeIdServicesReferenceType = 3065
	OpcuaNodeIdServicesReferenceType_HasTrueSubState                     OpcuaNodeIdServicesReferenceType = 9004
	OpcuaNodeIdServicesReferenceType_HasFalseSubState                    OpcuaNodeIdServicesReferenceType = 9005
	OpcuaNodeIdServicesReferenceType_HasCondition                        OpcuaNodeIdServicesReferenceType = 9006
	OpcuaNodeIdServicesReferenceType_HasPubSubConnection                 OpcuaNodeIdServicesReferenceType = 14476
	OpcuaNodeIdServicesReferenceType_DataSetToWriter                     OpcuaNodeIdServicesReferenceType = 14936
	OpcuaNodeIdServicesReferenceType_HasGuard                            OpcuaNodeIdServicesReferenceType = 15112
	OpcuaNodeIdServicesReferenceType_HasDataSetWriter                    OpcuaNodeIdServicesReferenceType = 15296
	OpcuaNodeIdServicesReferenceType_HasDataSetReader                    OpcuaNodeIdServicesReferenceType = 15297
	OpcuaNodeIdServicesReferenceType_HasAlarmSuppressionGroup            OpcuaNodeIdServicesReferenceType = 16361
	OpcuaNodeIdServicesReferenceType_AlarmGroupMember                    OpcuaNodeIdServicesReferenceType = 16362
	OpcuaNodeIdServicesReferenceType_HasEffectDisable                    OpcuaNodeIdServicesReferenceType = 17276
	OpcuaNodeIdServicesReferenceType_HasDictionaryEntry                  OpcuaNodeIdServicesReferenceType = 17597
	OpcuaNodeIdServicesReferenceType_HasInterface                        OpcuaNodeIdServicesReferenceType = 17603
	OpcuaNodeIdServicesReferenceType_HasAddIn                            OpcuaNodeIdServicesReferenceType = 17604
	OpcuaNodeIdServicesReferenceType_HasEffectEnable                     OpcuaNodeIdServicesReferenceType = 17983
	OpcuaNodeIdServicesReferenceType_HasEffectSuppressed                 OpcuaNodeIdServicesReferenceType = 17984
	OpcuaNodeIdServicesReferenceType_HasEffectUnsuppressed               OpcuaNodeIdServicesReferenceType = 17985
	OpcuaNodeIdServicesReferenceType_HasWriterGroup                      OpcuaNodeIdServicesReferenceType = 18804
	OpcuaNodeIdServicesReferenceType_HasReaderGroup                      OpcuaNodeIdServicesReferenceType = 18805
	OpcuaNodeIdServicesReferenceType_AliasFor                            OpcuaNodeIdServicesReferenceType = 23469
	OpcuaNodeIdServicesReferenceType_IsDeprecated                        OpcuaNodeIdServicesReferenceType = 23562
	OpcuaNodeIdServicesReferenceType_HasStructuredComponent              OpcuaNodeIdServicesReferenceType = 24136
	OpcuaNodeIdServicesReferenceType_AssociatedWith                      OpcuaNodeIdServicesReferenceType = 24137
	OpcuaNodeIdServicesReferenceType_UsesPriorityMappingTable            OpcuaNodeIdServicesReferenceType = 25237
	OpcuaNodeIdServicesReferenceType_HasLowerLayerInterface              OpcuaNodeIdServicesReferenceType = 25238
	OpcuaNodeIdServicesReferenceType_IsExecutableOn                      OpcuaNodeIdServicesReferenceType = 25253
	OpcuaNodeIdServicesReferenceType_Controls                            OpcuaNodeIdServicesReferenceType = 25254
	OpcuaNodeIdServicesReferenceType_Utilizes                            OpcuaNodeIdServicesReferenceType = 25255
	OpcuaNodeIdServicesReferenceType_Requires                            OpcuaNodeIdServicesReferenceType = 25256
	OpcuaNodeIdServicesReferenceType_IsPhysicallyConnectedTo             OpcuaNodeIdServicesReferenceType = 25257
	OpcuaNodeIdServicesReferenceType_RepresentsSameEntityAs              OpcuaNodeIdServicesReferenceType = 25258
	OpcuaNodeIdServicesReferenceType_RepresentsSameHardwareAs            OpcuaNodeIdServicesReferenceType = 25259
	OpcuaNodeIdServicesReferenceType_RepresentsSameFunctionalityAs       OpcuaNodeIdServicesReferenceType = 25260
	OpcuaNodeIdServicesReferenceType_IsHostedBy                          OpcuaNodeIdServicesReferenceType = 25261
	OpcuaNodeIdServicesReferenceType_HasPhysicalComponent                OpcuaNodeIdServicesReferenceType = 25262
	OpcuaNodeIdServicesReferenceType_HasContainedComponent               OpcuaNodeIdServicesReferenceType = 25263
	OpcuaNodeIdServicesReferenceType_HasAttachedComponent                OpcuaNodeIdServicesReferenceType = 25264
	OpcuaNodeIdServicesReferenceType_IsExecutingOn                       OpcuaNodeIdServicesReferenceType = 25265
	OpcuaNodeIdServicesReferenceType_HasPushedSecurityGroup              OpcuaNodeIdServicesReferenceType = 25345
	OpcuaNodeIdServicesReferenceType_AlarmSuppressionGroupMember         OpcuaNodeIdServicesReferenceType = 32059
	OpcuaNodeIdServicesReferenceType_HasKeyValueDescription              OpcuaNodeIdServicesReferenceType = 32407
	OpcuaNodeIdServicesReferenceType_HasEngineeringUnitDetails           OpcuaNodeIdServicesReferenceType = 32558
	OpcuaNodeIdServicesReferenceType_HasQuantity                         OpcuaNodeIdServicesReferenceType = 32559
	OpcuaNodeIdServicesReferenceType_HasCurrentData                      OpcuaNodeIdServicesReferenceType = 32633
	OpcuaNodeIdServicesReferenceType_HasCurrentEvent                     OpcuaNodeIdServicesReferenceType = 32634
	OpcuaNodeIdServicesReferenceType_HasReferenceDescription             OpcuaNodeIdServicesReferenceType = 32679
)

var OpcuaNodeIdServicesReferenceTypeValues []OpcuaNodeIdServicesReferenceType

func init() {
	_ = errors.New
	OpcuaNodeIdServicesReferenceTypeValues = []OpcuaNodeIdServicesReferenceType{
		OpcuaNodeIdServicesReferenceType_References,
		OpcuaNodeIdServicesReferenceType_NonHierarchicalReferences,
		OpcuaNodeIdServicesReferenceType_HierarchicalReferences,
		OpcuaNodeIdServicesReferenceType_HasChild,
		OpcuaNodeIdServicesReferenceType_Organizes,
		OpcuaNodeIdServicesReferenceType_HasEventSource,
		OpcuaNodeIdServicesReferenceType_HasModellingRule,
		OpcuaNodeIdServicesReferenceType_HasEncoding,
		OpcuaNodeIdServicesReferenceType_HasDescription,
		OpcuaNodeIdServicesReferenceType_HasTypeDefinition,
		OpcuaNodeIdServicesReferenceType_GeneratesEvent,
		OpcuaNodeIdServicesReferenceType_Aggregates,
		OpcuaNodeIdServicesReferenceType_HasSubtype,
		OpcuaNodeIdServicesReferenceType_HasProperty,
		OpcuaNodeIdServicesReferenceType_HasComponent,
		OpcuaNodeIdServicesReferenceType_HasNotifier,
		OpcuaNodeIdServicesReferenceType_HasOrderedComponent,
		OpcuaNodeIdServicesReferenceType_FromState,
		OpcuaNodeIdServicesReferenceType_ToState,
		OpcuaNodeIdServicesReferenceType_HasCause,
		OpcuaNodeIdServicesReferenceType_HasEffect,
		OpcuaNodeIdServicesReferenceType_HasHistoricalConfiguration,
		OpcuaNodeIdServicesReferenceType_HasSubStateMachine,
		OpcuaNodeIdServicesReferenceType_HasArgumentDescription,
		OpcuaNodeIdServicesReferenceType_HasOptionalInputArgumentDescription,
		OpcuaNodeIdServicesReferenceType_AlwaysGeneratesEvent,
		OpcuaNodeIdServicesReferenceType_HasTrueSubState,
		OpcuaNodeIdServicesReferenceType_HasFalseSubState,
		OpcuaNodeIdServicesReferenceType_HasCondition,
		OpcuaNodeIdServicesReferenceType_HasPubSubConnection,
		OpcuaNodeIdServicesReferenceType_DataSetToWriter,
		OpcuaNodeIdServicesReferenceType_HasGuard,
		OpcuaNodeIdServicesReferenceType_HasDataSetWriter,
		OpcuaNodeIdServicesReferenceType_HasDataSetReader,
		OpcuaNodeIdServicesReferenceType_HasAlarmSuppressionGroup,
		OpcuaNodeIdServicesReferenceType_AlarmGroupMember,
		OpcuaNodeIdServicesReferenceType_HasEffectDisable,
		OpcuaNodeIdServicesReferenceType_HasDictionaryEntry,
		OpcuaNodeIdServicesReferenceType_HasInterface,
		OpcuaNodeIdServicesReferenceType_HasAddIn,
		OpcuaNodeIdServicesReferenceType_HasEffectEnable,
		OpcuaNodeIdServicesReferenceType_HasEffectSuppressed,
		OpcuaNodeIdServicesReferenceType_HasEffectUnsuppressed,
		OpcuaNodeIdServicesReferenceType_HasWriterGroup,
		OpcuaNodeIdServicesReferenceType_HasReaderGroup,
		OpcuaNodeIdServicesReferenceType_AliasFor,
		OpcuaNodeIdServicesReferenceType_IsDeprecated,
		OpcuaNodeIdServicesReferenceType_HasStructuredComponent,
		OpcuaNodeIdServicesReferenceType_AssociatedWith,
		OpcuaNodeIdServicesReferenceType_UsesPriorityMappingTable,
		OpcuaNodeIdServicesReferenceType_HasLowerLayerInterface,
		OpcuaNodeIdServicesReferenceType_IsExecutableOn,
		OpcuaNodeIdServicesReferenceType_Controls,
		OpcuaNodeIdServicesReferenceType_Utilizes,
		OpcuaNodeIdServicesReferenceType_Requires,
		OpcuaNodeIdServicesReferenceType_IsPhysicallyConnectedTo,
		OpcuaNodeIdServicesReferenceType_RepresentsSameEntityAs,
		OpcuaNodeIdServicesReferenceType_RepresentsSameHardwareAs,
		OpcuaNodeIdServicesReferenceType_RepresentsSameFunctionalityAs,
		OpcuaNodeIdServicesReferenceType_IsHostedBy,
		OpcuaNodeIdServicesReferenceType_HasPhysicalComponent,
		OpcuaNodeIdServicesReferenceType_HasContainedComponent,
		OpcuaNodeIdServicesReferenceType_HasAttachedComponent,
		OpcuaNodeIdServicesReferenceType_IsExecutingOn,
		OpcuaNodeIdServicesReferenceType_HasPushedSecurityGroup,
		OpcuaNodeIdServicesReferenceType_AlarmSuppressionGroupMember,
		OpcuaNodeIdServicesReferenceType_HasKeyValueDescription,
		OpcuaNodeIdServicesReferenceType_HasEngineeringUnitDetails,
		OpcuaNodeIdServicesReferenceType_HasQuantity,
		OpcuaNodeIdServicesReferenceType_HasCurrentData,
		OpcuaNodeIdServicesReferenceType_HasCurrentEvent,
		OpcuaNodeIdServicesReferenceType_HasReferenceDescription,
	}
}

func OpcuaNodeIdServicesReferenceTypeByValue(value int32) (enum OpcuaNodeIdServicesReferenceType, ok bool) {
	switch value {
	case 117:
		return OpcuaNodeIdServicesReferenceType_HasSubStateMachine, true
	case 129:
		return OpcuaNodeIdServicesReferenceType_HasArgumentDescription, true
	case 131:
		return OpcuaNodeIdServicesReferenceType_HasOptionalInputArgumentDescription, true
	case 14476:
		return OpcuaNodeIdServicesReferenceType_HasPubSubConnection, true
	case 14936:
		return OpcuaNodeIdServicesReferenceType_DataSetToWriter, true
	case 15112:
		return OpcuaNodeIdServicesReferenceType_HasGuard, true
	case 15296:
		return OpcuaNodeIdServicesReferenceType_HasDataSetWriter, true
	case 15297:
		return OpcuaNodeIdServicesReferenceType_HasDataSetReader, true
	case 16361:
		return OpcuaNodeIdServicesReferenceType_HasAlarmSuppressionGroup, true
	case 16362:
		return OpcuaNodeIdServicesReferenceType_AlarmGroupMember, true
	case 17276:
		return OpcuaNodeIdServicesReferenceType_HasEffectDisable, true
	case 17597:
		return OpcuaNodeIdServicesReferenceType_HasDictionaryEntry, true
	case 17603:
		return OpcuaNodeIdServicesReferenceType_HasInterface, true
	case 17604:
		return OpcuaNodeIdServicesReferenceType_HasAddIn, true
	case 17983:
		return OpcuaNodeIdServicesReferenceType_HasEffectEnable, true
	case 17984:
		return OpcuaNodeIdServicesReferenceType_HasEffectSuppressed, true
	case 17985:
		return OpcuaNodeIdServicesReferenceType_HasEffectUnsuppressed, true
	case 18804:
		return OpcuaNodeIdServicesReferenceType_HasWriterGroup, true
	case 18805:
		return OpcuaNodeIdServicesReferenceType_HasReaderGroup, true
	case 23469:
		return OpcuaNodeIdServicesReferenceType_AliasFor, true
	case 23562:
		return OpcuaNodeIdServicesReferenceType_IsDeprecated, true
	case 24136:
		return OpcuaNodeIdServicesReferenceType_HasStructuredComponent, true
	case 24137:
		return OpcuaNodeIdServicesReferenceType_AssociatedWith, true
	case 25237:
		return OpcuaNodeIdServicesReferenceType_UsesPriorityMappingTable, true
	case 25238:
		return OpcuaNodeIdServicesReferenceType_HasLowerLayerInterface, true
	case 25253:
		return OpcuaNodeIdServicesReferenceType_IsExecutableOn, true
	case 25254:
		return OpcuaNodeIdServicesReferenceType_Controls, true
	case 25255:
		return OpcuaNodeIdServicesReferenceType_Utilizes, true
	case 25256:
		return OpcuaNodeIdServicesReferenceType_Requires, true
	case 25257:
		return OpcuaNodeIdServicesReferenceType_IsPhysicallyConnectedTo, true
	case 25258:
		return OpcuaNodeIdServicesReferenceType_RepresentsSameEntityAs, true
	case 25259:
		return OpcuaNodeIdServicesReferenceType_RepresentsSameHardwareAs, true
	case 25260:
		return OpcuaNodeIdServicesReferenceType_RepresentsSameFunctionalityAs, true
	case 25261:
		return OpcuaNodeIdServicesReferenceType_IsHostedBy, true
	case 25262:
		return OpcuaNodeIdServicesReferenceType_HasPhysicalComponent, true
	case 25263:
		return OpcuaNodeIdServicesReferenceType_HasContainedComponent, true
	case 25264:
		return OpcuaNodeIdServicesReferenceType_HasAttachedComponent, true
	case 25265:
		return OpcuaNodeIdServicesReferenceType_IsExecutingOn, true
	case 25345:
		return OpcuaNodeIdServicesReferenceType_HasPushedSecurityGroup, true
	case 3065:
		return OpcuaNodeIdServicesReferenceType_AlwaysGeneratesEvent, true
	case 31:
		return OpcuaNodeIdServicesReferenceType_References, true
	case 32:
		return OpcuaNodeIdServicesReferenceType_NonHierarchicalReferences, true
	case 32059:
		return OpcuaNodeIdServicesReferenceType_AlarmSuppressionGroupMember, true
	case 32407:
		return OpcuaNodeIdServicesReferenceType_HasKeyValueDescription, true
	case 32558:
		return OpcuaNodeIdServicesReferenceType_HasEngineeringUnitDetails, true
	case 32559:
		return OpcuaNodeIdServicesReferenceType_HasQuantity, true
	case 32633:
		return OpcuaNodeIdServicesReferenceType_HasCurrentData, true
	case 32634:
		return OpcuaNodeIdServicesReferenceType_HasCurrentEvent, true
	case 32679:
		return OpcuaNodeIdServicesReferenceType_HasReferenceDescription, true
	case 33:
		return OpcuaNodeIdServicesReferenceType_HierarchicalReferences, true
	case 34:
		return OpcuaNodeIdServicesReferenceType_HasChild, true
	case 35:
		return OpcuaNodeIdServicesReferenceType_Organizes, true
	case 36:
		return OpcuaNodeIdServicesReferenceType_HasEventSource, true
	case 37:
		return OpcuaNodeIdServicesReferenceType_HasModellingRule, true
	case 38:
		return OpcuaNodeIdServicesReferenceType_HasEncoding, true
	case 39:
		return OpcuaNodeIdServicesReferenceType_HasDescription, true
	case 40:
		return OpcuaNodeIdServicesReferenceType_HasTypeDefinition, true
	case 41:
		return OpcuaNodeIdServicesReferenceType_GeneratesEvent, true
	case 44:
		return OpcuaNodeIdServicesReferenceType_Aggregates, true
	case 45:
		return OpcuaNodeIdServicesReferenceType_HasSubtype, true
	case 46:
		return OpcuaNodeIdServicesReferenceType_HasProperty, true
	case 47:
		return OpcuaNodeIdServicesReferenceType_HasComponent, true
	case 48:
		return OpcuaNodeIdServicesReferenceType_HasNotifier, true
	case 49:
		return OpcuaNodeIdServicesReferenceType_HasOrderedComponent, true
	case 51:
		return OpcuaNodeIdServicesReferenceType_FromState, true
	case 52:
		return OpcuaNodeIdServicesReferenceType_ToState, true
	case 53:
		return OpcuaNodeIdServicesReferenceType_HasCause, true
	case 54:
		return OpcuaNodeIdServicesReferenceType_HasEffect, true
	case 56:
		return OpcuaNodeIdServicesReferenceType_HasHistoricalConfiguration, true
	case 9004:
		return OpcuaNodeIdServicesReferenceType_HasTrueSubState, true
	case 9005:
		return OpcuaNodeIdServicesReferenceType_HasFalseSubState, true
	case 9006:
		return OpcuaNodeIdServicesReferenceType_HasCondition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesReferenceTypeByName(value string) (enum OpcuaNodeIdServicesReferenceType, ok bool) {
	switch value {
	case "HasSubStateMachine":
		return OpcuaNodeIdServicesReferenceType_HasSubStateMachine, true
	case "HasArgumentDescription":
		return OpcuaNodeIdServicesReferenceType_HasArgumentDescription, true
	case "HasOptionalInputArgumentDescription":
		return OpcuaNodeIdServicesReferenceType_HasOptionalInputArgumentDescription, true
	case "HasPubSubConnection":
		return OpcuaNodeIdServicesReferenceType_HasPubSubConnection, true
	case "DataSetToWriter":
		return OpcuaNodeIdServicesReferenceType_DataSetToWriter, true
	case "HasGuard":
		return OpcuaNodeIdServicesReferenceType_HasGuard, true
	case "HasDataSetWriter":
		return OpcuaNodeIdServicesReferenceType_HasDataSetWriter, true
	case "HasDataSetReader":
		return OpcuaNodeIdServicesReferenceType_HasDataSetReader, true
	case "HasAlarmSuppressionGroup":
		return OpcuaNodeIdServicesReferenceType_HasAlarmSuppressionGroup, true
	case "AlarmGroupMember":
		return OpcuaNodeIdServicesReferenceType_AlarmGroupMember, true
	case "HasEffectDisable":
		return OpcuaNodeIdServicesReferenceType_HasEffectDisable, true
	case "HasDictionaryEntry":
		return OpcuaNodeIdServicesReferenceType_HasDictionaryEntry, true
	case "HasInterface":
		return OpcuaNodeIdServicesReferenceType_HasInterface, true
	case "HasAddIn":
		return OpcuaNodeIdServicesReferenceType_HasAddIn, true
	case "HasEffectEnable":
		return OpcuaNodeIdServicesReferenceType_HasEffectEnable, true
	case "HasEffectSuppressed":
		return OpcuaNodeIdServicesReferenceType_HasEffectSuppressed, true
	case "HasEffectUnsuppressed":
		return OpcuaNodeIdServicesReferenceType_HasEffectUnsuppressed, true
	case "HasWriterGroup":
		return OpcuaNodeIdServicesReferenceType_HasWriterGroup, true
	case "HasReaderGroup":
		return OpcuaNodeIdServicesReferenceType_HasReaderGroup, true
	case "AliasFor":
		return OpcuaNodeIdServicesReferenceType_AliasFor, true
	case "IsDeprecated":
		return OpcuaNodeIdServicesReferenceType_IsDeprecated, true
	case "HasStructuredComponent":
		return OpcuaNodeIdServicesReferenceType_HasStructuredComponent, true
	case "AssociatedWith":
		return OpcuaNodeIdServicesReferenceType_AssociatedWith, true
	case "UsesPriorityMappingTable":
		return OpcuaNodeIdServicesReferenceType_UsesPriorityMappingTable, true
	case "HasLowerLayerInterface":
		return OpcuaNodeIdServicesReferenceType_HasLowerLayerInterface, true
	case "IsExecutableOn":
		return OpcuaNodeIdServicesReferenceType_IsExecutableOn, true
	case "Controls":
		return OpcuaNodeIdServicesReferenceType_Controls, true
	case "Utilizes":
		return OpcuaNodeIdServicesReferenceType_Utilizes, true
	case "Requires":
		return OpcuaNodeIdServicesReferenceType_Requires, true
	case "IsPhysicallyConnectedTo":
		return OpcuaNodeIdServicesReferenceType_IsPhysicallyConnectedTo, true
	case "RepresentsSameEntityAs":
		return OpcuaNodeIdServicesReferenceType_RepresentsSameEntityAs, true
	case "RepresentsSameHardwareAs":
		return OpcuaNodeIdServicesReferenceType_RepresentsSameHardwareAs, true
	case "RepresentsSameFunctionalityAs":
		return OpcuaNodeIdServicesReferenceType_RepresentsSameFunctionalityAs, true
	case "IsHostedBy":
		return OpcuaNodeIdServicesReferenceType_IsHostedBy, true
	case "HasPhysicalComponent":
		return OpcuaNodeIdServicesReferenceType_HasPhysicalComponent, true
	case "HasContainedComponent":
		return OpcuaNodeIdServicesReferenceType_HasContainedComponent, true
	case "HasAttachedComponent":
		return OpcuaNodeIdServicesReferenceType_HasAttachedComponent, true
	case "IsExecutingOn":
		return OpcuaNodeIdServicesReferenceType_IsExecutingOn, true
	case "HasPushedSecurityGroup":
		return OpcuaNodeIdServicesReferenceType_HasPushedSecurityGroup, true
	case "AlwaysGeneratesEvent":
		return OpcuaNodeIdServicesReferenceType_AlwaysGeneratesEvent, true
	case "References":
		return OpcuaNodeIdServicesReferenceType_References, true
	case "NonHierarchicalReferences":
		return OpcuaNodeIdServicesReferenceType_NonHierarchicalReferences, true
	case "AlarmSuppressionGroupMember":
		return OpcuaNodeIdServicesReferenceType_AlarmSuppressionGroupMember, true
	case "HasKeyValueDescription":
		return OpcuaNodeIdServicesReferenceType_HasKeyValueDescription, true
	case "HasEngineeringUnitDetails":
		return OpcuaNodeIdServicesReferenceType_HasEngineeringUnitDetails, true
	case "HasQuantity":
		return OpcuaNodeIdServicesReferenceType_HasQuantity, true
	case "HasCurrentData":
		return OpcuaNodeIdServicesReferenceType_HasCurrentData, true
	case "HasCurrentEvent":
		return OpcuaNodeIdServicesReferenceType_HasCurrentEvent, true
	case "HasReferenceDescription":
		return OpcuaNodeIdServicesReferenceType_HasReferenceDescription, true
	case "HierarchicalReferences":
		return OpcuaNodeIdServicesReferenceType_HierarchicalReferences, true
	case "HasChild":
		return OpcuaNodeIdServicesReferenceType_HasChild, true
	case "Organizes":
		return OpcuaNodeIdServicesReferenceType_Organizes, true
	case "HasEventSource":
		return OpcuaNodeIdServicesReferenceType_HasEventSource, true
	case "HasModellingRule":
		return OpcuaNodeIdServicesReferenceType_HasModellingRule, true
	case "HasEncoding":
		return OpcuaNodeIdServicesReferenceType_HasEncoding, true
	case "HasDescription":
		return OpcuaNodeIdServicesReferenceType_HasDescription, true
	case "HasTypeDefinition":
		return OpcuaNodeIdServicesReferenceType_HasTypeDefinition, true
	case "GeneratesEvent":
		return OpcuaNodeIdServicesReferenceType_GeneratesEvent, true
	case "Aggregates":
		return OpcuaNodeIdServicesReferenceType_Aggregates, true
	case "HasSubtype":
		return OpcuaNodeIdServicesReferenceType_HasSubtype, true
	case "HasProperty":
		return OpcuaNodeIdServicesReferenceType_HasProperty, true
	case "HasComponent":
		return OpcuaNodeIdServicesReferenceType_HasComponent, true
	case "HasNotifier":
		return OpcuaNodeIdServicesReferenceType_HasNotifier, true
	case "HasOrderedComponent":
		return OpcuaNodeIdServicesReferenceType_HasOrderedComponent, true
	case "FromState":
		return OpcuaNodeIdServicesReferenceType_FromState, true
	case "ToState":
		return OpcuaNodeIdServicesReferenceType_ToState, true
	case "HasCause":
		return OpcuaNodeIdServicesReferenceType_HasCause, true
	case "HasEffect":
		return OpcuaNodeIdServicesReferenceType_HasEffect, true
	case "HasHistoricalConfiguration":
		return OpcuaNodeIdServicesReferenceType_HasHistoricalConfiguration, true
	case "HasTrueSubState":
		return OpcuaNodeIdServicesReferenceType_HasTrueSubState, true
	case "HasFalseSubState":
		return OpcuaNodeIdServicesReferenceType_HasFalseSubState, true
	case "HasCondition":
		return OpcuaNodeIdServicesReferenceType_HasCondition, true
	}
	return 0, false
}

func OpcuaNodeIdServicesReferenceTypeKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesReferenceTypeValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesReferenceType(structType any) OpcuaNodeIdServicesReferenceType {
	castFunc := func(typ any) OpcuaNodeIdServicesReferenceType {
		if sOpcuaNodeIdServicesReferenceType, ok := typ.(OpcuaNodeIdServicesReferenceType); ok {
			return sOpcuaNodeIdServicesReferenceType
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesReferenceType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesReferenceType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesReferenceTypeParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesReferenceType, error) {
	return OpcuaNodeIdServicesReferenceTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesReferenceTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesReferenceType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesReferenceType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesReferenceType")
	}
	if enum, ok := OpcuaNodeIdServicesReferenceTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesReferenceType")
		return OpcuaNodeIdServicesReferenceType(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesReferenceType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesReferenceType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesReferenceType", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesReferenceType) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesReferenceType_HasSubStateMachine:
		return "HasSubStateMachine"
	case OpcuaNodeIdServicesReferenceType_HasArgumentDescription:
		return "HasArgumentDescription"
	case OpcuaNodeIdServicesReferenceType_HasOptionalInputArgumentDescription:
		return "HasOptionalInputArgumentDescription"
	case OpcuaNodeIdServicesReferenceType_HasPubSubConnection:
		return "HasPubSubConnection"
	case OpcuaNodeIdServicesReferenceType_DataSetToWriter:
		return "DataSetToWriter"
	case OpcuaNodeIdServicesReferenceType_HasGuard:
		return "HasGuard"
	case OpcuaNodeIdServicesReferenceType_HasDataSetWriter:
		return "HasDataSetWriter"
	case OpcuaNodeIdServicesReferenceType_HasDataSetReader:
		return "HasDataSetReader"
	case OpcuaNodeIdServicesReferenceType_HasAlarmSuppressionGroup:
		return "HasAlarmSuppressionGroup"
	case OpcuaNodeIdServicesReferenceType_AlarmGroupMember:
		return "AlarmGroupMember"
	case OpcuaNodeIdServicesReferenceType_HasEffectDisable:
		return "HasEffectDisable"
	case OpcuaNodeIdServicesReferenceType_HasDictionaryEntry:
		return "HasDictionaryEntry"
	case OpcuaNodeIdServicesReferenceType_HasInterface:
		return "HasInterface"
	case OpcuaNodeIdServicesReferenceType_HasAddIn:
		return "HasAddIn"
	case OpcuaNodeIdServicesReferenceType_HasEffectEnable:
		return "HasEffectEnable"
	case OpcuaNodeIdServicesReferenceType_HasEffectSuppressed:
		return "HasEffectSuppressed"
	case OpcuaNodeIdServicesReferenceType_HasEffectUnsuppressed:
		return "HasEffectUnsuppressed"
	case OpcuaNodeIdServicesReferenceType_HasWriterGroup:
		return "HasWriterGroup"
	case OpcuaNodeIdServicesReferenceType_HasReaderGroup:
		return "HasReaderGroup"
	case OpcuaNodeIdServicesReferenceType_AliasFor:
		return "AliasFor"
	case OpcuaNodeIdServicesReferenceType_IsDeprecated:
		return "IsDeprecated"
	case OpcuaNodeIdServicesReferenceType_HasStructuredComponent:
		return "HasStructuredComponent"
	case OpcuaNodeIdServicesReferenceType_AssociatedWith:
		return "AssociatedWith"
	case OpcuaNodeIdServicesReferenceType_UsesPriorityMappingTable:
		return "UsesPriorityMappingTable"
	case OpcuaNodeIdServicesReferenceType_HasLowerLayerInterface:
		return "HasLowerLayerInterface"
	case OpcuaNodeIdServicesReferenceType_IsExecutableOn:
		return "IsExecutableOn"
	case OpcuaNodeIdServicesReferenceType_Controls:
		return "Controls"
	case OpcuaNodeIdServicesReferenceType_Utilizes:
		return "Utilizes"
	case OpcuaNodeIdServicesReferenceType_Requires:
		return "Requires"
	case OpcuaNodeIdServicesReferenceType_IsPhysicallyConnectedTo:
		return "IsPhysicallyConnectedTo"
	case OpcuaNodeIdServicesReferenceType_RepresentsSameEntityAs:
		return "RepresentsSameEntityAs"
	case OpcuaNodeIdServicesReferenceType_RepresentsSameHardwareAs:
		return "RepresentsSameHardwareAs"
	case OpcuaNodeIdServicesReferenceType_RepresentsSameFunctionalityAs:
		return "RepresentsSameFunctionalityAs"
	case OpcuaNodeIdServicesReferenceType_IsHostedBy:
		return "IsHostedBy"
	case OpcuaNodeIdServicesReferenceType_HasPhysicalComponent:
		return "HasPhysicalComponent"
	case OpcuaNodeIdServicesReferenceType_HasContainedComponent:
		return "HasContainedComponent"
	case OpcuaNodeIdServicesReferenceType_HasAttachedComponent:
		return "HasAttachedComponent"
	case OpcuaNodeIdServicesReferenceType_IsExecutingOn:
		return "IsExecutingOn"
	case OpcuaNodeIdServicesReferenceType_HasPushedSecurityGroup:
		return "HasPushedSecurityGroup"
	case OpcuaNodeIdServicesReferenceType_AlwaysGeneratesEvent:
		return "AlwaysGeneratesEvent"
	case OpcuaNodeIdServicesReferenceType_References:
		return "References"
	case OpcuaNodeIdServicesReferenceType_NonHierarchicalReferences:
		return "NonHierarchicalReferences"
	case OpcuaNodeIdServicesReferenceType_AlarmSuppressionGroupMember:
		return "AlarmSuppressionGroupMember"
	case OpcuaNodeIdServicesReferenceType_HasKeyValueDescription:
		return "HasKeyValueDescription"
	case OpcuaNodeIdServicesReferenceType_HasEngineeringUnitDetails:
		return "HasEngineeringUnitDetails"
	case OpcuaNodeIdServicesReferenceType_HasQuantity:
		return "HasQuantity"
	case OpcuaNodeIdServicesReferenceType_HasCurrentData:
		return "HasCurrentData"
	case OpcuaNodeIdServicesReferenceType_HasCurrentEvent:
		return "HasCurrentEvent"
	case OpcuaNodeIdServicesReferenceType_HasReferenceDescription:
		return "HasReferenceDescription"
	case OpcuaNodeIdServicesReferenceType_HierarchicalReferences:
		return "HierarchicalReferences"
	case OpcuaNodeIdServicesReferenceType_HasChild:
		return "HasChild"
	case OpcuaNodeIdServicesReferenceType_Organizes:
		return "Organizes"
	case OpcuaNodeIdServicesReferenceType_HasEventSource:
		return "HasEventSource"
	case OpcuaNodeIdServicesReferenceType_HasModellingRule:
		return "HasModellingRule"
	case OpcuaNodeIdServicesReferenceType_HasEncoding:
		return "HasEncoding"
	case OpcuaNodeIdServicesReferenceType_HasDescription:
		return "HasDescription"
	case OpcuaNodeIdServicesReferenceType_HasTypeDefinition:
		return "HasTypeDefinition"
	case OpcuaNodeIdServicesReferenceType_GeneratesEvent:
		return "GeneratesEvent"
	case OpcuaNodeIdServicesReferenceType_Aggregates:
		return "Aggregates"
	case OpcuaNodeIdServicesReferenceType_HasSubtype:
		return "HasSubtype"
	case OpcuaNodeIdServicesReferenceType_HasProperty:
		return "HasProperty"
	case OpcuaNodeIdServicesReferenceType_HasComponent:
		return "HasComponent"
	case OpcuaNodeIdServicesReferenceType_HasNotifier:
		return "HasNotifier"
	case OpcuaNodeIdServicesReferenceType_HasOrderedComponent:
		return "HasOrderedComponent"
	case OpcuaNodeIdServicesReferenceType_FromState:
		return "FromState"
	case OpcuaNodeIdServicesReferenceType_ToState:
		return "ToState"
	case OpcuaNodeIdServicesReferenceType_HasCause:
		return "HasCause"
	case OpcuaNodeIdServicesReferenceType_HasEffect:
		return "HasEffect"
	case OpcuaNodeIdServicesReferenceType_HasHistoricalConfiguration:
		return "HasHistoricalConfiguration"
	case OpcuaNodeIdServicesReferenceType_HasTrueSubState:
		return "HasTrueSubState"
	case OpcuaNodeIdServicesReferenceType_HasFalseSubState:
		return "HasFalseSubState"
	case OpcuaNodeIdServicesReferenceType_HasCondition:
		return "HasCondition"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesReferenceType) String() string {
	return e.PLC4XEnumName()
}
