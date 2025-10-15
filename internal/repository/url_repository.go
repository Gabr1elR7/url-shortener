package repository

import (
	"time"

	"github.com/Gabr1elR7/url-shortener/internal/domain/model"
	"gorm.io/gorm"
)

type URLRepository interface {
	Create(url *model.URL) error
	GetByCode(code string) (*model.URL, error)
	IncrementVisit(code string) error
	GetStats(code string) (*model.URL, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) Create(u *model.URL) error {
	return r.db.Create(u).Error
}

func (r *urlRepository) GetByCode(code string) (*model.URL, error) {
	var url model.URL
	if err := r.db.First(&url, "code = ?", code).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) IncrementVisit(code string) error {
	now := time.Now()
	return r.db.Model(&model.URL{}).
		Where("code = ?", code).
		Updates(map[string]interface{}{
			"visits":      gorm.Expr("visits + ?", 1),
			"last_visit": now,
		}).Error
}

func (r *urlRepository) GetStats(code string) (*model.URL, error) {
	return r.GetByCode(code)
}