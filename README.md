# Log

[![Build Status](https://travis-ci.org/taybart/log.svg)](https://travis-ci.org/taybart/log)

Simple level logging

### Levels
Levels can be:
1)	TRACE
1)	VERBOSE
1)	DEBUG
1)	INFO
1)	WARN
1)	ERROR
1)	FATAL

### Usage

```go
package main

import "github.com/taybart/log"


// All levels
func main() {
  log.SetLevel(log.TRACE)
  log.SetTimeFmt("2006-01-02 15:04:05") // Default time format

  log.SetOutput("./logifle.log") // defaults to stdout

  log.UseColors(true) // defaults to true

  amount := 1
  thingy := "thingy"

  log.Traceln(amount, thingy)
  log.Trace("test %s", thingy)

  log.Verboseln(amount, thingy)
  log.Verbose("test %s", thingy)

  log.Debugln(amount)
  log.Debug("test %s", thingy)

  log.Infoln(amount, thingy)
  log.Info("test %s", thingy)

  log.Warnln(amount, thingy)
  log.Warn("test %s", thingy)

  log.Errorln(amount, thingy)
  log.Error("test %s", thingy)
}
```
