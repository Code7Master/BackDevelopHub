package access

import (
	"BackDevelopHub/dhjson"
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
	query := "SELECT COUNT(*) FROM question_table"

	var rowsCount int
	if err := db.QueryRow(query).Scan(&rowsCount); err != nil {
		fmt.Println("QueryRow error: ", err)
	}
	return rowsCount
}

func GetQuestionPreviewByIndex(question_idx int) dhjson.QuestionPreview {
	query := fmt.Sprintf("SELECT * FROM question_table WHERE idx=%d", question_idx)
	
	var questionPreview dhjson.QuestionPreview
	var ignore interface{}
	if err := db.QueryRow(query).Scan(&ignore, &questionPreview.Writer, &questionPreview.Title, &ignore, &questionPreview.View_count, &questionPreview.Answer_count, &questionPreview.Vote_count); err != nil {
		fmt.Println("GetQuestionPreviewByIndex error: ", err)
	}
	return questionPreview
}

func GetQuestionDetailByIndex(question_idx int) dhjson.QuestionDetail {
	query := fmt.Sprintf("SELECT * FROM question_table WHERE idx=%d", question_idx) 

	var questionDetail dhjson.QuestionDetail
	var ignore interface{}
	if err := db.QueryRow(query).Scan(&ignore, &questionDetail.Writer, &questionDetail.Title, &questionDetail.Content, &questionDetail.View_count, &questionDetail.Answer_count, &questionDetail.Vote_count); err != nil {
		fmt.Println("GetQuestionDetailByIndex error: ", err)
	}
	return questionDetail
}