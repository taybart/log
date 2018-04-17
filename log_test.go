package logger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	toutput(t, LEVELTRACE, "\033[37m[TRACE]\033[0m")
}
func TestVerbose(t *testing.T) {
	toutput(t, LEVELVERBOSE, "\033[35m[VERBOSE]\033[0m")
}
func TestDebug(t *testing.T) {
	toutput(t, LEVELDEBUG, "\033[34m[DEBUG]\033[0m")
}
func TestInfo(t *testing.T) {
	toutput(t, LEVELINFO, "\033[32m[INFO]\033[0m")
}
func TestTest(t *testing.T) {
	toutput(t, LEVELTEST, "\033[32m[TEST]\033[0m")
}
func TestWarn(t *testing.T) {
	toutput(t, LEVELWARN, "\033[33m[WARN]\033[0m")
}
func TestError(t *testing.T) {
	toutput(t, LEVELERROR, "\033[31m[ERROR]\033[0m")
}

func toutput(t *testing.T, level Level, label string) {
	stdo := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	l := NewLogger(level)
	oc := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		if err != nil {
			t.Error(err)
		}
		oc <- buf.String()
	}()
	switch level {
	case LEVELTRACE:
		l.Trace("test")
	case LEVELVERBOSE:
		l.Verbose("test")
	case LEVELDEBUG:
		l.Debug("test")
	case LEVELINFO:
		l.Info("test")
	case LEVELTEST:
		l.Test("test")
	case LEVELWARN:
		l.Warn("test")
	case LEVELERROR:
		l.Error("test")
	}
	err := w.Close()
	if err != nil {
		t.Error(err)
	}
	os.Stdout = stdo
	output := <-oc
	str := time.Now().Format("2006-01-02 15:04:05") + " " + label + " test\n"
	if strings.Compare(str, output) != 0 {
		t.Error(str, output)
	}
}
