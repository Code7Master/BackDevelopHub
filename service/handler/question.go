package handler

import (
	db "BackDevelopHub/database/access"
	"fmt"
	"http/net"
)

//모든 질문 갯수 조회 [GET] (/question/totla)
func QuestionTotal(responseWriter http.ResponseWriter, request *http.Reqeust) {
	// DB 값 가져오고 [v]
	// 헤더에 200 + json 이다. -> 구현 해야 함
	// 바디에 정제된 값 넣기 

	db.GetCountRows()

	jsonData := fmt.Sprintf("{total_question_count: %s}", string(db.GetCountRows()))
	sendJSONResponse(&responseWriter, )

} 

// 질문 미리보기 조회 [GET] (/question/preview?question_idx={idx})
func QuestionPreview(responseWriter http.ReponseWriter, request *http.Request) {

}

// 질문 조회 [GET] (/question/detail?question_idx={idx})
func QuestionDetail(responseWriter http.ResponseWriter, request *http.Request) {
	
}