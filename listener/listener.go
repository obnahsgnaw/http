package listener

import (
	"errors"
	"github.com/obnahsgnaw/application/pkg/url"
	"github.com/obnahsgnaw/application/pkg/utils"
	"github.com/soheilhy/cmux"
	"net"
	"strconv"
)

type PortedListener struct {
	l net.Listener
	h url.Host
	m cmux.CMux
}

func New(network string, host url.Host) (*PortedListener, error) {
	if network == "" {
		network = "tcp"
	}
	if host.Ip == "" {
		if ip, err := utils.GetLocalIp(); err != nil || ip == "" {
			return nil, errors.New("ip is required")
		} else {
			host.Ip = ip
		}
	}
	if host.Port <= 0 {
		return nil, errors.New("port is required")
	}
	l, err := net.Listen(network, ":"+strconv.Itoa(host.Port))
	if err != nil {
		return nil, errors.New("listener listen failed, err=" + err.Error())
	}

	return &PortedListener{
		l: l,
		h: host,
		m: cmux.New(l),
	}, nil
}

func Default(host url.Host) (*PortedListener, error) {
	return New("tcp", host)
}

func (s *PortedListener) Network() string {
	return s.l.Addr().Network()
}

func (s *PortedListener) Listener() net.Listener {
	return s.l
}

func (s *PortedListener) Http1Listener() net.Listener {
	return s.m.Match(cmux.HTTP1Fast())
}

func (s *PortedListener) GrpcListener() net.Listener {
	return s.m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
}

func (s *PortedListener) Host() url.Host {
	return s.h
}

func (s *PortedListener) Serve() error {
	return s.m.Serve()
}

func (s *PortedListener) Ip() string {
	return s.h.Ip
}

func (s *PortedListener) Port() int {
	return s.h.Port
}

func (s *PortedListener) Close() {
	_ = s.l.Close()
}
