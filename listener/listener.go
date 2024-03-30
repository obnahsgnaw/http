package listener

import (
	"errors"
	"github.com/soheilhy/cmux"
	"net"
	"strconv"
)

type PortedListener struct {
	m    cmux.CMux
	l    net.Listener
	ip   string
	port int
}

func New(network string, ip string, port int) (*PortedListener, error) {
	if network == "" {
		network = "tcp"
	}
	if ip == "" {
		return nil, errors.New("ip is required")
	}
	if port <= 0 {
		return nil, errors.New("port is required")
	}
	l, err := net.Listen(network, ":"+strconv.Itoa(port))
	if err != nil {
		return nil, errors.New("listener listen failed, err=" + err.Error())
	}

	return &PortedListener{
		m:    cmux.New(l),
		l:    l,
		ip:   ip,
		port: port,
	}, nil
}

func Default(ip string, port int) (*PortedListener, error) {
	return New("tcp", ip, port)
}

func (s *PortedListener) Network() string {
	return s.l.Addr().Network()
}

func (s *PortedListener) RawListener() net.Listener {
	return s.l
}

func (s *PortedListener) HttpListener() net.Listener {
	return s.m.Match(cmux.HTTP1Fast())
}

func (s *PortedListener) GrpcListener() net.Listener {
	return s.m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
}

func (s *PortedListener) Serve() error {
	return s.m.Serve()
}

func (s *PortedListener) Ip() string {
	return s.ip
}

func (s *PortedListener) Port() int {
	return s.port
}

func (s *PortedListener) Host() string {
	return s.Ip() + ":" + strconv.Itoa(s.Port())
}

func (s *PortedListener) Close() {
	_ = s.l.Close()
}
