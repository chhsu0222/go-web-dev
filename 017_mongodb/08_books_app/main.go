package main

import (
	"net/http"

	"github.com/chhsu0222/go-web-dev/017_mongodb/08_books_app/books"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/books", books.Index)
	r.GET("/books/show", books.Show)
	r.GET("/books/create", books.Create)
	r.POST("/books/create/process", books.CreateProcess)
	r.GET("/books/update", books.Update)
	r.POST("/books/update/process", books.UpdateProcess)
	r.GET("/books/delete/process", books.DeleteProcess)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
