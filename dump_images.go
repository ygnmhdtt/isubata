package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dump()
}
func dump() {
	dsn := "isucon:isucon@tcp(localhost:3306)/isubata?parseTime=true&loc=Local&charset=utf8mb4"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT distinct name, data FROM image")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		var data []byte
		err = rows.Scan(&name, &data)
		file, err := os.Create("/home/isucon/images/" + name)
		if err != nil {
			panic(err)
		}
		file.Write(data)
		fmt.Println("saving " + name)
	}
}
