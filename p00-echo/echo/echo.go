package echo

import (
	"fmt"
	"io"
	"net"
)

// Server represents a TCP echo server
type Server struct {
	host     string
	port     int
	listener net.Listener
}

// New creates a new echo server
func NewEchoServer() *Server {
	return &Server{
		host: "127.0.0.1",
		port: 7,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.GetAddress())
	if err != nil {
		return fmt.Errorf("❌ Failed to start server on %s: %w", s.GetAddress(), err)
	}
	s.listener = listener
	defer listener.Close()

	fmt.Printf("👂 Echo server listening on %s\n", s.GetAddress())

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("❌ Failed to accept connection: %w", err)
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("🤝 New connection from %s\n", remoteAddr)

	_, err := io.Copy(conn, conn)
	if err != nil {
		fmt.Printf("❌ Error handling connection from %s: %v\n", remoteAddr, err)
	}

	fmt.Printf("👋 Connection from %s closed\n", remoteAddr)
}

func (s *Server) GetAddress() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}
	return fmt.Sprintf("%s:%d", s.host, s.port)
}
