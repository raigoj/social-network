package dbFunc

import (
	"database/sql"
	"fmt"
	"time"

	"social-network/structs"
)

func SetSession(dbname string, sessionid string, userid int) error {
	//open database
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	//prepare statement
	lastactivity := time.Now()
	insertSQL := `INSERT INTO sessions(sessionid, userid, lastactivity) VALUES (?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	//insert data to db
	_, err = statement.Exec(sessionid, userid, lastactivity)
	if err != nil {
		return err
	}
	return nil
	//check for errors

}

func UpdateSession(dbname string, userid int) error {
	forumDB := OpenDatabase((dbname))
	defer forumDB.Close()
	t := time.Now()
	sql := `UPDATE sessions SET lastactivity = $2 WHERE userid = $1`
	stmt, err := forumDB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userid, t)
	if err != nil {
		return err
	}
	return nil
}

func SessionCheck(dbname string, userid int) (int, error) {
	ses, err := GetSessionByUserId(dbname, userid)
	if err != nil {
		return 0, err
	}
	diff := time.Now().Sub(ses.Lastactivity).Seconds()
	if 600 < diff {
		err := DeleteSessionBySessionId(dbname, ses.Sessionid)
		if err != nil {
			return 0, err
		}
		fmt.Println("session kustutatud")
		return 1, nil
	}
	return 0, nil
}

func GetSessionByUserId(dbname string, userid int) (structs.Session, error) {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()

	var session structs.Session
	session.Userid = userid

	row := forumDB.QueryRow("SELECT sessionid, lastactivity FROM sessions WHERE userid = ?", userid)
	switch err := row.Scan(&session.Sessionid, &session.Lastactivity); err {
	case nil:
		return session, nil
	case sql.ErrNoRows:
		return session, err
	default:
		panic(err)
	}
}

func GetSessionBySessionId(dbname string, sessionid string) (structs.Session, error) {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()

	var session structs.Session
	session.Sessionid = sessionid

	row := forumDB.QueryRow("SELECT userid, lastactivity FROM sessions WHERE sessionid = ?", sessionid)
	switch err := row.Scan(&session.Userid, &session.Lastactivity); err {
	case nil:
		return session, nil
	case sql.ErrNoRows:
		return session, err
	default:
		panic(err)
	}
}

func GetAllSessions(dbName string) ([]structs.Session, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var session structs.Session
	sessions := make([]structs.Session, 0)
	rows, err := forumDB.Query("SELECT sessionid, userid, lastactivity FROM sessions")
	if err != nil {
		return sessions, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&session.Sessionid, &session.Userid, &session.Lastactivity); err {
		case nil:
			sessions = append(sessions, session)
		default:
			return sessions, err
		}
	}
	return sessions, nil
}

func DeleteSessionBySessionId(dbname string, sessionid string) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()

	deleteSQL := `DELETE FROM sessions WHERE sessionid = ?`
	statement, err := forumDB.Prepare(deleteSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(sessionid)
	if err != nil {
		return err
	}
	return nil

}
