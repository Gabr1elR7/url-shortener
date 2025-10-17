package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Ctx = context.Background()

func NewRedisClient(addr, pass string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Error conectando a Redis: %v", err)
	}

	log.Println("✅ Conectado a Redis")
	return client
}
