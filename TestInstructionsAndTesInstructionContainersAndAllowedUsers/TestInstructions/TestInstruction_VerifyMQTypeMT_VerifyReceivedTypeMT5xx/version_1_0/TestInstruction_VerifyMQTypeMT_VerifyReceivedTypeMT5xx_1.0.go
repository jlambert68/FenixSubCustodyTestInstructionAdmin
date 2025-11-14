package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_VerifyMQTypeMT_VerifyReceivedTypeMT5xx"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

const (

	// *************************************
	// *** TestInstruction *** 'VerifyReceivedTypeMT5xx'
	TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx
	TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx               TypeAndStructs.TestInstructionNameType = "VerifyReceivedTypeMT5xx"
	TestInstructionTypeUUID_SubCustody_VerifyReceivedTypeMT5xx                                                  = TestInstructions.TestInstructionTypeUUID_SubCustody_VerifMQMessageTypeMT
	TestInstructionTypeName_SubCustody_VerifyReceivedTypeMT5xx                                                  = TestInstructions.TestInstructionTypeName_SubCustody_VerifMQMessageTypeMT
	TestInstructionDescription_SubCustody_VerifyReceivedTypeMT5xx        string                                 = "Validate a MT5xx message received from MQ"
	TestInstructionMouseOverText_SubCustody_VerifyReceivedTypeMT5xx      string                                 = "Validate a MT5xx message received from MQ"
	TestInstructionDeprecated_SubCustody_VerifyReceivedTypeMT5xx         bool                                   = false
	TestInstructionEnabled_SubCustody_VerifyReceivedTypeMT5xx            bool                                   = true
	TestInstructionMajorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx int                                    = 1
	TestInstructionMinorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx int                                    = 0
	TestInstructionColor_SubCustody_VerifyReceivedTypeMT5xx              TypeAndStructs.ColorType               = "#0DF36EAA"
	TCRuleDeletion_SubCustody_VerifyReceivedTypeMT5xx                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_SubCustody_VerifyReceivedTypeMT5xx                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                                     TypeAndStructs.UpdatedTimeStampType    = "2023-11-27 13:00:00"
	ExpectedMaxTestInstructionExecutionDurationInSeconds                 int64                                  = 300

	// *** DropZone *** 'VerifyReceivedTypeMT5xx_ExpectsToSucceed'
	TestInstructionDropZoneUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "b6213a14-a8bd-484f-8767-fe8ad187c857"
	TestInstructionDropZoneName_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "VerifyReceivedTypeMT5xx_ExpectsToSucceed"
	TestInstructionDropZoneDescription_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed       TypeAndStructs.ColorType        = "#0DF36EAA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "a355e44b-9405-43b8-95a5-ed52f0436e66" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeDescription_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeActionCommand_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE

	// Attribute - 'RelatedReference_5xx_20CRELA'
	TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA          TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA          TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeType_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA          TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_RelatedReference_54x_20CRELA
	TestInstructionAttributeDescription_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA   string                                          = "Extracts the response parameter from 5xx_20CSEME that was sent on MQ"
	TestInstructionAttributeMouseOverText_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA string                                          = "Extracts the response parameter from 5xx_20CSEME that was sent on MQ"
)

// TestInstruction_SubCustody_VerifyReceivedTypeMT5xx
// Variable that holds the data for the TestInstruction
var TestInstruction_SubCustody_VerifyReceivedTypeMT5xx *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SubCustody_VerifyReceivedTypeMT5xx
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SubCustody_VerifyReceivedTypeMT5xx() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Generate Response variables for the TestInstruction
	Initate_TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx()

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
		strconv.Itoa(TestInstructionMajorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx) + "_" +
			strconv.Itoa(TestInstructionMinorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx))

	// Create 'TestApiEngineClassesMethodsAttributesVersionMap' for this TestInstruction-version
	testApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString] = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct{
		TestInstructionOriginalUUID: TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionName:         TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx,
		TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_VerifyMQTypeMT,
		TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_VerifyMQTypeMT,
		TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_UUID_SubCustody_VerifyMQMessageTypeMT_VerifyReceivedTypeMT5xx,
		TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_Name_SubCustody_VerifyMQMessageTypeMT_VerifyReceivedTypeMT5xx,
		Attributes:                  &testApiEngineMethodAttributeMap,
	}

	// Create testInstructionsMap for this TestInstruction
	testInstructionsMap[TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx] = &testApiEngineClassesMethodsAttributesVersionMap

	// Initiate variable to store all TestInstruction data
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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

	// TestInstruction - VerifyReceivedTypeMT5xx
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionName:          TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_VerifyReceivedTypeMT5xx,
		Deprecated:                   TestInstructionDeprecated_SubCustody_VerifyReceivedTypeMT5xx,
		Enabled:                      TestInstructionEnabled_SubCustody_VerifyReceivedTypeMT5xx,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - VerifyReceivedTypeMT5xx
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionName:          TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionTypeName:      TestInstructionTypeName_SubCustody_VerifyReceivedTypeMT5xx,
		Deprecated:                   TestInstructionDeprecated_SubCustody_VerifyReceivedTypeMT5xx,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SubCustody_VerifyReceivedTypeMT5xx,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SubCustody_VerifyReceivedTypeMT5xx,
		TCRuleDeletion:               TCRuleDeletion_SubCustody_VerifyReceivedTypeMT5xx,
		TCRuleSwap:                   TCRuleSwap_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionDescription:   TestInstructionDescription_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SubCustody_VerifyReceivedTypeMT5xx,
		Enabled:                      TestInstructionEnabled_SubCustody_VerifyReceivedTypeMT5xx,
	}

	// DropZone 'VerifyReceivedTypeMT5xx_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: VerifyReceivedTypeMT5xx_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		TestInstructionUUID:          TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionName:          TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx,
		DropZoneUUID:                 TestInstructionDropZoneUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_SubCustody_VerifyReceivedTypeMT5xx_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.ImmatureTestInstructionInformation = append(
		TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.ImmatureTestInstructionInformation,
		TestInstruction_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionName:                           TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
	}
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.TestInstructionAttribute = append(
		TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeName:         TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed
	//TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.LocalExecutionMethods.TestInstructionsMap.Attributes[TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// TestInstruction Attribute - 'RelatedReference_5xx_20CRELA'
	var TestInstructionAttribute_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_SubCustody,
		DomainName:                                    DomainData.DomainName_SubCustody,
		TestInstructionUUID:                           TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionName:                           TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
		TestInstructionAttributeName:                  TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
	}
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.TestInstructionAttribute = append(
		TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA)

	// Add TestApiEngine relation for Attribute - 'RelatedReference_5xx_20CRELA'
	var tempTestApiEngineAttributeRelatedReference_5xx_20CRELA *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeRelatedReference_5xx_20CRELA = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
		TestInstructionAttributeName:         TestInstructionAttributeName_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_VerifyMQMessageTypeMT_RelatedReference_5xx_20CRELA,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_VerifyMQMessageTypeMT_RelatedReference_5xx_20CRELA,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA] = tempTestApiEngineAttributeRelatedReference_5xx_20CRELA
	//TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.LocalExecutionMethods.TestInstructionsMap.Attributes[TestInstructionAttributeUUID_SubCustody_VerifyReceivedTypeMT5xx_RelatedReference_5xx_20CRELA] = tempTestApiEngineAttributeRelatedReference_5xx_20CRELA

	// ImmatureElementModel - VerifyReceivedTypeMT5xx
	var TestInstructionImmatureElementModel_SubCustody_VerifyReceivedTypeMT5xx *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SubCustody_VerifyReceivedTypeMT5xx = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SubCustody_VerifyReceivedTypeMT5xx),
		PreviousElementUUID:      TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		NextElementUUID:          TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		FirstChildElementUUID:    TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		ParentElementUUID:        TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		TopImmatureElementUUID:   TestInstructionUUID_SubCustody_VerifyReceivedTypeMT5xx,
		IsTopElement:             true,
	}
	TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.ImmatureElementModel = append(
		TestInstruction_SubCustody_VerifyReceivedTypeMT5xx.ImmatureElementModel,
		TestInstructionImmatureElementModel_SubCustody_VerifyReceivedTypeMT5xx)

	return TestInstruction_SubCustody_VerifyReceivedTypeMT5xx
}
