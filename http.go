package http

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/http/engine"
	"github.com/obnahsgnaw/http/listener"
)

type Http struct {
	e      *gin.Engine
	l      *listener.PortedListener
	runKey string
}

func New(e *gin.Engine, l *listener.PortedListener) *Http {
	return &Http{
		e: e,
		l: l,
	}
}

func Default(ip string, port int, config *engine.Config) (*Http, error) {
	var e *gin.Engine
	var l *listener.PortedListener
	var err error

	if e, err = engine.New(config); err != nil {
		return nil, err
	}
	if l, err = listener.Default(ip, port); err != nil {
		return nil, err
	}
	return New(e, l), nil
}

func (s *Http) Serve(key string) error {
	return s.l.Serve(key)
}

func (s *Http) Run(key string) error {
	if s.runKey != "" {
		return nil
	}
	err := s.e.RunListener(s.l.HttpListener())
	if err == nil {
		s.runKey = key
	}
	return err
}

func (s *Http) RunAndServ(key string, cb func(error)) {
	go func() {
		if err := s.Run(key); err != nil {
			cb(err)
		}
	}()
	if err := s.Serve(key); err != nil {
		cb(err)
	}
	return
}

func (s *Http) Close(key string) {
	if key != s.runKey {
		return
	}
	s.l.Close(key)
}

func (s *Http) Engine() *gin.Engine {
	return s.e
}

func (s *Http) Listener() *listener.PortedListener {
	return s.l
}

func (s *Http) Ip() string {
	return s.l.Ip()
}

func (s *Http) Port() int {
	return s.l.Port()
}

func (s *Http) Host() string {
	return s.l.Host()
}
