package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port            int
	OrderServiceUrl string
}

func NewServer(port int, orderServiceUrl string) *Server {
	return &Server{
		Port:            port,
		OrderServiceUrl: orderServiceUrl,
	}
}

func (s *Server) Start() error {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from order server!")

		res, err := http.Get(s.OrderServiceUrl)
		if err != nil {
			fmt.Fprintf(w, "Error calling order service: %s", err)
		} else {
			fmt.Fprint(w, "Got response!")
			fmt.Fprintf(w, "Status code: %d", res.StatusCode)
		}
	})

	url := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Server is running on http://localhost%s\n", url)
	return http.ListenAndServe(url, mux)
}
