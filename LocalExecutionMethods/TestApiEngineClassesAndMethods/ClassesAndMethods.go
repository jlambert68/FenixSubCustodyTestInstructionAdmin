package TestApiEngineClassesAndMethodsAndAttributes

import "github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"

// Types for TestApiEngine
type TestApiEngine_MethodNameVersion_SubCustody_Type string // "1_0 or 13_3" or similar

type TestApiEngine_ClassName_UUID_SubCustody_Type string
type TestApiEngine_ClassName_Name_SubCustody_Type string
type TestApiEngine_MethodName_UUID_SubCustody_Type string
type TestApiEngine_MethodName_Name_SubCustody_Type string
type TestApiEngine_AttributeName_UUID_SubCustody_Type string
type TestApiEngine_AttributeName_Name_SubCustody_Type string

// TestApiEngineClassesMethodsAttributesStruct
// Type this is used for specifying Classes, Methods and Attributes for TestApiEngine which is used by SubCustodys TestAutomation
type TestApiEngineClassesMethodsAttributesStruct struct {
	TestInstructionOriginalUUID TypeAndStructs.OriginalElementUUIDType                                              `json:"TestInstructionOriginalUUID"`
	TestInstructionName         TypeAndStructs.TestInstructionNameType                                              `json:"TestInstructionName"`
	TestApiEngineClassNameUUID  TestApiEngine_ClassName_UUID_SubCustody_Type                                        `json:"TestApiEngineClassNameUUID"`
	TestApiEngineClassNameNAME  TestApiEngine_ClassName_Name_SubCustody_Type                                        `json:"TestApiEngineClassNameNAME"`
	TestApiEngineMethodNameUUID TestApiEngine_MethodName_UUID_SubCustody_Type                                       `json:"TestApiEngineMethodNameUUID"`
	TestApiEngineMethodNameNAME TestApiEngine_MethodName_Name_SubCustody_Type                                       `json:"TestApiEngineMethodNameNAME"`
	Attributes                  *map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineAttributesStruct `json:"Attributes"`
}

// TestApiEngineAttributesStruct
// Type for Attributes
type TestApiEngineAttributesStruct struct {
	TestInstructionAttributeUUID         TypeAndStructs.TestInstructionAttributeUUIDType     `json:"TestInstructionAttributeUUID"`
	TestInstructionAttributeName         TypeAndStructs.TestInstructionAttributeNameType     `json:"TestInstructionAttributeName"`
	TestInstructionAttributeTypeUUID     TypeAndStructs.TestInstructionAttributeTypeUUIDType `json:"TestInstructionAttributeTypeUUID"`
	TestApiEngineAttributeNameUUID       TestApiEngine_AttributeName_UUID_SubCustody_Type    `json:"TestApiEngineAttributeNameUUID"`
	TestApiEngineAttributeNameName       TestApiEngine_AttributeName_Name_SubCustody_Type    `json:"TestApiEngineAttributeNameName"`
	AttributeShouldBeSentToTestApiEngine bool                                                `json:"AttributeShouldBeSentToTestApiEngine"`
}

// TestApiEngineClassesMethodsAttributesVersionMapType
// Type for map holding connecting Attributes for TestInstruction with attributes for TestApiEngine, for a specific version of the TestInstruction/TestApiEngine-call
type TestApiEngineClassesMethodsAttributesVersionMapType map[TestApiEngine_MethodNameVersion_SubCustody_Type]*TestApiEngineClassesMethodsAttributesStruct

// TestInstructionsMapType
// Type for map connecting Original UUID for TestInstruction with the all versions of the TestInstruction
type TestInstructionsMapType map[TypeAndStructs.OriginalElementUUIDType]*TestApiEngineClassesMethodsAttributesVersionMapType

// Classes, Methods and their Parameters in TestApiEngine for SubCustody
const (

	// General Attribute - ''
	TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_UUID_SubCustody_Type = "9edb935e-e683-4fee-bfca-f05faa0af595"
	TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_Name_SubCustody_Type = "expectedToBePassed"

	// ClassName - ***** 'SendOnMQTypeMT' *****
	TestApiEngine_ClassName_UUID_SubCustody_SendOnMQTypeMT TestApiEngine_ClassName_UUID_SubCustody_Type = "8baf20d0-3882-45f9-aa2e-a6eef24fbe21"
	TestApiEngine_ClassName_Name_SubCustody_SendOnMQTypeMT TestApiEngine_ClassName_Name_SubCustody_Type = "SendOnMQTypeMT"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMTGeneral'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMTGeneral TestApiEngine_MethodName_UUID_SubCustody_Type = "e59b252b-47a7-4a37-9059-75fc7dbdb355"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMTGeneral TestApiEngine_MethodName_Name_SubCustody_Type = "SendMTGeneral_v1_0"

	// ClassName:'SendOnMQTypeMT' - MethodName: 'SendMTGeneral' - Attributes
	TestApiEngine_AttributeName_UUID_SubCustody_SendOnMQTypeMT_SwiftMessageType   TestApiEngine_AttributeName_UUID_SubCustody_Type = "2dfba095-f4fa-411b-9f0e-a09142fa8b97"
	TestApiEngine_AttributeName_Name_SubCustody_SendOnMQTypeMT_SwiftMessageType   TestApiEngine_AttributeName_Name_SubCustody_Type = "SwiftMessageType"
	TestApiEngine_AttributeName_UUID_SubCustody_SendOnMQTypeMT_TemplateAsString   TestApiEngine_AttributeName_UUID_SubCustody_Type = "c3411ccd-91f4-4a13-ad7d-28967a24bef9"
	TestApiEngine_AttributeName_Name_SubCustody_SendOnMQTypeMT_TemplateAsString   TestApiEngine_AttributeName_Name_SubCustody_Type = "TemplateAsString"
	TestApiEngine_AttributeName_UUID_SubCustody_SendOnMQTypeMT_ReplacePlaceholers TestApiEngine_AttributeName_UUID_SubCustody_Type = "73b3cc20-19d2-43e8-b83a-bac71ac3f005"
	TestApiEngine_AttributeName_Name_SubCustody_SendOnMQTypeMT_ReplacePlaceholers TestApiEngine_AttributeName_Name_SubCustody_Type = "ReplacePlaceholders"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMT540'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT540 TestApiEngine_MethodName_UUID_SubCustody_Type = "573c604d-62ad-4c2b-9bc8-57f46b3434a3"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT540 TestApiEngine_MethodName_Name_SubCustody_Type = "SendMT540_v1_0"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMT542'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT542 TestApiEngine_MethodName_UUID_SubCustody_Type = "f6d8a81a-127e-4582-8aaf-e62c4b7c142a"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT542 TestApiEngine_MethodName_Name_SubCustody_Type = "SendMT542_v1_0"

	// ClassName - ***** 'ValidateMQTypeMT54x' *****
	TestApiEngine_ClassName_UUID_SubCustody_VerifyMQMessageTypeMT TestApiEngine_ClassName_UUID_SubCustody_Type = "6db25acd-1a2a-471d-a068-62ce01208747"
	TestApiEngine_ClassName_Name_SubCustody_VerifyMQMessageTypeMT TestApiEngine_ClassName_Name_SubCustody_Type = "ValidateMQTypeMT54x"

	// ClassName: 'ValidateMQTypeMT54x' - MethodName: 'ValidateMT544'
	TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT544 TestApiEngine_MethodName_UUID_SubCustody_Type = "09292093-751b-4b41-afd0-8e0a93d49e6a"
	TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT544 TestApiEngine_MethodName_Name_SubCustody_Type = "ValidateMT544_v1_0"

	// ClassName:'ValidateMQTypeMT54x' - MethodName: 'ValidateMT544' - Attributes
	TestApiEngine_AttributeName_UUID_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA TestApiEngine_AttributeName_UUID_SubCustody_Type = "9d99a553-23fe-4f02-942f-0d53ddc8acfb"
	TestApiEngine_AttributeName_Name_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA TestApiEngine_AttributeName_Name_SubCustody_Type = "RelatedReference_54x_20CRELA"

	// ClassName: 'ValidateMQTypeMT54x' - MethodName: 'ValidateMT546'
	TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT546 TestApiEngine_MethodName_UUID_SubCustody_Type = "db919745-f09b-4afa-9698-c0b9596fc67c"
	TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT546 TestApiEngine_MethodName_Name_SubCustody_Type = "ValidateMT546_v1_0"

	// ClassName: 'ValidateMQTypeMT54x' - MethodName: 'ValidateMT548'
	TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT548 TestApiEngine_MethodName_UUID_SubCustody_Type = "f7b0b0da-6b3a-40b8-bdbb-a032a6d8d056"
	TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT548 TestApiEngine_MethodName_Name_SubCustody_Type = "ValidateMT548_v1_0"
)
