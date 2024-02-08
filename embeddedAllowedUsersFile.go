package main

import (
	_ "embed"
)

//go:embed TestInstructionsAndTesInstructionContainersAndAllowedUsers/allowedUsers/allowedUsers.json
var allowedUsers []byte
