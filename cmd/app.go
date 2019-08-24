package cmd

import (
	"bufio"
	"fmt"
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/service"
	"github.com/jojoarianto/vending-machine-cli/utils"
	"os"
)

var (
	Coin model.Coin
	Item []model.Item
	Svc  service.VendingMachineService
)

func init() {
	// build list item for sale
	Item = []model.Item{
		model.Item{
			Name: "Canned coffee",
			CoinValue: 120,
			Qty: 5,
		},
		model.Item{
			Name: "Water PET bottle",
			CoinValue: 100,
			Qty: 0,
		},
		model.Item{
			Name: "Sport drinks",
			CoinValue: 150,
			Qty: 2,
		},
	}

	Svc = service.NewInsertService(Coin, Item)

	fmt.Println(utils.WelcomeMsg())
	fmt.Println(utils.Display(Coin, Item))
}

func ReadInput() {
	// read input from typing
	reader := bufio.NewReader(os.Stdin)

	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if len(cmdString) <= 1 {
			// for handling empty command
			continue
		}

		// sent to router
		err = Router(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}