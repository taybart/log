package log

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// The client must receive a flushed chunk before the handler returns.
func TestMiddlewareStreams(t *testing.T) {
	release := make(chan struct{})
	srv := httptest.NewServer(Middleware(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "first\n")
		w.(http.Flusher).Flush()
		<-release
		io.WriteString(w, "second\n")
	}))
	defer srv.Close()
	defer close(release)

	// read in a goroutine: with a buffering middleware even the headers
	// never arrive, so http.Get itself would block until the handler returns
	got := make(chan string, 1)
	go func() {
		res, err := http.Get(srv.URL)
		if err != nil {
			got <- "error: " + err.Error()
			return
		}
		defer res.Body.Close()
		line, _ := bufio.NewReader(res.Body).ReadString('\n')
		got <- line
	}()
	select {
	case line := <-got:
		if line != "first\n" {
			t.Fatalf("got %q, want %q", line, "first\n")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("first chunk not received before handler returned; response is buffered")
	}
}

// A connection upgrade (as websockets do) must be able to hijack through Middleware.
func TestMiddlewareHijack(t *testing.T) {
	srv := httptest.NewServer(Middleware(func(w http.ResponseWriter, r *http.Request) {
		h, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "no hijacker", http.StatusInternalServerError)
			return
		}
		conn, rw, err := h.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		rw.WriteString("HTTP/1.1 101 Switching Protocols\r\nUpgrade: test\r\nConnection: Upgrade\r\n\r\nhello\n")
		rw.Flush()
	}))
	defer srv.Close()

	conn, err := net.Dial("tcp", strings.TrimPrefix(srv.URL, "http://"))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(2 * time.Second))
	io.WriteString(conn, "GET / HTTP/1.1\r\nHost: x\r\nUpgrade: test\r\nConnection: Upgrade\r\n\r\n")

	br := bufio.NewReader(conn)
	res, err := http.ReadResponse(br, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusSwitchingProtocols {
		t.Fatalf("status %d, want 101", res.StatusCode)
	}
	line, err := br.ReadString('\n')
	if err != nil || line != "hello\n" {
		t.Fatalf("got %q, %v; want %q", line, err, "hello\n")
	}
}

func TestMiddlewareRecordsStatus(t *testing.T) {
	cases := []struct {
		name string
		h    http.HandlerFunc
		want int
	}{
		{"explicit", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusTeapot) }, http.StatusTeapot},
		{"implicit write", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "x") }, http.StatusOK},
		{"first wins", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.WriteHeader(http.StatusOK)
		}, http.StatusNotFound},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			sw := &statusWriter{ResponseWriter: rec}
			c.h(sw, httptest.NewRequest("GET", "/", nil))
			if sw.status != c.want {
				t.Fatalf("recorded %d, want %d", sw.status, c.want)
			}
			if rec.Code != c.want {
				t.Fatalf("passed through %d, want %d", rec.Code, c.want)
			}
		})
	}
}
