package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

const (

	// *************************************
	// *** TestInstruction *** 'ValidateMT544'
	TestInstructionUUID_SubCustody_ValidateMT544               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_ValidateMT544
	TestInstructionName_SubCustody_ValidateMT544               TypeAndStructs.TestInstructionNameType = "ValidateMT544"
	TestInstructionTypeUUID_SubCustody_ValidateMT544                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_VerifMQMessageTypeMT
	TestInstructionTypeName_SubCustody_ValidateMT544                                                  = TestInstructions.TestInstructionTypeName_SubCustody_VerifMQMessageTypeMT
	TestInstructionDescription_SubCustody_ValidateMT544        string                                 = "Validate a MT544 message received from MQ"
	TestInstructionMouseOverText_SubCustody_ValidateMT544      string                                 = "Validate a MT544 message received from MQ"
	TestInstructionDeprecated_SubCustody_ValidateMT544         bool                                   = false
	TestInstructionEnabled_SubCustody_ValidateMT544            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_ValidateMT544 int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_ValidateMT544 int                                    = 0
	TestInstructionColor_SubCustody_ValidateMT544              TypeAndStructs.ColorType               = "#0DF36EAA"
	TCRuleDeletion_SubCustody_ValidateMT544                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_ValidateMT544                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                           TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"
	ExpectedMaxTestInstructionExecutionDurationInSeconds       int64                                  = 300

	// *** DropZone *** 'ValidateMT544_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_ValidateMT544_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "6435925e-dc84-4953-a3ef-d6575bcb387f"
	TestInstructionDropZoneName_SubCustody_ValidateMT544_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "ValidateMT544_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_ValidateMT544_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_ValidateMT544_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_ValidateMT544_ExpectsToSucceed       TypeAndStructs.ColorType        = "#0DF36EAA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_ValidateMT544_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "3b10634a-aaf5-4b62-89e6-90011d76b21d" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_ValidateMT544_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_ValidateMT544_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeDescription_SubCustody_ValidateMT544_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_ValidateMT544_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeActionCommand_SubCustody_ValidateMT544_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_ValidateMT544_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_ValidateMT544_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE

	// Attribute - 'RelatedReference_54x_20CRELA'
	TestInstructionAttributeUUID_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA          TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeName_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA          TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeType_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA          TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeDescription_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA   string                                          = "Extracts the response parameter from 54x_20CSEME that was sent on MQ"
	TestInstructionAttributeMouseOverText_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA string                                          = "Extracts the response parameter from 54x_20CSEME that was sent on MQ"
)

// TestInstruction_SubCustody_ValidateMT544
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_ValidateMT544 *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_ValidateMT544
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_ValidateMT544() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Generate Response variables for the TestInstruction
	Initate_TestInstructionResponseVariables_SubCustody_ValidateMT544()

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
		strconv.Itoa(TestInstructionMajorVersionNumber_SubCustody_ValidateMT544) + "_" +
			strconv.Itoa(TestInstructionMinorVersionNumber_SubCustody_ValidateMT544))

	// Create 'TestApiEngineClassesMethodsAttributesVersionMap' for this TestInstruction-version
	testApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString] = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct{
		TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstructionName:         TestInstructionName_SubCustody_ValidateMT544,
		TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_VerifyMQMessageTypeMT,
		TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_VerifyMQMessageTypeMT,
		TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyMT544,
		TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyMT544,
		Attributes:                  &testApiEngineMethodAttributeMap,
	}

	// Create testInstructionsMap for this TestInstruction
	testInstructionsMap[TestInstructionUUID_SubCustody_ValidateMT544] = &testApiEngineClassesMethodsAttributesVersionMap

	// Initiate variable to store all TestInstruction data
	TestInstruction_SubCustody_ValidateMT544 = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: ExpectedMaxTestInstructionExecutionDurationInSeconds,
				},
				TestInstructionsMap: &testInstructionsMap},
		},
	}

	// TestInstruction - ValidateMT544
	TestInstruction_SubCustody_ValidateMT544.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstructionName:          TestInstructionName_SubCustody_ValidateMT544,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_ValidateMT544,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_ValidateMT544,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_ValidateMT544,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_ValidateMT544,
		Deprecated:                   TestInstructionDeprecated_SubCustody_ValidateMT544,
		Enabled:                      TestInstructionEnabled_SubCustody_ValidateMT544,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_ValidateMT544,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_ValidateMT544,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - ValidateMT544
	TestInstruction_SubCustody_ValidateMT544.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstructionName:          TestInstructionName_SubCustody_ValidateMT544,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_ValidateMT544,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_ValidateMT544,
		Deprecated:                   TestInstructionDeprecated_SubCustody_ValidateMT544,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_ValidateMT544,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_ValidateMT544,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_ValidateMT544,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_ValidateMT544,
		TCRuleSwap:                   TCRuleSwap_SubCustody_ValidateMT544,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_ValidateMT544,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_ValidateMT544,
		Enabled:                      TestInstructionEnabled_SubCustody_ValidateMT544,
	}

	// DropZone 'ValidateMT544_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: ValidateMT544_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_ValidateMT544_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_ValidateMT544_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstructionName:          TestInstructionName_SubCustody_ValidateMT544,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_ValidateMT544_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_ValidateMT544_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_ValidateMT544_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_ValidateMT544_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_ValidateMT544_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_ValidateMT544_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_ValidateMT544_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_ValidateMT544_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_ValidateMT544,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_ValidateMT544_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_ValidateMT544.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_ValidateMT544.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_ValidateMT544_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_ValidateMT544_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_ValidateMT544_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstructionName:                           TestInstructionName_SubCustody_ValidateMT544,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_SubCustody_ExpectedToPass,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             "^(true|false)$",
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_ValidateMT544_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_ValidateMT544.TestInstructionAttribute = append(
		TestInstruction_SubCustody_ValidateMT544.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_ValidateMT544_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeName:         TestInstructionAttributeName_SubCustody_ValidateMT544_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_ValidateMT544_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed
	//TestInstruction_SubCustody_ValidateMT544.LocalExecutionMethods.TestInstructionsMap.Attributes[TestInstructionAttributeUUID_SubCustody_ValidateMT544_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// TestInstruction Attribute - 'RelatedReference_54x_20CRELA'
	var TestInstructionAttribute_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_ValidateMT544,
		TestInstructionName:                           TestInstructionName_SubCustody_ValidateMT544,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_SubCustody_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             "^.+$",
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
	}
	TestInstruction_SubCustody_ValidateMT544.TestInstructionAttribute = append(
		TestInstruction_SubCustody_ValidateMT544.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA)

	// Add TestApiEngine relation for Attribute - 'RelatedReference_54x_20CRELA'
	var tempTestApiEngineAttributeRelatedReference_54x_20CRELA *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeRelatedReference_54x_20CRELA = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
		TestInstructionAttributeName:         TestInstructionAttributeName_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_VerifyMQMessageTypeMT_RelatedReference_54x_20CRELA,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA] = tempTestApiEngineAttributeRelatedReference_54x_20CRELA
	//TestInstruction_SubCustody_ValidateMT544.LocalExecutionMethods.TestInstructionsMap.Attributes[TestInstructionAttributeUUID_SubCustody_ValidateMT544_RelatedReference_54x_20CRELA] = tempTestApiEngineAttributeRelatedReference_54x_20CRELA

	// ImmatureElementModel - ValidateMT544
	var TestInstructionImmatureElementModel_SubCustody_ValidateMT544 *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_ValidateMT544 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_ValidateMT544,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_ValidateMT544),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_ValidateMT544,
		NextElementUUID:          TestInstructionUUID_SubCustody_ValidateMT544,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_ValidateMT544,
		ParentElementUUID:        TestInstructionUUID_SubCustody_ValidateMT544,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_ValidateMT544,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_ValidateMT544,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_ValidateMT544.ImmatureElementModel = append(
		TestInstruction_SubCustody_ValidateMT544.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_ValidateMT544)

	return TestInstruction_SubCustody_ValidateMT544
}
