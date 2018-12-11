package models

import (
	"log"
)

func InsertPhone(name string, number string) {

	res, err := db.Exec("INSERT INTO Phone VALUES($1, $2)", name, number)
	if err != nil{
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected = %d\n", rowCnt)
}