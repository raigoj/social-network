package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"social-network/dbFunc"
)

func createpost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type post struct {
			Title    string
			Content  string
			Category string
			Username string
		}
		var postObj post
		json.Unmarshal(data, &postObj)
		userInt, err := dbFunc.GetUserByUsername(dbName, postObj.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		categoryInt, err := dbFunc.GetCategoryByName(dbName, postObj.Category)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("hello")
		err = dbFunc.SetPost(dbName, postObj.Title, postObj.Content, userInt.Id, categoryInt.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}
