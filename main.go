package main

import (
	_ "embed"
	"fmt"
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"log"
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

	var publicKey string
	var err error
	publicKey, err = shared_code.GeneratePublicKeyAsBase64StringFromPrivateKey()
	if err != nil {
		log.Fatalln("Couldn't generate 'Public Key' from 'Private Key'")
	}

	fmt.Println(fmt.Sprintf("Public Key = %s", publicKey))
}
