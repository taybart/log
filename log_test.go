package logger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/logrusorgru/aurora"
)

func TestTrace(t *testing.T) {
	toutput(t, LEVELTRACE, Gray("[TRACE]").String())
}
func TestVerbose(t *testing.T) {
	toutput(t, LEVELVERBOSE, Magenta("[VERBOSE]").String())
}
func TestDebug(t *testing.T) {
	toutput(t, LEVELDEBUG, Blue("[DEBUG]").String())
}
func TestInfo(t *testing.T) {
	toutput(t, LEVELINFO, Green("[INFO]").String())
}
func TestWarn(t *testing.T) {
	toutput(t, LEVELWARN, Brown("[WARN]").String())
}
func TestError(t *testing.T) {
	toutput(t, LEVELERROR, Red("[ERROR]").String())
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
