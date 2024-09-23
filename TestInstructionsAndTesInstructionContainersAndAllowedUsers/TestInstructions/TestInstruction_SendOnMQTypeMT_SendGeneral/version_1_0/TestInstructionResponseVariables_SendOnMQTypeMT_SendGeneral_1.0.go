package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
)

const ()

const (
	ResponseVariableUuid_SendOnMQTypeMT_SendGeneral_20CSEME TypeAndStructs.ResponseVariableUuidType = "c002af12-114c-4e55-9574-05fae7cf674c"
	ResponseVariableName_SendOnMQTypeMT_SendGeneral_20CSEME TypeAndStructs.ResponseVariableNameType = ":20C::SEME//"
)

// TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral
// Variable that holds the data for the TestInstructionResponseVariables
var TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct

// Initate_TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral
// Function that creates all data for the TestInstructionResponseVariables
func Initate_TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral() {

	// Initiate the TestInstructionResponseVariable
	TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

	TestInstructionResponseVariables_SendOnMQTypeMT_SendGeneral = map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
		ResponseVariableUuid_SendOnMQTypeMT_SendGeneral_20CSEME: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
			ResponseVariable: TypeAndStructs.ResponseVariableStruct{
				ResponseVariableUuid:        ResponseVariableUuid_SendOnMQTypeMT_SendGeneral_20CSEME,
				ResponseVariableName:        ResponseVariableName_SendOnMQTypeMT_SendGeneral_20CSEME,
				ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType("An identifier for the SendOnMQTypeMT_SendGeneral-message"),
				ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(true),
				ResponseVariableTypeUuid:    TestInstructions.ResponseVariableTypeUuid_54x_20CSEME,
				ResponseVariableTypeName:    TestInstructions.ResponseVariableTypeName_54x_20CSEME,
			},
			ResponseVariableHash: shared_code.InitialValueBeforeHashed,
		},
	}

	return
}
