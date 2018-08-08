package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	id   int
	name string
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/mac")

	rows, err := db.Query("select id, name from user where id > ?", 0)

	if err != nil {
		log.Fatal(err)
	}

	// 延迟释放
	defer db.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
