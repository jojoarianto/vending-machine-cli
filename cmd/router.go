package cmd

import (
	"fmt"
	"github.com/jojoarianto/vending-machine-cli/constant"
	"github.com/jojoarianto/vending-machine-cli/utils"
	"os"
	"strconv"
	"strings"
)

// Router is orchestration of command input
func Router(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)

	switch arrCommandStr[0] {
	case "1": // command for insert

		// validasi args must exist
		if len(arrCommandStr) <= 1 {
			// return error args not exist
		}

		in, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return constant.ErrInputInvalid
		}

		InsertedCoins, err = Svc.Insert(int64(in))
		if err != nil {
			return err
		}

		// sent message
		fmt.Println("your coin insert successfully")
		fmt.Println(utils.Display(InsertedCoins, Item))
	case "help":
		utils.HelpMsg()
	case "exit":
		os.Exit(0)
	}

	return nil
}
