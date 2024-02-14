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

	// Run 'init()' to load environment variables
	Init()

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

			// Pointer to the Map holding a TestInstructions, with their versions
			var testInstructionsMapPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType

			// Check if TestInstructionUuid exists in map
			testInstructionsMapPtr, existInMap = FullTestApiEngineClassesMethodsAttributesVersionMap[tempOriginalTestInstructionUuid]
			if existInMap == false {

				// TestInstruction does not exist, so it should be created together with this version

				// Extract 'LocalExecutionMethods'
				var tempLocalExecutionMethodsForSingleTestInstructionDefinition *TestInstructionAndTestInstuctionContainerTypes.AnyType
				tempLocalExecutionMethodsForSingleTestInstructionDefinition = &tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods

				// Cast 'AnyType' into type used for mapping towards TestApiEngine
				if tempMethodsForLocalExecutionsForSingleTestInstructionDefinition, ok := tempLocalExecutionMethodsForSingleTestInstructionDefinition.Value.(*MethodsForLocalExecutionsStruct); ok {
					// If ok is true, cast was successful

					// Pointer to the Map holding all TestInstructions, with their versions
					var TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinitionPtr *TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
					TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinitionPtr = tempMethodsForLocalExecutionsForSingleTestInstructionDefinition.TestInstructionsMap

					// Map holding all TestInstructions, with their versions
					var TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
					TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition = make(TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType)

					TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition = *TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinitionPtr

					// Pointer to the Map holding each TestInstructionVersion for the TestInstruction
					var tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinitionPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinitionPtr, existInMap := TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition[tempOriginalTestInstructionUuid]
					if existInMap == false {
						fmt.Println("TestInstruction should not be missing")
						os.Exit(1)
					}

					// Map holding each TestInstructionVersion for the TestInstruction
					var tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinition TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinition = *tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinitionPtr

					// Add back data into map that holds all TestInstructions and all their versions
					FullTestApiEngineClassesMethodsAttributesVersionMap[tempOriginalTestInstructionUuid] = &tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinition

				} else {
					// If ok is false, the cast failed and castedValue will be the zero value of the type
					fmt.Println("Casting failed")
					os.Exit(1)
				}

			} else {
				// TestInstruction does exist, so just add this version to the map

				// Extract 'LocalExecutionMethods'
				var tempLocalExecutionMethodsForSingleTestInstructionDefinition *TestInstructionAndTestInstuctionContainerTypes.AnyType
				tempLocalExecutionMethodsForSingleTestInstructionDefinition = &tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods

				// Cast 'AnyType' into type used for mapping towards TestApiEngine
				if tempMethodsForLocalExecutionsForSingleTestInstructionDefinition, ok := tempLocalExecutionMethodsForSingleTestInstructionDefinition.Value.(*MethodsForLocalExecutionsStruct); ok {
					// If ok is true, cast was successful

					// Pointer to the Map holding all TestInstructions, with their versions
					var TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinitionPtr *TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
					TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinitionPtr = tempMethodsForLocalExecutionsForSingleTestInstructionDefinition.TestInstructionsMap

					// Map holding all TestInstructions, with their versions
					var TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType
					TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition = make(TestApiEngineClassesAndMethodsAndAttributes.TestInstructionsMapType)

					TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition = *TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinitionPtr

					// Pointer to the Map holding each TestInstructionVersion for the TestInstruction
					var tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinitionPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinitionPtr, existInMap := TestInstructionsMapFromTestInstructionVersionForSingleTestInstructionDefinition[tempOriginalTestInstructionUuid]
					if existInMap == false {
						fmt.Println("TestInstruction should not be missing")
						os.Exit(1)
					}

					// Map holding each TestInstructionVersion for the TestInstruction
					var tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinition TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinition = *tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinitionPtr

					// Pointer to the structure holding mapping towards TestApiEngine Classes, Methods, and Attribute namings for a specific TestInstruction-version
					var tempTestApiEngineClassesMethodsAttributesForSingleTestInstructionDefinitionPtr *TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct
					tempTestApiEngineClassesMethodsAttributesForSingleTestInstructionDefinitionPtr, existInMap = tempTestApiEngineClassesMethodsAttributesVersionMapForSingleTestInstructionDefinition[versionNumberAsString]
					if existInMap == false {
						fmt.Println("TestInstructionVersion should not be missing")
						os.Exit(1)
					}

					// structure holding mapping towards TestApiEngine Classes, Methods, and Attribute namings
					var tempTestApiEngineClassesMethodsAttributes TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesStruct
					tempTestApiEngineClassesMethodsAttributes = *tempTestApiEngineClassesMethodsAttributesForSingleTestInstructionDefinitionPtr

					// Map holding a TestInstructions, with their versions
					var testInstructionsMap TestApiEngineClassesAndMethodsAndAttributes.TestApiEngineClassesMethodsAttributesVersionMapType
					testInstructionsMap = *testInstructionsMapPtr

					// Add back data into map that holds all their versions for a specific TestInstruction
					testInstructionsMap[versionNumberAsString] = &tempTestApiEngineClassesMethodsAttributes

				} else {
					// If ok is false, the cast failed and castedValue will be the zero value of the type
					fmt.Println("Casting failed")
					os.Exit(1)
				}
			}
		}
	}
	return
}
