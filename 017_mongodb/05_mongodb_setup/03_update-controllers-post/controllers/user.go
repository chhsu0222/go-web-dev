package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/chhsu0222/go-web-dev/017_mongodb/05_mongodb_setup/03_update-controllers-post/models"
	"github.com/julienschmidt/httprouter"
)

// UserController is the struct with methods to handle different routes.
type UserController struct {
	session *mgo.Session
}

// NewUserController returns a pointer to UserController.
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// Methods have to be capitalized to be exported, eg, GetUser and not getUser

// GetUser sends user's data back in json format.
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// using "Encode" instead of "Marshal"
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateUser receives a new user's data in json format and creates a new user.
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// composite literal - type and curly braces
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = bson.NewObjectId()

	// store the user in mongodb
	err := uc.session.DB("go-web-dev-db").C("users").Insert(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write content-type, statuscode
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteUser deletes the user's data from server.
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	fmt.Fprint(w, "Write code to delete user\n")
}
