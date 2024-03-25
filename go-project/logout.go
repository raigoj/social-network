package main

import (
	"fmt"
	"net/http"

	"social-network/dbFunc"
)

// vaata yle
func logout(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	//fmt.Println(w.Header())
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	// delete the session
	err = dbFunc.DeleteSessionBySessionId(dbName, c.Value)
	if err != nil {
		fmt.Println(err)
	}
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusOK)
}
