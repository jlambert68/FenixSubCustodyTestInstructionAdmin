package TestInstructionsAndTesInstructionContainersAndAllowedUsers

import (
	"fmt"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	SendOnMQTypeMT_SendMT540 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540"
	SendOnMQTypeMT_SendMT540_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT540/version_1_0"
	SendOnMQTypeMT_SendMT542 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542"
	SendOnMQTypeMT_SendMT542_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendOnMQTypeMT_SendMT542/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"os"
	"time"
)

var TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody(allowedUsers []byte) {

	var err error

	// Load Allowed users from json-file
	err = shared_code.ParseAllowedUsersFromEmbeddedFile(allowedUsers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set AllUsersAuthorizationRights and initial Hash-value
	shared_code.AllowedUsersLoadFromJsonFile.AllUsersAuthorizationRights = &TestInstructionAndTestInstuctionContainerTypes.AllUsersAuthorizationRightsStruct{
		AllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain:  true,
		AllUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomain: true,
	}
	shared_code.AllowedUsersLoadFromJsonFile.AllowedUsersHash = shared_code.InitialValueBeforeHashed
	// Generate TestInstructions
	// SendOnMQTypeMT::SendMT540
	SendOnMQTypeMT_SendMT540_1_0.Initate_TestInstruction_SubCustody_SendMT540()

	// SendOnMQTypeMT::SendMT542
	SendOnMQTypeMT_SendMT542_1_0.Initate_TestInstruction_SubCustody_SendMT542()

	// Build structure for all TestInstructions & TestInstructionContainers to be sent over gRPC to Fenix Backend
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'SendOnMQTypeMT_SendMT540'
				SendOnMQTypeMT_SendMT540.TestInstructionUUID_SubCustody_SendMT540: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'SendOnMQTypeMT_SendMT540_1.0'
						{
							TestInstructionInstance:             SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540,
							TestInstructionInstanceMajorVersion: SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.MinorVersionNumber,
							Deprecated:                          SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.Deprecated,
							Enabled:                             SendOnMQTypeMT_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
				//TestInstruction 'SendOnMQTypeMT_SendMT542'
				SendOnMQTypeMT_SendMT542.TestInstructionUUID_SubCustody_SendMT542: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'SendOnMQTypeMT_SendMT542_1.0'
						{
							TestInstructionInstance:             SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542,
							TestInstructionInstanceMajorVersion: SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.MinorVersionNumber,
							Deprecated:                          SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.Deprecated,
							Enabled:                             SendOnMQTypeMT_SendMT542_1_0.TestInstruction_SubCustody_SendMT542.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
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

	// Calculate hashes that is included in the Supported TestInstructions, TestInstructionContainers and Allowed Users message
	err = shared_code.CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
		TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
