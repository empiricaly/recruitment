package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playgroundServer("Empirica Recruitment GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

var playgroudPage = template.Must(template.New("graphiql").Parse(`<!DOCTYPE html>
<html>
<head>
	<meta charset=utf-8/>
	<meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
	<link rel="shortcut icon" href="https://graphcool-playground.netlify.com/favicon.png">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/static/css/index.css"
		integrity="{{ .cssSRI }}" crossorigin="anonymous"/>
	<link rel="shortcut icon" href="https://cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/favicon.png"
		crossorigin="anonymous"/>
	<script src="https://cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/static/js/middleware.js"
		integrity="{{ .jsSRI }}" crossorigin="anonymous"></script>
	<title>{{.title}}</title>
</head>
<body>
<style type="text/css">
	html { font-family: "Open Sans", sans-serif; overflow: hidden; }
	body { margin: 0; background: #172a3a; }
</style>
<div id="root"/>
<script type="text/javascript">
	window.addEventListener('load', function (event) {
		const root = document.getElementById('root');
		root.classList.add('playgroundIn');
		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
		GraphQLPlayground.init(root, {
			endpoint: location.protocol + '//' + location.host + '{{.endpoint}}',
			subscriptionsEndpoint: wsProto + '//' + location.host + '{{.endpoint }}',
			shareEnabled: true,
			workspaceName: "{{.title}}",
			settings: {
				'request.credentials': 'same-origin',
				'schema.polling.enable': false
			}
		})
	})
</script>
</body>
</html>
`))

func playgroundServer(title string, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := playgroudPage.Execute(w, map[string]string{
			"title":    title,
			"endpoint": endpoint,
			"version":  "1.7.26",
			"cssSRI":   "sha256-dKnNLEFwKSVFpkpjRWe+o/jQDM6n/JsvQ0J3l5Dk3fc=",
			"jsSRI":    "sha256-SG9YAy4eywTcLckwij7V4oSCG3hOdV1m+2e1XuNxIgk=",
		})
		if err != nil {
			panic(err)
		}
	}
}
