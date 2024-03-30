package http

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/http/engine"
	"github.com/obnahsgnaw/http/listener"
)

type Http struct {
	e *gin.Engine
	l *listener.PortedListener
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

func (s *Http) Run() error {
	return s.e.RunListener(s.l.HttpListener())
}

func (s *Http) RunAndServ() (err error) {
	go func() {
		err = s.Run()
	}()
	err = s.Serve()
	return
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
