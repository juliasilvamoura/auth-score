package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConnectRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Erro ao conectar ao Redis:", err)
	}

	log.Println("Conectado ao Redis com sucesso!")
}

func AddToBlacklist(token string, expiration time.Duration) error {
	return RedisClient.Set(Ctx, "blacklist:"+token, true, expiration).Err()
}

func IsTokenBlacklisted(token string) bool {
	_, err := RedisClient.Get(Ctx, "blacklist:"+token).Result()
	return err == nil
}
