package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "goappuser"
	password = "password"
	dbname   = "goapp"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	// // Insert data
	// insertUser := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	// var userID int
	// err = db.QueryRow(insertUser, "John Doe", 30).Scan(&userID)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Inserted user with ID %d\n", userID)

	fmt.Printf("%T", psqlInfo)
}
