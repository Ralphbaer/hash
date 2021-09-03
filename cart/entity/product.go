package entity

// Product represents a collection of products data,
// including its amount and if its a gift or not.
// swagger:model Product
type Product struct {
	// ID is a unique field
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}
