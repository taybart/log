package log

import (
	"io"
	"os"
	"strings"
)

// Level type for level logging
type Level int

const (
	// TRACE lowest = most verbose
	TRACE Level = iota + 1
	// DEBUG level
	DEBUG
	// VERBOSE level
	VERBOSE
	// TEST special level for testing
	TEST
	// INFO level
	INFO
	// WARN level
	WARN
	// ERROR level
	ERROR
	// FATAL level
	FATAL
)

const (
	addnewline = true
)

var (
	// level : sets the log level, anything under will not be sent to "Output"
	level = INFO
	// timeFmt : go format for the time in the logs
	timeFmt = "2006-01-02 15:04:05"
	// plain : don't add level and time to logs (true is essentially fmt.Print() with levels)
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

func init() {
	if os.Getenv("LOG_NO_COLOR") != "" {
		useColors = false
	}
	lvl := os.Getenv("LOG_LEVEL")
	switch strings.ToUpper(lvl) {
	case "TRACE":
		level = TRACE
	case "DEBUG":
		level = DEBUG
	case "VERBOSE":
		level = VERBOSE
	case "TEST":
		level = TEST
	case "INFO":
		level = INFO
	case "WARN":
		level = WARN
	case "ERROR":
		level = ERROR
	case "FATAL":
		level = FATAL
	}
	if os.Getenv("LOG_PLAIN") != "" {
		plain = true
	}
}

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

// SetFancy : output, will print time and level
func SetFancy() {
	plain = false
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
	fd, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	Output = fd
	return nil
}
