package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

// rClient 声明一个全局的redis.Client量
var rClient *redis.Client

var once sync.Once

type Settings struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolSize int    `json:"poolSize"`
	Timeout  int    `json:"timeout"`
}

// InitClient 初始化连接
func InitClient(settings Settings) error {
	var err error
	once.Do(func() {
		rClient = redis.NewClient(&redis.Options{
			Addr:     settings.Addr,
			Password: settings.Password, // no password set
			DB:       settings.Db,       // use default DB
			PoolSize: settings.PoolSize, // 连接池大小
		})

		timeout := settings.Timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()

		_, err = rClient.Ping(ctx).Result()
	})
	return err
}

func Set(ctx context.Context, key string, object interface{}, expiration time.Duration) error {
	jsonStr, err := json.Marshal(object)
	if err != nil {
		return err
	}
	if err = rClient.Set(ctx, key, jsonStr, expiration).Err(); err != nil {
		return err
	}
	return err
}

func Get(ctx context.Context, key string) (string, error) {
	return rClient.Get(ctx, key).Result()
}

func Del(ctx context.Context, key string) error {
	_, err := rClient.Get(ctx, key).Result()
	if err == redis.Nil { //key不存在
		return nil
	}
	if err != nil {
		return err
	}
	_, err = rClient.Del(ctx, key).Result()
	return err
}
