package json_prime

import (
	"fmt"
	"net"

	"github.com/andy1li/protohackers/internal"
)

func NewJSONPrimeServer() *internal.Server {
	return internal.NewServer(handleJSONPrime, "0.0.0.0", 8001)
}

func handleJSONPrime(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	defer fmt.Printf("👋 Connection from %s closed\n", remoteAddr)

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf("📦 Received from %s:\n%q\n", conn.RemoteAddr().String(), buffer[:n])

		response := "{\"method\":\"isPrime\",\"prime\":false}"
		_, err = conn.Write([]byte(response))
		if err != nil {
			break
		}
	}
}
