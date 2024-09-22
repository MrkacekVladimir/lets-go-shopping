package main

import (
	"log"

	"letsgoshopping.com/order/internal/server"
)

func main() {
	server := server.NewServer(8090)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
