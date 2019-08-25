package utils

import "github.com/jojoarianto/vending-machine-cli/model"

// SumCoin total coin counter
func SumCoin(coins []model.Coin) (totalCoin int64) {
	for _, value := range coins {
		totalCoin += value.Value
	}

	return totalCoin
}

/*
	GiveCoinChanges mengembalikan list array dari total koin untuk di kembalikan
	kembalian hanya dapat berupa coin 10 atau 100
*/
func GiveCoinChanges(totalCoin int64) (coins []model.Coin, left int64) {

	for totalCoin >= 10 {
		if totalCoin >= 100 {
			totalCoin = totalCoin - 100
			coins = append(coins, model.Coin{Value:100})
			continue
		}

		if totalCoin >= 10 {
			totalCoin = totalCoin - 10
			coins = append(coins, model.Coin{Value:10})
			continue
		}
	}

	left = totalCoin
	return
}

