package json_prime

import (
	"encoding/json"
	"fmt"
	"math"
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

		var request IsPrimeRequest
		err = json.Unmarshal(buffer[:n], &request)
		if err != nil {
			break
		}

		fmt.Printf("🔢 Parsed as %+v\n", request)

		if request.Method != "isPrime" {
			break
		}

		prime := isPrime(request.Number)

		response, err := json.Marshal(IsPrimeResponse{
			Method: "isPrime",
			Prime:  prime,
		})
		if err != nil {
			break
		}
		_, err = conn.Write([]byte(response))
		if err != nil {
			break
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
