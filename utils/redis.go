package utils

import (
    "os"
    "github.com/redis/go-redis/v9"
    "github.com/joho/godotenv"
)

func init() {
    godotenv.Load()
}

var RedisClient = redis.NewClient(&redis.Options{
    Addr:     os.Getenv("REDIS_ADDR"),     
    Password: os.Getenv("REDIS_PASSWORD"), 
    DB:       0,
})