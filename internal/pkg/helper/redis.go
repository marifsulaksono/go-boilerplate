package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Helper functions for setting data to Redis
func SetRedisJSONCache(ctx context.Context, rds *redis.Client, key string, data interface{}, ttl time.Duration) error {
	// Marshal the value into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data for Redis: %w", err)
	}

	// Set the JSON data in Redis with the specified TTL
	if err := rds.Set(ctx, key, jsonData, ttl).Err(); err != nil {
		return fmt.Errorf("error setting data in Redis: %w", err)
	}
	return nil
}
