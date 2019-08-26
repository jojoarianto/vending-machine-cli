package cmd

import (
	"bufio"
	"fmt"
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/service"
	"github.com/jojoarianto/vending-machine-cli/utils"
	"os"
)

var ( // global state
	InsertedCoins []model.Coin
	VendingCoins  []model.Coin
	VendingItem   []model.Item
	VendingOutlet []model.Item
	Svc           service.VendingMachineService
)

/*
	init method to construct this app
*/
func init() {
	// init items for sale
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

	// init vending machine service
	Svc = service.NewVendingService(
		InsertedCoins,
		VendingItem,
		VendingCoins,
		VendingOutlet,
		)

	welcomeMsg := utils.WelcomeMsg()
	displayMsg := utils.Display(InsertedCoins, VendingItem)

	fmt.Println(welcomeMsg)
	fmt.Println(displayMsg)
}

/*
	ReadInput method to read interactive terminal input from user
*/
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
