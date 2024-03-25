package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"social-network/dbFunc"
	"social-network/structs"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	switch r.Method {
	case "GET":
		fmt.Println("get")
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var userObj structs.User
		json.Unmarshal(data, &userObj)
		type Errors struct {
			invalidUserOrPass bool
		}
		var message Errors
		var user structs.User

		if strings.Contains(userObj.Username, "@") {
			user, err = dbFunc.GetUserByEmail(dbName, userObj.Username)
			if err != nil {
				message.invalidUserOrPass = true
				fmt.Println("ei saa kasutajat k2tte")
			}
		} else {
			user, err = dbFunc.GetUserByUsername(dbName, userObj.Username)
			if err != nil {
				message.invalidUserOrPass = true
				fmt.Println("ei saa kasutajat k2tte")
			}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userObj.Password)); err != nil {
			message.invalidUserOrPass = true
			fmt.Println("passwordid ei matchi")
		}
		if message.invalidUserOrPass {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(message)
			return
		}

		sID := uuid.NewV4()

		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		err = dbFunc.SetSession(dbName, c.Value, user.Id)
		/*		for i, val := range dbSessions {
					if val.un == form.Username {
						delete(dbSessions, i)
					}
				}
				dbSessions[c.Value] = session{form.Username, time.Now()}*/
		w.WriteHeader(http.StatusOK)
	}
}
