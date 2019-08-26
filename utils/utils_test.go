package utils

import (
	"reflect"
	"testing"

	"github.com/jojoarianto/vending-machine-cli/model"
)

func TestValidate(t *testing.T) {
	type args struct {
		coin int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1. Coin valid",
			args: args{coin: 10},
			want: true,
		},
		{
			name: "2. Coin invalid",
			args: args{coin: 40},
			want: false,
		},
		{
			name: "3. Coin valid",
			args: args{coin: 100},
			want: true,
		},
		{
			name: "4. Coin valid",
			args: args{coin: 500},
			want: true,
		},
		{
			name: "5. Coin invalid",
			args: args{coin: 520},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.coin); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			args: args{totalCoin: 40},
			wantCoins: []model.Coin{
				model.Coin{Value: 10},
				model.Coin{Value: 10},
				model.Coin{Value: 10},
				model.Coin{Value: 10},
			},
			wantLeft: 0,
		},
		{
			name: "test coin return 2",
			args: args{totalCoin: 230},
			wantCoins: []model.Coin{
				model.Coin{Value: 100},
				model.Coin{Value: 100},
				model.Coin{Value: 10},
				model.Coin{Value: 10},
				model.Coin{Value: 10},
			},
			wantLeft: 0,
		},
		{
			name: "test coin return 3",
			args: args{totalCoin: 35},
			wantCoins: []model.Coin{
				model.Coin{Value: 10},
				model.Coin{Value: 10},
				model.Coin{Value: 10},
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

func TestCheckChangeExist(t *testing.T) {
	type args struct {
		coins []model.Coin
	}
	tests := []struct {
		name                 string
		args                 args
		wantIsExist10Change  bool
		wantIsExist100Change bool
	}{
		{
			name: "success",
			args: args{
				coins: []model.Coin{
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

					model.Coin{Value: 100},
					model.Coin{Value: 100},
					model.Coin{Value: 100},
					model.Coin{Value: 100},
					model.Coin{Value: 100},
				},
			},
			wantIsExist10Change: true,
			wantIsExist100Change: true,
		},
		{
			name: "success_2",
			args: args{
				coins: []model.Coin{
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

					model.Coin{Value: 100},
					model.Coin{Value: 100},
					model.Coin{Value: 100},
				},
			},
			wantIsExist10Change: true,
			wantIsExist100Change: false,
		},
		{
			name: "success_3",
			args: args{
				coins: []model.Coin{
					model.Coin{Value: 10},

					model.Coin{Value: 100},
					model.Coin{Value: 100},
					model.Coin{Value: 100},
				},
			},
			wantIsExist10Change: false,
			wantIsExist100Change: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsExist10Change, gotIsExist100Change := CheckChangeExist(tt.args.coins)
			if gotIsExist10Change != tt.wantIsExist10Change {
				t.Errorf("CheckChangeExist() gotIsExist10Change = %v, want %v", gotIsExist10Change, tt.wantIsExist10Change)
			}
			if gotIsExist100Change != tt.wantIsExist100Change {
				t.Errorf("CheckChangeExist() gotIsExist100Change = %v, want %v", gotIsExist100Change, tt.wantIsExist100Change)
			}
		})
	}
}


func TestDisplay(t *testing.T) {
	DataStorage := model.Storage{
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
				Qty:       2,
			},
		},
		InsertedCoins: []model.Coin{
			model.Coin{Value: 10},
		},
	}

	type args struct {
		data model.Storage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{data: DataStorage},
			want: ` ----------------------------------
 [Input amount]    10 JPY
 [Change]          100 JPY      No change
                   10 JPY       Change 
 [Return gate]
                   Empty
 [Items for sale]  
                   1. Canned coffee (120 JPY) 
                   2. Water PET bottle (100 JPY) Sold out
                   3. Sport drinks (150 JPY) 

 [Outlet]
                   Empty

 for list command please reply 'help'
 ----------------------------------`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Display(tt.args.data); got != tt.want {
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

func Test_buildOutletList(t *testing.T) {
	type args struct {
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
				item: []model.Item{
					model.Item{
						Name:      "Sport drinks",
						CoinValue: 150,
						Qty:       1,
					},
					model.Item{
						Name:      "Water PET bottle",
						CoinValue: 100,
						Qty:       1,
					},
				},
			},
			want: `                   Sport drinks
                   Water PET bottle
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildOutletList(tt.args.item); got != tt.want {
				t.Errorf("buildOutletList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWelcomeMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WelcomeMsg(); got != tt.want {
				t.Errorf("WelcomeMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildReturnList(t *testing.T) {
	type args struct {
		coin []model.Coin
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildReturnList(tt.args.coin); got != tt.want {
				t.Errorf("buildReturnList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelpMsg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelpMsg(); got != tt.want {
				t.Errorf("HelpMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

