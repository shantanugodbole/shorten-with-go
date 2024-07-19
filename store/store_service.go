package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"encoding/json"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const cacheDuration = 6 * time.Hour

type Config struct{
	redisPassword string `json:"store_pwd"`
}


func loadConfig(file string) (*Config, error){
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
        return nil, err
    }
    defer configFile.Close()

    bytes, err := ioutil.ReadAll(configFile)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(bytes, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}


func InitializeStore() *StorageService{
	config, err := loadConfig("config.json")
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis-16324.c81.us-east-1-2.ec2.redns.redis-cloud.com:16324",
		Password: config.redisPassword,
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