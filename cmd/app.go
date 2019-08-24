package cmd

import (
	"bufio"
	"fmt"
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/service"
	"os"
)

var (
	Coin model.Coin
	Item model.Item
	Svc  service.VendingMachineService
)

func init() {
	Svc = service.NewInsertService(Coin, Item)

	welcomeMessage()
	display()
	inputCommand()
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

func welcomeMessage() {
	msg := `
========================================
 Welcome to Vending Maching Application
========================================`

	fmt.Println(msg)
}

func display() {
	msg := ` ----------------------------------
 [Input amount]    %d JPY
 [Change]          100 JPY      No change
                   10 JPY       Change 
 [Return gate]     Empty
 [Items for sale]  1. Canned coffee    120 JPY
                   2. Water PET bottle 100 JPY   Sold out
                   3. Sport drinks     150 JPY
 [Outlet]          Empty
 ----------------------------------`
	msg = fmt.Sprintf(msg, Coin.Value)
	fmt.Println(msg)
}

func inputCommand()  {
	msg := ` please input your command : 
     [command number + space + arguments]

 available command :
  1. insert coin
  2. choose item to purchase
  3. get items
  4. return coins
  5. get returned coins
 ----------------------------------`
	fmt.Println(msg)
}
