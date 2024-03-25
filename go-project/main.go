package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"social-network/dbFunc"

	_ "github.com/mattn/go-sqlite3"
)

const dbName string = "./database/forum-database.db"
const dbTableName string = "database/Forum.sql"
const sessionLength = 6000 //100minutit

func enableCors(w *http.ResponseWriter) {

	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	(*w).Header().Add("Access-Control-Allow-Credentials", "true")
	(*w).Header().Add("Access-Control-Allow-Method", "GET")
	(*w).Header().Add("Access-Control-Allow-Method", "POST")
	(*w).Header().Add("Content-Type", "application/json")
}

/*
type Session struct {
	Sessionid    string
	Userid       int
	Lastactivity time.Time
}
*/

func checkSessions() {
	for {
		sessions, err := dbFunc.GetAllSessions(dbName)
		if err != nil {
			fmt.Println(err)
		}
		for _, session := range sessions {
			diff := time.Now().Sub(session.Lastactivity).Seconds()
			if sessionLength < diff {
				err := dbFunc.DeleteSessionBySessionId(dbName, session.Sessionid)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("session kustutatud")
			}
		}
		duration := time.Duration(10) * time.Second
		time.Sleep(duration)
	}
}

/*
func sessionMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		fmt.Println(c)
		if err != nil {
			fmt.Println("cookiet, ei ole")
		} else {
			fmt.Println(c)
		}
		h(w, r)
	}
}*/

// if cookie expires, send info to frontend
func main() {
	dbFunc.StartDatabase(dbName, dbTableName)
	//dbFunc.SetCategory(dbName, "buy")
	//dbFunc.SetCategory(dbName, "sell")
	//dbFunc.SetCategory(dbName, "exchange")
	hub := newHub()
	go hub.run()
	go checkSessions()
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/createpost", createpost)
	http.HandleFunc("/home", home)
	http.HandleFunc("/", loadcomments)
	http.HandleFunc("/chats", chats)
	http.HandleFunc("/createcomment", createcomment)
	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Fatal(http.ListenAndServe(":8009", nil))
}

/*
CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "email" varchar UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "dateofbirth" DATE
*/

/*
responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObj BlackBox
	json.Unmarshal(responseData, &responseObj)
	var out []float64
	out = append(out, responseObj.Features[0].Geometry.Coordinates[0])
	out = append(out, responseObj.Features[0].Geometry.Coordinates[1])
	return out
*/
