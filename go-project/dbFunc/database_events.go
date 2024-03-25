package dbFunc

import (
	"database/sql"

	"social-network/structs"
)

type tmp struct {
	Yes []structs.UserEvent
	No  []structs.UserEvent
}

func CreateEvent(dbname string, creator int, name string, text string, time string) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	insertSQL := `INSERT INTO events(creator, title, description, time) VALUES (?, ?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(creator, name, text, time)
	if err != nil {
		return err
	}
	t, _ := GetEventByName(dbname, name)
	AddUserEventConnection(dbname, creator, t.Id, "Yes")
	return nil
}

func GetEventById(dbName string, eid int) (structs.Event, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var event structs.Event
	row := forumDB.QueryRow("SELECT * FROM events WHERE id = ?", eid)
	switch err := row.Scan(&event.Id, &event.Creator, &event.Name, &event.Text); err {
	case nil:
		return event, nil
	case sql.ErrNoRows:
		return event, err
	default:
		return event, err
	}
}

func GetEventByName(dbName string, name string) (structs.Event, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var event structs.Event
	row := forumDB.QueryRow("SELECT * FROM events WHERE id = ?", name)
	switch err := row.Scan(&event.Id, &event.Creator, &event.Name, &event.Text); err {
	case nil:
		return event, nil
	case sql.ErrNoRows:
		return event, err
	default:
		return event, err
	}
}

func AddUserEventConnection(dbname string, uid int, eid int, status string) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	insertSQL := `INSERT INTO userevents(uid, eid, status) VALUES (?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uid, eid, status)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserEventConnection(dbname string, uid int, eid int, status string) error {
	forumDB := OpenDatabase((dbname))
	defer forumDB.Close()
	sql := `UPDATE sessions SET status = $3 WHERE uid = $1 AND eid = $2`
	stmt, err := forumDB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uid, eid, status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(dbname string, eid int) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	insertSQL := `DELETE * FROM events WHERE eid = ?`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(eid)
	if err != nil {
		return err
	}
	return nil
}

func GetEventUsers(dbName string, eid int) (tmp, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var user structs.UserEvent
	var users tmp
	users.Yes = make([]structs.UserEvent, 0)
	users.No = make([]structs.UserEvent, 0)
	rows, err := forumDB.Query("SELECT * FROM userevents WHERE eid = ?", eid)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&user.Uid, &user.Eid, &user.Status); err {
		case nil:
			switch user.Status {
			case "Yes":
				users.Yes = append(users.Yes, user)
			case "No":
				users.No = append(users.No, user)
			}
		default:
			return users, err
		}
	}
	return users, nil
}
