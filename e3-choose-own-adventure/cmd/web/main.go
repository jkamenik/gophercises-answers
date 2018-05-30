package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"../../cyoa"
)

var tpl *template.Template
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
		<style>
			body {
				font-family: helvetica, arial;
			}
			h1 {
				text-align: center;
				position: relative;
			}
			.page {
				width: 80%;
				max-width: 500px;
				margin: auto;
				margin-top: 40px;
				margin-bottom: 40px;
				padding: 80px;
				background: #FFFCF6;
				border: 1px solid #eee;
				box-shadow: 0 10px 6px -6 px #777;
			}
			ul {
				border-top: 1px dotted #ccc;
				padding: 10px 0 0 0;
				-webkit-padding-start: 0;
			}
			li {
				padding-top: 10px;
			}
			a, a:visited {
				text-decoration: none;
				color: #6295b5;
			}
			a:active, a:hover {
				color: #7792a2;
			}
			p {
				text-indent: 1em;
			}
		</style>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>
			{{ range .Paragraphs }}
			<p>{{.}}</p>
			{{ end }}
							<ul>
			{{range .Options}}
				<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
			<ul>
		</section>
	</body>
</html>
`

func main() {
	port := flag.Int("port", 3000, "the port to start the server on.")
	file := flag.String("file", "gopher.json", "the JSON file with the story.")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.ParseJSON(f)
	if err != nil {
		panic(err)
	}

	h := newHandler(story)
	fmt.Printf("Stargind server on %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func newHandler(s cyoa.Story) http.Handler {
	return handler{s}
}

type handler struct {
	s cyoa.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// remove leading "/"
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v\n", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}
