package main

import (
	"github.com/Gabr1elR7/url-shortener/internal/config"
	urlHttp "github.com/Gabr1elR7/url-shortener/internal/delivery/http"
	"github.com/Gabr1elR7/url-shortener/internal/domain/model"
	redisCache "github.com/Gabr1elR7/url-shortener/internal/infrastructure/cache"
	"github.com/Gabr1elR7/url-shortener/internal/infrastructure/database"
	"github.com/Gabr1elR7/url-shortener/internal/repository"
	"github.com/Gabr1elR7/url-shortener/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Configuration
	cfg := config.Load()

	// DB
	db := database.NewPostgres(cfg.DatabaseURL, &model.URL{})

	// Redis
	cache := redisCache.NewRedisClient(cfg.RedisADDR, cfg.RedisPass)

	// DI
	repo := repository.NewURLRepository(db, cache)
	uc := usecase.NewURLUsecase(repo)
	handler := urlHttp.NewURLHandler(uc)

	e := echo.New()

	// Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Routing
	e.POST("/api/shorten", handler.Shorten)
	e.GET("/:code", handler.Redirect)
	e.GET("/api/stats/:code", handler.Stats)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
