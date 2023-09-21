package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// RClient 声明一个全局的rdb变量
var RClient *redis.Client

type Settings struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolSize int    `json:"poolSize"`
	Timeout  int    `json:"timeout"`
}

// InitClient 初始化连接
func InitClient(settings Settings) (err error) {
	RClient = redis.NewClient(&redis.Options{
		Addr:     settings.Addr,
		Password: settings.Password, // no password set
		DB:       settings.Db,       // use default DB
		PoolSize: settings.PoolSize, // 连接池大小
	})

	timeout := settings.Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	_, err = RClient.Ping(ctx).Result()
	return err
}
