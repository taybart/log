# Log

[![Build Status](https://travis-ci.org/taybartski/log.svg)](https://travis-ci.org/taybartski/log)

### Levels
Levels can be:
1)	LEVELTRACE
1)	LEVELVERBOSE
1)	LEVELDEBUG
1)	LEVELINFO
1)	LEVELWARN
1)	LEVELERROR

### Usage

Note: Error level does not panic on its own. This is ment to give the user process control.


```go
import "github.com/taybartski/log"


// All levels
func main() {
  myVar := "test"
  log.SetLevel(log.LEVELTRACE)
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

Why do some levels have "ln"? Because I said so.
