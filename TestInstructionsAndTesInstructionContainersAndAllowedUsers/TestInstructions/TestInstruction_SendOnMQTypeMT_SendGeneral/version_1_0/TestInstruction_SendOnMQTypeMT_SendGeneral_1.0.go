package version_1_0

import (
	fenixTestInstructions "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	TestInstruction_SendTemplateToThisDomain_version_1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	TestApiEngineClassesAndMethodsAndAttributes "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendGeneral"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

const (

	// *************************************
	// *** TestInstruction *** 'SendSwiftMessageOnMQTypeMT'
	TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate               TypeAndStructs.OriginalElementUUIDType     = TestInstruction_SendOnMQTypeMT_SendGeneral.TestInstructionUUID_SendOnMQTypeMT_SendGeneral
	TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate               TypeAndStructs.TestInstructionNameType     = "SendSwiftMessageOnMQTypeMT"
	TestInstructionTypeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate           TypeAndStructs.TestInstructionTypeUUIDType = fenixTestInstructions.TestInstructionTypeUUID_FenixSentToUsersDomain_FenixSendTemplateAddOn
	TestInstructionTypeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate           TypeAndStructs.TestInstructionTypeNameType = fenixTestInstructions.TestInstructionTypeName_FenixSentToUsersDomain_FenixSendTemplateAddOn
	TestInstructionDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate        string                                     = "This TestInstruction sends Swift MT-message on MQ"
	TestInstructionMouseOverText_SendOnMQTypeMT_FenixAddonSendGeneralTemplate      string                                     = "This TestInstruction sends Swift MT-message on MQ"
	TestInstructionDeprecated_SendOnMQTypeMT_FenixAddonSendGeneralTemplate         bool                                       = false
	TestInstructionEnabled_SendOnMQTypeMT_FenixAddonSendGeneralTemplate            bool                                       = true
	TestInstructionMajorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate int                                        = 1
	TestInstructionMinorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate int                                        = 0
	TestInstructionColor_SendOnMQTypeMT_FenixAddonSendGeneralTemplate              TypeAndStructs.ColorType                   = "#DBBC0CAA"
	TCRuleDeletion_SendOnMQTypeMT_FenixAddonSendGeneralTemplate                    TypeAndStructs.TCRuleDeletionType          = "TCRuleDeletion020"
	TCRuleSwap_SendOnMQTypeMT_FenixAddonSendGeneralTemplate                        TypeAndStructs.TCRuleSwapType              = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                                               TypeAndStructs.UpdatedTimeStampType        = "2024-09-19 13:00:00"
	ExpectedMaxTestInstructionExecutionDurationInSeconds                           int64                                      = 30

	// Attribute - 'SwiftMessageTypeCombobox'
	TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox        TypeAndStructs.TestInstructionAttributeUUIDType = "79e71319-4dcf-4939-8b6f-0c4b1b0cb6b1" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_SendOnMQTypeMT_SendTemplateToThisExecutionDomain
	TestInstructionAttributeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox        TypeAndStructs.TestInstructionAttributeNameType = "Swift MT-message type"
	TestInstructionAttributeType_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox        TypeAndStructs.TestInstructionAttributeTypeType = "COMBOBOX"
	TestInstructionAttributeDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox string                                          = "The Template is of this Swift MT-message Type"
	TestInstructionAttributeMouseOver_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox   string                                          = "The Template is of this Swift MT-message Type"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed          TypeAndStructs.TestInstructionAttributeUUIDType = "a160b777-caa7-44c7-a4c4-7814f8fd3691" //TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed
	TestInstructionAttributeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed          TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_SubCustody_ExpectedToBePassedComboBox
	TestInstructionAttributeType_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed          TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_SubCustody_ExpectedToBePassedComboBox
	TestInstructionAttributeDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed   string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

var (
	// Attribute - 'FenixOwnedSendTemplateComboBox' and 'FenixOwnedSendTemplateToThisExecutionDomainComboBox'
	TestInstructionAttributeComboBoxPredefinedValues_SendOnMQTypeMT_EmptyListForComboBox TypeAndStructs.TestInstructionAttributeComboBoxPredefinedValuesType

	// Attribute - 'FenixOwnedSendRawTemplateComboBox'
	TestInstructionAttributeComboBoxPredefinedValues_FenixAddonSendGeneralTemplate_SwiftMessageTypeComboboxValues = TypeAndStructs.
															TestInstructionAttributeComboBoxPredefinedValuesType{"MT540", "MT542"}

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeComboBoxPredefinedValues_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed = TypeAndStructs.
																TestInstructionAttributeComboBoxPredefinedValuesType{"true", "false"}
)

// TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate
// Variable that holds the data for the TestInstruction
var TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Generate Response variables for the TestInstruction
	Initate_TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral()

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
		strconv.Itoa(TestInstructionMajorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate) + "_" +
			strconv.Itoa(TestInstructionMinorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate))

	// Create 'TestApiEngineClassesMethodsAttributesVersionMap' for this TestInstruction-version
	testApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString] = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct{
		TestInstructionOriginalUUID: TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionName:         TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestApiEngineClassNameUUID:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_SendOnMQTypeMT,
		TestApiEngineClassNameNAME:  TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_SendOnMQTypeMT,
		TestApiEngineMethodNameUUID: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_UUID_SubCustody_SendOnMQTypeMT_SendMTGeneral,
		TestApiEngineMethodNameNAME: TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodName_Name_SubCustody_SendOnMQTypeMT_SendMTGeneral,
		Attributes:                  &testApiEngineMethodAttributeMap,
	}

	// Create testInstructionsMap for this TestInstruction
	testInstructionsMap[TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate] = &testApiEngineClassesMethodsAttributesVersionMap

	// Initiate variable to store all TestInstruction data
	TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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

	// TestInstruction - FenixOwnedSendTemplateToThisDomain
	TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionName:          TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionTypeName:      TestInstructionTypeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionDescription:   TestInstructionDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		Deprecated:                   TestInstructionDeprecated_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		Enabled:                      TestInstructionEnabled_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - FenixOwnedSendTemplateToThisDomain
	TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_SubCustody,
		DomainName:                   DomainData.DomainName_SubCustody,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_SubCustody_OnPrem,
		ExecutionDomainName:          DomainData.ExecutionDomainName_SubCustody_OnPrem,
		TestInstructionUUID:          TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionName:          TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionTypeName:      TestInstructionTypeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		Deprecated:                   TestInstructionDeprecated_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TCRuleDeletion:               TCRuleDeletion_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TCRuleSwap:                   TCRuleSwap_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionDescription:   TestInstructionDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionMouseOverText: TestInstructionMouseOverText_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		Enabled:                      TestInstructionEnabled_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
	}

	// TestInstruction Attribute - 'SwiftMessageTypeCombobox'
	var TestInstructionAttribute_SubCustody_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_SubCustody_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_SubCustody,
		DomainName:                                       DomainData.DomainName_SubCustody,
		TestInstructionUUID:                              TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionName:                              TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeName:                     TestInstructionAttributeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOver_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeTypeUUID:                 TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestInstructionAttributeTypeName:                 TestInstructions.TestInstructionAttributeTypeName_SubCustody_Standard,
		TestInstructionAttributeValueAsString:            Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:                Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:                  true,
		TestInstructionAttributeEnabled:                  true,
		TestInstructionAttributeMandatory:                true,
		TestInstructionAttributeVisibleInTestCaseArea:    false,
		TestInstructionAttributeIsDeprecated:             false,
		TestInstructionAttributeInputMask:                "^.+$",
		TestInstructionAttributeType:                     TestInstructionAttributeType_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeComboBoxPredefinedValues: TestInstructionAttributeComboBoxPredefinedValues_FenixAddonSendGeneralTemplate_SwiftMessageTypeComboboxValues,
	}
	TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstructionAttribute = append(
		TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstructionAttribute,
		TestInstructionAttribute_SubCustody_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed)

	// Add TestApiEngine relation for Attribute - 'SwiftMessageTypeCombobox'
	var tempTestApiEngineAttributeSwiftMessageTypeCombobox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeSwiftMessageTypeCombobox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeName:         TestInstructionAttributeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_SendOnMQTypeMT_SwiftMessageType,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_SendOnMQTypeMT_SwiftMessageType,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_SwiftMessageTypeCombobox] = tempTestApiEngineAttributeSwiftMessageTypeCombobox

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassedCombobox *TypeAndStructs.TestInstructionAttributeStruct
	tempTestApiEngineAttributeExpectedToBePassedCombobox = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                       DomainData.DomainUUID_SubCustody,
		DomainName:                                       DomainData.DomainName_SubCustody,
		TestInstructionUUID:                              TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionName:                              TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestInstructionAttributeUUID:                     TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeName:                     TestInstructionAttributeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeDescription:              TestInstructionAttributeDescription_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:                TestInstructionAttributeMouseOverText_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:                 TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestInstructionAttributeTypeName:                 TestInstructions.TestInstructionAttributeTypeName_SubCustody_ExpectedToPass,
		TestInstructionAttributeValueAsString:            Domains.TestInstructionAttributeValueAsStringValue_TRUE,
		TestInstructionAttributeValueUUID:                Domains.TestInstructionAttributeValueUUID_TRUE,
		TestInstructionAttributeVisible:                  true,
		TestInstructionAttributeEnabled:                  true,
		TestInstructionAttributeMandatory:                true,
		TestInstructionAttributeVisibleInTestCaseArea:    false,
		TestInstructionAttributeIsDeprecated:             false,
		TestInstructionAttributeInputMask:                "^(true|false)$",
		TestInstructionAttributeType:                     TestInstructionAttributeType_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeComboBoxPredefinedValues: TestInstructionAttributeComboBoxPredefinedValues_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
	}
	TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstructionAttribute = append(
		TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstructionAttribute,
		tempTestApiEngineAttributeExpectedToBePassedCombobox)

	// Add TestApiEngine relation for Attribute - 'ExpectedToBePassed'
	var tempTestApiEngineAttributeExpectedToBePassed *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeExpectedToBePassed = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeName:         TestInstructionAttributeName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_ExpectedToPass,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_UUID_SubCustody_GeneralAttribute_ExpectedToBePassed,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_ClassName_Name_SubCustody_GeneralAttribute_ExpectedToBePassed,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstructionAttributeUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate_ExpectedToBePassed] = tempTestApiEngineAttributeExpectedToBePassed

	// *** Add Attributes that comes from Fenix ***

	// Add TestApiEngine relation for Attribute - 'FenixOwnedSendTemplateComboBox' which not should be sent to TestApiEngine
	var tempTestApiEngineAttributeFenixOwnedSendTemplateComboBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeFenixOwnedSendTemplateComboBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_Type(""),
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_Type(""),
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox] =
		tempTestApiEngineAttributeFenixOwnedSendTemplateComboBox

	// Add TestApiEngine relation for Fenix-Attribute - 'ReplacePlaceholers'
	var tempTestApiEngineAttributeReplacePlaceholders *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeReplacePlaceholders = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateReplacePlaceholersComboBox,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateReplacePlaceholersComboBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_SendOnMQTypeMT_ReplacePlaceholers,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_SendOnMQTypeMT_ReplacePlaceholers,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateReplacePlaceholersComboBox] = tempTestApiEngineAttributeReplacePlaceholders

	// Add TestApiEngine relation for Fenix-Attribute - 'TemplateAsString'
	var tempTestApiEngineAttributeTemplateAsString *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeTemplateAsString = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedTemplateAsString,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_FenixOwnedTemplateAsString,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_SendOnMQTypeMT_TemplateAsString,
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_SendOnMQTypeMT_TemplateAsString,
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedTemplateAsString] = tempTestApiEngineAttributeTemplateAsString

	// Add TestApiEngine relation for Attribute - 'FenixOwnedSendTemplateToThisExecutionDomainComboBox' which not should be sent to TestApiEngine
	var tempTestApiEngineAttributeFenixOwnedSendTemplateToThisExecutionDomainComboBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeFenixOwnedSendTemplateToThisExecutionDomainComboBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainComboBox,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainComboBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_Type(""),
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_Type(""),
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainComboBox] =
		tempTestApiEngineAttributeFenixOwnedSendTemplateToThisExecutionDomainComboBox

	// Add TestApiEngine relation for Attribute - 'FenixOwnedSendTemplateToThisDomainTextBox' which not should be sent to TestApiEngine
	var tempTestApiEngineAttributeFenixOwnedSendTemplateToThisDomainTextBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeFenixOwnedSendTemplateToThisDomainTextBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_Type(""),
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_Type(""),
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisDomainTextBox] =
		tempTestApiEngineAttributeFenixOwnedSendTemplateToThisDomainTextBox

	// Add TestApiEngine relation for Attribute - 'FenixOwnedSendTemplateToThisExecutionDomainTextBox' which not should be sent to TestApiEngine
	var tempTestApiEngineAttributeFenixOwnedSendTemplateToThisExecutionDomainTextBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeFenixOwnedSendTemplateToThisExecutionDomainTextBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainTextBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_Type(""),
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_Type(""),
		AttributeShouldBeSentToTestApiEngine: false,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateToThisExecutionDomainTextBox] =
		tempTestApiEngineAttributeFenixOwnedSendTemplateToThisExecutionDomainTextBox

	// Add TestApiEngine relation for Attribute - 'FenixOwnedSendTemplateToThisDomainFenixOwnedSendTemplateComboBox' which should be sent to TestApiEngine
	var tempTestApiEngineAttributeFenixOwnedSendTemplateToThisDomainFenixOwnedSendTemplateComboBox *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct
	tempTestApiEngineAttributeFenixOwnedSendTemplateToThisDomainFenixOwnedSendTemplateComboBox = &TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineAttributesStruct{
		TestInstructionAttributeUUID:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox,
		TestInstructionAttributeName:         TestInstruction_SendTemplateToThisDomain_version_1_0.TestInstructionAttributeName_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox,
		TestInstructionAttributeTypeUUID:     TestInstructions.TestInstructionAttributeTypeUUID_SubCustody_Standard,
		TestApiEngineAttributeNameUUID:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_UUID_SubCustody_Type(""),
		TestApiEngineAttributeNameName:       TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_AttributeName_Name_SubCustody_Type(""),
		AttributeShouldBeSentToTestApiEngine: true,
	}

	testApiEngineMethodAttributeMap[TestInstruction_SendTemplateToThisDomain_version_1_0.
		TestInstructionAttributeUUID_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain_FenixOwnedSendTemplateComboBox] =
		tempTestApiEngineAttributeFenixOwnedSendTemplateToThisDomainFenixOwnedSendTemplateComboBox

	// ImmatureElementModel - FenixOwnedSendTemplateToThisDomain
	var TestInstructionImmatureElementModel_SendOnMQTypeMT_FenixAddonSendGeneralTemplate *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_SendOnMQTypeMT_FenixAddonSendGeneralTemplate = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_SubCustody,
		DomainName:               DomainData.DomainName_SubCustody,
		ImmatureElementUUID:      TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_SendOnMQTypeMT_FenixAddonSendGeneralTemplate),
		PreviousElementUUID:      TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		NextElementUUID:          TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		FirstChildElementUUID:    TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		ParentElementUUID:        TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		TopImmatureElementUUID:   TestInstructionUUID_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
		IsTopElement:             true,
	}
	TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.ImmatureElementModel = append(
		TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.ImmatureElementModel,
		TestInstructionImmatureElementModel_SendOnMQTypeMT_FenixAddonSendGeneralTemplate)

	return TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate
}
