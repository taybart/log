package log

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

var level = LEVELINFO
var timeFmt = "2006-01-02 15:04:05"

func SetLevel(l Level) {
	level = l
}

func SetFmt(f string) {
	timeFmt = f
}

func Trace(format string, v ...interface{}) {
	if level <= LEVELTRACE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[37m[TRACE]\033[0m", f)
	}
}

func Verbose(format string, v ...interface{}) {
	if level <= LEVELVERBOSE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[35m[VERBOSE]\033[0m", f)
	}
}

func Debug(format string, v ...interface{}) {
	if level <= LEVELDEBUG {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[34m[DEBUG]\033[0m", f)
	}
}

func Infoln(v ...interface{}) {
	if level <= LEVELINFO {
		f := fmt.Sprintf("%v", v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[32m[INFO]\033[0m", f)
	}
}
func Info(format string, v ...interface{}) {
	if level <= LEVELINFO {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[32m[INFO]\033[0m", f)
	}
}

func Test(format string, v ...interface{}) {
	if level <= LEVELTEST {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[32m[TEST]\033[0m", f)
	}
}

func Warn(format string, v ...interface{}) {
	if level <= LEVELWARN {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[33m[WARN]\033[0m", f)
	}
}

func Error(format string, v ...interface{}) {
	if level <= LEVELERROR {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[31m[ERROR]\033[0m", f)
	}
}

func Errorln(v ...interface{}) {
	if level <= LEVELERROR {
		f := fmt.Sprintf("%v", v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[31m[ERROR]\033[0m", f)
	}
}
