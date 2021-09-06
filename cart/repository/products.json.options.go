package repository

import (
	"math/rand"
	"time"

	e "github.com/Ralphbaer/hash/cart/entity"
)

// withIDs filters for given IDs and return its content
var withIDs = func(slice []*ProductJSONModel, IDs []int32) []*e.Product {
	set := make(map[int32]*ProductJSONModel, len(slice))
	for _, s := range slice {
		set[s.ID] = s
	}

	products := make([]*e.Product, 0)
	for _, v := range IDs {
		if _, ok := set[v]; ok {
			products = append(products, set[v].ToEntity())
		}
	}

	return products
}

// withRandom finds a random product inside the Product slice
var withRandom = func(slice []*ProductJSONModel) *e.Product {
	rand.Seed(time.Now().Unix())

	return slice[rand.Intn(len(slice))].ToEntity()
}
