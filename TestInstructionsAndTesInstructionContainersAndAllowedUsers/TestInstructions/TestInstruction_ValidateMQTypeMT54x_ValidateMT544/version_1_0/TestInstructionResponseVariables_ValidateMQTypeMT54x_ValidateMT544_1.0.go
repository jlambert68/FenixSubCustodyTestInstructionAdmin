package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
)

const (
	ResponseVariableUuid_ValidateMT544_20CSEME TypeAndStructs.ResponseVariableUuidType = "39818ba1-676d-42d0-87da-e1080e9d5ffd"
	ResponseVariableName_ValidateMT544_20CSEME TypeAndStructs.ResponseVariableNameType = ":20C::SEME//"
)

// TestInstructionResponseVariables_SubCustody_ValidateMT544
// Variable that holds the data for the TestInstructionResponseVariables
var TestInstructionResponseVariables_SubCustody_ValidateMT544 map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct

// Initate_TestInstructionResponseVariables_SubCustody_ValidateMT544
// Function that creates all data for the TestInstructionResponseVariables
func Initate_TestInstructionResponseVariables_SubCustody_ValidateMT544() {

	// Initiate the TestInstructionResponseVariable
	TestInstructionResponseVariables_SubCustody_ValidateMT544 = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

	TestInstructionResponseVariables_SubCustody_ValidateMT544 = map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
		ResponseVariableUuid_ValidateMT544_20CSEME: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
			ResponseVariable: TypeAndStructs.ResponseVariableStruct{
				ResponseVariableUuid:        ResponseVariableUuid_ValidateMT544_20CSEME,
				ResponseVariableName:        ResponseVariableName_ValidateMT544_20CSEME,
				ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType("An identifier for the MT544-message"),
				ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(true),
				ResponseVariableTypeUuid:    TestInstructions.ResponseVariableTypeUuid_54x_20CSEME,
				ResponseVariableTypeName:    TestInstructions.ResponseVariableTypeName_54x_20CSEME,
			},
			ResponseVariableHash: shared_code.InitialValueBeforeHashed,
		},
	}

	return
}
