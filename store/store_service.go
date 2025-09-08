package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw redis client
type StorageService struct {
	redisClient *redis.Client
}

// top level delarations for storeservice and redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// initialize the store service (redis client), and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService

}

// save the mapping between the originalURL and shortURL
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

// retrieve the initial long url once the short is provided and users will call it
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).err()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
