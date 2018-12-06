package log

import (
	"fmt"
	"os"
	"time"
)

// Level type for level logging
type Level int

const (
	// TRACE lowest = most verbose
	TRACE Level = iota
	// VERBOSE level
	VERBOSE
	// DEBUG level
	DEBUG
	// INFO level
	INFO
	// TEST special level for testing
	TEST
	// WARN level
	WARN
	// ERROR level
	ERROR
)

var level = INFO
var timeFmt = "2006-01-02 15:04:05"
var output = os.Stdout

// PanicOnErrors panic with error instead of logging
var PanicOnErrors = false

// SetLevel used to set logging level
func SetLevel(l Level) {
	level = l
}

// SetFmt used to adjust time format for logs
func SetFmt(f string) {
	timeFmt = f
}

// SetOutput used to set log output
func SetOutput(f *os.File) {
	output = f
}

// Trace logging
func Trace(format string, v ...interface{}) {
	if level <= TRACE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[37m[TRACE]\033[0m", f)
	}
}

// Verbose logging
func Verbose(format string, v ...interface{}) {
	if level <= VERBOSE {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[35m[VERBOSE]\033[0m", f)
	}
}

// Debugln print var with no format
func Debugln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[34m[DEBUG]\033[0m", f)
	}
}

// Debug logging
func Debug(format string, v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[34m[DEBUG]\033[0m", f)
	}
}

// Infoln info with no format
func Infoln(v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf("%v", v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[32m[INFO]\033[0m", f)
	}
}

// Info logging
func Info(format string, v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[32m[INFO]\033[0m", f)
	}
}

// Test logging
func Test(format string, v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[32m[TEST]\033[0m", f)
	}
}

// Warn logging
func Warn(format string, v ...interface{}) {
	if level <= WARN {
		f := fmt.Sprintf(format, v...)
		fmt.Println(time.Now().Format(timeFmt), "\033[33m[WARN]\033[0m", f)
	}
}

// Errorln error var with no format
func Errorln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v...)
		err := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), "\033[31m[ERROR]\033[0m", f)
		if PanicOnErrors {
			panic(err)
		}
		fmt.Println(err)
	}
}

// Error logging
func Error(format string, v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf(format, v...)
		err := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), "\033[31m[ERROR]\033[0m", f)
		if PanicOnErrors {
			panic(err)
		}
		fmt.Println(err)
	}
}
