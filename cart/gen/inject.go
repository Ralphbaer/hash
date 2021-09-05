//+build wireinject
//golint:ignore

package gen

import (
	"net/http"

	"github.com/Ralphbaer/hash/cart/app"
	h "github.com/Ralphbaer/hash/cart/handler"
	uc "github.com/Ralphbaer/hash/cart/usecase"
	"github.com/Ralphbaer/hash/common"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

var applicationSet = wire.NewSet(
	common.InitLocalEnvConfig,
	app.NewConfig,
	app.NewRouter,
	app.NewServer,
	wire.Struct(new(uc.CartUseCase), "*"),
	wire.Struct(new(h.CartHandler), "*"),
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
