package log

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"time"
)

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
	if level <= VERBOSE {
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
	if level <= WARN {
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
		panic(f)
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
			Reset)
		log(o)

		for k, v := range rec.Result().Header {
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
		return color + label + Reset
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
		fmt.Printf("%s[LOG ERROR] Issue with log package: %s%s\n", Red, err, Reset)
	}
}
