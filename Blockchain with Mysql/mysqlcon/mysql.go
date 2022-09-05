package mysqlcon

import (
	"database/sql"
	"log"
)

func getMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/blockchain/blockinfo?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

