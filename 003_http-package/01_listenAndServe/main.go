package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// hotdog implicitly implements Handler interface with this method.
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
