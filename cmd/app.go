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
	Svc         service.VendingMachineService
	DataStorage model.Storage
)

/*
	init method to construct this app
*/
func init() {
	DataStorage = model.Storage{
		VendingItems: []model.Item{
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
		},
		VendingCoins: []model.Coin{
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},

			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},

			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
			model.Coin{Value: 10},
		},
	}

	// init vending machine service
	Svc = service.NewVendingService(&DataStorage)

	welcomeMsg := utils.WelcomeMsg()
	displayMsg := utils.Display(DataStorage)

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
