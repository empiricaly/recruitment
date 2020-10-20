package server

import (
	"net/http"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

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
