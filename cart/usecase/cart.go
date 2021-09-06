package usecase

import (
	"context"
	"log"

	e "github.com/Ralphbaer/hash/cart/entity"
	g "github.com/Ralphbaer/hash/cart/gateway"
	r "github.com/Ralphbaer/hash/cart/repository"
)

// CartUseCase represents a collection of use cases for Cart operations
type CartUseCase struct {
	Repository r.ProductRepository
	Gateway    g.DiscountServiceGateway
}

// Create creates a new Cart with a list of products
func (uc *CartUseCase) Create(ctx context.Context, cpi *CreateCartInput) (*e.Cart, error) {
	log.Printf("CartUseCase.Create.CreateCartInput %+v", cpi.Products)

	IDs := make([]int32, len(cpi.Products))
	for k, v := range cpi.Products {
		IDs[k] = v.ID
	}

	products, err := uc.Repository.List(ctx, &r.ProductFilter{IDs: IDs})
	if err != nil {
		return nil, err
	}

	pc, err := uc.mapProductCriteria(ctx, products, cpi.Products)
	if err != nil {
		return nil, err
	}

	cart := &e.Cart{}
	cart.TotalAmount = uc.calculateTotalAmount(ctx, pc)
	totalAmountWithDiscount, err := uc.calculateTotalAmountWithDiscount(ctx, pc)
	if err != nil {
		return nil, err
	}
	cart.TotalAmountWithDiscount = totalAmountWithDiscount
	cart.TotalDiscount = cart.TotalAmount - cart.TotalAmountWithDiscount
	cart.Products = uc.mapProductsAmount(ctx, pc)

	return cart, nil
}
