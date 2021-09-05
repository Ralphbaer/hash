//+build wireinject
//golint:ignore

package gen

import (
	"net/http"

	"github.com/Ralphbaer/hash/cart/app"
	"github.com/Ralphbaer/hash/common"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"github.com/gorilla/mux"

	h "github.com/Ralphbaer/hash/cart/handler"
	uc "github.com/Ralphbaer/hash/cart/usecase"
	g "github.com/Ralphbaer/hash/cart/gateway"
)

func setupDiscountGatewayClient(cfg *app.Config) *grpc.ClientConn {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
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
	setupDiscountGatewayClient,
	g.NewDiscountClient,
	wire.Struct(new(g.DiscountGateway), "*"),
	wire.Struct(new(uc.CartUseCase), "*"),
	wire.Struct(new(h.CartHandler), "*"),
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
