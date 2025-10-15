package main

import (
	"github.com/Gabr1elR7/url-shortener/internal/config"
	urlHttp "github.com/Gabr1elR7/url-shortener/internal/delivery/http"
	"github.com/Gabr1elR7/url-shortener/internal/domain/model"
	"github.com/Gabr1elR7/url-shortener/internal/infrastructure/database"
	"github.com/Gabr1elR7/url-shortener/internal/repository"
	"github.com/Gabr1elR7/url-shortener/internal/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg.DatabaseURL, &model.URL{})

	repo := repository.NewURLRepository(db)
	uc := usecase.NewURLUsecase(repo)
	handler := urlHttp.NewURLHandler(uc)

	e := echo.New()
	e.POST("/api/shorten", handler.Shorten)
	e.GET("/api/:code", handler.Redirect)
	e.GET("/api/stats/:code", handler.Stats)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}