package books

import (
	"net/http"

	"github.com/chhsu0222/go-web-dev/017_mongodb/08_books_app/config"

	"github.com/julienschmidt/httprouter"
)

// Index handles the request for showing all the books.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bks, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
}

// Show handles the request for a specific book.
func Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.gohtml", bk)
}

// Create renders a page for creating a new book.
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

// CreateProcess handles the form submission for creating a new book.
func CreateProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bk, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", bk)
}

// Update renders a page for updating a book.
func Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", bk)
}

// UpdateProcess handles the form submission for updating a book.
func UpdateProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bk, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "updated.gohtml", bk)
}

// DeleteProcess handles the request for deleting a book.
func DeleteProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
