package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"social-network/dbFunc"
)

func makeevent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type uevent struct {
			Creator string
			Name    string
			Text    string
		}
		var event uevent
		json.Unmarshal(data, &event)
		userInt, err := dbFunc.GetUserByUsername(dbName, event.Creator)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dbFunc.CreateGroup(dbName, userInt.Id, event.Name, event.Text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func seteventstatus(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type ueve struct {
			event  string
			user   string
			status string
		}
		var eve ueve
		json.Unmarshal(data, &eve)
		reciver, err := dbFunc.GetUserByUsername(dbName, eve.user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		event, err := dbFunc.GetEventByName(dbName, eve.event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dbFunc.UpdateUserEventConnection(dbName, reciver.Id, event.Id, eve.status)
		if err == sql.ErrNoRows {
			err = dbFunc.AddUserEventConnection(dbName, reciver.Id, event.Id, eve.status)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
