package cache

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func SetInstancesInCache(client *redis.Client, instances interface{}) {
	data, err := json.Marshal(instances)
	if err != nil {
		log.Fatalf("Error marshalling instances: %v", err)
	}
	err = client.Set(ctx, "instances", data, 0).Err()
	if err != nil {
		log.Fatalf("Error setting instances in cache: %v", err)
	}
}

func GetInstancesFromCache(client *redis.Client) []byte {
	val, err := client.Get(ctx, "instances").Result()
	if err != nil {
		log.Fatalf("Error getting instances from cache: %v", err)
	}
	return []byte(val)
}
