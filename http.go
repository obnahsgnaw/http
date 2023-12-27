package http

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/application/pkg/url"
	"github.com/obnahsgnaw/http/engine"
	"github.com/obnahsgnaw/http/listener"
)

type Http struct {
	e          *gin.Engine
	l          *listener.PortedListener
	reusedPort bool
}

func New(e *gin.Engine, l *listener.PortedListener) *Http {
	return &Http{
		e: e,
		l: l,
	}
}

func Default(host url.Host, config *engine.Config) (*Http, error) {
	var e *gin.Engine
	var l *listener.PortedListener
	var err error

	if e, err = engine.New(config); err != nil {
		return nil, err
	}
	if l, err = listener.Default(host); err != nil {
		return nil, err
	}
	return New(e, l), nil
}

func (s *Http) Serve() error {
	return s.l.Serve()
}

func (s *Http) Run() error {
	if s.reusedPort {
		return s.e.RunListener(s.l.Http1Listener())
	}
	return s.e.RunListener(s.l.Listener())
}

func (s *Http) RunAndServ() (err error) {
	if s.reusedPort {
		go func() {
			err = s.Run()
		}()
		err = s.Serve()
	} else {
		err = s.Run()
	}
	return
}

func (s *Http) Engine() *gin.Engine {
	return s.e
}

func (s *Http) Listener() *listener.PortedListener {
	return s.l
}

func (s *Http) Host() url.Host {
	return s.l.Host()
}

func (s *Http) Ip() string {
	return s.l.Host().Ip
}

func (s *Http) Port() int {
	return s.l.Host().Port
}

func (s *Http) ReusedPort() {
	s.reusedPort = true
}
