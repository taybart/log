package log

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"time"
)

// Level type for level logging
type Level int

const (
	// TRACE lowest = most verbose
	TRACE Level = iota + 1
	// VERBOSE level
	VERBOSE
	// DEBUG level
	DEBUG
	// INFO level
	INFO
	// HTTP level
	HTTP
	// TEST special level for testing
	TEST
	// WARN level
	WARN
	// ERROR level
	ERROR
	// FATAL level
	FATAL
)

// Colors
const (
	// Color escape
	ce = "\033[0;"

	// Normal Colors
	Gray   = ce + "37m"
	Purple = ce + "35m"
	Blue   = ce + "34m"
	Yellow = ce + "33m"
	Green  = ce + "32m"
	Red    = ce + "31m"

	// Bold Colors
	ceBold     = "\033[1;"
	BoldGray   = ceBold + "37m"
	BoldPurple = ceBold + "35m"
	BoldBlue   = ceBold + "34m"
	BoldYellow = ceBold + "33m"
	BoldGreen  = ceBold + "32m"
	BoldRed    = ceBold + "31m"

	// Italic Colors
	ceItalic     = "\033[3;"
	ItalicGray   = ceItalic + "37m"
	ItalicPurple = ceItalic + "35m"
	ItalicBlue   = ceItalic + "34m"
	ItalicYellow = ceItalic + "33m"
	ItalicGreen  = ceItalic + "32m"
	ItalicRed    = ceItalic + "31m"

	// Underlined Colors
	ceUnderlined     = "\033[4;"
	UnderlinedGray   = ceUnderlined + "37m"
	UnderlinedPurple = ceUnderlined + "35m"
	UnderlinedBlue   = ceUnderlined + "34m"
	UnderlinedYellow = ceUnderlined + "33m"
	UnderlinedGreen  = ceUnderlined + "32m"
	UnderlinedRed    = ceUnderlined + "31m"

	// Blinking Colors
	ceBlinking     = "\033[5;"
	BlinkingGray   = ceBlinking + "37m"
	BlinkingPurple = ceBlinking + "35m"
	BlinkingBlue   = ceBlinking + "34m"
	BlinkingYellow = ceBlinking + "33m"
	BlinkingGreen  = ceBlinking + "32m"
	BlinkingRed    = ceBlinking + "31m"

	// Return to default
	Rtd = ce + "0m"
)

const (
	addnewline = true
)

var (
	// level : sets the log level, anything under will not be sent to "Output"
	level = INFO
	// timeFmt : go format for the time in the logs
	timeFmt = "2006-01-02 15:04:05"
	// plain : don't add level and time to logs (false is essentially fmt.Print()
	plain = false
	// noTime : don't add time to output, is not used if plain is set
	noTime = false
	// timeOnly : don't show level in logs, is not used if plain is set
	timeOnly = false

	// Output writer for log
	Output io.Writer = os.Stdout

	// UseColors allow console coloring
	useColors = true
)

// UseColors : used to set colors
func UseColors(use bool) {
	useColors = use
}

// SetLevel : used to set logging level
func SetLevel(l Level) {
	level = l
}

// SetTimeFmt : used to adjust time format for logs
func SetTimeFmt(f string) {
	timeFmt = f
}

// SetPlain : output, will not print time or level
func SetPlain() {
	plain = true
}

// SetTimeOnly : output, will not print time or level
func SetNoTime() {
	noTime = true
}

// SetTimeOnly : output, will not print time or level
func SetTimeOnly() {
	timeOnly = true
}

// SetOutputWriter : set log io.Writer, if logs should be streamed, the io.Writer can be passed here
func SetOutputWriter(w io.Writer) {
	Output = w
}

// SetOutput : set log output to a specific file, default is stdout
func SetOutput(filename string) error {
	logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	Output = logfile
	return nil
}

// Print :  print var with no format
func Print(v ...interface{}) {
	Traceln(v)
}

// Println : print var with no format
func Println(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		if plain {
			log(fmt.Sprintf("%s\n", f))
			return
		}
		l := getlabel(Gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Printf : logging
func Printf(format string, v ...interface{}) {
	f := fmt.Sprintf(format, v...)
	if plain {
		log(f)
		return
	}
	o := fmt.Sprintf("%v %v", time.Now().Format(timeFmt), f)
	log(o)
}

// Trace : print var with no format
func Trace(v ...interface{}) {
	Traceln(v)
}

// Traceln : print var with no format
func Traceln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		if plain {
			log(fmt.Sprintf("%s\n", f))
			return
		}
		l := getlabel(Gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v\n", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Tracef : logging
func Tracef(format string, v ...interface{}) {
	if level <= TRACE {
		f := fmt.Sprintf(format, v...)
		if plain {
			log(f)
			return
		}
		l := getlabel(Gray, "[TRACE]")
		o := fmt.Sprintf("%v %v %v", time.Now().Format(timeFmt), l, f)
		log(o)
	}
}

// Verbose : print var with no format
func Verbose(v ...interface{}) {
	Verboseln(v)
}

// Verboseln : print var with no format
func Verboseln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		if plain {
			log(fmt.Sprintf("%s\n", f))
			return
		}
		l := getlabel(Purple, "[VERBOSE]")
		log(createOutput(l, f, addnewline))
	}
}

// Verbosef : logging
func Verbosef(format string, v ...interface{}) {
	if level <= VERBOSE {
		f := fmt.Sprintf(format, v...)
		if plain {
			log(f)
			return
		}
		l := getlabel(Purple, "[VERBOSE]")
		log(createOutput(l, f, !addnewline))
	}
}

// Debug : print var with no format
func Debug(v ...interface{}) {
	Debugln(v)
}

// Debugln : print var with no format
func Debugln(v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Blue, "[DEBUG]")
		log(createOutput(l, f, addnewline))
	}
}

// Debugf : logging
func Debugf(format string, v ...interface{}) {
	if level <= DEBUG {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Blue, "[DEBUG]")
		log(createOutput(l, f, !addnewline))
	}
}

// Info : info with no format
func Info(v ...interface{}) {
	Infoln(v)
}

// Infoln : info with no format
func Infoln(v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Green, "[INFO]")
		log(createOutput(l, f, addnewline))
	}
}

// Infof : logging
func Infof(format string, v ...interface{}) {
	if level <= INFO {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Green, "[INFO]")
		log(createOutput(l, f, !addnewline))
	}
}

// Test : test with no format
func Test(v ...interface{}) {
	Testln(v)
}

// Testln : test with no format
func Testln(v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Green, "[TEST]")
		log(createOutput(l, f, addnewline))
	}
}

// Testf : logging
func Testf(format string, v ...interface{}) {
	if level <= TEST {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Green, "[TEST]")
		log(createOutput(l, f, !addnewline))
	}
}

// Warn : error var with no format
func Warn(v ...interface{}) {
	Warnln(v)
}

// Warnln : error var with no format
func Warnln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Yellow, "[WARN]")
		log(createOutput(l, f, addnewline))
	}
}

// Warnf : logging
func Warnf(format string, v ...interface{}) {
	if level <= WARN {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Yellow, "[WARN]")
		log(createOutput(l, f, !addnewline))
	}
}

// Error : error var with no format
func Error(v ...interface{}) {
	Errorln(v)
}

// Errorln : error var with no format
func Errorln(v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Red, "[ERROR]")
		log(createOutput(l, f, addnewline))
	}
}

// Errorf : logging
func Errorf(format string, v ...interface{}) {
	if level <= ERROR {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Red, "[ERROR]")
		log(createOutput(l, f, !addnewline))
	}
}

// Fatal : error var with no format
func Fatal(v ...interface{}) {
	Fatalln(v)
}

// Fatalln : error var with no format
func Fatalln(v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf("%v", v)
		f = strings.Trim(f, "[]")
		l := getlabel(Red, "[FATAL]")
		log(createOutput(l, f, addnewline))
		_, file, line, _ := runtime.Caller(2)
		f = fmt.Sprintf("%v: %v\n", file, line)
		log(createOutput(l, f, addnewline))
		os.Exit(1)
	}
}

// Fatalf : logging
func Fatalf(format string, v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf(format, v...)
		l := getlabel(Red, "[FATAL]")
		log(createOutput(l, f, addnewline))
		_, file, line, _ := runtime.Caller(2)
		f = fmt.Sprintf("%v: %v\n", file, line)
		log(createOutput(l, f, addnewline))
		os.Exit(1)
	}
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if level > HTTP {
			next(w, r)
			return
		}

		rec := httptest.NewRecorder()
		start := time.Now()
		next(rec, r)
		statusCode := rec.Result().StatusCode
		codeLabel := fmt.Sprintf("[%d]", statusCode)
		l := getlabel(Blue, codeLabel)
		if statusCode >= 400 {
			l = getlabel(Red, codeLabel)
		}
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
		o := fmt.Sprintln(time.Now().Format(timeFmt), l,
			ip,
			time.Since(start),
			r.Method,
			r.URL.Path,
			Rtd)
		log(o)

		for k, v := range rec.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		rec.Body.WriteTo(w)
	}
}

func getlabel(color, label string) string {
	if timeOnly {
		return ""
	}
	if useColors {
		return color + label + Rtd
	}
	return label
}

func createOutput(label, rest string, addnewline bool) string {
	nl := ""
	if addnewline {
		nl = "\n"
	}
	if plain {
		return fmt.Sprintf("%s%s", rest, nl)
	}
	if noTime {
		return fmt.Sprintf("%v %v%s", label, rest, nl)
	}
	return fmt.Sprintf("%v %v %v%s", time.Now().Format(timeFmt), label, rest, nl)
}

func log(s string) {
	_, err := Output.Write([]byte(s))
	if err != nil {
		fmt.Println("[LOG ERROR] Issue with log package", err)
	}
}
