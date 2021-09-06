package gateway

import "context"

//go:generate mockgen -destination=../gen/mock/gateway_mock.go -package=mock . DiscountServiceGateway

// DiscountServiceGateway represents the HTTP comunication With Notifier Service
type DiscountServiceGateway interface {
	GetDiscount(ctx context.Context, productID int32) (float32, error)
}
