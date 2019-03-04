package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// req.Form is only available after ParseForm is called.
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// The type of req.Form is Value, which is map[string][]string
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
