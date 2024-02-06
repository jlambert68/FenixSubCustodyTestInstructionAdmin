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

// Classes, Methods and their Parameters in TestApiEngine for _SubCustody
const (

	// General Attribute - ''
	TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_UUID_SubCustody_Type = "9b9e4ca8-e9a3-4939-b9dc-184b4e44f60e"
	TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed TestApiEngine_AttributeName_Name_SubCustody_Type = "expectedToBePassed"

	// ClassName - ***** 'GeneralSetupTearDown' *****
	TestApiEngine_ClassName_UUID_SubCustody_GeneralSetupTearDown TestApiEngine_ClassName_UUID_SubCustody_Type = "85373d2b-30ec-49ee-823a-0d8a0b2d5599"
	TestApiEngine_ClassName_Name_SubCustody_GeneralSetupTearDown TestApiEngine_ClassName_Name_SubCustody_Type = "GeneralSetupTearDown"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'Setup'
	TestApiEngine_MethodName_UUID_SubCustody_GeneralSetupTearDown_Setup TestApiEngine_MethodName_UUID_SubCustody_Type = "093cfe88-0970-427e-9548-82568bfede8c"
	TestApiEngine_MethodName_Name_SubCustody_GeneralSetupTearDown_Setup TestApiEngine_MethodName_Name_SubCustody_Type = "Setup"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'TearDown'
	TestApiEngine_MethodName_UUID_SubCustody_GeneralSetupTearDown_TearDown TestApiEngine_MethodName_UUID_SubCustody_Type = "0db4a61c-5a85-49a2-b039-4f411de0edd9"
	TestApiEngine_MethodName_Name_SubCustody_GeneralSetupTearDown_TearDown TestApiEngine_MethodName_Name_SubCustody_Type = "TearDown"
)
