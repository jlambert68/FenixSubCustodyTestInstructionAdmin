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

// Type this is used for specifying Classes, Methods and Attributes for TestApiEngine which is used by SubCustodys TestAutomation
type TestApiEngineClassesMethodsAttributesStruct struct {
	TestInstructionOriginalUUID TypeAndStructs.OriginalElementUUIDType                                             `json:"TestInstructionOriginalUUID"`
	TestInstructionName         TypeAndStructs.TestInstructionNameType                                             `json:"TestInstructionName"`
	TestApiEngineClassNameUUID  TestApiEngine_ClassName_UUID_SubCustody_Type                                       `json:"TestApiEngineClassNameUUID"`
	TestApiEngineClassNameNAME  TestApiEngine_ClassName_Name_SubCustody_Type                                       `json:"TestApiEngineClassNameNAME"`
	TestApiEngineMethodNameUUID TestApiEngine_MethodName_UUID_SubCustody_Type                                      `json:"TestApiEngineMethodNameUUID"`
	TestApiEngineMethodNameNAME TestApiEngine_MethodName_Name_SubCustody_Type                                      `json:"TestApiEngineMethodNameNAME"`
	Attributes                  map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineAttributesStruct `json:"Attributes"`
}

// Type for Attributes
type TestApiEngineAttributesStruct struct {
	TestInstructionAttributeUUID     TypeAndStructs.TestInstructionAttributeUUIDType     `json:"TestInstructionAttributeUUID"`
	TestInstructionAttributeName     TypeAndStructs.TestInstructionAttributeNameType     `json:"TestInstructionAttributeName"`
	TestInstructionAttributeTypeUUID TypeAndStructs.TestInstructionAttributeTypeUUIDType `json:"TestInstructionAttributeTypeUUID"`
	TestApiEngineAttributeNameUUID   TestApiEngine_AttributeName_UUID_SubCustody_Type    `json:"TestApiEngineAttributeNameUUID"`
	TestApiEngineAttributeNameName   TestApiEngine_AttributeName_Name_SubCustody_Type    `json:"TestApiEngineAttributeNameName"`
}

// TestApiEngineClassesMethodsAttributesType
// Type for map holding connecting Attributes for TestInstruction with attributes for TestApiEngine, for a specific version of the TestInstruction/TestApiEngine-call
type TestApiEngineClassesMethodsAttributesType map[TestApiEngine_MethodNameVersion_SubCustody_Type]*TestApiEngineClassesMethodsAttributesStruct

// TestInstructionVersionMapType
// Type for map connecting Original UUID for TestInstruction with the all versions of the TestInstruction
type TestInstructionVersionMapType map[TypeAndStructs.OriginalElementUUIDType]TestApiEngineClassesMethodsAttributesType

// Classes, Methods and their Parameters in TestApiEngine for SubCustody
const (

	// General Attribute - ''
	TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_UUID_SubCustody_Type = "9edb935e-e683-4fee-bfca-f05faa0af595"
	TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_Name_SubCustody_Type = "expectedToBePassed"

	// ClassName - ***** 'SendOnMQTypeMT' *****
	TestApiEngine_ClassName_UUID_SubCustody_SendOnMQTypeMT TestApiEngine_ClassName_UUID_SubCustody_Type = "8baf20d0-3882-45f9-aa2e-a6eef24fbe21"
	TestApiEngine_ClassName_Name_SubCustody_SendOnMQTypeMT TestApiEngine_ClassName_Name_SubCustody_Type = "SendOnMQTypeMT"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMT540'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT540 TestApiEngine_MethodName_UUID_SubCustody_Type = "573c604d-62ad-4c2b-9bc8-57f46b3434a3"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT540 TestApiEngine_MethodName_Name_SubCustody_Type = "SendMT540_v1_0"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMT542'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT542 TestApiEngine_MethodName_UUID_SubCustody_Type = "f6d8a81a-127e-4582-8aaf-e62c4b7c142a"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT542 TestApiEngine_MethodName_Name_SubCustody_Type = "SendMT542_v1_0"

	// ClassName - ***** 'VerifyMQMessageTypeMT' *****
	TestApiEngine_ClassName_UUID_SubCustody_VerifyMQMessageTypeMT TestApiEngine_ClassName_UUID_SubCustody_Type = "6db25acd-1a2a-471d-a068-62ce01208747"
	TestApiEngine_ClassName_Name_SubCustody_VerifyMQMessageTypeMT TestApiEngine_ClassName_Name_SubCustody_Type = "VerifyMQMessageTypeMT"

	// ClassName: 'VerifyMQMessageTypeMT' - MethodName: 'VerifyMT544'
	TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT544 TestApiEngine_MethodName_UUID_SubCustody_Type = "09292093-751b-4b41-afd0-8e0a93d49e6a"
	TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT544 TestApiEngine_MethodName_Name_SubCustody_Type = "VerifyMT544_v1_0"

	// ClassName:'VerifyMQMessageTypeMT' - MethodName: 'VerifyMT544' - Attributes
	FangEngine_AttributeName_UUID_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA TestApiEngine_AttributeName_UUID_SubCustody_Type = "9d99a553-23fe-4f02-942f-0d53ddc8acfb"
	FangEngine_AttributeName_Name_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA TestApiEngine_AttributeName_Name_SubCustody_Type = "RelatedReference_54x_20CRELA"

	// ClassName: 'VerifyMQMessageTypeMT' - MethodName: 'VerifyMT546'
	TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT546 TestApiEngine_MethodName_UUID_SubCustody_Type = "db919745-f09b-4afa-9698-c0b9596fc67c"
	TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT546 TestApiEngine_MethodName_Name_SubCustody_Type = "VerifyMT546_v1_0"

	// ClassName: 'VerifyMQMessageTypeMT' - MethodName: 'VerifyMT548'
	TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT548 TestApiEngine_MethodName_UUID_SubCustody_Type = "f7b0b0da-6b3a-40b8-bdbb-a032a6d8d056"
	TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT548 TestApiEngine_MethodName_Name_SubCustody_Type = "VerifyMT548_v1_0"
)
