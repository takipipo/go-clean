package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/takipipo/go-clean/cmd/api/handler"
	"github.com/takipipo/go-clean/cmd/api/router"
	"github.com/takipipo/go-clean/pkg/adapter"
	"github.com/takipipo/go-clean/pkg/repository"
	"github.com/takipipo/go-clean/pkg/usecase"
)

func main() {
	var (
		mockDatabase      = adapter.NewMockDatabase()
		productRepository = repository.NewProduct(mockDatabase)
		productUseCase    = usecase.NewProduct(productRepository)

		productHandler = handler.NewProduct(productUseCase)

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	)

	productApp := echo.New()
	router.NewProductRouter(productApp, productHandler)

	go func() {
		if err := productApp.Start(":8000"); err != nil && err != http.ErrServerClosed {
			productApp.Logger.Fatal("shutting down the server")
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	defer cancel()
	if err := productApp.Shutdown(ctx); err != nil {
		productApp.Logger.Fatal(err)
		panic(err)
	}
}
