package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"path"
	"strings"
	"time"
)

const defaultMaxBodyLen = 1024 * 1024

// request is always scoped to a single http request handled by the server
type request struct {
	file, path string

	ctx context.Context
	w   http.ResponseWriter
	r   *http.Request

	body []byte

	start       time.Time
	rid         uint64 // random request id
	read, wrote int
	maxBodyLen  int64
	ip, port    string
	err, logerr error
}

// newRequest initializes request scoped structures, context and counters, returning
// a function that can be deferred to log request details, newrelic, etc
func newRequest(w http.ResponseWriter, rq *http.Request) request {
	r := request{
		path:  rq.URL.Path,
		ctx:   rq.Context(),
		r:     rq,
		w:     w,
		start: time.Now(),
		rid:   rand.Uint64(),
	}
	r.rid |= 1 << 63 // sacrifice one bit of entropy so they always have the same # digits
	r.ip = r.r.Header.Get("X-Forwarded-For")
	r.port = r.r.Header.Get("X-Forwarded-Port")
	if r.ip == "" {
		r.ip, r.port, _ = net.SplitHostPort(r.r.RemoteAddr)
	}
	r.log(
		"ip", r.ip,
		"port", r.port,
		"raddr", r.r.RemoteAddr,
		"method", r.r.Method,
		"path", r.r.URL.Path,
		"ref", r.r.Referer(),
		"ua", r.r.UserAgent(),
	)
	return r
}

func (r *request) finalize() {
	if r.logerr == nil {
		r.logerr = r.err
	}
	r.log(
		"rx", r.read,
		"tx", r.wrote,
		"err", r.logerr,
	)
}

func (s *request) ok() bool {
	return s.err == nil
}

// Body reads the request body at most once and
// returns it.
func (s *request) Body() []byte {
	if !s.ok() {
		return nil
	}
	if s.body != nil {
		return s.body
	}
	s.body, s.err = ioutil.ReadAll(io.LimitReader(s.r.Body, defaultMaxBodyLen))
	s.read = len(s.body)
	return s.body
}

func (s *request) writeerror(msg string, code int, err error) bool {
	s.log(
		"msg", msg,
		"code", code,
		"err", err,
	)
	s.w.Header().Set("content-type", "application/json")
	s.w.WriteHeader(code)
	fmt.Fprintln(s.w, PlatformError{
		Ok:     false,
		Status: code,
		Rid:    s.rid,
		Msg:    msg,
	}.String())
	return false
}

func (s *request) log(kv ...interface{}) {
	logkv(append([]interface{}{
		"t", time.Now().UnixNano(),
		"rid", s.rid,
	}, kv...)...)
}

func (s *request) writebody(data interface{}, mimeType ...string) bool {
	if len(mimeType) != 0 {
		s.w.Header().Set("Content-Type", mimeType[0])
	}
	switch t := data.(type) {
	case io.WriterTo:
		n, err := t.WriteTo(s.w)
		s.wrote, s.err = int(n), err
	case []byte:
		s.wrote, s.err = s.w.Write(t)
	case string:
		s.wrote, s.err = s.w.Write([]byte(t))
	case interface{}:
		data, _ := json.Marshal(t)
		s.wrote, s.err = s.w.Write(data)
	}
	return s.ok()
}

func (s *request) UnmarshalJSON(body interface{}) (ok bool) {
	data := s.Body()
	if !s.ok() {
		// log specific error
		return false
	}
	if s.err = json.Unmarshal(data, body); s.err != nil {
		// log specific error
		return false
	}
	return s.ok()
}

func (s *request) chop() string {
	s.file, s.path = chop(s.path)
	return s.file
}

func chop(p string) (file, next string) {
	p = path.Clean(p)[1:]
	if n := strings.Index(p, "/"); n >= 0 {
		return p[:n], p[n:]
	}
	return p, "/"
}
