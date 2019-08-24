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
 Welcome to Vending Maching Application`

	return msg
}

/*
	display untuk menampilkan status dari vending machine
	parameter : coin, item barang
*/
func Display(coin model.Coin, item []model.Item) string {
	msg := ` ----------------------------------
 [Input amount]    %d JPY
 [Change]          100 JPY      No change
                   10 JPY       Change 
 [Return gate]     Empty
 [Items for sale]  
%s
 [Outlet]          Empty

 for list command please reply 'help'
 ----------------------------------`

	itemStr := buildItemList(item)
	msg = fmt.Sprintf(msg, coin.Value, itemStr)
	return msg
}

func buildItemList(item []model.Item) string {
	msg := ""

	// build item
	for key, value := range item {
		msg = msg + fmt.Sprintf(`                   %d. %s (%d JPY) `, (key+1), value.Name, value.CoinValue)
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

