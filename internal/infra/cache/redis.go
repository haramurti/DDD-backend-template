package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Connect(addr string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}

// TODO: helper cache-aside (GetOrSet) + invalidasi (version key / delete key)
