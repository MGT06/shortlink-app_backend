package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	User string
	Pass string
	Host string
	Port string
}

func NewRedisClient(user, pass, host, port string) *RedisClient {
	return &RedisClient{
		User: user,
		Pass: pass,
		Host: host,
		Port: port,
	}
}

func (r *RedisClient) ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Username: r.User,
		Password: r.Pass,
	})
}
