package usecase

import (
	"controle-estoque-beleza/domain"
)

type ProductUsecase struct {
	// Adicione dependências como repositórios aqui
}

func (uc *ProductUsecase) CreateProduct(product domain.Product) error {
	// Lógica para criar um produto
	return nil
}

func (uc *ProductUsecase) GetProducts() ([]domain.Product, error) {
	// Lógica para listar produtos
	return nil, nil
}