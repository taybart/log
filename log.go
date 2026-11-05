package log

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
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
		f := strings.Trim(fmt.Sprintf("%v", v), "[]")
		_, file, line, _ := runtime.Caller(2)
		f = fmt.Sprintf("%s\n%v: %v", f, file, line)
		l := getlabel(Red, "[FATAL]")
		log(createOutput(l, f, addnewline))
		os.Exit(1)
	}
}

// Fatalf : logging
func Fatalf(format string, v ...interface{}) {
	if level <= FATAL {
		f := fmt.Sprintf(format, v...)
		_, file, line, _ := runtime.Caller(2)
		f = fmt.Sprintf("%s\n%v: %v", f, file, line)
		l := getlabel(Red, "[FATAL]")
		log(createOutput(l, f, addnewline))
		os.Exit(1)
	}
}

// statusWriter wraps an http.ResponseWriter, recording the status code while
// passing everything straight through to the underlying writer.
type statusWriter struct {
	http.ResponseWriter
	status   int
	hijacked bool
}

func (w *statusWriter) WriteHeader(code int) {
	if w.status == 0 {
		w.status = code
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

// Flush implements http.Flusher for code that type-asserts directly.
func (w *statusWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		if w.status == 0 {
			w.status = http.StatusOK
		}
		f.Flush()
	}
}

// Hijack implements http.Hijacker so connection upgrades (websockets) work.
func (w *statusWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("log: underlying ResponseWriter does not implement http.Hijacker")
	}
	conn, rw, err := h.Hijack()
	if err == nil {
		w.hijacked = true
	}
	return conn, rw, err
}

// Unwrap lets http.ResponseController reach the real writer
// (Flush, Hijack, SetReadDeadline, SetWriteDeadline, ...).
func (w *statusWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }

// Middleware logs each request's status, remote address, duration, method and path.
// The response is passed through unbuffered, so streaming and upgrades work.
func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w}
		next(sw, r)

		statusCode := sw.status
		if statusCode == 0 {
			// nothing written, or the connection was hijacked (e.g. via
			// http.ResponseController, which bypasses our Hijack method)
			if sw.hijacked || r.Header.Get("Upgrade") != "" {
				statusCode = http.StatusSwitchingProtocols
			} else {
				statusCode = http.StatusOK
			}
		}
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
