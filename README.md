# Log

[![Build Status](https://travis-ci.org/pmpbar/log.svg)](https://travis-ci.org/pmpbar/log)

```go
import "github.com/pmpbar/log"


// All levels
func main() {
  l = log.SetLevel(log.LEVELTRACE)
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
