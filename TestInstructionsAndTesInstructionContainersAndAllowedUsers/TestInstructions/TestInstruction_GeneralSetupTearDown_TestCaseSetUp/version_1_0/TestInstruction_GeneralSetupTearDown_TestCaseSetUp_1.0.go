package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseSetUp'
	TestInstructionUUID_SubCustody_TestCaseSetUp               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_TestCaseSetUp
	TestInstructionName_SubCustody_TestCaseSetUp               TypeAndStructs.TestInstructionNameType = "TestCaseSetUp"
	TestInstructionTypeUUID_SubCustody_TestCaseSetUp                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_GeneralSetUpTearDown
	TestInstructionTypeName_SubCustody_TestCaseSetUp                                                  = TestInstructions.TestInstructionTypeName_SubCustody_GeneralSetUpTearDown
	TestInstructionDescription_SubCustody_TestCaseSetUp        string                                 = "Initiate _SubCustodys execution engine to be able to execute TestInstructionsMap"
	TestInstructionMouseOverText_SubCustody_TestCaseSetUp      string                                 = "Initiate _SubCustodys execution engine to be able to execute TestInstructionsMap"
	TestInstructionDeprecated_SubCustody_TestCaseSetUp         bool                                   = false
	TestInstructionEnabled_SubCustody_TestCaseSetUp            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_TestCaseSetUp int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_TestCaseSetUp int                                    = 0
	TestInstructionColor_SubCustody_TestCaseSetUp              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SubCustody_TestCaseSetUp                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_TestCaseSetUp                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                           TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"

	// *** DropZone *** 'TestCaseSetUp_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "f0576c45-bc12-44dc-beb8-0f115e029920"
	TestInstructionDropZoneName_SubCustody_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseSetUp_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_TestCaseSetUp_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_TestCaseSetUp_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_TestCaseSetUp_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "e945cf17-7f4f-49df-9800-6447a25637cb" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SubCustody_TestCaseSetUp_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SubCustody_TestCaseSetUp_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_TestCaseSetUp_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_SubCustody_TestCaseSetUp
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_TestCaseSetUp *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_TestCaseSetUp
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_TestCaseSetUp() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	var fangMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct
	fangMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct)

	// Initiate variable to store all TestInstruction data
	TestInstruction_SubCustody_TestCaseSetUp = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				TestApiEngineClassesMethodsAttributes: &TestApiEngineClassesAndMethods.TestApiEngineClassesMethodsAttributesStruct{
					TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_TestCaseSetUp,
					TestInstructionName:         TestInstructionName_SubCustody_TestCaseSetUp,
					TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethods.TestApiEngine_ClassName_UUID_SubCustody_GeneralSetupTearDown,
					TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_GeneralSetupTearDown,
					TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethods.TestApiEngine_MethodName_UUID_SubCustody_GeneralSetupTearDown_Setup,
					TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethods.TestApiEngine_MethodName_Name_SubCustody_GeneralSetupTearDown_Setup,
					Attributes:                  fangMethodAttributeMap,
				},
			},
		},
	}

	// TestInstruction - TestCaseSetUp
	TestInstruction_SubCustody_TestCaseSetUp.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SubCustody_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_SubCustody_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_SubCustody_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_TestCaseSetUp,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseSetUp
	TestInstruction_SubCustody_TestCaseSetUp.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SubCustody_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_SubCustody_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_TestCaseSetUp,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_TestCaseSetUp,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_TestCaseSetUp,
		TCRuleSwap:                   TCRuleSwap_SubCustody_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_SubCustody_TestCaseSetUp,
	}

	// DropZone 'TestCaseSetUp_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseSetUp_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_SubCustody_TestCaseSetUp,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_TestCaseSetUp_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_TestCaseSetUp_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_TestCaseSetUp_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_TestCaseSetUp_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_TestCaseSetUp_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_TestCaseSetUp,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_TestCaseSetUp.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_TestCaseSetUp_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_TestCaseSetUp,
		TestInstructionName:                           TestInstructionName_SubCustody_TestCaseSetUp,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_TestCaseSetUp_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_TestCaseSetUp.TestInstructionAttribute = append(
		TestInstruction_SubCustody_TestCaseSetUp.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_TestCaseSetUp_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethods.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SubCustody_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:   TestApiEngineClassesAndMethods.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:   TestApiEngineClassesAndMethods.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
	}
	fangMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed
	//TestInstruction_SubCustody_TestCaseSetUp.LocalExecutionMethods.TestApiEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_SubCustody_TestCaseSetUp_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseSetUp
	var TestInstructionImmatureElementModel_SubCustody_TestCaseSetUp *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_TestCaseSetUp = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_TestCaseSetUp,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_TestCaseSetUp),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_TestCaseSetUp,
		NextElementUUID:          TestInstructionUUID_SubCustody_TestCaseSetUp,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_TestCaseSetUp,
		ParentElementUUID:        TestInstructionUUID_SubCustody_TestCaseSetUp,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_TestCaseSetUp,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_TestCaseSetUp,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_TestCaseSetUp.ImmatureElementModel = append(
		TestInstruction_SubCustody_TestCaseSetUp.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_TestCaseSetUp)

	return TestInstruction_SubCustody_TestCaseSetUp
}
