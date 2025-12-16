package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

// OwnerDomain and ExecutionDomain for TestInstruction // CHANGE HERE
const (
	domainUUID          TypeAndStructs.DomainUUIDType = DomainData.DomainUUID_SubCustody
	domainName          TypeAndStructs.DomainNameType = DomainData.DomainName_SubCustody
	executionDomainUUID TypeAndStructs.DomainUUIDType = DomainData.ExecutionDomainUUID_SubCustody_OnPrem
	executionDomainName TypeAndStructs.DomainNameType = DomainData.ExecutionDomainName_SubCustody_OnPrem
)

// Constants for TestInstruction "SubCustody_SendMT540"  // CHANGE HERE
const (
	testInstructionUUID                                  TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_SendMT540
	testInstructionTypeUUID                                                                     = TestInstructions.TestInstructionTypeUUID_SubCustody_SendOnMQTypeMT
	testInstructionTypeName                                                                     = TestInstructions.TestInstructionTypeName_SubCustody_SendOnMQTypeMT
	testInstructionName                                  TypeAndStructs.TestInstructionNameType = "SendMT540"
	testInstructionDescription                           string                                 = "Sends a MT540 message on MQ"
	testInstructionMouseOverText                         string                                 = "Sends a MT540 message on MQ"
	testInstructionDeprecated                            bool                                   = false
	testInstructionEnabled                               bool                                   = true
	testInstructionMajorVersionNumber                    int                                    = 1
	testInstructionMinorVersionNumber                    int                                    = 0
	testInstructionColor                                 TypeAndStructs.ColorType               = "#09AA4DAA"
	tCRuleDeletion                                       TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	tCRuleSwap                                           TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	testInstructionCreatingTimeStamp                     TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"
	expectedMaxTestInstructionExecutionDurationInSeconds int64                                  = 30
)

const (

	// *************************************
	// *** TestInstruction *** 'SendMT540'
	TestInstructionUUID_SubCustody_SendMT540                                  TypeAndStructs.OriginalElementUUIDType = testInstructionUUID
	TestInstructionName_SubCustody_SendMT540                                  TypeAndStructs.TestInstructionNameType = testInstructionName
	TestInstructionTypeUUID_SubCustody_SendMT540                                                                     = testInstructionTypeUUID
	TestInstructionTypeName_SubCustody_SendMT540                                                                     = testInstructionTypeName
	TestInstructionDescription_SubCustody_SendMT540                           string                                 = testInstructionDescription
	TestInstructionMouseOverText_SubCustody_SendMT540                         string                                 = testInstructionMouseOverText
	TestInstructionDeprecated_SubCustody_SendMT540                            bool                                   = testInstructionDeprecated
	TestInstructionEnabled_SubCustody_SendMT540                               bool                                   = testInstructionEnabled
	TestInstructionMajorVersionNumber_SubCustody_SendMT540                    int                                    = testInstructionMajorVersionNumber
	TestInstructionMinorVersionNumber_SubCustody_SendMT540                    int                                    = testInstructionMinorVersionNumber
	TestInstructionColor_SubCustody_SendMT540                                 TypeAndStructs.ColorType               = testInstructionColor
	TCRuleDeletion_SubCustody_SendMT540                                       TypeAndStructs.TCRuleDeletionType      = tCRuleDeletion
	TCRuleSwap_SubCustody_SendMT540                                           TypeAndStructs.TCRuleSwapType          = tCRuleSwap
	TestInstructionCreatingTimeStamp_SubCustody_SendMT540                     TypeAndStructs.UpdatedTimeStampType    = testInstructionCreatingTimeStamp
	ExpectedMaxTestInstructionExecutionDurationInSeconds_SubCustody_SendMT540 int64                                  = expectedMaxTestInstructionExecutionDurationInSeconds

	// DropZone // CHANGE HERE
	// *** DropZone *** 'SendMT540_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_SendMT540_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "ae344db4-2d34-474e-bd91-1b24ac408b75"
	TestInstructionDropZoneName_SubCustody_SendMT540_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "SendMT540_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_SendMT540_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_SendMT540_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_SendMT540_ExpectsToSucceed       TypeAndStructs.ColorType        = "#09AA4DAA"

	// Attribute  - SendMT540 - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_SendMT540_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "3b10634a-aaf5-4b62-89e6-90011d76b21d"
	TestInstructionAttributeName_SubCustody_SendMT540_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_SendMT540_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeActionCommand_SubCustody_SendMT540_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_SendMT540_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_SendMT540_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_SubCustody_SendMT540_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_SendMT540_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

/*
// var TestInstructionAttribute_SubCustody_SendMT540_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
var TestInstructionAttribute_SubCustody_SendMT540_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
	DomainUUID:                                    domainUUID,
	DomainName:                                    domainName,
	TestInstructionUUID:                           testInstructionUUID,
	TestInstructionName:                           testInstructionName,
	TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_SendMT540_ExpectedToBePassed,
	TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_SendMT540_ExpectedToBePassed,
	TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_SendMT540_ExpectedToBePassed,
	TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_SendMT540_ExpectedToBePassed,
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
	TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_SendMT540_ExpectedToBePassed,
}

*/

// TestInstruction_SubCustody_SendMT540
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_SendMT540 *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_SendMT540
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_SendMT540() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Generate Response variables for the TestInstruction
	Initate_TestInstructionResponseVariables_SubCustody_SendMT540()

	// Create and Initialize maps used in 'LocalExecutionMethods'
	var testApiEngineMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	testApiEngineMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct)

	var testInstructionsMap TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
	testInstructionsMap = make(TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType)

	var testApiEngineClassesMethodsAttributesVersionMap TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
	testApiEngineClassesMethodsAttributesVersionMap = make(TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType)

	// Create version as 'string'
	var versionNumberAsString TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type // "1_0" or "13_3" ...)
	versionNumberAsString = TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type(
		strconv.Itoa(TestInstructionMajorVersionNumber_SubCustody_SendMT540) + "_" +
			strconv.Itoa(TestInstructionMinorVersionNumber_SubCustody_SendMT540))

	// Create 'TestApiEngineClassesMethodsAttributesVersionMap' for this TestInstruction-version
	testApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString] = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct{
		TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_SendMT540,
		TestInstructionName:         TestInstructionName_SubCustody_SendMT540,
		TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_SendOnMQTypeMT,
		TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_SendOnMQTypeMT,
		TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMT540,
		TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMT540,
		Attributes:                  &testApiEngineMethodAttributeMap,
	}

	// Create testInstructionsMap for this TestInstruction
	testInstructionsMap[TestInstructionUUID_SubCustody_SendMT540] = &testApiEngineClassesMethodsAttributesVersionMap

	// Initiate variable to store all TestInstruction data
	TestInstruction_SubCustody_SendMT540 = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: ExpectedMaxTestInstructionExecutionDurationInSeconds_SubCustody_SendMT540,
				},
				TestInstructionsMap: &testInstructionsMap},
		},
	}

	// TestInstruction - SendMT540
	TestInstruction_SubCustody_SendMT540.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   domainUUID,
		DomainName:                   domainName,
		ExecutionDomainUUID:          executionDomainUUID,
		ExecutionDomainName:          executionDomainName,
		TestInstructionUUID:          testInstructionUUID,
		TestInstructionName:          testInstructionName,
		TestInstructionTypeUUID:      testInstructionTypeUUID,
		TestInstructionTypeName:      testInstructionTypeName,
		TestInstructionDescription:   testInstructionDescription,
		TestInstructionMouseOverText: testInstructionMouseOverText,
		Deprecated:                   testInstructionDeprecated,
		Enabled:                      testInstructionEnabled,
		MajorVersionNumber:           testInstructionMajorVersionNumber,
		MinorVersionNumber:           testInstructionMinorVersionNumber,
		UpdatedTimeStamp:             testInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - SendMT540
	TestInstruction_SubCustody_SendMT540.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   domainUUID,
		DomainName:                   domainName,
		ExecutionDomainUUID:          executionDomainUUID,
		ExecutionDomainName:          executionDomainName,
		TestInstructionUUID:          testInstructionUUID,
		TestInstructionName:          testInstructionName,
		TestInstructionTypeUUID:      testInstructionTypeUUID,
		TestInstructionTypeName:      testInstructionTypeName,
		Deprecated:                   testInstructionDeprecated,
		MajorVersionNumber:           testInstructionMajorVersionNumber,
		MinorVersionNumber:           testInstructionMinorVersionNumber,
		UpdatedTimeStamp:             testInstructionCreatingTimeStamp,
		TestInstructionColor:         testInstructionColor,
		TCRuleDeletion:               tCRuleDeletion,
		TCRuleSwap:                   tCRuleSwap,
		TestInstructionDescription:   testInstructionDescription,
		TestInstructionMouseOverText: testInstructionMouseOverText,
		Enabled:                      testInstructionEnabled,
	}

	// DropZone 'SendMT540_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: SendMT540_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_SendMT540_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_SendMT540_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   domainUUID,
		DomainName:                   domainName,
		TestInstructionUUID:          testInstructionUUID,
		TestInstructionName:          testInstructionName,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_SendMT540_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_SendMT540_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_SendMT540_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_SendMT540_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_SendMT540_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_SendMT540_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_SendMT540_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_SendMT540_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_SendMT540,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_SendMT540_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_SendMT540.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_SendMT540.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_SendMT540_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_SendMT540_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_SendMT540_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    domainUUID,
		DomainName:                                    domainName,
		TestInstructionUUID:                           testInstructionUUID,
		TestInstructionName:                           testInstructionName,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_SendMT540_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_SendMT540_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_SendMT540.TestInstructionAttribute = append(
		TestInstruction_SubCustody_SendMT540.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_SendMT540_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeName:         TestInstructionAttributeName_SubCustody_SendMT540_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_SendMT540_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// ImmatureElementModel - SendMT540
	var TestInstructionImmatureElementModel_SubCustody_SendMT540 *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_SendMT540 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               domainUUID,
		DomainName:               domainName,
		ImmatureElementUUID:      testInstructionUUID,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(testInstructionName),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_SendMT540,
		NextElementUUID:          TestInstructionUUID_SubCustody_SendMT540,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_SendMT540,
		ParentElementUUID:        TestInstructionUUID_SubCustody_SendMT540,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_SendMT540,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_SendMT540,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_SendMT540.ImmatureElementModel = append(
		TestInstruction_SubCustody_SendMT540.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_SendMT540)

	return TestInstruction_SubCustody_SendMT540
}
