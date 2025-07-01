package echo

import (
	"io"
	"net"
	"testing"
)

func TestEchoServer(t *testing.T) {
	server := NewEchoServer()

	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("Failed to accept: %v", err)
			return
		}
		server.Handler(conn)
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	testData := "test data"
	_, err = conn.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Failed to write: %v", err)
	}

	conn.(*net.TCPConn).CloseWrite()

	response, err := io.ReadAll(conn)
	if err != nil {
		t.Fatalf("Failed to read: %v", err)
	}

	if string(response) != testData {
		t.Errorf("Expected '%s', got '%s'", testData, string(response))
	}
}
