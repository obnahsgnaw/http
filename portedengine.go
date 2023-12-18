package http

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/application/pkg/url"
)

type PortedEngine struct {
	e    *gin.Engine
	host url.Host
}

func NewPortedEngine(e *gin.Engine, h url.Host) *PortedEngine {
	return &PortedEngine{
		e:    e,
		host: h,
	}
}

func (s *PortedEngine) Run() error {
	return s.e.Run(s.host.String())
}
func (s *PortedEngine) Engine() *gin.Engine {
	return s.e
}
func (s *PortedEngine) Host() url.Host {
	return s.host
}
