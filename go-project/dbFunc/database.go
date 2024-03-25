package dbFunc

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
)

// This is the one call you'll need to create and initialise the database. It won't reset the data in it if there's already an existing database.
func StartDatabase(dbName string, dbTableName string) {
	log.Println("Starting " + dbName + "...")
	f, err := os.OpenFile(dbName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	CreateTables(dbName, dbTableName)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println(dbName + " started")
}

/*
const dbName string = "./database/forum-database.db"
const dbTableName string = "database/Forum.sql"

func main() {
	dbFunc.StartDatabase(dbName, dbTableName)
*/

// Called internally here by StartDatabase()
func CreateTables(dbName string, dbTableName string) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	query, err := ioutil.ReadFile(dbTableName)
	if err != nil {
		log.Println(query)
		panic(err)
	}
	if _, err := forumDB.Exec(string(query)); err != nil {
		log.Println(string(query))
		panic(err)
	}
}

// Called internally here by various functions
func OpenDatabase(dbName string) *sql.DB {
	forumDB, _ := sql.Open("sqlite3", dbName)
	return forumDB
}

// Some querying example
/*
func canPurchase(id int, quantity int) (bool, error) {
    var enough bool
    // Query for a value based on a single row.
    if err := db.QueryRow("SELECT (quantity >= ?) from album where id = ?",
        quantity, id).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return false, fmt.Errorf("canPurchase %d: unknown album", id)
        }
        return false, fmt.Errorf("canPurchase %d: %v", id)
    }
    return enough, nil
}
*/
