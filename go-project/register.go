package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"social-network/dbFunc"
	"social-network/structs"

	"golang.org/x/crypto/bcrypt"
)

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	//w.WriteHeader(http.StatusCreated)

	switch r.Method {
	case "GET":
		fmt.Println("get")
	case "POST":
		fmt.Println("post")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		type Errors struct {
			UsernameTaken bool
			EmailTaken    bool
		}
		var message Errors

		var userObj structs.User
		json.Unmarshal(data, &userObj)
		userObj.Online = 0
		if _, err := dbFunc.GetUserByUsername(dbName, userObj.Username); err == nil {
			message.UsernameTaken = true
		}

		if _, err := dbFunc.GetUserByEmail(dbName, userObj.Email); err == nil {
			message.EmailTaken = true
		}
		if message.EmailTaken || message.UsernameTaken {
			//
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(message)
			return
		}

		password := []byte(userObj.Password)
		hashedPassword, err2 := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err2 != nil {
			panic(err2)
		}
		dbFunc.SetUser(dbName, userObj.Email, userObj.Username, string(hashedPassword), userObj.Firstname, userObj.Lastname, userObj.Gender, userObj.Age, userObj.Online)
		registrationSuccess := "User was registred"
		w.Write([]byte(registrationSuccess))

	}
}
