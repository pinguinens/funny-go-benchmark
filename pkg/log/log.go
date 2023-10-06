package log

import "fmt"

type Logger struct{}

func (l *Logger) Info(v ...interface{}) {
	fmt.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}
