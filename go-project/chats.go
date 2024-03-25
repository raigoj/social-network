package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"social-network/dbFunc"
	"social-network/structs"
)

func chatHistory(w http.ResponseWriter, r *http.Request) {

}

func chats(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	switch r.Method {
	case "POST":
		fmt.Println("hi")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("hi2")
		type input struct {
			Username string
		}
		var loggedInUser input
		json.Unmarshal(data, &loggedInUser)
		var users []structs.User
		users, err = dbFunc.GetAllUsers(dbName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		/*	var sessions []structs.Session
			sessions, err = dbFunc.GetAllSessions(dbName)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}*/
		fmt.Println(users)
		fmt.Println("tere")
		json.NewEncoder(w).Encode(users)

	}
}
