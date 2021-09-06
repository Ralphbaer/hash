package entity

// Cart represents a collection of Cart data,
// including its amount and if its a gift or not.
// swagger:model Cart
type Cart struct {
	TotalAmount             int32            `json:"total_amount"`
	TotalAmountWithDiscount int32            `json:"total_amount_with_discount"`
	TotalDiscount           int32            `json:"total_discount"`
	Products                []*ProductAmount `json:"products"`
}
