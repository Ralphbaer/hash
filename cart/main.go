package main

import (
	"log"

	"github.com/Ralphbaer/hash/cart/gen"
	"github.com/Ralphbaer/hash/common"
)

func main() {
	common.InitLocalEnvConfig()
	gen.InitializeApp().Run()
	log.Print("cart service terminated")
}
