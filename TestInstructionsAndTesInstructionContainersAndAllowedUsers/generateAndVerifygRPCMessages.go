package TestInstructionsAndTesInstructionContainersAndAllowedUsers

import (
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"os"
)

func GenerateAndVerifyRPCMessages() {

	var err error

	// Worker Server - gRPC-message

	// Convert supported TestInstructions, TestInstructionContainers and Allowed Users message into a gRPC-Worker version of the message
	var supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersGrpcWorkerMessage *fenixExecutionWorkerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersGrpcWorkerMessage, err = shared_code.
		GenerateTestInstructionAndTestInstructionContainerAndUserGrpcWorkerMessage(string(DomainData.DomainUUID_SubCustody),
			TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert back supported TestInstructions, TestInstructionContainers and Allowed Users message from a gRPC-Worker version of the message and check correctness of Hashes
	var testInstructionsAndTestInstructionContainersFromGrpcWorkerMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct
	testInstructionsAndTestInstructionContainersFromGrpcWorkerMessage, err = shared_code.
		GenerateStandardFromGrpcWorkerMessageForTestInstructionsAndUsers(
			supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersGrpcWorkerMessage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Verify recreated Hashes from gRPC-Worker-message
	var errorSliceWorker []error
	errorSliceWorker = shared_code.VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashesAndDomain(
		DomainData.DomainUUID_SubCustody,
		testInstructionsAndTestInstructionContainersFromGrpcWorkerMessage)
	if errorSliceWorker != nil {
		for _, errFromWorker := range errorSliceWorker {
			fmt.Println(errFromWorker)
		}
		os.Exit(1)
	}

	// Builder Server - gRPC-message
	// Convert supported TestInstructions, TestInstructionContainers and Allowed Users message into a gRPC-Builder version of the message
	var supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersGrpcBuilderMessage *fenixTestCaseBuilderServerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersGrpcBuilderMessage, err = shared_code.
		GenerateTestInstructionAndTestInstructionContainerAndUserGrpcBuilderMessage(string(DomainData.DomainUUID_SubCustody),
			TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert back supported TestInstructions, TestInstructionContainers and Allowed Users message from a gRPC-Builder version of the message and check correctness of Hashes
	var testInstructionsAndTestInstructionContainersFromGrpcBuilderMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct
	testInstructionsAndTestInstructionContainersFromGrpcBuilderMessage, err = shared_code.
		GenerateStandardFromGrpcBuilderMessageForTestInstructionsAndUsers(
			supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersGrpcBuilderMessage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Verify recreated Hashes from gRPC-Builder-message
	var errorSliceBuilder []error
	errorSliceBuilder = shared_code.VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashesAndDomain(
		DomainData.DomainUUID_SubCustody,
		testInstructionsAndTestInstructionContainersFromGrpcBuilderMessage)
	if errorSliceBuilder != nil {
		for _, errFromBuilder := range errorSliceBuilder {
			fmt.Println(errFromBuilder)
		}
		os.Exit(1)
	}

}
