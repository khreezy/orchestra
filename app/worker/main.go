package main

import (
	"github.com/adjust/rmq/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func main() {
	errChan := make(chan error)

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		log.Fatal(err.Error())
	}

	redisClient := redis.NewClient(opt)

	rc, err := rmq.OpenConnectionWithRedisClient(uuid.NewString(), redisClient, errChan)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = rc.OpenQueue("spotify")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully opened connection")

	<-errChan
}
