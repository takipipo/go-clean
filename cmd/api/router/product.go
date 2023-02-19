package router

import (
	"github.com/labstack/echo/v4"
	"github.com/takipipo/go-clean/cmd/api/handler"
)

func NewProductRouter(productApp *echo.Echo, productHandler handler.IProduct) {
	productApp.GET("/", productHandler.HealthCheck)
	productApp.GET("/products", productHandler.GetAll)
}
