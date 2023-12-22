package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/application/pkg/url"
	"github.com/obnahsgnaw/application/pkg/utils"
	"net"
	"strconv"
)

type PortedEngine struct {
	e *gin.Engine
	l net.Listener
	h url.Host
}

func NewPortedEngine(e *gin.Engine, h url.Host) *PortedEngine {
	return &PortedEngine{
		e: e,
		h: h,
	}
}

func NewListenerEngine(e *gin.Engine, h url.Host, l net.Listener) *PortedEngine {
	return &PortedEngine{
		e: e,
		l: l,
		h: h,
	}
}

func NewListener(h url.Host) (net.Listener, error) {
	return net.Listen("tcp", ":"+strconv.Itoa(h.Port))
}

func (s *PortedEngine) Run() error {
	if s.h.Ip == "" {
		if ip, err := utils.GetLocalIp(); err != nil {
			return errors.New("ip invalid")
		} else {
			s.h.Ip = ip
		}
	}
	if s.l != nil {
		return s.e.RunListener(s.l)
	}
	return s.e.Run(s.h.String())
}

func (s *PortedEngine) Engine() *gin.Engine {
	return s.e
}

func (s *PortedEngine) Host() url.Host {
	return s.h
}

func (s *PortedEngine) Listener() net.Listener {
	return s.l
}
