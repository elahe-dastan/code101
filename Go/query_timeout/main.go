package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"log"
)

func main() {
	// Open up our database connection.
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	db, err := sql.Open("mysql", "root:parham@tcp(127.0.0.1:3306)")
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil{
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://insert_data.sql",
		"",
		driver,
	)
	if err != nil{
		log.Fatal(err)
	}

	m.Steps(2)

}