package server

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/empiricaly/recruitment/internal/graph"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

func (s *Server) startGraphqlServer() {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	r := &graph.Resolver{}

	// router.Use(MachinesLockMiddleware(r))

	gqlsrv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	gqlsrv.AddTransport(transport.Options{})
	gqlsrv.AddTransport(transport.GET{})
	gqlsrv.AddTransport(transport.POST{})
	gqlsrv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	gqlsrv.Use(extension.Introspection{})

	router.Handle("/play", playground.Handler("Empirica Recruitment GraphQL", "/query"))
	router.Handle("/query", c.Handler(gqlsrv))

	srv := &http.Server{
		Addr:    s.config.GQLAddr,
		Handler: router,
	}

	s.wg.Add(1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("failed start graphql server")
			s.wg.Done()
		}
	}()

	go func() {
		log.Debug().Msgf("Started GraphQL server at %s", s.config.GQLAddr)

		<-s.done
		log.Debug().Msg("stopping graphql server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("graphql server shutdown failed")
		} else {
			log.Debug().Msg("graphql server stopped")
		}
		s.wg.Done()
	}()
}
