package main

import (
	"html/template"
	"log"
	"net/http"
)

type mainHandler struct{}

var tpl *template.Template

//must check for error
func must(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.tpl"))
}

func main() {
	var mainHandler mainHandler
	m := http.NewServeMux()

	m.Handle("/", mainHandler)
	log.Fatal(http.ListenAndServe(":7749", m))
}

func (h mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	must(tpl.ExecuteTemplate(w, "index.tpl", "hello gopher"))
}
