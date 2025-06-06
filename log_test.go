package log

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	toutput(t, TRACE, "\033[0;37m[TRACE]\033[0;0m")
}
func TestVerbose(t *testing.T) {
	toutput(t, VERBOSE, "\033[0;35m[VERBOSE]\033[0;0m")
}
func TestDebug(t *testing.T) {
	toutput(t, DEBUG, "\033[0;34m[DEBUG]\033[0;0m")
}
func TestInfo(t *testing.T) {
	toutput(t, INFO, "\033[0;32m[INFO]\033[0;0m")
}
func TestTest(t *testing.T) {
	toutput(t, TEST, "\033[0;32m[TEST]\033[0;0m")
}
func TestWarn(t *testing.T) {
	toutput(t, WARN, "\033[0;33m[WARN]\033[0;0m")
}
func TestError(t *testing.T) {
	toutput(t, ERROR, "\033[0;31m[ERROR]\033[0;0m")
}
func TestFatal(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("logger did not die")
		}
	}()
	toutput(t, FATAL, "\033[0;31m[FATAL]\033[0;0m")
}
func TestNone(t *testing.T) {
	toutput(t, NONE, "")
}

func toutput(t *testing.T, level Level, label string) {
	stdo := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	oc := make(chan string)

	SetLevel(level)
	SetOutputWriter(w)
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
	case NONE:
		Fatalln("test")
	}
	err := w.Close()
	if err != nil {
		t.Error(err)
	}
	output := <-oc
	os.Stdout = stdo
	str := fmt.Sprintf("%v %v test\n", time.Now().Format("2006-01-02 15:04:05"), label)
	if level == NONE {
		if output != "" {
			t.Errorf("expected no output, got: %v", output)
		}
	} else if strings.Compare(str, output) != 0 {
		t.Errorf("\n%v != \n%v\n", []byte(str), []byte(strings.Trim(output, "\n")))
	}
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
