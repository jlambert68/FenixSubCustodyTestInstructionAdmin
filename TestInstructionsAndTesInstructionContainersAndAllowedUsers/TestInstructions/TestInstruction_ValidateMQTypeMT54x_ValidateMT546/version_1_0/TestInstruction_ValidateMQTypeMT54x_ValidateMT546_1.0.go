package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

const (

	// *************************************
	// *** TestInstruction *** 'ValidateMT546'
	TestInstructionUUID_SubCustody_ValidateMT546               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_ValidateMT546
	TestInstructionName_SubCustody_ValidateMT546               TypeAndStructs.TestInstructionNameType = "ValidateMT546"
	TestInstructionTypeUUID_SubCustody_ValidateMT546                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_VerifMQMessageTypeMT
	TestInstructionTypeName_SubCustody_ValidateMT546                                                  = TestInstructions.TestInstructionTypeName_SubCustody_VerifMQMessageTypeMT
	TestInstructionDescription_SubCustody_ValidateMT546        string                                 = "Validate a MT546 message received from MQ"
	TestInstructionMouseOverText_SubCustody_ValidateMT546      string                                 = "Validate a MT546 message received from MQ"
	TestInstructionDeprecated_SubCustody_ValidateMT546         bool                                   = false
	TestInstructionEnabled_SubCustody_ValidateMT546            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_ValidateMT546 int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_ValidateMT546 int                                    = 0
	TestInstructionColor_SubCustody_ValidateMT546              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_SubCustody_ValidateMT546                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_ValidateMT546                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                           TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"

	// *** DropZone *** 'ValidateMT546_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_ValidateMT546_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "f8ef8b0d-bf91-47b4-a5e5-b80d3a4542aa"
	TestInstructionDropZoneName_SubCustody_ValidateMT546_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "ValidateMT546_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_ValidateMT546_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_ValidateMT546_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_ValidateMT546_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_ValidateMT546_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "8f36c241-abfb-4a49-9a53-0617ed2167fd" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_ValidateMT546_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_ValidateMT546_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeDescription_SubCustody_ValidateMT546_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_ValidateMT546_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeActionCommand_SubCustody_ValidateMT546_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_ValidateMT546_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_ValidateMT546_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE

	// Attribute - 'RelatedReference_54x_20CRELA'
	TestInstructionAttributeUUID_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA          TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeName_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA          TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeType_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA          TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeDescription_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA   string                                          = "Extracts the response parameter from 54x_20CSEME that was sent on MQ"
	TestInstructionAttributeMouseOverText_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA string                                          = "Extracts the response parameter from 54x_20CSEME that was sent on MQ"
)

// TestInstruction_SubCustody_ValidateMT546
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_ValidateMT546 *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_ValidateMT546
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_ValidateMT546() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Generate Response variables for the TestInstruction
	Initate_TestInstructionResponseVariables_SubCustody_ValidateMT546()

	// Create and Initialize maps used in 'LocalExecutionMethods'
	var testApiEngineMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	testApiEngineMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct)

	var testInstructionsMap TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
	testInstructionsMap = make(TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType)

	var testApiEngineClassesMethodsAttributesVersionMap TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
	testApiEngineClassesMethodsAttributesVersionMap = make(TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType)

	// Create version as string
	var versionNumberAsString TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type // "1_0" or "13_3" ...)
	versionNumberAsString = TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type(
		strconv.Itoa(TestInstructionMajorVersionNumber_SubCustody_ValidateMT546) + "_" +
			strconv.Itoa(TestInstructionMinorVersionNumber_SubCustody_ValidateMT546))

	// Create 'TestApiEngineClassesMethodsAttributesVersionMap' for this TestInstruction-version
	testApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString] = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct{
		TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstructionName:         TestInstructionName_SubCustody_ValidateMT546,
		TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_VerifyMQMessageTypeMT,
		TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_VerifyMQMessageTypeMT,
		TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT546,
		TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT546,
		Attributes:                  &testApiEngineMethodAttributeMap,
	}

	// Create testInstructionsMap for this TestInstruction
	testInstructionsMap[TestInstructionUUID_SubCustody_ValidateMT546] = &testApiEngineClassesMethodsAttributesVersionMap

	// Initiate variable to store all TestInstruction data
	TestInstruction_SubCustody_ValidateMT546 = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: 300,
				},
				TestInstructionsMap: &testInstructionsMap},
		},
	}

	// TestInstruction - ValidateMT546
	TestInstruction_SubCustody_ValidateMT546.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstructionName:          TestInstructionName_SubCustody_ValidateMT546,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_ValidateMT546,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_ValidateMT546,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_ValidateMT546,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_ValidateMT546,
		Deprecated:                   TestInstructionDeprecated_SubCustody_ValidateMT546,
		Enabled:                      TestInstructionEnabled_SubCustody_ValidateMT546,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_ValidateMT546,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_ValidateMT546,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - ValidateMT546
	TestInstruction_SubCustody_ValidateMT546.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstructionName:          TestInstructionName_SubCustody_ValidateMT546,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_ValidateMT546,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_ValidateMT546,
		Deprecated:                   TestInstructionDeprecated_SubCustody_ValidateMT546,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_ValidateMT546,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_ValidateMT546,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_ValidateMT546,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_ValidateMT546,
		TCRuleSwap:                   TCRuleSwap_SubCustody_ValidateMT546,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_ValidateMT546,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_ValidateMT546,
		Enabled:                      TestInstructionEnabled_SubCustody_ValidateMT546,
	}

	// DropZone 'ValidateMT546_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: ValidateMT546_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_ValidateMT546_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_ValidateMT546_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstructionName:          TestInstructionName_SubCustody_ValidateMT546,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_ValidateMT546_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_ValidateMT546_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_ValidateMT546_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_ValidateMT546_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_ValidateMT546_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_ValidateMT546_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_ValidateMT546_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_ValidateMT546_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_ValidateMT546,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_ValidateMT546_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_ValidateMT546.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_ValidateMT546.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_ValidateMT546_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_ValidateMT546_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_ValidateMT546_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstructionName:                           TestInstructionName_SubCustody_ValidateMT546,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_ValidateMT546_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_ValidateMT546_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_ValidateMT546.TestInstructionAttribute = append(
		TestInstruction_SubCustody_ValidateMT546.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_ValidateMT546_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_SubCustody_ValidateMT546_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:   TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:   TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
	}
	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_ValidateMT546_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed
	//TestInstruction_SubCustody_ValidateMT546.LocalExecutionMethods.TestInstructionsMap.Attributes[TestInstructionAttributeUUID_SubCustody_ValidateMT546_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// TestInstruction Attribute - 'RelatedReference_54x_20CRELA'
	var TestInstructionAttribute_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_ValidateMT546,
		TestInstructionName:                           TestInstructionName_SubCustody_ValidateMT546,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
	}
	TestInstruction_SubCustody_ValidateMT546.TestInstructionAttribute = append(
		TestInstruction_SubCustody_ValidateMT546.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA)

	// Add TestApiEngine relation for Attribute - 'RelatedReference_54x_20CRELA'
	var tempTestApiEngineAttributeRelatedReference_54x_20CRELA *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeRelatedReference_54x_20CRELA = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
		TestInstructionAttributeName:     TestInstructionAttributeName_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:   TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA,
		TestApiEngineAttributeNameName:   TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA,
	}
	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA] = tempTestApiEngineAttributeRelatedReference_54x_20CRELA
	//TestInstruction_SubCustody_ValidateMT546.LocalExecutionMethods.TestInstructionsMap.Attributes[TestInstructionAttributeUUID_SubCustody_ValidateMT546_RelatedReference_54x_20CRELA] = tempTestApiEngineAttributeRelatedReference_54x_20CRELA

	// ImmatureElementModel - ValidateMT546
	var TestInstructionImmatureElementModel_SubCustody_ValidateMT546 *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_ValidateMT546 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_ValidateMT546,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_ValidateMT546),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_ValidateMT546,
		NextElementUUID:          TestInstructionUUID_SubCustody_ValidateMT546,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_ValidateMT546,
		ParentElementUUID:        TestInstructionUUID_SubCustody_ValidateMT546,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_ValidateMT546,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_ValidateMT546,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_ValidateMT546.ImmatureElementModel = append(
		TestInstruction_SubCustody_ValidateMT546.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_ValidateMT546)

	return TestInstruction_SubCustody_ValidateMT546
}
