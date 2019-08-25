package service

import (
	"github.com/jojoarianto/vending-machine-cli/constant"
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/utils"
)

type VendingMachineService interface {
	Insert(newCoin int64) ([]model.Coin, error)
	Purchase(idxItem int64) ([]model.Coin, []model.Item, []model.Coin, error)
}

type vendingMachineService struct {
	insertedCoins []model.Coin
	items         []model.Item
	vendingCoins  []model.Coin
}

func NewInsertService(insertedCoins []model.Coin, items []model.Item, vendingCoins []model.Coin) VendingMachineService {
	return &vendingMachineService{insertedCoins, items, vendingCoins}
}

/*
	Insert method untuk melakukan input inserted coin
*/
func (svc *vendingMachineService) Insert(newCoin int64) ([]model.Coin, error) {
	// cek validate input
	if utils.Validate(newCoin) != true {
		// input coin is not valid
		return svc.insertedCoins, constant.ErrCoinInvalid
	}

	svc.insertedCoins = append(svc.insertedCoins, model.Coin{Value: newCoin})
	return svc.insertedCoins, nil
}

/*
	Purchase method untuk melakukan pembelian barang
*/
func (svc *vendingMachineService) Purchase(idxItem int64)(insertedCoins []model.Coin, items []model.Item, vendingCoins []model.Coin, err error) {
	var (
		newItems []model.Item
		userCoin int64
	)

	userCoin = utils.SumCoin(svc.insertedCoins)

	// find item
	for key, item := range svc.items {

		if int64(key) == idxItem { // jika ini ada lah barang yang di cari

			if userCoin < item.CoinValue { // check your coin is enough or not
				err = constant.ErrCoinNotEnough
				return svc.insertedCoins, svc.items, svc.vendingCoins, err
			}

			if item.Qty <= 0 {
				err = constant.ErrItemsStockNotAvailable
				return svc.insertedCoins, svc.items, svc.vendingCoins, err
			}

			item.Qty -= 1                        // if valid then decreate item qty
			userCoin = userCoin - item.CoinValue // decrease coin user
			newItems = append(newItems, item)    // copy to new variable

		} else { // jiks bukan barang yang di cari

			newItems = append(newItems, item) // copy to new variable
			continue
		}

	}

	// calculate for change
	newReturnCoins, left := utils.GiveCoinChanges(userCoin)
	if left != 0 {
		// return ganjil
	}

	return newReturnCoins, newItems, vendingCoins, err
}
