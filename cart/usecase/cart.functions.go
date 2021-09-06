package usecase

import (
	"context"

	e "github.com/Ralphbaer/hash/cart/entity"
)

// ProductCriteria represents the aggregate information about Product criteria
type ProductCriteria struct {
	ID       int32
	Quantity int32
	Amount   int32
	Discount float32
	IsGift   bool
}

func (uc *CartUseCase) mapProductCriteria(ctx context.Context, p []*e.Product, pr []*ProductRequest) ([]*ProductCriteria, error) {
	pc := make([]*ProductCriteria, len(p))
	for k, v := range p {
		for _, w := range pr {
			if v.ID == w.ID {
				discount, err := uc.Gateway.GetDiscount(ctx, v.ID)
				if err != nil {
					discount = 0
				}
				pc[k] = &ProductCriteria{
					ID:       v.ID,
					Quantity: w.Quantity,
					Amount:   v.Amount,
					Discount: discount,
					IsGift:   v.IsGift,
				}
			}
		}
	}

	return pc, nil
}

func (uc *CartUseCase) calculateTotalAmount(ctx context.Context, pc []*ProductCriteria) int32 {
	result := int32(0)
	for _, v := range pc {
		result += (v.Amount * v.Quantity)
	}

	return result
}

func (uc *CartUseCase) calculateTotalAmountWithDiscount(ctx context.Context, pc []*ProductCriteria) (int32, error) {
	result := int32(0)
	for _, v := range pc {
		for i := 0; i < int(v.Quantity); i++ {
			result += uc.applyDiscount(ctx, v.Discount, v.Amount)
		}
	}

	return result, nil
}

func (uc *CartUseCase) applyDiscount(ctx context.Context, discount float32, productAmount int32) int32 {
	return int32(float32(productAmount) * (1 - discount))
}

func (uc *CartUseCase) mapProductsAmount(ctx context.Context, pc []*ProductCriteria) []*e.ProductAmount {
	productsAmount := make([]*e.ProductAmount, len(pc))
	for k, v := range pc {
		productsAmount[k] = &e.ProductAmount{
			ID:          v.ID,
			Quantity:    v.Quantity,
			UnitAmount:  v.Amount,
			TotalAmount: (v.Quantity) * (v.Amount),
			Discount:    v.Amount - uc.applyDiscount(ctx, v.Discount, v.Amount),
			IsGift:      v.IsGift,
		}
	}

	return productsAmount
}
