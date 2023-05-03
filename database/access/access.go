package access

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/DevelopHub")
	if err != nil {
		fmt.Println("sql.Open error: ", err)
	}
}

func GetRowsCount() int {
	var rowsCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM question_table").Scan(&rowsCount); err != nil {
		fmt.Println("db.QueryRow error: ", err)
	}
	return rowsCount
}