package main

import (
	"code101/Go/sqlboiler_userauth/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open handle to database like normal
	db, err := sql.Open("mysql", "root:parham@tcp(127.0.0.1:3306)/parham?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	//Insert
	//p1 := models.Passenger{
	//	ID:                        12,
	//	IsCellphoneVerified:       false,
	//	IsEmailVerified:           false,
	//	IsRegisteredWithGoogle:    false,
	//	EmailVerificationCode:     "73",
	//	CellphoneVerificationCode: "78",
	//	IsBlocked:                 false,
	//	AdjustFingerprint:         "panda",
	//	ComapnyID:                 0,
	//}
	//
	//err = p1.Insert(context.Background(), db, boil.Infer())
	//if err != nil{
	//	log.Fatal(err)
	//}

	// Update
	//passengers, err := models.Passengers(Where("adjust_fingerprint = 'chi shode'")).All(context.Background(), db)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//
	//rowsAff, err := passengers.UpdateAll(context.Background(), db, models.M{"adjust_fingerprint": "love gerdali"})
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(rowsAff)

	// SELECT
	//var p []models.Passenger
	//count, err := models.Passengers().Count(context.Background(), db)
	//fmt.Println(count)
	//err = models.NewQuery(qm.Select("*"), qm.From("passenger")).Bind(context.Background(), db, &p)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(p)

	// DELETE
	
}