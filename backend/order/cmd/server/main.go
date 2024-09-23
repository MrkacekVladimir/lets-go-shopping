package main

import (
	"log"
	"os"
	"strconv"

	"letsgoshopping.com/order/internal/server"
)

func main() {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("Failed to parse application port.", err)
	}

	server := server.NewServer(port)
	if err := server.Start(); err != nil {
		log.Fatal("Server failed.", err)
	}
}
