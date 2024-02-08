package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseTearDown'
	TestInstructionUUID_SubCustody_TestCaseTearDown               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_TestCaseTearDown
	TestInstructionName_SubCustody_TestCaseTearDown               TypeAndStructs.TestInstructionNameType = "TestCaseTearDown"
	TestInstructionTypeUUID_SubCustody_TestCaseTearDown                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_GeneralSetUpTearDown
	TestInstructionTypeName_SubCustody_TestCaseTearDown                                                  = TestInstructions.TestInstructionTypeName_SubCustody_GeneralSetUpTearDown
	TestInstructionDescription_SubCustody_TestCaseTearDown        string                                 = "TestCaseTearDown, runs last for every TestCase for OnPrem-demo"
	TestInstructionMouseOverText_SubCustody_TestCaseTearDown      string                                 = "TestCaseTearDown, runs last for every TestCase for OnPrem-demo"
	TestInstructionDeprecated_SubCustody_TestCaseTearDown         bool                                   = false
	TestInstructionEnabled_SubCustody_TestCaseTearDown            bool                                   = false
	TestInstructionMajorVersionNumber_SubCustody_TestCaseTearDown int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_TestCaseTearDown int                                    = 0
	TestInstructionColor_SubCustody_TestCaseTearDown              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SubCustody_TestCaseTearDown                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_TestCaseTearDown                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                              TypeAndStructs.UpdatedTimeStampType    = "2023-11-28 14:00:00"

	// *** DropZone *** 'TestCaseTearDown_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "bdf1b383-a4c5-4fb2-a9a1-17a7bea75efb"
	TestInstructionDropZoneName_SubCustody_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseTearDown_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_TestCaseTearDown_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_TestCaseTearDown_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_TestCaseTearDown_ExpectsToSucceed       TypeAndStructs.ColorType        = "#000000AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "737b77d1-f468-4f07-acef-bdc087d71fe0" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SubCustody_TestCaseTearDown_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SubCustody_TestCaseTearDown_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_TestCaseTearDown_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	TestApiEngine_Class_Name_SubCustody_TestCaseTearDown = "GeneralSetupTearDown"
)

// TestInstruction_SubCustody_TestCaseTearDown
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_TestCaseTearDown *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_TestCaseTearDown
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_TestCaseTearDown() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	var testApiEngineMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct
	testApiEngineMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct)

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_SubCustody_TestCaseTearDown = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: 360,
				},
				TestApiEngineClassesMethodsAttributes: &TestApiEngineClassesAndMethods.TestApiEngineClassesMethodsAttributesStruct{
					TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_TestCaseTearDown,
					TestInstructionName:         TestInstructionName_SubCustody_TestCaseTearDown,
					TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethods.TestApiEngine_ClassName_UUID_SubCustody_GeneralSetupTearDown,
					TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_GeneralSetupTearDown,
					TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethods.TestApiEngine_MethodName_UUID_SubCustody_GeneralSetupTearDown_TearDown,
					TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethods.TestApiEngine_MethodName_Name_SubCustody_GeneralSetupTearDown_TearDown,
					Attributes:                  testApiEngineMethodAttributeMap,
				},
			},
		},
	}

	// TestInstruction - TestCaseTearDown
	TestInstruction_SubCustody_TestCaseTearDown.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SubCustody_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_SubCustody_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_SubCustody_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_TestCaseTearDown,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseTearDown
	TestInstruction_SubCustody_TestCaseTearDown.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SubCustody_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_SubCustody_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_TestCaseTearDown,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_TestCaseTearDown,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_TestCaseTearDown,
		TCRuleSwap:                   TCRuleSwap_SubCustody_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_SubCustody_TestCaseTearDown,
	}

	// DropZone 'TestCaseTearDown_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseTearDown_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_TestCaseTearDown_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_TestCaseTearDown_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_SubCustody_TestCaseTearDown,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_TestCaseTearDown_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_TestCaseTearDown_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_TestCaseTearDown_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_TestCaseTearDown_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_TestCaseTearDown_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_TestCaseTearDown,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_TestCaseTearDown.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_TestCaseTearDown.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_TestCaseTearDown_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_TestCaseTearDown_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_TestCaseTearDown_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_TestCaseTearDown,
		TestInstructionName:                           TestInstructionName_SubCustody_TestCaseTearDown,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_TestCaseTearDown_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_TestCaseTearDown.TestInstructionAttribute = append(
		TestInstruction_SubCustody_TestCaseTearDown.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_TestCaseTearDown_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SubCustody_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:   TestApiEngineClassesAndMethods.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:   TestApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
	}
	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed
	//TestInstruction_SubCustody_TestCaseTearDown.LocalExecutionMethods.TestApiEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SubCustody_TestCaseTearDown_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseTearDown
	var TestInstructionImmatureElementModel_SubCustody_TestCaseTearDown *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_TestCaseTearDown = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_TestCaseTearDown,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_TestCaseTearDown),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_TestCaseTearDown,
		NextElementUUID:          TestInstructionUUID_SubCustody_TestCaseTearDown,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_TestCaseTearDown,
		ParentElementUUID:        TestInstructionUUID_SubCustody_TestCaseTearDown,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_TestCaseTearDown,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_TestCaseTearDown,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_TestCaseTearDown.ImmatureElementModel = append(
		TestInstruction_SubCustody_TestCaseTearDown.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_TestCaseTearDown)

	return TestInstruction_SubCustody_TestCaseTearDown
}
