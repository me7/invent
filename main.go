package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type mainHandler struct{}

type page struct {
	Title   int
	Content string
	Footer  string
}

var tpl = make(map[string]*template.Template)

//must check for error
func must(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

func init() {
	tpl["index"] = template.Must(template.ParseFiles("tpl/base", "tpl/index"))
	tpl["index2"] = template.Must(template.ParseFiles("tpl/base", "tpl/index2"))
	p1 := page{Title: 1, Content: "custom content"}
	tpl["index"].Execute(os.Stdout, p1)
}

func main() {
	var mainHandler mainHandler
	m := http.NewServeMux()

	m.Handle("/", mainHandler)
	must(http.ListenAndServe(":7749", m))
}

func (h mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// must(tpl.ExecuteTemplate(w, "index.tpl", "hello gopher"))
}

func add(x, y int) int {
	return x + y
}
