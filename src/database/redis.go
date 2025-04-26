package database

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis default address
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// Ping Redis to check connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Erro ao conectar ao Redis:", err)
	}

	log.Println("Conectado ao Redis com sucesso!")
}

// AddToBlacklist adiciona um token à blacklist
func AddToBlacklist(token string, expiration time.Duration) error {
	return RedisClient.Set(Ctx, "blacklist:"+token, true, expiration).Err()
}

// IsTokenBlacklisted verifica se um token está na blacklist
func IsTokenBlacklisted(token string) bool {
	_, err := RedisClient.Get(Ctx, "blacklist:"+token).Result()
	return err == nil // se não houver erro, o token está na blacklist
}
