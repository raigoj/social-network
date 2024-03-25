package dbFunc

import (
	"database/sql"

	"social-network/structs"
)

func CreateGroup(dbname string, creator int, name string, text string) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	insertSQL := `INSERT INTO groups(creator, name, description) VALUES (?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(creator, name, text)
	if err != nil {
		return err
	}
	t, _ := GetGroupByName(dbname, name)
	AddUserGroupConnection(dbname, creator, t.Id)
	return nil
}

func GetGroupById(dbName string, gid int) (structs.Group, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var group structs.Group
	row := forumDB.QueryRow("SELECT * FROM groups WHERE id = ?", gid)
	switch err := row.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
	case nil:
		return group, nil
	case sql.ErrNoRows:
		return group, err
	default:
		return group, err
	}
}

func GetGroupsByCreator(dbName string, creator int) ([]structs.Group, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var group structs.Group
	groups := make([]structs.Group, 0)
	rows, err := forumDB.Query("SELECT * FROM groups WHERE creator = ?", creator)
	if err != nil {
		return groups, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
		case nil:
			groups = append(groups, group)
		default:
			return groups, err
		}
	}
	return groups, nil
}

func GetGroupByName(dbName string, name string) (structs.Group, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var group structs.Group
	row := forumDB.QueryRow("SELECT * FROM groups WHERE name = ?", name)
	switch err := row.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
	case nil:
		return group, nil
	case sql.ErrNoRows:
		return group, err
	default:
		return group, err
	}
}

func AddUserGroupConnection(dbname string, uid int, gid int) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	insertSQL := `INSERT INTO usergroups(uid, gid) VALUES (?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uid, gid)
	if err != nil {
		return err
	}
	return nil
}

func GetUserGroups(dbname string, uid int) ([]structs.UserGoup, error) {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	var ug structs.UserGoup
	ugs := make([]structs.UserGoup, 0)
	rows, err := forumDB.Query("SELECT * FROM usergroups WHERE uid = ?", uid)
	if err != nil {
		return ugs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&ug.Uid, ug.Gid); err {
		case nil:
			ugs = append(ugs, ug)
		default:
			return ugs, err
		}
	}
	return ugs, nil
}

func GetGroupUsers(dbname string, gid int) ([]structs.UserGoup, error) {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	var ug structs.UserGoup
	ugs := make([]structs.UserGoup, 0)
	rows, err := forumDB.Query("SELECT * FROM usergroups WHERE gid = ?", gid)
	if err != nil {
		return ugs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&ug.Uid, ug.Gid); err {
		case nil:
			ugs = append(ugs, ug)
		default:
			return ugs, err
		}
	}
	return ugs, nil
}

func RemoveUserFromGroup(dbname string, uid int, gid int) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	_, err := forumDB.Exec("DELETE FROM usergroups WHERE uid = ? AND gid = ?", uid, gid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func CreateGroupInvite(dbname string, uid int, rid int, gid int) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	insertSQL := `INSERT INTO groupinvs(uid, rid, gid) VALUES (?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uid, rid, gid)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInvites(dbname string, rid int) ([]structs.Invite, error) {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	var inv structs.Invite
	invs := make([]structs.Invite, 0)
	rows, err := forumDB.Query("SELECT * FROM groupinvs WHERE rid = ?", rid)
	if err != nil {
		return invs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&inv.Sid, &inv.Rid, &inv.Gid); err {
		case nil:
			invs = append(invs, inv)
		default:
			return invs, err
		}
	}
	return invs, nil
}

func DeleteSpecificInv(dbname string, sid int, rid int, gid int) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	_, err := forumDB.Exec("DELETE FROM groupinvs WHERE sid = ? AND rid = ? AND gid = ?", sid, rid, gid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}
func DeleteGroupInv(dbname string, rid int, gid int) error {
	forumDB := OpenDatabase(dbname)
	defer forumDB.Close()
	_, err := forumDB.Exec("DELETE FROM groupinvs WHERE rid = ? AND gid = ?", rid, gid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}
