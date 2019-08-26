package service

import (
	"github.com/jojoarianto/vending-machine-cli/constant"
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/utils"
)

type VendingMachineService interface {
	Insert(newCoin int64) error
	Purchase(idxItem int64) error
	GetItem() error
}

type vendingMachineService struct {
	storage *model.Storage
}

// NewInsertService service contructor
func NewVendingService(stg *model.Storage) VendingMachineService {
	return &vendingMachineService{storage: stg}
}

/*
	Insert method untuk melakukan input inserted coin
*/
func (svc *vendingMachineService) Insert(newCoin int64) error {

	if utils.Validate(newCoin) != true { // cek validate input
		return constant.ErrCoinInvalid // input coin is not valid
	}

	svc.storage.InsertedCoins = append(svc.storage.InsertedCoins, model.Coin{Value: newCoin})
	return nil
}

/*
	Purchase method untuk melakukan pembelian barang
*/
func (svc *vendingMachineService) Purchase(idxItem int64) (error) {

	var (
		newItems []model.Item
		userCoin int64
		err error
	)

	userCoin = utils.SumCoin(svc.storage.InsertedCoins)

	for key, item := range svc.storage.VendingItems { // proses untuk menjelajahi semuaitem
		if int64(key) == idxItem { // jika ini ada lah barang yang dicari

			if userCoin < item.CoinValue { // check your coin is enough or not
				err = constant.ErrCoinNotEnough
				return err
			}

			if item.Qty <= 0 {
				err = constant.ErrItemsStockNotAvailable
				return err
			}

			item.Qty -= 1                        // if valid then decreate item qty
			userCoin = userCoin - item.CoinValue // decrease coin user
			newItems = append(newItems, item)    // copy to new variable

			itemToCart := item
			itemToCart.Qty = 1
			svc.storage.VendingOutlet = append(svc.storage.VendingOutlet, itemToCart)

		} else { // jiks bukan barang yang di cari

			newItems = append(newItems, item) // copy to new variable
			continue
		}

	}

	// calculate for change
	svc.storage.InsertedCoins, _ = utils.GiveCoinChanges(userCoin)
	svc.storage.VendingItems = newItems
	return err
}

/*
	GetItem
*/
func (svc *vendingMachineService) GetItem() error {
	svc.storage.VendingOutlet = nil
	return nil
}