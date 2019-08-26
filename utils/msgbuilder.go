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
 [Change]          100 JPY      %s
                   10 JPY       %s 
 [Return gate]
%s
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
	returnStr := buildReturnList(storage.ReturnCoins)

	isExist10Change, isExist100Change := CheckChangeExist(storage.VendingCoins)

	change10Str := "No change"
	if isExist10Change {
		change10Str = "Change"
	}

	change100Str := "No change"
	if isExist100Change {
		change100Str = "Change"
	}

	msg = fmt.Sprintf(msg, totalCoin, change100Str, change10Str, returnStr, itemStr, outletStr)
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


func buildReturnList(coin []model.Coin) string {
	msg := ""
	if len(coin) == 0 || coin == nil {
		msg = msg + fmt.Sprintf(`                   Empty`)
		return msg
	}

	for _, value := range coin {
		msg = msg + fmt.Sprintf(`                   %d JPY`, value.Value)
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
