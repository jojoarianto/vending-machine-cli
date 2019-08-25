package utils

import (
	"testing"

	"github.com/jojoarianto/vending-machine-cli/model"
)

func TestDisplay(t *testing.T) {
	type args struct {
		coin []model.Coin
		item []model.Item
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				coin: []model.Coin{
					model.Coin{Value: 10},
				},
				item: []model.Item{
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
						Qty:       2,
					},
				},
			},
			want: ` ----------------------------------
 [Input amount]    10 JPY
 [Change]          100 JPY      No change
                   10 JPY       Change 
 [Return gate]     Empty
 [Items for sale]  
                   1. Canned coffee (120 JPY) 
                   2. Water PET bottle (100 JPY) Sold out
                   3. Sport drinks (150 JPY) 

 [Outlet]          Empty

 for list command please reply 'help'
 ----------------------------------`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Display(tt.args.coin, tt.args.item); got != tt.want {
				t.Errorf("Display() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildItemList(t *testing.T) {
	item := []model.Item{
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
			Qty:       2,
		},
	}

	type args struct {
		coin []model.Coin
		item []model.Item
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success 1",
			args: args{[]model.Coin{
				model.Coin{Value: 10},
			}, item},
			want: `                   1. Canned coffee (120 JPY) 
                   2. Water PET bottle (100 JPY) Sold out
                   3. Sport drinks (150 JPY) 
`,
		},
		{
			name: "success 2",
			args: args{[]model.Coin{
				model.Coin{Value: 500},
			}, item},
			want: `                   1. Canned coffee (120 JPY) Available for purchase
                   2. Water PET bottle (100 JPY) Sold out
                   3. Sport drinks (150 JPY) Available for purchase
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildItemList(tt.args.coin, tt.args.item); got != tt.want {
				t.Errorf("buildItemList() = %v, want %v", got, tt.want)
			}
		})
	}
}
