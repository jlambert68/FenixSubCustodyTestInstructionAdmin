package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'SendMT542'
	TestInstructionUUID_SubCustody_SendMT542               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_SendMT542
	TestInstructionName_SubCustody_SendMT542               TypeAndStructs.TestInstructionNameType = "SendMT542"
	TestInstructionTypeUUID_SubCustody_SendMT542                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_SendOnMQTypeMT
	TestInstructionTypeName_SubCustody_SendMT542                                                  = TestInstructions.TestInstructionTypeName_SubCustody_SendOnMQTypeMT
	TestInstructionDescription_SubCustody_SendMT542        string                                 = "Sends a MT542 message on MQ"
	TestInstructionMouseOverText_SubCustody_SendMT542      string                                 = "Sends a MT542 message on MQ"
	TestInstructionDeprecated_SubCustody_SendMT542         bool                                   = false
	TestInstructionEnabled_SubCustody_SendMT542            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_SendMT542 int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_SendMT542 int                                    = 0
	TestInstructionColor_SubCustody_SendMT542              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SubCustody_SendMT542                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_SendMT542                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                       TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"

	// *** DropZone *** 'SendMT542_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_SendMT542_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "e9c6191d-1a3d-42ae-9d78-fd0e3a24069a"
	TestInstructionDropZoneName_SubCustody_SendMT542_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "SendMT542_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_SendMT542_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_SendMT542_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_SendMT542_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_SendMT542_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "11b48728-0228-43cc-8674-af2c72b2dfbf" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_SendMT542_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_SendMT542_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SubCustody_SendMT542_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_SendMT542_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_SendMT542_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SubCustody_SendMT542_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_SendMT542_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_SubCustody_SendMT542
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_SendMT542 *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_SendMT542
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_SendMT542() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	var testApiEngineMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct
	testApiEngineMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct)

	// Initiate variable to store all TestInstruction data
	TestInstruction_SubCustody_SendMT542 = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				TestApiEngineClassesMethodsAttributes: &TestApiEngineClassesAndMethods.TestApiEngineClassesMethodsAttributesStruct{
					TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_SendMT542,
					TestInstructionName:         TestInstructionName_SubCustody_SendMT542,
					TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethods.TestApiEngine_ClassName_UUID_SubCustody_SendOnMQTypeMT,
					TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_SendOnMQTypeMT,
					TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethods.TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT542,
					TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethods.TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT542,
					Attributes:                  testApiEngineMethodAttributeMap,
				},
			},
		},
	}

	// TestInstruction - SendMT542
	TestInstruction_SubCustody_SendMT542.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_SendMT542,
		TestInstructionName:          TestInstructionName_SubCustody_SendMT542,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_SendMT542,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_SendMT542,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_SendMT542,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_SendMT542,
		Deprecated:                   TestInstructionDeprecated_SubCustody_SendMT542,
		Enabled:                      TestInstructionEnabled_SubCustody_SendMT542,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_SendMT542,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_SendMT542,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - SendMT542
	TestInstruction_SubCustody_SendMT542.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_SendMT542,
		TestInstructionName:          TestInstructionName_SubCustody_SendMT542,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_SendMT542,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_SendMT542,
		Deprecated:                   TestInstructionDeprecated_SubCustody_SendMT542,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_SendMT542,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_SendMT542,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_SendMT542,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_SendMT542,
		TCRuleSwap:                   TCRuleSwap_SubCustody_SendMT542,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_SendMT542,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_SendMT542,
		Enabled:                      TestInstructionEnabled_SubCustody_SendMT542,
	}

	// DropZone 'SendMT542_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: SendMT542_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_SendMT542_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_SendMT542_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_SendMT542,
		TestInstructionName:          TestInstructionName_SubCustody_SendMT542,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_SendMT542_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_SendMT542_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_SendMT542_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_SendMT542_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_SendMT542_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_SendMT542_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_SendMT542_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_SendMT542_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_SendMT542,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_SendMT542_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_SendMT542.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_SendMT542.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_SendMT542_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_SendMT542_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_SendMT542_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_SendMT542,
		TestInstructionName:                           TestInstructionName_SubCustody_SendMT542,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_SubCustody_ExpectedToPass,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_SendMT542_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_SendMT542.TestInstructionAttribute = append(
		TestInstruction_SubCustody_SendMT542.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_SendMT542_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SubCustody_SendMT542_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:   TestApiEngineClassesAndMethods.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:   TestApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
	}
	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_SendMT542_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed
	//TestInstruction_SubCustody_SendMT542.LocalExecutionMethods.TestApiEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SubCustody_SendMT542_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// ImmatureElementModel - SendMT542
	var TestInstructionImmatureElementModel_SubCustody_SendMT542 *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_SendMT542 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_SendMT542,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_SendMT542),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_SendMT542,
		NextElementUUID:          TestInstructionUUID_SubCustody_SendMT542,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_SendMT542,
		ParentElementUUID:        TestInstructionUUID_SubCustody_SendMT542,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_SendMT542,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_SendMT542,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_SendMT542.ImmatureElementModel = append(
		TestInstruction_SubCustody_SendMT542.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_SendMT542)

	return TestInstruction_SubCustody_SendMT542
}
