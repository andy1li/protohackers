package internal

import (
	"fmt"
	"net"
)

type Server struct {
	Handler func(conn net.Conn)
	host    string
	port    int
}

func NewServer(handler func(conn net.Conn), host string, port int) *Server {
	return &Server{
		Handler: handler,
		host:    host,
		port:    port,
	}
}

func (s *Server) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.GetAddress())
	if err != nil {
		return fmt.Errorf("❌ Failed to start server on %s: %w", s.GetAddress(), err)
	}
	defer listener.Close()

	fmt.Printf("👂 Echo server listening on %s\n", s.GetAddress())

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("❌ Failed to accept connection: %w", err)
		}

		go s.Handler(conn)
	}
}
