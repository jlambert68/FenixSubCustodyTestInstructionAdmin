package TestInstructionsAndTesInstructionContainersAndAllowedUsers

import (
	"fmt"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	testInstructionContainer_SpecialSerialBaseContainer "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	testInstructionContainer_SpecialSerialBaseContainer_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer/version_1_0"
	generalSetupTearDown_SendMT540 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_SendMT540"
	generalSetupTearDown_SendMT540_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_SendMT540/version_1_0"
	generalSetupTearDown_SendMT540_1_1 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_SendMT540/version_1_1"
	generalSetupTearDown_TestCaseTearDown "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	generalSetupTearDown_TestCaseTearDown_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/version_1_0"
	isDateAPublicHoliday "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsDateAPublicHoliday"
	isDateAPublicHoliday_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsDateAPublicHoliday/version_1_0"
	isServerAlive "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive"
	isServerAlive_v_1_0 "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"os"
	"time"
)

var TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody() {

	var err error

	// Load Allowed users from json-file
	err = shared_code.ImportAllowedUsersFromFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate TestInstructions
	// GeneralSetupTearDown::SendMT540
	generalSetupTearDown_SendMT540_1_1.Initate_TestInstruction_SubCustody_SendMT540()
	generalSetupTearDown_SendMT540_1_0.Initate_TestInstruction_SubCustody_SendMT540()

	// GeneralSetupTearDown::TestCaseTearDown
	generalSetupTearDown_TestCaseTearDown_1_0.Initate_TestInstruction_SubCustody_TestCaseTearDown()

	// Standard::IsDateAPublicHoliday
	isDateAPublicHoliday_1_0.Initate_TestInstruction_SubCustody_IsDateAPublicHoliday()

	// Standard::IsServerAlive
	isServerAlive_v_1_0.Initate_TestInstruction_SubCustody_IsServerAlive()

	// Build structure for all TestInstructions & TestInstructionContainers to be sent over gRPC to Fenix Backend
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'generalSetupTearDown_SendMT540'
				generalSetupTearDown_SendMT540.TestInstructionUUID_SubCustody_SendMT540: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						//Version 'generalSetupTearDown_SendMT540_1.1'
						{
							TestInstructionInstance:             generalSetupTearDown_SendMT540_1_1.TestInstruction_SubCustody_SendMT540,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_SendMT540_1_1.TestInstruction_SubCustody_SendMT540.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_SendMT540_1_1.TestInstruction_SubCustody_SendMT540.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_SendMT540_1_1.TestInstruction_SubCustody_SendMT540.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_SendMT540_1_1.TestInstruction_SubCustody_SendMT540.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},

						// Version 'generalSetupTearDown_SendMT540_1.0'
						{
							TestInstructionInstance:             generalSetupTearDown_SendMT540_1_0.TestInstruction_SubCustody_SendMT540,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_SendMT540_1_0.TestInstruction_SubCustody_SendMT540.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				// TestInstruction 'generalSetupTearDown_SendMT540'
				generalSetupTearDown_TestCaseTearDown.TestInstructionUUID_SubCustody_TestCaseTearDown: {
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'generalSetupTearDown_SendMT540_1.0'
						{
							TestInstructionInstance:             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SubCustody_TestCaseTearDown,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SubCustody_TestCaseTearDown.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SubCustody_TestCaseTearDown.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SubCustody_TestCaseTearDown.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_SubCustody_TestCaseTearDown.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				// TestInstruction 'IsDateAPublicHoliday'
				isDateAPublicHoliday.TestInstructionUUID_SubCustody_IsDateAPublicHoliday: {
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'isDateAPublicHoliday_1_0'
						{
							TestInstructionInstance:             isDateAPublicHoliday_1_0.TestInstruction_SubCustody_IsDateAPublicHoliday,
							TestInstructionInstanceMajorVersion: isDateAPublicHoliday_1_0.TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: isDateAPublicHoliday_1_0.TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstruction.MinorVersionNumber,
							Deprecated:                          isDateAPublicHoliday_1_0.TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstruction.Deprecated,
							Enabled:                             isDateAPublicHoliday_1_0.TestInstruction_SubCustody_IsDateAPublicHoliday.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				// TestInstruction 'IsServerAlive'
				isServerAlive.TestInstructionUUID_SubCustody_IsServerAlive: {
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'IsServerAlive_1_0'
						{
							TestInstructionInstance:             isServerAlive_v_1_0.TestInstruction_SubCustody_IsServerAlive,
							TestInstructionInstanceMajorVersion: isServerAlive_v_1_0.TestInstruction_SubCustody_IsServerAlive.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: isServerAlive_v_1_0.TestInstruction_SubCustody_IsServerAlive.TestInstruction.MinorVersionNumber,
							Deprecated:                          isServerAlive_v_1_0.TestInstruction_SubCustody_IsServerAlive.TestInstruction.Deprecated,
							Enabled:                             isServerAlive_v_1_0.TestInstruction_SubCustody_IsServerAlive.TestInstruction.Enabled,
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
	// testInstructionContainer_SpecialSerialBaseContainer
	testInstructionContainer_SpecialSerialBaseContainer_1_0.Initiate_TestInstructionContainer_SubCustody_Serial(TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)

	// TestInstructionContainers
	var testInstructionContainers *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct
	testInstructionContainers = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{
		TestInstructionContainersMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{

			// 'SC_SpecialSerialBaseContainer'
			testInstructionContainer_SpecialSerialBaseContainer.TestInstructionContainerUUID_SubCustody_SpecialSerialBaseContainer: {
				TestInstructionContainerVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{

					//Version 'TestInstructionContainer_SpecialSerialBaseContainer_1.0'
					{
						TestInstructionContainerInstance:             testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SubCustody_SpecialSerialBase,
						TestInstructionContainerInstanceMajorVersion: testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SubCustody_SpecialSerialBase.TestInstructionContainer.MajorVersionNumber,
						TestInstructionContainerInstanceMinorVersion: testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SubCustody_SpecialSerialBase.TestInstructionContainer.MinorVersionNumber,
						Deprecated:                           testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SubCustody_SpecialSerialBase.TestInstructionContainer.Deprecated,
						Enabled:                              testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_SubCustody_SpecialSerialBase.TestInstructionContainer.Enabled,
						TestInstructionContainerInstanceHash: shared_code.InitialValueBeforeHashed,
					},
				},
				TestInstructionContainerVersionsHash: shared_code.InitialValueBeforeHashed,
			},
		},

		TestInstructionContainersHash: shared_code.InitialValueBeforeHashed,
	}
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody.TestInstructionContainers = testInstructionContainers

	// Define temporary storage for 'LocalExecutionMethods.MethodsForLocalExecutionsStruct'
	/*
		var tempStorageLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType


		var PushToTempStoreFunction = func(incomingLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType) {
			tempStorageLocalExecutionMethodsObject = incomingLocalExecutionMethodsObject

		}

		var PullFromTempStoreFunction = func() (outgoingLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType) {
			return tempStorageLocalExecutionMethodsObject
		}

	*/
	//type PushToTempStoreFunctionType[ T, any] func(T)
	//var PushToTempStore = shared_code.PushToTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType](PushToTempStoreFunction)
	//var PullFromTempStore = shared_code.PullFromTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType](PullFromTempStoreFunction)

	// Calculate hashes that is included in the Supported TestInstructions, TestInstructionContainers and Allowed Users message
	err = shared_code.CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
		TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
