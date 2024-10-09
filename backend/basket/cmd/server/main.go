package main

import (
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"letsgoshopping.com/basket/internal/server"
)

func main() {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("Failed to parse application port.", err)
	}

	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	rdb := redis.NewClient(options)

	orderServiceUrl := os.Getenv("ORDER_SERVICE_URL")
	server := server.NewServer(port, orderServiceUrl, rdb)
	if err := server.Start(); err != nil {
		log.Fatal("Server failed.", err)
	}
}
