package main

import (
	"log"
	"net/http"
	"time"

	"github.com/scotty-c/novasearch/awsclient"
	"github.com/scotty-c/novasearch/cache"
	"github.com/scotty-c/novasearch/config"

	"github.com/go-redis/redis/v8"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	// Load the configuration using Viper
	config.LoadConfig()

	// Start the background job to update instances every 5 minutes
	go updateInstances()

	// Set up HTTP route
	http.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		instances := cache.GetInstancesFromCache(redisClient)
		w.Header().Set("Content-Type", "application/json")
		w.Write(instances)
	})

	// Start the HTTP server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

// Background job to query AWS every 5 minutes
func updateInstances() {
	for {
		time.Sleep(5 * time.Minute)
		instances := awsclient.FindInstancesByTags(config.AppConfig.Tags, config.AppConfig.Region)
		cache.SetInstancesInCache(redisClient, instances)
	}
}
