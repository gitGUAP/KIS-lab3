package main

import (
	"database/sql"  // Интерфейс для работы со SQL-like БД
	"html/template" // Шаблоны для выдачи html страниц
	"log"           // Вывод информации в консоль
	"net/http"      // Для запуска HTTP сервера

	_ "github.com/mattn/go-sqlite3" // Драйвер для работы со SQLite3
)

// DB указатель на соединение с базой данных
var DB *sql.DB

// Реализация обработчика запроса
func handler(w http.ResponseWriter, r *http.Request) {
	// Выполнение запроса к базе данных
	rows, err := DB.Query(`SELECT name_cat FROM Category`)
	if err != nil {
		log.Fatal(err)
	}

	// Чтение шаблона из файла
	tmpl, _ := template.ParseFiles("tmpl/index.html")

	// Создание массива для снятия слепка с базы данных,
	names := []string{}
	// Итерируемся по всем строкам, который вернул запрос SQLite
	for rows.Next() {
		var temp string
		rows.Scan(&temp)
		// Записывает возвращенные данные в слепок
		names = append(names, temp)
	}

	// Вписываем данные в шаблон HTML страницы, дл отдачи пользователю
	tmpl.Execute(w, names)
}

func main() {
	var err error
	// Открытие соединения с БД SQLite3.db
	DB, err = sql.Open("sqlite3", "./sqlite.db")
	// Проверка установки соединения
	if err != nil {
		log.Fatal(err)
	}
	// Закрытие соединения с БД по выходу из функции main
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

	// Установка обработчика запроса по данному запросу
	http.HandleFunc("/", handler)

	log.Println("Listening...")
	// Запуск локального сервека на 8080 порту
	log.Fatal(http.ListenAndServe(":8080", nil))
}
