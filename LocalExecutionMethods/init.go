package LocalExecutionMethods

import (
	"fmt"
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
	"os"
	"strconv"
)

func Init() {
	var err error

	// Extract environment variable for Network address to TestApiEngine
	TestApiEngineAddress = environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("TestApiEngineAddress")

	// Extract environment variable for URL-path used when calling TestApiEngine
	TestApiEngineUrlPath = environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("TestApiEngineUrlPath")

	// Extract environment variable for Port used by TestApiEngine web server
	_, err = strconv.ParseInt(environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("TestApiEnginePort"), 10, 64)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'TestApiEnginePort:' to an integer, error: ", err)
		os.Exit(0)
	}
	TestApiEnginePort = environmentVariables.
		ExtractEnvironmentVariableOrInjectedEnvironmentVariable("TestApiEnginePort")

}
