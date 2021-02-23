package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	transcoding "github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/cbsinteractive/transcode-orchestrator/service/exceptions"
	"github.com/sirupsen/logrus"
	"github.com/zsiec/pkg/tracing"
)

var ErrProvider = errors.New("provider error")
var ErrStorage = errors.New("storage error")

type Server struct {
	Config      *config.Config
	DB          *job.Client
	logger      *logrus.Logger
	errReporter exceptions.Reporter
	tracer      tracing.Tracer

	request
}

func (s Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.request = newRequest(rw, r)
	s.serve()
	defer s.request.finalize()
}

func (s *Server) serve() bool {
	switch s.chop() {
	case "jobs":
		job := &job.Job{ID: s.chop()}
		switch s.method() {
		case "POST":
			if !s.request.UnmarshalJSON(job) {
				return false
			}
			stat, err := s.putJob0(job)
			if err != nil {
				return s.writeerror("put job failed", 400, err)
			}
			return s.writebody(stat)
		case "GET":
			stat, err := s.getJob0(job, false)
			if err != nil {
				return s.writeerror("get job failed", 400, err)
			}
			return s.writebody(stat)
		case "DELETE":
			stat, err := s.getJob0(job, false)
			if err != nil {
				return s.writeerror("del job failed", 400, err)
			}
			return s.writebody(stat)
		}
	case "providers":
	default:
		s.writeerror("bad request path", 400, nil)
	}
	return false
}

func (s *Server) provider0(job *job.Job) (transcoding.Provider, error) {
	fn, err := transcoding.GetFactory(job.ProviderName)
	if err != nil {
		return nil, err
	}
	return fn(s.Config)
}

func (s *Server) putJob0(job *job.Job) (*job.Status, error) {
	p, err := s.provider0(job)
	if err != nil {
		return nil, err
	}
	stat, err := p.Create(s.request.ctx, job)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrProvider, err)
	}
	stat.ID = job.ID
	job.ProviderJobID = stat.ProviderJobID
	job.CreationTime = time.Now()
	if err = s.DB.Put(job.ID, &job); err != nil {
		return stat, fmt.Errorf("%w: %v", ErrStorage, err)
	}
	return stat, nil
}

func (s *Server) getJob0(job *job.Job, del bool) (*job.Status, error) {
	if err := s.DB.Get(job.ID, &job); err != nil {
		return nil, err
	}
	p, err := s.provider0(job)
	if err != nil {
		return nil, err
	}
	if del {
		if err = p.Cancel(s.request.ctx, job.ProviderJobID); err != nil {
			return nil, err
		}
	}
	//TODO(as): provider name
	return p.Status(s.request.ctx, job)
}

func (s *Server) method() string {
	return s.request.r.Method
}

// PlatformError implements a well-known error response for http clients
// encountering an error when using the service.
type PlatformError struct {
	Ok     bool   `json:"ok"`
	Status int    `json:"status"`
	Rid    uint64 `json:"rid"`
	Msg    string `json:"msg,omitempty"`
}

// String returns the json-formatted platform response
func (p PlatformError) String() string {
	data, _ := json.Marshal(p)
	return string(data)
}

type StatusError struct {
	Code int
	Msg  string
	body string
}

func (e StatusError) NotFound() bool {
	return e.Code == 404
}
func (e StatusError) Error() string {
	return fmt.Sprintf("http status: %d: %q", e.Code, e.body)
}

func logkv(kv ...interface{}) bool {
	msg := `{`
	sep := " "
	for i := 0; i+1 < len(kv); i += 2 {
		v := kv[i+1]
		if v == nil {
			v = ""
		} else {
			switch v.(type) {
			case fmt.Stringer:
				v = fmt.Sprint(v)
			case error:
				v = fmt.Sprint(v)
			}
		}
		value, _ := json.Marshal(v)
		msg += fmt.Sprintf(`%s%q:%s`, sep, kv[i], string(value))
		sep = ", "
	}
	msg += `}`
	log.Println(msg)
	return true
}
