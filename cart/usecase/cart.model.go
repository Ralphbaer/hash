package usecase

// CreateCartInput is the set of information that will be used to enter data through our handlers.
// We can understand it as a Command. It is used in CREATE operations.
// swagger:model CreateCartInput
type CreateCartInput struct {
	Products []*Product `json:"products"`
}

// Product represents a collection of products data,
// including its amount and if its a gift or not.
// swagger:model Product
type Product struct {
	ID       string `json:"id"`
	Quantity int32  `json:"quantity"`
}
