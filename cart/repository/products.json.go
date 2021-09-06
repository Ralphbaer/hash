package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	e "github.com/Ralphbaer/hash/cart/entity"
)

// ProductJSONRepository represents a local JSON implementation of ProductRepository interface
type ProductJSONRepository struct {
	FilePath string
}

// NewProductJSONRepository creates an instance of repository.ProductJSONRepository
func NewProductJSONRepository() *ProductJSONRepository {
	return &ProductJSONRepository{
		FilePath: "./products.json",
	}
}

// List returns all products
func (c *ProductJSONRepository) List(ctx context.Context, f *ProductFilter) ([]*e.Product, error) {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/repository/%s", pwd, c.FilePath))
	if err != nil {
		return nil, err
	}

	var obj []*ProductJSONModel
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	return withIDs(obj, f.IDs), nil
}
