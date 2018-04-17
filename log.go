package logger

import (
	"fmt"
	"time"
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
		fmt.Println(time.Now().Format(l.timeFmt), "\033[37m[TRACE]\033[0m", f)
	}
}

func (l Logger) Verbose(format string, v ...interface{}) {
	if l.level <= LEVELVERBOSE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), "\033[35m[VERBOSE]\033[0m", f)
	}
}

func (l Logger) Debug(format string, v ...interface{}) {
	if l.level <= LEVELDEBUG {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), "\033[34m[DEBUG]\033[0m", f)
	}
}

func (l Logger) Info(format string, v ...interface{}) {
	if l.level <= LEVELINFO {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), "\033[32m[INFO]\033[0m", f)
	}
}

func (l Logger) Test(format string, v ...interface{}) {
	if l.level <= LEVELTEST {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), "\033[32m[TEST]\033[0m", f)
	}
}

func (l Logger) Warn(format string, v ...interface{}) {
	if l.level <= LEVELWARN {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), "\033[33m[WARN]\033[0m", f)
	}
}

func (l Logger) Error(format string, v ...interface{}) {
	if l.level <= LEVELERROR {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(l.timeFmt), "\033[31m[ERROR]\033[0m", f)
	}
}
