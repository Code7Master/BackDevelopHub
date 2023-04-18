package database

import (
	"BackDevelopHub/Program/log"
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

var logger log.Logger = *log.NewLogger()
var DBEnvironment string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", USER, PASSWORD, HOST, PORT, DBNAME)

type DatabaseQuerier struct {
}

// table structs
var user User = User{}

func InitDB() *sql.DB {
	logger.Record("[BackDevelopHub/Program/database/database.go (func InitDB)__GOOD] start")
	db, err := sql.Open("mysql", DBEnvironment)
	if err != nil {
		logger.Record(fmt.Sprintf("[BackDevelopHub/Program/database/database.go (func InitDB)__ERROR] Do not opne mysql: %s", err))
	}
	return db
}

func NewDatabaseQuerier() *DatabaseQuerier {
	return &DatabaseQuerier{}
}

// SELECT COUNT(*) FROM user_table WHERE username='JunBeomHan' OR email='1234'"
func (d *DatabaseQuerier) IsSameUsernameOREmail(db *sql.DB, username, email string) bool {
	logger.Record("[BackDevelopHub/Program/database/database.go (func IsSameUsernameOREmail)__GOOD] start")
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM user_table WHERE username='%s' OR email='%s'`, username, email)
	res := db.QueryRow(query)
	err := res.Scan(&count)
	if err != nil {
		logger.Record(fmt.Sprintf("[BackDevelopHub/Program/database/database.go (func IsSameUsernameOREmail)__ERROR] Do not read to sql query: %s", err))
		return false
	}
	return count > 0
}

// INSERT INTO user_table(username, email, passwrod) VALUES('JunBeomHan', 'agc@gmail.com', '123132')
func (d *DatabaseQuerier) SingUp(db *sql.DB, username, email, password string) {
	logger.Record("[BackDevelopHub/Program/database/database.go (func SingUp)__GOOD] start")
	query := fmt.Sprintf(`INSERT INTO user_table(username, email, password) VALUES('%s', '%s', '%s')`, username, email, password)
	_, err := db.Exec(query)

	if err != nil {
		logger.Record(fmt.Sprintf("[BackDevelopHub/Program/database/database.go (func SingUp)__ERROR] Do not insert to mysql: %s", err))
	}
}
