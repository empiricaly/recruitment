package server

import (
	"context"
	"net/http"
	"time"

	"github.com/empiricaly/recruitment/internal/admin"
	ginlogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate sh ./buildweb.sh
//go:generate rice embed-go

func (s *Server) startHTTPServer() {
	// Enable GIN debug mode or not
	if s.config.HTTP.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	rgin := gin.New()

	// CORS
	configureCORS(s, rgin)

	// Admin authentication
	rgin.Use(admin.GinMiddleware(s.storeConn, []byte(s.config.SecretKey)))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	rgin.Use(gin.Recovery())

	srvlogger := log.With().Str("pkg", "http").Logger()
	if s.config.HTTP.Debug {
		srvlogger = srvlogger.Level(zerolog.DebugLevel)
		rgin.Use(ginlogger.SetLogger(ginlogger.Config{
			Logger: &srvlogger,
			// UTC:            true,
			// SkipPath:       []string{"/skip"},
			// SkipPathRegexp: rxURL,
		}))
	} else {
		srvlogger = srvlogger.Level(zerolog.InfoLevel)
	}

	queryHandler := graphqlHandler(s)
	rgin.POST("/query", queryHandler)
	rgin.POST("/a/:id", ginAnswersHandler(s))
	rgin.GET("/*rest", webHandler(queryHandler, playgroundHandler(), ginQuestionsHandler(s)))

	srv := &http.Server{
		Addr:    s.config.HTTP.Addr,
		Handler: rgin,
	}

	s.wg.Add(1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srvlogger.Error().Err(err).Msg("Failed start GraphQL server")
			s.wg.Done()
		}
	}()

	go func() {
		srvlogger.Debug().Msgf("Started GraphQL server at %s", s.config.HTTP.Addr)

		<-s.done
		srvlogger.Debug().Msg("Stopping GraphQL server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srvlogger.Error().Err(err).Msg("GraphQL server shutdown failed")
		} else {
			srvlogger.Debug().Msg("GraphQL server gracefully shutdown")
		}
		s.wg.Done()
	}()

}
