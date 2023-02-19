package usecase

import (
	"context"

	"github.com/takipipo/go-clean/pkg/entity"
	"github.com/takipipo/go-clean/pkg/repository"
)

type IProduct interface {
	GetAll(ctx context.Context) ([]entity.Product, error)
}

type Product struct{ productRepository repository.IProduct }

func NewProduct(productRepository repository.IProduct) IProduct { return &Product{productRepository} }

func (u *Product) GetAll(ctx context.Context) ([]entity.Product, error) {
	result := []entity.Product{}
	err := u.productRepository.GetAll(ctx, &result)
	if err != nil {
		return result, err
	}
	return result, nil

}
