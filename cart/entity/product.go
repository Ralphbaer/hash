package entity

// ProductAmount represents a collection of products data,
// all its amounts, discounts and if its a gift or not.
// swagger:model ProductAmount
type ProductAmount struct {
	// ID is a unique field
	ID          int32 `json:"id"`
	Quantity    int32 `json:"quantity"`
	UnitAmount  int64 `json:"unit_amount"`
	TotalAmount int64 `json:"total_amount"`
	Discount    int64 `json:"discount"`
	IsGift      bool  `json:"is_gift"`
}

// Product represents a collection of products data,
// including its amount and if its a gift or not.
// swagger:model Product
type Product struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}
