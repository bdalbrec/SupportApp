package models

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/bdalbrec/sll"
	"github.com/bdalbrec/SupportApp/configs"
)

var db *sql.DB



func InitDB(dataSourceName string) {
	logname := configs.Configs.Logname
    var e, err error
    db, err = sql.Open("mssql", dataSourceName)
    if err != nil {
		e = sll.LogError("Error opening datatase.", logname, err)
		if e != nil {
			fmt.Println(e)
		}
    }

    if err = db.Ping(); err != nil {
		e = sll.LogError("Error pinging database", logname, err)
		if e != nil {
			fmt.Println(e)
		}
	}
	e = sll.LogInfo("Database connect and ping successful.", logname)
	if e  != nil {
		fmt.Println(e)
	}
}