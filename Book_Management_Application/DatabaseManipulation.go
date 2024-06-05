package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var BookNum int = 0

type Book struct {
	id    int
	Title string
}

var BooksList []Book

// func initializeDatabase() {
// 	// dbHost := os.Getenv("DB_HOST")
// 	// dbUser := os.Getenv("DB_USER")
// 	// dbPassword := os.Getenv("DB_PASSWORD")
// 	// dbName := os.Getenv("DB_NAME")

// 	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
// 	// fmt.Printf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
// 	// db, err := sql.Open("mysql", dsn)
// 	// if err != nil {
// 	// 	log.Fatal("Error connecting to the database:", err)
// 	// }
// 	// defer db.Close()

// 	ConnectToDatabase()

// 	Db.Close()

// }

func ConnectToDatabase() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "Nta21012001",
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: "BooksList",
	}
	// Get a database handle.
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		fmt.Println("trouble here")
	}
	fmt.Println(Db)

	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	// var err error
	// Db, err = sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal("Error connecting to the database:", err)
	// }

	fmt.Println("Connected!HeHeHe")
}

func GetBooksList() {

	BooksList = make([]Book, 0)
	ConnectToDatabase()

	var sqlStatement string = "SELECT * FROM BooksList"
	rows, err := Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	BookNum = 0

	for rows.Next() {

		var book Book
		err := rows.Scan(&book.id, &book.Title)
		if err != nil {
			log.Fatal(err)
		}
		BookNum++
		BooksList = append(BooksList, book)
	}
	//rows.Close()
	Db.Close()

	//return BookNum, BooksList

}

func addBook(bookTitle string) {
	ConnectToDatabase()

	var sqlStatement string = "INSERT INTO BooksList (name) VALUES (?)"

	result, err := Db.Exec(sqlStatement, bookTitle)
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %d rows\n", rowCount)

	Db.Close()

	GetBooksList()
}
