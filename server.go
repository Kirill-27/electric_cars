package electric_cars

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

//handler http.Handler
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		//Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}