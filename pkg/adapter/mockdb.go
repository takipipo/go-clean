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
	FindByID(ctx context.Context, result interface{}, ID string) error
}

type MockDatabase struct {
	products []entity.Product
}

func NewMockDatabase() IDatabase { return &database }
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
func (a *MockDatabase) FindByID(ctx context.Context, result interface{}, ID string) error {
	for _, p := range a.products {
		if p.ID == ID {
			origJSON, err := json.Marshal(p)
			json.Unmarshal(origJSON, result)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
