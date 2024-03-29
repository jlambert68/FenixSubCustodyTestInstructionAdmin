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

// FullTestApiEngineClassesMethodsAttributesVersionMap
// Holds the reference to information about Local Methods to be used for each version
var FullTestApiEngineClassesMethodsAttributesVersionMap TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType

// LocalParametersUsedInRunTimeStruct
// Struct for holding the local parameters used during runtime
type LocalParametersUsedInRunTimeStruct struct {
	ExpectedTestInstructionExecutionDurationInSeconds int64 `json:"TimeOutTimeInSeconds"`
}
