package main

import (
	"database/sql"
	"fmt"
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

	rows, err := db.Query("SELECT id, username FROM " + UserTable)
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

	err := db.QueryRow("SELECT id, username FROM "+UserTable+" WHERE id = $1", id).Scan(&usr.ID, &usr.Username)
	if err != nil {
		log.Fatalln(err)
	}

	return usr
}

func insertUser(usr user) error {
	db := getDb()
	defer db.Close()

	password, _ := HashPassword(usr.Password)

	_, err := db.Exec("INSERT INTO "+UserTable+" (username, password) VALUES ($1, $2)", usr.Username, password)
	if err != nil {
		return err
	}

	return nil
}

func dbLogin(usr user) {
	db := getDb()
	defer db.Close()

	username := usr.Username
	password := usr.Password
	var foundPassword string

	err := db.QueryRow("SELECT password FROM "+UserTable+" WHERE username = $1", username).Scan(&foundPassword)

	if err != nil {
		panic(err)
	}

	if CheckPasswordHash(password, foundPassword) {
		fmt.Printf("User %v exists\n", username)
	} else {
		fmt.Printf("User %v does not exist\n", username)
	}
}
