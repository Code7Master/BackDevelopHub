package log

import "fmt"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Record(str string) {
	// TODO: 파일 하나를 열어 현재 시각과 str을 저장해야 함
	fmt.Println(str)
}
