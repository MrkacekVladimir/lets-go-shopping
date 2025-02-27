package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start() error {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from order server!")
	})

	url := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Server is running on http://localhost%s\n", url)
	return http.ListenAndServe(url, mux)
}
