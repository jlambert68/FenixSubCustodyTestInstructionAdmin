package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'IsServerAlive'
	TestInstructionUUID_SubCustody_IsServerAlive               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_IsServerAlive
	TestInstructionName_SubCustody_IsServerAlive               TypeAndStructs.TestInstructionNameType = "IsServerAlive"
	TestInstructionTypeUUID_SubCustody_IsServerAlive                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_Standard
	TestInstructionTypeName_SubCustody_IsServerAlive                                                  = TestInstructions.TestInstructionTypeName_SubCustody_Standard
	TestInstructionDescription_SubCustody_IsServerAlive        string                                 = "This TestInstruction Checks if server 'X' is alive and responding "
	TestInstructionMouseOverText_SubCustody_IsServerAlive      string                                 = "This TestInstruction Checks if server 'X' is alive and responding"
	TestInstructionDeprecated_SubCustody_IsServerAlive         bool                                   = false
	TestInstructionEnabled_SubCustody_IsServerAlive            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_IsServerAlive int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_IsServerAlive int                                    = 0
	TestInstructionColor_SubCustody_IsServerAlive              TypeAndStructs.ColorType               = "#fff000AA"
	TCRuleDeletion_SubCustody_IsServerAlive                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_IsServerAlive                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                           TypeAndStructs.UpdatedTimeStampType    = "2024-01-15 20:00:00"

	// *** DropZone *** 'IsServerAlive_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_IsServerAlive_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "280b8ccd-3ed8-4c88-985a-3190da72016d"
	TestInstructionDropZoneName_SubCustody_IsServerAlive_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "IsServerAlive_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_IsServerAlive_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_IsServerAlive_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_IsServerAlive_ExpectsToSucceed       TypeAndStructs.ColorType        = "#ffff00AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_IsServerAlive_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_IsServerAlive_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_IsServerAlive_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SubCustody_IsServerAlive_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_IsServerAlive_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_IsServerAlive_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SubCustody_IsServerAlive_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_IsServerAlive_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'ServerServerIpAddress'
	TestInstructionAttributeUUID_SubCustody_IsServerAlive_ServerIpAddress          TypeAndStructs.TestInstructionAttributeUUIDType = "7cd1932e-1842-48ee-b4f7-c5782d6e7453" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SubCustody_ServerIpAddress
	TestInstructionAttributeName_SubCustody_IsServerAlive_ServerIpAddress          TypeAndStructs.TestInstructionAttributeNameType = "The IP-address of the server"
	TestInstructionAttributeType_SubCustody_IsServerAlive_ServerIpAddress          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_SubCustody_IsServerAlive_ServerIpAddress   string                                          = "The IP-address to the server to be checked"
	TestInstructionAttributeMouseOverText_SubCustody_IsServerAlive_ServerIpAddress string                                          = "The IP-address to the server to be checked"
)

// TestInstruction_SubCustody_IsServerAlive
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_IsServerAlive *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_IsServerAlive
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_IsServerAlive() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_SubCustody_IsServerAlive = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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
				TestApiEngineClassesMethodsAttributes: &TestApiEngineClassesAndMethods.TestApiEngineClassesMethodsAttributesStruct{},
			},
		},
	}

	// TestInstruction - IsServerAlive
	TestInstruction_SubCustody_IsServerAlive.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_IsServerAlive,
		TestInstructionName:          TestInstructionName_SubCustody_IsServerAlive,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_IsServerAlive,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_IsServerAlive,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_IsServerAlive,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_IsServerAlive,
		Deprecated:                   TestInstructionDeprecated_SubCustody_IsServerAlive,
		Enabled:                      TestInstructionEnabled_SubCustody_IsServerAlive,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_IsServerAlive,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_IsServerAlive,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - IsServerAlive
	TestInstruction_SubCustody_IsServerAlive.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_IsServerAlive,
		TestInstructionName:          TestInstructionName_SubCustody_IsServerAlive,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_IsServerAlive,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_IsServerAlive,
		Deprecated:                   TestInstructionDeprecated_SubCustody_IsServerAlive,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_IsServerAlive,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_IsServerAlive,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_IsServerAlive,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_IsServerAlive,
		TCRuleSwap:                   TCRuleSwap_SubCustody_IsServerAlive,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_IsServerAlive,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_IsServerAlive,
		Enabled:                      TestInstructionEnabled_SubCustody_IsServerAlive,
	}

	// DropZone 'IsServerAlive_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: IsServerAlive_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_IsServerAlive_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_IsServerAlive_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_IsServerAlive,
		TestInstructionName:          TestInstructionName_SubCustody_IsServerAlive,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_IsServerAlive_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_IsServerAlive_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_IsServerAlive_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_IsServerAlive_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_IsServerAlive_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_IsServerAlive_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_IsServerAlive_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_IsServerAlive_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_IsServerAlive,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_IsServerAlive_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_IsServerAlive.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_IsServerAlive.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_IsServerAlive_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_IsServerAlive_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_IsServerAlive_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_IsServerAlive,
		TestInstructionName:                           TestInstructionName_SubCustody_IsServerAlive,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_IsServerAlive_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_IsServerAlive_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_IsServerAlive.TestInstructionAttribute = append(
		TestInstruction_SubCustody_IsServerAlive.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_IsServerAlive_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// TestInstruction Attribute - 'ServerIpAddress'
	var TestInstructionAttribute_SubCustody_IsServerAlive_ServerIpAddress *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_IsServerAlive_ServerIpAddress = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_IsServerAlive,
		TestInstructionName:                           TestInstructionName_SubCustody_IsServerAlive,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_SubCustody_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_IsServerAlive_ServerIpAddress,
	}
	TestInstruction_SubCustody_IsServerAlive.TestInstructionAttribute = append(
		TestInstruction_SubCustody_IsServerAlive.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_IsServerAlive_ServerIpAddress)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// ImmatureElementModel - IsServerAlive
	var TestInstructionImmatureElementModel_SubCustody_IsServerAlive *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_IsServerAlive = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_IsServerAlive,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_IsServerAlive),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_IsServerAlive,
		NextElementUUID:          TestInstructionUUID_SubCustody_IsServerAlive,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_IsServerAlive,
		ParentElementUUID:        TestInstructionUUID_SubCustody_IsServerAlive,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_IsServerAlive,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_IsServerAlive,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_IsServerAlive.ImmatureElementModel = append(
		TestInstruction_SubCustody_IsServerAlive.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_IsServerAlive)

	return TestInstruction_SubCustody_IsServerAlive
}
