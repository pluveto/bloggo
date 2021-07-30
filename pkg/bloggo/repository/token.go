package repository

import (
	"context"

	"github.com/go-redis/redis"
)

func (r *Repository) TokenIsDestroyed(token string) bool {
	var ctx = context.Background()
	_, err := r.rdb.Get(ctx, "token_blacklist_"+token).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		panic(err)
	} else {
		return true
	}
}

