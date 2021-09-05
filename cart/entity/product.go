package entity

// Product represents a collection of products data,
// including its amount and if its a gift or not.
// swagger:model Product
type Product struct {
	// ID is a unique field
	ID          int  `json:"id"`
	Quantity    int  `json:"quantity"`
	UnitAmount  int  `json:"unit_amount"`
	TotalAmount int  `json:"total_amount"`
	Discount    int  `json:"discount"`
	IsGift      bool `json:"is_gift"`
}
