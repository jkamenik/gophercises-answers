package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"../../cyoa"
)

var tpl *template.Template
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{ range .Paragraphs }}
    <p>{{.}}</p>
    {{ end }}
    
    <ul>
    {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
    <ul>
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
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}
