package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var Ctx = context.Background()
var RDB *redis.Client

func InitRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err = RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis connect error: %v", err)
	}

	fmt.Println("✅ Redis connected")
}
