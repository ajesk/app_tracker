package main

import (
	"database/sql"
	"log"
)

func convertToUser(rows *sql.Rows) user {
	var usr user

	rows.Scan(&usr.ID, &usr.Username)
	return usr
}

func getDbUsers() []user {
	db := getDb()
	defer db.Close()

	rows, err := db.Query("SELECT id, username FROM app_track.users")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var users []user

	for rows.Next() {
		users = append(users, convertToUser(rows))
	}

	return users
}

func getDbUserById(id int) (usr user) {
	db := getDb()
	defer db.Close()

	err := db.QueryRow("SELECT id, username FROM app_track.users WHERE id = $1", id).Scan(&usr.ID, &usr.Username)
	if err != nil {
		log.Fatalln(err)
	}

	return usr
}

func insertUser(usr user) error {
	db := getDb()
	defer db.Close()

	// Insert a new record
	password, _ := HashPassword(usr.Password)

	_, err := db.Exec("INSERT INTO app_track.users (username, password) VALUES ($1, $2)", usr.Username, password)
	if err != nil {
		return err
	}

	return nil
}
