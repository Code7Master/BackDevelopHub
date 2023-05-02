package access

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *mysql.DB

func Init() {
	var err error
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/DevelopHub")
	if err != nil {
		fmt.Println("sql.Open error: ", err)
	}
}

func GetCountRows() int {
	var rowCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM question_table").Scan(&rowCount); err != nil {
		fmt.Println("db.QueryRow error: ", err)
	}
	return rowCount
}