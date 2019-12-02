package stats

import (
	"fmt"

	"github.com/go-redis/redis"
)

const redisHitsCountKey = "fizzbuzz-api:hits"

// RedisStore is redis-backed stats store
type RedisStore struct {
	client *redis.Client
}

// NewRedisStatsStore returns a new *RedisStatsStore
func NewRedisStatsStore(addr string) *RedisStore {
	clt := redis.NewClient(&redis.Options{Addr: addr})
	return &RedisStore{clt}
}

// Increment increments an endpoint's hits
func (r *RedisStore) Increment(endpoint string) error {
	return r.client.ZIncrBy(redisHitsCountKey, 1, endpoint).Err()
}

// GetMax returns endpoint with most hits
func (r *RedisStore) GetMax() (*Stats, error) {
	cmd := r.client.ZRevRangeByScoreWithScores(redisHitsCountKey, redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: 0, Count: 1})
	if err := cmd.Err(); err != nil {
		return nil, fmt.Errorf("could not fetch max hits from redis: %v", err)
	}
	result := cmd.Val()
	if len(result) == 0 {
		return &Stats{}, nil
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("unexpected result from redis: %v", result)
	}
	parameters := result[0].Member.(string)
	count := int(result[0].Score)
	return &Stats{parameters, count}, nil
}

// Flush resets all stats
func (r *RedisStore) Flush() error {
	return r.client.Del(redisHitsCountKey).Err()
}

// Close closes the redis store
func (r *RedisStore) Close() error {
	return r.client.Close()
}
