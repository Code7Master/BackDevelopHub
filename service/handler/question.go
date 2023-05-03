package handler

import (
	query "BackDevelopHub/database/access"
	"fmt"
	"net/http"
)

//모든 질문 갯수 조회 [GET] (/question/totla)
func QuestionTotal(responseWriter http.ResponseWriter, request *http.Request) {
	jsonData := fmt.Sprintf("{total_question_count: %d}", query.GetRowsCount())
	sendJSONResponse(responseWriter, http.StatusOK, jsonData)
} 

// 질문 미리보기 조회 [GET] (/question/preview?question_idx={idx})
func QuestionPreview(responseWriter http.ResponseWriter, request *http.Request) {

}

// 질문 조회 [GET] (/question/detail?question_idx={idx})
func QuestionDetail(responseWriter http.ResponseWriter, request *http.Request) {
	
}