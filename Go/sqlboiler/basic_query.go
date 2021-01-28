package main

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"

	// Import this so we don't have to use qm.Limit etc.
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:parham@tcp(127.0.0.1:3306)/parham")
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)

}
