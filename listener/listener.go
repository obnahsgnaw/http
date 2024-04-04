package listener

import (
	"errors"
	"fmt"
	"github.com/soheilhy/cmux"
	"net"
	"strconv"
)

type PortedListener struct {
	m        cmux.CMux
	l        net.Listener
	ip       string
	port     int
	startKey string
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
		return nil, fmt.Errorf("listener listen failed: %w", err)
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
func (s *PortedListener) ServeWithKey(key string) error {
	if s.startKey != "" {
		return nil
	}
	err := s.m.Serve()
	if err == nil {
		s.startKey = key
	}
	return err
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

func (s *PortedListener) CloseWithKey(key string) {
	if key != s.startKey {
		return
	}
	s.m.Close()
	_ = s.l.Close()
	s.startKey = ""
}
