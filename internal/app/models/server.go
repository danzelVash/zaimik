package models

import (
	"context"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) setConfig(port string, handler http.Handler) {
	maxHeaderBytes := viper.GetInt("http.max_header_bytes")
	readHeaderTimeout := viper.GetDuration("http.read_header_timeout")
	readTimeout := viper.GetDuration("http.read_timeout")
	writeTimeout := viper.GetDuration("http.write_timeout")

	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    maxHeaderBytes,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
	}

}

func (s *Server) Run(port string, handler http.Handler) error {
	s.setConfig(port, handler)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
