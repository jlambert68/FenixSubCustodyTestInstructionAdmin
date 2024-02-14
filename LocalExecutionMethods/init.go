package LocalExecutionMethods

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func Init() {
	var err error

	// Extract environment variable for Network address to TestApiEngine
	TestApiEngineAddress = mustGetenv("TestApiEngineAddress")

	// Extract environment variable for URL-path used when calling TestApiEngine
	TestApiEngineUrlPath = mustGetenv("TestApiEngineUrlPath")

	// Extract environment variable for Port used by TestApiEngine web server
	_, err = strconv.ParseInt(mustGetenv("TestApiEnginePort"), 10, 64)
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'TestApiEnginePort:' to an integer, error: ", err)
		os.Exit(0)
	}
	TestApiEnginePort = mustGetenv("TestApiEnginePort")

}
