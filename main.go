package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	echo "github.com/andy1li/protohackers/p00-echo"
)

func main() {
	echoServer := echo.NewEchoServer()
	go func() {
		if err := echoServer.Start(); err != nil {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("✅ Echo server started. Press Ctrl+C to stop.")
	<-sigChan
	fmt.Println("\n👋 Shutting down...")
}
