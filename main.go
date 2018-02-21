package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

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

	defer db.Close()
}
