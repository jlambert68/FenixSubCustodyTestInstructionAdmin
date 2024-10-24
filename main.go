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
	fmt.Println("\nSuccess when generating and verifying all messages for TestInstructions, TestInstructionContainers and AllowedUsers")

	var publicKey string
	var err error
	publicKey, err = shared_code.GeneratePublicKeyAsBase64StringFromPrivateKey()
	if err != nil {
		log.Fatalln("\nCouldn't generate 'Public Key' from 'Private Key'")
	}

	fmt.Println(fmt.Sprintf("\nPublic Key from Private key stored in environment variable = %s", publicKey))

	fmt.Println("\n****START**** New Private-Public key par ****START****")
	var newPrivateKey string
	newPrivateKey, err = shared_code.GenerateNewPrivateKeyAsBase64String()
	if err != nil {
		log.Fatalln("Couldn't generate a new 'Private Key'")
	}

	// Create Public key from Private Key
	var newPublicKey string
	newPublicKey, err = shared_code.GeneratePublicKeyAsBase64StringFromPrivateKeyInput(newPrivateKey)
	if err != nil {
		log.Fatalln("Couldn't generate 'Public Key' from new 'Private Key'")
	}

	fmt.Println(fmt.Sprintf("New Private Key = '%s'", newPrivateKey))
	fmt.Println(fmt.Sprintf("New Public Key from new Private key = '%s'", newPublicKey))
	fmt.Println("****END**** New Private-Public key par ****END****")

}
