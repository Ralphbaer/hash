package repository

import (
	"context"

	e "github.com/Ralphbaer/hash/cart/entity"
)

//go:generate mockgen -destination=../gen/mock/repository_mock.go -package=mock . ProductRepository

// ProductRepository manages product repository operations
type ProductRepository interface {
	List(ctx context.Context, f *ProductFilter) ([]*e.Product, error)
	FindRandom(ctx context.Context) (*e.Product, error)
}
