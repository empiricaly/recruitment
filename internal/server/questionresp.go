package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/empiricaly/recruitment/internal/ent"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	stepRunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	mdExtension "github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

const htmlHead = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Questions</title><style>*,::after,::before{box-sizing:border-box}ol[class],ul[class]{padding:0}blockquote,body,dd,dl,figcaption,figure,h1,h2,h3,h4,li,ol[class],p,ul[class]{margin:0}body{min-height:100vh;scroll-behavior:smooth;text-rendering:optimizeSpeed;line-height:1.5;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif}ol[class],ul[class]{list-style:none}a:not([class]){text-decoration-skip-ink:auto}img{max-width:100%;display:block}article>*+*{margin-top:1em}button,input,select,textarea{font:inherit}@media (prefers-reduced-motion:reduce){*{animation-duration:.01ms!important;animation-iteration-count:1!important;transition-duration:.01ms!important;scroll-behavior:auto!important}}</style></head><body>`
const htmlFoot = `</body></html>`

type questionHandler struct {
	*Server
}

func ginQuestionsHandler(s *Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := strings.TrimPrefix(c.Request.URL.Path, "/q/")

		stepRun, err := s.storeConn.StepRun.
			Query().
			WithStep(func(step *ent.StepQuery) {
				step.WithTemplate()
			}).
			WithRun().
			Where(stepRunModel.UrlTokenEQ(id)).
			First(c.Request.Context())
		if err != nil {
			log.Error().Err(err).Msg("get stepRun")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		step, err := stepRun.Edges.StepOrErr()
		if err != nil {
			log.Error().Err(err).Msg("get step")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if step.Type != stepModel.TypeMTURK_HIT {
			log.Error().Err(err).Msg("is mturk HIT step")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		msgArgs := &model.MessageStepArgs{}
		err = json.Unmarshal(step.MsgArgs, msgArgs)
		if err != nil {
			log.Error().Err(err).Msg("decode mturk HIT step args")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		switch msgArgs.MessageType {
		case model.ContentTypeHTML:
			var out string
			if strings.Contains(msgArgs.Message, "<html>") {
				out = msgArgs.Message
			} else {
				out = htmlHead + msgArgs.Message + htmlFoot
			}

			c.Header("Content-Type", "text/html; charset=utf-8")
			log.Debug().Str("content", out).Msg("html message")
			c.String(200, out)
		case model.ContentTypeMarkdown:
			md := goldmark.New(
				goldmark.WithExtensions(mdExtension.GFM),
				goldmark.WithParserOptions(
					parser.WithAutoHeadingID(),
				),
				goldmark.WithRendererOptions(
					html.WithHardWraps(),
					html.WithXHTML(),
					html.WithUnsafe(),
				),
			)
			log.Debug().Str("content", msgArgs.Message).Msg("markdown message")
			var buf bytes.Buffer
			if err := md.Convert([]byte(msgArgs.Message), &buf); err != nil {
				log.Error().Err(err).Msg("convert markdown")
				c.AbortWithStatus(http.StatusNotFound)
				return
			}

			c.Header("Content-Type", "text/html; charset=utf-8")
			log.Debug().Str("content", buf.String()).Msg("html message")
			c.String(200, htmlHead+buf.String()+htmlFoot)
		case model.ContentTypeSvelte:
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, "<html>react not yet supported</html>")
		case model.ContentTypeReact:
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, "<html>react not yet supported</html>")
		default:
			log.Error().Msgf("unknown step message type: %s", msgArgs.MessageType.String())
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}

func (q *questionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/q/")

	stepRun, err := q.storeConn.StepRun.
		Query().
		WithStep(func(step *ent.StepQuery) {
			step.WithTemplate()
		}).
		WithRun().
		Where(stepRunModel.UrlTokenEQ(id)).
		First(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("get stepRun")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	step, err := stepRun.Edges.StepOrErr()
	if err != nil {
		log.Error().Err(err).Msg("get step")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if step.Type != stepModel.TypeMTURK_HIT {
		log.Error().Err(err).Msg("is mturk HIT step")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	msgArgs := &model.MessageStepArgs{}
	err = json.Unmarshal(step.MsgArgs, msgArgs)
	if err != nil {
		log.Error().Err(err).Msg("decode mturk HIT step args")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch msgArgs.MessageType {
	case model.ContentTypeHTML:
		var out string
		if strings.Contains(msgArgs.Message, "<html>") {
			out = msgArgs.Message
		} else {
			out = htmlHead + msgArgs.Message + htmlFoot
		}

		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		log.Debug().Str("content", out).Msg("html message")
		w.Write([]byte(out))
	case model.ContentTypeMarkdown:
		md := goldmark.New(
			goldmark.WithExtensions(mdExtension.GFM),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
				html.WithUnsafe(),
			),
		)
		log.Debug().Str("content", msgArgs.Message).Msg("markdown message")
		var buf bytes.Buffer
		if err := md.Convert([]byte(msgArgs.Message), &buf); err != nil {
			log.Error().Err(err).Msg("convert markdown")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		log.Debug().Str("content", buf.String()).Msg("html message")
		w.Write([]byte(htmlHead + buf.String() + htmlFoot))
	case model.ContentTypeSvelte:
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte("<html>react not yet supported</html>"))
	case model.ContentTypeReact:
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte("<html>react not yet supported</html>"))
	default:
		log.Error().Msgf("unknown step message type: %s", msgArgs.MessageType.String())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
