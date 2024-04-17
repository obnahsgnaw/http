package http

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/http/engine"
	"github.com/obnahsgnaw/http/listener"
)

type Http struct {
	e            *gin.Engine
	l            *listener.PortedListener
	runKey       string // for refuse multi run and close
	initializers []func() error
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

func (s *Http) Serve() error {
	return s.l.Serve()
}

func (s *Http) ServeWithKey(key string) error {
	return s.l.ServeWithKey(key)
}

func (s *Http) Run() error {
	for _, fn := range s.initializers {
		if err := fn(); err != nil {
			return err
		}
	}
	return s.e.RunListener(s.l.HttpListener())
}

func (s *Http) RunWithKey(key string) error {
	if s.runKey != "" {
		return nil
	}
	err := s.Run()
	if err == nil {
		s.runKey = key
	}
	return err
}

func (s *Http) RunAndServ() (err error) {
	go func() {
		err = s.Run()
	}()
	err = s.Serve()
	return
}

func (s *Http) RunAndServWithKey(key string, cb func(error)) {
	go func() {
		if err := s.RunWithKey(key); err != nil {
			cb(err)
		}
	}()
	if err := s.ServeWithKey(key); err != nil {
		cb(err)
	}
	return
}

func (s *Http) Close() {
	s.l.Close()
}

func (s *Http) CloseWithKey(key string) {
	if key != s.runKey {
		return
	}
	s.l.CloseWithKey(key)
	s.runKey = ""
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

func (s *Http) AddInitializer(initializer func() error) {
	s.initializers = append(s.initializers, initializer)
}
