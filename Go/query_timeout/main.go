package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	// Open up our database connection.
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	db, err := sql.Open("mysql", "root:parham@tcp(127.0.0.1:3306)/parham")
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
		"file://create_table.sql",
		"parham",
		driver,
	)
	if err != nil{
		fmt.Println("a")
		log.Fatal(err)
	}

	if err := m.Up(); err != nil{
		log.Fatal(err)
	}

}