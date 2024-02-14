package LocalExecutionMethods

import (
	"github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
)

// MethodsForLocalExecutionsStruct
// Struct for holding all information of how to execute a TestInstruction
// Even when there are no information about the methods this struct must exist
type MethodsForLocalExecutionsStruct struct {
	LocalParametersUsedInRunTime *LocalParametersUsedInRunTimeStruct                                  `json:"LocalParametersUsedInRunTime"`
	TestInstructionsMap          *TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType `json:"TestInstructionsMap"`
}

var FullTestApiEngineClassesMethodsAttributesVersionMap TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType

// LocalParametersUsedInRunTimeStruct
// Struct for holding the local parameters used during runtime
type LocalParametersUsedInRunTimeStruct struct {
	ExpectedTestInstructionExecutionDurationInSeconds int64 `json:"TimeOutTimeInSeconds"`
}

// Environment variables
var (
	// TestApiEngineAddress
	// Network address to TestApiEngine
	TestApiEngineAddress string

	// TestApiEngineAddressPath
	// URL-path used when calling TestApiEngine
	TestApiEngineAddressPath string
)
