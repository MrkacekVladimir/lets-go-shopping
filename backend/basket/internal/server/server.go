package server

import (
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	Port            int
	OrderServiceUrl string
	RedisClient     *redis.Client
}

func NewServer(port int, orderServiceUrl string, redisClient *redis.Client) *Server {
	return &Server{
		Port:            port,
		OrderServiceUrl: orderServiceUrl,
		RedisClient:     redisClient,
	}
}

func (s *Server) Start() error {

	mux := http.NewServeMux()

	mux.HandleFunc("/my-basket", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from order server!")

		/*
			res, err := http.Get(s.OrderServiceUrl)
			if err != nil {
				fmt.Fprintf(w, "Error calling order service: %s", err)
			} else {
				fmt.Fprint(w, "Got response!")
				fmt.Fprintf(w, "Status code: %d", res.StatusCode)
			}
		*/
	})

	url := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Server is running on http://localhost%s\n", url)
	return http.ListenAndServe(url, mux)
}
