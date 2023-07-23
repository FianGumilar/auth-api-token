package component

import (
	"context"
	"log"
	"time"

	"github.com/FianGumilar/auth-api-token/domain"
	"github.com/allegro/bigcache/v3"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("Failed connect to cache %s", err)
	}
	return cache
}
