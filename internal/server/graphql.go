package server

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	rice "github.com/GeertJohan/go.rice"
	"github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/graph"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	logger "github.com/empiricaly/recruitment/internal/log"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

//go:generate rice embed-go

func (s *Server) startGraphqlServer() {
	router := chi.NewRouter()

	copts := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}
	if s.config.HTTP.Debug {
		copts.Debug = true
	}
	c := cors.New(copts)

	// m, _ := storage.NewMapping(s.storeConn)
	r := &graph.Resolver{
		MTurk:       s.mturk,
		MTurkSanbox: s.mturkSandbox,
		Store:       s.storeConn,
		Admins:      s.config.Admins,
		SecretKey:   s.config.SecretKey,
	}

	// router.Use(MachinesLockMiddleware(r))
	if s.config.HTTP.Debug {
		router.Use(logger.HTTPLogger())
	}
	router.Use(admin.Middleware(s.storeConn, []byte(s.config.SecretKey)))

	gconf := generated.Config{Resolvers: r}
	gconf.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		user := admin.ForContext(ctx)
		if role == model.RoleAdmin && user == nil {
			return nil, errors.New("Access Denied")
		}

		return next(ctx)
	}

	gqlsrv := handler.New(generated.NewExecutableSchema(gconf))

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

	box := rice.MustFindBox("../../web/public")
	router.Handle("/q/*", c.Handler(&questionHandler{s}))
	router.Handle("/play", c.Handler(playground.Handler("Empirica Recruitment GraphQL", "/query")))
	router.Handle("/query", c.Handler(gqlsrv))
	router.Handle("/*", c.Handler(http.FileServer(MakeSPABox(box))))

	srv := &http.Server{
		Addr:    s.config.HTTP.Addr,
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
		log.Debug().Msgf("Started GraphQL server at %s", s.config.HTTP.Addr)

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

// HTTPBox implements http.FileSystem which allows the use of Box with a http.FileServer.
//   e.g.: http.Handle("/", http.FileServer(rice.MustFindBox("http-files").HTTPBox()))
type SPABox struct {
	box   *rice.Box
	index http.File
}

// MakeSPABox creates a new SPABox from an existing Box
func MakeSPABox(box *rice.Box) *SPABox {
	index, err := box.Open("index.html")
	if err != nil {
		panic("public/ must contain an index.html file")
	}
	return &SPABox{box, index}
}

// Open returns a File using the http.File interface
func (box *SPABox) Open(name string) (http.File, error) {
	f, err := box.box.Open(name)
	if err != nil {
		return box.index, nil
	}
	return f, nil
}
