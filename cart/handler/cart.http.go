package handler

import (
	"log"
	"net/http"

	uc "github.com/Ralphbaer/hash/cart/usecase"
	commonHTTP "github.com/Ralphbaer/hash/common/net/http"
)

// CartHandler represents a handler which deal with Cart resource operations
type CartHandler struct {
	UseCase *uc.CartUseCase
}

// Create creates a Cart
// swagger:operation POST /carts carts Create
// Creates a Cart and applies a discount to each product (if a discount is available).
// ---
// parameters:
//  -  name: input
//     in: body
//     type: string
//     description: The payload
//     required: true
//     schema:
//       "$ref": "#/definitions/CreateCartInput"
//
// security:
//  - Definitions: []
//
// responses:
//   '200':
//     description: Success Operation
//     schema:
//       "$ref": "#/definitions/Cart"
//   '400':
//     description: Invalid Input - Input has invalid/missing values
//     schema:
//       "$ref": "#/definitions/ValidationError"
//     examples:
//       "application/json":
//         code: 400
//         message: message
func (handler *CartHandler) Create(p interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := p.(*uc.CreateCartInput)

		log.Printf("Trying to mount a Cart with payload %+v", payload)

		cart, err := handler.UseCase.Create(r.Context(), payload)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		commonHTTP.OK(w, cart)
	})
}
