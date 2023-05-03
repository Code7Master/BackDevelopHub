package dhjson

// question관련 JSON
type QuestionPreview struct {
	Writer 			string `json:"writer"`
    Title 		 	string `json:"title"` 
    View_count 		int    `json:"view_count`      
    Answer_count 	int    `json:"answer_count`
    Vote_count 		int    `json:"vote_count"`
}

type QuestionDetail struct {
	Writer 			string `json:"writer"`
    Title 		 	string `json:"title"` 
	Content			string `json:"content"`
    View_count 		int    `json:"view_count"`      
    Answer_count 	int    `json:"answer_count"`
    Vote_count 		int    `json:"vote_count"`
}