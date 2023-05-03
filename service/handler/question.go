package handler

import (
	"fmt"
	"BackDevelopHub/service/utility"
	"net/http"
	"encoding/json"
	query "BackDevelopHub/database/access"
)

//모든 질문 갯수 조회 [GET] (/question/total)
func QuestionTotal(responseWriter http.ResponseWriter, request *http.Request) {
	var jsonData []byte	

	jsonData = []byte( fmt.Sprintf("{total_question_count: %d}", query.GetRowsCount()) )
	utility.SendJSONResponse(responseWriter, http.StatusOK, jsonData)
} 

// 질문 미리보기 조회 [GET] (/question/preview?question_idx={idx})
func QuestionPreview(responseWriter http.ResponseWriter, request *http.Request) {
	var jsonData []byte
	var question_idx int
	var statusCode int
	var err error

	question_idx = utility.Atoi(request.URL.Query().Get("question_idx"))
	
	if question_idx > query.GetRowsCount() {
		jsonData, err = json.Marshal(utility.GetInvalidQuestionPreviewIndexStruct())
		statusCode = http.StatusNotFound;
	} else {
		jsonData, err = json.Marshal(query.GetQuestionPreviewByIndex(question_idx)) 
		statusCode = http.StatusOK
	}

	if err != nil {
		fmt.Println("QuestionPreview error: ", err)
	}

	utility.SendJSONResponse(responseWriter, statusCode, jsonData)
}

// 질문 조회 [GET] (/question/detail?question_idx={idx})
func QuestionDetail(responseWriter http.ResponseWriter, request *http.Request) {
	var jsonData []byte
	var question_idx int
	var statusCode int
	var err error

	question_idx = utility.Atoi(request.URL.Query().Get("question_idx"))

	if question_idx > query.GetRowsCount() {
		jsonData, err = json.Marshal(utility.GetInvalidQuestionDetailByIndex())
		statusCode = http.StatusNotFound;
	} else {
		jsonData, err = json.Marshal(query.GetQuestionDetailByIndex(question_idx))
		statusCode = http.StatusOK
	}

	if err != nil {
		fmt.Println("QuestionDetail error: ", err)
	}

	utility.SendJSONResponse(responseWriter, statusCode, jsonData)
}


