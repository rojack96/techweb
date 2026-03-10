package server

import (
	// go imports
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// external imports
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	// project imports
)

type Server struct {
	host   string
	port   uint16
	log    *zap.Logger
	router *gin.Engine
}

func NewServer(opts ...Options) *Server {
	s := &Server{
		host: "0.0.0.0",
		port: 8201,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Options - designed follow the idiomatic Go Functional options pattern
type Options func(s *Server)

func WithHost(host string, port uint16) Options {
	return func(s *Server) {
		s.host = host
		s.port = port
	}
}

func WithRouter(router *gin.Engine) Options {
	return func(s *Server) {
		s.router = router
	}
}

func WithSwagger(enabled bool, service string, auth bool, user, passwd string) Options {
	return func(s *Server) {
		if !enabled {
			return
		}

		workdir, _ := os.Getwd()
		workdir = strings.Replace(workdir, "/cmd/rest_api", "", 1)
		swaggerFile := filepath.Join(workdir, "api-docs", "swagger.json")

		s.router.StaticFile("/notification-service/api-docs", swaggerFile)

		// path base per swagger
		swaggerBase := "/" + service + "/swagger"
		url := ginSwagger.URL("/notification-service/api-docs")
		if !auth {
			s.router.GET(swaggerBase+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
			return
		}

		// route protetta con BasicAuth
		swaggerGroup := s.router.Group(
			swaggerBase,
			gin.BasicAuth(gin.Accounts{user: passwd}),
		)
		swaggerGroup.GET("/*any",
			ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}

func WithZapLogger(log *zap.Logger) Options {
	return func(s *Server) {
		s.log = log
	}
}

func (s *Server) Serve() {
	if err := s.router.Run(fmt.Sprintf("%s:%d", s.host, s.port)); err != nil {
		return
	}
}
