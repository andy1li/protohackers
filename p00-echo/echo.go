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
		host: "0.0.0.0",
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

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("❌ Error handling connection from %s: %v\n", remoteAddr, err)
			continue
		}

		fmt.Printf("📦 Received from %s:\n%q\n", remoteAddr, buffer[:n])

		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Printf("❌ Error writing response to %s: %v\n", remoteAddr, err)
		}
	}

	fmt.Printf("👋 Connection from %s closed\n", remoteAddr)
}

func (s *Server) GetAddress() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}
	return fmt.Sprintf("%s:%d", s.host, s.port)
}
