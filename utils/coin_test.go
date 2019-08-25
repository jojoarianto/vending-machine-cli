package utils

import (
	"reflect"
	"testing"

	"github.com/jojoarianto/vending-machine-cli/model"
)

func TestSumCoin(t *testing.T) {
	type args struct {
		coins []model.Coin
	}
	tests := []struct {
		name          string
		args          args
		wantTotalCoin int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotalCoin := SumCoin(tt.args.coins); gotTotalCoin != tt.wantTotalCoin {
				t.Errorf("SumCoin() = %v, want %v", gotTotalCoin, tt.wantTotalCoin)
			}
		})
	}
}

func TestGiveCoinChanges(t *testing.T) {
	type args struct {
		totalCoin int64
	}
	tests := []struct {
		name      string
		args      args
		wantCoins []model.Coin
		wantLeft  int64
	}{
		{
			name: "test coin return 1",
			args: args{totalCoin:40},
			wantCoins: []model.Coin{
				model.Coin{Value:10},
				model.Coin{Value:10},
				model.Coin{Value:10},
				model.Coin{Value:10},
			},
			wantLeft: 0,
		},
		{
			name: "test coin return 2",
			args: args{totalCoin:230},
			wantCoins: []model.Coin{
				model.Coin{Value:100},
				model.Coin{Value:100},
				model.Coin{Value:10},
				model.Coin{Value:10},
				model.Coin{Value:10},
			},
			wantLeft: 0,
		},
		{
			name: "test coin return 3",
			args: args{totalCoin:35},
			wantCoins: []model.Coin{
				model.Coin{Value:10},
				model.Coin{Value:10},
				model.Coin{Value:10},
			},
			wantLeft: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCoins, gotLeft := GiveCoinChanges(tt.args.totalCoin)
			if !reflect.DeepEqual(gotCoins, tt.wantCoins) {
				t.Errorf("GiveCoinChanges() gotCoins = %v, want %v", gotCoins, tt.wantCoins)
			}
			if gotLeft != tt.wantLeft {
				t.Errorf("GiveCoinChanges() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
		})
	}
}
