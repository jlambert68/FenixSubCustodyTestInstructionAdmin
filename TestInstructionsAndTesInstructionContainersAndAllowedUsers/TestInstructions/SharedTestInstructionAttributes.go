package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

// *************************************
// *** Shared TestInstruction Attributes ***
const (
	// Attribute 'ExpectedToBePassed' is the parameter sent to CA-execution engine telling if the TestInstruction is expected to pass or fail
	TestInstructionAttributeUUID_SubCustody_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeUUIDType = "43fef553-2424-4a82-ac7e-01db0a9a8e17"
	TestInstructionAttributeName_SubCustody_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeNameType = "ExpectedToBePassed"
	TestInstructionAttributeType_SubCustody_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"

	TestInstructionAttributeName_SubCustody_ExpectedToBePassedComboBox TypeAndStructs.TestInstructionAttributeNameType = "ExpectedToBePassed"
	TestInstructionAttributeType_SubCustody_ExpectedToBePassedComboBox TypeAndStructs.TestInstructionAttributeTypeType = "COMBOBOX"

	TestInstructionAttributeTypeTypeName_SubCustody_ExpectedToBePassed

	// Attribute 'RelatedReference_54x_20CRELA' is the parameter that gets its value from another TestInstructionExecutions output via a Response variable
	TestInstructionAttributeUUID_SubCustody_RelatedReference_54x_20CRELA TypeAndStructs.TestInstructionAttributeUUIDType = "7de76297-9495-41f0-8861-26d03bca3d6a"
	TestInstructionAttributeName_SubCustody_RelatedReference_54x_20CRELA TypeAndStructs.TestInstructionAttributeNameType = "RelatedReference_54x_20CRELA"
	TestInstructionAttributeType_SubCustody_RelatedReference_54x_20CRELA TypeAndStructs.TestInstructionAttributeTypeType = "RESPONSE_VARIABLE_COMBOBOX"

	TestInstructionAttributeTypeUUID_StandardAttributes TypeAndStructs.TestInstructionAttributeTypeUUIDType = "a0161997-4323-45a8-93fa-8ff7d64d1b7a"
	TestInstructionAttributeTypeName_StandardAttributes TypeAndStructs.TestInstructionAttributeTypeNameType = "Standard Attributes"
)
