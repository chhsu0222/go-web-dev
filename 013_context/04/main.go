package main

import (
	"context"
	"fmt"
	"net/http"
)

type favContextKey string

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, favContextKey("userID"), 777)
	ctx = context.WithValue(ctx, favContextKey("fname"), "Bond")

	results := dbAccess(ctx)

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value(favContextKey("userID")).(int)
	return uid
}

// Assertion example: https://tour.golang.org/methods/15
