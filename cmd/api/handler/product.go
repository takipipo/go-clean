package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takipipo/go-clean/pkg/usecase"
)

type IProduct interface {
	HealthCheck(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetByID(ctx echo.Context) error
}

type product struct{ productUseCase usecase.IProduct }

func NewProduct(usecase usecase.IProduct) IProduct { return &product{usecase} }

func (h product) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello from Product API")
}
func (h product) GetAll(ctx echo.Context) error {
	prods, err := h.productUseCase.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &prods)

}
func (h product) GetByID(ctx echo.Context) error {
	ID := ctx.Param("id")
	prod, err := h.productUseCase.GetByID(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &prod)
}
