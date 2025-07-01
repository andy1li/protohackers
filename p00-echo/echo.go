package echo

import (
	"fmt"
	"io"
	"net"

	"github.com/andy1li/protohackers/internal"
)

func NewEchoServer() *internal.Server {
	return internal.NewServer(handleEcho, "0.0.0.0", 7)
}

func handleEcho(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("🤝 New connection from %s\n", remoteAddr)

	io.Copy(conn, conn)
	fmt.Printf("👋 Connection from %s closed\n", remoteAddr)
}
