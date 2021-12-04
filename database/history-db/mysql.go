package historydb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SelectAll() *[]string {
	var entries []string

	rows, err := db.Query("select * from history")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var entry string
		if err := rows.Scan(&entry); err != nil {
			log.Println(err.Error())
		} else {
			entries = append(entries, entry)
		}
	}
	return &entries
}

func Insert(entry string) {
	_, err := db.Exec("insert into history(entry) values(?)", entry)
	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	if db != nil {
		return
	}

	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUsername := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbDatabase := os.Getenv("MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS history (entry char(3))"); err != nil {
		panic(err.Error())
	}
}
