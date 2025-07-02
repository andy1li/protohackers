package echo

import (
	"fmt"
	"io"
	"net"

	"github.com/andy1li/protohackers/internal"
)

func NewEchoServer() *internal.Server {
	return internal.NewServer(handleEcho, "0.0.0.0", 8000)
}

func handleEcho(conn net.Conn) {
	defer conn.Close()
	defer fmt.Printf("👋 Connection from %s closed\n", conn.RemoteAddr().String())

	io.Copy(conn, conn)
}
