package conf

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

type RedisClient struct{}

var (
	rdClient  *redis.Client
	nDuration = 30 * 24 * 60 * 60 * time.Second
	ctx       = context.Background()
)

// InitRedis redis-server.exe 启动
func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{DB: 0, Addr: viper.GetString("redis.url"), Password: ""})
	_, err := rdClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any, reset ...any) error {
	d := nDuration
	if len(reset) > 0 {
		if v, ok := reset[0].(time.Duration); ok {
			d = v
		}
	}
	return rdClient.Set(key, value, d).Err()
}
func (rc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(key).Result()
}
func (rc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(key...).Err()
}

func (rc *RedisClient) GetExpireDuration(key string) (time.Duration, error) {
	return rdClient.TTL(key).Result()
}
