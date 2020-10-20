package server

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/graph"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

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
