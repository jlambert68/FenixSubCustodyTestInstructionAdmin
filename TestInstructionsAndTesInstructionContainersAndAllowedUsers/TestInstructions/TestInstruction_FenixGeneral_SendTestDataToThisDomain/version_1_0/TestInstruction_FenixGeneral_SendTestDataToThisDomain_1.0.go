package version_1_0

import (
	TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

const (
	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "5d354d15-5612-4e83-a0a3-98654f21f989" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassed
	TestInstructionAttributeType_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassed
	TestInstructionAttributeValueAsStringValue_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_FenixGeneral_SendTestDataToThisDomain
// Variable that holds the data for the TestInstruction
var TestInstruction_FenixGeneral_SendTestDataToThisDomain *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_FenixGeneral_SendTestDataToThisDomain
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_FenixGeneral_SendTestDataToThisDomain() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Generate Response variables for the TestInstruction
	Initate_TestInstructionResponseVariables_FenixGeneral_SendTestDataToThisDomain()

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
		strconv.Itoa(TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionMajorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain) + "_" +
			strconv.Itoa(TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
				TestInstructionMinorVersionNumber_FenixSentToUsersDomain_SendTestDataToThisDomain))

	// Create 'TestApiEngineClassesMethodsAttributesVersionMap' for this TestInstruction-version
	testApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString] = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct{
		TestInstructionOriginalUUID: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestInstructionName:         TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.TestInstructionName_FenixSentToUsersDomain_SendTestDataToThisDomain,
		TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_FenixGeneral,
		TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name__FenixGeneral,
		TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_UUID_FenixGeneral_SendTestDataToThisDomain,
		TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_Name_FenixGeneralT_SendTestDataToThisDomain,
		Attributes:                  &testApiEngineMethodAttributeMap,
	}

	// Create testInstructionsMap for this TestInstruction
	testInstructionsMap[TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
		TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain] = &testApiEngineClassesMethodsAttributesVersionMap

	// Initiate variable to store all TestInstruction data
	TestInstruction_FenixGeneral_SendTestDataToThisDomain = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
						ExpectedMaxTestInstructionExecutionDurationInSeconds,
				},
				TestInstructionsMap: &testInstructionsMap},
		},
	}

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed,
		TestInstructionAttributeName:         TestInstructionAttributeName_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_FenixGeneral_SendTestDataToThisDomain_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// Add TestApiEngine relation for Attribute - 'ChosenTestDataAsJsonString'
	var tempTestApiEngineAttributeChosenTestDataAsJsonString *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeChosenTestDataAsJsonString = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeUUID_SendTestDataToThisDomain_ChosenTestDataAsJsonString,
		TestInstructionAttributeName: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeName_SendTestDataToThisDomain_ChosenTestDataAsJsonString,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_FenixGeneral_SendTestDataToThisDomain_ChosenTestDataAsJsonString,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SFenixGeneral_SendTestDataToThisDomain_ChosenTestDataAsJsonString,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
		TestInstructionAttributeUUID_SendTestDataToThisDomain_ChosenTestDataAsJsonString] = tempTestApiEngineAttributeChosenTestDataAsJsonString

	// Should not be sent to TestApiEngine
	// Add TestApiEngine relation for Attribute - 'SendTestDataToThisExecutionDomain'
	var tempSendTestDataToThisExecutionDomain *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempSendTestDataToThisExecutionDomain = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeName: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeName_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_ZeroEmptyUuid,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_ZeroEmptyName,
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
		TestInstructionAttributeUUID_FenixSentToUsersDomain_SendTestDataToThisDomain_SendTestDataToThisExecutionDomain] = tempSendTestDataToThisExecutionDomain

	// Should not be sent to TestApiEngine
	// Add TestApiEngine relation for Attribute - 'SendTestDataToThisDomainTextBox'
	var tempSendTestDataToThisDomainTextBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempSendTestDataToThisDomainTextBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeUUID_SendTestDataToThisDomain_SendTestDataToThisDomainTextBox,
		TestInstructionAttributeName: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeName_SendTestDataToThisDomain_SendTestDataToThisDomainTextBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_ZeroEmptyUuid,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_ZeroEmptyName,
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
		TestInstructionAttributeUUID_SendTestDataToThisDomain_SendTestDataToThisDomainTextBox] = tempSendTestDataToThisDomainTextBox

	// Should not be sent to TestApiEngine
	// Add TestApiEngine relation for Attribute - 'SendTestDataToThisExecutionDomainTextBox'
	var tempSendTestDataToThisExecutionDomainTextBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempSendTestDataToThisExecutionDomainTextBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeUUID_SendTestDataToThisDomain_SendTestDataToThisExecutionDomainTextBox,
		TestInstructionAttributeName: TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
			TestInstructionAttributeName_SendTestDataToThisDomain_SendTestDataToThisExecutionDomainTextBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_ZeroEmptyUuid,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_ZeroEmptyName,
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_FenixGeneral_SendTestDataToThisDomain_v1_0.
		TestInstructionAttributeUUID_SendTestDataToThisDomain_SendTestDataToThisExecutionDomainTextBox] = tempSendTestDataToThisExecutionDomainTextBox

	return TestInstruction_FenixGeneral_SendTestDataToThisDomain
}
