package dbFunc

import (
	"database/sql"

	"social-network/structs"
)

func GetCategoryById(dbName string, id int) (structs.Category, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var category structs.Category
	category.Id = id

	row := forumDB.QueryRow("SELECT name FROM categories WHERE id = ?", id)
	switch err := row.Scan(&category.Name); err {
	case nil:
		return category, nil
	case sql.ErrNoRows:
		return category, err
	default:
		return category, err
	}
}

func GetCategoryByName(dbName string, name string) (structs.Category, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()
	var category structs.Category
	category.Name = name

	row := forumDB.QueryRow("SELECT id FROM categories WHERE name = ?", name)
	switch err := row.Scan(&category.Id); err {
	case nil:
		return category, nil
	case sql.ErrNoRows:
		return category, err
	default:
		return category, err
	}
}

func SetCategory(dbName string, name string) error {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	insertSQL := `INSERT INTO CATEGORIES(name) VALUES (?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	if err != nil {
		return err
	}
	return nil
}
