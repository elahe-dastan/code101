package main

import (
	//"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/golang-migrate/migrate"
	//"github.com/golang-migrate/migrate/database/mysql"
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

	//driver, err := mysql.WithInstance(db, &mysql.Config{})
	//if err != nil{
	//	log.Fatal(err)
	//}
	//
	//m, err := migrate.NewWithDatabaseInstance(
	//	"file://migrations",
	//	"parham",
	//	driver,
	//)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//
	//if err := m.Up(); err != nil{
	//	log.Fatal(err)
	//}

	// fill the table with data
	//for i := 1000; i < 5000; i++ {
	//	_, err := db.Exec("INSERT INTO parham VALUES (?, 'parham');", i)
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//}

	// execution time
	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	//defer cancel()
	start := time.Now()
	//rows, err := db.QueryContext(ctx,"SELECT * FROM parham;")
	rows := db.QueryRow("SELECT count(*) FROM parham;")
	end := time.Now()
	//if err != nil{
	//	log.Fatal(err)
	//}

	fmt.Println(end.Sub(start))

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, );err != nil{
			log.Fatal(err)
		}
		fmt.Println(id)
	}
}