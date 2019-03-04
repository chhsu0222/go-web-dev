package main

import (
	"net/http"
	"text/template"

	"github.com/satori/go.uuid"
)

type user struct {
	Email string
	First string
	Last  string
}

var tpl *template.Template
var dbUsers = map[string]user{}         // user email, user
var dbSession = make(map[string]string) // session ID, user email

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	if e, ok := dbSession[c.Value]; ok {
		u = dbUsers[e]
	}

	// process form submission
	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{
			Email: e,
			First: f,
			Last:  l,
		}
		dbSession[c.Value] = e
		dbUsers[e] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	e, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[e]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
