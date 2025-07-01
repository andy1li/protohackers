package echo

import (
	"io"
	"net"
	"testing"
)

func TestHandleConnection(t *testing.T) {
	server := NewEchoServer()

	// Create a listener
	listener, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	// Accept connection in goroutine and handle it
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("Failed to accept: %v", err)
			return
		}
		server.handleConnection(conn)
	}()

	// Connect client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Send data
	testData := "test data"
	_, err = conn.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Failed to write: %v", err)
	}

	// Close write side
	conn.(*net.TCPConn).CloseWrite()

	// Read response
	response, err := io.ReadAll(conn)
	if err != nil {
		t.Fatalf("Failed to read: %v", err)
	}

	if string(response) != testData {
		t.Errorf("Expected '%s', got '%s'", testData, string(response))
	}
}
