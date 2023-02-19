package adapter

import (
	"context"
	"encoding/json"

	"github.com/takipipo/go-clean/pkg/entity"
)

var database = MockDatabase{
	products: []entity.Product{
		{ID: "0", Category: "Electronics", Name: "Computer"},
		{ID: "1", Category: "Electronics", Name: "IPhone"},
	},
}

type IDatabase interface {
	FindAll(ctx context.Context, result interface{}) error
	// FindOne(ctx context.Context, result interface{}, filter interface{}) error
}

type MockDatabase struct {
	products []entity.Product
}
func NewMockDatabase() IDatabase {return &database}
func (a *MockDatabase) FindAll(ctx context.Context, result interface{}) error {
	origJSON, err := json.Marshal(a.products)
	if err != nil {
		return err
	}
	err = json.Unmarshal(origJSON, result)
	if err != nil {
		return err
	}
	return nil
}
