package service

import (
	"github.com/jojoarianto/vending-machine-cli/model"
	"github.com/jojoarianto/vending-machine-cli/utils"
)

type VendingMachineService interface {
	Insert(newCoin int64) (coin model.Coin, err error)
}

type vendingMachineService struct {
	coin model.Coin
	item []model.Item
}

func NewInsertService(coin model.Coin, item []model.Item) VendingMachineService {
	return &vendingMachineService{coin, item}
}

func (svc *vendingMachineService) Insert(newCoin int64) (coin model.Coin, err error) {
	// cek validate input
	if utils.Validate(newCoin) != true {
		// input coin is not valid
		return svc.coin, err
	}

	svc.coin.Value = svc.coin.Value + newCoin
	return svc.coin, err
}
