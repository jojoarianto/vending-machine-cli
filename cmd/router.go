package cmd

import (
	"fmt"
	"os"
	"strings"
)

// Router is orchestration of command input
func Router(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)

	switch arrCommandStr[0] {
	case "1": // insert

		// validasi args must exist
		if len(arrCommandStr) > 1 {

		}

		fmt.Println("insert coin ", arrCommandStr[1] )
		display()
	case "exit":
		os.Exit(0)
	}

	return nil
}
