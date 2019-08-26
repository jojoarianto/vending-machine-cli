package utils

import (
	"fmt"
	"github.com/jojoarianto/vending-machine-cli/model"
)

/*
	WelcomeMsg untuk menampilakn pesan selamat datang
*/
func WelcomeMsg() string {
	msg := `
 ----------------------------------
 Welcome to Vending Machine Application`

	return msg
}

/*
	display untuk menampilkan status dari vending machine
	parameter : coin, item barang
*/
func Display(storage model.Storage) string {
	msg := ` ----------------------------------
 [Input amount]    %d JPY
 [Change]          100 JPY      No change
                   10 JPY       Change 
 [Return gate]     Empty
 [Items for sale]  
%s
 [Outlet]
%s

 for list command please reply 'help'
 ----------------------------------`

	var totalCoin int64
	for _, value := range storage.InsertedCoins {
		totalCoin += value.Value
	}

	itemStr := buildItemList(storage.InsertedCoins, storage.VendingItems)
	outletStr := buildOutletList(storage.VendingOutlet)

	msg = fmt.Sprintf(msg, totalCoin, itemStr, outletStr)
	return msg
}

func buildItemList(coins []model.Coin, item []model.Item) string {
	userCoin := SumCoin(coins)
	msg := ""

	// build item
	for key, value := range item {
		msg = msg + fmt.Sprintf(`                   %d. %s (%d JPY) `, (key+1), value.Name, value.CoinValue)
		if value.Qty == 0 {
			msg = msg + "Sold out"
		}

		if userCoin >= value.CoinValue && value.Qty > 0 {
			msg = msg + "Available for purchase"
		}
		msg = msg + "\n"
	}

	return msg
}

func buildOutletList(item []model.Item) string {
	msg := ""
	if len(item) == 0 || item == nil {
		msg = msg + fmt.Sprintf(`                   Empty`)
		return msg
	}

	for _, value := range item {
		msg = msg + fmt.Sprintf(`                   %s`, value.Name)
		if value.Qty == 0 {
			msg = msg + "Sold out"
		}

		msg = msg + "\n"
	}

	return msg
}

/*
	helpMessage
*/
func HelpMsg() string {
	msg := ` ----------------------------------
 please input your command : 
     [command number + space + arguments]

 available command :
  1. insert coin
  2. choose item to purchase
  3. get items
  4. return coins
  5. get returned coins
 ----------------------------------`

	return msg
}

