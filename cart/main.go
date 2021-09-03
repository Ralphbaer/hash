package main

import (
	"log"

	"github.com/Ralphbaer/hash/common"
)

func main() {
	common.InitLocalEnvConfig()
	log.Print("cart service terminated")
}
