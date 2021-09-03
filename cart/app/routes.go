// Package app Cart API.
//
// This guide describes all Hash Cart API and usecase.
//
//     Schemes: http, https
//     BasePath: /cart
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package app

import (
	"github.com/Ralphbaer/hash/cart/handler"
	uc "github.com/Ralphbaer/hash/cart/usecase"
	lib "github.com/Ralphbaer/hash/common/net/http"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(c *handler.CartHandler) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	r.Use(lib.WithCorrelationID)

	r.Handle("/cart/carts", lib.WithBody(new(uc.CreateCartInput), c.Create)).Methods("POST")

	// Common

	r.HandleFunc("/cart/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "cart", "Cart API Documentation", r)

	return r
}
