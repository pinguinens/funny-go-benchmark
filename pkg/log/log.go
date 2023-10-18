package log

import "fmt"

type Logger struct{}

func (l *Logger) Infoln(v ...interface{}) {
	fmt.Println(v...)
}

func (l *Logger) Infolnf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
