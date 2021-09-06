package repository

import (
	e "github.com/Ralphbaer/hash/cart/entity"
)

// withIDs is a criteria Ids filter
var withIDs = func(slice []*ProductJSONModel, item []int32) []*e.Product {
	set := make(map[int32]*ProductJSONModel, len(slice))
	for _, s := range slice {
		set[s.ID] = s
	}

	products := make([]*e.Product, 0)
	for _, v := range item {
		if _, ok := set[v]; ok {
			products = append(products, set[v].ToEntity())
		}
	}

	return products
}
