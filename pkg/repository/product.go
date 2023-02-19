package repository

import (
	"context"

	"github.com/takipipo/go-clean/pkg/adapter"
	"github.com/takipipo/go-clean/pkg/entity"
)

type IProduct interface {
	GetAll(ctx context.Context, result *[]entity.Product) (error)
}
type Product struct{ database adapter.IDatabase }

func NewProduct(database adapter.IDatabase) IProduct { return &Product{database} }

func (r *Product) GetAll(ctx context.Context, result *[]entity.Product) (error) {
	err := r.database.FindAll(ctx, result)
	if err != nil {
		return err
	}
	return nil
}
