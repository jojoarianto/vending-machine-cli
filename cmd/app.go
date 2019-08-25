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
	InsertedCoins []model.Coin
	VendingCoins  []model.Coin
	VendingItem   []model.Item
	Svc           service.VendingMachineService
)

func init() {
	// build list item for sale
	VendingItem = []model.Item{
		model.Item{
			Name:      "Canned coffee",
			CoinValue: 120,
			Qty:       5,
		},
		model.Item{
			Name:      "Water PET bottle",
			CoinValue: 100,
			Qty:       0,
		},
		model.Item{
			Name:      "Sport drinks",
			CoinValue: 150,
			Qty:       1,
		},
	}

	// init coin vending machine
	VendingCoins = []model.Coin{
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},

		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},

		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
		model.Coin{Value:10},
	}

	Svc = service.NewInsertService(InsertedCoins, VendingItem, VendingCoins)

	fmt.Println(utils.WelcomeMsg())
	fmt.Println(utils.Display(InsertedCoins, VendingItem))
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
