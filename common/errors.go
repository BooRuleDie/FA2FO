package common

import "errors"

var (
	ErrNoItems = errors.New("items array should have at least one item")
	ErrNoID = errors.New("items id is required")
	ErrInvalidQuantity = errors.New("item quantity must be bigger than zero")
	ErrOutOfStock = errors.New("at least one item is out of stock")
	
)
