package constant

import "errors"

var (
	ErrInputRequired          = errors.New("argument input is required")
	ErrInputInvalid           = errors.New("your input is invalid")
	ErrCoinInvalid            = errors.New("Your coin is not valid")
	ErrCoinNotEnough          = errors.New("Your coin is not enough")
	ErrItemsStockNotAvailable = errors.New("Item stock is not available")
)
