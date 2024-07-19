package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"

)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const cacheDuration = 6 * time.Hour


func InitializeStore() *StorageService{
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis-16324.c81.us-east-1-2.ec2.redns.redis-cloud.com:16324",
		Password: "ODaymZn9TezzCDBKYyI8BgKiZ2W4G3Bn",
		DB: 0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil{
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\n Redis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}


func SaveUrlMapping(shortUrl string, originalUrl string, userId string){

	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, cacheDuration).Err()
	if err != nil{
		panic(fmt.Sprintf("Failed saving key URL | Error %v - shorturl %s - orginalurl %s \n", err, shortUrl,originalUrl))
	}
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

func RetrieveInitialUrl(shortUrl string) string{
	result, err := storeService.redisClient.Get(ctx,shortUrl).Result()
	if err != nil{
		panic(fmt.Sprintf("Failed Retrieve Initial URL url | Error %v - shorturl %s \n", err, shortUrl))
	}
	return result
}