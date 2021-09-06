package usecase

// CreateCartInput is the set of information that will be used to enter data through our handlers.
// We can understand it as a Command. It is used in CREATE operations.
// swagger:model CreateCartInput
type CreateCartInput struct {
	Products []*ProductRequest `json:"products"`
}

// ProductRequest represents the incoming product request data,
// swagger:model ProductRequest
type ProductRequest struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}