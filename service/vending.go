package service

import (
	"github.com/jojoarianto/vending-machine-cli/constant"
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/utils"
)

type VendingMachineService interface {
	Insert(newCoin int64) ([]model.Coin, error)
	Purchase(idxItem int64) ([]model.Coin, []model.Item, []model.Coin, []model.Item, error)
}

type vendingMachineService struct {
	insertedCoins []model.Coin
	items         []model.Item
	vendingCoins  []model.Coin
	vendingOutlet []model.Item
}

// NewInsertService service contructor
func NewVendingService(
	insertedCoins []model.Coin,
	items []model.Item,
	vendingCoins []model.Coin,
	vendingOutlet []model.Item,
) VendingMachineService {

	return &vendingMachineService{
		insertedCoins,
		items,
		vendingCoins,
		vendingOutlet,
	}
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
func (svc *vendingMachineService) Purchase(idxItem int64) (
	insertedCoins []model.Coin,
	items []model.Item,
	vendingCoins []model.Coin,
	vendingOutlet []model.Item,
	err error,
) {

	var (
		newItems []model.Item
		userCoin int64
	)

	userCoin = utils.SumCoin(svc.insertedCoins)

	for key, item := range svc.items { // proses untuk menjelajahi semuaitem
		if int64(key) == idxItem { // jika ini ada lah barang yang dicari

			if userCoin < item.CoinValue { // check your coin is enough or not
				err = constant.ErrCoinNotEnough
				return svc.insertedCoins, svc.items, svc.vendingCoins, svc.vendingOutlet, err
			}

			if item.Qty <= 0 {
				err = constant.ErrItemsStockNotAvailable
				return svc.insertedCoins, svc.items, svc.vendingCoins, svc.vendingOutlet, err
			}

			item.Qty -= 1                        // if valid then decreate item qty
			userCoin = userCoin - item.CoinValue // decrease coin user
			newItems = append(newItems, item)    // copy to new variable

			itemToCart := item
			itemToCart.Qty = 1
			svc.vendingOutlet = append(svc.vendingOutlet, itemToCart)

		} else { // jiks bukan barang yang di cari

			newItems = append(newItems, item) // copy to new variable
			continue
		}

	}

	// calculate for change
	newReturnCoins, _ := utils.GiveCoinChanges(userCoin)

	return newReturnCoins, newItems, vendingCoins, svc.vendingOutlet, err
}
