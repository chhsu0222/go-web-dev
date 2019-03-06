package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/chhsu0222/go-web-dev/017_mongodb/06_hands-on/models"
	"github.com/julienschmidt/httprouter"
)

// UserController is the struct with methods to handle different routes.
type UserController struct {
	session map[string]models.User
}

// NewUserController returns a pointer to UserController.
func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

// GetUser sends user's data back in json format.
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u, ok := uc.session[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateUser receives a new user's data in json format and creates a new user.
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	id, _ := uuid.NewV4()
	u.Id = id.String()

	// store the user
	uc.session[u.Id] = u

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteUser deletes the user's data from server.
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
