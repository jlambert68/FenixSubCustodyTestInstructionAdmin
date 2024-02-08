package TestApiEngineClassesAndMethods

import "github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"

// Types for TestApiEngine
type TestApiEngine_ClassName_UUID_SubCustody_Type string
type TestApiEngine_ClassName_Name_SubCustody_Type string
type TestApiEngine_MethodName_UUID_SubCustody_Type string
type TestApiEngine_MethodName_Name_SubCustody_Type string
type TestApiEngine_AttributeName_UUID_SubCustody_Type string
type TestApiEngine_AttributeName_Name_SubCustody_Type string

// Type this is used for specifying Classes, Methods and Attributes for TestApiEngine which is used by _SubCustodys TestAutomation
type TestApiEngineClassesMethodsAttributesStruct struct {
	TestInstructionOriginalUUID TypeAndStructs.OriginalElementUUIDType                                             `json:"TestInstructionOriginalUUID"`
	TestInstructionName         TypeAndStructs.TestInstructionNameType                                             `json:"TestInstructionName"`
	TestApiEngineClassNameUUID  TestApiEngine_ClassName_UUID_SubCustody_Type                                       `json:"TestApiEngineClassNameUUID"`
	TestApiEngineClassNameNAME  TestApiEngine_ClassName_Name_SubCustody_Type                                       `json:"TestApiEngineClassNameNAME"`
	TestApiEngineMethodNameUUID TestApiEngine_MethodName_UUID_SubCustody_Type                                      `json:"TestApiEngineMethodNameUUID"`
	TestApiEngineMethodNameNAME TestApiEngine_MethodName_Name_SubCustody_Type                                      `json:"TestApiEngineMethodNameNAME"`
	Attributes                  map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineAttributesStruct `json:"Attributes"`
}

type TestApiEngineAttributesStruct struct {
	TestInstructionAttributeUUID     TypeAndStructs.TestInstructionAttributeUUIDType     `json:"TestInstructionAttributeUUID"`
	TestInstructionAttributeName     TypeAndStructs.TestInstructionAttributeNameType     `json:"TestInstructionAttributeName"`
	TestInstructionAttributeTypeUUID TypeAndStructs.TestInstructionAttributeTypeUUIDType `json:"TestInstructionAttributeTypeUUID"`
	TestApiEngineAttributeNameUUID   TestApiEngine_AttributeName_UUID_SubCustody_Type    `json:"TestApiEngineAttributeNameUUID"`
	TestApiEngineAttributeNameName   TestApiEngine_AttributeName_Name_SubCustody_Type    `json:"TestApiEngineAttributeNameName"`
}

// Classes, Methods and their Parameters in TestApiEngine for -SubCustody
const (

	// General Attribute - ''
	TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_UUID_SubCustody_Type = "9edb935e-e683-4fee-bfca-f05faa0af595"
	TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_Name_SubCustody_Type = "expectedToBePassed"

	// ClassName - ***** 'SendOnMQTypeMT' *****
	TestApiEngine_ClassName_UUID_SubCustody_SendOnMQTypeMT TestApiEngine_ClassName_UUID_SubCustody_Type = "8baf20d0-3882-45f9-aa2e-a6eef24fbe21"
	TestApiEngine_ClassName_Name_SubCustody_SendOnMQTypeMT TestApiEngine_ClassName_Name_SubCustody_Type = "SendOnMQTypeMT"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMT540'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT540 TestApiEngine_MethodName_UUID_SubCustody_Type = "573c604d-62ad-4c2b-9bc8-57f46b3434a3"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT540 TestApiEngine_MethodName_Name_SubCustody_Type = "SendMT540"

	// ClassName: 'SendOnMQTypeMT' - MethodName: 'SendMT542'
	TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT542 TestApiEngine_MethodName_UUID_SubCustody_Type = "f6d8a81a-127e-4582-8aaf-e62c4b7c142a"
	TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT542 TestApiEngine_MethodName_Name_SubCustody_Type = "SendMT542"
)
