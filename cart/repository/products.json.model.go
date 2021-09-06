package repository

import (
	e "github.com/Ralphbaer/hash/cart/entity"
)

// ProductFilter represents a filter utility
type ProductFilter struct {
	IDs []int32
}

// ProductJSONModel is the model of entity.Product
type ProductJSONModel struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}

// ToEntity converts a ProductJSONModel to e.Product
func (d *ProductJSONModel) ToEntity() *e.Product {
	product := &e.Product{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Amount:      d.Amount,
		IsGift:      d.IsGift,
	}

	return product
}

// FromEntity converts an entity.Product to ProductJSONModel
func (d *ProductJSONModel) FromEntity(product *e.Product) error {
	d.ID = product.ID
	d.Title = product.Title
	d.Description = product.Description
	d.Amount = product.Amount
	d.IsGift = product.IsGift

	return nil
}
