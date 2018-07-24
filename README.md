# Log

[![Build Status](https://travis-ci.org/taybartski/log.svg)](https://travis-ci.org/taybartski/log)

```go
import "github.com/taybartski/log"


// All levels
func main() {
  log.SetLevel(log.LEVELTRACE)
  log.SetFmt("2006-01-02 15:04:05") // Default time format
  log.Trace("test")
  log.Verbose("test")
  log.Debug("test")
  log.Info("test")
  log.Warn("test")
  log.Error("test")
}
```

### Levels
Levels can be:
1)	LEVELTRACE
1)	LEVELVERBOSE
1)	LEVELDEBUG
1)	LEVELINFO
1)	LEVELWARN
1)	LEVELERROR
