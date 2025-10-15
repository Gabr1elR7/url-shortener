package usecase

import (
	"github.com/Gabr1elR7/url-shortener/internal/domain/model"
	"github.com/Gabr1elR7/url-shortener/internal/repository"
	"github.com/google/uuid"
)

type URLUsecase interface {
	Shorten(original string) (*model.URL, error)
	GetByCode(code string) (*model.URL, error)
	GetStats(code string) (*model.URL, error)
}

type urlUsecase struct {
	repo repository.URLRepository
}

func NewURLUsecase(repo repository.URLRepository) URLUsecase {
	return &urlUsecase{repo: repo}
}

func (u *urlUsecase) Shorten(original string) (*model.URL, error) {
	code := uuid.New().String()[:8]
	url := &model.URL{
		Code:        code,
		OriginalURL: original,
	}
	if err := u.repo.Create(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (u *urlUsecase) GetByCode(code string) (*model.URL, error) {
	if err := u.repo.IncrementVisit(code); err != nil {
		return nil, err
	}
	return u.repo.GetByCode(code)
}

func (u *urlUsecase) GetStats(code string) (*model.URL, error) {
	return u.repo.GetStats(code)
}