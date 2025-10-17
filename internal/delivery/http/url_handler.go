package http

import (
	"fmt"
	"net/http"

	"github.com/Gabr1elR7/url-shortener/internal/usecase"
	"github.com/labstack/echo/v4"
)

type URLHandler struct {
	usecase usecase.URLUsecase
}

func NewURLHandler(u usecase.URLUsecase) *URLHandler {
	return &URLHandler{usecase: u}
}

type shortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func (h *URLHandler) Shorten(c echo.Context) error {
	var req shortenRequest
	if err := c.Bind(&req); err != nil || req.URL == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "URL inválida"})
	}

	url, err := h.usecase.Shorten(req.URL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"shortUrl": fmt.Sprintf("http://localhost:8080/%s", url.Code),
	})
}

func (h *URLHandler) Redirect(c echo.Context) error {
	code := c.Param("code")
	url, err := h.usecase.GetByCode(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "URL no encontrada"})
	}
	return c.Redirect(http.StatusFound, url.OriginalURL)
}

func (h *URLHandler) Stats(c echo.Context) error {
	code := c.Param("code")
	stats, err := h.usecase.GetStats(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "URL no encontrada"})
	}
	return c.JSON(http.StatusOK, stats)
}