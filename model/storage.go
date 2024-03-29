package model

type Storage struct {
	InsertedCoins []Coin `json:"inserted_coins"`
	ReturnCoins   []Coin `json:"return_coins"`
	VendingCoins  []Coin `json:"vending_coins"`
	VendingItems  []Item `json:"vending_items"`
	VendingOutlet []Item `json:"vending_outlet"`
}
