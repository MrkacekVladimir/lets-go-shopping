package main

import (
	"log"
	"os"
	"strconv"

	"letsgoshopping.com/basket/internal/server"
)

func main() {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("Failed to parse application port.", err)
	}

	orderServiceUrl := os.Getenv("ORDER_SERVICE_URL")
	server := server.NewServer(port, orderServiceUrl)
	if err := server.Start(); err != nil {
		log.Fatal("Server failed.", err)
	}
}
