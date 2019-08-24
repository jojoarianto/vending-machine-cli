package model

type Item struct {
	Name      string `json:"name"`
	CoinValue int64  `json:"coin_value"`
	State     string `json:"state"`
	Qty       int64  `json:"qty"`
}
