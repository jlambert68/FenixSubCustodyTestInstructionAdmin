package LocalExecutionMethods

import "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"

// MethodsForLocalExecutionsStruct
// Struct for holding all information of how to execute a TestInstruction
// Even when there are no information about the methods this struct must exist
type MethodsForLocalExecutionsStruct struct {
	LocalParametersUsedInRunTime          *LocalParametersUsedInRunTimeStruct                                         `json:"LocalParametersUsedInRunTime"`
	TestApiEngineClassesMethodsAttributes *TestApiEngineClassesAndMethods.TestApiEngineClassesMethodsAttributesStruct `json:"TestApiEngineClassesMethodsAttributes"`
}

// LocalParametersUsedInRunTimeStruct
// Struct for holding the local parameters used during runtime
type LocalParametersUsedInRunTimeStruct struct {
	ExpectedTestInstructionExecutionDurationInSeconds int64 `json:"TimeOutTimeInSeconds"`
}
