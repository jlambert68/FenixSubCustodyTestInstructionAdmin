package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsDateAPublicHoliday"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'IsDateAPublicHoliday'
	TestInstructionUUID_SubCustody_IsDateAPublicHoliday               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_IsDateAPublicHoliday
	TestInstructionName_SubCustody_IsDateAPublicHoliday               TypeAndStructs.TestInstructionNameType = "IsDateAPublicHoliday"
	TestInstructionTypeUUID_SubCustody_IsDateAPublicHoliday                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_Standard
	TestInstructionTypeName_SubCustody_IsDateAPublicHoliday                                                  = TestInstructions.TestInstructionTypeName_SubCustody_Standard
	TestInstructionDescription_SubCustody_IsDateAPublicHoliday        string                                 = "This TestInstruction validates if a date is public holiday or not"
	TestInstructionMouseOverText_SubCustody_IsDateAPublicHoliday      string                                 = "This TestInstruction validates if a date is public holiday or not"
	TestInstructionDeprecated_SubCustody_IsDateAPublicHoliday         bool                                   = false
	TestInstructionEnabled_SubCustody_IsDateAPublicHoliday            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_IsDateAPublicHoliday int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_IsDateAPublicHoliday int                                    = 0
	TestInstructionColor_SubCustody_IsDateAPublicHoliday              TypeAndStructs.ColorType               = "#ffff00AA"
	TCRuleDeletion_SubCustody_IsDateAPublicHoliday                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_IsDateAPublicHoliday                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                                  TypeAndStructs.UpdatedTimeStampType    = "2024-01-08 15:00:00"

	// *** DropZone *** 'IsDateAPublicHoliday_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "dc713785-9bed-4e6e-aeab-964a5dbf648a"
	TestInstructionDropZoneName_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "IsDateAPublicHoliday_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed       TypeAndStructs.ColorType        = "#ffff00AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'HolidayDateToCheck'
	TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck          TypeAndStructs.TestInstructionAttributeUUIDType = "ebfeb406-75fc-4b5d-8e83-28d02d0d69d6" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SubCustody_HolidayDateToCheck
	TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck          TypeAndStructs.TestInstructionAttributeNameType = "Holiday Date To Check"
	TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck   string                                          = "The date that should be checked if it is a Public Holiday, use format 'YYYY-MM-DD'"
	TestInstructionAttributeMouseOverText_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck string                                          = "The date that should be checked if it is a Public Holiday, use format 'YYYY-MM-DD'"

	// Attribute - 'CountryCode'
	TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_CountryCode          TypeAndStructs.TestInstructionAttributeUUIDType = "371349ff-787d-4031-9e2e-4a937703cc0c" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SubCustody_CountryCode
	TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_CountryCode          TypeAndStructs.TestInstructionAttributeNameType = "Country Code"
	TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_CountryCode          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_SubCustody_IsDateAPublicHoliday_CountryCode   string                                          = "The Country for which to check the Public Holiday, use format 'XX'"
	TestInstructionAttributeMouseOverText_SubCustody_IsDateAPublicHoliday_CountryCode string                                          = "The Country for which to check the Public Holiday, use format 'XX'"
)

// TestInstruction_SubCustody_IsDateAPublicHoliday
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_IsDateAPublicHoliday *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_IsDateAPublicHoliday
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_IsDateAPublicHoliday() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_SubCustody_IsDateAPublicHoliday = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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

	// TestInstruction - IsDateAPublicHoliday
	TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_InGCP,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_InGCP,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionName:          TestInstructionName_SubCustody_IsDateAPublicHoliday,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_IsDateAPublicHoliday,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_IsDateAPublicHoliday,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_IsDateAPublicHoliday,
		Deprecated:                   TestInstructionDeprecated_SubCustody_IsDateAPublicHoliday,
		Enabled:                      TestInstructionEnabled_SubCustody_IsDateAPublicHoliday,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_IsDateAPublicHoliday,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_IsDateAPublicHoliday,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - IsDateAPublicHoliday
	TestInstruction_SubCustody_IsDateAPublicHoliday.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_InGCP,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_InGCP,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionName:          TestInstructionName_SubCustody_IsDateAPublicHoliday,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_IsDateAPublicHoliday,
		Deprecated:                   TestInstructionDeprecated_SubCustody_IsDateAPublicHoliday,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_IsDateAPublicHoliday,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_IsDateAPublicHoliday,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_IsDateAPublicHoliday,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_IsDateAPublicHoliday,
		TCRuleSwap:                   TCRuleSwap_SubCustody_IsDateAPublicHoliday,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_IsDateAPublicHoliday,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_IsDateAPublicHoliday,
		Enabled:                      TestInstructionEnabled_SubCustody_IsDateAPublicHoliday,
	}

	// DropZone 'IsDateAPublicHoliday_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: IsDateAPublicHoliday_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionName:          TestInstructionName_SubCustody_IsDateAPublicHoliday,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_IsDateAPublicHoliday_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_IsDateAPublicHoliday.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_IsDateAPublicHoliday.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionName:                           TestInstructionName_SubCustody_IsDateAPublicHoliday,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstructionAttribute = append(
		TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// TestInstruction Attribute - 'HolidayDateToCheck'
	var TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionName:                           TestInstructionName_SubCustody_IsDateAPublicHoliday,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck,
	}
	TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstructionAttribute = append(
		TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_HolidayDateToCheck)

	// TestInstruction Attribute - 'CountryCode'
	var TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_CountryCode *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_CountryCode = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestInstructionName:                           TestInstructionName_SubCustody_IsDateAPublicHoliday,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_IsDateAPublicHoliday_CountryCode,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_IsDateAPublicHoliday_CountryCode,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_IsDateAPublicHoliday_CountryCode,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_IsDateAPublicHoliday_CountryCode,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_IsDateAPublicHoliday_CountryCode,
	}
	TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstructionAttribute = append(
		TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_IsDateAPublicHoliday_CountryCode)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// ImmatureElementModel - IsDateAPublicHoliday
	var TestInstructionImmatureElementModel_SubCustody_IsDateAPublicHoliday *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_IsDateAPublicHoliday = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_IsDateAPublicHoliday),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		NextElementUUID:          TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		ParentElementUUID:        TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_IsDateAPublicHoliday,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_IsDateAPublicHoliday.ImmatureElementModel = append(
		TestInstruction_SubCustody_IsDateAPublicHoliday.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_IsDateAPublicHoliday)

	return TestInstruction_SubCustody_IsDateAPublicHoliday
}
