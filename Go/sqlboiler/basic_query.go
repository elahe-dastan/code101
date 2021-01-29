package main

import (
	"code101/Go/sqlboiler/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:parham@127.0.0.1:5432/parham?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	count, err := models.Pilots().Count(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

	p1 := models.Pilot{
		Name: "parham",
	}

	err = p1.Insert(context.Background(), db, boil.Infer())
	if err != nil{
		log.Fatal(err)
	}
}
