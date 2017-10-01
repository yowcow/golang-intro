package socket

import (
	"net"
)

type Server struct {
	listener net.Listener
}

func NewServer(proto, addr string) (*Server, error) {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		return nil, err
	}
	return &Server{listener}, nil
}

func (s Server) Accept() (*Conn, error) {
	c, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}
	return NewConn(c), nil
}

func (s Server) Close() error {
	return s.listener.Close()
}
