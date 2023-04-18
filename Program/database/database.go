package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USER     = "root"
	PASSWORD = "1234"
	PORT     = 3306
	HOST     = "localhost"
	DBNAME   = "DevelopHub"
)

var DBInfo string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USER, PASSWORD, HOST, PORT, DBNAME)

type DatabaseQuery struct {
}

// table structs
var user User = User{}

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", DBInfo)
	if err != nil {
		fmt.Println("[Database ERROR]: ", err)
	}
	return db
}

func NewDatabaseQuery() *DatabaseQuery {
	return &DatabaseQuery{}
}

func (d *DatabaseQuery) IsSameUsernameEmail(db *sql.DB, username string, email string) bool {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM user_table WHERE username='%s' OR email='%s'", username, email)
	res := db.QueryRow(query)
	err := res.Scan(&count)
	if err != nil {
		fmt.Println("[Database IsSameUsernameEmail ERROR]: ", err)
		return false
	}
	fmt.Println("return value", count > 0)
	return count > 0
}
