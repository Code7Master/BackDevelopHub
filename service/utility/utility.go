package utility  

import (
	"BackDevelopHub/dhjson"
	"net/http"
)

func SendJSONResponse(responseWriter http.ResponseWriter, statusCode int, json []byte) {
	responseWriter.Header().Set("Content-Type", "application/json");
	responseWriter.WriteHeader(statusCode)
	responseWriter.Write(json)
}

func Atoi(str string) int {
	var resVal int
	strLen := len(str)

	for i := 0; i < strLen; i++ {
		resVal += int(str[i]) - '0'
		if i+1 != strLen {
			resVal *= 10
		}
	}
	return resVal
}

func GetInvalidQuestionPreviewIndexStruct() dhjson.QuestionPreview{
	return dhjson.QuestionPreview {
		Writer: "No information was found corresponding to the passed query",
		Title: "No information was found corresponding to the passed query",
		View_count: 0,
		Answer_count: 0,
		Vote_count: 0,
	}
}

func GetInvalidQuestionDetailByIndex() dhjson.QuestionDetail {
	return dhjson.QuestionDetail {
		Writer: "No information was found corresponding to the passed query",
		Title: "No information was found corresponding to the passed query",
		Content: "No information was found corresponding to the passed query",
		View_count: 0,
		Answer_count: 0,
		Vote_count: 0,
	}
}