package LocalExecutionMethods

import (
	"fmt"
	TestApiEngineClassesAndMethodsAndAttributes "github.com/jlambert68/FenixSubCustodyTestInstructionAdmin/LocalExecutionMethods/TestApiEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"os"
	"strconv"
)

func InitiateFullTestApiEngineClassesMethodsAttributesMap(
	testInstructionsAndTestInstructionContainersAndAllowedUsers *TestInstructionAndTestInstuctionContainerTypes.
		TestInstructionsAndTestInstructionsContainersStruct) {

	// Initiate map
	FullTestApiEngineClassesMethodsAttributesVersionMap = make(TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType)

	//var testApiEngineClassesMethodsAttributesVersionMap TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
	//testApiEngineClassesMethodsAttributesVersionMap = make(TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType)

	// Extract all TestInstructions
	var allTestInstructionsMap map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct
	allTestInstructionsMap = testInstructionsAndTestInstructionContainersAndAllowedUsers.TestInstructions.TestInstructionsMap

	// Version as 'string'
	var versionNumberAsString TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type // "1_0" or "13_3" ...)

	var existInMap bool

	// Loop all TestInstructions
	for tempOriginalTestInstructionUuid, tempTestInstruction := range allTestInstructionsMap {

		// Loop all versions for this TestInstructions
		for _, tempTestInstructionVersion := range tempTestInstruction.TestInstructionVersions {

			// Create the version number as a 'string'
			versionNumberAsString = TestApiEngineClassesAndMethodsAndAttributes.TestApiEngine_MethodNameVersion_SubCustody_Type(
				strconv.Itoa(tempTestInstructionVersion.TestInstructionInstanceMajorVersion) + "_" +
					strconv.Itoa(tempTestInstructionVersion.TestInstructionInstanceMinorVersion))

			// Check if TestInstructionUuid exists in map
			var testInstructionsMap *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
			testInstructionsMap, existInMap = FullTestApiEngineClassesMethodsAttributesVersionMap[tempOriginalTestInstructionUuid]
			if existInMap == false {

				// TestInstruction does not exist so it should be created
				//var testInstructionsMapNew TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
				//testInstructionsMapNew = make(TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType)

				//var testApiEngineClassesMethodsAttributesVersionMapNew TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
				//testApiEngineClassesMethodsAttributesVersionMapNew = make(TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType)

				// Extract 'LocalExecutionMethods'
				var tempLocalExecutionMethods *TestInstructionAndTestInstuctionContainerTypes.AnyType
				tempLocalExecutionMethods = &tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods

				//var testInstructionsMapForTestInstructionVersion *MethodsForLocalExecutionsStruct

				if tempMethodsForLocalExecutions, ok := tempLocalExecutionMethods.Value.(*MethodsForLocalExecutionsStruct); ok {
					// If ok is true, cast was successful
					var TestInstructionsMapFromtestInstructionVersionPtr *TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
					TestInstructionsMapFromtestInstructionVersionPtr = tempMethodsForLocalExecutions.TestInstructionsMap

					var TestInstructionsMapFromtestInstructionVersion TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
					TestInstructionsMapFromtestInstructionVersion = make(TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType)

					TestInstructionsMapFromtestInstructionVersion = *TestInstructionsMapFromtestInstructionVersionPtr

					var tempTestApiEngineClassesMethodsAttributesVersionMapPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					tempTestApiEngineClassesMethodsAttributesVersionMapPtr, existInMap := TestInstructionsMapFromtestInstructionVersion[tempOriginalTestInstructionUuid]
					if existInMap == false {
						fmt.Println("TestInstruction should not be missing")
						os.Exit(1)
					}

					var tempTestApiEngineClassesMethodsAttributesVersionMap TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					tempTestApiEngineClassesMethodsAttributesVersionMap = *tempTestApiEngineClassesMethodsAttributesVersionMapPtr

					var tempTestApiEngineClassesMethodsAttributesPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct
					tempTestApiEngineClassesMethodsAttributesPtr, existInMap = tempTestApiEngineClassesMethodsAttributesVersionMap[versionNumberAsString]
					if existInMap == false {
						fmt.Println("TestInstructionVersion should not be missing")
						os.Exit(1)
					}

					var tempTestApiEngineClassesMethodsAttributes TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct
					tempTestApiEngineClassesMethodsAttributes = *tempTestApiEngineClassesMethodsAttributesPtr
					fmt.Println(tempTestApiEngineClassesMethodsAttributes)

					fmt.Println(testInstructionsMap)

				} else {
					// If ok is false, the cast failed and castedValue will be the zero value of the type
					fmt.Println("Casting failed")
					os.Exit(1)
				}

			} else {

			}

		}
	}

}
