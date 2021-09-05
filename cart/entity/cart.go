package entity

// Cart represents a collection of Cart data,
// including its amount and if its a gift or not.
// swagger:model Cart
type Cart struct {
	TotalAmount             string `json:"total_amount"`
	TotalAmountWithDiscount string `json:"total_amount_with_discount"`
	TotalDiscount           string `json:"total_discount"`
	Products                []*Product    `json:"amount"`
}
