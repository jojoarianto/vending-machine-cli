package service

import (
	"github.com/jojoarianto/vending-machine-cli/utils"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jojoarianto/vending-machine-cli/model"
)

var (
	InsertedCoins []model.Coin
	VendingCoins  []model.Coin
	VendingItem   []model.Item
	Svc           VendingMachineService
)

func setupTesting() {
	// init item for sale
	VendingItem = []model.Item{
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
	}

	// init coin vending machine
	VendingCoins = []model.Coin{
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
	}

	Svc = NewInsertService(InsertedCoins, VendingItem, VendingCoins)
}

func TestNewInsertService(t *testing.T) {
	setupTesting()
}

func Test_vendingMachineService_Insert(t *testing.T) {
	var (
		newInsertedCoins []model.Coin
		err              error
		ncoin            int64
	)

	setupTesting()

	// insert 10 coin then
	newInsertedCoins, err = Svc.Insert(10)
	ncoin = utils.SumCoin(newInsertedCoins)

	// assert
	assert.Equal(t, int64(10), ncoin)
	assert.Nil(t, err)

	// insert 10 coin again
	newInsertedCoins, err = Svc.Insert(10)
	ncoin = utils.SumCoin(newInsertedCoins)

	// assert
	assert.Equal(t, int64(20), ncoin)
	assert.NotNil(t, newInsertedCoins)

	assert.Nil(t, err)
}
