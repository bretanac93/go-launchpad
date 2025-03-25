package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

type Server struct {
    addr string
    r *gin.Engine
}

func New(cfg Config) *Server {
    r := gin.New()

	r.Use(sloggin.New(slog.Default()))
    r.Use(gin.Recovery())

    return &Server{
        addr: cfg.Addr(),
        r: r,
    }
} 

func (s *Server) AddMiddlewares(middleware ...gin.HandlerFunc) {
    s.r.Use(middleware...)
}

func (s *Server) AddHandlers(handlers ...Handler) {
    for _, h := range handlers {
        h.Register(s.r)
    }
}

func (s *Server) Run() {
    if err := s.r.Run(s.addr); err != nil {
        slog.Error(err.Error())
    }
}

