package server

import (
	"bytes"
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

		switch step.MsgArgs.MessageType {
		case model.ContentTypeHTML:
			var out string
			if strings.Contains(step.MsgArgs.Message, "<html>") {
				out = step.MsgArgs.Message
			} else {
				out = htmlHead + step.MsgArgs.Message + htmlFoot
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
			log.Debug().Str("content", step.MsgArgs.Message).Msg("markdown message")
			var buf bytes.Buffer
			if err := md.Convert([]byte(step.MsgArgs.Message), &buf); err != nil {
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
			log.Error().Msgf("unknown step message type: %s", step.MsgArgs.MessageType.String())
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
