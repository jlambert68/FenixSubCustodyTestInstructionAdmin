package main

// Variables that can be populated while building the application
// Used when building an executable to be run by itself
var (
	useInjectedEnvironmentVariables string

	Injected_TestApiEngineAddress string
	Injected_TestApiEngineUrlPath string
	Injected_TestApiEnginePort    string
	Injected_PrivateKey           string
)

var falseValue string = "false"

// Map used when Extracting the value of the injected variable
var injectedVariablesMap = map[string]*string{
	// Will Injected values och real Environment Variables be used
	"useInjectedEnvironmentVariables": &falseValue,

	// Environment Variables used for testing "FenixSubCustodyTestInstructionAdmin" by itself
	"Injected_TestApiEngineAddress": &Injected_TestApiEngineAddress,
	"Injected_TestApiEngineUrlPath": &Injected_TestApiEngineUrlPath,
	"Injected_TestApiEnginePort":    &Injected_TestApiEnginePort,
	"Injected_PrivateKey":           &Injected_PrivateKey,
}
