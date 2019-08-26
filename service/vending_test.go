package service

import (
	"github.com/jojoarianto/vending-machine-cli/utils"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jojoarianto/vending-machine-cli/model"
)

var (
	DataStorage model.Storage
	Svc         VendingMachineService
)

func setupTesting() {
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
	Svc = NewVendingService(&DataStorage)
}

func TestNewInsertService(t *testing.T) {
	setupTesting()
}

func Test_vendingMachineService_Insert(t *testing.T) {
	var (
		err              error
		ncoin            int64
	)

	setupTesting()

	// insert 10 coin then
	err = Svc.Insert(10)
	ncoin = utils.SumCoin(DataStorage.InsertedCoins)

	// assert
	assert.Equal(t, int64(10), ncoin)
	assert.Nil(t, err)

	// insert 10 coin again
	err = Svc.Insert(10)
	ncoin = utils.SumCoin(DataStorage.InsertedCoins)

	// assert
	assert.Equal(t, int64(20), ncoin)
	ncoin = utils.SumCoin(DataStorage.InsertedCoins)

	assert.Nil(t, err)
}
