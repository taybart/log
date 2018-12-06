# Log

[![Build Status](https://travis-ci.org/taybartski/log.svg)](https://travis-ci.org/taybartski/log)

Level logging to stdout.

### Levels
Levels can be:
1)	TRACE
1)	VERBOSE
1)	DEBUG
1)	INFO
1)	WARN
1)	ERROR

### Usage

```go
import "github.com/taybartski/log"


// All levels
func main() {
  myVar := "test"
  log.SetLevel(log.TRACE)
  log.SetFmt("2006-01-02 15:04:05") // Default time format
  log.Trace("test %s", myVar)
  log.Verbose("test %s", myVar)
  log.Debugln(myVar)
  log.Debug("test %s", myVar)
  log.Infoln(myVar)
  log.Info("test %s", myVar)
  log.Warn("test %s", myVar)
  log.Errorln(myVar)
  log.Error("test %s", myVar)
}
```
