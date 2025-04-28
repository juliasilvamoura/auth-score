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
		Addr:     "localhost:6379", 
		Password: "",              
		DB:       0,                
	})

	// Ping Redis to check connection
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
	return err == nil // se não houver erro, o token está na blacklist
}
