package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		id, idErr := uuid.NewV4()
		if idErr != nil {
			log.Fatalln(idErr)
		}
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
	}
	fmt.Fprintln(w, c.String())
}
