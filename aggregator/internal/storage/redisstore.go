package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func SaveToRedis(rdb *redis.Client, date string, articles []AggregatedArticle) error {
	key := fmt.Sprintf("feed:%s", date)
	data, err := json.Marshal(articles)
	if err != nil {
		return err
	}
	// Set TTL to 7 days
	expiration := 2 * 24 * time.Hour
	return rdb.Set(ctx, key, data, expiration).Err()
}