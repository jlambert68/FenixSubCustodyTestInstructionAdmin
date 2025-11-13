package version_1_0

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
)

const (
	ResponseVariableUuid_VerifyReceivedTypeMT5xx_20CSEME TypeAndStructs.ResponseVariableUuidType = "17bdce39-6c5a-4d1b-b0c0-fc2c4800deb3"
	ResponseVariableName_VerifyReceivedTypeMT5xx_20CSEME TypeAndStructs.ResponseVariableNameType = ":20C::SEME//"
)

// TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx
// Variable that holds the data for the TestInstructionResponseVariables
var TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct

// Initate_TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx
// Function that creates all data for the TestInstructionResponseVariables
func Initate_TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx() {

	// Initiate the TestInstructionResponseVariable
	TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

	TestInstructionResponseVariables_SubCustody_VerifyReceivedTypeMT5xx = map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
		ResponseVariableUuid_VerifyReceivedTypeMT5xx_20CSEME: &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
			ResponseVariable: TypeAndStructs.ResponseVariableStruct{
				ResponseVariableUuid:        ResponseVariableUuid_VerifyReceivedTypeMT5xx_20CSEME,
				ResponseVariableName:        ResponseVariableName_VerifyReceivedTypeMT5xx_20CSEME,
				ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType("An identifier for the MT5xx-message"),
				ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(true),
				ResponseVariableTypeUuid:    TestInstructions.ResponseVariableTypeUuid_5xx_20CSEME,
				ResponseVariableTypeName:    TestInstructions.ResponseVariableTypeName_5xx_20CSEME,
			},
			ResponseVariableHash: shared_code.InitialValueBeforeHashed,
		},
	}

	return
}
