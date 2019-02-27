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
	// FATAL level
	FATAL
)

/* // Colors
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
) */

// Colors
const (
	// Color escape
	ce     = "\033["
	Gray   = ce + "37m"
	Purple = ce + "35m"
	Blue   = ce + "34m"
	Yellow = ce + "33m"
	Green  = ce + "32m"
	Red    = ce + "31m"
	// Return to default
	Rtd = ce + "0m"
)

var level = INFO
var timeFmt = "2006-01-02 15:04:05"
var output = os.Stdout

// UseColors allow console coloring
var useColors = true

// UseColors used to set colors
func UseColors(use bool) {
}

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
		l := getlabel(Gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Trace logging
func Trace(format string, v ...interface{}) {
	if level <= TRACE {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Verboseln print var with no format
func Verboseln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Purple, "[VERBOSE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Verbose logging
func Verbose(format string, v ...interface{}) {
	if level <= VERBOSE {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Purple, "[VERBOSE]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Debugln print var with no format
func Debugln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Blue, "[DEBUG]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Debug logging
func Debug(format string, v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Blue, "[DEBUG]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Infoln info with no format
func Infoln(v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Green, "[INFO]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Info logging
func Info(format string, v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Green, "[INFO]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Testln test with no format
func Testln(v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Green, "[TEST]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Test logging
func Test(format string, v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Green, "[TEST]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Warnln error var with no format
func Warnln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Yellow, "[WARN]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Warn logging
func Warn(format string, v ...interface{}) {
	if level <= WARN {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Yellow, "[WARN]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Errorln error var with no format
func Errorln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Red, "[ERROR]")
		err := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(err)
	}
}

// Error logging
func Error(format string, v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Red, "[ERROR]")
		err := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(err)
	}
}

// Fatalln error var with no format
func Fatalln(v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Red, "[FATAL]")
		err := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(err)
		os.Exit(1)
	}
}

// Fatal logging
func Fatal(format string, v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Red, "[FATAL]")
		err := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(err)
		os.Exit(1)

	}
}

func getlabel(color, label string) string {
	if useColors {
		return color + label + Rtd
	}
	return label
}
func log(s string) {
	output.Write([]byte(s))
}
