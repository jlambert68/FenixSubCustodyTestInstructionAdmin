package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
)

const (
	ResponseVariableUuid_ValidateMT548_20CSEME TypeAndStructs.ResponseVariableUuidType = "8ed1ead9-741b-4115-9f78-f8a7db1d6274"
	ResponseVariableName_ValidateMT548_20CSEME TypeAndStructs.ResponseVariableNameType = ":20C::SEME//"
)

// TestInstructionResponseVariables_SubCustody_ValidateMT548
// Variable that holds the data for the TestInstructionResponseVariables
var TestInstructionResponseVariables_SubCustody_ValidateMT548 map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct

// Initate_TestInstructionResponseVariables_SubCustody_ValidateMT548
// Function that creates all data for the TestInstructionResponseVariables
func Initate_TestInstructionResponseVariables_SubCustody_ValidateMT548() {

	// Initiate the TestInstructionResponseVariable
	TestInstructionResponseVariables_SubCustody_ValidateMT548 = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

	TestInstructionResponseVariables_SubCustody_ValidateMT548 = map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
		ResponseVariableUuid_ValidateMT548_20CSEME: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
			ResponseVariable: TypeAndStructs.ResponseVariableStruct{
				ResponseVariableUuid:        ResponseVariableUuid_ValidateMT548_20CSEME,
				ResponseVariableName:        ResponseVariableName_ValidateMT548_20CSEME,
				ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType("An identifier for the MT548-message"),
				ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(true),
				ResponseVariableTypeUuid:    TestInstructions.ResponseVariableTypeUuid_54x_20CSEME,
				ResponseVariableTypeName:    TestInstructions.ResponseVariableTypeName_54x_20CSEME,
			},
			ResponseVariableHash: shared_code.InitialValueBeforeHashed,
		},
	}

	return
}
