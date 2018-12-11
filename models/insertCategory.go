package models

import (
	"log"
)

func InsertCategory(name string, number int) {

	res, err := db.Exec("INSERT INTO Category VALUES($1, $2)", name, number)
	if err != nil{
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected = %d\n", rowCnt)
}