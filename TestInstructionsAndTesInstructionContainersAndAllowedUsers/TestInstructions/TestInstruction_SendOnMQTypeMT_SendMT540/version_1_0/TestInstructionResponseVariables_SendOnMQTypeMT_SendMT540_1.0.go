package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
)

const (
	ResponseVariableUuid_SendMT540_20CSEME TypeAndStructs.ResponseVariableUuidType = "39818ba1-676d-42d0-87da-e1080e9d5ffd"
)

// TestInstructionResponseVariables_SubCustody_SendMT540
// Variable that holds the data for the TestInstructionResponseVariables
var TestInstructionResponseVariables_SubCustody_SendMT540 map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct

// Initate_TestInstructionResponseVariables_SubCustody_SendMT540
// Function that creates all data for the TestInstructionResponseVariables
func Initate_TestInstructionResponseVariables_SubCustody_SendMT540() {

	// Initiate the TestInstructionResponseVariable
	TestInstructionResponseVariables_SubCustody_SendMT540 = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

	TestInstructionResponseVariables_SubCustody_SendMT540 = map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
		ResponseVariableUuid_SendMT540_20CSEME: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
			ResponseVariable: TypeAndStructs.ResponseVariableStruct{
				ResponseVariableUuid:        ResponseVariableUuid_SendMT540_20CSEME,
				ResponseVariableName:        TypeAndStructs.ResponseVariableNameType(":20C::SEME//"),
				ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType("An identifier for the MT540-message"),
				ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(true),
				ResponseVariableTypeUuid:    TestInstructions.ResponseVariableTypeUuid_54x_20CSEME,
				ResponseVariableTypeName:    TestInstructions.ResponseVariableTypeName_54x_20CSEME,
			},
			ResponseVariableHash: shared_code.InitialValueBeforeHashed,
		},
	}

	return
}
