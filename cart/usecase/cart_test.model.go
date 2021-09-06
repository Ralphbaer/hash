package usecase

import (
	e "github.com/Ralphbaer/hash/cart/entity"
	r "github.com/Ralphbaer/hash/cart/repository"
)

var happyScenario = []*e.Product{
	products[0],
	products[1],
	products[2],
}

var giftScenario = []*e.Product{
	products[3],
}

var products = []*e.Product{
	{
		ID:          1,
		Title:       "XPTO Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      60000,
		IsGift:      false,
	},
	{
		ID:          2,
		Title:       "ABC Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      52000,
		IsGift:      false,
	},
	{
		ID:          3,
		Title:       "DEF Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      32000,
		IsGift:      false,
	},
	{
		ID:          4,
		Title:       "GHI Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      43210,
		IsGift:      true,
	},
	{
		ID:          5,
		Title:       "JKL Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      17555,
		IsGift:      false,
	},
	{
		ID:          6,
		Title:       "MNO Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      90111,
		IsGift:      true,
	},
	{
		ID:          7,
		Title:       "PQR Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      87312,
		IsGift:      true,
	},
	{
		ID:          8,
		Title:       "STU Product",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Amount:      56781,
		IsGift:      true,
	},
}

var createCartInput = &CreateCartInput{
	Products: []*ProductRequest{
		{
			ID:       1,
			Quantity: 2,
		},
		{
			ID:       2,
			Quantity: 3,
		},
		{
			ID:       3,
			Quantity: 1,
		},
	},
}

var createCartWithGiftInput = &CreateCartInput{
	Products: []*ProductRequest{
		{
			ID:       4,
			Quantity: 2,
		},
	},
}

var giftProductFilter = &r.ProductFilter{
	IDs: []int32{4},
}
