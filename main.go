package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func handler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`SELECT name_cat FROM Category`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}

func main() {
	DB, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	// CREATE TABLE Category (
	// 	id_cat   INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name_cat TEXT    NOT NULL
	// 					 UNIQUE,
	// 	url_cat  TEXT    NOT NULL
	// 					 UNIQUE
	// );

	// CREATE TABLE Subcategory (
	// 	id_subc   INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	id_cat            REFERENCES Category (id_cat),
	// 	name_subc TEXT    NOT NULL
	// 					  UNIQUE,
	// 	url_subc  TEXT    NOT NULL
	// 					  UNIQUE
	// );

	// CREATE TABLE Model (
	// 	id_mod      INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name_mod    TEXT    NOT NULL
	// 						UNIQUE ON CONFLICT ROLLBACK,
	// 	price       INTEGER,
	// 	country     TEXT,
	// 	manufacture TEXT,
	// 	weight      INTEGER
	// );

	http.HandleFunc("/", handler)
}
