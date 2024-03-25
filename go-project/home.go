package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"social-network/dbFunc"
	"social-network/structs"
)

func home(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
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
