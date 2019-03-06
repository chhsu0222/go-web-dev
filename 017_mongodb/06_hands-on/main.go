package main

import (
	"net/http"

	"github.com/chhsu0222/go-web-dev/017_mongodb/06_hands-on/controllers"
	"github.com/chhsu0222/go-web-dev/017_mongodb/06_hands-on/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}
