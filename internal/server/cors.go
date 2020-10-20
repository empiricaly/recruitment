package server

import (
	"bytes"
	stdlog "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

type corsLogWriter struct{}

func (c *corsLogWriter) Write(b []byte) (int, error) {
	log.Debug().Str("pkg", "cors").Msg(string(bytes.TrimSpace(b)))
	return len(b), nil
}

func configureCORS(s *Server, rgin *gin.Engine) {
	opts := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}
	if s.config.HTTP.Debug {
		opts.Debug = true
	}

	c := cors.New(opts)
	if s.config.HTTP.Debug {
		c.Log = stdlog.New(&corsLogWriter{}, "", 0)
	}

	rgin.Use(func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		if ctx.Request.Method == http.MethodOptions &&
			ctx.GetHeader("Access-Control-Request-Method") != "" {
			// Abort processing next Gin middlewares.
			ctx.AbortWithStatus(http.StatusOK)
		}
	})
}
