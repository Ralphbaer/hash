//go:build wireinject
// +build wireinject

//golint:ignore

package gen

import (
	"net/http"

	"github.com/Ralphbaer/hash/cart/app"
	"github.com/Ralphbaer/hash/common"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	g "github.com/Ralphbaer/hash/cart/gateway"
	h "github.com/Ralphbaer/hash/cart/handler"
	r "github.com/Ralphbaer/hash/cart/repository"
	uc "github.com/Ralphbaer/hash/cart/usecase"
)

func setupDiscountGatewayClient(cfg *app.Config) *grpc.ClientConn {
	conn, err := grpc.Dial("hash-discount-service:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}

var applicationSet = wire.NewSet(
	common.InitLocalEnvConfig,
	app.NewConfig,
	app.NewRouter,
	app.NewServer,
	r.NewProductJSONRepository,
	setupDiscountGatewayClient,
	g.NewDiscountClient,
	wire.Struct(new(g.DiscountGateway), "*"),
	wire.Struct(new(uc.CartUseCase), "*"),
	wire.Struct(new(h.CartHandler), "*"),
	wire.Bind(new(r.ProductRepository), new(*r.ProductJSONRepository)),
	wire.Bind(new(g.DiscountServiceGateway), new(*g.DiscountGateway)),
	wire.Bind(new(http.Handler), new(*mux.Router)),
)

// InitializeApp setup the dependencies and returns a new *app.App instance
func InitializeApp() *app.App {
	wire.Build(
		applicationSet,
		wire.Struct(new(app.App), "*"),
	)
	return nil
}
