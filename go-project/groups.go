package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"social-network/dbFunc"
	"social-network/structs"
)

func makegroup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type ugroup struct {
			Creator string
			Name    string
			Text    string
		}
		var group ugroup
		json.Unmarshal(data, &group)
		userInt, err := dbFunc.GetUserByUsername(dbName, group.Creator)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dbFunc.CreateGroup(dbName, userInt.Id, group.Name, group.Text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func acceptinvite(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type uinv struct {
			group   string
			reciver string
		}
		var inv uinv
		json.Unmarshal(data, &inv)
		reciver, err := dbFunc.GetUserByUsername(dbName, inv.reciver)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		group, err := dbFunc.GetGroupByName(dbName, inv.group)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dbFunc.AddUserGroupConnection(dbName, reciver.Id, group.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = dbFunc.DeleteGroupInv(dbName, reciver.Id, group.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func viewgroups(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type input struct {
			Category int
		}
		var category input
		json.Unmarshal(data, &category)
		var posts []structs.Posts
		if category.Category == 0 {
			posts, err = dbFunc.GetAllPosts(dbName)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			posts, err = dbFunc.GetPostsByCategory(dbName, category.Category)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(posts)
	}
}
