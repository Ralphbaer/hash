package usecase

import "errors"

// ErrCantBuyAGift is throwed when the client tries to create a cart with a pre-load gift product
var ErrCantBuyAGift = errors.New("you cannot buy a gift product")
