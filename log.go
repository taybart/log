package logger

import (
	"fmt"
	"time"

	. "github.com/logrusorgru/aurora"
)

type Level int

const (
	LEVELTRACE Level = iota
	LEVELVERBOSE
	LEVELDEBUG
	LEVELINFO
	LEVELTEST
	LEVELWARN
	LEVELERROR
)

type Logger struct {
	level   Level
	timeFmt string
}

func NewLogger(l Level) Logger {
	timeFmt := "2006-01-02 15:04:05"
	return Logger{l, timeFmt}
}

func (l Logger) Trace(format string, v ...interface{}) {
	if l.level <= LEVELTRACE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Gray("[TRACE]"), f)
	}
}

func (l Logger) Verbose(format string, v ...interface{}) {
	if l.level <= LEVELVERBOSE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Magenta("[VERBOSE]"), f)
	}
}

func (l Logger) Debug(format string, v ...interface{}) {
	if l.level <= LEVELDEBUG {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Blue("[DEBUG]"), f)
	}
}

func (l Logger) Info(format string, v ...interface{}) {
	if l.level <= LEVELINFO {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Green("[INFO]"), f)
	}
}

func (l Logger) Test(format string, v ...interface{}) {
	if l.level <= LEVELTEST {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Green("[TEST]"), f)
	}
}

func (l Logger) Warn(format string, v ...interface{}) {
	if l.level <= LEVELWARN {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Brown("[WARN]"), f)
	}
}

func (l Logger) Error(format string, v ...interface{}) {
	if l.level <= LEVELERROR {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), Red("[ERROR]"), f)
	}
}
