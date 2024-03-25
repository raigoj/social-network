package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"social-network/dbFunc"
	"social-network/structs"
)

func loadcomments(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type input struct {
			Postid string
		}
		var postid2 input
		json.Unmarshal(data, &postid2)

		type input2 struct {
			Postid int
		}
		var postid input2
		postid.Postid, _ = strconv.Atoi(postid2.Postid)
		var comments []structs.Comment
		comments, err = dbFunc.GetCommentsByPostId(dbName, postid.Postid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(comments)
	}
}

func createcomment(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "POST":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// todo: vaata yle
		type comment struct {
			Content  string
			Username string
			Postid   string
		}
		var commentObj2 comment
		json.Unmarshal(data, &commentObj2)

		type comment2 struct {
			Content  string
			Username string
			Postid   int
		}
		var commentObj comment2
		commentObj.Content = commentObj2.Content
		commentObj.Username = commentObj2.Username
		commentObj.Postid, _ = strconv.Atoi(commentObj2.Postid)
		fmt.Println(commentObj)
		err = dbFunc.SetComment(dbName, commentObj.Content, commentObj.Username, commentObj.Postid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	}
}
