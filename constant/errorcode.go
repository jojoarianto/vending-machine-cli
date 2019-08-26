package constant

import "errors"

var (
	ErrInputRequired          = errors.New("Argument input is required")
	ErrInputInvalid           = errors.New("Your input is invalid")
	ErrCoinInvalid            = errors.New("Your coin is not valid")
	ErrCoinNotEnough          = errors.New("Your coin is not enough")
	ErrItemsStockNotAvailable = errors.New("Item stock is not available")
)
