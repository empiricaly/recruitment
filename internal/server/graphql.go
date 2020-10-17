package server

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"time"

	stdlog "log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	rice "github.com/GeertJohan/go.rice"
	"github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/graph"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	ginlogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate rice embed-go

// Defining the Graphql handler
func graphqlHandler(s *Server) gin.HandlerFunc {
	r := &graph.Resolver{
		MTurk:       s.mturk,
		MTurkSanbox: s.mturkSandbox,
		Store:       s.storeConn,
		Admins:      s.config.Admins,
		SecretKey:   s.config.SecretKey,
	}

	gconf := generated.Config{Resolvers: r}
	gconf.Directives.HasRole = func(ctx context.Context, _ interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
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

	return func(c *gin.Context) {
		gqlsrv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playgroundServer("Empirica Recruitment GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Web site handler
func webHandler(queryHandler, playHandler, questionsHandler gin.HandlerFunc) gin.HandlerFunc {
	box := rice.MustFindBox("../../web/public")
	httpFS := MakeSPABox(box)

	return func(c *gin.Context) {
		parts := strings.Split(c.Request.URL.Path, "/")

		var group string
		if len(parts) > 1 {
			group = parts[1]
		}

		switch group {
		case "query":
			queryHandler(c)
		case "q":
			questionsHandler(c)
		case "play":
			playHandler(c)
		default:
			c.FileFromFS(c.Request.URL.Path, httpFS)
		}
	}
}

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

func (s *Server) startGraphqlServer() {
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

// SPABox implements http.FileSystem which allows the use of Box with a http.FileServer.
//   e.g.: http.Handle("/", http.FileServer(rice.MustFindBox("http-files").HTTPBox()))
type SPABox struct {
	box *rice.Box
}

// MakeSPABox creates a new SPABox from an existing Box
func MakeSPABox(box *rice.Box) *SPABox {
	return &SPABox{box}
}

// Open returns a File using the http.File interface
func (box *SPABox) Open(name string) (http.File, error) {
	f, err := box.box.Open(name)
	if err != nil {
		return box.Open("index.html")
	}
	return f, nil
}
