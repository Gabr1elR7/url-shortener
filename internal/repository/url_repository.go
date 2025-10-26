package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/Gabr1elR7/url-shortener/internal/domain/model"
	redisCache "github.com/Gabr1elR7/url-shortener/internal/infrastructure/cache"
)

type URLRepository interface {
	Create(url *model.URL) error
	GetByCode(code string) (*model.URL, error)
	IncrementVisit(code string) error
	GetStats(code string) (*model.URL, error)
}

type urlRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewURLRepository(db *gorm.DB, cache *redis.Client) URLRepository {
	return &urlRepository{db: db, cache: cache}
}

func (r *urlRepository) Create(u *model.URL) error {
	return r.db.Create(u).Error
}

func (r *urlRepository) GetByCode(code string) (*model.URL, error) {
	cacheKey := fmt.Sprintf("urlCode:%s", code)

	// Search in redis
	val, err := r.cache.Get(redisCache.Ctx, cacheKey).Result()
	if err == nil {
		var cached model.URL
		_ = json.Unmarshal([]byte(val), &cached)
		fmt.Println("📦 URL obtenida desde Redis")
		return &cached, nil
	}

	// Search in db
	var url model.URL
	if err := r.db.First(&url, "code = ?", code).Error; err != nil {
		return nil, err
	}

	// Review for 5 minutes
	data, _ := json.Marshal(url)
	r.cache.Set(redisCache.Ctx, cacheKey, data, 5 * time.Minute)
	fmt.Println("💾 URL guardada en Redis")

	return &url, nil
}

func (r *urlRepository) IncrementVisit(code string) error {
	now := time.Now()
	return r.db.Model(&model.URL{}).
		Where("code = ?", code).
		Updates(map[string]interface{}{
			"visits":     gorm.Expr("visits + ?", 1),
			"last_visit": now,
		}).Error
}

func (r *urlRepository) GetStats(code string) (*model.URL, error) {
	return r.GetByCode(code)
}
