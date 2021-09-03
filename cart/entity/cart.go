package entity

// Cart represents a collection of identification data about a Hash Cart,
// including its amount and if its a gift ir not.
// swagger:model Cart
type Cart struct {
	// ID is a unique field
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}
