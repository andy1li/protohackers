package echo

import (
	"io"
	"net"
	"testing"
)

func TestEchoServer(t *testing.T) {
	conn, cleanup := setupEchoTest(t)
	defer cleanup()

	testData := "test data"
	_, err := conn.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Failed to write: %v", err)
	}

	conn.(*net.TCPConn).CloseWrite()

	response, err := io.ReadAll(conn)
	if err != nil {
		t.Fatalf("Failed to read: %v", err)
	}

	if string(response) != testData {
		t.Errorf(`
🎯 Expected: '%s'
📦 Received: '%s'`, testData, string(response))
	}
}

func setupEchoTest(t *testing.T) (net.Conn, func()) {
	server := NewEchoServer()

	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

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
		listener.Close()
		t.Fatalf("Failed to connect: %v", err)
	}

	cleanup := func() {
		conn.Close()
		listener.Close()
	}

	return conn, cleanup
}
