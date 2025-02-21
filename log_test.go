package log

import (
	"bytes"
	"errors"
	"io"
	"os"

	// "strings"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	toutput(t, TRACE, "\033[37m[TRACE]\033[0m")
}
func TestVerbose(t *testing.T) {
	toutput(t, VERBOSE, "\033[35m[VERBOSE]\033[0m")
}
func TestDebug(t *testing.T) {
	toutput(t, DEBUG, "\033[34m[DEBUG]\033[0m")
}
func TestInfo(t *testing.T) {
	toutput(t, INFO, "\033[32m[INFO]\033[0m")
}
func TestTest(t *testing.T) {
	toutput(t, TEST, "\033[32m[TEST]\033[0m")
}
func TestWarn(t *testing.T) {
	toutput(t, WARN, "\033[33m[WARN]\033[0m")
}
func TestError(t *testing.T) {
	toutput(t, ERROR, "\033[31m[ERROR]\033[0m")
}
func TestFatal(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("logger did not die")
		}
	}()
	toutput(t, FATAL, "\033[31m[FATAL]\033[0m")
}

func TestFileOutput(t *testing.T) {
	testlog := "./test.txt"
	// cleanup
	defer func() {
		if err := os.Remove(testlog); err != nil {
			t.Errorf("could not remove test file: %s", err)
		}
	}()

	if err := os.Remove(testlog); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			t.Error(err)
		}
	}
	// prep logger
	SetOutput(testlog)
	SetPlain()
	SetLevel(INFO) // level must be set because other tests set it

	// test
	Infof("test")

	// check contents
	b, err := os.ReadFile(testlog)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "test" {
		t.Errorf("expected file to contain log contents, got: %s", b)
	}
}

func toutput(t *testing.T, level Level, label string) {
	stdo := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	SetLevel(level)
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
	case TRACE:
		Traceln("test")
	case VERBOSE:
		Verboseln("test")
	case DEBUG:
		Debugln("test")
	case INFO:
		Infoln("test")
	case TEST:
		Testln("test")
	case WARN:
		Warnln("test")
	case ERROR:
		Errorln("test")
	case FATAL:
		Fatalln("test")
	}
	err := w.Close()
	if err != nil {
		t.Error(err)
	}
	os.Stdout = stdo
	output := <-oc
	str := time.Now().Format("2006-01-02 15:04:05") + " " + label + " test\n"
	if bytes.Equal([]byte(str), []byte(output)) {
		t.Error(str, output)
	}
}
