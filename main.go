package main

import (
	_ "embed"
	"fmt"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
)

//go:embed TestInstructionsAndTesInstructionContainersAndAllowedUsers/allowedUsers/allowedUsers.json
var allowedUsers []byte

func main() {

	environmentVariables.InitiateInjectedVariablesMap(&injectedVariablesMap)

	tempTestInstructionsAndTestInstructionsContainers, tempTestApiEngineClassesAndMethodsAndAttributes :=
		TestInstructionsAndTesInstructionContainersAndAllowedUsers.
			GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody(allowedUsers)
	TestInstructionsAndTesInstructionContainersAndAllowedUsers.GenerateAndVerifyRPCMessages()

	fmt.Println(tempTestInstructionsAndTestInstructionsContainers)
	fmt.Println(tempTestApiEngineClassesAndMethodsAndAttributes)
	fmt.Println(TestInstructionsAndTesInstructionContainersAndAllowedUsers.TestInstructionsAndTestInstructionContainersAndAllowedUsers_SubCustody)
	fmt.Println("Success when generating and verifying all messages for TestInstructions, TestInstructionContainers and AllowedUsers")
}
