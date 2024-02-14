package LocalExecutionMethods

import (
	"log"
	"os"
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
	// Extract environment variable for Network address to TestApiEngine
	TestApiEngineAddress = mustGetenv("TestApiEngineAddress")

	// Extract environment variable for URL-path used when calling TestApiEngine
	TestApiEngineAddressPath = mustGetenv("TestApiEngineAddressPath")
}
