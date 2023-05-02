package handler

import (
	"http/net"
)

//모든 질문 갯수 조회 [GET] (/question/totla)
func QuestionTotal(reponseWriter http.ResponseWriter, request *http.Reqeust) {

} 

// 질문 미리보기 조회 [GET] (/question/preview?question_idx={idx})
func QuestionPreview(responseWriter http.ReponseWriter, request *http.Request) {

}

// 질문 조회 [GET] (/question/detail?question_idx={idx})
func QuestionDetail(responseWriter http.ResponseWriter, request *http.Request) {
	
}