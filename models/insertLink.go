package models

import (
	"log"
	"fmt"
)

func InsertLink(name string, address string, category string, tags string) {

	fmt.Println("inserting a link for " + name)

	res, err := db.Exec("INSERT INTO Links VALUES($1, $2, $3, $4)", name, address, tags, category)
	if err != nil{
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected = %d\n", rowCnt)
}