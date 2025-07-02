package json_prime

import (
	"io"
	"net"
	"testing"
)

func TestJSONPrimeServer(t *testing.T) {
	conn, cleanup := setupJSONPrimeTest(t)
	defer cleanup()

	testRequest := "{\"method\":\"isPrime\",\"number\":123}"
	_, err := conn.Write([]byte(testRequest))
	if err != nil {
		t.Fatalf("Failed to write: %v", err)
	}

	conn.(*net.TCPConn).CloseWrite()

	response, err := io.ReadAll(conn)
	if err != nil {
		t.Fatalf("Failed to read: %v", err)
	}

	expectedResponse := "{\"method\":\"isPrime\",\"prime\":false}"
	if string(response) != expectedResponse {
		t.Errorf(`
🎯 Expected: '%s'
📦 Received: '%s'`, expectedResponse, string(response))
	}
}

func setupJSONPrimeTest(t *testing.T) (net.Conn, func()) {
	server := NewJSONPrimeServer()

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
