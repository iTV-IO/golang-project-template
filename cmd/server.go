package server

import (
	"net/http"
	"time"

	"github.com/quvonchbe05/golang-clean-architecture-template/app/http/middleware"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	// Cors
	handler = middleware.CORSMiddleware(handler)
	
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}