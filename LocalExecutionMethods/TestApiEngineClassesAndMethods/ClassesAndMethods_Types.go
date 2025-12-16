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
