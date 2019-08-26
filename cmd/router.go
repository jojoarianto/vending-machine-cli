package cmd

import (
	"fmt"
	"github.com/jojoarianto/vending-machine-cli/constant"
	"github.com/jojoarianto/vending-machine-cli/service"
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

		// renew service
		Svc = service.NewVendingService(
			InsertedCoins,
			VendingItem,
			VendingCoins,
			VendingOutlet,
			)

		InsertedCoins, err = Svc.Insert(int64(in))
		if err != nil {
			return err
		}

		// sent message
		fmt.Println("your coin insert successfully")
		fmt.Println(utils.Display(InsertedCoins, VendingItem))

	case "2": // command for purchase

		if len(arrCommandStr) <= 1 { // validasi args must exist
			return constant.ErrInputRequired
		}

		in, err := strconv.Atoi(arrCommandStr[1])
		if err != nil {
			return constant.ErrInputInvalid
		}

		// renew service
		Svc = service.NewVendingService(
			InsertedCoins,
			VendingItem,
			VendingCoins,
			VendingOutlet,
			)

		InsertedCoins, VendingItem, VendingCoins, VendingOutlet, err = Svc.Purchase(int64(in-1))
		if err != nil {
			return err
		}

		// sent message
		fmt.Println("your purchase is successfully")
		fmt.Println(utils.Display(InsertedCoins, VendingItem))
		fmt.Println(VendingOutlet)

	case "help":

		msg := utils.HelpMsg()
		fmt.Println(msg)

	case "exit":

		os.Exit(0)

	}

	return nil
}
