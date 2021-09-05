package usecase

import (
	"context"
	"fmt"
	"log"

	e "github.com/Ralphbaer/hash/cart/entity"
	g "github.com/Ralphbaer/hash/cart/gateway"
)

// CartUseCase represents a collection of use cases for Cart operations
type CartUseCase struct{
	Gateway     g.DiscountServiceGateway
}

// Create creates a new Cart with a list of products
func (uc *CartUseCase) Create(ctx context.Context, cpi *CreateCartInput) (*e.Cart, error) {
	log.Printf("CartUseCase.Create.CreateCartInput %+v", cpi)

	r, err := uc.Gateway.GetDiscount(ctx, 1)
	if err != nil {
		return nil, err
	}

	fmt.Println(r)
	
	return nil, nil
}
