package usecase

import (
	"context"
	"log"

	e "github.com/Ralphbaer/hash/cart/entity"
)

// CartUseCase represents a collection of use cases for Cart operations
type CartUseCase struct{}

// Create creates a new Cart with a list of products
func (uc *CartUseCase) Create(ctx context.Context, cpi *CreateCartInput) (*e.Cart, error) {
	log.Printf("CartUseCase.Create.CreateCartInput %+v", cpi)

	return nil, nil
}
