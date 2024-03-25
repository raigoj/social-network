package dbFunc

import (
	"database/sql"

	"social-network/structs"
)

// add error, if not able to setuser
func SetUser(dbName string, email string, username string, password string, firstname string, lastname string, gender string, dateofbirth string, online int) error {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	insertSQL := `INSERT INTO users(email, username, password, firstname, lastname, gender, dateofbirth, online) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec(email, username, password, firstname, lastname, gender, dateofbirth, online)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(dbName string, email string) (structs.User, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var user structs.User
	user.Email = email

	// Query for a value based on a single row.
	row := forumDB.QueryRow("SELECT id, username, password, firstname, lastname, gender, dateofbirth, online FROM users WHERE email = ?", user.Email)

	switch err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Online); err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return user, err
	default:
		panic(err)
	}
}

func GetUserByUsername(dbName string, username string) (structs.User, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var user structs.User
	user.Username = username

	// Query for a value based on a single row.
	row := forumDB.QueryRow("SELECT id, email, password, firstname, lastname, gender, dateofbirth, online FROM users WHERE username = ?", user.Username)

	switch err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Online); err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return user, err
	default:
		panic(err)
	}
}

func GetUserById(dbName string, userId int) (structs.User, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var user structs.User
	user.Id = userId

	// Query for a value based on a single row.
	row := forumDB.QueryRow("SELECT email, password, username, firstname, lastname, gender, dateofbirth, online FROM users WHERE id = ?", user.Id)

	switch err := row.Scan(&user.Email, &user.Password, &user.Username, &user.Firstname, &user.Lastname, &user.Gender, &user.Age, &user.Online); err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return user, err
	default:
		panic(err)
	}
}

func GetAllUsers(dbName string) ([]structs.User, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var user structs.User
	users := make([]structs.User, 0)
	rows, err := forumDB.Query("SELECT id, username, firstname, lastname, online FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		switch err = rows.Scan(&user.Id, &user.Username, &user.Firstname, &user.Lastname, &user.Online); err {
		case nil:
			users = append(users, user)
		default:
			return users, err
		}
	}
	return users, nil
}

func ChangeUserOnlineStatus(dbName string, username string, status int) error {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	_, err := forumDB.Exec("UPDATE users SET online = ? WHERE username = ?", status, username)
	if err != nil {
		return err
	}
	return nil

}
