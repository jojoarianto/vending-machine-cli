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

		if len(arrCommandStr) <= 1 { // validasi args must exist
			return constant.ErrInputRequired
		}

		in, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return constant.ErrInputInvalid
		}

		err = Svc.Insert(int64(in))
		if err != nil {
			return err
		}

		displayMsg := utils.Display(DataStorage)
		msg := "Your coin insert successfully"

		fmt.Println(msg)
		fmt.Println(displayMsg)

	case "2": // command for purchase

		if len(arrCommandStr) <= 1 { // validasi args must exist
			return constant.ErrInputRequired
		}

		in, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return constant.ErrInputInvalid
		}

		err = Svc.Purchase(int64(in-1))
		if err != nil {
			return err
		}

		displayMsg := utils.Display(DataStorage)
		msg := "Your purchase is successfully"

		fmt.Println(msg)
		fmt.Println(displayMsg)

	case "3": // command for get item

		err := Svc.GetItem()
		if err != nil {
			return err
		}

		displayMsg := utils.Display(DataStorage)
		msg := "Get item is successfully"

		fmt.Println(msg)
		fmt.Println(displayMsg)

	case "4":

		err := Svc.ReturnCoin()
		if err != nil {
			return err
		}

		displayMsg := utils.Display(DataStorage)
		msg := "Get item is successfully"

		fmt.Println(msg)
		fmt.Println(displayMsg)

	case "help":

		msg := utils.HelpMsg()
		fmt.Println(msg)

	case "exit":

		os.Exit(0)

	}

	return nil
}
