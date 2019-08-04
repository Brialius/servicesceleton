package server

import (
	"fmt"
	"net/http"
)

type HTTPServer struct {
	httpPort int
	handler  http.HandlerFunc
}

func (s *HTTPServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.httpPort), s.handler)
}

func NewServer(p int, h http.HandlerFunc) *HTTPServer {
	return &HTTPServer{p, h}
}
