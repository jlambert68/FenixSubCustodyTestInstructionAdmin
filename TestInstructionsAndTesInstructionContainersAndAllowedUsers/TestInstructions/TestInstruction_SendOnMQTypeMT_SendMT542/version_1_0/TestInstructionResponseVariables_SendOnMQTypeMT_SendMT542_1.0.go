package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
)

const (
	ResponseVariableUuid_SendMT542_20CSEME TypeAndStructs.ResponseVariableUuidType = "9dd57f25-75e0-4024-862b-e0728c066604"
)

// TestInstructionResponseVariables_SubCustody_SendMT542
// Variable that holds the data for the TestInstructionResponseVariables
var TestInstructionResponseVariables_SubCustody_SendMT542 map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct

// Initate_TestInstructionResponseVariables_SubCustody_SendMT542
// Function that creates all data for the TestInstructionResponseVariables
func Initate_TestInstructionResponseVariables_SubCustody_SendMT542() {

	// Initiate the TestInstructionResponseVariable
	TestInstructionResponseVariables_SubCustody_SendMT542 = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

	TestInstructionResponseVariables_SubCustody_SendMT542 = map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
		ResponseVariableUuid_SendMT542_20CSEME: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
			ResponseVariable: TypeAndStructs.ResponseVariableStruct{
				ResponseVariableUuid:        ResponseVariableUuid_SendMT542_20CSEME,
				ResponseVariableName:        TypeAndStructs.ResponseVariableNameType(":20C::SEME//"),
				ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType("An identifier for the MT542-message"),
				ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(true),
				ResponseVariableTypeUuid:    TestInstructions.ResponseVariableTypeUuid_54x_20CSEME,
				ResponseVariableTypeName:    TestInstructions.ResponseVariableTypeName_54x_20CSEME,
			},
			ResponseVariableHash: shared_code.InitialValueBeforeHashed,
		},
	}

	return
}
