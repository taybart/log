package log

import (
	"fmt"
	"time"

	. "github.com/logrusorgru/aurora"
)

const (
	err = iota
	info
)

type Logger struct {
	level   int
	timeFmt string
}

func NewLogger() Logger {
	level := err
	timeFmt := "2006-01-02 15:04:05"
	return Logger{level, timeFmt}
}

func (l Logger) Info(format string, v ...interface{}) {
	if l.level <= info {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Green("[INFO]"), f)
	}
}

func (l Logger) Error(format string, v ...interface{}) {
	if l.level <= err {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Red("[ERROR]"), f)
	}
}
