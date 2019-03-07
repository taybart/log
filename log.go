package log

import (
	"fmt"
	"os"
	"runtime"
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

// Trace print var with no format
func Trace(v ...interface{}) {
	Traceln(v)
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

// Tracef logging
func Tracef(format string, v ...interface{}) {
	if level <= TRACE {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Verbose print var with no format
func Verbose(v ...interface{}) {
	Verboseln(v)
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

// Verbosef logging
func Verbosef(format string, v ...interface{}) {
	if level <= VERBOSE {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Purple, "[VERBOSE]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Debug print var with no format
func Debug(v ...interface{}) {
	Debugln(v)
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

// Debugf logging
func Debugf(format string, v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Blue, "[DEBUG]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Info info with no format
func Info(v ...interface{}) {
	Infoln(v)
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

// Infof logging
func Infof(format string, v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Green, "[INFO]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Test test with no format
func Test(v ...interface{}) {
	Testln(v)
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

// Testf logging
func Testf(format string, v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Green, "[TEST]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Warn error var with no format
func Warn(v ...interface{}) {
	Warnln(v)
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

// Warnf logging
func Warnf(format string, v ...interface{}) {
	if level <= WARN {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Yellow, "[WARN]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Error error var with no format
func Error(v ...interface{}) {
	Errorln(v)
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

// Errorf logging
func Errorf(format string, v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Red, "[ERROR]")
		err := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(err)
	}
}

// Fatal error var with no format
func Fatal(v ...interface{}) {
	Fatalln(v)
}

// Fatalln error var with no format
func Fatalln(v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Red, "[FATAL]")
		_, file, line, _ := runtime.Caller(1)
		err := fmt.Sprintf("%v %v %v\n%v %v\n", time.Now().Format(timeFmt), l, f, file, line)
		log(err)
		os.Exit(1)
	}
}

// Fatalf logging
func Fatalf(format string, v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Red, "[FATAL]")

		_, file, line, _ := runtime.Caller(1)
		err := fmt.Sprintf("%v %v %v\n%v %v", time.Now().Format(timeFmt), l, f, file, line)
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
