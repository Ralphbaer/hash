package gateway

import "context"

// DiscountServiceGateway represents the HTTP comunication With Notifier Service
type DiscountServiceGateway interface {
	GetDiscount(ctx context.Context, productID int32) (float32, error)
}
