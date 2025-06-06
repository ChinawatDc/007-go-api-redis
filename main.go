package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"007-go-api-redis/controllers"
	"007-go-api-redis/routes"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()
var rdb *redis.Client

func initRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis connect error: %v", err)
	}

	fmt.Println("✅ Redis connected")
}

func main() {
	initRedis()
	controllers.SetRedisClient(rdb) // inject redis client

	r := routes.SetupRouter()
	r.Run(":8080")
}
