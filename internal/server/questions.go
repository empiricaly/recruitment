package server

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aymerick/raymond"
	"github.com/empiricaly/recruitment/internal/ent"
	participantModel "github.com/empiricaly/recruitment/internal/ent/participant"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	stepRunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/empiricaly/recruitment/internal/js"
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
		console.log(params.get("assignmentId"));
		if (params.get("assignmentId") === "ASSIGNMENT_ID_NOT_AVAILABLE") {
			const notAssigned = document.querySelectorAll(".notAssigned");
			console.log(notAssigned);
			for (let i = 0; i < notAssigned.length; i++) {
				notAssigned[i].style.display = "block";
			}
			const assigned = document.querySelectorAll(".assigned");
			for (let i = 0; i < assigned.length; i++) {
				assigned[i].style.display = "none";
			}
		} else {
			const notAssigned = document.querySelectorAll(".notAssigned");
			console.log(1, notAssigned);
			for (let i = 0; i < notAssigned.length; i++) {
				notAssigned[i].style.display = "none";
			}
			const assigned = document.querySelectorAll(".assigned");
			for (let i = 0; i < assigned.length; i++) {
				assigned[i].style.display = "block";
			}
		}

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
					let value = el.value;
					if (el.type === "number") {
						value = parseFloat(value);
					}
					data[el.name] = value;
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
							method: "POST",
							headers: {
								"Content-Type": "application/json",
							},
							body: JSON.stringify(data),
						});

						loading.remove();

						const submitHit = () => {
							const assignmentId = document.createElement("input");
							assignmentId.name = "assignmentId";
							assignmentId.type = "hidden";
							assignmentId.value = params.get("assignmentId");

							f.action = params.get("turkSubmitTo") + "/mturk/externalSubmit";
							f.method = "POST";
							f.appendChild(assignmentId);
							f.submit();
						};

						if (response.ok) {
							submitHit();
							return;
						}

						console.error("error ", response.status, response.statusText);

						let textMessage =
							"Something went wrong. Please try again or report back to us.";
						const errorResponse = await response.json();
						let showButton = true;

						if (errorResponse && errorResponse.message) {
							if (errorResponse.message === "stepRunEnded") {
								submitHit();
							}
						}

						const div = document.createElement("div");
						div.style.backgroundColor = "#e53e3e";
						div.style.padding = "20px";
						div.style.color = "#f7fafc";

						const text = document.createElement("p");
						text.style.fontSize = "24px";
						text.innerText = textMessage;
						div.appendChild(text);

						if (showButton) {
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
							div.appendChild(button);
						}

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

		workerID := c.Query("workerId")
		if id == "" {
			log.Error().Msg("answers handler: missing workerID")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		assignmentID := c.Query("assignmentId")
		if id == "" {
			log.Error().Msg("answers handler: missing assignmentID")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		hitID := c.Query("hitId")
		if id == "" {
			log.Error().Msg("answers handler: missing hitID")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		stepRun, err := s.storeConn.StepRun.
			Query().
			WithStep(func(step *ent.StepQuery) {
				step.WithTemplate(func(template *ent.TemplateQuery) {
					template.WithSteps()
				})
			}).
			WithRun(func(run *ent.RunQuery) {
				run.WithSteps()
			}).
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

		if stepRun.StartedAt == nil {
			log.Error().Err(err).Msg("stepRun is not running yet")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if stepRun.EndedAt != nil {
			timeExtension := stepRun.EndedAt.Add(time.Minute * time.Duration(step.HitArgs.Timeout))
			remainingTime := timeExtension.Sub(time.Now())

			if remainingTime < 0 {
				log.Error().Err(err).Msg("stepRun has ended")
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
		}

		run, err := stepRun.Edges.RunOrErr()
		if err != nil {
			log.Error().Err(err).Msg("get run")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		stepRuns, err := run.Edges.StepsOrErr()
		if err != nil {
			log.Error().Err(err).Msg("get step runs")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		template, err := step.Edges.TemplateOrErr()
		if err != nil {
			log.Error().Err(err).Msg("get template")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		steps, err := template.Edges.StepsOrErr()
		if err != nil {
			log.Error().Err(err).Msg("get steps")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		switch step.MsgArgs.MessageType {
		case model.ContentTypeHTML:
			content := step.MsgArgs.Message
			if step.MsgArgs.URL != nil {
				currentParticipant, err := stepRun.QueryParticipants().Where(participantModel.MturkWorkerID(workerID)).First(c.Request.Context())
				if err != nil {
					log.Error().Err(err).Msg("get steps")
					c.AbortWithStatus(http.StatusNotFound)
				}
				urlString, err := js.UrlJS(currentParticipant, stepRun, steps, run, *step.MsgArgs.URL)
				u, err := url.Parse(urlString)
				if err != nil {
					log.Error().Err(err).Msg("invalid HIT message URL")
				} else {
					q := u.Query()
					q.Set("workerId", workerID)
					q.Set("assignmentId", assignmentID)
					q.Set("hitId", hitID)
					u.RawQuery = q.Encode()

					rsteps := make([]*model.RenderStep, len(steps))
					t := *run.StartedAt
					for i, s := range steps {
						var startsAt, startedAt, endedAt string
						if stepRuns[i].Index <= stepRun.Index {
							startedAt = stepRun.StartedAt.Format(time.Kitchen)
							startsAt = startedAt

							if stepRuns[i].Index < stepRun.Index && stepRun.EndedAt != nil {
								endedAt = stepRun.EndedAt.Format(time.Kitchen)
							}
						} else {
							startsAt = t.Format(time.Kitchen)
							t = t.Add(time.Duration(step.Duration) * time.Minute)
						}
						rsteps[i] = &model.RenderStep{
							Index:             s.Index,
							Duration:          s.Duration,
							ParticipantsCount: stepRuns[i].ParticipantsCount,
							Type:              s.Type.String(),
							StartsAt:          startsAt,
							StartedAt:         startedAt,
							EndedAt:           endedAt,
						}
					}

					startedAt := stepRun.StartedAt.Format(time.Kitchen)
					renderCtx := &model.RenderContext{
						URL: u.String(),
						Template: &model.RenderTemplate{
							Adult:            template.Adult,
							Sandbox:          template.Sandbox,
							SelectionType:    template.SelectionType.String(),
							Name:             template.Name,
							ParticipantCount: template.ParticipantCount,
						},
						Run: &model.RenderRun{
							Name:      run.Name,
							StartedAt: run.StartedAt.Format(time.Kitchen),
						},
						Step: &model.RenderStep{
							Index:             step.Index,
							Duration:          step.Duration,
							ParticipantsCount: stepRun.ParticipantsCount,
							Type:              step.Type.String(),
							StartsAt:          startedAt,
							StartedAt:         startedAt,
						},
						Steps: rsteps,
						Participant: &model.RenderParticipant{
							WorkerID:     workerID,
							HITID:        hitID,
							AssignmentID: assignmentID,
						},
					}
					r, err := raymond.Render(content, renderCtx)
					if err != nil {
						log.Error().Err(err).Msg("failed to render HTML message")
					} else {
						content = r
					}
				}
			}

			var out string
			if strings.Contains(step.MsgArgs.Message, "<html>") {
				out = content
			} else {
				out = htmlHead + content + htmlFoot
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

			content := buf.String()
			if step.MsgArgs.URL != nil {
				u, err := url.Parse(*step.MsgArgs.URL)
				if err != nil {
					log.Error().Err(err).Msg("invalid HIT message URL")
				} else {
					q := u.Query()
					q.Set("workerId", workerID)
					q.Set("assignmentId", assignmentID)
					q.Set("hitId", hitID)
					u.RawQuery = q.Encode()
					renderCtx := &model.RenderContext{
						URL: u.String(),
					}
					r, err := raymond.Render(content, renderCtx)
					if err != nil {
						log.Error().Err(err).Msg("failed to render HTML message")
					} else {
						content = r
					}
				}
			}

			c.Header("Content-Type", "text/html; charset=utf-8")
			log.Debug().Str("content", content).Msg("html message")
			c.String(200, htmlHead+content+htmlFoot)
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
