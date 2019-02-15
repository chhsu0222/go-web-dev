package main

import "net/http"

func getUser(req *http.Request) user {
	var u user
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if e, ok := dbSession[c.Value]; ok {
		u = dbUsers[e]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	e := dbSession[c.Value]
	_, ok := dbUsers[e]
	return ok
}
