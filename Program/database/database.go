package database

import (
	"database/sql"
	"fmt"
	"log"

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

// SELECT COUNT(*) FROM user_table WHERE username='JunBeomHan' OR email='1234'"
func (d *DatabaseQuery) IsSameUsernameEmail(db *sql.DB, username, email string) bool {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM user_table WHERE username='%s' OR email='%s'`, username, email)
	res := db.QueryRow(query)
	err := res.Scan(&count)
	if err != nil {
		fmt.Println("[Database IsSameUsernameEmail ERROR]: ", err)
		return false
	}
	return count > 0
}

// INSERT INTO user_table(username, email, passwrod) VALUES('JunBeomHan', 'agc@gmail.com', '123132')
func (d *DatabaseQuery) SingUp(db *sql.DB, username, email, password string) {
	query := fmt.Sprintf(`INSERT INTO user_table(username, email, password) VALUES('%s', '%s', '%s')`, username, email, password)
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}
