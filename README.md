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

  log.SetOutput("./logfile.log") // defaults to stdout

  log.UseColors(true) // defaults to true

  amount := 1
  thingy := "thingy"

  log.Trace(amount, thingy) // alias for "ln" version
  log.Traceln(amount, thingy)
  log.Tracef("test %s", thingy)

  log.Verbose(amount, thingy) // alias for "ln" version
  log.Verboseln(amount, thingy)
  log.Verbosef("test %s", thingy)

  log.Debug(amount) // alias for "ln" version
  log.Debugln(amount)
  log.Debugf("test %s", thingy)

  log.Info(amount, thingy) // alias for "ln" version
  log.Infoln(amount, thingy)
  log.Infof("test %s", thingy)

  log.Warn(amount, thingy) // alias for "ln" version
  log.Warnln(amount, thingy)
  log.Warnf("test %s", thingy)

  log.Error(amount, thingy) // alias for "ln" version
  log.Errorln(amount, thingy)
  log.Errorf("test %s", thingy)
}
```

## HTTP logging

```go
package main

import (
  "fmt"

  "github.com/taybart/log"
)

type server struct {
  router *http.ServeMux
}

func (s *server) routes() {
  // Add logging to a route
  s.router.HandleFunc("/ping", log.Middleware(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "pong")
  }))
}

func main() {
  srv := &http.Server{
    Handler:      s.router,
    Addr:         c.addr,
    WriteTimeout: 15 * time.Second,
    ReadTimeout:  15 * time.Second,
  }

  log.Infof("Serving at %s\n", c.addr)
  log.Fatal(srv.ListenAndServe())
}
```

## Templates

```go
package main

import (
	"bytes"
	"text/template"

	"github.com/taybart/log"
)

type Block struct {
	Text string
}

func main() {
	log.SetPlain()
	b := Block{Text: "hello world!"}
	tmpl := `{{green "~~~~~START~~~~~" }}
{{red .Text}}
{{ green "~~~~~~END~~~~~~" }}`

	t := template.Must(template.New("block").Funcs(log.TmplFuncs).Parse(tmpl))

	var buf bytes.Buffer
	t.Execute(&buf, b)
	log.Info(buf.String())
}
```
