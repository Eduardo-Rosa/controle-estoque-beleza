package usecase

import (
	"testing"
	"controle-estoque-beleza/domain"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	uc := ProductUsecase{}
	product := domain.Product{
		Name:     "Shampoo",
		Brand:    "Marca X",
		Category: "Cuidados com o Cabelo",
		Price:    19.99,
		Quantity: 50,
		SKU:      "12345",
	}
	err := uc.CreateProduct(product)
	assert.Nil(t, err)
}

func TestGetProducts(t *testing.T) {
	uc := ProductUsecase{}
	products, err := uc.GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
}