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

const htmlHead = `<!DOCTYPE html><html lang="en"><head>
<style>
	.loader {
		margin: auto;
		border: 16px solid #f3f3f3; 
		border-top: 16px solid #3498db; 
		border-radius: 50%;
		width: 80px;
		height: 80px;
		animation: spin 2s linear infinite;
	}

	@keyframes spin {
		0% { transform: rotate(0deg); }
		100% { transform: rotate(360deg); }
	}
</style>

<script>
  window.addEventListener("load", () => {
    const params = new URL(document.location).searchParams;
    if (!params.get("turkSubmitTo")) {
      console.log("can't find redirect param", params);
      return;
    }

    const forms = document.querySelectorAll("form");
    forms.forEach((f) => {
      f.addEventListener("submit", (e) => {
        e.preventDefault();
        f.style.display = "none";
        const loading = document.createElement("div");
        loading.className = "loader";
        document.body.appendChild(loading);

        const data = {};
        e.currentTarget.querySelectorAll("input").forEach((el) => {
          data[el.name] = el.value;
        });
        try {
          (async () => {
            const url =
              document.location.origin +
              "/a" +
              document.location.pathname.slice(2) +
              "?" +
              params.toString();
            const response = await fetch(url, {
              method: "POST", // *GET, POST, PUT, DELETE, etc.
              headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
              },
              body: JSON.stringify(data),
            });

            loading.remove();

            if (response.ok) {
							const assignmentId = document.createElement("input");
              assignmentId.name = "assignmentId";
              assignmentId.type = "hidden";
              assignmentId.value = params.get("assignmentId");
							
              f.action = params.get("turkSubmitTo") + "/mturk/externalSubmit";
              f.method = "POST";
							f.appendChild(assignmentId);
							f.submit();
							
              return;
            }

            console.error("error ", response.status, response.statusText);

            const div = document.createElement("div");
            div.style.backgroundColor = "#e53e3e";
            div.style.padding = "20px";
            div.style.color = "#f7fafc";

            const text = document.createElement("p");
            text.style.fontSize = "24px";
            text.innerText =
              "Something went wrong. Please try again or report back to us.";

            const button = document.createElement("button");
            button.style.backgroundColor = "#f7fafc";
            button.style.color = "#e53e3e";
            button.style.borderRadius = "5px";
            button.style.borderColor = "transparent";
            button.style.padding = "3px 20px";
            button.style.marginTop = "10px";
            button.style.fontSize = "16px";
            button.style.cursor = "pointer";
            button.innerText = "Try Again";
            button.addEventListener("click", (e) => {
              e.preventDefault();
              f.style.display = "block";
              div.remove();
            });

            div.appendChild(text);
            div.appendChild(button);
            document.body.appendChild(div);
          })();
        } catch (error) {
          console.error("Submitting Form: ", error);
        }
      });
    });
  });
</script>



<meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Questions</title><style>*,::after,::before{box-sizing:border-box}ol[class],ul[class]{padding:0}blockquote,body,dd,dl,figcaption,figure,h1,h2,h3,h4,li,ol[class],p,ul[class]{margin:0}body{min-height:100vh;scroll-behavior:smooth;text-rendering:optimizeSpeed;line-height:1.5;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif}ol[class],ul[class]{list-style:none}a:not([class]){text-decoration-skip-ink:auto}img{max-width:100%;display:block}article>*+*{margin-top:1em}button,input,select,textarea{font:inherit}@media (prefers-reduced-motion:reduce){*{animation-duration:.01ms!important;animation-iteration-count:1!important;transition-duration:.01ms!important;scroll-behavior:auto!important}}</style></head><body>`
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
