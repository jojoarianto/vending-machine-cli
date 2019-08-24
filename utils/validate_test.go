package utils

import "testing"

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
			name: "Coin valid",
			args: args{coin: 40},
			want: true,
		},
		{
			name: "Coin ivalid",
			args: args{coin: 5},
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
