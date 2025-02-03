package TestInstructionsAndTesInstructionContainersAndAllowedUsers

import (
	"fmt"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods"
	TestApiEngineClassesAndMethodsAndAttributes "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	FenixGeneral_SendTestDataToThisDomain_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_FenixGeneral_SendTestDataToThisDomain/version_1_0"
	SendOnMQTypeMT_SendGeneral "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendGeneral"
	SendOnMQTypeMT_SendGeneral_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendGeneral/version_1_0"
	SendOnMQTypeMT_SendMT540 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	SendOnMQTypeMT_SendMT540_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0"
	SendOnMQTypeMT_SendMT542 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542"
	SendOnMQTypeMT_SendMT542_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	VerifyMQMessageTypeMT_VerifyMT544 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544"
	VerifyMQMessageTypeMT_VerifyMT544_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT544/version_1_0"
	VerifyMQMessageTypeMT_VerifyMT546 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546"
	VerifyMQMessageTypeMT_VerifyMT546_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT546/version_1_0"
	VerifyMQMessageTypeMT_VerifyMT548 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548"
	VerifyMQMessageTypeMT_VerifyMT548_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_ValidateMQTypeMT54x_ValidateMT548/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"os"
	"time"
)

// TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody
// Overall Message
var TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody(allowedUsers []byte) (
	*TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct,
	*TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType) {

	var err error

	// Load Allowed users from json-file
	err = shared_code.ParseAllowedUsersFromEmbeddedFile(allowedUsers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set AllUsersAuthorizationRights and initial Hash-value
	shared_code.AllowedUsersLoadFromJsonFile.AllUsersAuthorizationRights = &TestInstructionAndTestInstuctionContainerTypes.AllUsersAuthorizationRightsStruct{
		AllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain:  false,
		AllUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomain: false,
	}
	shared_code.AllowedUsersLoadFromJsonFile.AllowedUsersHash = shared_code.InitialValueBeforeHashed

	// Generate TestInstructions
	// SendOnMQTypeMT::SendMT540
	SendOnMQTypeMT_SendMT540_1_0.Initate_TestInstruction_SubCustody_SendMT540()

	// SendOnMQTypeMT::SendMT542
	SendOnMQTypeMT_SendMT542_1_0.Initate_TestInstruction_SubCustody_SendMT542()

	// VerifyMQMessageTypeMT::VerifyMT544
	VerifyMQMessageTypeMT_VerifyMT544_1_0.Initate_TestInstruction_SubCustody_ValidateMT544()

	// VerifyMQMessageTypeMT::VerifyMT546
	VerifyMQMessageTypeMT_VerifyMT546_1_0.Initate_TestInstruction_SubCustody_ValidateMT546()

	// VerifyMQMessageTypeMT::VerifyMT548
	VerifyMQMessageTypeMT_VerifyMT548_1_0.Initate_TestInstruction_SubCustody_ValidateMT548()

	// SendOnMQTypeMT::FenixAddonSendGeneralTemplate
	SendOnMQTypeMT_SendGeneral_1_0.Initate_TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate()

	// Initate_TestInstruction_FenixGeneral_SendTestDataToThisDomain
	FenixGeneral_SendTestDataToThisDomain_1_0.Initate_TestInstruction_FenixGeneral_SendTestDataToThisDomain()

	// Build structure for all TestInstructions & TestInstructionContainers to be sent over gRPC to Fenix Backend
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'SendOnMQTypeMT_SendMT540'
				SendOnMQTypeMT_SendGeneral.TestInstructionUUID_SendOnMQTypeMT_SendGeneral: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// SendOnMQTypeMT::FenixAddonSendGeneralTemplate_1_0
						{
							TestInstructionInstance:             SendOnMQTypeMT_SendGeneral_1_0.TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate,
							TestInstructionInstanceMajorVersion: SendOnMQTypeMT_SendGeneral_1_0.TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: SendOnMQTypeMT_SendGeneral_1_0.TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstruction.MinorVersionNumber,
							Deprecated:                          SendOnMQTypeMT_SendGeneral_1_0.TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstruction.Deprecated,
							Enabled:                             SendOnMQTypeMT_SendGeneral_1_0.TestInstruction_SendOnMQTypeMT_FenixAddonSendGeneralTemplate.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							ResponseVariablesMapStructure: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
								ResponseVariablesMap:     SendOnMQTypeMT_SendGeneral_1_0.TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral,
								ResponseVariablesMapHash: shared_code.InitialValueBeforeHashed,
							},
							TestInstructionInstanceVersionAndResponseVariablesHash: shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
				//TestInstruction 'SendOnMQTypeMT_SendMT540'
				SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'SendOnMQTypeMT_SendMT540_1_0'
						{
							TestInstructionInstance:             SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540,
							TestInstructionInstanceMajorVersion: SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.MinorVersionNumber,
							Deprecated:                          SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.Deprecated,
							Enabled:                             SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							ResponseVariablesMapStructure: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
								ResponseVariablesMap:     SendOnMQTypeMT_SendMT540_1_0.TestInstructionResponseVariables_SubCustody_SendMT540,
								ResponseVariablesMapHash: shared_code.InitialValueBeforeHashed,
							},
							TestInstructionInstanceVersionAndResponseVariablesHash: shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
				//TestInstruction 'SendOnMQTypeMT_SendMT542'
				SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'SendOnMQTypeMT_SendMT542_1_0'
						{
							TestInstructionInstance:             SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542,
							TestInstructionInstanceMajorVersion: SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.MinorVersionNumber,
							Deprecated:                          SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.Deprecated,
							Enabled:                             SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							ResponseVariablesMapStructure: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
								ResponseVariablesMap:     SendOnMQTypeMT_SendMT542_1_0.TestInstructionResponseVariables_SubCustody_SendMT542,
								ResponseVariablesMapHash: shared_code.InitialValueBeforeHashed,
							},
							TestInstructionInstanceVersionAndResponseVariablesHash: shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				//TestInstruction 'VerifyMQMessageTypeMT_VerifyMT544'
				VerifyMQMessageTypeMT_VerifyMT544.TestInstructionUUID_SubCustody_ValidateMT544: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'VerifyMQMessageTypeMT_VerifyMT544_1_0'
						{
							TestInstructionInstance:             VerifyMQMessageTypeMT_VerifyMT544_1_0.TestInstruction_SubCustody_ValidateMT544,
							TestInstructionInstanceMajorVersion: VerifyMQMessageTypeMT_VerifyMT544_1_0.TestInstruction_SubCustody_ValidateMT544.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: VerifyMQMessageTypeMT_VerifyMT544_1_0.TestInstruction_SubCustody_ValidateMT544.TestInstruction.MinorVersionNumber,
							Deprecated:                          VerifyMQMessageTypeMT_VerifyMT544_1_0.TestInstruction_SubCustody_ValidateMT544.TestInstruction.Deprecated,
							Enabled:                             VerifyMQMessageTypeMT_VerifyMT544_1_0.TestInstruction_SubCustody_ValidateMT544.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							ResponseVariablesMapStructure: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
								ResponseVariablesMap:     VerifyMQMessageTypeMT_VerifyMT544_1_0.TestInstructionResponseVariables_SubCustody_ValidateMT544,
								ResponseVariablesMapHash: shared_code.InitialValueBeforeHashed,
							},
							TestInstructionInstanceVersionAndResponseVariablesHash: shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				//TestInstruction 'VerifyMQMessageTypeMT_VerifyMT546'
				VerifyMQMessageTypeMT_VerifyMT546.TestInstructionUUID_SubCustody_ValidateMT546: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'VerifyMQMessageTypeMT_VerifyMT546_1_0'
						{
							TestInstructionInstance:             VerifyMQMessageTypeMT_VerifyMT546_1_0.TestInstruction_SubCustody_ValidateMT546,
							TestInstructionInstanceMajorVersion: VerifyMQMessageTypeMT_VerifyMT546_1_0.TestInstruction_SubCustody_ValidateMT546.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: VerifyMQMessageTypeMT_VerifyMT546_1_0.TestInstruction_SubCustody_ValidateMT546.TestInstruction.MinorVersionNumber,
							Deprecated:                          VerifyMQMessageTypeMT_VerifyMT546_1_0.TestInstruction_SubCustody_ValidateMT546.TestInstruction.Deprecated,
							Enabled:                             VerifyMQMessageTypeMT_VerifyMT546_1_0.TestInstruction_SubCustody_ValidateMT546.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							ResponseVariablesMapStructure: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
								ResponseVariablesMap:     VerifyMQMessageTypeMT_VerifyMT546_1_0.TestInstructionResponseVariables_SubCustody_ValidateMT546,
								ResponseVariablesMapHash: shared_code.InitialValueBeforeHashed,
							},
							TestInstructionInstanceVersionAndResponseVariablesHash: shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
				//TestInstruction 'VerifyMQMessageTypeMT_VerifyMT548'
				VerifyMQMessageTypeMT_VerifyMT548.TestInstructionUUID_SubCustody_ValidateMT548: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'VerifyMQMessageTypeMT_VerifyMT548_1_0'
						{
							TestInstructionInstance:             VerifyMQMessageTypeMT_VerifyMT548_1_0.TestInstruction_SubCustody_ValidateMT548,
							TestInstructionInstanceMajorVersion: VerifyMQMessageTypeMT_VerifyMT548_1_0.TestInstruction_SubCustody_ValidateMT548.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: VerifyMQMessageTypeMT_VerifyMT548_1_0.TestInstruction_SubCustody_ValidateMT548.TestInstruction.MinorVersionNumber,
							Deprecated:                          VerifyMQMessageTypeMT_VerifyMT548_1_0.TestInstruction_SubCustody_ValidateMT548.TestInstruction.Deprecated,
							Enabled:                             VerifyMQMessageTypeMT_VerifyMT548_1_0.TestInstruction_SubCustody_ValidateMT548.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
							ResponseVariablesMapStructure: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
								ResponseVariablesMap:     VerifyMQMessageTypeMT_VerifyMT548_1_0.TestInstructionResponseVariables_SubCustody_ValidateMT548,
								ResponseVariablesMapHash: shared_code.InitialValueBeforeHashed,
							},
							TestInstructionInstanceVersionAndResponseVariablesHash: shared_code.InitialValueBeforeHashed,
						},
					},

					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
			},
			TestInstructionsHash: shared_code.InitialValueBeforeHashed,
		},

		// TestInstructionContainers
		TestInstructionContainers: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{},
		AllowedUsers:              shared_code.AllowedUsersLoadFromJsonFile,
		TestInstructionsAndTestInstructionsContainersAndUsersMessageHash: shared_code.InitialValueBeforeHashed,
		MessageCreationTimeStamp: time.Now(),
		ForceNewBaseLineForTestInstructionsAndTestInstructionContainers: false,
		ConnectorsDomain: TestInstructionAndTestInstuctionContainerTypes.ConnectorsDomainStruct{
			ConnectorsDomainUUID: DomainData.DomainUUID_SubCustody,
			ConnectorsDomainName: DomainData.DomainName_SubCustody,
			ConnectorsDomainHash: shared_code.InitialValueBeforeHashed,
		},
	}

	// Generate TestInstructionContainers

	// TestInstructionContainers

	// Create structure for Local Methods to call in TestApiEngine
	var tempFullTestApiEngineClassesMethodsAttributesVersionMap *TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
	tempFullTestApiEngineClassesMethodsAttributesVersionMap = LocalExecutionMethods.
		InitiateFullTestApiEngineClassesMethodsAttributesMap(TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)

	// Calculate hashes that is included in the Supported TestInstructions, TestInstructionContainers and Allowed Users message
	err = shared_code.CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
		TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody, tempFullTestApiEngineClassesMethodsAttributesVersionMap

}
