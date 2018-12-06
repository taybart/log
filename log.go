package log

import (
	"fmt"
	"os"
	"strings"
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

// Colors
const (
	// Color escape
	ce     = "\033["
	gray   = ce + "37m"
	purple = ce + "35m"
	blue   = ce + "34m"
	yellow = ce + "33m"
	green  = ce + "32m"
	red    = ce + "31m"
	// Return to default
	rtd = ce + "0m"
)

var level = INFO
var timeFmt = "2006-01-02 15:04:05"
var output = os.Stdout

// PanicOnErrors panic with error instead of logging
var PanicOnErrors = false

// UseColors allow console coloring
var UseColors = true

// SetLevel used to set logging level
func SetLevel(l Level) {
	level = l
}

// SetTimeFmt used to adjust time format for logs
func SetTimeFmt(f string) {
	timeFmt = f
}

// SetOutput used to set log output
func SetOutput(filename string) error {

	logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	output = logfile
	return nil
}

// Traceln print var with no format
func Traceln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Trace logging
func Trace(format string, v ...interface{}) {
	if level <= TRACE {
		f := fmt.Sprintf(format, v...)
		l := getlabel(gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Verboseln print var with no format
func Verboseln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(purple, "[VERBOSE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Verbose logging
func Verbose(format string, v ...interface{}) {
	if level <= VERBOSE {
		f := fmt.Sprintf(format, v...)
		l := getlabel(purple, "[VERBOSE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Debugln print var with no format
func Debugln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(blue, "[DEBUG]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Debug logging
func Debug(format string, v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf(format, v...)
		l := getlabel(blue, "[DEBUG]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Infoln info with no format
func Infoln(v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(green, "[INFO]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Info logging
func Info(format string, v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf(format, v...)
		l := getlabel(green, "[INFO]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Test logging
func Test(format string, v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf(format, v...)
		l := getlabel(green, "[TEST]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Warnln error var with no format
func Warnln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(yellow, "[WARN]")
		err := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		if PanicOnErrors {
			panic(err)
		}
		log(err)
	}
}

// Warn logging
func Warn(format string, v ...interface{}) {
	if level <= WARN {
		f := fmt.Sprintf(format, v...)
		l := getlabel(yellow, "[WARN]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Errorln error var with no format
func Errorln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(red, "[ERROR]")
		err := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		if PanicOnErrors {
			panic(err)
		}
		log(err)
	}
}

// Error logging
func Error(format string, v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf(format, v...)
		l := getlabel(red, "[ERROR]")
		err := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		if PanicOnErrors {
			panic(err)
		}
		log(err)
	}
}

func getlabel(color, label string) string {
	if UseColors {
		return color + label + rtd
	}
	return label
}
func log(s string) {
	output.Write([]byte(s))
}
