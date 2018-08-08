package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var (
	id   int
	name string
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/mac")
	rows, err := db.Query("select id, name from user where id > ?", 0)
	checkErr(err)

	// 延迟释放
	defer db.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		checkErr(err)

		log.Println(id, name)
		fmt.Println(id, name)
	}
	err = rows.Err()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
