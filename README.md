# Log

```go
import "github.com/pmpbar/log"

var l Logger

// All levels
func main() {
  l = logger.NewLogger(logger.LEVELTRACE)
  l.Trace("test")
  l.Verbose("test")
  l.Debug("test")
  l.Info("test")
  l.Warn("test")
  l.Error("test")
}
```

### Levels
Levels can be:
*	LEVELTRACE
*	LEVELVERBOSE
*	LEVELDEBUG
*	LEVELINFO
*	LEVELWARN
*	LEVELERROR
