package gateway

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Ralphbaer/hash/grpc/discount"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// DiscountGateway abstracts the attributes of DepositClient
type DiscountGateway struct {
	*DepositClient
}
 
// DepositClient represents the client to invoke discount service
type DepositClient struct {
	conn    *grpc.ClientConn
}

// NewDiscountClient creates an instance of DepositClient
func NewDiscountClient(conn *grpc.ClientConn) *DepositClient {
	return &DepositClient{
		conn:    conn,
	}
}
 
// GetDiscount creates a new Notification and return it's ID
func (g *DiscountGateway) GetDiscount(ctx context.Context, productID int32) (float32, error) {
	client := pb.NewDiscountClient(g.conn)
 
	request := &pb.GetDiscountRequest{ProductID: productID}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()
 
	response, err := client.GetDiscount(ctx, request)
	if err != nil {
		if er, ok := status.FromError(err); ok {
			return 0, fmt.Errorf("grpc: %s, %s", er.Code(), er.Message())
		}
		return 0, fmt.Errorf("server: %s", err.Error())
	}
 
	return response.GetPercentage(), nil
}
