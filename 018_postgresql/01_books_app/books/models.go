package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/chhsu0222/go-web-dev/018_postgresql/01_books_app/config"
)

// Book is the struct represents a book's info.
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllBooks returns all the books in the collection as a slice of Book.
func AllBooks() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

// OneBook returns the book with the provided isbn
func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. bad request")
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1;", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// PutBook inserts a book into the collection.
func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. bad request. all fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. not acceptable. price must be a number")
	}
	bk.Price = float32(f64)

	// insert values
	insertStr := "INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4);"
	_, err = config.DB.Exec(insertStr, bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. internal server error." + err.Error())
	}
	return bk, nil
}

// UpdateBook updatess the book with the provided isbn
func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. bad request. all fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. not acceptable. price must be a number")
	}
	bk.Price = float32(f64)

	// update values
	updateStr := "UPDATE books SET isbn = $1, title = $2, author = $3, price = $4 WHERE isbn = $1;"
	_, err = config.DB.Exec(updateStr, bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// DeleteBook deletes the book with provided isbn
func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. bad request")
	}

	_, err := config.DB.Exec("DELETE FROM books WHERE isbn = $1;", isbn)
	if err != nil {
		return errors.New("500. internal server error")
	}
	return nil
}
