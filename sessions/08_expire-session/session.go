package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if s, ok := dbSession[c.Value]; ok {
		// update lastActivity
		s.lastActivity = time.Now()
		dbSession[c.Value] = s
		u = dbUsers[s.e]
		// refresh session
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSession[c.Value]
	if ok {
		// update lastActivity
		s.lastActivity = time.Now()
		dbSession[c.Value] = s
		// refresh session
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
	}
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes

	for k, v := range dbSession {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSession, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("*************")
	for k, v := range dbSession {
		fmt.Println(k, v.e)
	}
	fmt.Println("")
}
