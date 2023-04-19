package service

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var developHubDB *sql.DB = initDevelopHubDatabase()

// Database information
const (
	USER     = "root"
	PASSWORD = "1234"
	PORT     = 3306
	HOST     = "localhost"
	DBNAME   = "DevelopHub"
)

// Initionalizer DevelopHub Database
func initDevelopHubDatabase() *sql.DB {
	developHubDatabase, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USER, PASSWORD, HOST, PORT, DBNAME))
	if err != nil {
		// TODO: Recoder send!
	}
	return developHubDatabase
}

// Querys
func isSameUserNameEmail(username, email string) bool {
	var count int
	res := developHubDB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM user_table WHERE username='%s' OR email='%s'`, username, email))
	err := res.Scan(&count)

	if err != nil {
		// TODO: Recoder send!
	}
	return count > 0
}

func singUp(username, email, password string) {
	_, err := developHubDB.Exec(fmt.Sprintf(`INSERT INTO user_table(username, email, password) VALUES('%s', '%s', '%s')`, username, email, password))
	if err != nil {
		// TODO: Recoder send!
	}
}
